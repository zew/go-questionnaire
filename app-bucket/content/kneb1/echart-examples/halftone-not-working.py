# pip install halftone


import PIL.Image
import os

# import halftone as ht

from typing import Tuple, Union, Callable

import numpy as np
import PIL.Image

ImgType = Union[PIL.Image.Image, np.ndarray]


def halftone(img: ImgType, spot_fn: Callable[[int, int], float]) -> ImgType:
    is_pil = isinstance(img, PIL.Image.Image)
    if is_pil:
        img = array_from_pil(img)
    else:
        print("error no PIL")
        quit()

    # print("evaluating dims %s func %s" % (img.shape, spot_fn))
    halftoned = img > evaluate_2d_func(img.shape, spot_fn)
    if is_pil:
        halftoned = pil_from_array(halftoned)
    return halftoned


def euclid_dot(spacing: float, angle: float = 0):
    """Circular dot changing to square at 50% grey."""
    pixel_divisor = 2.0 / spacing

    def fn(x: int, y: int):
        x, y = rotate(x * pixel_divisor, y * pixel_divisor, angle)
        # to make an ellipse multiply sin/cos
        # arguments with different multipliers
        return (0.5 - (0.25 * (
                np.sin(np.pi * (x + 0.5)) +
                np.cos(np.pi * y))))

    return fn


def square_dot(spacing: float, angle: float = 0):
    def fn(x, y):
        x, y = rotate(x, y, angle)
        coords = np.array([x, y])
        x, y = np.abs(((np.abs(coords) / (spacing * 0.5)) % 2) - 1)
        return 1 - (0.5 * (x + y))

    return fn


def circle_dot(spacing: float, angle: float = 0):
    def fn(x, y):
        x, y = rotate(x, y, angle)
        coords = np.array([x, y])
        coords = coords / spacing
        coords = coords % 1
        coords = coords * 2 - 1
        x, y = coords
        return 1 - np.sqrt((x * x) + (y * y)) / np.sqrt(2)

    return fn


def triangle_dot(spacing: float, angle: float = 0):
    inv_spacing = 1 / spacing

    def fn(x, y):
        x, y = rotate(x, y, angle)
        coords = np.array([x, y])
        coords = coords * inv_spacing
        coords = coords % 1
        x, y = coords
        return (x + y) * 0.5

    return fn


def line(spacing: float, angle: float = 0):
    angle %= np.pi * 0.5
    inv_spacing = 1 / spacing
    sin_inv_spacing = np.sin(angle) * inv_spacing
    cos_inv_spacing = np.cos(angle) * inv_spacing

    def fn(x, y):
        value = (sin_inv_spacing * x) + (cos_inv_spacing * y)
        return np.abs(np.sign(value) * value % np.sign(value))

    return fn


def evaluate_2d_func(img_shape, fn):
    w, h, depth = img_shape
    xaxis = np.arange(w)
    yaxis = np.arange(h)
    print("   ")
    print("   ")
    print("   callable %s with w%d - h%d" % (fn, w,h))
    return fn(xaxis[:, None], yaxis[None, :])


def rotate(x: float, y: float, angle: float) -> Tuple[float, float]:
    """
    Rotate coordinates (x, y) by given angle.

    angle: Rotation angle in degrees
    """
    angle_rad = angle / 360 * 2 * np.pi
    sin, cos = np.sin(angle_rad), np.cos(angle_rad)
    return x * cos - y * sin, x * sin + y * cos


def transform_grid(spacing, angle):
    def fn(x, y):
        x, y = rotate(x, y, angle)
        coords = np.array([x, y])
        coords = coords / spacing
        coords = coords % 1
        return coords

    return fn


def pil_from_array(arr):
    return PIL.Image.fromarray((arr * 255).astype('uint8'))


def array_from_pil(img):
    return np.array(img) / 255


spot_functions = {
    'euclid': euclid_dot,
    'square': square_dot,
    'circle': circle_dot,
    'triangle': triangle_dot,
    'line': line,
}


fileNames = [
    'director-schnabel.png',
    'heinemann.png',
]

print("start")








for idx,fn in enumerate(fileNames):
    print("%d - %s" %  (idx,fn))
    fullFN = os.path.join('./img',fn)
    img = PIL.Image.open(fullFN)
    htImg = halftone(img, euclid_dot(spacing=8, angle=30))
    fullFNH = fullFN.replace(".png", "_ht.png")
    htImg.save(fullFNH)

print("fin")
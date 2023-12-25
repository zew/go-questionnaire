from PIL import Image
import os.path, sys

inPath = "C:\\Users\\xie\\Desktop\\tiff\\Bmp"

inPath = ".\\tmp-raw"
outPath = ".\\out"

def makeWhiteTransparent(filePath):
    img = Image.open(filePath)
    img = img.convert("RGBA")
    datas = img.getdata()
    newData = []
    for item in datas:
        if item[0] == 255 and item[1] == 255 and item[2] == 255:
            newData.append((255, 255, 255, 0))
        else:
            newData.append(item)
    img.putdata(newData)
    img.save(filePath, "PNG")    


def crop(subDir):
    outDir = os.path.join(outPath,subDir)
    # os.mkdir( outDir )
    os.makedirs(outDir, exist_ok=True)
    inDir = os.path.join(inPath,subDir)
    print("in dir is %s" % inDir)
    fileNames = os.listdir( inDir )
    for idx, fn in enumerate(fileNames):
        fullPath = os.path.join(inDir,fn)
        print("  file is %s" % fullPath)
        if os.path.isfile(fullPath):
            im = Image.open(fullPath)
            # 1920x1080
            # 1280x720
            cx = 128 + 48
            xD =  24 +  0    # crop more left than right
            cy = 14
            imCrop = im.crop((cx + xD, cy, 1280 - (cx - xD), 720-cy))
            # save
            noExt, oldExt = os.path.splitext(fullPath)
            baseName = os.path.basename(noExt)
            newFn = "%s_cropped.png" % baseName
            newFn = "%02d.png" % idx
            imCrop.save(  os.path.join( outDir, newFn ), "PNG", quality=100)
            makeWhiteTransparent( os.path.join( outDir, newFn ) )


crop( "fin" )
crop( "umw" )
import os
import halftone

fileNames = [
    'woman.jpg',
    'director-schnabel.png',
    'heinemann.png',
]

print("start")

for idx,fn in enumerate(fileNames):
    if fn.startswith("out-"):
        continue
    pureFn, ext = os.path.splitext(fn)    
    fullFN = os.path.join('.','img',fn)
    print("%d - %s%s - %s" %  (idx, pureFn, ext, fullFN))
    # halftone.halftone(fullFN, fg_color = (55,44,44))    
    
    # side does not make any difference
    # halftone.halftone(fullFN, side = 40, bg_color = (255,255,255), fg_color = (55,44,44), alpha = 1.8 )    
    
    #  
    # halftone.halftone(fullFN, bg_color = (244,244,244), fg_color = (55,44,44), alpha = 1.6 )    

    # keep jump derived from resolution 
    # halftone.halftone(fullFN, jump = 6, bg_color = (244,244,244), fg_color = (55,44,44), alpha = 1.6 )    
    
    halftone.halftone(fullFN, bg_color = (244,244,244), fg_color = (44,44,44), alpha = 1.6 )    

print("fin")
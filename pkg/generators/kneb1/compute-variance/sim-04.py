import matplotlib.pyplot as plt

from numpy.random import seed
# seed(1) # reproducible


from numpy.random import normal

import numpy as np
import random

from   scipy import stats




desMn  = 1.059 # desired mean
desStd = 0.1462  # desired standard deviation

smplSize = 200
smplSize = 20*1000*1000

loops = 10*1000
loops = 100*1000
loops = 1000*1000

mnsSum  = 0
stdsSum = 0
wrSum   = 0

wrss = []


for i1 in range(0,loops):

    if i1 % (loops/10) == 0:
        # creating a synthetic MSCI world distribution
        mciworld = normal(loc=desMn, scale=desStd, size=smplSize)
        nmLp  = np.mean(mciworld)                            
        stdLp = np.std( mciworld, ddof=1)  #  ddof - Delta Degrees of Freedom - 0 or 1 - default 0
        print("mciworld synth dist %5d: %5.4f - %5.4f" % (i1, nmLp, stdLp) )


    # T = 10
    # T = 15
    T = 20
    s = 1200 # savings
    ws = 0     # wealth safe  asset
    wr = 0     # wealth risky asset
    wrs = 0    # wealth risky asset - stochastic
    for i2 in range(T+1): 
        draw = random.randint(0,smplSize-1)
        # print("     random draw %d from %d: %5.4f" % ( draw, smplSize, mciworld[r]))
        # futureValue(mciworld[draw])
        ws = ws*1.01  + s
        wr = ws*1.059 + s

        r = mciworld[draw]
        wrs = wrs*r + s
        # print("    t = %2d r = %5.4f  ws = %7.1f  wr = %7.1f wrs = %7.1f" % (i2,r,ws,wr,wrs))


        # count, bins, ignored = plt.hist(mciworld, bins='auto', histtype='stepfilled', alpha=0.8, density=True)
        # plt.show()

    wrSum += wrs
    wrss.append(wrs) 
    # mnsSum  += mn
    # stdsSum += std

    if i1 % (loops/10) == 0 and i1 > 0  :
        mn  = np.mean(wrss)                            
        std = np.std( wrss, ddof=1)  
        print("wealth risky mn and std - %5d: %7.1f - %6.5f" % (i1,mn, std/mn) )


    # print("end of loop %d" % i1)






nmLp  = np.mean(wrss)                            
stdLp = np.std( wrss, ddof=1)  
print("wealth risky", wrss[:4], "...")
print("wealth risky mn and std: %7.1f - %6.5f" % (nmLp, stdLp/nmLp) )


# 1.000.000

# s=400
# wealth risky mn and std:189726.0 - 0.44289

# s=200
# wealth risky mn and std: 94843.3 - 0.44189

# s=50
# wealth risky mn and std: 23729.7 - 0.44286

# s=25
# wealth risky mn and std: 11859.2 - 0.44280


# T = 20
# s = 100
# wealth risky mn and std: 47476.8 - 0.44302
# wealth risky mn and std: 47449.5 - 0.44280


# T = 15
# s = 100
# wealth risky mn and std: 30543.4 - 0.36453

# T = 10
# s = 100
# wealth risky mn and std: 17867.6 - 0.28114

# T = 5
# s = 100
# wealth risky mn and std:  8346.7 - 0.18650



import matplotlib.pyplot as plt

from numpy.random import seed
# seed(1) # reproducible

import numpy as np
import random

from   scipy import stats


desMn  = 1.059 # desired mean
desStd = 0.15  # desired standard deviation

smplSize = 200
smplSize = 200*100

loops = 1000*50

mns  = []
mnsSum  = 0
stds = []
stdsSum = 0

for i in range(0,loops):
    mciworld = stats.norm.rvs(loc=desMn, scale=desStd, size=smplSize) # random variates

    # print(mciworld[0:5])
    # print(mciworld[:])

    mn  = np.mean(mciworld)                            
    std = np.std( mciworld, ddof=1)  #  ddof - Delta Degrees of Freedom - 0 or 1 - default 0

    if i % (loops/10) == 0: 
        print("    dist %5d: %5.4f - %5.4f" % (i, mn, std) )
        # draw
        rn = random.randint(0,smplSize-1)
        print("     random draw %d from %d: %5.4f" % ( rn, smplSize, mciworld[rn]))

        # bins='auto'
        count, bins, ignored = plt.hist(mciworld, bins=30, histtype='stepfilled', alpha=0.8, density=True)
        plt.show()


    mns.append(mn)
    mnsSum += mn
    stds.append(std)
    stdsSum += std





print("==========================" )
print("%5.4f - %5.4f" % (mnsSum/loops, stdsSum/loops) )


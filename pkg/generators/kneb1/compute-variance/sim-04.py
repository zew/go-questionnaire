'''
    The MCI world stock index provides historical total returns for the stocks contained.
    These total returns are observed daily and then computed into annualized total returns
    and also into a number for annualized standard deviation of total returns of 0.1462.

    The total returns are understood to be the result of billions of random variables
    from within the contained private companies. Company 231 gets a better CEO. Its input prices rise. 
    Its technology improves. A factory burns down. 
    The historical total returns are also affected by global secular trends. For the past
    thirty years these include: China coming into the global market and adding 300 million 
    workers to the global economy. It also includes an unprecedented increase of female labor
    force participation rate, the first wave of digitalization (personal computers), favorable demographics (boomer cohort)
    and some more. 
    For the next thirty years, these global secular trends will definitely shift. 
    The female labor force rate cannot rise much more. China will not grow as before. 
    The green transformation and artificial intelligence might cause even more growth
    than the previous factors, or they might not.
    So whereas the many "small" variables of the MCI World companies will not change their mean nor standard deviation much,
    the secular macro variables might shift in unpredictable ways, which cannot be captured stochastically
    from past observations. 

    The past mean of 1.059 and the standard deviation of 0.1462, on which we base our projections for risky assets outcomes,
    must therefore be taken with two grains of salt.

    That said, we need to transform the yearly standard deviation of the MCI world 
    into the standard deviation for a 20-year-annuity savings scheme.
    This scheme comprises summation and multiplication of stochastic variables.
    Summation, because another payment is added to the savings account every year.
    Multiplication, because the previous savings are compounded with the stochastic real return
    each period (year).

    Since the computation of the overall standard deviation of our savings scheme is impossible 
    with _analytical_  methods and might even change [depending on the mean](https://en.wikipedia.org/wiki/Distribution_of_the_product_of_two_random_variables#Variance_of_the_product_of_independent_random_variables),
    we simply simulate ten million instances of the entire scheme with random draws
    from a smooth synthetic MCI world distribution.
    We save all outcomes and compute the standard deviation of these outcomes. 

    Details and results in the code below

'''

from numpy.random import seed

# normal distribution
# provides densitity func, cumulativ density func and func for random sampling 
from numpy.random import normal

# for "drawing" an observation from a synthetic sample
from random import randint

import numpy as np


# image of the generated distribution
# for debugging
import matplotlib.pyplot as plt


# In order to make a run of this simulation reproducible,
# we could set a seed for random number generation
if False:
    seed(1) 


# Historical annualized total return of MCIWorld stock index is 1.059 for Euro-Zone investments.
# Historical annualized stddev of MCIWorld stock index is 0.1462.
desMn  = 1.059   # desired mean - equal to MCI world historical mean
desStd = 0.1462  # desired standard deviation - equal to MCI world historical stddev


# size of the synthetic distribution
#   for a fine grained synthetic distribution to draw from 
smplSize = 100*1000*1000

# 10 million loops
loops = 10*1000*1000

wrSum  = 0 # store the sum of the risky asset future value for every loop
wrss = [] # achieved future values for for the risky asset from stochastic total returns


# each loops represents one 20-year annuity saving
for i1 in range(0,loops):

    # creating a synthetic distribution of the MCI world.
    # we create a new distribution from time to time
    # even though we should have perfect randomness without
    if i1 % (loops/10) == 0:
        mciworld = normal(loc=desMn, scale=desStd, size=smplSize)
        nmLp  = np.mean(mciworld)                            
        stdLp = np.std( mciworld, ddof=1)  #  ddof - Delta Degrees of Freedom - 0 or 1 - default 0
        print("mciworld synth dist %5d: %5.4f - %5.4f" % (i1, nmLp, stdLp) )


    # we can vary the number of years 
    # to get an idea how the standard deviation increases with longer savings periods
    # T = 10
    T = 20

    s = 1200 # savings per year
    ws =  0    # wealth if investing 100% into safe  asset - for debugging and comparison
    wr =  0    # wealth if investing 100% into risky asset - mean total returns - for debugging and comparison
    wrs = 0    # wealth if investing 100% into risky asset - stochastic total returns

    # saving over twenty years...
    for i2 in range(T+1): 

        ws = ws*1.01  + s
        wr = ws*1.059 + s

        draw = randint(0,smplSize-1)
        # print("     random draw %d from %d: %5.4f" % ( draw, smplSize, mciworld[r]))
        # futureValue(mciworld[draw])
        r = mciworld[draw]
        wrs = wrs*r + s
        # print("    t = %2d r = %5.4f  ws = %7.1f  wr = %7.1f wrs = %7.1f" % (i2,r,ws,wr,wrs))

        if False:
            # visualize for debugging
            count, bins, ignored = plt.hist(mciworld, bins='auto', histtype='stepfilled', alpha=0.8, density=True)
            plt.show()

    wrSum += wrs     # storing results of looop
    wrss.append(wrs) 


    if i1 % (loops/10) == 0 and i1 > 0  :
        # from time to time print progress to the command line
        mn  = np.mean(wrss)                    
        std = np.std( wrss, ddof=1)  
        print("wealth risky mn and std - %5d: %7.1f - %6.5f" % (i1,mn, std/mn) )


    # print("end of loop %d" % i1)




# final analysis of the simulated savings accounts
nmLp  = np.mean(wrss)                            
stdLp = np.std( wrss, ddof=1)  
print("wealth risky", wrss[:4], "...") # look at some concrete FVs
print("wealth risky mn and std: %7.1f - %6.5f" % (nmLp, stdLp/nmLp) )

'''
results from execution on a 2023 Thinkpad T14S2 computer with Intel CPU and 16GB RAM,
  though this should not matter to the results in any way.

1.)
Changing the yearly contribution s
did _not_ affect the stddev in any way

Some results for changes in s.
  Only 200.000 loops => thus more variation

  s=400
  wealth risky mn and std:189726.0 - 0.44289

  s=200
  wealth risky mn and std: 94843.3 - 0.44189

  s=50
  wealth risky mn and std: 23729.7 - 0.44286

  s=25
  wealth risky mn and std: 11859.2 - 0.44280

2.) Results for the standard deviation of the "risky" savings.

  T = 20
  s = 100
  wealth risky mn and std: 47447.5 - 0.42993
  wealth risky mn and std: 47446.5 - 0.42999


  T = 10
  s = 100
  wealth risky mn and std: 17876.0 - 0.27449
  

'''

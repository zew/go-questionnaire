import random



for idx in range(0,100):
    rn = random.uniform(0,100)
    if idx % 10 == 0:
        print("%03d - %6.4f" % (idx, rn))


# 
# https://docs.scipy.org/doc/scipy/reference/generated/scipy.stats.norm.html

import numpy as np
from   scipy.stats import norm
import matplotlib.pyplot as plt

print("imports complete")

fig, ax = plt.subplots(1, 1)


# compute 'moments'

# loc     - mean
# scale   - std dev
# moments - standard deviation, kurtosis, and mean
#


mean, var, skew, kurt = norm.stats(moments='mvsk')

print("  mean %f, var %f, skew %f, kurt %f" % (mean, var, skew, kurt))

# ppf - percent point function
x = np.linspace(  norm.ppf(0.01),  norm.ppf(0.99),  100 )
print("norm dist func vals f(0) %4.3f f(49) %4.3f f(50) %4.3f f(99) %4.3f" % (x[0], x[49], x[50], x[99])  )



# freeze the probability density function  - pdf - by plotting it
rv = norm()
ax.plot(x, rv.pdf(x), 'k-', lw=2, label='frozen pdf')

# check - by using the cumulative distribution function - cdf
checks = [0.001, 0.5, 0.999]
vals = norm.ppf(checks)
isRoughlyEqual = np.allclose(checks, norm.cdf(vals))
print(checks, " functioned and inverted %r" % isRoughlyEqual )



# generate random results
#   rvs - random variates
r = norm.rvs(size=1000)
ax.hist(r, density=True, bins='auto', histtype='stepfilled', alpha=0.2)

ax.set_xlim([x[0], x[-1]])
ax.legend(loc='best', frameon=False)

plt.show()

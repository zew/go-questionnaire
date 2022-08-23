// github.com/chen0040/js-stats

var jsstats = jsstats || {};

(function (jss) {

    var NormalDistribution = function (mean, sd) {
        if (!mean) {
            mean = 0.0;
        }
        if (!sd) {
            sd = 1.0;
        }
        this.mean = mean;
        this.sd = sd;
        this.Sqrt2 = 1.4142135623730950488016887;
        this.Sqrt2PI = 2.50662827463100050242E0;
        this.lnconstant = -Math.log(this.Sqrt2PI * sd);
    };

    NormalDistribution.prototype.sample = function () {

    };

    NormalDistribution.prototype.cumulativeProbability = function (x) {
        var z = (x - this.mean) / (this.Sqrt2 * this.sd);
        return 0.5 + 0.5 * this.errorFunc(z);
    };

    NormalDistribution.prototype.invCumulativeProbability = function (p) {
        var Z = this.Sqrt2 * this.invErrorFunc(2 * p - 1);
        return Z * this.sd + this.mean;
    };
 
    NormalDistribution.prototype.errorFunc = function (z) {
        var t = 1.0 / (1.0 + 0.5 * Math.abs(z));

        // use Horner's method
        var ans = 1 - t * Math.exp(-z * z - 1.26551223 +
            t * (1.00002368 +
                t * (0.37409196 +
                    t * (0.09678418 +
                        t * (-0.18628806 +
                            t * (0.27886807 +
                                t * (-1.13520398 +
                                    t * (1.48851587 +
                                        t * (-0.82215223 +
                                            t * (0.17087277))))))))));
        if (z >= 0) return ans;
        else return -ans;
    };

    NormalDistribution.prototype.invErrorFunc = function (x) {
        var z;
        var a = 0.147;
        var the_sign_of_x;
        if (0 == x) {
            the_sign_of_x = 0;
        }
        else if (x > 0) {
            the_sign_of_x = 1;
        }
        else {
            the_sign_of_x = -1;
        }

        if (0 != x) {
            var ln_1minus_x_sqrd = Math.log(1 - x * x);
            var ln_1minusxx_by_a = ln_1minus_x_sqrd / a;
            var ln_1minusxx_by_2 = ln_1minus_x_sqrd / 2;
            var ln_etc_by2_plus2 = ln_1minusxx_by_2 + (2 / (Math.PI * a));
            var first_sqrt = Math.sqrt((ln_etc_by2_plus2 * ln_etc_by2_plus2) - ln_1minusxx_by_a);
            var second_sqrt = Math.sqrt(first_sqrt - ln_etc_by2_plus2);
            z = second_sqrt * the_sign_of_x;
        }
        else { // x is zero
            z = 0;
        }
        return z;
    };
        
    jss.NormalDistribution = NormalDistribution;

})(jsstats);

var module = module || {};
if (module) {
    module.exports = jsstats;
}
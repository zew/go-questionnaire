// stackoverflow.com/questions/5259421
// normal (distribution) - cumulative distribution function
function normalCDF1(mean, sigma, to) {
    var z = (to - mean) / Math.sqrt(2 * sigma * sigma);
    var t = 1 / (1 + 0.3275911 * Math.abs(z));
    var a1 = 0.254829592;
    var a2 = -0.284496736;
    var a3 = 1.421413741;
    var a4 = -1.453152027;
    var a5 = 1.061405429;
    var erf = 1 - (((((a5 * t + a4) * t) + a3) * t + a2) * t + a1) * t * Math.exp(-z * z);
    var sign = 1;
    if (z < 0) {
        sign = -1;
    }
    return (1 / 2) * (1 + sign * erf);
}

function normalCDF2(x, mean, std) {
    x = (x - mean) / std
    let t = 1 / (1 + .2315419 * Math.abs(x))
    let d = .3989423 * Math.exp(-x * x / 2)
    let prob = d * t * (.3193815 + t * (-.3565638 + t * (1.781478 + t * (-1.821256 + t * 1.330274))))
    if (x > 0) prob = 1 - prob
    return prob
}

/* 
    console.log(normalCDF1(30, 25, 1.4241), `//-> 0.12651187738346226`, `wolframalpha.com  0.12651200000000000`);

    console.log(normalCDF2(0.0, 0, 17));
    console.log(normalCDF2(0.1, 0, 17));
    console.log(normalCDF2(0.5, 0, 17));
    console.log(normalCDF2(0.9, 0, 17));
    console.log(normalCDF2(1.0, 0, 17));
 */


var mn = 0.0; // mean
var sd = 1.0; // standard deviation
var normDist = new jsstats.NormalDistribution(mn, sd);

var x = 10.0; // point estimate value 
var p = normDist.cumulativeProbability(x); // cumulative probability
console.log(`x ${x}  => p ${p}`);

var p = 0.7; // cumulative probability
var x = normDist.invCumulativeProbability(p); // point estimate value
console.log(`p ${p}  => p ${x}`);


{

    var chartDom2 = document.getElementById('chart_container_2');
    var myChart2 = echarts.init(chartDom2);
    var option2;


    function dataNormalCDF2(sigma) {
        let data = [];
        for (let idx = -5.0; idx < 5.0; idx+= 0.2) {
            let item = [idx, normalCDF2(idx, 0, sigma)];
            data.push(item);
        }
        return data;
    }

    function dataNormalCDF3(sigma) {
        let data = [];
        // small steps
        for (let idx = 0.001; idx < 0.02; idx += 0.001) {
            let item = [idx, normDist.invCumulativeProbability(idx)];
            data.push(item);
        }
        // big steps
        for (let idx = 0.02; idx < 1.0; idx += 0.02) {
        // for (let idx = -4.0; idx <= 4.0; idx += 0.1) {
            let item = [idx, normDist.invCumulativeProbability(idx)];
            data.push(item);
        }
        // small steps
        for (let idx = 0.98; idx < 1.00; idx += 0.001) {
            let item = [idx, normDist.invCumulativeProbability(idx)];
            data.push(item);
        }
        console.log(data);
        return data;
    }


    option2 = {
        tooltip: {},
        xAxis:   {},
        yAxis:   {},
        series: [
            // {
            //     type: 'line',
            //     data: dataNormalCDF2(1),
            // },
            {
                type: 'line',
                data: dataNormalCDF3(1),
            },
        ],
    };
    option2 && myChart2.setOption(option2);
}



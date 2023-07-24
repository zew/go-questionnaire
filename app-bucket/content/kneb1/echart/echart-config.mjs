// echart configuration




// Carolin-01-start


let yr = new Date().getFullYear()
let az = 20; // "Anlagehorizont"
let sb = 100.0; // sparbetrag
let sby = 12* sb; // sparbetrag per year

// standardized normal distribution
let mn = 0.0; // mean
let sd = 1.0; // standard deviation

// normal distribution of
// MSCI world for € investments since 1998 (25yrs)
mn = 0.059
sd = 0.1462

// 90 confidence interval - multiple of sd
let ci90 = 1.645 * sd

let p05 = mn * (1-ci90)
let p95 = mn * (1+ci90)

console.log(`pct05 ${p05}  -- mn ${mn}   pct95 ${p95}`)

let p05p1 = 1 + p05  // worst case plus one
let mnp1  = 1 + mn  // mean plus one
let p95p1 = 1 + p95  // worst case plus one

console.log(`pct05+1 ${p05p1}  -- mn+1 ${mnp1}   pct95+1 ${p95p1}`)



// github.com/chen0040/js-stats
// used for converting random draws from [0,1] into norm dist probs
let normDist = new jsstats.NormalDistribution(mn, sd);




// unused primitive series...
let dataXAxix = []; // unused
for (let i = yr; i <= yr+az; i++) { dataXAxix.push(i); }
let dataReturns = []; // unused
for (let i = 0; i <= az; i++) {     dataReturns.push(250+i*2000); }


// console.log(ds2.source);




// stackoverflow.com/questions/1479319/
var myInstance = (function () {
    var privateVar = '';
    function privateMethod() {
    }
    // public interface
    return {
        publicMethod1: function () {
        },
        publicMethod2: function () {
        }
    };
})();


var dataObject = (function () {

    var ds = []; // private

    var resetDataPriv = () => {ds = [] }

    // computeDataPriv compiles data for eChart options object
    // usage:
    //       myChart.setOption({
    //          dataset: dataObject.computeData(),
    //       });
    var computeDataPriv = () => {
        //
        // random draws - mapped to normal dist.
        if (false) {
            for (let i = ds1.source.length; i < (counterDraws+1); i++) {
                let linDraw = Math.random(); // a number from 0 to <1

                while (linDraw == 0.0) {
                    // just avoid 0.0, because it creates infinity below
                    linDraw = Math.random();
                }

                let draw  = normDist.invCumulativeProbability(linDraw)
                // console.log(`   lin draw ${linDraw} => draw  ${draw}`);

                let subAr = ["draw", draw];
                ds1.source.push(subAr);
            }
        }
        // console.log(`counterDraws ${counterDraws} - ds1a: `, ds1a.source );

        if (ds === undefined || ds.length == 0) {
            ds = []
            let c0=0, c1=0, c2=0
            for (let i = 0; i <= az; i++) {
                // return on existing balance
                c0 *= p05p1; c1 *= mnp1; c2 *=p95p1;
                // additional annuity
                c0 += sby; c1 += sby; c2 +=sby;
                let row = [yr+i, c0, c1, c2, `item${i}` ]
                // console.log(i, i+yr, mnp1**i);
                // console.log(row);
                ds.push( row );
            }
            // console.log(ds);
            console.log(`dataObject - ds recomputed - length ${ds.length}`);

            console.log(ds);

        }

        return ds



        return [
            // [col1, col2, col3 ... ]
            // [dimX, dimY, other dimensions ...
            // In cartesian (grid), "dimX" and "dimY" correspond to xAxis and yAxis respectively.
            //    see      https://echarts.apache.org/en/option.html#series-line
            //    search   'Relationship between "value" and axis.type'
            //
            [2023,     950+a,    175+a , 'item-1'   ],
            [2024,    2900+a,   2200+a , 'item-2'   ],
            [2025,    4400+a,   4000+a , 'item-3'   ],
            [2026,    5000+a,   4000+a , 'item-4'   ],
            [2027,    6500+a,   4500+a , 'item-5'   ],
            [2029,   13500+a,   4500+a , 'item-6'   ],
            [2029.5, 13800+a,   7800+a , 'item-7'   ],
            [2030,          ,   8000+a , 'item-8'   ],
            [2031,   22000+a,  20000+a , 'item-9'   ],
            [2034,   24000+a,  23000+a , 'item-10'  ],
            [2036,   26000+a,  24000+a , 'item-11'  ],
            [2037,   36000+a,  33000+a , 'item-12'  ],
            [2043,   38000+a,  34000+a , 'item-12'  ],
        ];

    }

    // public interface
    return {
        resetData:   resetDataPriv,
        computeData: computeDataPriv,
    };


})();



let vertMarkerYr = yr + az/2;
let vertMarker1 = [
    {
        name: 'Ihr gewählter Anlagehorizont',
        xAxis: 2029-0.3,
        xAxis: vertMarkerYr - 0.12,
    },
    {
        xAxis: vertMarkerYr + 0.12,
    }
];
let vertMarker2 = [
    {
        name: 'Evening Peak',
        xAxis: 2034,
    },
    {
        xAxis: 2036,
    }
];
// used on second series in setOptions
let markArea = {
    itemStyle: {
      color: 'rgba(255, 173, 177, 0.4)'
    },
    data: [vertMarker1, vertMarker2],
    data: [vertMarker1],
};



// chart config variables
var seriesIdx = -1;
var animDuration = 800;
var colorPalette = [
    'rgba( 2,134,228,0.6)',
    'rgba( 0,105,180,0.9)',
    'rgba( 2,134,228,0.6)',
    '#229',
    '#22b',
    '#229',
    '#22c',
    '#22d',
    // 'var(--clr-pri-hov);',
    ];




var optEchart = {
    // echarts.apache.org/handbook/en/concepts/dataset/
    // dataset: [],
    title: {
        // text: 'ECharts Getting Started Example'
        text: 'Auszahlungen',
        left: '1%'
    },
    tooltip: {},
    toolbox: {
        show: true,
        right: 10,
        feature: {
            saveAsImage: { show: true },
            // magicType:   { show: true, type: ['stack', 'tiled'] },
            // dataZoom: { yAxisIndex: 'none' },
            // restore: {},
        }
    },
    grid: {
        left:  '12%',
        left:  '13%',
        right:  '3%',
        top:    '8.5%',
        top:    '9%',
        bottom: '7%',
      },
    legend: {
        // data: ['sales']
    },
    xAxis: {
        // type: 'category',
        type: 'time',
        type: 'value',

        // only in numerical axis, i.e., type: 'value'.
        //    show zero position only, if justified by data
        //    if min and max are set, this setting has no meaning
        scale: true,

        // animation only makes sense for series
        animation: false,

        axisLabel: {
            // compare  axisLabel. formatter
            formatter: function (vl, idx) {
                return vl + ' ';
            },
            textStyle: {
                // color: function (vl, idx) {
                //     return vl >= 2030 ? 'green' : 'red';
                // }
            },
        },
        // data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri']
        // data: dataXAxix,
        // min: dataXAxix[0]-2,
        // min: 2000,
        // min: 'dataMin',
        min: function (vl) {
            // this effectively returns dataReturns.min and max
            console.log(`min ${vl.min} max ${vl.max} `)
            return vl.min;
        },
        min: yr+0,
        max: yr+az,

        // show label of the min/max tick
        //    seems ineffective, labels are shown anyway
        showMinLabel: false,
        showMaxLabel: false,

        // number of ticks - recommendation
        splitNumber: 8,
        // number of ticks - mandatory
        interval: 3,
        interval: 2,


        axisTick: {
            show: false,
            show: true,
            length: 4,
        },
        minorTick: {
            show: true,
        },
        // inverse: true,
        axisLine: {
            // show: false,
            // whether axis lies on the other's origin position - where value is 0 on axis.
            onZero: false,
            onZero: true,
        },

    },
    yAxis: {
        type: 'value',
        // name: 'y-axis-name',
        min: 0,
        max: 40*1000,

        axisLabel: {
            // compare  axisLabel.formatter
            formatter: function (vl, index) {
                // adjust grid.left
                let vl1 = vl.toFixed(0)
                vl1 = vl1 + ' €';
                vl1 = vl1.replace("000 €", ".000 €", )
                return vl1;
            },
        },

    },
    series: [
        {
            // name - only if we want it to be shown
            // name: 'series1',
            type: 'line',
            dummy: seriesIdx++,
            color: colorPalette[seriesIdx],
            
            showSymbol: true,
            showSymbol: false,
            symbol: 'emptyCircle',
            symbolSize: 4,

            animation: false,
            animation: true,
            animationDelay:    seriesIdx * animDuration,
            animationDuration: animDuration,

            // explanation for encode:
            //      see 10 lines below - 'data'
            //      see     https://echarts.apache.org/en/option.html#
            //      search  'series-line. encode'
            encode: {
                x: 0,
                y: 1,
                itemName: 4,
                tooltip: [0, 1, 4],
             },
            data: [
                // [col1, col2, col3 ... ]
                // [dimX, dimY, other dimensions ...
                // In cartesian (grid), "dimX" and "dimY" correspond to xAxis and yAxis respectively.
                //    see      https://echarts.apache.org/en/option.html#series-line
                //    search   'Relationship between "value" and axis.type'
                //
                [2023,     950,   175  , 'item-1'   ],
                [2024,    2900,   2200 , 'item-2'   ],
                [2025,    4400,   4000 , 'item-3'   ],
                [2026,    5000,   4000 , 'item-4'   ],
                [2027,    6500,   4500 , 'item-5'   ],
                [2029,   13500,   4500 , 'item-6'   ],
                [2029.5, 13800,   7800 , 'item-7'   ],
                [2030,        ,   8000 , 'item-8'   ],
                [2031,   22000,  20000 , 'item-9'   ],
                [2034,   24000,  23000 , 'item-10'  ],
                [2036,   26000,  24000 , 'item-11'  ],
                [2037,   36000,  33000 , 'item-12'  ],
            ],
            data: dataObject.computeData(),
        },

        {
            // name - only if we want it to be shown
            // name: 'series2',
            type: 'line',
            dummy: seriesIdx++,
            color: colorPalette[seriesIdx],
            
            showSymbol: true,
            symbol: 'circle',
            symbolSize: 6,

            animation: false,
            animation: true,
            animationDelay:    seriesIdx * animDuration,
            animationDuration: animDuration,

            // same data struct, but
            // y: 2 instead of 1
            encode: {
                x: 0,
                y: 2,
                itemName: 4,
                tooltip: [0, 2, 4],
             },
             data: dataObject.computeData(),
             markArea: markArea,
        },



        {
            type: 'line',
            dummy: seriesIdx++,
            color: colorPalette[seriesIdx],
            
            showSymbol: true,
            showSymbol: false,
            symbol: 'emptyCircle',
            symbolSize: 4,

            animation: false,
            animation: true,
            animationDelay:    seriesIdx * animDuration,
            animationDelay:            0 * animDuration,
            animationDuration: animDuration,

            // same data struct, but
            // y: 2 instead of 1
            encode: {
                x: 0,
                y: 3,
                itemName: 4,
                tooltip: [0, 3, 4],
             },
             data: dataObject.computeData(),
        },





    ]
};

// creation of chart object => common.js - initPage()


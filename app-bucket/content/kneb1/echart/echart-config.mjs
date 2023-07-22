// echart configuration

// github.com/apache/echarts
// github.com/ecomfe/echarts-stat
// github.com/ecomfe/awesome-echarts
// echarts.apache.org/examples/en/editor.html?c=bar-histogram

// import does not work; possibly because of stackoverflow.com/questions/71022803
// import * as echarts from './echarts.min.js';
// import ecStat from './ecStat.js';



var chartDom = document.getElementById('chart_container');
// console.log(chartDom);
var myChart = echarts.init(chartDom);
var opt1;
var opt2;



// var colorPalette = ['#d87c7c', '#919e8b', '#d7ab82', '#6e7074', '#61a0a8', '#efa18d', '#787464', '#cc7e63', '#724e58', '#4b565b'];
var colorPalette = [
    '#229',
    '#22b',
    '#22c',
    '#22d',
    // 'var(--clr-pri-hov);',
    ];
function getColor() {
    let idx = colorPalette.length % counterDraws;
    return colorPalette[seriesIdx];
}


// histogram config
// ======================

const w  = 10;   // width
const wh = w/2;  // width half

var maxXHisto = 0;
function getMax() {
    return maxXHisto + 2;
}



// risky asset random draws
var ds1Example = {
    source: [
        [10.3, 143],
        [10.6, 214],
        [10.8, 251],
        [10.0, 176],
        [10.1, 221],
        [10.2, 188],
        [10.4, 191],
        [10.0, 196],
        [10.9, 177],
        [10.9, 153],
        [10.3, 201],
    ]
};



var counterDrawsInit = 4 ;
var counterDraws = counterDrawsInit;  // counter for getData


// Carolin-01-start
var mn = 0.0; // mean
var sd = 1.0; // standard deviation

var mn = 153.0; // mean
var sd = 15.89; // standard deviation

// github.com/chen0040/js-stats
var normDist = new jsstats.NormalDistribution(mn, sd);


var ds2 = {
    source: []
};
for (let i = 0; i <= 300/w; i++) {
    ds2.source.push([wh + i*w, 0]);
}

// console.log(ds2.source);

// getData compiles data for eChart options object
// usage:
//       myChart.setOption({
//          dataset: getData(),
//       });
function getData() {

    counterDraws++;
    let a = 600*counterDraws;

    if (false) {
        try {
            // both cases lead to infinity
            console.log( "icnp(0.0):", normDist.invCumulativeProbability(0)    );
            console.log( "icnp(1.0):", normDist.invCumulativeProbability(1.0)  );
        } catch (error) {
            console.error(error);
        }
    }

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
    ];

}



var dataXAxix = []; // unused
let iStart = new Date().getFullYear()
for (let i = iStart; i <= iStart+15; i++) {
    dataXAxix.push(i);
}
var dataReturns = []; // unused
for (let i = 0; i <= 15; i++) {
    dataReturns.push(250+i*2000);
}



let seriesIdx = -1;
let animDuration = 800;

opt2 = {
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
        top:    '7%',
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
        min: iStart+0,
        max: iStart+15,

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
            symbol: 'emptyCircle',
            symbolSize: 6,
            showSymbol: true,
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
                itemName: 3, 
                tooltip: [0, 1, 3],
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
            data: getData(),
        },

        {
            // name - only if we want it to be shown
            // name: 'series2',
            type: 'line',
            dummy: seriesIdx++,
            color: colorPalette[seriesIdx],

            symbol: 'emptyCircle',
            symbolSize: 4,
            showSymbol: true,
            animation: false,
            animation: true,
            animationDelay:    seriesIdx * animDuration,
            animationDuration: animDuration,

            // same data struct, but
            // y: 2 instead of 1
            encode: { 
                x: 0, 
                y: 2, 
                itemName: 3, 
                tooltip: [0, 2, 3],
             },
             data: getData(),
        },



    ]
};

// opt1 && myChart.setOption(opt1);
opt2 && myChart.setOption(opt2);
console.log(`echart config and creation complete`)



// https://github.com/apache/echarts
// https://github.com/ecomfe/echarts-stat
// https://github.com/ecomfe/awesome-echarts
// https://echarts.apache.org/examples/en/editor.html?c=bar-histogram

// import simply does not work; possibly because of stackoverflow.com/questions/71022803
// import * as echarts from './echarts.min.js';
// import ecStat from './ecStat.js';



var chartDom = document.getElementById('chart_container');
// console.log(chartDom);

var myChart = echarts.init(chartDom);
var option;



// see https://github.com/ecomfe/echarts-stat
echarts.registerTransform(ecStat.transform.histogram);

var ds1 = {
    source: [
        [10.3, 143],
        [10.6, 214],
        [10.8, 251],
        [10.7, 86],
        [10.8, 93],
        [10.0, 176],
        [10.1, 221],
        [10.2, 188],
        [10.4, 91],
        [10.4, 191],
        [10.0, 196],
        [10.9, 177],
        [10.9, 153],
        [10.3, 201],
        [10.7, 199],
        [10.2, 98],
        [10.5, 121],
        [10.3, 105],
        [10.5, 168],
        [10.9, 84],
        [10.0, 197],
        [10.0, 155],
        [10.6, 125]
    ]
};

var ds2 = {
    source: [
        [ 25,  0],
        [ 75,  5],
        [125,  4],
        [175, 12],
        [225,  3],
        [275,  1],
    ]
};

var counterDraws = -1  // counter for getData

function getData() {

    counterDraws++;

    let ds1a = ds1;
    let ds2a = ds2;


    ds1a.source = [];
    for (let i = 0; i < (counterDraws+4); i++) {
        let val = 90.0 + 90*Math.random();
        let subAr = ["draw", val];
        ds1a.source.push(subAr);
    }

    console.log(`counterDraws ${counterDraws} - ds1a: `, ds1a.source );
    
    ds2a.source = [
        [ 25, 0],
        [ 75, 0],
        [125, 0],
        [175, 0],
        [225, 0],
        [275, 0],
    ];
    for (let i = 0; i < ds1a.source.length; i++) {

        let val = Math.floor(ds1a.source[i][1]);

        let binId = Math.round(val/50)*50 + 25;
        
        let binIdx = (binId - 25) / 50;
        
        // console.log(`   val ${val} => binId ${binId} - => binIdx ${binIdx}`);

        ds2a.source[binIdx][1]++;
    }

    console.log(`counterDraws ${counterDraws} - ds2a: `, ds2a.source);

    return [
        ds1a,
        ds2a,
        {
            transform: {
                type: 'ecStat:histogram',
                // print: true,
                config: { dimensions: [1] }
            }
        },
    ];
}

option = {
    dataset: getData(),
    tooltip: {},
    grid: [
        {
            top:   '04%',
            right: '60%',
        },
        {
            top:    '04%',
            left:   '56%',
        }
    ],
    xAxis: [
        {
            type:  'value',
            type:  'category',
            // do not include zero position 
            scale:  true,  
            gridIndex: 0
        },
        {
            scale: true,
            gridIndex: 1,
            inverse: true,
        }
    ],
    yAxis: [
        {
            gridIndex: 0,
            min: 0,
            max: 300,
        },
        {
            gridIndex: 1,
            // min: 0,
            // max: 350,
            // must be category: https://github.com/apache/echarts/issues/15960
            //      or https://echarts.apache.org/en/option.html#series-custom
            type: 'category',
            // axisTick:  { show: false },
            // axisLabel: { show: false },
            // axisLine:  { show: false },

            axisLine: {
                // necessary for position: right to take effect
                onZero: false,
            },
            position: 'right',

        }
    ],
    series: [
        {
            name: 'origianl scatter',
            type: 'scatter',
            xAxisIndex: 0,
            yAxisIndex: 0,
            encode: { tooltip: [0, 1] },
            datasetIndex: 0
        },
        {
            name: 'histogram',
            type: 'bar',
            xAxisIndex: 1,
            yAxisIndex: 1,
            barWidth: '99.3%',
            barWidth: '2px',
            label: {
                show: true,
                position: 'right'
            },
            encode: { x: 1, y: 0, itemName: 4 },
            datasetIndex: 1
        }
    ]
};

option && myChart.setOption(option);


setInterval(() => {
    myChart.setOption({
        dataset: getData(),
        // series: {
        //     data: makeRandomData()
        // }
    });
}, 2000);
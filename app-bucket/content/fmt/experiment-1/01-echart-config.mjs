"strict mode";


// the fields for permanent storage of response value in go-questionnaire 
//   come in distinct instances, suffixed by  "_0"  and  "_1"
const getBackgroundField = core => {
    const inst = document.getElementById("simtool_instance");
    if (!inst) {
        throw `hidden input 'simtool_instance' required`;
    }
    const nameInstance = `${core}_${inst.value}`;
    const el = document.getElementById(nameInstance);
    if (!el) {
        throw `hidden input ${core}_${inst.value} required`;
    }
    return el
};


// stuff declare using 'var xyz' is intentionally for usage across scopes 

var simHist = [];
var simHistInp   = getBackgroundField("sim_history");


var param1Inp   = getBackgroundField("param1");
var param1InpBG = getBackgroundField("param1_bg");
if (param1InpBG && param1InpBG.value !== "") {
    param1Inp.value = parseInt(param1InpBG.value);  // restore from before
}


var param2Inp   = getBackgroundField("param2");
var param2InpBG = getBackgroundField("param2_bg");
if (param2InpBG && param2InpBG.value !== "") {
    param2Inp.value = parseInt(param2InpBG.value);  // restore from before
}


var az = 50;


var slider01 = param1Inp;

let initSlider01 = (inst) => {
    slider01.value = 25;
    // slider01.click();
    const evt = new Event("input");
    slider01.dispatchEvent(evt);
}

// init sliders;
window.addEventListener('load', initSlider01, false);


// update sliders
slider01.oninput = function () {
    let incr = parseInt(this.value) + 5;
    // slider01Legend.value =  `${this.value}  -  ${incr}`;
}




const yr=2025;

function getChartTitle(yr) {
    return `Prognostizierte Entwicklung bis ${yr}`;
    // return `Prognostizierte Entwicklung der Ernte über das Jahr`;
}

function getXAxisType() {
    return 'value';
    // return 'category';
}
function getXAxisMin(yr) {
    return yr + 0;
    // return 'Okt';
}
function getXAxisMax(yr) {
    return yr + 0;
    // return 'Juli';
}

function getXAxisFormatter(vl, yr) {
    return vl + ' ';
}








function getXInteerval() {
    // return 2.285;
    return 2;
}



function getYAxisTitle() {
    return `Ertrag in Euro`;
}
function getVerticalMarkerTitle() {
    // return `Ernte`;
    return `Bäume werden\n gefällt`;
}










console.log(`init sb ${param1Inp.value}`)




// stackoverflow.com/questions/1479319/
// "class" dataObjectCreate() below is created in this pattern:
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


var dataObjectCreate = (function () {

    // private members
    var ds = [];
    var maxY = 40 * 1000;

    // private method
    var pResetData = () => {
        ds = [];
        maxY = 40 * 1000;
    }

    // private method
    // get max Y
    var pMaxY = () => {
        pComputeData()
        if (maxY < 40 * 1000) {
            return 40 * 1000
        }
        return maxY
    }


    // private method
    // computeDataPriv compiles data for eChart options object
    // usage:
    //       myChart.setOption({
    //          dataset: dataObject.computeData(),
    //       });
    var pComputeData = () => {

        const a = 23;
        return [
            // [col1, col2, col3 ... ]
            // [dimX, dimY, other dimensions ...
            // In cartesian (grid), "dimX" and "dimY" correspond to xAxis and yAxiis respectively.
            //    see      https://echarts.apache.org/en/option.html#series-line
            //    search   'Relationship between "value" and axis.type'
            //
            [2023,  950 + a, 175 + a, 'item-1'],
            [2024, 2900 + a, 2200 + a, 'item-2'],
            [2025, 4400 + a, 4000 + a, 'item-3'],
            [2026, 5000 + a, 4000 + a, 'item-4'],
            [2027, 6500 + a, 4500 + a, 'item-5'],
            [2029, 13500 + a, 4500 + a, 'item-6'],
            [2029.5, 13800 + a, 7800 + a, 'item-7'],
            [2030, , 8000 + a, 'item-8'],
            [2031, 22000 + a, 20000 + a, 'item-9'],
            [2034, 24000 + a, 23000 + a, 'item-10'],
            [2036, 26000 + a, 24000 + a, 'item-11'],
            [2037, 36000 + a, 33000 + a, 'item-12'],
            [2043, 38000 + a, 34000 + a, 'item-12'],
        ];

    }

    // public interface
    return {
        resetData: pResetData,
        maxY: pMaxY,
        computeData: pComputeData,
    };


});


var dataObject = dataObjectCreate();


getVerticalArea = function (argYryr, argAzV) {

    let vertMarkerYr = argYryr + argAzV;
    let vertMarker1 = [
        {
            name: getVerticalMarkerTitle(),
            xAxis: 2029 - 0.3,
            xAxis: vertMarkerYr - 0.08,
        },
        {
            xAxis: vertMarkerYr + 0.08,
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
        label: {
            // show: false,
            color: 'rgba( 0,105,180,0.99)',
        },
        // animation: true,
        // animationDurationUpdate: 200,
        itemStyle: {
            color: 'rgba(255, 188, 188, 0.6)',
            color: 'rgba( 0,105,180,0.299)',

        },
        data: [vertMarker1, vertMarker2],
        data: [vertMarker1],
    };

    return markArea;

}




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
        text: 'Angespartes Vermögen',
        text: getChartTitle(yr),
        left: '1%',
        textStyle: {
            fontSize: 18,
            fontSize: 15,
        }
    },
    tooltip: {},
    toolbox: {
        show: true,
        right: 10,
        feature: {
            saveAsImage: { show: true },
            // magicType:   { show: true, type: ['stack', 'tiled'] },
            // dataZoom: { yAxiisIndex: 'none' },
            // restore: {},
        }
    },
    grid: {
        left: '12%',
        left: '13%',
        right: '3%',
        top: '8.5%',
        top: '9.8%',
        bottom: '7%',
    },
    legend: {
        // data: ['sales']
    },
    xAxis: {
        // type: 'category',
        type: 'time',
        type: 'value',
        type: getXAxisType(),

        // only in numerical axis, i.e., type: 'value'.
        //    show zero position only, if justified by data
        //    if min and max are set, this setting has no meaning
        scale: true,

        // animation only makes sense for series
        animation: false,

        axisLabel: {
            // compare  axisLabel. formatter
            formatter: function (vl, idx) {
                return getXAxisFormatter(vl, yr);
            },
            textStyle: {
                // color: function (vl, idx) {
                //     return vl >= 2030 ? 'green' : 'red';
                // }
            },
        },
        // data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri'],
        // data: dataXPrimitive,
        // min:  dataXPrimitive[0]-2,
        // min:  2000,
        // min: 'dataMin',
        min: function (vl) {
            // this effectively returns dataReturns.min and max
            console.log(`x-axis min ${vl.min} max ${vl.max} `)
            return vl.min - 10;
        },
        min: getXAxisMin(yr),
        max: getXAxisMax(yr),

        // show label of the min/max tick
        //    seems ineffective, labels are shown anyway
        showMinLabel: false,
        showMaxLabel: false,

        // number of ticks - recommendation
        splitNumber: 8,
        // number of ticks - mandatory
        interval: 3,
        interval: 2,
        interval: getXInteerval(),

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
        max: 40 * 1000, // init
        max: dataObject.maxY(),

        //
        name: getYAxisTitle(),
        nameLocation: 'middle',
        nameGap: 62,
        nameTextStyle: {
            fontSize: 12,
            fontSize: 14,
        },


        axisLabel: {
            // compare  axisLabel.formatter
            formatter: function (vl, index) {
                // adjust grid.left
                let vl1 = vl.toFixed(0)
                vl1 = vl1 + ' €';
                vl1 = vl1.replace("000 €", ".000 €",)
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
            animationDelay: seriesIdx * animDuration,
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
                [2023, 950, 175, 'item-1'],
                [2024, 2900, 2200, 'item-2'],
                [2025, 4400, 4000, 'item-3'],
                [2026, 5000, 4000, 'item-4'],
                [2027, 6500, 4500, 'item-5'],
                [2029, 13500, 4500, 'item-6'],
                [2029.5, 13800, 7800, 'item-7'],
                [2030, , 8000, 'item-8'],
                [2031, 22000, 20000, 'item-9'],
                [2034, 24000, 23000, 'item-10'],
                [2036, 26000, 24000, 'item-11'],
                [2037, 36000, 33000, 'item-12'],
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
            animationDelay: seriesIdx * animDuration,
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

            markArea: getVerticalArea(yr, 50),
            //  markPoint: getMarkpointConfig( [30000, 33000, 34000] ),
            //  markPoint: "getMarkpointConfig( dataObject.FVs() )",


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
            animationDelay: seriesIdx * animDuration,
            animationDelay: 0 * animDuration,
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

function refresh(chartObj, dataObj) {

    dataObj.resetData()

    // setOption or resize
    // chartObj.resize();

    if (true) {
        chartObj.setOption({
            'xAxis': {
                max: yr + az,
            },
            'yAxis': {
                max: dataObject.maxY(),
            },
            series: [
                {
                    data: dataObj.computeData(),
                },
                {
                    data: dataObj.computeData(),
                    markArea: getVerticalArea(yr, 50),
                    // markPoint: getMarkpointConfig( [30000, 33000, 34000] ),
                    // markPoint: getMarkpointConfig( dataObj.FVs() ),
                },
                {
                    data: dataObj.computeData(),
                },
            ]
        });
    }

    let arrayFVs = dataObj.FVs()



}


// creation of chart object => common.js - initPage()


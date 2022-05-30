// import * as echarts from 'echarts';

const symbolSize = 20;


function showTooltip(dataIndex) {
    chrt.dispatchAction({
        type: 'showTip',
        seriesIndex: 0,
        dataIndex: dataIndex
    });
}

function hideTooltip(dataIndex) {
    chrt.dispatchAction({
        type: 'hideTip'
    });
}

// recompute position of data points
//   on zoom events, on page resize events
function updatePosition() {
    chrt.setOption({
        graphic: dataSeries.map(function (item, dataIndex) {
            let pos = chrt.convertToPixel('grid', item);
            // console.log(`pos is ${pos}`);
            return {
                position: pos
            };
        })
    });
}

function onPointDrag(dataIndex, pos) {
    dataSeries[dataIndex] = chrt.convertFromPixel('grid', pos);
    console.log(`drag pos ${pos}`);  // pos is an array with index 0,1
    // update data
    chrt.setOption({
        series: [
            {
                id: 'a',
                data: dataSeries
            }
        ]
    });
}



const dataSeries = [
  [ 40  , -10],
  [-30  , -05],
  [-76.5,  20],
  [-63.5,  40],
  [-22.1,  50]
];

const zoomConfig1 = [];
const zoomConfig2 = [
    {
        type: 'slider',
        xAxisIndex: 0,
        filterMode: 'none'
    },
    {
        type: 'slider',
        yAxisIndex: 0,
        filterMode: 'none'
    },
    {
        type: 'inside',
        xAxisIndex: 0,
        filterMode: 'none'
    },
    {
        type: 'inside',
        yAxisIndex: 0,
        filterMode: 'none'
    }
]




let option = {
    title: {
        text: 'Points drag',
        left: 'center'
    },
    tooltip: {
        triggerOn: 'none',
        formatter: function (params) {
            return (
                'X: ' +
                params.data[0].toFixed(2) +
                '<br>Y: ' +
                params.data[1].toFixed(2)
            );
        }
    },

    grid: {
        top: '8%',
        bottom: '12%'
    },

    xAxis: {
        min: -100,
        max:   70,
        type: 'value',
        axisLine: { onZero: false }
    },

    yAxis: {
        min: -30,
        max:  60,
        type: 'value',
        axisLine: { onZero: false }
    },

    dataZoom: zoomConfig1,

    series: [
        {
            id: 'a',
            type: 'line',
            smooth: true,
            symbolSize: symbolSize,
            data: dataSeries
        }
    ]

};


// additional option settings for dragging functionality
//   with minimal event queue deference; why not directly?
setTimeout(  () => {
    // Add shadow circles (which is not visible) to enable drag.
    chrt.setOption({
        graphic: dataSeries.map(function (item, dataIndex) {
            return {
                type: 'circle',
                position: chrt.convertToPixel('grid', item),
                shape: {
                    cx: 0,
                    cy: 0,
                    r: symbolSize / 2
                },
                invisible: true,
                draggable: true,
                ondrag: function (dx, dy) {
                    onPointDrag(dataIndex, [this.x, this.y]);
                },
                onmousemove: function () {
                    showTooltip(dataIndex);
                },
                onmouseout: function () {
                    hideTooltip(dataIndex);
                },
                z: 100
            };
        })
    });
}, 0);


let chartDom = document.getElementById('container');
// let chrt     = echarts.init(chartDom, 'dark');
let chrt     = echarts.init(chartDom);
console.log(`echart object created`);

if (option && typeof option === 'object') {
    chrt.setOption(option);
    console.log(`echart options set`);
}

window.addEventListener('resize', updatePosition);
// window.addEventListener('resize', chrt.resize);
chrt.on('dataZoom', updatePosition);

console.log(`echart all success`);
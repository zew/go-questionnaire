const chartIDs = ['chart_container',];
let chartObjs = [];





document.addEventListener('DOMContentLoaded', () => {


    chartIDs.forEach(function (id) {
        let el = document.getElementById(id);
        const ch = echarts.init(el);
        chartObjs.push(ch);
    });


    function updateCharts1(state=0, bValue=3, cValue=4) {

        // console.log(`updateCharts1() start`);

        function swapping(nodes, bValue=3, cValue=4){

            if (nodes.length===1){
                return nodes;
            }
            if (nodes.length===2){
                return nodes;
            }

            if (bValue < cValue){
                return nodes;
            }

            console.log(`swapping`);

            el1 = nodes[0];
            el2 = nodes[1];
            el3 = nodes[2];

            el2["name"] = 'c';
            el3["name"] = 'b';
            return [el1, el2, el3];


        }

        let statesNodes = [
            [
                {
                    name: 'a',
                    label: { show: true, formatter: '100 €'  },
                    itemStyle: { color: '#444' },
                },
                {
                    name: 'b',
                    label: { show: true, formatter: 'to allocate'  },
                    itemStyle: { color: '#444' },
                },

            ],
            [
                {
                    name: 'a',
                    label: { show: true, formatter: '100 €'  },
                    itemStyle: { color: '#444' },
                },
                {
                    name: 'b',
                    label: { show: true, formatter: 'Bonds' },
                    itemStyle: { color: '#494' },
                    // order: 0,
                },
                {
                    name: 'c',
                    label: { show: true, formatter: 'Stocks' },
                    itemStyle: { color: '#944' },
                    // order: 1,
                },

            ],
        ];

        let statesLinks = [
            [
                { source: 'a', target: 'b', value: bValue },
            ],
            [
                { source: 'a', target: 'b', value: bValue },
                { source: 'a', target: 'c', value: cValue },
            ],
        ];



        let opt = {
            tooltip: {
                trigger:   'item',
                triggerOn: 'mousemove'
            },
            animation: false,
            series: [
                {
                    type: 'sankey',

                    top:   '15%',
                    bottom: '1%',
                    left:   '1%',
                    right:  '1%',
                    width: '98%',

                    nodeAlign: 'justify',        // alignment only; order comes from data[]
                    layoutIterations: 0,       
                    layoutIterations: 20,       // prevent over-compression of layout

                    nodeGap:   30,             // ↑ increase space between streams (default ~8)
                    nodeWidth: 20,             // thickness of the node
                    draggable: false,


                    emphasis: {
                        // focus: 'adjacency'
                    },
                    data:  swapping(statesNodes[state], bValue=3, cValue=4)   ,
                    // data: anchors.concat(nodes),
                    links: statesLinks[state],


                    orient: 'vertical',
                    label: {
                        position: 'top'
                    },
                    lineStyle: {
                        color:    'source',
                        // curveness: 0.5
                        curveness: 0.9,
                    },
                    emphasis: { focus: 'none' }                    ,
                }
            ]
        };

        chartObjs[0].setOption(opt);
        // console.log(`distance chart ${chartObjs[0]}  `);

        console.log(`updateCharts1() stop`);

    }



    updateCharts1();


    // handle radio changes
    const radios = document.querySelectorAll('input[name="flowRatio"]');
    radios.forEach(function (radio) {
        radio.addEventListener('change', function (event) {

            const n = parseInt(event.target.value);

            const bValue = n;
            const cValue = 7 - n;
            // console.log(`Radio changed → b=${bValue}, c=${cValue}`);
            updateCharts1(
                1,
                bValue,
                cValue,
            );
        });
    });



    // charts responsive on browser resize
    window.addEventListener('resize', function () {
        chartObjs.forEach(function (chartObj) {
            chartObj.resize();
        });
    });


    console.log(`pageLoaded() pg3 complete`)




});

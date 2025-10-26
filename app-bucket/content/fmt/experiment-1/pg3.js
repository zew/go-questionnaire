// parsedData - see data.js
// const dbg = JSON.stringify(parsedData[participantIdx], null, 2);
// console.log(` participantIdx ${participantIdx} - data ${dbg} `);

// forecastData - inserted by server
const dbg = JSON.stringify(forecastData, null, 2);
console.log(` data ${dbg} `);
console.log(` user ID  ${forecastData['user_id']} - grp ${forecastData['group']} `);

const chartIDs  = ['distanceChart', 'forecastChart', 'consensusChart'];
let   chartObjs = [];



document.addEventListener('DOMContentLoaded', () => {

    const shareComparisonText = document.getElementById('shareComparisonText');
    shareComparisonText.innerHTML = `shareComparisonText content`;

    chartIDs.forEach(function (id) {
        let el = document.getElementById(id);
        const ch = echarts.init(el) ;
        chartObjs.push(ch);
    });



    function updateCharts() {

        function formatDE(value, digits = 2) {
            if (isNaN(value)) return '';
            return value.toLocaleString("de-DE", {
                minimumFractionDigits: digits,
                maximumFractionDigits: digits
            });
        }

        /*
        function getSelectedQuarter() {
            return 'Q4_2025';  // Q1_2026, Q2_2026
            return document.querySelector('input[name="quarter"]:checked').value;
        }


        const participantIdx = parseInt(10, 10);
        const quarter        = getSelectedQuarter();

        if (isNaN(participantIdx) || !parsedData[participantIdx]) {

            console.log(` nothing found for participantIdx ${participantIdx} - cleaning out  `);

            chartObjs.forEach(function (chartObj) {
                chartObj.clear();
            });

            if (shareComparisonText) {
                shareComparisonText.textContent = '';
            }
            return;
        }
        
        const participantDta = parsedData[participantIdx];
        */

        const participantDta = forecastData;

        //
        let   userShare      = 50.0;
        if (true){
            const userSharePrev = document.getElementById('user_share_prev');
            userShare      = parseFloat(userSharePrev.value) || 0;
            console.log(`userShare from prev page ${userShare} parsed from ${userSharePrev.value}`);
        }

        // const actualShareRaw = participantDta[`grshare${quarter}`];
        // const actualShare    = parseFloat(actualShareRaw) * 100;
        // const forecast       = parseFloat(participantDta[`growth${quarter}`]);
        // const consensus      = parseFloat(participantDta[`consensus${quarter}`]);

        // chart 1  - userShare vs actualShare
        const actualShareRaw = participantDta[`share_lower_Q42025`];
        const actualShare    = parseFloat(actualShareRaw) * 100;


        // chart 2+3
        let  quarter         = participantDta[`quarter`];
             quarter         = "Forecast\nfür Q4 2025\nin Prozent";

        const forecast       = parseFloat(participantDta[`Q42025`]);
        const consensus      = parseFloat(participantDta[`consensus`]);


        console.log(`actual ${actualShare} - forecast ${forecast} - consensus ${consensus}   `);


        if (!isNaN(actualShare)) {
            shareComparisonText.innerHTML = `Sie haben   
                    <strong>${formatDE(userShare,0)}%</strong> 
                    als 
                    <i style="color:#EE6666">Anteil mit niedrigerer Wachstumsprognose</i> angegeben. 
                <br>
                Tatsächlich lag der <i>Anteil unter allen Befragten</i>,
                die im August 2025 ein <i>niedrigeres</i> Wachstum als Sie angegeben haben,
                bei&nbsp;<strong><span style="color:#EE6666">${formatDE(actualShare,1)}%</span></strong>.`;
        } else {
            shareComparisonText.innerHTML = `Ihr Anteil: <strong>${formatDE(userShare,1)}%</strong> |
                Tatsächlicher Anteil: <strong>N/A</strong>`;
        }




        const lblCntr1 = document.getElementById('cou_nter1');
        const lblCntr2 = document.getElementById('cou_nter2');
        const lblCntr3 = document.getElementById('cou_nter3');
        if( forecastData['group'] !=="T" ){
            const domChart1 = document.getElementById('distanceChart');
            domChart1.style.display = "none";
            const chart1Header = document.getElementById('distanceChartHeader');
            chart1Header.style.display = "none";
            try {
                lblCntr1.innerHTML = "";
               
                lblCntr1.parentNode.style.display = "none"

                lblCntr2.innerHTML = "1";
                lblCntr3.innerHTML = "2";
            
            } catch (error) {
                console.error(error)                
            }
            



        } else {
            lblCntr1.innerHTML = "1";
            lblCntr2.innerHTML = "2";
            lblCntr3.innerHTML = "3";

        }


        // chart #1 - distance
        const distance = Math.abs(userShare - actualShare);
        const distanceChartOption = {
            grid:    { top: 20, right: 40, bottom: 20, left: 40 },
            xAxis:   { type: 'value',
                       min: 0, max: 100,
                       axisLabel: { formatter: '{value}%' },
                       splitLine: { show: false }
                     },
            yAxis:   { type: 'category', data: [''], show: false },
            series: [{
                type: 'scatter',
                symbolSize: 26,
                data: [
                    {
                        name: 'Ihr Anteil',
                        value: [ userShare, 'Ihr Anteil'],
                        itemStyle: { color: '#546a7b' },
                        label: {
                            show: true, position: 'top',
                            distance: 18,
                            fontWeight: 'bold', color: '#000000',
                            formatter: (params) => formatDE(params.value[0],0) + '%',
                        }
                    },
                    {
                        name: 'Tatsächlicher Anteil',
                        value: [actualShare, ''],
                        itemStyle: { color: '#EE6666' },
                        label: {
                            show: true, position: 'bottom',
                            distance: 8,
                            // padding: [-2,0,0,0],
                            fontWeight: 'bold', color: '#EE6666',
                            formatter: (params) => ' ' + formatDE(params.value[0],1) + '%',
                        }
                    }
                ],
                markLine: !isNaN(actualShare) ? {
                    symbol: ['none', 'none'],
                    label: { show: true, position: 'middle',
                        distance: 13,
                        afontWeight: 'bold',
                        color: '#111' ,
                        color: '#6b7280',
                        formatter: `Abstand ~${formatDE(distance,0)}%`,
                    },
                    lineStyle: { type: 'solid', width: 3, color: '#6b7280' },
                    data: [[
                        { coord: [userShare,   ''] },
                        { coord: [actualShare, ''] },
                    ]]
                } : {}
            }]
        };
        chartObjs[0].setOption(distanceChartOption);
        // console.log(`distance chart ${chartObjs[0]}  `);




        // chart #2 - forecast
        const forecastOption = {
            grid:    { top: 10, right: 60, bottom: 30, left: 100 },  // y-axis outset
            grid:    { top: 10, right: 20, bottom: 30, left:  20 },  // y-axis inset
            xAxis:   { type: 'value',
                       min: -3, max: 3,
                       interval: 0.5,
                     },
            yAxis:   {
                       type: 'category',
                       data: [quarter],
                        axisLabel: {
                            inside: true,     // label inside the grid area
                            align: 'left',    // aligns text to the left edge of each tick position
                            margin: -8,       // outdent
                            lineHeight: 16,
                            verticalAlign: 'middle',
                            padding: [-3, 0, 0, 0],
                        },
                    },
            series: [{
                type: 'bar',
                data: [forecast],
                itemStyle: { color: '#62929e' },
                label: { show: true, position: 'insideRight',
                    color: '#000', 
                    backgroundColor: 'rgba(255,255,255,0.8)', borderRadius: 4,

                    // fix - adapted padding - but only effective, if we add the "rich" property
                    padding: [3, 6, 1, 6],
                    rich: {},                    

                    // these did not help
                    // offset:  [0, 0],
                    // position: [0, '25%'],


                    // fontWeight: 'bold',
                    fontSize:   13,
                    lineHeight: 20,
                    

                    formatter: (val) => (isNaN(val.value) ? '' : formatDE(val.value) ) ,
                }
            }]
        };
        chartObjs[1].setOption(forecastOption);



        // chart #3 - consensus
        const consensusOption = {
            grid:    { top: 10, right: 60, bottom: 30, left: 100 },  // y-axis outset
            grid:    { top: 10, right: 20, bottom: 30, left:  20 },  // y-axis inset
            xAxis:   { type: 'value', min: -3, max: 3, interval: 0.5, axisLabel: { formatter: '{value}' } },
            yAxis:   {
                        type: 'category',
                        data: [quarter],

                        axisLabel: {
                            inside: true,     // label inside the grid area
                            align: 'left',    // aligns text to the left edge of each tick position
                            margin: -8,       // outdent

                            lineHeight: 16,
                            verticalAlign: 'middle',
                            padding: [-3, 0, 0, 0],
                        },

                     },
            series: [{
                type: 'bar',
                data: [consensus],
                itemStyle: { color: '#546a7b' },
                label: { show: true, position: 'insideRight',
                    color: '#000', 
                    backgroundColor:  'rgba(255,255,255,0.8)' , borderRadius: 4, 

                    // top padding smaller than bottom padding
                    padding: [2, 6],

                    // fix - adapted padding - but only effective, if we add the "rich" property
                    padding: [3, 6, 1, 6],
                    rich: {},                    

                    // these did not help
                    // offset:  [0, 0],
                    // position: [0, '25%'],


                    // fontWeight: 'bold',
                    fontSize:   13,
                    lineHeight: 20,
                    
                    formatter: (val) => (isNaN(val.value) ? '' : formatDE(val.value) ),
                }
            }]
        };
        chartObjs[2].setOption(consensusOption);
    }


    chartObjs.forEach(function (chartObj) {
        // chartObj.resize();
    });


    console.log(`pageLoaded() pg3 complete`)


    setTimeout(function() {

        console.log(`updateCharts() start`)
        updateCharts();
        chartObjs.forEach(function (chartObj) {
            chartObj.resize();
            // var inst = echarts.getInstanceByDom(el);
        });
        console.log(`updateCharts() stop`)

    }, 10);


});

"strict mode";

// declared and initialized in echart-config.mjs
//      param1, param2
//      myChart, dataObject too

let pageLoaded = (inst) => {

    const evt = new Event("change");

    let sbChange = (evt) => {
        let src = evt.srcElement;
        const chVal = src.value;
        refresh(myChart, dataObject);
        console.log(`sbChange ${chVal}`)

        simHist.push({ "sb_ch": chVal });
        simHistBG.value = JSON.stringify(simHist);
        // console.log(`simHistBG.value ${simHistBG.value}`)
    }

    param1Inp.onchange = sbChange


    //
    let chartDom = document.getElementById('chart_container');
    myChart = echarts.init(chartDom);

    optEchart && myChart.setOption(optEchart);
    console.log(`echart config and creation complete`)



    console.log(`pageLoaded() complete`)
}


window.addEventListener('load', pageLoaded, false);



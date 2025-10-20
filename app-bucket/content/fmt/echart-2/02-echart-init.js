"strict mode";

// declared and initialized in echart-config.mjs
//      param1, param2
//      myChart, dataObject too

let pageLoaded = (inst) => {

    const evt = new Event("change");

    let paramChange = (evt) => {
        let src = evt.srcElement;
        const chVal = src.value;
        refresh(myChart, dataObject);
        console.log(`   ${evt.srcElement.name} - new val  ${chVal}`)

        simHist.push({ "sb_ch": chVal });
        simHistInp.value = JSON.stringify(simHist);
        // console.log(`simHistInp.value ${simHistInp.value}`)
    }

    param1Inp.onchange = paramChange
    param2Inp.onchange = paramChange

    //
    let chartDom = document.getElementById('chart_container');
    myChart = echarts.init(chartDom);

    optEchart && myChart.setOption(optEchart);
    console.log(`pageLoaded() - echart config and creation complete`)

}


window.addEventListener('load', pageLoaded, false);



document.addEventListener('DOMContentLoaded', () => {

    function getTimeFlooredBy10Seconds() {
        const nowMillis = Date.now();
        const nowSeconds = Math.floor(nowMillis / 1000);
        const floored10Sec = Math.floor(nowSeconds / 10) * 10;
        return floored10Sec;
    }


    const userShareInput      = document.getElementById('userShareInput');
    const userShareSlider     = document.getElementById('userShareSlider');

    // Sync slider and input
    userShareInput.addEventListener('input', () => {
        userShareSlider.value = userShareInput.value;
        updateCharts();
    });


    userShareSlider.addEventListener('input', () => {
        userShareInput.value = userShareSlider.value;
        updateCharts();
    });
    userShareSlider.addEventListener('change', () => {
        userShareInput.value = userShareSlider.value;
        updateCharts();
    });




    let  simHist     = {};
    let  simHistInp  = document.getElementById("change_history_pg2");


    let  param1Inp   = document.getElementById("userShareInput");
    let  param1InpBG = document.getElementById("param1_pg2_bg");
    if (param1InpBG && param1InpBG.value !== "") {
        param1Inp.value = parseInt(param1InpBG.value);  // restore from before
        userShareSlider.value = parseInt(param1InpBG.value);  // restore from before
    }



    console.log(`init param1 ${param1Inp.value}, bg ${param1InpBG.value} `)





    const evt = new Event("change");

    let updateCharts = () => {}

    let paramChange = (evt) => {

        let src = evt.srcElement;
        const chVal = src.value;

        param1InpBG.value = chVal;

        // refresh(myChart, dataObject);

        console.log(`   ${evt.srcElement.name} - new val  ${chVal}`)

        const nm = src.name.trim();
        const entry = {}
        entry[nm] = chVal

        simHist[ getTimeFlooredBy10Seconds() ] = entry;
        simHistInp.value = JSON.stringify(simHist);

        console.log(`simHistInp.value ${simHistInp.value}`)
    }

    param1Inp.onchange = paramChange
    userShareSlider.onchange = paramChange




    userShareSlider.focus();

    console.log(`pageLoaded() pg2 complete`)




});




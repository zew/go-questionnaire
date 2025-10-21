document.addEventListener('DOMContentLoaded', () => {

    function getTimeFlooredBy10Seconds() {
        const nowMillis = Date.now();
        const nowSeconds = Math.floor(nowMillis / 1000);
        const floored10Sec = Math.floor(nowSeconds / 10) * 10;
        return floored10Sec;
    }


    // // Sync slider and input
    // userShareInput.addEventListener('input', () => {
    //     userShareSlider.value = userShareInput.value;
    //     updateCharts();
    // });


    // userShareSlider.addEventListener('input', () => {
    //     userShareInput.value = userShareSlider.value;
    //     updateCharts();
    // });
    // userShareSlider.addEventListener('change', () => {
    //     userShareInput.value = userShareSlider.value;
    //     updateCharts();
    // });




    const  simHist     = {};
    const  simHistInp  = document.getElementById("change_history_pg2");

    const  userShareSld   = document.getElementById('userShareSlider');
    const  userShareInp   = document.getElementById('userShareInput');

    const  userShareBG    = document.getElementById("param1_pg2_bg");

    if (userShareBG && userShareBG.value !== "") {
        userShareInp.value = parseInt(userShareBG.value);  // restore from before
        userShareSld.value = parseInt(userShareBG.value);  
    }

    console.log(`init param1 ${userShareInp.value}, bg ${userShareBG.value} `)





    const evt = new Event("change");

    let updateCharts = () => {}

    let paramChange = (evt) => {

        let src = evt.srcElement;
        const chVal = src.value;

        userShareBG.value = chVal;

        // refresh(myChart, dataObject);

        console.log(`   ${evt.srcElement.name} - new val  ${chVal}`)


        if (evt.srcElement.name=="userShareSlider") {
           userShareInp.value = userShareSld.value 
        } else {
           userShareSld.value = userShareInp.value  
        }


        const nm = src.name.trim();
        const entry = {}
        entry[nm] = chVal
        entry["userShare"] = chVal

        simHist[ getTimeFlooredBy10Seconds() ] = entry;
        simHistInp.value = JSON.stringify(simHist);

        // console.log(`simHistInp.value ${simHistInp.value}`)
    }

    userShareInp.onchange = paramChange
    userShareSld.onchange = paramChange




    userShareSld.focus();

    console.log(`pageLoaded() pg2 complete`)




});




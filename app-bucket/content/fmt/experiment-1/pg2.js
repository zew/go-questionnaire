document.addEventListener('DOMContentLoaded', () => {


    const  userShareSld   = document.getElementById('userShareSlider');
    const  userShareInp   = document.getElementById('userShareInput');
    const  userShareBG    = document.getElementById("param1_pg2_bg");

    if (userShareBG && userShareBG.value !== "") {
        // restore from before
        const  restored = parseFloat(userShareBG.value)
        userShareInp.value = restored;
        userShareSld.value = restored;
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
        recordChange("userShare", chVal);

    }

    // userShareInp.onchange = paramChange
    // userShareSld.onchange = paramChange
    userShareInp.oninput = paramChange
    userShareSld.oninput = paramChange




    userShareInp.focus();
    // userShareSld.focus();

    console.log(`pageLoaded() pg2 complete`)




});




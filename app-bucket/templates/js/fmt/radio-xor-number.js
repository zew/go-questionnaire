// inp1Id - radio input name
// inp2Id - number input - assigned to radio 6
//    any radio checking triggers the disabling of the number input - except radio 6
//    initial disabling NOT set 			...gr.Style.Desktop....


// function block isolates multiple instances
(function () {

    // must be var
    var inp1Id  = "{{.inp1}}";
    var inp2Id  = "{{.inp2}}";
    var radioOn = "{{.radioOn}}";

    console.log( `names for inp1-2 -${inp1Id}- -${inp2Id}-  radioOn -${radioOn}- `);

    var inp2   = document.getElementById(inp2Id);
    var radio6 = document.getElementById(radioOn);

    function checkHandler(evt) {
        // console.log("selected:", evt.target.value);
        if (evt.target.checked) {
            if (evt.target.id === radioOn){
                inp2.placeholder = "0.00";
                inp2.disabled = false;
            } else {
                inp2.placeholder = "";
                inp2.disabled = true;
            }
        }
    }




    // addEventListener is cumulative
    window.addEventListener("load", function (evt) {

        const selector = `input[type="radio"][name="${inp1Id}"]`
        // console.log(`selector ${selector}`)
        const radioList = document.querySelectorAll(selector);

        let anyChecked = false;
        for (let idx1 = 0; idx1 < radioList.length; idx1++) {
            const inpSrcRadio = radioList[idx1];

            inpSrcRadio.addEventListener('change', checkHandler);
            if (inpSrcRadio.checked){
                anyChecked = inpSrcRadio;
            }
            // console.log(`change listener assigned to ${inpSrcRadio.id} - ${inpSrcRadio.type}`);
        }


        function numberInput(evt) {
            radio6.checked = true;
        }
        inp2.addEventListener('input', numberInput);


        // init
        if (anyChecked) {
            const evtInit = new Event("change");
            anyChecked.dispatchEvent(evtInit);         
        }    


        
    });

})();

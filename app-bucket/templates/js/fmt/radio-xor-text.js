// inp1Id - radio  input name
// inp2Id - text   input - assigned to radio 6
//    any radio checking triggers the disabling of the text input - except radio 6
//    initial disabling NOT set 			...gr.Style.Desktop....


// function block isolates multiple instances
(function () {

    // must be var
    var inp1Id  = "{{.inp1}}";
    var inp2Id  = "{{.inp2}}";
    var radioOn = "{{.radioOn}}";


    console.log( `ids for inp1-2 -${inp1Id}- -${inp2Id}-  radioOn -${radioOn}- `);

    // declaring outer references before dom parse
    var inp1   = null;
    var inp2   = null;
    var radio6 = null;
    // console.log( `el  for radio6.name -${radio6.tagName}- `);

    function checkHandler(evt) {
        // console.log("selected:", evt.target.value);
        if (evt.target.checked) {
            if (evt.target.id === radioOn){
                // inp2.placeholder = "0.00";
                inp2.disabled = false;
            } else {
                // inp2.placeholder = "";
                inp2.disabled = true;
            }
        }
    }




    // addEventListener is cumulative
    window.addEventListener("load", function (evt) {

        // querying elements after full dom parse
        inp1   = document.getElementById(inp1Id);
        inp2   = document.getElementById(inp2Id);
        radio6 = document.getElementById(radioOn);

        // if (inp1) {
        //     console.log( `el  for inp1.name   -${inp1.tagName}-  -${inp1.name}- `);
        // }
        if (!inp2 || !radio6) {
            console.error( `missing dom elements inp2=${inp2Id} radioOn=${radioOn}` );
            return;
        }

        // preventing duplicate binding across repeated script blocks
        if (inp2.dataset.xorBound) {
            return;
        }
        inp2.dataset.xorBound = "1";

        // resolving shared name attribute from element or fallback
        const radioName = inp1 ? inp1.name : inp1Id;
        const selector = `input[type="radio"][name="${radioName}"]`
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

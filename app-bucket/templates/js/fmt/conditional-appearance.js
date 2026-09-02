// inp1Id - radio input name
//    any radio checking triggers the appearance of a second question below


// function block isolates multiple instances
(function () {

    // must be var
    var inp1Id    = "{{.inp1Id}}";
    var groupID1  = "{{.groupID1}}";
    var groupID2  = "{{.groupID2}}";

    var grp1 = document.getElementsByClassName(groupID1)[0]
    var grp2 = document.getElementsByClassName(groupID2)[0]

    console.log( ` inp1 -${inp1Id}-   groupID -${groupID1}- `);
    console.log( grp1 );
    console.log( grp2 );


    function checkHandler(evt) {
        // console.log("selected:", evt.target.value);
        if (evt.target.checked) {
            console.log("radio checked:", evt.target.value);
            if ( evt.target.checked ){
                grp1.style.display  = "grid";
                grp2.style.display  = "grid";
            } else {
                grp1.style.display  = "none";
                grp2.style.display  = "none";
            }
        }
    }




    // addEventListener is cumulative
    window.addEventListener("load", function (evt) {


        const selector = `input[type="radio"][name="${inp1Id}"]`
        console.log(`selector ${selector}`)
        const radioList = document.querySelectorAll(selector);


        for (let idx1 = 0; idx1 < radioList.length; idx1++) {
            const inpSrcRadio = radioList[idx1];
            inpSrcRadio.addEventListener('change', checkHandler);
            console.log(`change listener assigned to ${inpSrcRadio.id} - ${inpSrcRadio.type}`);
        }


        let anyChecked = false;
        for (let idx1 = 0; idx1 < radioList.length; idx1++) {
            const inpSrcRadio = radioList[idx1];
            if (inpSrcRadio.checked){
                anyChecked = inpSrcRadio;
                console.log(`\t  anyChecked ${anyChecked.id} `);
            }
        }



        // init
        if (anyChecked) {
            const evtInit = new Event("change");
            anyChecked.dispatchEvent(evtInit);         
        }    


        
    });

})();

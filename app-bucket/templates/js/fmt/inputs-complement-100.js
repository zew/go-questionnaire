// inpSrcId - radio input CSS class name
// inpDstId - group div
//  any radio checking triggers the group to become visible (display: grid)
//  initial invisibility set with 			gr.Style.Desktop.StyleBox.Display = "none"


// function block isolates multiple instances
(function () {

    // must be var
    var inp1Id = "{{.inp1}}";
    var inp2Id = "{{.inp2}}";
    var inp3Id = "{{.inp3}}";

    console.log( `names for inp1-3 -${inp1Id}- -${inp2Id}- -${inp3Id}- `);

    var inp1 = document.getElementById(inp1Id);
    var inp2 = document.getElementById(inp2Id);
    var inp3 = document.getElementById(inp3Id);


    function funcOnChange(evt) {
        console.log("selected:", evt.target.value);
        if (evt.target.value){
            inp2.value = 100 - parseInt(evt.target.value) ;
        } else {
            inp2.value = "";
        }
    }


    function funcOnChecked(evt) {
        console.log("selected:", evt.target.value);
        if (evt.target.checked){
            inp1.value = "";
            inp2.value = "";
            inp1.placeholder = "";
            inp2.placeholder = "";
            inp1.disabled = true;
        } else {
            inp1.disabled = false;
            inp1.placeholder = "0";
            // inp2.placeholder = "0";
        }
    }


    // addEventListener is cumulative
    window.addEventListener("load", function (evt) {

        // inp1.addEventListener('change', funcOnChange);
        inp1.addEventListener('input', funcOnChange);
        inp3.addEventListener('change', funcOnChecked);

        inp1.placeholder = "0";
        inp2.placeholder = "";

        if (inp3.checked) {
            const evtInit = new Event("change");
            inp3.dispatchEvent(evtInit);         
        }

    });

})();

// inpSrcId - radio input CSS class name
// inpDstId - group div
//  any radio checking triggers the group to become visible (display: grid)
//  initial invisibility set with 			gr.Style.Desktop.StyleBox.Display = "none"

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


// addEventListener is cumulative
window.addEventListener("load", function (evt) {

    inp1.addEventListener('change', funcOnChange);
    // console.log(`change listener assigned to ${inpSrcRadio.id} - ${inpSrcRadio.type}`);


    let anyChecked = false;
    // for (let idx1 = 0; idx1 < radioList.length; idx1++) {
    //     const inpSrcRadio = radioList[idx1];

    //     inpSrcRadio.addEventListener('change', checkHandler);
    //     if (inpSrcRadio.checked){
    //         anyChecked = inpSrcRadio;
    //     }
    //     // console.log(`change listener assigned to ${inpSrcRadio.id} - ${inpSrcRadio.type}`);
    // }

    // init
    if (anyChecked) {
        // const evtInit = new Event("change");
        // anyChecked.dispatchEvent(evtInit);         
    }

});







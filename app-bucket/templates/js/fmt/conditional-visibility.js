// inpSrcId - radio input CSS class name
// inpDstId - group div
//  any radio checking triggers the group to become visible (display: grid)
//  initial invisibility set with 			gr.Style.Desktop.StyleBox.Display = "none"

// must be var
var inp1Id = "{{.inpSrc}}";
var inp2Id = "{{.inpDst}}";

console.log( `names for src, dst -${inp1Id}- -${inp2Id}- `);

var inp1 = document.getElementsByClassName(inp2Id)[0];
console.log("id for dst ", inp1.id, inp1.type);



function checkHandler(evt) {
    // console.log("selected:", evt.target.value);
    if (evt.target.checked) {
        inp1.style.display = "grid";
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

    // init
    if (anyChecked) {
        const evtInit = new Event("change");
        anyChecked.dispatchEvent(evtInit);         
    }

});







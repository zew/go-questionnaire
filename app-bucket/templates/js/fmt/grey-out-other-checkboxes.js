// must be var
var coreId = "{{.core_id}}";
console.log( `core ID is  -${coreId}- `);


var trigger = document.getElementById(`${coreId}1`);
console.log("id for trigger ", trigger.id, trigger.type);



function checkHandler(evt) {
    console.log(`\tchanged ${evt.target.name}`);

    for (let idx = 0; idx < 10; idx++) {
        if (idx < 2) {
            continue
        }
        var dst = document.getElementById(`${coreId}${idx}`);
        // console.log(`\t  idx ${idx}`);
        if (dst) {
            dst.disabled = trigger.checked;
            console.log(`\t  toggling ${dst.name}`);
        }
    }

    var free = document.getElementById(`${coreId}free`);
    free.disabled = trigger.checked;


    console.log(`\thandler success`);
}


// addEventListener is cumulative
window.addEventListener("load", function (evt) {


    trigger.addEventListener('change', checkHandler);

    const evtInit = new Event("change");
    trigger.dispatchEvent(evtInit);         

});







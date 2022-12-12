{
    // non global block

    // unused;
    // because the init of range, display, noanswer 
    // is done using server side HTML generation

    let range = document.getElementById("{{.inputName }}");

    let initRange = (inst) => {
        const evt = new Event("input");
        range.dispatchEvent(evt);
    }

    // on windows load event
    // we will trigger an oninput event (see above);
    // this will init the range-input
    window.addEventListener('load', initRange, false);


    // console.log("JS tpl 'rangeAuto.js' successfully added for {{.inputName }}");
}


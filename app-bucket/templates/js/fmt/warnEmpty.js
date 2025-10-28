function emptyUserShareWarning(evt) {
    
    if (evt.type == "submit") {
        // console.log(`checking named ${evt.target.name}-${evt.target.type} - sum at ${sum}`);
        const elUserShare = document.getElementById("userShareInput");
        const vl = elUserShare.value;
        if(vl === ""){
            /** alert(" curly-brace-open curly-brace-open  .msg2 curly-brace-close curly-brace-close "); */
            // alert("Bitte geben Sie Ihre Einschätzung ab.");
            let doContinue = window.confirm("Bitte geben Sie Ihre Einschätzung ab.");
            if (doContinue) {
                return true;
            }
            // not only return false - but also preventDefault()
            if (evt.preventDefault) {
                evt.preventDefault(); 
            }
            return false;
        }
    }

}







// addEventListener is cumulative
window.addEventListener("load", function (evt) {
    // const evtInit = new Event("change");
    // inp01.dispatchEvent(evtInit);
    document.forms.frmMain.addEventListener('submit', emptyUserShareWarning);

});

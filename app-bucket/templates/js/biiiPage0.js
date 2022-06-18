function disableSiblings(disable) {

    var elements = document.forms.frmMain.elements;
    for (var i = 0, element; element = elements[i++];) {

        if (element.type === "hidden") {
            continue;
        }

        if (element.type !== "checkbox") {
            continue;
        }

        if (element.name.substring(0, 4) == "q03_") {

            if (element.name == "q03_no_investing") {
                continue;
            }

            if (disable) {
                element.disabled = true;
                element.checked = false;
            } else {
                element.disabled = false;
            }

        }

    }
    return true;
}


function myOnchange(event) {
    let checked = document.forms.frmMain.q03_no_investing[0].checked;
    checked = event.target.checked
    disableSiblings(checked);
    return true;
}



if (document.forms.frmMain.q03_no_investing) {
    document.forms.frmMain.q03_no_investing[0].addEventListener('change', myOnchange);
    console.log("handlers assigned");
}

// addEventListener is cumulative
window.addEventListener("load", function (event) {
    disableSiblings(document.forms.frmMain.q03_no_investing[0].checked);
});
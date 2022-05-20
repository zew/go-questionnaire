/*

    Three checkboxes hide and show a subset.

    Each must be selected with [0],
    to separate it from the hidden sister input

    document.forms.frmMain.q39_western_europe[0]

*/

function showHideSub(idx,show) {

    var elements = document.forms.frmMain.elements;
    for (var i = 0, element; element = elements[i++];) {

        if (element.type === "hidden") {
            continue;
        }

        if (element.type !== "checkbox") {
            continue;
        }

        let prefix = "q39_sub" + String(idx)
        // console.log(`going for prefix ${prefix}`);
        if (element.name.substring(0, 8) == prefix) {

            // console.log(`found prefix ${prefix} - ${element.name}`);

            if (show) {
                // element.style.display = ""; // not block
                let anc = element.closest(".grid-item-lvl-1");
                anc.style.display = "";
            } else {
                // element.style.display = "none";
                let anc = element.closest(".grid-item-lvl-1");
                anc.style.display = "none";
            }

        }

    }
    return true;
}


function q39_3(event) {
    let checked = document.forms.frmMain.q39_western_europe[0].checked;
    checked = event.target.checked
    showHideSub(3, checked);
    return true;
}

function q39_4(event) {
    // let radio1Val = document.forms.frmMain.q39_western_europe.value;
    let checked = document.forms.frmMain.q39_western_europe[0].checked;
    checked = event.target.checked
    showHideSub(4, checked);
    return true;
}

function q39_5(event) {
    let checked = document.forms.frmMain.q39_western_europe[0].checked;
    checked = event.target.checked
    showHideSub(5, checked);
    return true;
}


if (document.forms.frmMain.q39_western_europe) {
    document.forms.frmMain.q39_western_europe[0].addEventListener('change', q39_3 );
    document.forms.frmMain.q39_middle_east_eur[0].addEventListener('change', q39_4 );
    document.forms.frmMain.q39_worldwide[0].addEventListener('change', q39_5 );
    console.log("handlers assigned");
}

// addEventListener is cumulative
window.addEventListener("load", function (event) {
    showHideSub(3, document.forms.frmMain.q39_western_europe[0].checked);
    showHideSub(4, document.forms.frmMain.q39_middle_east_eur[0].checked);
    showHideSub(5, document.forms.frmMain.q39_worldwide[0].checked);
});
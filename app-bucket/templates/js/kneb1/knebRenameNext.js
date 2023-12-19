// addEventListener is cumulative
window.addEventListener("load", function (event) {
    const btns  = document.getElementsByName("submitBtn");
    btns.forEach( btn => {
        if (btn.value === "next") {
            btn.innerHTML = "<b>&nbsp;&nbsp; Werte speichern und weiter &nbsp;&nbsp;</b>";
        }
    });
});
// addEventListener is cumulative
window.addEventListener("load", function (event) {
    const btns  = document.getElementsByName("submitBtn");
    btns.forEach( btn => {
        if (btn.value === "prev") {
            btn.style.display = "block";
            btn.innerHTML = "<b>&nbsp;&nbsp; Zurück zum Tool &nbsp;&nbsp;</b>";
        }
    });
});
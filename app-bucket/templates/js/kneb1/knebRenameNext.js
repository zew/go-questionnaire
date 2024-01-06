const msg = "{{.msg}}";
// alert(msg);

// addEventListener is cumulative
window.addEventListener("load", function (event) {
    const btns  = document.getElementsByName("submitBtn");
    btns.forEach( btn => {
        if (btn.value === "next") {
            btn.innerHTML = `<b>&nbsp;&nbsp; ${msg} &nbsp;&nbsp;</b>`;
        }
    });
});
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

    try {
        // const page = document.getElementsByClassName("pg31");
        const simTool = document.getElementsByClassName("panel-flex-container")[0];
        const pageDiv = simTool.parentElement.parentElement.parentElement.parentElement.parentElement;
        if (pageDiv) {
            pageDiv.style.marginTop = "0.1rem !important";
            pageDiv.classList.add("echart-mobile");
            console.log(`pageDiv is ${pageDiv.tagName} - ${pageDiv.classList}`);
        } else {
            // console.log(`par not found`);
        }
    } catch (err) {
        console.log(err)
    }


});
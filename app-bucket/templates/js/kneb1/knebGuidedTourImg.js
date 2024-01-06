const appPrefix = "{{.appPrefix}}";
const treatIdx  = "{{.treatIdx}}";
const pageIdx   = "{{.pageIdx}}";


window.onload = function() {
    try {
        // const strURL = (window.location.href).toLowerCase();
        // const url = new URL(strURL);
        // const treatIdx = url.searchParams.get("treatIdx"); // 0 or 1
        let treatment = "ntrl";
        if (treatIdx==="1") {
            treatment = "fin";
        }
        console.log(`treatIdx ${treatIdx}  - treatment ${treatment}  - pageIdx ${pageIdx}`);
        
        const img = document.createElement("img");
        const imgURL = `./${appPrefix}/doc/kneb1/slide-show/out/${treatment}/0${pageIdx}.png` ;
        console.log(`img url ${imgURL}`)

        img.setAttribute("src", imgURL)
        img.setAttribute("alt", "")
        // img.style.display = "inline-block"
        img.style.margin = "0 auto";
        img.style.maxHeight = "calc(100vh - 8rem)";

        // const anchor = document.getElementById("anchor");
        const anchor = document.getElementsByClassName("grid-container")[0];
        anchor.appendChild(img);


        try {
            const pg = anchor.parentElement;
            console.log(`hopefully the page:`, pg )
            pg.style.marginTop = 0;
            // pg.style.maxWidth = 'unset';
        } catch (err) {
            console.error(`could not step up to page - ${pg}`);
            console.error(err);
        }

        try {
            const btn = document.getElementsByName("submitBtn")[0];
            btn.style.animation = `guidedTourNextButtonFadeIn 2s`
        } catch (err) {
            console.error(`could not animate the next btn - ${btn}`);
            console.error(err);            
        }


    } catch (err) {
        console.error(err);
    }




}		

const appPrefix = "{{.appPrefix}}";
const treatIdx  = "{{.treatIdx}}";
const pageIdx   = "{{.pageIdx}}";

function createImgEl(treatment, mobile){

    const img = document.createElement("img");
    let imgURL = `./${appPrefix}/doc/kneb1/slide-show/out/${treatment}/0${pageIdx}.png` ;
    if (mobile) {
        imgURL = `./${appPrefix}/doc/kneb1/slide-show/out/${treatment}-mobile/0${pageIdx}.png` ;
    }
    console.log(`img url ${imgURL}`)

    img.setAttribute("src", imgURL)
    img.setAttribute("alt", "")

    if (!mobile) {
        img.classList.add("img-guided-tour-desktop")      
    } else {
        img.classList.add("img-guided-tour-mobile")      
    }

    // img.style.display = "inline-block"
    img.style.margin = "0 auto";
    img.style.maxHeight = "calc(100vh - 8rem)";

    return img;
}


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
        
        const img1 = createImgEl(treatment, false);
        const img2 = createImgEl(treatment, true );


        // const anchor = document.getElementById("anchor");
        const anchor = document.getElementsByClassName("grid-container")[0];
        anchor.appendChild(img1);
        anchor.appendChild(img2);


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

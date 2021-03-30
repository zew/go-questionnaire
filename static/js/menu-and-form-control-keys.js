// non essentiell JS helpers
function keyControls(e) {

    // [enter] key opens  2nd level menu, just as space bar does
    if (e.key === "Enter") {
        var menuCheckbox = document.getElementById("mnu-1st-lvl-toggler");
        var isFocused = (document.activeElement === menuCheckbox);
        if (isFocused) {
            menuCheckbox.checked = true;
            console.log("key listener ENTER fired");
        }
    }

    // [esc]   key closes 2nd level menu, if its expanded
    if (e.key === "Escape") {
        document.getElementById("mnu-1st-lvl-toggler").checked = false;


        // ExcelDB: hide all control-menu-2 
        // var mnu2s = document.getElementsByClassName("control-menu-2");
        // for (var i = 0; i < mnu2s.length; i++) {
        // 	mnu2s[i].style.display = 'none';
        // }
        // console.log("key listener ESC fired");
    }

    // [enter] on inputs transformed into focus next input.
    // Sending events to inputs is security forbidden.
    // We find the next element and focus() it.
    //
    // TEXTAREA: SHIFT+ENTER mode is impossible on mobile - 
    // thus we cannot inlude TEXTAREA into the func	 
    // 
    //	optionally restrict to certain user agens: && /Android/.test(navigator.userAgent)
    if (e.key === "Enter") {

        var isShift = !!e.shiftKey; // convert to boolean
        if (isShift) {
            console.log("let SHIFT ENTER pass");
            return;
        }

        var el = document.activeElement;

        // skip for <input type=submit>  and <button>... 
        if ((el.tagName == "INPUT" && el.type != "submit") || el.tagName == "SELECT") {

            e.preventDefault();
            var nextEl = null;


            if (false) {
                // first method for finding next element:
                // adding succinct tab indize
                // then taking current tab index and incrementing it
                var elements = el.form.elements;
                var cntr = 1;
                for (var i = 0, lpEl; lpEl = elements[i++];) {
                    if (lpEl.type !== "hidden" && lpEl.type !== "fieldset") {
                        lpEl.tabIndex = cntr;
                        cntr++;
                        // console.log("tab index", element.name, " to ", i);
                    } else {
                        // console.log("SKIPPING tab index ", element.name, " - ", i);
                    }
                }
                var nextTabIndex = el.tabIndex + 1;
                nextEl = el.form.elements[nextTabIndex];
                if (nextEl && nextEl.focus) nextEl.focus();
            }


            // second method: simply follow the form elements order
            var found = false;
            if (el.form) {
                for (var i = 0, lpEl; lpEl = el.form.elements[i++];) {
                    if (lpEl.type !== "hidden" && lpEl.type !== "fieldset") {
                        if (found) {
                            nextEl = lpEl;
                            console.log(`found next	   ${lpEl.name} type ${lpEl.type} at `, i);
                            break;
                        }
                        if (el === lpEl) {
                            console.log(`found current ${lpEl.name} type ${lpEl.type} at `, i);
                            found = true;
                        }
                        // console.log("iterating form elements", element.name, " to ", i);
                    } else {
                        // console.log("iterating form elements - skipping ", element.name, " - ", i);
                    }
                }
            }
            if (nextEl && nextEl.focus) nextEl.focus();


            if (nextEl) {
                // console.log("key listener ENTER - transformed into TAB:", el.tagName, el.name, nextEl.tagName, nextEl.name );
            } else {
                // console.log("key listener ENTER - transformed into TAB:", el.tagName, el.name, " next element not found" );
            }

        } else {
            // console.log("key listener ENTER on tagname:", el.tagName, el.name );
        }
    }

}

// click outside menu closes it
function outsideMenu(event) {
    var elNav = document.getElementsByTagName('nav');
    var nav = elNav[0];
    // event.preventDefault();
    if (!nav.contains(event.target)) {
        // console.log('click outside menu');
        document.getElementById("mnu-1st-lvl-toggler").checked = false;
    }
}

// click on nde-2nd-lvl pulls up mnu-3rd-lvl
//
// we would love to change li.nde-2nd-lvl::before
// into an upward arrow too, but pseudo elements 
// cannot be selected / styled via javascript
var closeLevel3 = function () {
    for (let i = 0; i < this.children.length; i++) {
        if (this.children[i].tagName == "UL") {
            var el = this.children[i];
            var style = window.getComputedStyle(el);
            if (style.opacity < 0.5) {
                el.classList.remove("mnu-3rd-lvl-pull-up");  // remove means *show* ;this is the show / init branch - opacity 0 and growing
            } else {
                el.classList.add("mnu-3rd-lvl-pull-up");	 // add	means *hide*
            }
            break;
        }
    }
};


// window.onload = ...   is *not* cumulative
// window.onload = function () {
//     //    
// };
// 
// addEventListener is cumulative
window.addEventListener("load", function (event) {

    document.addEventListener("keydown", keyControls, false);
    console.log("global key listener registered");


    var html = document.body.parentNode;
    html.addEventListener("touchstart", outsideMenu, false);
    html.addEventListener('click', outsideMenu, false);

    var nodesLvl2 = document.getElementsByClassName("nde-2nd-lvl");
    for (var i = 0; i < nodesLvl2.length; i++) {
        nodesLvl2[i].addEventListener('click', closeLevel3, false);
    }
    console.log("outsideMenu and closeLevel3 registered");


    // focus on first visible input
    // TODO: first input after p.error-block
    var elements = document.forms.frmMain.elements;
    for (var i = 0, element; element = elements[i++];) {
        if (element.type !== "hidden") {
            // pbu 2021-02: dropped - first control is sometimes too far down - init view scrolls down
            // element.focus();
            // console.log("focus on form main input number", i, element.name);
            break;
        }
    }

});

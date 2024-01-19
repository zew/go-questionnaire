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
    // thus we cannot include TEXTAREA into the func
    //
    //	optionally restrict to certain user agents: && /Android/.test(navigator.userAgent)
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
                            // console.log(`found next	   ${lpEl.name} type ${lpEl.type} at `, i);
                            break;
                        }
                        if (el === lpEl) {
                            // console.log(`found current ${lpEl.name} type ${lpEl.type} at `, i);
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

    try {
        document.addEventListener("keydown", keyControls, false);
        // console.log("global key listener registered");

        var html = document.body.parentNode;
        html.addEventListener("touchstart", outsideMenu, false);
        html.addEventListener('click', outsideMenu, false);

        var nodesLvl2 = document.getElementsByClassName("nde-2nd-lvl");
        for (var i = 0; i < nodesLvl2.length; i++) {
            nodesLvl2[i].addEventListener('click', closeLevel3, false);
        }
        // console.log("outsideMenu and closeLevel3 registered");

    } catch (err) {
        console.error(`error in registering 'keycontrol' 'outsideMenu' or 'closeLevel3' `);
        console.error(err);
    }



    var invalidInputs = false; // invalid by HTML5
    var invalidFields = document.querySelectorAll("form :invalid");  // excluding invalid form itself
    for (var i = 0; i < invalidFields.length; i++) {
        /*  first pages with first element after long text
                => scrolls down
            preventScroll supported only since 2018
         */
        try {
            invalidFields[i].focus({
                preventScroll: true
            });
        } catch (error) {
            // forgoing initial focussing
        }
        // console.log(`focus on first invalid input ${invalidFields[i].name}`);
        invalidInputs = true;
        break;
    }

    var invalidServerFields = document.querySelectorAll(".error-block-input"); // invalid by server rules
    var firstErrMsgTop = 0;
    if (invalidServerFields.length > 0) {
        firstErrMsgTop = invalidServerFields[0].getBoundingClientRect().y
        // console.log(`.error-block-input found at ${topPosOfErr}`);
    }


    if (!invalidInputs) {
        // focus on first visible input
        var elements = document.forms.frmMain.elements;
        for (var i = 0, el; el = elements[i++];) {
            if (el.type === "hidden") {
                continue;
            }

            if (firstErrMsgTop > 0 && el.getBoundingClientRect().y < firstErrMsgTop) {
                // console.log(`.error-block-input found ${element.getBoundingClientRect().y} < ${topPosOfErr}`);
                continue;
            }

            if (el.type === "range") {
                // if the first input is a range;
                // then we dont want to focus - because otherwise hitting PageDown
                // would not scroll down the page, but change the range-slider thumb
                break;
            }

            /*  first pages with first element after long text
                    => scrolls down
                preventScroll supported only since 2018
                */
            try {
                el.focus({
                    preventScroll: true
                });
            } catch (error) {
                // forgoing initial focussing
            }
            // console.log(`focus on ${i}th input ${element.name} of form main`);
            break;

        }
    }


});


/**
 * We try to control the browser back buttons.
 * 
 * We cannot disable them.
 * 
 * We cannot distinguish them from regular page load wie browser address bar or clicking a link
 * 
 * We want to prevent the "Confirm form submit" message, when going back to previous post page.
 * 
 * The previous/next browser buttons show the page from the 
 *    back-forward cache   (https://web.dev/articles/bfcache)
 * 
 * This back-forward cache is *separate* 
 * from teh http cache (https://web.dev/articles/http-cache)
 * 
 * We disable the http cache by responding the header
 *     Cache-Control: no-cache="Set-Cookie"
 * 
 * We cannot disable the back-forward cache.
 * It is a version of the HTML-page completely disconnected from the server
 * We dont get any notice on the server, if a back-forward page is shown.
 * 
 * Suppose we are on page 7.
 * Suppose we press 'previous', 'previous', 'previous' to a cached view of page 4.
 * Support now, we click on the "Next page" button on page 4.
 * Then the server will respond with page 8,
 * because he does not know about the 'previous'
 * 
 * window event "beforeunload" is useless for our purpose.
 * 
 * window event "popstate"     is useless because its selective:
 * Quote from  the docs:
 *      Note that just calling history.pushState() or history.replaceState() 
 *      won't trigger a popstate event. 
 *      The popstate event will be triggered by doing a browser action 
 *      such as a click on the back or forward button 
 *      or calling history.back() or -forward() in JavaScript.
 * 
 * We can *only* prevent the "Confirm post submit" message on 'previous'.
 * We do this by caling  window.history.replaceState(...) for every window event "pageshow".
 * See below.
 * We dont change the browser URL.
 *   urlPathSuffix remains "/".
 * This leaves the URL unchanged, but still prevents the "form resubmit" question.
 * 'previous' now simply goes back in the back-forward cache
 * or sometimes even remains on the same page;
 * effectively neutering 'previous'.
 * The behaviour may depend upon whether is identical to golang "URLPathPrefix".
 * Development config is URLPathPrefix = "/survey".
 * Production  config is URLPathPrefix = "/".
 * 
 */



// const urlPathSuffix = "/replaced-state";
const urlPathSuffix = "/";

const manageBrowserHistory = (evt) => {

    if (evt.persisted) {
        console.log('restored from bfcache.');
    } else {
        // default; not worth reporting
        // console.log('loaded without bfcache.');
    }

    // get array of PerformanceEntry  - https://developer.mozilla.org/en-US/docs/Web/API/PerformanceEntry
    let perfEntries = performance.getEntriesByType("navigation");
    if ( perfEntries.length > 0 ) {

        if (perfEntries[0].type !== "navigate") {
            // "navigate", "reload", "back_forward" or "prerender"
            console.log(`perfEntries[0].type -${perfEntries[0].type}-`); 
        }
        if (perfEntries[0].initiatorType !== "navigation") {
            console.log(`perfEntries[0].initiatorType -${perfEntries[0].initiatorType}-`);
        }

        if (perfEntries[0].name !== window.location.href ) {
            let a = window.location.href;
            let b = perfEntries[0].name;
            a = a.trim();
            b = b.trim();

            if (  a.endsWith(urlPathSuffix) ) {
                a = a.substring(0, a.length - urlPathSuffix.length );
                // console.log(`urlPathSuffix cut`);
            }
            if ( a === b ) {
                console.log(`case0 -${a}-`);
            } else if (  a.startsWith(b) ) {
                // const pos = a.indexOf(b);
                const dff = a.substring(b.length);
                console.log(`case1 -${b}- vs -...${dff}-`);
            } else if (  b.startsWith(a) ) {
                const dff = b.substring(a.length);
                console.log(`case2 -${a}- vs -...${dff}-`);
            } else {
                console.log(`case3 -${a}- vs -${b}-`);
            }
        }
        if (perfEntries[0].redirectCount !== 0 ) {
            console.log(`perfEntries[0].redirectCount -${perfEntries[0].redirectCount}-`);
        }
        // console.log(perfEntries[0]);
    }

    /*
        window.history only reports its size and not much else
        you use its methods   back(), forward(), go(), pushState(), replaceState()
     */
    if (window.history && window.history.state !== null) {
        console.log(`hist size ${window.history.length} - state ${window.history.state}`   );
        console.log(`state`,window.history.state);
    }

    // replace *current* history state.
    //   first argument 'state' is saved as 'history.state' for later read operations on the history
    //   title - does not what it says, its simply unused
    //   location - changes relative URL of current own page, not URL of any other pages. 
    //        current own page URL is changed, but no reload occurs
    //        .back() and .forward() show changed URL
    //        .back() and .forward() trigger popstate, but no load() nor unload()
    //      use for changing query string - not full URL
    window.history.replaceState(
        {
            state1: 'stateval1',
            state2: 'stateval2',
        },
          document.getElementsByTagName("title"),
          urlPathSuffix,
    );


}
window.addEventListener( "pageshow", manageBrowserHistory);









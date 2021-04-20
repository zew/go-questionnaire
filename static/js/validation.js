/* 
    using a factory function
    https://isamatov.com/encapsulation-in-javascript-es6/
    first concept

*/
function Validator(argForm) {

    this.PublicVar = "we dont want public vars"

    let form = argForm;

    // internet suggests three modes to accomplish suppression of builtin bubbles
    // see if() branches below
    let suppressBuiltinBubbles = 3;  
    
    let onFocusRemove          = false; // unexported parameter
    let onInputRemove          = true;
    let onInputShowAndRemove   = false;
    let lockFocus              = false;


    // default is showing custom popups for every faulty input
    // onlyOne changes this
    let onlyOne = true;


    // custom popups are sized and positioned relative 
    // to the input's parent or grandparent.
    // Parent or grandparent provide us with a reliable width 
    // that does not overflow the screen width;
    // unexported parameter
    let attachGrandparent = true;


    this.SetOnInputRemove = function(newVal) {
        onInputRemove = newVal;
    }

    this.SetOnInputShowAndRemove = function(newVal) {
        onInputShowAndRemove = newVal;
    }

    this.SetLockFocus = function(newVal) {
        lockFocus = newVal;
    }

    this.SetOnlySingleCustomPopup = function(newVal) {
        onlyOne = newVal;
    }

    // "exporting" the func
    //  keeping the internal version, because internal callers have another 'this'
    this.ShowCustomPopup = showPopup;


    // if any input elements are not clean yet 
    //   => dont show any compound errors
    this.IsCleanForm = function(event) {

        let frmLoc = null;
        if (event.target.tagName == "FORM") {
            frmLoc = event.target;
        } else {
            frmLoc = event.target.form;
        }

        try {
            if (!frmLoc.checkValidity()) {
                return false;
            }
        } catch (error) {
            logFn("Exception: isCleanForm() was fired for non-form-element; event.target: ", event.target.tagName)
        }
        return true;
    }


    function hasPopup(el) {
        var elErrors = el.parentNode.querySelectorAll(":scope > .popup-invalid-anchor");
        if (attachGrandparent) {
            elErrors = el.parentNode.parentNode.querySelectorAll(":scope > .popup-invalid-anchor");
        }
        for (var i = 0; i < elErrors.length; i++) {
            // console.log(`found-a ${i + 1}of${elErrors.length} - oldID${oldChild.getAttribute('id')} `);
            return true;
        }
        return false;
    }


    // removing previous message from element el
    function clearPopup(el) {
        var elErrors = el.parentNode.querySelectorAll(":scope > .popup-invalid-anchor");
        if (attachGrandparent) {
            elErrors = el.parentNode.parentNode.querySelectorAll(":scope > .popup-invalid-anchor");
        }
        for (var i = 0; i < elErrors.length; i++) {
            var oldChild = elErrors[i].parentNode.removeChild(elErrors[i]);
            // console.log(`removed-a ${i + 1}of${elErrors.length} - oldID${oldChild.getAttribute('id')} `);
        }
    }


    // removing any previous custom messages
    function clearAllPopups() {
        var errorMessages = form.querySelectorAll(".popup-invalid-anchor");
        for (var i = 0; i < errorMessages.length; i++) {
            var oldChild = errorMessages[i].parentNode.removeChild(errorMessages[i]);
            // console.log(`removed-b ${i + 1}of${errorMessages.length} - oldID${oldChild.getAttribute('id')} `);
        }
    }

    // clearing and re-creating a custom message 
    // right-beside or -below DOM element el
    function showPopup(el, msg, overrideCheckValidity) {

        if (!el) {
            console.log("flagInvalid() el not defined - return ");
            return;
        }

        if (msg === undefined) {  // typeof msg == "undefined"
            msg = el.dataset.validation_msg
            if (msg === undefined) {
                msg = el.validationMessage // not localized, too mathematical
            }
        }

        if (onlyOne) {
            clearAllPopups();
        } else {
            clearPopup(el);
        }

        if (!el.checkValidity() || overrideCheckValidity === true) {
            var parent = el.parentNode;
            if (attachGrandparent) {
                parent = el.parentNode.parentNode;
            }
            // el.validationMessage is mathematical has is always in browser local
            parent.insertAdjacentHTML(
                "beforeend",
                `<div class='popup-invalid-anchor'  id='err-${el.getAttribute('name')}' >
                    <div class='popup-invalid-content'>
                    ${msg}
                    </div>
                </div>`
            );
        }

    }

    // for onsubmit
    // for each invalid input element of a form
    // a custom popup message is displayed right-next or -below
    function onSubmitCustomPopupsForInvalids(event) {

        clearAllPopups();

        // insert new messages at the end of parent
        // `this` to select descendents of <form> - excluding invalid <form> itself 
        var invalidFields = this.querySelectorAll(":invalid");
        for (var i = 0; i < invalidFields.length; i++) {
            showPopup(invalidFields[i]);
            if (onlyOne) {
                break;
            }
        }

        // focus first invalid field
        if (invalidFields.length > 0) {
            invalidFields[0].focus();
        }

        if (invalidFields.length > 0) {
            return false;
        }
        return true;
    }





    this.ValidateFormWithCustomPopups = function() {


        if (suppressBuiltinBubbles == 1) {
            // => form.submit() no longer works; only submit buttons clicks still effect a submission
            form.addEventListener(
                "invalid",
                function (event) {
                    console.log("form invalid: ", event.target.getAttribute("name"), " - default prevented");
                    event.preventDefault();
                },
                true
            );
        }


        if (suppressBuiltinBubbles == 2) {
            // => form.submit() no longer works; only submit buttons clicks still effect a submission
            var inputs = form.querySelectorAll("input[type=number]");
            for (var i = 0; i < inputs.length; i++) {
                var inp = inputs[i];
                var funcInv = function (event) {
                    console.log("input invalid: ", event.target.getAttribute("name"), " - default prevented");
                    event.preventDefault();
                };
                inp.addEventListener("invalid", funcInv, true);
            }
        }


        if (suppressBuiltinBubbles == 3) {

            // disable form validation
            // form.submit() validation disabled
            // form.submit() works and still goes through
            form.setAttribute("novalidate", true);


            // "re-enable" validation of inputs using an explicit event handler
            form.addEventListener(
                "submit",
                function (event) {
                    if (!this.checkValidity()) {
                        var name = event.target.getAttribute("name");
                        console.log(`prevented submitting form ${name}: invalid inputs`);
                        // emulating cancellation of form.submit() on invalid input
                        event.preventDefault();
                    }
                },
                true
            );
            // form submit now stalls on invalid
            // but without any bubbles nor any other messages
            console.log(`suppressBuiltinBubbles complete`);
        }

        // add custom popups
        form.addEventListener(
            "submit",
            onSubmitCustomPopupsForInvalids,
            true
        );
        console.log(`on submit: custom popups for invalids attached`);


    }



    // showing custom popups on form submit is too late?
    //   => show them on blur
    this.ShowPopupOnBlurOrInput = function() {

        // if we dont apply ValidateFormWithCustomPopups(),
        // then we still need this for standalone functionality
        form.setAttribute("novalidate", true);

        var funcReport = function (event) {
            // event.target.reportValidity();
            if (onInputRemove && event.type == "input") {
                if (event.target.checkValidity()) {
                    clearPopup(event.target);
                }
            } else {
                showPopup(event.target);
            }
            var lgMsg = "blur";
            if (onInputShowAndRemove || onInputRemove) {
                lgMsg = "blur+input";
            }
            console.log(`  ${lgMsg} inp.reportValidity() ${event.target.getAttribute('name')} ${event.target.checkValidity()}`);
            if (lockFocus) {
                if (event.type == "blur") {
                    if (!event.target.checkValidity()) {
                        event.target.focus();
                        console.log(`  blur focus reclaimed ${event.target.getAttribute('name')}`);
                    }
                }
            }
        };

        var inputs = form.querySelectorAll("input[type=number]");
        for (var i = 0; i < inputs.length; i++) {
            var inp = inputs[i];
            inp.addEventListener("blur", funcReport);  // blur does not bubble up
            if (onInputShowAndRemove || onInputRemove) {
                inp.addEventListener("input", funcReport);
            }

            if (onFocusRemove) { // remove on entering input
                var removeOnEntering = function (event) {
                    clearPopup(event.target);
                };
                inp.addEventListener("focus", removeOnEntering);    
            } else {
                var flagOnEntry = function (event) {
                    if (!event.target.checkValidity()) {
                        // console.log(`  show popup on focus -  ${event.target.name}`)
                        showPopup(event.target);
                    }
                };
                inp.addEventListener("focus", flagOnEntry);
            }
            var lgMsg = "blur";
            if (onInputShowAndRemove || onInputRemove) {
                lgMsg = "blur+input";
            }

            var logLen = 1
            if (i < logLen || i > (inputs.length - 1 - logLen)) {
                console.log(`     ${lgMsg} handler added to ${inp.getAttribute('name')}`);
            }
            if (i == logLen) {
                console.log(`          ...`);
            }
        }
    }



}

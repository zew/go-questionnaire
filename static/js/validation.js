/* 
    using a factory function
    https://isamatov.com/encapsulation-in-javascript-es6/
    first concept

*/
function Validator(argForm) {

    this.publicVar = "we dont want public vars"

    let form = argForm;

    let suppressBuiltinBubbles = 3;  // internet suggests three modes to accomplish suppression of builtin bubbles
    
    let checkOnInput        = true;
    let reclaimFocus        = true;
    let removeBubbleOnEntry = false;


    /* bubble messages are positioned relative to parent's parent
        providing us with a reliable width that does not overflow;
    */    
    let attachOuterOuter = true;


    this.SetCheckOnInput = function(newVal) {
        checkOnInput = newVal;
    }

    this.SetReclaimFocus = function(newVal) {
        reclaimFocus = newVal;
    }

    // "exporting" the func
    //  keeping the internal version, because internal callers have another 'this'
    this.ShowBubblePopup = showBubblePopup;



    // removing any previous custom messages
    function clearAllPreviousBubbles() {
        var errorMessages = form.querySelectorAll(".bubble-invalid-anchor");
        for (var i = 0; i < errorMessages.length; i++) {
            var oldChild = errorMessages[i].parentNode.removeChild(errorMessages[i]);
            // console.log(`removed-b ${i + 1}of${errorMessages.length} - oldID${oldChild.getAttribute('id')} `);
        }
    }

    // clearing and re-creating a custom message right-beside or -below DOM element el
    function showBubblePopup(el, msg, overrideCheckValidity) {

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

        // removing previous message from this element
        var elErrors = el.parentNode.querySelectorAll(":scope > .bubble-invalid-anchor");
        if (attachOuterOuter) {
            elErrors = el.parentNode.parentNode.querySelectorAll(":scope > .bubble-invalid-anchor");
        }
        for (var i = 0; i < elErrors.length; i++) {
            var oldChild = elErrors[i].parentNode.removeChild(elErrors[i]);
            // console.log(`removed-a ${i + 1}of${elErrors.length} - oldID${oldChild.getAttribute('id')} `);
        }

        if (!el.checkValidity() || overrideCheckValidity === true) {
            var parent = el.parentNode;
            if (attachOuterOuter) {
                parent = el.parentNode.parentNode;
            }
            // el.validationMessage is mathematical has is always in browser local
            parent.insertAdjacentHTML(
                "beforeend",
                `<div class='bubble-invalid-anchor'  id='err-${el.getAttribute('name')}' >
                    <div class='bubble-invalid-content'>
                    ${msg}
                    </div>
                </div>`
            );
        }

    }

    // for onsubmit
    // for each invalid input element of a form
    // a bubble message is displayed right-next or -below
    function onSubmitCustomBubblesForInvalids(event) {

        clearAllPreviousBubbles();

        // insert new messages at the end of parent
        var invalidFields = this.querySelectorAll(":invalid");
        for (var i = 0; i < invalidFields.length; i++) {
            showBubblePopup(invalidFields[i]);
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





    this.validateFormWithCustomBubbles = function() {


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

        // add custom bubbles
        form.addEventListener(
            "submit",
            onSubmitCustomBubblesForInvalids,
            true
        );
        console.log(`on submit: custom bubbles for invalids attached`);


    }



    // showing custom bubbles on form submit is too late?
    //   => show them on blur
    this.showBubbleOnBlurOrInput = function() {

        // if we dont apply validateFormWithCustomBubbles(),
        // then we still need this for standalone functionality
        form.setAttribute("novalidate", true);

        var funcReport = function (event) {
            // event.target.reportValidity();
            showBubblePopup(event.target);
            var lgMsg = "blur";
            if (checkOnInput) {
                lgMsg = "blur+input";
            }
            console.log(`  ${ lgMsg } inp.reportValidity() ${event.target.getAttribute('name')}`);
            if (reclaimFocus) {
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
            if (checkOnInput) {
                inp.addEventListener("input", funcReport);
            }

            if (removeBubbleOnEntry) {
                var removeOnEntering = function (event) {
                    var el = event.target;
                    var elErrors = el.parentNode.querySelectorAll(":scope > .bubble-invalid-anchor");
                    for (var i = 0; i < elErrors.length; i++) {
                        var oldChild = elErrors[i].parentNode.removeChild(elErrors[i]);
                        // console.log(`removed ${i + 1}of${elErrors.length} - oldID${oldChild.getAttribute('id')} `);
                    }
                };
                inp.addEventListener("focus", removeOnEntering);    // remove on entering input
            }
            var lgMsg = "blur";
            if (checkOnInput) {
                lgMsg = "blur+input";
            }
            console.log(`     ${lgMsg} handler added to ${inp.getAttribute('name')}`);
        }
    }



}

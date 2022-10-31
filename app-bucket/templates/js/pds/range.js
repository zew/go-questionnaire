// non global block
{

    // the wide rectangle, not the `thumb` 
    let range = document.getElementById("{{.inputName }}");

    // an disabled input, where the value is shown as digits
    let rangeDisplay = document.getElementsByName("{{.inputName }}_display")[0];

    let initRange = (inst) => {
        // range.value = 25; // init value should be set in the control
        // range.click();

        const evt = new Event("input");
        range.dispatchEvent(evt);

        range.setAttribute("list", "range-scale");
        rangeDisplay.setAttribute("disabled", "true");
    }


    // update slider
    range.oninput = function () {
        let incr = parseInt(this.value) + 5;
        // backticks formatted string not working  
        // sliderDisplay.value =   `${this.value}  -  ${incr}`;
        let out = ""
        if (this.value) {
            out += this.value;
        }
        out += " - ";
        out += incr;

        rangeDisplay.value = out;
    }

    // init slider;
    // triggers oninput above
    window.addEventListener('load', initRange, false);


}
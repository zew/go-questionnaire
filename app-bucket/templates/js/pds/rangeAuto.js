// non global block
{
    let range = document.getElementById("{{.inputName }}");

/*     
    // outer rectangle, not the *thumb*
    let displ = document.getElementsByName("{{.inputName }}_display")[0]; // range display

    // update slider
    range.oninput = function () {
        // let incr = parseInt(this.value) + parseInt(this.step);
        let incr = parseFloat(this.value) + parseFloat(this.step);
        let out = ""
        if (this.value) {
            out += this.value;
        }
        out += " - ";
        out += incr;

        displ.value = out;
    }
 */

    let initRange = (inst) => {
        const evt = new Event("input");
        range.dispatchEvent(evt);
    }

    // init slider;
    // triggers oninput above
    window.addEventListener('load', initRange, false);

}
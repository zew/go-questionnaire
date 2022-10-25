
var slider01 = document.getElementById("slider01");
var slider01Legend = document.getElementsByName("slider01Legend")[0];

var slider02 = document.getElementById("slider02");
var slider02Legend = document.getElementsByName("slider02Legend")[0];



let initSlider01 = (inst) => {
    slider01.value = 25; 
    // slider01.click();
    const evt = new Event("input");
    slider01.dispatchEvent(evt);    
}
let initSlider02 = (inst) => {
    slider02.value = 25; 
    // slider02.click();
    const evt = new Event("input");
    slider02.dispatchEvent(evt);    
}

// init sliders;
window.addEventListener('load', initSlider01, false);
window.addEventListener('load', initSlider02, false);


// update sliders
slider01.oninput = function () {
    let incr = parseInt(this.value) + 5;
    slider01Legend.value =  `${this.value}  -  ${incr}`;
}

slider02.oninput = function () {
    let incr = parseInt(this.value) + 5;
    slider02Legend.value =  `${this.value}  -  ${incr}`;
}



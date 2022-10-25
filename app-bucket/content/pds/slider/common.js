
var slider01 = document.getElementById("slider01");

var slider01cx1  = document.getElementsByName("slider01cx1")[0];
var slider01cx2  = document.getElementsByName("slider01cx2")[0];
var slider01Legend = document.getElementsByName("slider01Legend")[0];

// init output windows
// slider01cx1.value = 100 - slider01.value;
// slider01cx2.value = slider01.value;
// slider01Legend.value = slider01.value

// slider01.value = 40; // setting an init value



let initSlider = () => {
    slider01.value = 25; 
    // slider01.click();
    const evt = new Event("input");
    slider01.dispatchEvent(evt);    
}

// initSlider();
window.addEventListener('load', initSlider, false);


// update
slider01.oninput = function () {
    // slider01cx1.value = 100 - this.value;
    // slider01cx2.value = this.value;
    let incr = parseInt(this.value) + 5;
    slider01Legend.value =  `${this.value}  -  ${incr}`;
}



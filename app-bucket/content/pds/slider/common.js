
var slider01Thumb = document.getElementById("slider01Thumb");
slider01Thumb.value = 40; // setting an init value

var slider01cx1  = document.getElementsByName("slider01cx1")[0];
var slider01cx2  = document.getElementsByName("slider01cx2")[0];

// init output windows
slider01cx1.value = 100 - slider01Thumb.value;
slider01cx2.value = slider01Thumb.value;

// update
slider01Thumb.oninput = function () {
    slider01cx1.value = 100 - this.value;
    slider01cx2.value = this.value;
}



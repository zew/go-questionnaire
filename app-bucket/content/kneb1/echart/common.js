
var slider = document.getElementById("sliderInner");
var safe = document.getElementsByName("share_safe")[0];
var risky = document.getElementsByName("share_risky")[0];

// init
safe.value = 100 - slider.value;
risky.value = slider.value;

// update
slider.oninput = function () {
    safe.value = 100 - this.value;
    risky.value = this.value;
}



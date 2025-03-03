"use strict";
(() => {
    const showSlidesBtn = document.querySelector("#show-slides-btn");
    const slidesDiv = document.querySelector("#slides");
    showSlidesBtn.addEventListener("click", () => {
        slidesDiv.classList.toggle("h-0");
        slidesDiv.classList.toggle("min-h-fit");
    });
})();

"use strict";
(() => {
    const slides = document.querySelector("#slides-container");
    const btn = document.querySelector("#view-slides-btn");
    btn.addEventListener("click", () => {
        console.log("button clicked");
        slides.classList.toggle("h-0");
        slides.classList.toggle("h-full");
        slides.classList.toggle("opacity-0");
        slides.classList.toggle("opacity-100");
    });
})();

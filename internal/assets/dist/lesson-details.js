"use strict";
(() => {
    const showSlidesBtn = document.querySelector("#show-slides-btn");
    const slidesDiv = document.querySelector("#slides");
    showSlidesBtn === null || showSlidesBtn === void 0 ? void 0 : showSlidesBtn.addEventListener("click", () => {
        slidesDiv === null || slidesDiv === void 0 ? void 0 : slidesDiv.classList.toggle("h-0");
        slidesDiv === null || slidesDiv === void 0 ? void 0 : slidesDiv.classList.toggle("min-h-fit");
    });
})();

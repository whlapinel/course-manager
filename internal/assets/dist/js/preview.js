"use strict";
(() => {
    var _a, _b;
    console.log("hello from preview.js!");
    (_a = document.body) === null || _a === void 0 ? void 0 : _a.addEventListener("click", (e) => {
        const previewDialog = document.querySelectorAll("dialog");
        const clickedElement = e.target;
        console.log(clickedElement);
        if ((clickedElement === null || clickedElement === void 0 ? void 0 : clickedElement.id) === "open-dialog") {
            console.log("open-dialog clicked!");
            for (const el of previewDialog) {
                el === null || el === void 0 ? void 0 : el.showModal();
            }
        }
    });
    (_b = document.body) === null || _b === void 0 ? void 0 : _b.addEventListener("click", (e) => {
        const clickedElement = e.target;
        const previewDialog = document.querySelectorAll("dialog");
        if ((clickedElement === null || clickedElement === void 0 ? void 0 : clickedElement.id) === "close-dialog") {
            console.log("preview button clicked!");
            for (const el of previewDialog) {
                el === null || el === void 0 ? void 0 : el.close();
            }
        }
    });
})();

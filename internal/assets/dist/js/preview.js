"use strict";
(() => {
    var _a, _b;
    console.log("hello from preview.js!");
    (_a = document.body) === null || _a === void 0 ? void 0 : _a.addEventListener("click", (e) => {
        const previewDialog = document.querySelector("#dialog-preview");
        const clickedElement = e.target;
        console.log(clickedElement);
        if ((clickedElement === null || clickedElement === void 0 ? void 0 : clickedElement.id) === "preview-button") {
            console.log("preview button clicked!");
            previewDialog === null || previewDialog === void 0 ? void 0 : previewDialog.showModal();
        }
    });
    (_b = document.body) === null || _b === void 0 ? void 0 : _b.addEventListener("click", (e) => {
        const clickedElement = e.target;
        const previewDialog = document.querySelector("#dialog-preview");
        if ((clickedElement === null || clickedElement === void 0 ? void 0 : clickedElement.id) === "dialog-close-button") {
            console.log("preview button clicked!");
            previewDialog === null || previewDialog === void 0 ? void 0 : previewDialog.close();
        }
    });
})();

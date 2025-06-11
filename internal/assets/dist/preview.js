"use strict";
(() => {
    console.log("hello from preview.js!");
    const previewButton = document.querySelector("#preview-button");
    const previewDialog = document.querySelector("dialog");
    previewButton === null || previewButton === void 0 ? void 0 : previewButton.addEventListener("click", () => {
        console.log("preview button clicked!");
        previewDialog === null || previewDialog === void 0 ? void 0 : previewDialog.showModal();
    });
})();

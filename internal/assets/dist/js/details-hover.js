"use strict";
(() => {
    const script = document.currentScript;
    const parent = script === null || script === void 0 ? void 0 : script.parentElement;
    parent === null || parent === void 0 ? void 0 : parent.addEventListener("mouseenter", () => {
        parent.setAttribute("open", "");
    });
    parent === null || parent === void 0 ? void 0 : parent.addEventListener("mouseleave", () => {
        parent.removeAttribute("open");
    });
})();

"use strict";
(() => {
    const collapsibles = document.querySelectorAll('.collapsible');
    for (const collapsible of collapsibles) {
        collapsible.addEventListener('click', () => {
            var _a, _b;
            (_a = collapsible.nextElementSibling) === null || _a === void 0 ? void 0 : _a.classList.toggle('max-h-0');
            (_b = collapsible.nextElementSibling) === null || _b === void 0 ? void 0 : _b.classList.toggle('max-h-fit');
        });
    }
})();

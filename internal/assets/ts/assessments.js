"use strict";
/// <reference path="./@types/htmx.d.ts" />
console.log("hello from assessments.js!");
function addAssessmentListeners() {
    const filterButton = document.querySelector("#filter-button");
    const filterByActive = document.querySelector("#filter-by-active");
    const sortBy = document.querySelector("#sort-by");
    const startInput = document.querySelector("#start");
    const endInput = document.querySelector("#end");
    const categorySelect = document.querySelector("#filter-by-category");
    const { protocol, host } = window.location;
    const baseUrl = `${protocol}//${host}`;
    function updateQuery(key, val) {
        const hxVal = filterButton.getAttribute("hx-get");
        const url = new URL(hxVal, baseUrl);
        url.searchParams.set(key, val);
        filterButton.setAttribute("hx-get", url.toString());
        window.htmx.process(filterButton);
    }
    filterByActive === null || filterByActive === void 0 ? void 0 : filterByActive.addEventListener("change", () => {
        const param = filterByActive.value;
        updateQuery("active", param);
    });
    sortBy === null || sortBy === void 0 ? void 0 : sortBy.addEventListener("change", () => {
        const param = sortBy.value;
        updateQuery("sort-by", param);
    });
    startInput === null || startInput === void 0 ? void 0 : startInput.addEventListener("change", () => {
        const date = startInput.value;
        updateQuery("start", date);
    });
    endInput === null || endInput === void 0 ? void 0 : endInput.addEventListener("change", () => {
        const date = endInput.value;
        updateQuery("end", date);
    });
    categorySelect === null || categorySelect === void 0 ? void 0 : categorySelect.addEventListener("change", () => {
        const category = categorySelect.value;
        updateQuery("category", category);
    });
}
(() => {
    addAssessmentListeners();
    document.addEventListener("htmx:afterSwap", addAssessmentListeners);
})();

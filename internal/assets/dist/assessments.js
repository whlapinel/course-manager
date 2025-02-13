"use strict";
/// <reference path="./@types/htmx.d.ts" />
console.log("hello from assessments.js!");
(() => {
    const filterButton = document.querySelector("#filter-button");
    const startInput = document.querySelector("#start");
    const endInput = document.querySelector("#end");
    const categorySelect = document.querySelector("#category");
    const { protocol, host } = window.location;
    const baseUrl = `${protocol}//${host}`;
    function updateQuery(key, val) {
        const hxVal = filterButton.getAttribute("hx-get");
        const url = new URL(hxVal, baseUrl);
        url.searchParams.set(key, val);
        filterButton.setAttribute("hx-get", url.toString());
        window.htmx.process(filterButton);
    }
    startInput.addEventListener("change", () => {
        const date = startInput.value;
        updateQuery("start", date);
    });
    endInput.addEventListener("change", () => {
        const date = endInput.value;
        updateQuery("end", date);
    });
    categorySelect.addEventListener("change", () => {
        const category = categorySelect.value;
        updateQuery("category", category);
    });
})();

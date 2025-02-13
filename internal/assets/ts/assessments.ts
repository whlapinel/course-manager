/// <reference path="./@types/htmx.d.ts" />


console.log("hello from assessments.js!");
(() => {
    const filterButton: HTMLButtonElement = document.querySelector("#filter-button")!
    const startInput: HTMLInputElement = document.querySelector("#start")!
    const endInput: HTMLInputElement = document.querySelector("#end")!
    const categorySelect: HTMLSelectElement = document.querySelector("#category")!
    const { protocol, host } = window.location;
    const baseUrl = `${protocol}//${host}`;
    function updateQuery(key: string, val: string) {
        const hxVal: string = filterButton.getAttribute("hx-get")!
        const url = new URL(hxVal, baseUrl)!
        url.searchParams.set(key, val)
        filterButton.setAttribute("hx-get", url.toString())
        window.htmx.process(filterButton)

    }
    startInput.addEventListener("change", () => {
        const date: string = startInput.value
        updateQuery("start", date)
    })
    endInput.addEventListener("change", () => {
        const date: string = endInput.value
        updateQuery("end", date)
    })
    categorySelect.addEventListener("change", () => {
        const category: string = categorySelect.value
        updateQuery("category", category)
    })

})();
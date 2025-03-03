(() => {
    const showSlidesBtn: HTMLButtonElement = document.querySelector("#show-slides-btn")!
    const slidesDiv: HTMLDivElement = document.querySelector("#slides")!
    showSlidesBtn.addEventListener("click", () => {
        slidesDiv.classList.toggle("h-0")
        slidesDiv.classList.toggle("min-h-fit")
    })
})()
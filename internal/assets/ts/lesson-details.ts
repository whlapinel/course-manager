(() => {
    const showSlidesBtn: HTMLButtonElement | null = document.querySelector("#show-slides-btn")
    const slidesDiv: HTMLDivElement | null = document.querySelector("#slides")
    showSlidesBtn?.addEventListener("click", () => {
        slidesDiv?.classList.toggle("h-0")
        slidesDiv?.classList.toggle("min-h-fit")
    })
})()
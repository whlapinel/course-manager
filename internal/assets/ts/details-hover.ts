(() => {
    const script = document.currentScript
    const parent = script?.parentElement
    parent?.addEventListener("mouseenter", () => {
        parent.setAttribute("open", "")
    })
    parent?.addEventListener("mouseleave", () => {
        parent.removeAttribute("open")
    })
})()
(()=>{
    console.log("hello from preview.js!")
    document.body?.addEventListener("click", (e)=>{
        const previewDialog = document.querySelector("#dialog-preview") as HTMLDialogElement
        const clickedElement = e.target as HTMLElement
        console.log(clickedElement)
        if (clickedElement?.id === "preview-button") {
            console.log("preview button clicked!")
            previewDialog?.showModal()
        }
    })
    document.body?.addEventListener("click", (e)=>{
        const clickedElement = e.target as HTMLElement
        const previewDialog = document.querySelector("#dialog-preview") as HTMLDialogElement
        if (clickedElement?.id === "dialog-close-button") {
            console.log("preview button clicked!")
            previewDialog?.close()
        }
    })
})()
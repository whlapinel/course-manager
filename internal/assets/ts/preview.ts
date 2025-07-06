(()=>{
    console.log("hello from preview.js!")
    document.body?.addEventListener("click", (e)=>{
        const previewDialog = document.querySelectorAll("dialog")
        const clickedElement = e.target as HTMLElement
        console.log(clickedElement)
        if (clickedElement?.id === "open-dialog") {
            console.log("open-dialog clicked!")
            for (const el of previewDialog){
                el?.showModal()
            }
        }
    })
    document.body?.addEventListener("click", (e)=>{
        const clickedElement = e.target as HTMLElement
        const previewDialog = document.querySelectorAll("dialog")
        if (clickedElement?.id === "close-dialog") {
            console.log("preview button clicked!")
            for (const el of previewDialog) {
                el?.close()
            }
        }
    })
})()
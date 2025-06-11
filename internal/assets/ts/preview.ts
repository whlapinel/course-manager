(()=>{
    console.log("hello from preview.js!")
    const previewButton = document.querySelector("#preview-button")
    const previewDialog = document.querySelector("dialog")
    previewButton?.addEventListener("click", ()=>{
        console.log("preview button clicked!")
        previewDialog?.showModal()
    })
})()
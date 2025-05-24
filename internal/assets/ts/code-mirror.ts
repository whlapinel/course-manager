
(()=>{
    document.body.addEventListener("htmx:afterSettle", (e)=>{
        const editor = document.querySelector("#code-editor") as HTMLTextAreaElement;
        if (editor){
            const cm = CodeMirror.fromTextArea(editor, {
                lineNumbers: true,
                mode: "markdown",            
            })
            cm.on('change', (e) => {
                editor.value = cm.getValue();
            })
        }        
    })
})()

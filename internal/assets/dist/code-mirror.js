"use strict";
(() => {
    document.body.addEventListener("htmx:afterSettle", (e) => {
        const editor = document.querySelector("#code-editor");
        if (editor) {
            const cm = CodeMirror.fromTextArea(editor, {
                lineNumbers: true,
                mode: "markdown",
            });
            cm.on('change', (e) => {
                console.log("I'm updating the textarea!")
                editor.value = cm.getValue();
                console.log(editor.value)
            });
        }
    });
})();

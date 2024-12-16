(() => {
    const collapsibles = document.querySelectorAll('.collapsible')
    for (const collapsible of collapsibles) {
        collapsible.addEventListener('click', () => {
            collapsible.nextElementSibling?.classList.toggle('max-h-0')
            collapsible.nextElementSibling?.classList.toggle('max-h-fit')
        })
    }
})();
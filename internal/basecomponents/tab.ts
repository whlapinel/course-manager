(() => {
    setTimeout(() => {
        console.log("hello from tab.js!")
        // This narrows the queried nodes to HTMLAnchorElement, which helps with type-checking
        const tabs: any = document.querySelectorAll<HTMLAnchorElement>("#tab-menu ul li a");
        console.log(tabs)


        for (const tab of tabs) {
            tab.addEventListener("click", (event: MouseEvent) => {
                console.log("tab clicked!", event.currentTarget)
                // Remove "active" from every tab
                for (const t of tabs) {
                    t.classList.remove("bg-blue-500", "text-white", "font-bold");
                }
                // Cast the event currentTarget to HTMLAnchorElement, then add "active"
                (event.currentTarget as HTMLAnchorElement).classList.add("bg-blue-500", "text-white", "font-bold");
            });
        }
    }, 500);
})();

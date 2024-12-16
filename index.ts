(() => {
    console.log("Hello World");
    const navAbout = document.querySelector("#nav-about");
    const aboutMenu = document.querySelector("#about-menu");
    navAbout?.addEventListener("click", (e) => {
        console.log(e.target);
        aboutMenu?.classList.toggle("hidden");
        aboutMenu?.classList.toggle("flex");
    });
    document.addEventListener("click", (e) => {
        if (e.target !== navAbout && e.target !== aboutMenu) {
            aboutMenu?.classList.add("hidden");
            aboutMenu?.classList.remove("flex");
        }
    });
    document.addEventListener("htmx:afterSettle", () => {
        console.log("event triggered")
        const pageTitles = ["Index", "About", "Contact", "Blog", "Courses"];
        const currLocation = window.location.pathname;
        console.log(currLocation);
        for (const title of pageTitles) {
            if (currLocation.includes(title.toLowerCase())) {
                console.log(title);
                const titleElement = document.querySelector(`#nav-${title.toLowerCase()}`);
                const selectedElement = document.querySelector(".highlighted");
                console.log("titleElement", titleElement);
                console.log("selectedElement", selectedElement);
                if (selectedElement?.id !== titleElement?.id) {
                    console.log("selectedElement does not equal titleElement", selectedElement?.id, titleElement?.id);
                    selectedElement?.classList.remove("bg-green-700");
                    selectedElement?.classList.remove("highlighted");
                    console.log(titleElement);
                    titleElement?.classList.add("bg-green-700");
                    titleElement?.classList.add("highlighted");
                }
            }

        }
    })
})()
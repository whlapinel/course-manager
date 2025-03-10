"use strict";
(() => {
    console.log("Hello World from static index.ts!");
    const navAbout = document.querySelector("#nav-about");
    const aboutMenu = document.querySelector("#about-menu");
    navAbout === null || navAbout === void 0 ? void 0 : navAbout.addEventListener("click", (e) => {
        console.log(e.target);
        aboutMenu === null || aboutMenu === void 0 ? void 0 : aboutMenu.classList.toggle("hidden");
        aboutMenu === null || aboutMenu === void 0 ? void 0 : aboutMenu.classList.toggle("flex");
    });
    document.addEventListener("click", (e) => {
        if (e.target !== navAbout && e.target !== aboutMenu) {
            aboutMenu === null || aboutMenu === void 0 ? void 0 : aboutMenu.classList.add("hidden");
            aboutMenu === null || aboutMenu === void 0 ? void 0 : aboutMenu.classList.remove("flex");
        }
    });
    document.addEventListener("htmx:afterSettle", () => {
        console.log("event triggered");
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
                if ((selectedElement === null || selectedElement === void 0 ? void 0 : selectedElement.id) !== (titleElement === null || titleElement === void 0 ? void 0 : titleElement.id)) {
                    console.log("selectedElement does not equal titleElement", selectedElement === null || selectedElement === void 0 ? void 0 : selectedElement.id, titleElement === null || titleElement === void 0 ? void 0 : titleElement.id);
                    selectedElement === null || selectedElement === void 0 ? void 0 : selectedElement.classList.remove("bg-green-700");
                    selectedElement === null || selectedElement === void 0 ? void 0 : selectedElement.classList.remove("highlighted");
                    console.log(titleElement);
                    titleElement === null || titleElement === void 0 ? void 0 : titleElement.classList.add("bg-green-700");
                    titleElement === null || titleElement === void 0 ? void 0 : titleElement.classList.add("highlighted");
                }
            }
        }
    });
})();

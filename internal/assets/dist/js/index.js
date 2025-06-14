"use strict";
(() => {
    console.log("hello from server's index.ts converted to js!");
    const button = document.querySelector('#user-menu-button');
    const dropdown = document.querySelector('#dropdown-menu');
    let isOpen = false;
    const toggleDropdown = () => {
        isOpen = !isOpen;
        if (isOpen) {
            dropdown === null || dropdown === void 0 ? void 0 : dropdown.classList.remove('opacity-0', 'scale-95', 'ease-in', 'duration-75');
            dropdown === null || dropdown === void 0 ? void 0 : dropdown.classList.add('opacity-100', 'scale-100', 'ease-out', 'duration-100');
            button === null || button === void 0 ? void 0 : button.setAttribute('aria-expanded', 'true');
        }
        else {
            dropdown === null || dropdown === void 0 ? void 0 : dropdown.classList.remove('opacity-100', 'scale-100', 'ease-out', 'duration-100');
            dropdown === null || dropdown === void 0 ? void 0 : dropdown.classList.add('opacity-0', 'scale-95', 'ease-in', 'duration-75');
            button === null || button === void 0 ? void 0 : button.setAttribute('aria-expanded', 'false');
        }
    };
    button === null || button === void 0 ? void 0 : button.addEventListener('click', toggleDropdown);
    // Close dropdown if clicking outside
    document.addEventListener('click', (event) => {
        if (isOpen && !dropdown.contains(event.target) && !button.contains(event.target)) {
            toggleDropdown();
        }
    });
})();

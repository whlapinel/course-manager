(function () {
    setTimeout(function () {
        console.log("hello from tab.js!");
        // This narrows the queried nodes to HTMLAnchorElement, which helps with type-checking
        var tabs = document.querySelectorAll("#tab-menu ul li a");
        console.log(tabs);
        for (var _i = 0, tabs_1 = tabs; _i < tabs_1.length; _i++) {
            var tab = tabs_1[_i];
            tab.addEventListener("click", function (event) {
                console.log("tab clicked!", event.currentTarget);
                // Remove "active" from every tab
                for (var _i = 0, tabs_2 = tabs; _i < tabs_2.length; _i++) {
                    var t = tabs_2[_i];
                    t.classList.remove("bg-blue-500", "text-white", "font-bold");
                }
                // Cast the event currentTarget to HTMLAnchorElement, then add "active"
                event.currentTarget.classList.add("bg-blue-500", "text-white", "font-bold");
            });
        }
    }, 500);
})();

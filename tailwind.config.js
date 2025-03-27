/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./internal/templates/**/*.{html,js,templ}"
    ],
    theme: {
        extend: {},
    },
    plugins: [],
    safelist: [
        "max-h-0",
        "max-h-fit",
        "min-h-fit",
        "bg-blue-500",
        "text-white",
        "font-bold"
    ],

}
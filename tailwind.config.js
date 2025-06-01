/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./internal/**/*.{html,js,templ}",
        "./hugosites/**/*.{html,js,gotmpl}",
        "./internal/staticresources/layouts/**/*.{html,js,gotmpl}",
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
/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./**/*.{html,js,templ}",
        "!./node_modules/**/*",
    ],
    theme: {
        extend: {},
    },
    plugins: [],
    safelist: [
        "max-h-0",
        "max-h-fit",
    ],
}
package main

//go:generate templ generate
//go:generate npx tailwindcss -i ./manager-input.css -o ./internal/assets/dist/styles.css
//go:generate npx tailwindcss -i ./docs-input.css -o ./sites/assets/styles/styles.css
//go:generate tsc --build

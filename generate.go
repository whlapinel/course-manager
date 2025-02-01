package main

//go:generate templ generate
//go:generate npx tailwindcss -i ./manager-input.css -o ./internal/assets/dist/styles.css
//go:npx tailwindcss -i ./docs-input.css -o ./python/docs/styles/styles.css
//go:generate tsc --build

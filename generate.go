package main

//go:generate templ generate
//go:generate npx tailwindcss -i ./input.css -o ./python/docs/styles/styles.css
//go:generate npx tailwindcss -i ./manager-input.css -o ./cmd/web_app/assets/dist/styles.css
//go:generate tsc --build

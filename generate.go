package main

//go:generate sqlc generate -f sqlc_new.yml
//go:generate templ generate
//go:generate npx tailwindcss -i ./manager-input.css -o ./internal/assets/dist/styles.css
//go:generate npx tailwindcss -i ./static-input.css -o ./internal/staticresources/styles/styles.css
//go:generate npx tsc --build

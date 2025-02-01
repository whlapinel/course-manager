package main

//go:generate templ generate
//go:generate npx tailwindcss -c tailwind-manager.config.js -i ./manager-input.css -o ./cmd/web_app/assets/dist/styles.css && npx tailwindcss -c tailwind-docs.config.js -i ./docs-input.css -o ./python/docs/styles/styles.css
//go:generate tsc --build

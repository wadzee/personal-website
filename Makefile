run:
	@templ generate
	@go run cmd/main.go

templ:
	@templ generate -watch -proxy=http://localhost:3000

tailwind:
	@tailwindcss -i static/css/input.css -o static/css/output.css --watch

build:
	@templ generate
	@npx tailwindcss build -i static/css/input.css -o static/css/output.css
	@go build -o ./tmp/main ./cmd
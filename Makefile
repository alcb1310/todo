run:
	@go run cmd/api/main.go

build:
	@go build -o build/todo cmd/api/main.go

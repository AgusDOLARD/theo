build:
	@go build -o bin/theo cmd/main.go

run: build
	@bin/theo

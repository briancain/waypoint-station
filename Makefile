.PHONY: build
build: # Build the project
	@go build -o bin/ ./...

.PHONY: format
format: # Format all go code in project
	@gofmt -s -w ./

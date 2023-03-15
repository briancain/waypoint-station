.PHONY: build
build: # Build the project
	@go build -o bin/ ./...

.PHONY: run
run: # Build the project
	./bin/waypoint-station

.PHONY: format
format: # Format all go code in project
	@gofmt -s -w ./

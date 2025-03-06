# Run Wire
runwire:
	cd internal/app && ~/go/bin/wire

# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o bin/app

# Run tests
test:
	go test ./...

# Generate Swagger documentation
swagger:
	~/go/bin/swag init

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy
	go mod download

# Default target
.PHONY: runwire run build test swagger clean deps
default: deps swagger runwire build


# Run Dev Command, run wire and then run the application
dev:
	@cd internal/app && ~/go/bin/wire
	@go run main.go

# Run Air Command, run the application in a development environment
air:
	@cd internal/app && ~/go/bin/air
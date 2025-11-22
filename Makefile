.PHONY: build run format test clean install

# Build the application
build:
	@echo "Building aws-eye..."
	@go build -o aws-eye .

# Run the application
run: build
	@echo "Running aws-eye..."
	@./aws-eye describe

# Format the code
format:
	@echo "Formatting code..."
	@go fmt ./...

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f aws-eye
	@go clean

# Install dependencies
install:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Run with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"


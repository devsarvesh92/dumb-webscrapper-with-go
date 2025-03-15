# Makefile

# Go executable name
EXECUTABLE := dumb-webscrapper-with-go

# Build directory (optional, but good practice)
BUILD_DIR := bin

# Build the Go application
build:
	go build -o $(EXECUTABLE) cmd/main.go

# Run the Go application
run:
	go run cmd/main.go -urls "https://www.scrapethissite.com/pages"

# Clean the executable
clean:
	rm -f $(EXECUTABLE)

.PHONY: build run clean

# Makefile for building the project.

build:
	@echo "Running the build program..."
	@echo "Using hard-coded log file for now..."
	go run main.go analyse --input-file="logs.log"

test:
	@echo "Performing tests..."
	go test -v ./utils
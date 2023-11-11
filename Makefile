.DEFAULT_GOAL := build

.PHONY: fmt vet build
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	mkdir -p bin
	go build -o bin/nasa-apod
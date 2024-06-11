.DEFAULT_GOAL := run

.PHONY: fmt vet build run
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean

run: build
	./pointer_are_last_resort

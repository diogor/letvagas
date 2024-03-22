SHELL := /bin/bash
.PHONY: build clean

clean:
	go clean
	go mod tidy
	rm -rf ./build

build: clean
	mkdir -p ./build
	go build -o ./build/letvagas
	cp -r ./templates ./build/
	cp -r ./static ./build/

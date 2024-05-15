.PHONY: build run

build:
	go build

run: build
	go run main.go
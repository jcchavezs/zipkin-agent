.PHONY: build

run:
	go run cmd/zipkinagent/main.go

build:
	go build -o build/zipkin-agent cmd/zipkinagent/main.go
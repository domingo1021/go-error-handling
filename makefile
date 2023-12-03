build:
	go build -o ./bin/error-handling

run-binary:
	./bin/error-handling

run:
	go run main.go

.PHONY: build run-binary run
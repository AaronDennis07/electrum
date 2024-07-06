BINARY_NAME=electrum

run:
	go run ./cmd/main.go

build:
	go build -o bin/${BINARY_NAME} ./cmd/main.go

clean:
	rm -rf bin

start:
	go build -o bin/${BINARY_NAME} ./cmd/main.go
	./bin/${BINARY_NAME}
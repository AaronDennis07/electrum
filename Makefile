BINARY_NAME=electrum

run:
	go run .

build:
	go build -o bin/${BINARY_NAME} .

clean:
	rm -rf bin

start:
	./bin/${BINARY_NAME}
BINARY_NAME=playlist

build: 
	set GOARCH=amd64 
	set GOOS=linux 
	go build -o bin/${BINARY_NAME}-linux ./cmd/playlist/main.go
	set GOARCH=amd64 
	set GOOS=windows 
	go build -o bin/${BINARY_NAME}-windows.exe ./cmd/playlist/main.go

run:
	go run ./cmd/playlist/main.go

all:
	build
	run

clean:
	go clean
	rm ./bin/${BINARY_NAME}-linux
	rm ./bin/${BINARY_NAME}-windows.exe

tidy: 
	go mod tidy
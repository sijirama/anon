all:run

build:
	@go build -o bin/seapick

build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/seapick-linux-amd64

build-windows:
	@GOOS=windows GOARCH=amd64 go build -o bin/seapick-windows-amd64.exe

build-macos:
	@GOOS=darwin GOARCH=amd64 go build -o bin/seapick-darwin-amd64

release: build-linux build-windows build-macos


run: build
	@./bin/seapick

test:
	@go test ./... 

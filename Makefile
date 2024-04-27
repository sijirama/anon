all:run

build:
	@go build -o bin/seapick

run: build
	@./bin/seapick

test:
	@go test ./... 

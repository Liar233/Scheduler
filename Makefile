all: build

dependencies:
	go get .

build: dependencies
	go build -o ./scheduler_service .

build-dev: dependencies
	go build -o ./build/scheduler .

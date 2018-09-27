all: build

dependencies:
	go get .

build: dependencies
	go build -o ./scheduler .

build-dev: dependencies storage-driver-stub channel-driver-stub
	go build -o ./build/scheduler .

storage-driver-stub:
	go build -buildmode=plugin -o ./build/drivers/storage_driver_stub.so ./drivers/storage/storage_driver_stub.go
channel-driver-stub:
	go build -buildmode=plugin -o ./build/drivers/channel_driver_stub.so ./drivers/channel/channel_driver_stub.go

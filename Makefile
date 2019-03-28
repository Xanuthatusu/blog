test:
	go test -v ./...

install:
	go get ./...

generate:
	protoc -I protos protos/blog.proto --go_out=plugins=grpc:protos


test:
	go test -v ./...

install:
	curl -L -o protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-osx-x86_64.zip
	unzip protoc.zip -d protoc-3.11.2-osx-x86_64
	rm protoc.zip
	protoc-3.11.2-osx-x86_64/bin/protoc -I protos protos/blog.proto --go_out=plugins=grpc:protos
	go get ./...

clean:
	rm -rf protoc-3.11.2-osx-x86_64


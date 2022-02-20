lint:
	protolint -fix internal/protos
	golangci-lint run ./...



.PHONY: proto

proto:
	protolint -fix  protos
	protoc \
	--proto_path=protos \
	--go_out=. \
	--go-grpc_out=. \
	protos/*.proto

clients:
	protolint -fix  protos
	protoc \
    --proto_path=protos \
	--js_out=import_style=commonjs,binary:./client     \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./client \
	protos/*.proto
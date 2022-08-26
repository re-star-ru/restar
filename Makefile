run:
	CGO_ENABLED=0 go run ./cmd/restar

build:
	CGO_ENABLED=0 go build -o ./bin ./cmd/restar

generate:
	protoc --go_out . --go-grpc_out=. ./api/proto/v1/*.proto

example:
	CGO_ENABLED=0 go run ./example

.PHONY: example
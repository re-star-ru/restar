run:
	CGO_ENABLED=0 go run ./cmd/restar

dev-server:
	docker compose -f ./deployments/dev/docker-compose.yaml up
down-dev-server:
	docker compose -f ./deployments/dev/docker-compose.yaml down

build:
	CGO_ENABLED=0 go build -o ./.bin ./cmd/restar

generate:
	protoc --go_out . --go-grpc_out=. ./api/proto/v1/*.proto

example:
	CGO_ENABLED=0 go run ./example

.PHONY: example
run:
	CGO_ENABLED=0 go run ./cmd/restar

dev-server:
	docker compose -f ./deployments/dev/docker-compose.yaml up db
down-dev-server:
	sudo rm -rf ./.bin/db
	docker compose -f ./deployments/dev/docker-compose.yaml down

build:
	CGO_ENABLED=0 go build -o ./.bin ./cmd/restar

#generate: generate-go generate-flutter

generate:
	go generate ./...

generate-go:
	protoc --go_out . --go-grpc_out=. ./api/proto/v1/*.proto

generate-flutter:
	protoc -Iapi/proto/v1 --dart_out=grpc:api/flutter/v1  api/proto/v1/*.proto

example:
	CGO_ENABLED=0 go run ./example

test:
	docker compose -p restar-test -f ./test/docker-compose.yaml down || true
	docker compose -p restar-test -f ./test/docker-compose.yaml up -d
	go test ./test/... -v -coverpkg=./... -coverprofile=coverage.out || true

coverage:
	go tool cover -html=coverage.out

.PHONY: example test
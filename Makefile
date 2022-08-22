run:
	CGO_ENABLED=0 go run ./cmd/restar

build:
	CGO_ENABLED=0 go build -o ./bin ./cmd/restar
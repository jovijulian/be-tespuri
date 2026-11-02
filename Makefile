.PHONY: run migrate seed reset test build clean

run:
	go run ./cmd/api

migrate:
	go run ./cmd/api --migrate

seed:
	go run ./cmd/api --seed

reset:
	go run ./cmd/api --reset

test:
	go test -v ./...

build:
	go build -o bin/api ./cmd/api

clean:
	rm -rf bin/

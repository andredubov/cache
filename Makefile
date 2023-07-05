
.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

run: build
	./.bin/app

test:
	go test -v ./...
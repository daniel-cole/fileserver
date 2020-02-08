VERSION := $(shell git describe --always --long --dirty)
BRANCH := $(shell ([ ! -z ${BRANCH} ] && echo ${BRANCH}) || git rev-parse --abbrev-ref HEAD)

test:
	go test ./...

build:
	GOOS=linux GOARCH=amd64 go build -i -v -o bin/fileserver -ldflags="-X main.Version=${BRANCH}-${VERSION}"
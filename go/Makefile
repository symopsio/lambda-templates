VERSION := $(shell git describe --tags --always --dirty="-dev")
LDFLAGS := -ldflags='-X "main.Version=$(VERSION)"'

all: dist

clean:
	rm -rf ./bin ./dist

build:
	mkdir -p bin
	(cd cmd/handler && GOOS=linux GO111MODULE=on go build $(LDFLAGS) -o ../../bin/handler)

dist: build
	mkdir -p dist
	zip dist/handler.zip bin/handler

.PHONY: clean build dist all

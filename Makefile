.PHONY: all
all: build
FORCE: ;

SHELL := env ENV=$(ENV) $(SHELL)
ENV ?= dev

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	docker run --rm -v $(PWD):/app:Z -w /app golang:1.15 go mod download

build: clean dependencies build-cmd

build-cmd:
	docker run --rm -v $(PWD):/app:Z -w /app -e CGO_ENABLED=0 -e GOOS=linux -e TAG_NAME=$(ENV) golang:1.15 ./scripts/build.sh

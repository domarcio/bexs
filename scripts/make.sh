#!/usr/bin/env bash

case $1 in
  "build-cmd")
    go build -tags ${ENV} -o ./bin/cmd ./driver/cmd/main.go
  ;;
  "test")
    go clean -testcache && go test -tags "${ENV}" ./...
  ;;
esac

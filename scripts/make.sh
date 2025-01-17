#!/usr/bin/env bash

case $1 in
  "build-cmd")
    go build -tags ${ENV} -o ./bin/cmd ./driver/cmd/main.go
  ;;
  "run-cmd")
    ./bin/cmd $2
  ;;
  "build-api")
    go build -tags ${ENV} -o ./bin/api ./driver/api/main.go
  ;;
  "run-api")
    ./bin/api
  ;;
  "test")
    go clean -testcache && go test -tags "testing" ./...
  ;;
esac

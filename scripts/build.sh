#!/usr/bin/env bash

go build -a -installsuffix cgo -tags "${TAG_NAME} netgo" -installsuffix netgo -o ./bin/cmd ./driver/cmd/main.go
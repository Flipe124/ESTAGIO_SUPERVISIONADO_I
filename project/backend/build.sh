#!/bin/bash
GOOS=windows GOARCH=amd64 go build -o ./bin/benlimp-api.exe ./cmd/
GOOS=linux GOARCH=amd64 go build -o ./bin/benlimp-api.run ./cmd/

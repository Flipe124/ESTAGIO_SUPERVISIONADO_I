#!/bin/bash
. ./.env
swag fmt && swag init -g ./cmd/benlimp.go -o ./pkg/docs/ && go run -ldflags '-s -w' ./cmd/

#!/bin/bash
swag fmt && swag init -g ./cmd/benlimp.go -o ./pkg/docs/ && go run -ldflags '-s -w' ./cmd/

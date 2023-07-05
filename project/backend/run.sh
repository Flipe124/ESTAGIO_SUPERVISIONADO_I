#!/bin/bash
swag fmt && swag init -g ./cmd/openfinance.go -o ./pkg/docs/ && go run -ldflags '-s -w' ./cmd/

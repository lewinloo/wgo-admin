#!/bin/bash
printf "\nRegenerating swagger doc\n\n"
go install github.com/swaggo/swag/cmd/swag@latest
time swag init --parseDependency --generalInfo ./cmd/app/main.go
printf "\nDone.\n\n"

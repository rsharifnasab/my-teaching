#!/usr/bin/env bash

set -euo pipefail

go mod init my-app

go get -u github.com/gofiber/fiber/v2

go build main.go

./main

curl 'http://localhost:8080/' -i

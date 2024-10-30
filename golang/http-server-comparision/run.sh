#!/usr/bin/env bash

set -euo pipefail

go mod init my-app
go get -u github.com/labstack/echo/v4
go get -u github.com/gin-gonic/gin

#!/usr/bin/env bash

set -euo pipefail

go mod init my-app
go get -u github.com/tidwall/gjson

go get github.com/go-playground/validator/v10
go get github.com/xeipuuv/gojsonschema

#!/usr/bin/env bash

set -euo pipefail

go mod init my-app
go get -u github.com/tidwall/gjson

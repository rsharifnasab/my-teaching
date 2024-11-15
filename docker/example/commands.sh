#!/usr/bin/env bash

set -euo pipefail

docker build . -t my-app
docker run -p 8080:8080 my-app

docker attach ID
docker logs --follow ID

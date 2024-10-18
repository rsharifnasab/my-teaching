#!/usr/bin/env bash

set -euo pipefail

# normal get
curl https://httpbin.org/get

# print good info
curl -i https://httpbin.org/get

# verbose mode
curl -v https://httpbin.org/get

# save to file and be silent
curl -s -o result.json https://httpbin.org/get
curl -s https://httpbin.org/get >result.json

# send POST
curl -d '{"key":"value"}' -X POST "https://httpbin.org/anything"

curl -d '{"key":"value"}' -X POST "https://httpbin.org/anything" \
    -H "Content-Type: application/json"

curl -H "Authorization: Bearer YOUR_TOKEN" \
    "https://httpbin.org/anything"

# follow redirects
curl -L -X GET "https://httpbin.org/redirect/5" -H "accept: text/html"

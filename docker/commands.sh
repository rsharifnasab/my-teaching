#!/usr/bin/env bash

set -euo pipefail

dockerd

sudo -e /etc/docker/daemon.json # {"registry-mirrors": ["https://docker.iranserver.com"]}

docker run hello-world:latest
docker run ubuntu:latest
docker run alpine:latest

docker run -it ubuntu:latest /bin/sh
docker run -it alpine:latest /bin/sh

# base tools of alpine
docker run --rm -it busybox:latest sh

docker exec -it ID /bin/sh

docker ps
dockre ps -a
docker rm ID

docker pull debian
# https://hub.docker.com/_/debian

docker image ls
docker rmi ID

docker kill ID
docker stop ID

docker build
docker build -t
docker login
docker push

docker run --entry-point /bin/sh -it ubuntu:latest
docker builder prune

docker logs --follow ID
docker logs --tail ID
docker inspect ID

docker pull nginx:latest
docker create nginx:latest
docker start ID

# chaos testing
docker pause ID
docker unpause ID

# clean disk:
docker system df

docker save -o /tmp/nginx.tar nginx:latest
docker load -i /tmp/nginx

docker rm $(docker ps -a -q)

docker build . -t my-app -f 1.dockerfile
docker run -it my-app

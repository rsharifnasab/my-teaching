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

docker build .
docker build . -t my-app
docker login
docker push

docker run --entry-point /bin/sh -it ubuntu:latest
docker builder prune

docker attach ID
docker logs --follow ID
docker logs --tail ID
docker inspect ID

docker pull nginx:latest
docker create nginx:latest
docker start ID

# chaos testing
docker pause ID
docker unpause ID

# check docker disk usage
docker system df

docker save -o /tmp/nginx.tar nginx:latest
docker load -i /tmp/nginx.tar

docker rm $(docker ps -a -q)

docker build . -t my-app -f 1.dockerfile
docker run -it my-app

docker volume ls
docker volume create NAME
docker volume inspect NAME
docker volume rm NAME
docker volume prune

#                                     | image name
#                       | mount path in container
#                 | volome name
docker run -it -v NAME:/data my-app
docker run -d -v /path/on/host:/path/in/container my-app
docker run -ti -v /data busybox # anon volume

docker run -d -p 8080:80 nginx

docker network ls
docker network inspect bridge
docker network create my_network
docker network rm my_network
docker network connect my_network ID
docker network disconnect my_network ID
docker run --network my_network -d nginx:latest

docker run --network my_network --name web nginx
docker run --network my_network --name db mysql
docker run --network my_network --name web --network-alias web2 nginx # both web and web2
# ping db

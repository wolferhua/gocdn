#!/usr/bin/env bash
docker ps -a|grep gocdn |awk '{print $1}'|xargs docker kill >/dev/null 2>&1
docker ps -a|grep gocdn |awk '{print $1}'|xargs docker rm >/dev/null 2>&1

docker run \
-it \
-v ~/workspace/golang/gocdn:/go \
-v /home/wolferhua:/vhosts \
-w /go \
-p 81:8181 \
--name gocdn \
golang:1.9.4-stretch go run src/main/main.go


#!/usr/bin/env bash
docker ps -a|grep gocdn |awk '{print $1}'|xargs docker kill >/dev/null 2>&1
docker ps -a|grep gocdn |awk '{print $1}'|xargs docker rm >/dev/null 2>&1

docker run \
-it \
-v ~/workspace/golang/gocdn:/go \
-w /go \
-p 81:80 \
--name gocdn \
golang:1.9.4-stretch sh /go/start.sh


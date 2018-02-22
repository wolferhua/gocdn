#!/usr/bin/env bash

docker ps -a|grep gocdn |awk '{print $1}'|xargs docker kill >/dev/null 2>&1
docker ps -a|grep gocdn |awk '{print $1}'|xargs docker rm >/dev/null 2>&1

docker run -d \
-p 88:8181 \
--name gocdn \
wolferhua/gocdn
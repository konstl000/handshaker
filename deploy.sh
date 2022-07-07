#!/bin/bash
currDir=$(cd $(dirname $0) && pwd)
docker rmi -f handshaker:latest
docker build --no-cache $currDir -t handshaker:latest
docker run -t --env-file envfile --name handshaker handshaker:latest > shaker.log &


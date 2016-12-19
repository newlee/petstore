#!/bin/bash

docker rm dev && docker rmi dev
docker build -t dev .
docker run -ti --name dev dev /go/build.sh

if [ $? != 0 ]; then
    exit 1
fi

echo "try get bin file from docker"
docker cp dev:/go/petstore build/

echo "build production docker image"
docker rm -f petstore && docker rmi petstore
docker build -t petstore -f build/backend/Dockerfile .

echo "docker run"
docker run -tid --name petstore --publish 8090:8090 petstore

docker rm -f petstore-static && docker rmi petstore-static
docker build -t petstore-static -f build/frontend/Dockerfile .

docker run -tid --name petstore-static --publish 80:80 petstore-static
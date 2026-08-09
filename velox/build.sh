#!/bin/bash

proxy="http://192.168.31.2:1087"
tag=$(date +"%Y%m%d")

docker build \
  --build-arg HTTP_PROXY=${proxy} \
  --build-arg HTTPS_PROXY=${proxy} \
  -t velox-dev:${tag} .
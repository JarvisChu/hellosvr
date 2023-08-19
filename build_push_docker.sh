#!/bin/bash

version() {
  branch=$(git branch --show-current)
  commit=$(git rev-parse --short HEAD)
  echo "${branch}"_"${commit}"
}

SERVICE=hellosvr
REGISTRY=your_registry # e.g. docker.io
NAMESPACE=your_namespace
VERSION=$(version)

docker login --username=your_username --password=your_password ${REGISTRY}
docker build -t ${REGISTRY}/${NAMESPACE}/${SERVICE}:${VERSION} .
docker push ${REGISTRY}/${NAMESPACE}/${SERVICE}:${VERSION}
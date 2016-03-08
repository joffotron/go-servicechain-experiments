#!/bin/bash

echo "~~~~~~~~======== Golang build ========~~~~~~~~"

GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/public src/public.go
GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/backend src/backend.go
GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/crashinator src/crashinator.go


echo "~~~~~~~~======== Docker build ========~~~~~~~~"
docker build -t joffotron/public -f Dockerfile.public .
docker build -t joffotron/backend -f Dockerfile.backend  .
docker build -t joffotron/crashinator -f Dockerfile.crashinator  .
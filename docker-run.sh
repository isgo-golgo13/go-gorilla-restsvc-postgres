#! /bin/bash

TAG=1.0
docker run --name go-gorilla-restsvc-postgres -p 8080:8080 go-gorilla-restsvc-postgres:${TAG}
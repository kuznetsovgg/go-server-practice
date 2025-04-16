#!/bin/bash

docker build . -t go-server
docker run -d -p 8080:8080 --name go-server go-server
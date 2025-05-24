#!/bin/bash

docker build -t 'mikrotik-app:Dockerfile' .
docker run --detach 'mikrotik-app'
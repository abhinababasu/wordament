#!/bin/bash

# curl -fsSL get.docker.com -o get-docker.sh
# sh get-docker.sh

docker ps --filter="ancestor=bonggeek/wordament" -q | xargs docker stop
docker pull bonggeek/wordament
docker run -d --restart="always" -p 8090:8090 bonggeek/wordament
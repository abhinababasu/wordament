#!/bin/bash

VERSION=0.1
CONTAINER=wordament

echo Building go sources
go build .

echo Building Image $CONTAINER:$VERSION
docker build -t $CONTAINER:$VERSION .
docker tag $CONTAINER:$VERSION $DOCKER_ID_USER/$CONTAINER

echo Done ................................
echo 
echo Run Container using
echo docker run -d -p 8090:8090 $CONTAINER:$VERSION

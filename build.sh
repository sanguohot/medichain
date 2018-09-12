#!/bin/bash
set -o errexit
TAG=$MEDICHAIN_TAG
if [ ! $TAG ]; then  
  echo "MEDICHAIN_TAG env not set, e.g. 1.0,2.0,latest"
  exit 1
fi 
DOCKER_IMG=sanguohot/medichain:$TAG
docker build . -t ${DOCKER_IMG}
docker push ${DOCKER_IMG}

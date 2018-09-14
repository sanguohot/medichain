#!/bin/bash
set -o errexit
IMG=
if [ $1 ]; then
  IMG=$1
elif [ $MEDICHAIN_IMG ]; then
    IMG=$MEDICHAIN_IMG
else
  echo "MEDICHAIN_IMG env or param required, e.g. sanguohot/medichain:1.0,sanguohot/medichain:latest"
  exit 1
fi
docker build --force-rm=true -t ${IMG} .
docker push ${IMG}
docker images|grep none|awk '{print $3 }'|xargs docker rmi
docker rmi ${IMG}
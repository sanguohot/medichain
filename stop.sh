set -o errexit
TAG=$MEDICHAIN_TAG
if [ ! $TAG ]; then
  echo "MEDICHAIN_TAG env not set, e.g. 1.0,2.0,latest"
  exit 1
fi
IMG=sanguohot/medichain:$TAG
SERVER_PATH=/opt/medichain
docker stop ${IMG}
docker rm ${IMG}

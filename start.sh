set -o errexit
TAG=$MEDICHAIN_TAG
if [ ! $TAG ]; then  
  echo "MEDICHAIN_TAG env not set, e.g. 1.0,2.0,latest"
  exit 1
fi
IMG=sanguohot/medichain:$TAG
SERVER_PATH=/opt/medichain
docker run -it -d --name medichain \
-v ${SERVER_PATH}/key:/root/log \
-v ${SERVER_PATH}/etc:/root/etc \
-v ${SERVER_PATH}/databases:/root/databases \
-p 8443:8443 \
${IMG}

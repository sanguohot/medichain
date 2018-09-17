#!/bin/bash
set -o errexit
source shell/docker_img.sh
IMG=$(SetDockerImg $1)
if [[ ${IMG} == error:* ]]; then
    echo "${IMG}"
    exit 1
fi
SERVER_PATH=$(pwd)
docker run -it -d --name medichain \
-v ${SERVER_PATH}/key:/root/key \
-v ${SERVER_PATH}/etc:/root/etc \
-v ${SERVER_PATH}/keystore:/root/keystore \
-v ${SERVER_PATH}/databases:/root/databases \
-p 8443:8443 \
${IMG}


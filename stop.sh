#!/bin/bash
set -o errexit
NAME=medichain
docker stop ${NAME}
docker rm ${NAME}


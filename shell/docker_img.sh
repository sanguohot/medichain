#!/bin/bash
function SetDockerImg()
{
    IMG="error: MEDICHAIN_IMG env or param required, e.g. sanguohot/medichain:1.0,sanguohot/medichain:latest"
    if [ -n "$1" ]; then
      IMG=$1
    elif [ -n "$MEDICHAIN_IMG" ]; then
      IMG=$MEDICHAIN_IMG
    fi
    echo "${IMG}"
}

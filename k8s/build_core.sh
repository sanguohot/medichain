#!/bin/bash
IN_CONFIG_FILE=config_util

IN_YAML_TEMPLATE_FILE=$1
OUT_YAML_PATH=".tmp"
OUT_YAML_FILE=${OUT_YAML_PATH}/${IN_YAML_TEMPLATE_FILE/.template./.}

mkdir -p ${OUT_YAML_PATH}

config=`cat medichain/${IN_CONFIG_FILE}`
templ=`cat medichain/${IN_YAML_TEMPLATE_FILE}`
printf "$config\ncat << EOF\n$templ\nEOF" | bash > ${OUT_YAML_FILE}

#!/bin/bash
source config_util
yaml_array=(medichain-deploy.template.yaml medichain-configmap.template.yaml medichain-service.template.yaml)
for data in ${yaml_array[@]}
do
    ./build_core.sh ${data}
done



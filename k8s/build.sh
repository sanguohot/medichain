#!/bin/bash
source medichain/config_util
yaml_array=(medichain-deploy.template.yaml medichain-configmap.template.yaml medichain-service.template.yaml)
for data in ${yaml_array[@]}
do
    medichain/build_core.sh ${data}
done



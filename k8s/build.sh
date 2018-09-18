#!/bin/bash
source config_util
yaml_array=(\
medichain-secret.template.yaml \
#medichain-configmap.template.yaml \
medichain-pv-keystore.template.yaml \
medichain-pv-databases.template.yaml \
medichain-pvc-keystore.template.yaml \
medichain-pvc-databases.template.yaml \
medichain-deploy.template.yaml \
medichain-service.template.yaml  \
)
kubectl -n ${MEDICHAIN_NAMESPACE} create configmap ${MEDICHAIN_NAME}-configmap --from-file=../etc/config.json --dry-run -o yaml > .tmp/${MEDICHAIN_NAME}-configmap.yaml
echo "build ../etc/config.json to .tmp/${MEDICHAIN_NAME}-configmap.yaml successfully!"
for data in ${yaml_array[@]}
do
    ./build_core.sh ${data}
done
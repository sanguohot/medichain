#!/bin/bash
set -o errexit
echo "now going to do unit test..."
mkdir -p ./databases
export MEDICHAIN_PATH=$(pwd)
echo "with diretory ${MEDICHAIN_PATH}"
module_array=(\
chain \
datacenter \
)
for data in ${module_array[@]}
do
	echo "now test ${data} module..."
    go test -v ./${data}/ -cover=true
	if [[ $? == *FAIL* ]]; then
		exit 1
	fi
done
echo "all module test successfully!"

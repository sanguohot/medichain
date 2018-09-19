#!/bin/bash
set -o errexit
type go >/dev/null 2>&1 || { echo >&2 "go command required but it's not installed.  Aborting."; exit 1; }
echo "now going to do unit test..."
mkdir -p ./databases
export MEDICHAIN_PATH=$(pwd)
echo "with diretory ${MEDICHAIN_PATH}"
module_array=(\
chain \
datacenter \
service \
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

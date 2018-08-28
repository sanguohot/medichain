array=(EasyCns OrgsData UsersData FilesData Controller)
for data in ${array[@]}  
do  
    echo "generate ${data} ..."
	fisco-solc --overwrite --abi ${data}.sol -o build
	fisco-solc --overwrite --bin ${data}.sol -o build
	abigen --solc fisco-solc --bin=build/${data}.bin --abi=build/${data}.abi --pkg=medi --type=${data} --out=${data}.go	
done
tree ./
echo "generate all file successfully!"

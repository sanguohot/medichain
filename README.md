# medichain
更多示例请参考
https://github.com/miguelmota/ethereum-development-with-go-book

# 自动生成代码样例一
fisco-solc --abi Hello.sol | awk '/JSON ABI/{x=1;next}x' > Hello.abi
fisco-solc --bin Hello.sol | awk '/Binary:/{x=1;next}x' > Hello.bin
abigen --solc fisco-solc --bin=Hello.bin --abi=Hello.abi --pkg=hello --type=Hello --out=Hello.go
# 自动生成代码样例二
fisco-solc --abi OrgsData.sol | awk '/JSON ABI/{x=1;next}x' > OrgsData.abi;
fisco-solc --bin OrgsData.sol | awk '/Binary:/{x=1;next}x' > OrgsData.bin;
abigen --solc fisco-solc --bin=OrgsData.bin --abi=OrgsData.abi --pkg=medi --type=OrgsData --out=OrgsData.go;
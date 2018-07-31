# medichain
更多示例请参考
https://github.com/miguelmota/ethereum-development-with-go-book

自动生成代码
fisco-solc --abi Hello.sol | awk '/JSON ABI/{x=1;next}x' > Hello.abi
fisco-solc --bin Hello.sol | awk '/Binary:/{x=1;next}x' > Hello.bin
abigen --solc fisco-solc --bin=Hello.bin --abi=Hello.abi --pkg=hello --out=Hello.go
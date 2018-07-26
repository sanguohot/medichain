# medichain
更多示例请参考
https://github.com/miguelmota/ethereum-development-with-go-book

自动生成代码
fisco-solc --abi HelloWorld.sol | awk '/JSON ABI/{x=1;next}x' > HelloWorld.abi
fisco-solc --bin HelloWorld.sol | awk '/Binary:/{x=1;next}x' > HelloWorld.bin
abigen --solc fisco-solc --bin=HelloWorld.bin --abi=HelloWorld.abi --pkg=hello_world --out=HelloWorld.go
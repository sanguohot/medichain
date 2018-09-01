pragma solidity ^0.4.11;
// 专门写了一个用于验证签名的合约呀
contract Auth {
    function verifyWithPrefix( bytes32 hash, uint8 v, bytes32 r, bytes32 s) public constant returns(address retAddr) {
        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 prefixedHash = keccak256(prefix, hash);
        // 验证过程(这里就可以看出，通过v返回address)
        return ecrecover(prefixedHash, v, r, s);
    }
    function verifyWithoutPrefix( bytes32 hash, uint8 v, bytes32 r, bytes32 s) public constant returns(address retAddr) {
        return ecrecover(hash, v, r, s);
    }
}
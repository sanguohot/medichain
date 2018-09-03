pragma solidity ^0.4.11;
// 专门写了一个用于验证签名的合约呀
contract Auth {
    bytes prefix = "\x19Ethereum Signed Message:\n32";
    function verifySignatureWithPrefix( bytes32 hash, uint8 v, bytes32 r, bytes32 s) public constant returns(address retAddr) {
        bytes32 prefixedHash = sha3(prefix, hash);
        return ecrecover(prefixedHash, v, r, s);
    }
    function verifySignatureWithoutPrefix( bytes32 hash, uint8 v, bytes32 r, bytes32 s) public constant returns(address retAddr) {
        return ecrecover(hash, v, r, s);
    }
    function getSha3FromHash(bytes32 hash) public constant returns(bytes32) {
        return sha3(prefix, hash);
    }
    function getKeccak256FromHash(bytes32 hash) public constant returns(bytes32) {
        return keccak256(prefix, hash);
    }
}
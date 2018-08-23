pragma solidity ^0.4.11;

library Address {
    function getAddressFromPublicKeyV1(bytes _publicKey) internal constant returns (address) {
        return address(keccak256(_publicKey));
    }
    function getAddressFromPublicKeyV2(bytes32[2] _publicKey) internal constant returns (address) {
        return address(keccak256(_publicKey));
    }
}

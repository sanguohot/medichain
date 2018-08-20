pragma solidity ^0.4.11;
import "./Address.sol";

contract Validate is Address {
    modifier uuidNotZero(bytes16 uuid) {
        require(uuid != 0x0);
        _;
    }
    modifier publicKeyNotZero(bytes32[2] publicKey) {
        require(publicKey[0]!=0x0 && publicKey[1]!=0x0);
        _;
    }
    modifier addressNotZero(address account) {
        require(account != 0x0);
        _;
    }
    modifier stringNotEmpty(string str) {
        bytes memory temp = bytes(str);
        require(temp.length != 0x0);
        _;
    }
    modifier uintNotZero(uint u) {
        require(u != 0x0);
        _;
    }
    modifier bytes32NotZero(bytes32 b32) {
        require(b32 != 0x0);
        _;
    }
    modifier addressMatchPublicKey(address addr, bytes32[2] pub) {
        require(getAddressFromPublicKeyV2(pub) == addr);
        _;
    }
}

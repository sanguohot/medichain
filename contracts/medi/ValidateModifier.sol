pragma solidity ^0.4.11;

contract ValidateModifier {
    modifier uuidNotZero(bytes16 uuid) {
        require(uuid != 0x0);
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
}

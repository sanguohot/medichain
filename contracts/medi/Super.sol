pragma solidity ^0.4.11;
import "./ValidateModifier.sol";

contract Super is ValidateModifier {
    address owner;
    address[] supers;
    function Super() public {
        owner = msg.sender;
    }
    function isSuper(address addr) public constant returns (bool valid) {
        valid = false;
        for (uint i=0; i<supers.length; i++) {
            if (addr == supers[i]){
                valid = true;
                break;
            }
        }
        return valid;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }
    modifier onlySuper() {
        require(isSuper(msg.sender));
        _;
    }
    modifier onlySuperOrOwner() {
        require(msg.sender==owner || isSuper(msg.sender));
        _;
    }
    modifier notOutOfIndex(uint256 index) {
        require(index>=0 && index<supers.length);
        _;
    }

    // this should be done first if it is called by other contracts.
    function addSuper(address addr) public onlyOwner addressNotZero(addr) {
        supers.push(addr);
    }
    function delSuper(address addr) public onlyOwner addressNotZero(addr) {
        for (uint i=0; i<supers.length; i++) {
            if (addr == supers[i]){
                break;
            }
        }
        delete supers[i];
    }
    function getSuperSize() public constant returns (uint256) {
        return supers.length;
    }
    function getSuperByIndex(uint256 index) public notOutOfIndex(index) constant returns (address) {
        return supers[index];
    }
}

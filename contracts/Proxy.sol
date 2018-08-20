pragma solidity ^0.4.11;

//abi info base contract
contract Proxy{
    address owner;
    mapping(string => address) nameToAddressMap;
    //constrct function
    function Proxy() public {
        owner = msg.sender;
    }
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }
    //set member version
    function set(string name, address addr) public onlyOwner {
        nameToAddressMap[name] = addr;
    }
    //get member version
    function get(string name) constant public returns(address) {
        return nameToAddressMap[name];
    }
}

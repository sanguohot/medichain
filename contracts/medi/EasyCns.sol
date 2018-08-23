pragma solidity ^0.4.11;
import "./ValidateUtil.sol";

//abi info base contract
contract EasyCns {
    address owner;
    mapping(string => address) nameToAddressMap;
    
    event onSet(string name, address addr);
    
    //constrct function
    function EasyCns() public {
        owner = msg.sender;
    }
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }
    //set member version
    function set(string name, address addr) public onlyOwner {
        require(ValidateUtil.stringNotEmpty(name));
        nameToAddressMap[name] = addr;
        onSet(name, addr);
    }
    //get member version
    function get(string name) constant public returns(address) {
        return nameToAddressMap[name];
    }
}

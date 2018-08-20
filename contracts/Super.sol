pragma solidity ^0.4.11;
import "./Validate.sol";

contract Super is Validate{
    address owner;
    address[] superUsers;
    function Supper() public {
        owner = msg.sender;
    }
    function isSuper(address addr) public constant returns (bool valid) {
        valid = false;
        for (uint i=0; i<superUsers.length; i++) {
            if (addr == superUsers[i]){
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
    modifier onlySuperUser() {
        require(isSuper(msg.sender));
        _;
    }

    modifier onlySuperOrOwner() {
        require(msg.sender==owner || isSuper(msg.sender));
        _;
    }
    // this should be done first if it is called by other contracts.
    function addUserToSuperUsers(address addr) public onlyOwner addressNotZero(addr) {
        superUsers.push(addr);
    }
}

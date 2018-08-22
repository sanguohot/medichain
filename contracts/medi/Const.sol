pragma solidity ^0.4.11;

contract Const {
    string constant usersDataName = "UsersData";
    string constant orgsDataName = "OrgsData";
    string constant filesDataName = "FilesData";

    function getUsersDataName() public constant returns (string) {
        return usersDataName;
    }
    function getOrgsDataName() public constant returns (string) {
        return orgsDataName;
    }
    function getFilesDataName() public constant returns (string) {
        return filesDataName;
    }
}

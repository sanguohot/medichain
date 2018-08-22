pragma solidity ^0.4.11;
import "./EasyCns.sol";
import "./OrgsData.sol";
import "./UsersData.sol";
import "./FilesData.sol";

contract Controller {
    const string usersDataName = "";
    EasyCns easyCns;
    address orgsDataContractAddress;
    OrgsData orgsData;
    address usersDataContractAddress;
    UsersData usersData;
    address filesDataContractAddress;
    FilesData filesData;

    function UsersData(address easyCnsAddress) public {
        easyCns = EasyCns(easyCnsAddress);
    }
    function checkUsersDataOk() private returns (bool) {
        address addr = easyCns.get("UsersData");
        if(addr == 0x0){
            return false;
        }
        if(addr != usersDataContractAddress){
            usersData = UsersData(addr);
            usersDataContractAddress = addr;
        }
        return true;
    }
    function checkOrgsDataOk() private returns (bool) {
        address addr = easyCns.get("OrgsData");
        if(addr == 0x0){
            return false;
        }
        if(addr != orgsDataContractAddress){
            orgsData = OrgsData(addr);
            orgsDataContractAddress = addr;
        }
        return true;
    }
    function checkFilesDataOk() private returns (bool) {
        address addr = easyCns.get("FilesData");
        if(addr == 0x0){
            return false;
        }
        if(addr != filesDataContractAddress){
            filesData = FilesData(addr);
            filesDataContractAddress = addr;
        }
        return true;
    }
}

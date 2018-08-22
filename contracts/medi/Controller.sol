pragma solidity ^0.4.11;
import "./EasyCns.sol";
import "./OrgsData.sol";
import "./UsersData.sol";
import "./FilesData.sol";
import "./Const.sol";

contract Controller is Const {
    EasyCns easyCns;
    address orgsDataContractAddress;
    OrgsData orgsData;
    address usersDataContractAddress;
    UsersData usersData;
    address filesDataContractAddress;
    FilesData filesData;

    function Controller(address easyCnsAddress) public {
        easyCns = EasyCns(easyCnsAddress);
    }
    function checkUsersDataOk() private returns (bool) {
        address addr = easyCns.get(getUsersDataName());
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
        address addr = easyCns.get(getOrgsDataName());
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
        address addr = easyCns.get(getFilesDataName());
        if(addr == 0x0){
            return false;
        }
        if(addr != filesDataContractAddress){
            filesData = FilesData(addr);
            filesDataContractAddress = addr;
        }
        return true;
    }
    // 注意操作者都是自己，也就是msg.sender
    function addUser(bytes16 uuid, bytes16 orgUuid, bytes32[2] publicKey, bytes32 idCartNoHash, uint time)
    public {
        require(checkUsersDataOk());
        usersData.addUser(uuid, msg.sender, orgUuid, publicKey, idCartNoHash, time);
    }
    function addOrg(bytes16 uuid, bytes32[2] publicKey, bytes32 nameHash, bytes32[4] name, uint time)
    public {
        require(checkOrgsDataOk());
        orgsData.addOrg(uuid, msg.sender, publicKey, nameHash, name, time);
    }
    function addFile(bytes16 uuid, bytes16 ownerUuid, bytes32 fileType, bytes32[4] fileDesc, bytes32 keccak256Hash
    , bytes32 sha256Hash, bytes32 r, bytes32 s, uint8 v, uint time)
    public {
        require(checkUsersDataOk());
        require(checkFilesDataOk());
        // 这里可能是医生上传患者的资料
        filesData.addFile(uuid, ownerUuid, usersData.getUuidByUserAddress(msg.sender), fileType, fileDesc, keccak256Hash, sha256Hash, r, s, v, time);
    }
    function addSign(bytes16 uuid, bytes32 r, bytes32 s, uint8 v)
    public {
        require(checkUsersDataOk());
        require(checkFilesDataOk());
        filesData.addSign(uuid, usersData.getUuidByUserAddress(msg.sender), r, s, v);
    }
    // 链上的所有信息都是共享的，只要没有被删除都可以返回
    function getUserByUuid(bytes16 uuid) public constant returns (address, bytes16, bytes32[2], bytes32, uint) {
        require(checkUsersDataOk());
        return (
            usersData.getUserAddress(uuid),
            usersData.getOrgUuid(uuid),
            usersData.getPublicKey(uuid),
            usersData.getIdCartNoHash(uuid),
            usersData.getTime(uuid)
        );
    }
    function getOrgByUuid(bytes16 uuid) public constant returns (address, bytes32[2], bytes32, bytes32[4], uint) {
        require(checkOrgsDataOk());
        return (
            orgsData.getOrgAddress(uuid),
            orgsData.getPublicKey(uuid),
            orgsData.getNameHash(uuid),
            orgsData.getName(uuid),
            orgsData.getTime(uuid)
        );
    }
    function getFileByUuid(bytes16 uuid) public constant returns (bytes16, bytes16, bytes32, bytes32[4], bytes32, bytes32, uint) {
        require(checkFilesDataOk());
        return (
            filesData.getOwnerUuid(uuid),
            filesData.getUploaderUuid(uuid),
            filesData.getFileType(uuid),
            filesData.getFileDesc(uuid),
            filesData.getSha256Hash(uuid),
            filesData.getKeccak256Hash(uuid),
            filesData.getTime(uuid)
        );
    }
//    function getFileSignersAndDataByUuid(bytes16 uuid, uint256 start, uint256 limit) public constant returns (bytes16, bytes16, bytes32, bytes32[4], bytes32, bytes32, uint) {
//        require(checkFilesDataOk());
//        return (
//        filesData.getOwnerUuid(uuid),
//        filesData.getUploaderUuid(uuid),
//        filesData.getFileType(uuid),
//        filesData.getFileDesc(uuid),
//        filesData.getSha256Hash(uuid),
//        filesData.getKeccak256Hash(uuid),
//        filesData.getTime(uuid)
//        );
//    }
}

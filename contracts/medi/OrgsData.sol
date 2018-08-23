pragma solidity ^0.4.11;
import "./lib/EasyCns.sol";
import "./lib/ValidateUtil.sol";
import "./lib/Super.sol";
import "./lib/Address.sol";

contract OrgsData is Super {
    EasyCns easyCns;
    struct Org{
        bool active;
        address orgAddress;
        bytes32[2] publicKey;
        bytes32 nameHash;
        bytes32[4] name; // 华中科技大学同济医学院附属妇女儿童医疗保健中心 69 bytes
        uint time;
    }
    event onAddOrg(bytes16 uuid, address orgAddress, bytes32[2] publicKey, bytes32 nameHash, bytes32[4] name, uint time);
    event onDelOrg(bytes16 uuid);
    event onSetActive(bytes16 uuid, bool active);
    event onSetOrgAddressAndPublicKey(bytes16 uuid, address orgAddress, bytes32[2] publicKey);
    event onSetNameHashAndName(bytes16 uuid, bytes32 nameHash, bytes32[4] name);
    event onSetTime(bytes16 uuid, uint time);

    mapping(bytes16 => Org) uuidToOrgMap;
    mapping(bytes32 => bytes16) nameHashToUuidMap;
    mapping(address => bytes16) orgAddressToUuidMap;
    // 保存列表 方便以后导出 只增不减 注意检查唯一性
    bytes16[] uuidList;
    function OrgsData(address easyCnsAddress) public {
        require(Address.isContract(easyCnsAddress));
        easyCns = EasyCns(easyCnsAddress);
    }

    modifier onlyActive(bytes16 uuid) {
        require(uuidToOrgMap[uuid].active);
        _;
    }
    modifier onlyNotActive(bytes16 uuid) {
        require(!uuidToOrgMap[uuid].active);
        _;
    }
    function nameHashNotExist(bytes32 nameHash) private constant returns (bool) {
        return (nameHashToUuidMap[nameHash] == 0x0);
    }
    function nameHashExist(bytes32 nameHash) private constant returns (bool) {
        return (nameHashToUuidMap[nameHash] != 0x0);
    }
    function orgAddressNotExist(address orgAddress) private constant returns (bool) {
        return (orgAddressToUuidMap[orgAddress] == 0x0);
    }
    function orgAddressExist(address orgAddress) private constant returns (bool) {
        return (orgAddressToUuidMap[orgAddress] != 0x0);
    }
    function orgExist(bytes16 uuid) private constant returns (bool) {
        return (uuidToOrgMap[uuid].orgAddress != 0x0);
    }
    function orgNotExist(bytes16 uuid) private constant returns (bool) {
        return (uuidToOrgMap[uuid].orgAddress == 0x0);
    }
    function addOrg(bytes16 uuid, address orgAddress, bytes32[2] publicKey, bytes32 nameHash, bytes32[4] name, uint time)
    public onlySuperOrOwner onlyNotActive(uuid) {
        require(uuid!=0x0 && orgAddress!=0x0 && ValidateUtil.publicKeyNotZero(publicKey) && nameHash!=0x0 && time!=0x0);
        require(nameHashNotExist(nameHash));
        require(orgAddressNotExist(orgAddress));
        require(ValidateUtil.addressMatchPublicKey(orgAddress, publicKey));
        require(nameHash == keccak256(name));
        uuidToOrgMap[uuid] = Org(true, orgAddress, publicKey, nameHash, name, time);
        nameHashToUuidMap[nameHash] = uuid;
        orgAddressToUuidMap[orgAddress] = uuid;
        uuidList.push(uuid);
        onAddOrg(uuid, orgAddress, publicKey, nameHash, name, time);
    }
    function delOrg(bytes16 uuid)
    public onlySuperOrOwner onlyActive(uuid) {
        delete nameHashToUuidMap[uuidToOrgMap[uuid].nameHash];
        delete orgAddressToUuidMap[uuidToOrgMap[uuid].orgAddress];
        uuidToOrgMap[uuid].active = false;
        onDelOrg(uuid);
    }
    function setActive(bytes16 uuid, bool active)
    public onlySuperOrOwner {
        uuidToOrgMap[uuid].active = active;
        onSetActive(uuid, active);
    }
    // as address and publicKey are always a pair, so do not set them seperately.
    function setOrgAddressAndPublicKey(bytes16 uuid, address orgAddress, bytes32[2] publicKey)
    public onlySuperOrOwner onlyActive(uuid) {
        require(uuid!=0x0 && orgAddress!=0x0 && ValidateUtil.publicKeyNotZero(publicKey));
        require(orgAddressNotExist(orgAddress));
        require(ValidateUtil.addressMatchPublicKey(orgAddress, publicKey));
        uuidToOrgMap[uuid].orgAddress = orgAddress;
        uuidToOrgMap[uuid].publicKey = publicKey;
        orgAddressToUuidMap[orgAddress] = uuid;
        onSetOrgAddressAndPublicKey(uuid, orgAddress, publicKey);
    }
    function setNameHashAndName(bytes16 uuid, bytes32 nameHash, bytes32[4] name)
    public onlySuperOrOwner onlyActive(uuid) {
        require(uuid!=0x0 && nameHash!=0x0);
        require(nameHash == keccak256(name));
        delete nameHashToUuidMap[uuidToOrgMap[uuid].nameHash];
        uuidToOrgMap[uuid].nameHash = nameHash;
        uuidToOrgMap[uuid].name = name;
        nameHashToUuidMap[nameHash] = uuid;
        onSetNameHashAndName(uuid, nameHash, name);
    }
    function setTime(bytes16 uuid, uint time)
    public onlySuperOrOwner onlyActive(uuid) {
        require(uuid!=0x0 && time!=0x0);
        uuidToOrgMap[uuid].time = time;
        onSetTime(uuid, time);
    }

    // as it is seeable to everyone on the blockchain, so no need to check any input or the right.
    function getOrgAddress(bytes16 uuid)
    public onlyActive(uuid) constant returns (address) {
        return uuidToOrgMap[uuid].orgAddress;
    }
    function getPublicKey(bytes16 uuid)
    public onlyActive(uuid) constant returns (bytes32[2]) {
        return uuidToOrgMap[uuid].publicKey;
    }
    function getNameHash(bytes16 uuid)
    public onlyActive(uuid) constant returns (bytes32) {
        return uuidToOrgMap[uuid].nameHash;
    }
    function getName(bytes16 uuid)
    public onlyActive(uuid) constant returns (bytes32[4]) {
        return uuidToOrgMap[uuid].name;
    }
    function getTime(bytes16 uuid)
    public onlyActive(uuid) constant returns (uint) {
        return uuidToOrgMap[uuid].time;
    }
    function getUuidByNameHash(bytes32 nameHash)
    public constant returns (bytes16) {
        require(uuidToOrgMap[nameHashToUuidMap[nameHash]].active);
        return nameHashToUuidMap[nameHash];
    }
    function getUuidByOrgAddress(address orgAddress)
    public constant returns (bytes16) {
        require(uuidToOrgMap[orgAddressToUuidMap[orgAddress]].active);
        return orgAddressToUuidMap[orgAddress];
    }
    function getUuidListSize()
    public constant returns (uint256) {
        return uuidList.length;
    }
    function getUuidByIndex(uint256 index) public constant returns (bytes16) {
        require(index>=0 && index<getUuidListSize());
        return uuidList[index];
    }
    function isUuidExist(bytes16 uuid)
    public constant returns (bool) {
        if (uuidToOrgMap[uuid].active){
            return true;
        }
        return false;
    }
}

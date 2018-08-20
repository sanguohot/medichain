pragma solidity ^0.4.11;
import "./Proxy.sol";
import "./Validate.sol";
import "./Super.sol";

contract OrgsData  is Validate,Super {
    Proxy proxy;
    struct Org{
        bool active;
        address orgAddress;
        bytes32[2] publicKey;
        bytes32 nameHash;
        bytes32[4] name; // 华中科技大学同济医学院附属妇女儿童医疗保健中心 69 bytes
        uint time;
    }
    event onAddOrgCore(bytes16 uuid, address orgAddress, bytes32[2] publicKey, bytes32 nameHash, bytes32[4] name, uint time);
    event onDelOrg(bytes16 uuid);
    event onSetActive(bytes16 uuid, bool active);
    event onSetOrgAddressAndPublicKey(bytes16 uuid, address orgAddress, bytes32[2] publicKey);
    event onSetNameHashAndNameCore(bytes16 uuid, bytes32 nameHash, bytes32[4] name);
    event onSetTime(bytes16 uuid, uint time);

    mapping(bytes16 => Org) uuidToOrgMap;
    mapping(bytes32 => bytes16) nameHashToUuidMap;
    mapping(address => bytes16) orgAddressToUuidMap;
    // 保存用户列表 方便以后导出 只增不减 注意检查唯一性
    bytes16[] uuidList;
    function OrgsData(address proxyAddress) public {
        proxy = Proxy(proxyAddress);
    }

    modifier nameHashNotExist(bytes32 nameHash) {
        require(nameHashToUuidMap[nameHash] == 0x0);
        _;
    }
    modifier nameHashExist(bytes32 nameHash) {
        require(nameHashToUuidMap[nameHash] != 0x0);
        _;
    }
    modifier orgAddressNotExist(address orgAddress) {
        require(orgAddressToUuidMap[orgAddress] == 0x0);
        _;
    }
    modifier orgAddressExist(address orgAddress) {
        require(orgAddressToUuidMap[orgAddress] != 0x0);
        _;
    }
    modifier orgExist(bytes16 uuid) {
        require(uuidToOrgMap[uuid].orgAddress != 0x0);
        _;
    }
    modifier orgNotExist(bytes16 uuid) {
        require(uuidToOrgMap[uuid].orgAddress == 0x0);
        _;
    }
    modifier onlyActive(bytes16 uuid) {
        require(uuidToOrgMap[uuid].active);
        _;
    }
    modifier onlyNotActive(bytes16 uuid) {
        require(!uuidToOrgMap[uuid].active);
        _;
    }
    function addOrgCore(bytes16 uuid, address orgAddress, bytes32[2] publicKey, bytes32 nameHash, bytes32[4] name, uint time)
    private onlySuperOrOwner orgAddressNotExist(orgAddress) publicKeyNotZero(publicKey) uintNotZero(time)
    addressMatchPublicKey(orgAddress, publicKey) nameHashNotExist(nameHash){
        uuidToOrgMap[uuid] = Org(true, orgAddress, publicKey, nameHash, name, time);
        nameHashToUuidMap[nameHash] = uuid;
        orgAddressToUuidMap[orgAddress] = uuid;
        uuidList.push(uuid);
        onAddOrgCore(uuid, orgAddress, publicKey, nameHash, name, time);
    }
    function addOrg(bytes16 uuid, address orgAddress, bytes32[2] publicKey, bytes32[4] name, uint time)
    public {
        bytes32 nameHash = keccak256(name);
        addOrgCore(uuid, orgAddress, publicKey, nameHash, name, time);
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
    public onlySuperOrOwner onlyActive(uuid) orgAddressNotExist(orgAddress) publicKeyNotZero(publicKey) addressMatchPublicKey(orgAddress, publicKey) {
        uuidToOrgMap[uuid].orgAddress = orgAddress;
        uuidToOrgMap[uuid].publicKey = publicKey;
        orgAddressToUuidMap[orgAddress] = uuid;
        onSetOrgAddressAndPublicKey(uuid, orgAddress, publicKey);
    }
    function setNameHashAndNameCore(bytes16 uuid, bytes32 nameHash, bytes32[4] name)
    private onlySuperOrOwner onlyActive(uuid) bytes32NotZero(nameHash) nameHashNotExist(nameHash) {
        delete nameHashToUuidMap[uuidToOrgMap[uuid].nameHash];
        uuidToOrgMap[uuid].nameHash = nameHash;
        uuidToOrgMap[uuid].name = name;
        nameHashToUuidMap[nameHash] = uuid;
        onSetNameHashAndNameCore(uuid, nameHash, name);
    }
    function setNameHashAndName(bytes16 uuid, bytes32[4] name)
    public {
        bytes32 nameHash = keccak256(name);
        setNameHashAndNameCore(uuid, nameHash, name);
    }
    function setTime(bytes16 uuid, uint time)
    public onlySuperOrOwner onlyActive(uuid) uintNotZero(time) {
        uuidToOrgMap[uuid].time = time;
        onSetTime(uuid, time);
    }

    // as it is seeable to everyone on the blockchain, so no need to check any input or the right.
    function getOrgAddress(bytes16 uuid)
    public constant returns (address) {
        return uuidToOrgMap[uuid].orgAddress;
    }
    function getPublicKey(bytes16 uuid)
    public constant returns (bytes32[2]) {
        return uuidToOrgMap[uuid].publicKey;
    }
    function getNameHash(bytes16 uuid)
    public constant returns (bytes32) {
        return uuidToOrgMap[uuid].nameHash;
    }
    function getTime(bytes16 uuid)
    public constant returns (uint) {
        return uuidToOrgMap[uuid].time;
    }
    function getUuidByNameHash(bytes32 nameHash)
    public constant returns (bytes16) {
        return nameHashToUuidMap[nameHash];
    }
    function getUuidByOrgAddress(address orgAddress)
    public constant returns (bytes16) {
        return orgAddressToUuidMap[orgAddress];
    }
    function getUuidListSize()
    public constant returns (uint256) {
        return uuidList.length;
    }
    function getUuidByIndex(uint256 index) public constant returns (bytes16) {
        return uuidList[index];
    }
    function isUuidExist(bytes16 uuid)
    public constant returns (bool) {
        if (uuidToOrgMap[uuid].orgAddress != 0x0){
            return true;
        }
        return false;
    }
}

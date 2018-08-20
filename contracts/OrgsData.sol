pragma solidity ^0.4.11;
import "./Proxy.sol";

contract OrgsData{
    address owner;
    string[] superOrgs;
    Proxy proxy;
    struct Org{
        bool active;
        address orgAddress;
        bytes32[2] publicKey;
        bytes32 nameHash;
        bytes32[4] name; // 华中科技大学同济医学院附属妇女儿童医疗保健中心 69 bytes
        uint time;
    }

    mapping(bytes16 => Org) uuidToOrgMap;
    mapping(bytes32 => bytes16) nameHashToUuidMap;
    mapping(address => bytes16) orgAddressToUuidMap;
    // 保存机构列表 方便以后导出 只增不减 注意检查唯一性
    bytes16[] uuidList;
    function OrgsData(address proxyAddress) public {
        owner = msg.sender;
        proxy = Proxy(proxyAddress);
    }
    function isSuperOrg() public constant returns (bool valid) {
        valid = false;
        for (uint i=0; i<superOrgs.length; i++) {
            if (msg.sender == proxy.get(superOrgs[i])){
                valid = true;
                break;
            }
        }
        return valid;
    }
    modifier uuidNotZero(bytes16 uuid) {
        require(uuid != 0x0);
        _;
    }
    modifier publicKeyNotZero(bytes32[2] publicKey) {
        require(publicKey[0]!=0x0 && publicKey[1]!=0x0);
        _;
    }
    modifier addressNotZero(address account) {
        require(account != 0x0);
        _;
    }
    modifier stringNotEmpty(string str) {
        require(str != "");
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
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }
    modifier onlySuperOrg() {
        require(isSuperOrg());
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
    modifier onlySuperOrgOrOwner() {
        require(msg.sender==owner || isSuperOrg());
        _;
    }

    // this should be done first if it is called by other contracts.
    function addOrgToSuperOrgs(string name) public onlyOwner stringNotEmpty(name) {
        superOrgs.push(name);
    }
    function addOrg(bytes16 uuid, address orgAddress, bytes32[2] publicKey, bytes32 nameHash, uint time)
    public onlySuperOrgOrOwner orgAddressNotExist(orgAddress) publicKeyNotZero(publicKey) nameHashNotExist(nameHash) uintNotZero(time) {
        uuidToOrgMap[uuid] = Org(true, orgAddress, publicKey, nameHash, time);
        nameHashToUuidMap[nameHash] = uuid;
        orgAddressToUuidMap[orgAddress] = uuid;
        uuidList.push(uuid);
    }
    function setActive(bytes16 uuid, bool active)
    public onlyOwner {
        uuidToOrgMap[uuid].active = active;
    }
    function setOrgAddress(bytes16 uuid, address orgAddress)
    public onlySuperOrgOrOwner onlyActive(uuid) addressNotZero(orgAddress) {
        uuidToOrgMap[uuid].orgAddress = orgAddress;
    }
    function setPublicKey(bytes16 uuid, bytes32[2] publicKey)
    public onlySuperOrgOrOwner onlyActive(uuid) publicKeyNotZero(publicKey) {
        uuidToOrgMap[uuid].publicKey = publicKey;
    }
    function setNameHash(bytes16 uuid, bytes32 nameHash)
    public onlySuperOrgOrOwner onlyActive(uuid) bytes32NotZero(nameHash) nameHashNotExist(nameHash) {
        nameHashToUuidMap[uuidToOrgMap[uuid].nameHash] = 0x0;
        uuidToOrgMap[uuid].nameHash = nameHash;
        nameHashToUuidMap[nameHash] = uuid;
    }
    function setTime(bytes16 uuid, uint time)
    public onlySuperOrgOrOwner onlyActive(uuid) uintNotZero(time) {
        uuidToOrgMap[uuid].time = time;
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
}

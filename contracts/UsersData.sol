pragma solidity ^0.4.11;
import "./Proxy.sol";
import "./Validate.sol";
import "./Super.sol";

contract UsersData is Validate,Super {
    Proxy proxy;
    struct User{
        bool active;
        address userAddress;
        bytes16 orgUuid;
        bytes32[2] publicKey;
        bytes32 idCartNoHash;
        uint time;
    }

    mapping(bytes16 => User) uuidToUserMap;
    mapping(bytes32 => bytes16) idCartNoHashToUuidMap;
    mapping(address => bytes16) userAddressToUuidMap;
    // 保存用户列表 方便以后导出 只增不减 注意检查唯一性
    bytes16[] uuidList;
    function UsersData(address proxyAddress) public {
        proxy = Proxy(proxyAddress);
    }

    modifier idCartNoHashNotExist(bytes32 idCartNoHash) {
        require(idCartNoHashToUuidMap[idCartNoHash] == 0x0);
        _;
    }
    modifier idCartNoHashExist(bytes32 idCartNoHash) {
        require(idCartNoHashToUuidMap[idCartNoHash] != 0x0);
        _;
    }
    modifier userAddressNotExist(address userAddress) {
        require(userAddressToUuidMap[userAddress] == 0x0);
        _;
    }
    modifier userAddressExist(address userAddress) {
        require(userAddressToUuidMap[userAddress] != 0x0);
        _;
    }
    modifier userExist(bytes16 uuid) {
        require(uuidToUserMap[uuid].userAddress != 0x0);
        _;
    }
    modifier userNotExist(bytes16 uuid) {
        require(uuidToUserMap[uuid].userAddress == 0x0);
        _;
    }
    modifier onlyActive(bytes16 uuid) {
        require(uuidToUserMap[uuid].active);
        _;
    }
    modifier onlyNotActive(bytes16 uuid) {
        require(!uuidToUserMap[uuid].active);
        _;
    }

    function addUser(bytes16 uuid, address userAddress, bytes16 orgUuid, bytes32[2] publicKey, bytes32 idCartNoHash, uint time)
    public onlySuperOrOwner userAddressNotExist(userAddress) idCartNoHashNotExist(idCartNoHash) uintNotZero(time)
    publicKeyNotZero(publicKey) addressMatchPublicKey(userAddress, publicKey) {
        uuidToUserMap[uuid] = User(true, userAddress, orgUuid, publicKey, idCartNoHash, time);
        idCartNoHashToUuidMap[idCartNoHash] = uuid;
        userAddressToUuidMap[userAddress] = uuid;
        uuidList.push(uuid);
    }
    function delUser(bytes16 uuid)
    public onlySuperOrOwner onlyActive(uuid) {
        delete idCartNoHashToUuidMap[uuidToUserMap[uuid].idCartNoHash];
        delete userAddressToUuidMap[uuidToUserMap[uuid].userAddress];
        setActive(uuid, false);
    }
    function setActive(bytes16 uuid, bool active)
    public onlySuperOrOwner {
        uuidToUserMap[uuid].active = active;
    }
    // as address and publicKey are always a pair, so do not set them seperately.
    function setUserAddressAndPublicKey(bytes16 uuid, address userAddress, bytes32[2] publicKey)
    public onlySuperOrOwner onlyActive(uuid) userAddressNotExist(userAddress) publicKeyNotZero(publicKey) addressMatchPublicKey(userAddress, publicKey) {
        uuidToUserMap[uuid].userAddress = userAddress;
        uuidToUserMap[uuid].publicKey = publicKey;
        userAddressToUuidMap[userAddress] = uuid;
    }
    function setOrgUuid(bytes16 uuid, bytes16 orgUuid)
    public onlySuperOrOwner onlyActive(uuid){
        uuidToUserMap[uuid].orgUuid = orgUuid;
    }
    function setIdCartNoHash(bytes16 uuid, bytes32 idCartNoHash)
    public onlySuperOrOwner onlyActive(uuid) bytes32NotZero(idCartNoHash) idCartNoHashNotExist(idCartNoHash) {
        delete idCartNoHashToUuidMap[uuidToUserMap[uuid].idCartNoHash];
        uuidToUserMap[uuid].idCartNoHash = idCartNoHash;
        idCartNoHashToUuidMap[idCartNoHash] = uuid;
    }
    function setTime(bytes16 uuid, uint time)
    public onlySuperOrOwner onlyActive(uuid) uintNotZero(time) {
        uuidToUserMap[uuid].time = time;
    }

    // as it is seeable to everyone on the blockchain, so no need to check any input or the right.
    function getUserAddress(bytes16 uuid)
    public constant returns (address) {
        return uuidToUserMap[uuid].userAddress;
    }
    function getOrgUuid(bytes16 uuid)
    public constant returns (bytes16) {
        return uuidToUserMap[uuid].orgUuid;
    }
    function getPublicKey(bytes16 uuid)
    public constant returns (bytes32[2]) {
        return uuidToUserMap[uuid].publicKey;
    }
    function getIdCartNoHash(bytes16 uuid)
    public constant returns (bytes32) {
        return uuidToUserMap[uuid].idCartNoHash;
    }
    function getTime(bytes16 uuid)
    public constant returns (uint) {
        return uuidToUserMap[uuid].time;
    }
    function getUuidByIdCartNoHash(bytes32 idCartNoHash)
    public constant returns (bytes16) {
        return idCartNoHashToUuidMap[idCartNoHash];
    }
    function getUuidByUserAddress(address userAddress)
    public constant returns (bytes16) {
        return userAddressToUuidMap[userAddress];
    }
    function getUuidListSize()
    public constant returns (uint256) {
        return uuidList.length;
    }
    function getUuidByIndex(uint256 index) public constant returns (bytes16) {
        return uuidList[index];
    }
}

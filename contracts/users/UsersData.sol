pragma solidity ^0.4.2;
import "./ContractBase.sol";

contract UsersData is ContractBase("v1") {
    address owner;
    address[] superUsers;
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
    function UsersData() public {
        owner = msg.sender;
    }
    function isSuperUser() public constant returns (bool valid) {
        valid = false;
        for (uint i=0; i<superUsers.length; i++) {
            if (msg.sender == superUsers[i]){
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
    modifier uintNotZero(uint u) {
        require(u != 0x0);
        _;
    }
    modifier bytes32NotZero(bytes32 b32) {
        require(b32 != 0x0);
        _;
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
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }
    modifier onlySuperUser() {
        require(isSuperUser());
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
    modifier onlySuperUserOrOwner() {
        require(msg.sender==owner || isSuperUser());
        _;
    }

    // this should be done first if it is called by other contracts.
    function addUserToSuperUsers(address account) public onlyOwner addressNotZero(account) {
        superUsers.push(account);
    }
    function addUser(bytes16 uuid, address userAddress, bytes16 orgUuid, bytes32[2] publicKey, bytes32 idCartNoHash, uint time)
    public onlySuperUserOrOwner userAddressNotExist(userAddress) publicKeyNotZero(publicKey) idCartNoHashNotExist(idCartNoHash) uintNotZero(time) {
        uuidToUserMap[uuid] = User(true, userAddress, orgUuid, publicKey, idCartNoHash, time);
        idCartNoHashToUuidMap[idCartNoHash] = uuid;
        userAddressToUuidMap[userAddress] = uuid;
        uuidList.push(uuid);
    }
    function setActive(bytes16 uuid, bool active)
    public onlyOwner {
        uuidToUserMap[uuid].active = active;
    }
    function setUserAddress(bytes16 uuid, address userAddress)
    public onlySuperUserOrOwner onlyActive(uuid) addressNotZero(userAddress) {
        uuidToUserMap[uuid].userAddress = userAddress;
    }
    function setOrgUuid(bytes16 uuid, bytes16 orgUuid)
    public onlySuperUserOrOwner onlyActive(uuid){
        uuidToUserMap[uuid].orgUuid = orgUuid;
    }
    function setPublicKey(bytes16 uuid, bytes32[2] publicKey)
    public onlySuperUserOrOwner onlyActive(uuid) publicKeyNotZero(publicKey) {
        uuidToUserMap[uuid].publicKey = publicKey;
    }
    function setIdCartNoHash(bytes16 uuid, bytes32 idCartNoHash)
    public onlySuperUserOrOwner onlyActive(uuid) bytes32NotZero(idCartNoHash) idCartNoHashNotExist(idCartNoHash) {
        idCartNoHashToUuidMap[uuidToUserMap[uuid].idCartNoHash] = 0x0;
        uuidToUserMap[uuid].idCartNoHash = idCartNoHash;
        idCartNoHashToUuidMap[idCartNoHash] = uuid;
    }
    function setTime(bytes16 uuid, uint time)
    public onlySuperUserOrOwner onlyActive(uuid) uintNotZero(time) {
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

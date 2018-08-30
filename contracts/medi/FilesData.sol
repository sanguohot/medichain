pragma solidity ^0.4.11;
import "./EasyCns.sol";
import "./lib/Super.sol";
import "./UsersData.sol";
import "./OrgsData.sol";
import "./lib/Const.sol";

contract FilesData is Super {
    EasyCns easyCns;
    address usersDataContractAddress;
    address orgsDataContractAddress;
    UsersData usersData;
    OrgsData orgsData;
    struct File{
        bool active;
        bytes16 ownerUuid;
        bytes16 uploaderUuid;
        // 文件类型的keccak256 hash，类型在外部定义，在外部代码解释
        // 可以是标签云的hash，比如keccak256("门诊、急诊、住院")=0x116d70e0d1a135c065fc75f30d0035554486f00333e5505f39d414e775ddc42d
        // 由于类型比较规则，比较重要，我们作为hash存证，但同时由于其开放性和扩展性，我们交由上层应用定义
        bytes32 fileType;
        // 最大128字节的文件描述，方便后续查看分析
        // 由于描述是很不规则的信息，只起辅助作用，我们存储原文
        bytes32[4] fileDesc;
        bytes32 keccak256Hash;
        bytes32 sha256Hash;
        // 针对keccak256Hash的数字签名，增加前需要进行ecrecover校验，判断签名者身份
        // 默认包含上传者签名
        uint time;
        bytes32[] r;
        bytes32[] s;
        uint8[] v;
        bytes16[] signerUuidList;
        mapping(bytes16 => bool) signerUuidMap;
    }
    event onAddFile(bytes16 uuid, bytes16 ownerUuid, bytes16 uploaderUuid, bytes32 fileType, bytes32[4] fileDesc, bytes32 keccak256Hash
    , bytes32 sha256Hash, bytes32 r, bytes32 s, uint8 v, uint time);
    event onDelFile(bytes16 uuid);
    event onSetActive(bytes16 uuid, bool active);
    event onSetOwnerUuid(bytes16 uuid, bytes16 ownerUuid);
    event onSetUploaderUuid(bytes16 uuid, bytes16 uploaderUuid);
    event onSetFileType(bytes16 uuid, bytes32 fileType);
    event onSetFileDesc(bytes16 uuid, bytes32[4] fileDesc);
    event onAddSign(bytes16 uuid, bytes16 userUuid, bytes32 r, bytes32 s, uint8 v);
    event onSetTime(bytes16 uuid, uint time);

    mapping(bytes16 => File) uuidToFileMap;
    mapping(bytes32 => bytes16) keccak256HashToUuidMap;
    mapping(bytes32 => bytes16) sha256HashToUuidMap;
    // 保存列表 方便以后导出 只增不减 注意检查唯一性
    bytes16[] uuidList;
    function FilesData(address easyCnsAddress) public {
        require(Address.isContract(easyCnsAddress));
        easyCns = EasyCns(easyCnsAddress);
    }
    modifier onlyActive(bytes16 uuid) {
        require(uuidToFileMap[uuid].active);
        _;
    }
    modifier onlyNotActive(bytes16 uuid) {
        require(!uuidToFileMap[uuid].active);
        _;
    }
    function keccak256HashNotExist(bytes32 keccak256Hash) private constant returns (bool) {
        return (keccak256HashToUuidMap[keccak256Hash] == 0x0);
    }
    function keccak256HashExist(bytes32 keccak256Hash) private constant returns (bool) {
        return (keccak256HashToUuidMap[keccak256Hash] != 0x0);
    }
    function sha256HashNotExist(bytes32 sha256Hash) private constant returns (bool) {
        return (sha256HashToUuidMap[sha256Hash] == 0x0);
    }
    function sha256HashExist(bytes32 sha256Hash) private constant returns (bool) {
        return (sha256HashToUuidMap[sha256Hash] != 0x0);
    }
    function fileExist(bytes16 uuid) private constant returns (bool) {
        return (uuidToFileMap[uuid].keccak256Hash != 0x0);
    }
    function fileNotExist(bytes16 uuid) private constant returns (bool) {
        return (uuidToFileMap[uuid].keccak256Hash == 0x0);
    }

    function checkUsersDataOk() private returns (bool) {
        address addr = easyCns.get(Const.getUsersDataName());
        if(!Address.isContract(addr)){
            return false;
        }
        if(addr != usersDataContractAddress){
            usersData = UsersData(addr);
            usersDataContractAddress = addr;
        }
        return true;
    }
    function checkOrgsDataOk() private returns (bool) {
        address addr = easyCns.get(Const.getOrgsDataName());
        if(!Address.isContract(addr)){
            return false;
        }
        if(addr != orgsDataContractAddress){
            orgsData = OrgsData(addr);
            orgsDataContractAddress = addr;
        }
        return true;
    }
    function addFile(bytes16 uuid, bytes16 ownerUuid, bytes16 uploaderUuid, bytes32 fileType, bytes32[4] fileDesc, bytes32 keccak256Hash
    , bytes32 sha256Hash, bytes32 r, bytes32 s, uint8 v, uint time)
    public onlySuperOrOwner onlyNotActive(uuid) {
        require(uuid!=0x0 && ownerUuid!=0x0 && uploaderUuid!=0x0 && fileType!=0x0 && keccak256Hash!=0x0 && sha256Hash!=0x0 && r!=0x0 && s!=0x0 && time!=0x0);
        require(fileNotExist(uuid));
        require(sha256HashNotExist(sha256Hash));
        require(keccak256HashNotExist(keccak256Hash));
        require(sha256Hash != keccak256Hash);
        require(checkUsersDataOk());
        require(usersData.isUuidExist(ownerUuid));
        require(usersData.isUuidExist(uploaderUuid));
        // 上传文件的用户可以不属于一个机构
        if(usersData.getOrgUuid(uploaderUuid) != 0x0){
            require(checkOrgsDataOk());
            require(orgsData.isUuidExist(usersData.getOrgUuid(uploaderUuid)));
        }
        require(usersData.getUserAddress(uploaderUuid) == ecrecover(keccak256Hash, v, r, s));
        File memory file;
        file.active = true;
        file.ownerUuid = ownerUuid;
        file.uploaderUuid = uploaderUuid;
        file.fileType = fileType;
        file.fileDesc = fileDesc;
        file.keccak256Hash = keccak256Hash;
        file.sha256Hash = sha256Hash;
        file.time = time;
        uuidToFileMap[uuid] = file;
        uuidToFileMap[uuid].r.push(r);
        uuidToFileMap[uuid].s.push(s);
        uuidToFileMap[uuid].v.push(v);
        uuidToFileMap[uuid].signerUuidList.push(uploaderUuid);
        uuidToFileMap[uuid].signerUuidMap[uploaderUuid] = true;
        keccak256HashToUuidMap[keccak256Hash] = uuid;
        sha256HashToUuidMap[sha256Hash] = uuid;
        uuidList.push(uuid);
        onAddFile(uuid, ownerUuid, uploaderUuid, fileType, fileDesc, keccak256Hash
        , sha256Hash, r, s, v, time);
    }
    function delFile(bytes16 uuid)
    public onlySuperOrOwner onlyActive(uuid) {
        delete keccak256HashToUuidMap[uuidToFileMap[uuid].keccak256Hash];
        delete sha256HashToUuidMap[uuidToFileMap[uuid].sha256Hash];
        uuidToFileMap[uuid].active = false;
        onDelFile(uuid);
    }
    function setActive(bytes16 uuid, bool active)
    public onlySuperOrOwner {
        uuidToFileMap[uuid].active = active;
        onSetActive(uuid, active);
    }
    function setOwnerUuid(bytes16 uuid, bytes16 ownerUuid)
    public onlySuperOrOwner onlyActive(uuid) {
        require(uuid!=0x0 && ownerUuid!=0x0);
        require(checkUsersDataOk());
        require(usersData.isUuidExist(ownerUuid));
        uuidToFileMap[uuid].ownerUuid = ownerUuid;
        onSetOwnerUuid(uuid, ownerUuid);
    }
    function setUploaderUuid(bytes16 uuid, bytes16 uploaderUuid)
    public onlySuperOrOwner onlyActive(uuid) {
        require(uuid!=0x0 && uploaderUuid!=0x0);
        require(checkUsersDataOk());
        require(usersData.isUuidExist(uploaderUuid));
        uuidToFileMap[uuid].uploaderUuid = uploaderUuid;
        onSetUploaderUuid(uuid, uploaderUuid);
    }
    function setFileType(bytes16 uuid, bytes32 fileType)
    public onlySuperOrOwner onlyActive(uuid) {
        require(uuid!=0x0 && fileType!=0x0);
        uuidToFileMap[uuid].fileType = fileType;
        onSetFileType(uuid, fileType);
    }
    // fileDesc can be empty
    function setFileDesc(bytes16 uuid, bytes32[4] fileDesc)
    public onlySuperOrOwner onlyActive(uuid) {
        require(uuid != 0x0);
        uuidToFileMap[uuid].fileDesc = fileDesc;
        onSetFileDesc(uuid, fileDesc);
    }
    // as keccak256Hash, sha256Hash, signer, sign data are always together, and are possibly relate to many people, so no change please.
    //    function setKeccak256AndSha256Hash(bytes16 uuid, bytes32 keccak256Hash, bytes32 sha256Hash)
    //    public onlySuperOrOwner onlyActive(uuid) {
    //        require(uuid!=0x0 && keccak256Hash!=0x0 && sha256Hash!=0x0);
    //        require(sha256HashNotExist(sha256Hash));
    //        require(keccak256HashNotExist(keccak256Hash));
    //        require(sha256Hash != keccak256Hash);
    //        delete keccak256HashToUuidMap[uuidToFileMap[uuid].keccak256Hash];
    //        delete sha256HashToUuidMap[uuidToFileMap[uuid].sha256Hash];
    //        uuidToFileMap[uuid].keccak256Hash = keccak256Hash;
    //        uuidToFileMap[uuid].sha256Hash = sha256Hash;
    //        keccak256HashToUuidMap[keccak256Hash] = uuid;
    //        sha256HashToUuidMap[sha256Hash] = uuid;
    //    }
    function addSign(bytes16 uuid, bytes16 userUuid, bytes32 r, bytes32 s, uint8 v)
    public onlySuperOrOwner onlyActive(uuid) {
        require(uuid!=0x0 && userUuid!=0x0 && r!=0x0 && s!=0x0);
        require(!uuidToFileMap[uuid].signerUuidMap[userUuid]);
        require(checkUsersDataOk());
        require(usersData.isUuidExist(userUuid));
        require(usersData.getUserAddress(userUuid) == ecrecover(uuidToFileMap[uuid].keccak256Hash, v, r, s));
        uuidToFileMap[uuid].r.push(r);
        uuidToFileMap[uuid].s.push(s);
        uuidToFileMap[uuid].v.push(v);
        uuidToFileMap[uuid].signerUuidList.push(userUuid);
        uuidToFileMap[uuid].signerUuidMap[userUuid] = true;
        onAddSign(uuid, userUuid, r, s, v);
    }
    function setTime(bytes16 uuid, uint time)
    public onlySuperOrOwner onlyActive(uuid) {
        require(time != 0x0);
        uuidToFileMap[uuid].time = time;
        onSetTime(uuid, time);
    }

    // as it is seeable to everyone on the blockchain, so no need to check any input or the right.
    function getOwnerUuid(bytes16 uuid)
    public onlyActive(uuid) constant returns (bytes16) {
        return uuidToFileMap[uuid].ownerUuid;
    }
    function getUploaderUuid(bytes16 uuid)
    public onlyActive(uuid) constant returns (bytes16) {
        return uuidToFileMap[uuid].uploaderUuid;
    }
    function getSha256Hash(bytes16 uuid)
    public onlyActive(uuid) constant returns (bytes32) {
        return uuidToFileMap[uuid].sha256Hash;
    }
    function getKeccak256Hash(bytes16 uuid)
    public onlyActive(uuid) constant returns (bytes32) {
        return uuidToFileMap[uuid].keccak256Hash;
    }
    function getFileType(bytes16 uuid)
    public onlyActive(uuid) constant returns (bytes32) {
        return uuidToFileMap[uuid].fileType;
    }
    function getFileDesc(bytes16 uuid)
    public onlyActive(uuid) constant returns (bytes32[4]) {
        return uuidToFileMap[uuid].fileDesc;
    }
    function getTime(bytes16 uuid)
    public onlyActive(uuid) constant returns (uint) {
        return uuidToFileMap[uuid].time;
    }
    function getUuidByKeccak256Hash(bytes32 keccak256Hash)
    public constant returns (bytes16) {
        require(uuidToFileMap[keccak256HashToUuidMap[keccak256Hash]].active);
        return keccak256HashToUuidMap[keccak256Hash];
    }
    function getUuidBySha256Hash(bytes32 sha256Hash)
    public constant returns (bytes16) {
        require(uuidToFileMap[sha256HashToUuidMap[sha256Hash]].active);
        return sha256HashToUuidMap[sha256Hash];
    }
    function getUuidListSize()
    public constant returns (uint256) {
        return uuidList.length;
    }
    function getUuidByIndex(uint256 index) public constant returns (bytes16) {
        require(index>=0 && index<getUuidListSize());
        return uuidList[index];
    }
    function getFileSignerSize(bytes16 uuid) public onlyActive(uuid) constant returns (uint256) {
        return uuidToFileMap[uuid].signerUuidList.length;
    }
    function getFileSignerUuidByIndex(bytes16 uuid, uint256 index) public onlyActive(uuid) constant returns (bytes16) {
        require(index>=0 && index<getFileSignerSize(uuid));
        return uuidToFileMap[uuid].signerUuidList[index];
    }
    function getFileSignDataByIndex(bytes16 uuid, uint256 index) public onlyActive(uuid) constant returns (bytes32, bytes32, uint8) {
        require(index>=0 && index<getFileSignerSize(uuid));
        return (uuidToFileMap[uuid].r[index], uuidToFileMap[uuid].s[index], uuidToFileMap[uuid].v[index]);
    }
    function isUuidExist(bytes16 uuid)
    public constant returns (bool) {
        if (uuidToFileMap[uuid].active){
            return true;
        }
        return false;
    }
}

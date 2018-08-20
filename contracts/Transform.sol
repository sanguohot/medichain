pragma solidity ^0.4.11;

contract Transform {
    function bytes32ArrayToString(bytes32[] data) returns (string) {
        bytes memory bytesString = new bytes(data.length * 32);
        uint urlLength;
        for (uint i = 0; i< data.length; i++) {
            for (uint j = 0; j < 32; j++) {
                byte char = byte(bytes32(uint(data[i]) * 2 ** (8 * j)));
                if (char != 0) {
                    bytesString[urlLength] = char;
                    urlLength += 1;
                }
            }
        }
        bytes memory bytesStringTrimmed = new bytes(urlLength);
        for (i = 0; i < urlLength; i++) {
            bytesStringTrimmed[i] = bytesString[i];
        }
        return string(bytesStringTrimmed);
    }
    function StringToBytes(string memory source) returns (bytes result) {
        return bytes(source);
    }

    function stringToBytes32(string memory source) returns (bytes32 result) {
        assembly {
            result := mload(add(source, 32))
        }
    }
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package medi

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// UsersDataABI is the input ABI used to generate the binding from.
const UsersDataABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"setActive\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"idCartNoHash\",\"type\":\"bytes32\"}],\"name\":\"getUuidByIdCartNoHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"idCartNoHash\",\"type\":\"bytes32\"}],\"name\":\"setIdCartNoHash\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[2]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSuperSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"setTime\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getOrgUuid\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"userAddress\",\"type\":\"address\"},{\"name\":\"publicKey\",\"type\":\"bytes32[2]\"}],\"name\":\"setUserAddressAndPublicKey\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"delSuper\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUuidListSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getUuidByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"userAddress\",\"type\":\"address\"},{\"name\":\"orgUuid\",\"type\":\"bytes16\"},{\"name\":\"publicKey\",\"type\":\"bytes32[2]\"},{\"name\":\"idCartNoHash\",\"type\":\"bytes32\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"addUser\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"isUuidExist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUuidByUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSuperByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"orgUuid\",\"type\":\"bytes16\"}],\"name\":\"setOrgUuid\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addSuper\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getIdCartNoHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"delUser\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"easyCnsAddress\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"orgUuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"bytes32[2]\"},{\"indexed\":false,\"name\":\"idCartNoHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"onAddUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"onDelUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"onSetActive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"bytes32[2]\"}],\"name\":\"onSetUserAddressAndPublicKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"orgUuid\",\"type\":\"bytes16\"}],\"name\":\"onSetOrgUuid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"idCartNoHash\",\"type\":\"bytes32\"}],\"name\":\"onSetIdCartNoHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"onSetTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"onAddSuper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"onDelSuper\",\"type\":\"event\"}]"

// UsersDataBin is the compiled bytecode used for deploying new contracts.
const UsersDataBin = `606060405234156200000d57fe5b60405160208062002f0d833981016040528080519060200190919050505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b6200008d81620000e36401000000000262002bf9176401000000009004565b15156200009a5760006000fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5062000120565b6000600060008373ffffffffffffffffffffffffffffffffffffffff16141515156200010f5760006000fd5b823b90506000811191505b50919050565b612ddd80620001306000396000f3006060604052361561011b576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806324716b551461011d57806326c172b51461015b5780632a63069e146101b95780632e26787e146101f957806334349621146102765780633d1ef2b71461029c5780633f7a52b61461030f578063410f66e41461034b5780635249f18f146103b857806365f337f5146104365780636b3036821461046c5780637c333c0a1461049257806388fa6b71146104ec5780638eeded901461059c5780639bd027ce146105e7578063ab013c2614610657578063c13c7424146106b7578063e1e8637b14610706578063f0eaa2631461073c578063fc75c96a1461078b578063fe07283a146107be575bfe5b341561012557fe5b61015960048080356fffffffffffffffffffffffffffffffff19169060200190919080351515906020019091905050610805565b005b341561016357fe5b61017d60048080356000191690602001909190505061096c565b60405180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b34156101c157fe5b6101f760048080356fffffffffffffffffffffffffffffffff1916906020019091908035600019169060200190919050506109af565b005b341561020157fe5b61022a60048080356fffffffffffffffffffffffffffffffff1916906020019091905050610c57565b6040518082600260200280838360008314610264575b80518252602083111561026457602082019150602081019050602083039250610240565b50505090500191505060405180910390f35b341561027e57fe5b610286610d3c565b6040518082815260200191505060405180910390f35b34156102a457fe5b6102cd60048080356fffffffffffffffffffffffffffffffff1916906020019091905050610d4a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561031757fe5b61034960048080356fffffffffffffffffffffffffffffffff1916906020019091908035906020019091905050610e0a565b005b341561035357fe5b61037c60048080356fffffffffffffffffffffffffffffffff1916906020019091905050610fc0565b60405180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b34156103c057fe5b61043460048080356fffffffffffffffffffffffffffffffff191690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091908060400190600280602002604051908101604052809291908260026020028082843782019150505050509190505061107d565b005b341561043e57fe5b61046a600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506113ee565b005b341561047457fe5b61047c611600565b6040518082815260200191505060405180910390f35b341561049a57fe5b6104b0600480803590602001909190505061160e565b60405180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b34156104f457fe5b61059a60048080356fffffffffffffffffffffffffffffffff191690602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919080356fffffffffffffffffffffffffffffffff191690602001909190806040019060028060200260405190810160405280929190826002602002808284378201915050505050919080356000191690602001909190803590602001909190505061165b565b005b34156105a457fe5b6105cd60048080356fffffffffffffffffffffffffffffffff1916906020019091905050611cda565b604051808215151515815260200191505060405180910390f35b34156105ef57fe5b61061b600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611d3e565b60405180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b341561065f57fe5b6106756004808035906020019091905050611da5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156106bf57fe5b61070460048080356fffffffffffffffffffffffffffffffff19169060200190919080356fffffffffffffffffffffffffffffffff1916906020019091905050611e00565b005b341561070e57fe5b61073a600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612128565b005b341561074457fe5b61076d60048080356fffffffffffffffffffffffffffffffff19169060200190919050506122e6565b60405180826000191660001916815260200191505060405180910390f35b341561079357fe5b6107bc60048080356fffffffffffffffffffffffffffffffff1916906020019091905050612386565b005b34156107c657fe5b6107ef60048080356fffffffffffffffffffffffffffffffff191690602001909190505061266a565b6040518082815260200191505060405180910390f35b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061086657506108653361270a565b5b15156108725760006000fd5b600070010000000000000000000000000000000002826fffffffffffffffffffffffffffffffff1916141515156108a95760006000fd5b8060066000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160006101000a81548160ff0219169083151502179055507f713e572bb47b301d3993c6d748aa8f43912b3e2864d1fd701e16e586fbb53905828260405180836fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001821515151581526020019250505060405180910390a15b5b5050565b600060076000836000191660001916815260200190815260200160002060009054906101000a90047001000000000000000000000000000000000290505b919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610a105750610a0f3361270a565b5b1515610a1c5760006000fd5b8160066000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff161515610a735760006000fd5b600070010000000000000000000000000000000002836fffffffffffffffffffffffffffffffff191614158015610ab257506000600102826000191614155b1515610abe5760006000fd5b610ac782612784565b1515610ad35760006000fd5b6007600060066000866fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020600401546000191660001916815260200190815260200160002060006101000a8154906fffffffffffffffffffffffffffffffff02191690558160066000856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060040181600019169055508260076000846000191660001916815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff0219169083700100000000000000000000000000000000900402179055507f7811c8a53358a5d74aa42672dd20237aa4f645d7f260bd174b35851fc382ce55838360405180836fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200182600019166000191681526020019250505060405180910390a15b5b505b5050565b610c5f612c35565b8160066000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff161515610cb65760006000fd5b60066000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020600201600280602002604051908101604052809291908260028015610d2d576020028201915b81546000191681526020019060010190808311610d15575b505050505091505b5b50919050565b600060018054905090505b90565b60008160066000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff161515610da35760006000fd5b60066000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505b5b50919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610e6b5750610e6a3361270a565b5b1515610e775760006000fd5b8160066000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff161515610ece5760006000fd5b600070010000000000000000000000000000000002836fffffffffffffffffffffffffffffffff191614158015610f06575060008214155b1515610f125760006000fd5b8160066000856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020600501819055507f50843b091ea0bec4b6ead7f725b4fcb769671b6948c77b9cd98b25ff2440d84c838360405180836fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020018281526020019250505060405180910390a15b5b505b5050565b60008160066000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff1615156110195760006000fd5b60066000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060010160009054906101000a90047001000000000000000000000000000000000291505b5b50919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806110de57506110dd3361270a565b5b15156110ea5760006000fd5b8260066000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff1615156111415760006000fd5b600070010000000000000000000000000000000002846fffffffffffffffffffffffffffffffff19161415801561118f575060008373ffffffffffffffffffffffffffffffffffffffff1614155b80156111a0575061119f826127f0565b5b15156111ac5760006000fd5b6111b583612840565b15156111c15760006000fd5b6111cb83836128d0565b15156111d75760006000fd5b8260066000866fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160066000866fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060020190600261129d929190612c61565b5083600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff0219169083700100000000000000000000000000000000900402179055507fd64e6361ad680cd08aec52b6c8e50b6dda9a9c07b585c9ccc4375d984a488e0384848460405180846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001826002602002808383600083146113d2575b8051825260208311156113d2576020820191506020810190506020830392506113ae565b505050905001935050505060405180910390a15b5b505b505050565b60006000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561144f5760006000fd5b6114588361270a565b15156114645760006000fd5b60009150600090505b6001805490508110156115005760018181548110151561148957fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156114f25760019150611500565b5b808060010191505061146d565b81151561150957fe5b60018181548110151561151857fe5b906000526020600020900160005b6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690557f5561c1cda602083d14af2f0773152f273f6f9bf8feb17ce6265f94aec25bde0083604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a15b5b505050565b600060098054905090505b90565b600060098281548110151561161f57fe5b90600052602060002090600291828204019190066010025b9054906101000a90047001000000000000000000000000000000000290505b919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806116bc57506116bb3361270a565b5b15156116c85760006000fd5b8560066000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff161515156117205760006000fd5b600070010000000000000000000000000000000002876fffffffffffffffffffffffffffffffff19161415801561176e575060008673ffffffffffffffffffffffffffffffffffffffff1614155b801561177f575061177e846127f0565b5b801561179357506000600102836000191614155b80156117a0575060008214155b15156117ac5760006000fd5b600070010000000000000000000000000000000002856fffffffffffffffffffffffffffffffff19161415156118d1576117e4612912565b15156117f05760006000fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638eeded90866000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b15156118ac57fe5b6102c65a03f115156118ba57fe5b5050506040518051905015156118d05760006000fd5b5b6118da83612784565b15156118e65760006000fd5b6118ef86612840565b15156118fb5760006000fd5b61190586856128d0565b15156119115760006000fd5b60c0604051908101604052806001151581526020018773ffffffffffffffffffffffffffffffffffffffff168152602001866fffffffffffffffffffffffffffffffff19168152602001858152602001846000191681526020018381525060066000896fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160006101000a8154816fffffffffffffffffffffffffffffffff021916908370010000000000000000000000000000000090040217905550606082015181600201906002611a67929190612ca7565b506080820151816004019060001916905560a082015181600501559050508660076000856000191660001916815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff02191690837001000000000000000000000000000000009004021790555086600860008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff02191690837001000000000000000000000000000000009004021790555060098054806001018281611b659190612ced565b91600052602060002090600291828204019190066010025b89909190916101000a8154816fffffffffffffffffffffffffffffffff021916908370010000000000000000000000000000000090040217905550507ffc1ca23f6abe5b0a15b609b2202d84994ac33c4913c0a2644bb44f9dee400aa587878787878760405180876fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200184600260200280838360008314611ca4575b805182526020831115611ca457602082019150602081019050602083039250611c80565b5050509050018360001916600019168152602001828152602001965050505050505060405180910390a15b5b505b505050505050565b600060066000836fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff1615611d345760019050611d39565b600090505b919050565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a90047001000000000000000000000000000000000290505b919050565b6000611db082612b3b565b1515611dbc5760006000fd5b600182815481101515611dcb57fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480611e615750611e603361270a565b5b1515611e6d5760006000fd5b8160066000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff161515611ec45760006000fd5b600070010000000000000000000000000000000002836fffffffffffffffffffffffffffffffff191614151515611efb5760006000fd5b600070010000000000000000000000000000000002826fffffffffffffffffffffffffffffffff191614151561202057611f33612912565b1515611f3f5760006000fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638eeded90836000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b1515611ffb57fe5b6102c65a03f1151561200957fe5b50505060405180519050151561201f5760006000fd5b5b8160066000856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060010160006101000a8154816fffffffffffffffffffffffffffffffff0219169083700100000000000000000000000000000000900402179055507fc3f343d86c0f17b6b536a18799f6c0f89d37202680752b8e438d3f14d48291f2838360405180836fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019250505060405180910390a15b5b505b5050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156121855760006000fd5b60008173ffffffffffffffffffffffffffffffffffffffff16141515156121ac5760006000fd5b6121b58161270a565b1515156121c25760006000fd5b600180548060010182816121d69190612d27565b916000526020600020900160005b83909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f920999c03a7fd19172e4c657c22aa62a1bfb5e027a8cf6a2a0ee9daf4c5a347481604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a15b5b50565b60008160066000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff16151561233f5760006000fd5b60066000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020016000206004015491505b5b50919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806123e757506123e63361270a565b5b15156123f35760006000fd5b8060066000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff16151561244a5760006000fd5b600070010000000000000000000000000000000002826fffffffffffffffffffffffffffffffff1916141515156124815760006000fd5b6007600060066000856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020600401546000191660001916815260200190815260200160002060006101000a8154906fffffffffffffffffffffffffffffffff02191690556008600060066000856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a8154906fffffffffffffffffffffffffffffffff0219169055600060066000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160006101000a81548160ff0219169083151502179055507f7bc4ecd86fa11d1e4eb1e479a0fc84e9e48b8b799ff51ffa6375f7fac922cf158260405180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200191505060405180910390a15b5b505b50565b60008160066000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff1615156126c35760006000fd5b60066000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020016000206005015491505b5b50919050565b600060008273ffffffffffffffffffffffffffffffffffffffff161415801561277c5750600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b90505b919050565b600060007001000000000000000000000000000000000260076000846000191660001916815260200190815260200160002060009054906101000a9004700100000000000000000000000000000000026fffffffffffffffffffffffffffffffff19161490505b919050565b6000600060010282600060028110151561280657fe5b602002015160001916141580156128385750600060010282600160028110151561282c57fe5b60200201516000191614155b90505b919050565b6000600070010000000000000000000000000000000002600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a9004700100000000000000000000000000000000026fffffffffffffffffffffffffffffffff19161490505b919050565b60008273ffffffffffffffffffffffffffffffffffffffff166128f283612b5a565b73ffffffffffffffffffffffffffffffffffffffff161490505b92915050565b60006000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e61295c612bb4565b6000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252838181518152602001915080519060200190808383600083146129df575b8051825260208311156129df576020820191506020810190506020830392506129bb565b505050905090810190601f168015612a0b5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1515612a2657fe5b6102c65a03f11515612a3457fe5b505050604051805190509050612a4981612bf9565b1515612a585760009150612b37565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515612b325780600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600191505b5090565b600060008210158015612b52575060018054905082105b90505b919050565b6000816040518082600260200280838360008314612b97575b805182526020831115612b9757602082019150602081019050602083039250612b73565b50505090500191505060405180910390206001900490505b919050565b612bbc612d53565b604060405190810160405280600881526020017f4f7267734461746100000000000000000000000000000000000000000000000081525090505b90565b6000600060008373ffffffffffffffffffffffffffffffffffffffff1614151515612c245760006000fd5b823b90506000811191505b50919050565b6040604051908101604052806002905b600060001916815260200190600190039081612c455790505090565b8260028101928215612c96579160200282015b82811115612c95578251829060001916905591602001919060010190612c74565b5b509050612ca39190612d67565b5090565b8260028101928215612cdc579160200282015b82811115612cdb578251829060001916905591602001919060010190612cba565b5b509050612ce99190612d67565b5090565b815481835581811511612d22576001016002900481600101600290048360005260206000209182019101612d219190612d8c565b5b505050565b815481835581811511612d4e57818360005260206000209182019101612d4d9190612d8c565b5b505050565b602060405190810160405280600081525090565b612d8991905b80821115612d85576000816000905550600101612d6d565b5090565b90565b612dae91905b80821115612daa576000816000905550600101612d92565b5090565b905600a165627a7a7230582029017622e77da28f06c01dc67be824623a24b6dd89c1345634f32d4e73d1a8e90029`

// DeployUsersData deploys a new Ethereum contract, binding an instance of UsersData to it.
func DeployUsersData(auth *bind.TransactOpts, backend bind.ContractBackend, easyCnsAddress common.Address) (common.Address, *types.Transaction, *UsersData, error) {
	parsed, err := abi.JSON(strings.NewReader(UsersDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsersDataBin), backend, easyCnsAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UsersData{UsersDataCaller: UsersDataCaller{contract: contract}, UsersDataTransactor: UsersDataTransactor{contract: contract}, UsersDataFilterer: UsersDataFilterer{contract: contract}}, nil
}

// UsersData is an auto generated Go binding around an Ethereum contract.
type UsersData struct {
	UsersDataCaller     // Read-only binding to the contract
	UsersDataTransactor // Write-only binding to the contract
	UsersDataFilterer   // Log filterer for contract events
}

// UsersDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsersDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsersDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsersDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsersDataSession struct {
	Contract     *UsersData        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsersDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsersDataCallerSession struct {
	Contract *UsersDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UsersDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsersDataTransactorSession struct {
	Contract     *UsersDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UsersDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsersDataRaw struct {
	Contract *UsersData // Generic contract binding to access the raw methods on
}

// UsersDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsersDataCallerRaw struct {
	Contract *UsersDataCaller // Generic read-only contract binding to access the raw methods on
}

// UsersDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsersDataTransactorRaw struct {
	Contract *UsersDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsersData creates a new instance of UsersData, bound to a specific deployed contract.
func NewUsersData(address common.Address, backend bind.ContractBackend) (*UsersData, error) {
	contract, err := bindUsersData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsersData{UsersDataCaller: UsersDataCaller{contract: contract}, UsersDataTransactor: UsersDataTransactor{contract: contract}, UsersDataFilterer: UsersDataFilterer{contract: contract}}, nil
}

// NewUsersDataCaller creates a new read-only instance of UsersData, bound to a specific deployed contract.
func NewUsersDataCaller(address common.Address, caller bind.ContractCaller) (*UsersDataCaller, error) {
	contract, err := bindUsersData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsersDataCaller{contract: contract}, nil
}

// NewUsersDataTransactor creates a new write-only instance of UsersData, bound to a specific deployed contract.
func NewUsersDataTransactor(address common.Address, transactor bind.ContractTransactor) (*UsersDataTransactor, error) {
	contract, err := bindUsersData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsersDataTransactor{contract: contract}, nil
}

// NewUsersDataFilterer creates a new log filterer instance of UsersData, bound to a specific deployed contract.
func NewUsersDataFilterer(address common.Address, filterer bind.ContractFilterer) (*UsersDataFilterer, error) {
	contract, err := bindUsersData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsersDataFilterer{contract: contract}, nil
}

// bindUsersData binds a generic wrapper to an already deployed contract.
func bindUsersData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsersDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsersData *UsersDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsersData.Contract.UsersDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsersData *UsersDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsersData.Contract.UsersDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsersData *UsersDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsersData.Contract.UsersDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsersData *UsersDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsersData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsersData *UsersDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsersData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsersData *UsersDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsersData.Contract.contract.Transact(opts, method, params...)
}

// GetIdCartNoHash is a free data retrieval call binding the contract method 0xf0eaa263.
//
// Solidity: function getIdCartNoHash(uuid bytes16) constant returns(bytes32)
func (_UsersData *UsersDataCaller) GetIdCartNoHash(opts *bind.CallOpts, uuid [16]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _UsersData.contract.Call(opts, out, "getIdCartNoHash", uuid)
	return *ret0, err
}

// GetIdCartNoHash is a free data retrieval call binding the contract method 0xf0eaa263.
//
// Solidity: function getIdCartNoHash(uuid bytes16) constant returns(bytes32)
func (_UsersData *UsersDataSession) GetIdCartNoHash(uuid [16]byte) ([32]byte, error) {
	return _UsersData.Contract.GetIdCartNoHash(&_UsersData.CallOpts, uuid)
}

// GetIdCartNoHash is a free data retrieval call binding the contract method 0xf0eaa263.
//
// Solidity: function getIdCartNoHash(uuid bytes16) constant returns(bytes32)
func (_UsersData *UsersDataCallerSession) GetIdCartNoHash(uuid [16]byte) ([32]byte, error) {
	return _UsersData.Contract.GetIdCartNoHash(&_UsersData.CallOpts, uuid)
}

// GetOrgUuid is a free data retrieval call binding the contract method 0x410f66e4.
//
// Solidity: function getOrgUuid(uuid bytes16) constant returns(bytes16)
func (_UsersData *UsersDataCaller) GetOrgUuid(opts *bind.CallOpts, uuid [16]byte) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _UsersData.contract.Call(opts, out, "getOrgUuid", uuid)
	return *ret0, err
}

// GetOrgUuid is a free data retrieval call binding the contract method 0x410f66e4.
//
// Solidity: function getOrgUuid(uuid bytes16) constant returns(bytes16)
func (_UsersData *UsersDataSession) GetOrgUuid(uuid [16]byte) ([16]byte, error) {
	return _UsersData.Contract.GetOrgUuid(&_UsersData.CallOpts, uuid)
}

// GetOrgUuid is a free data retrieval call binding the contract method 0x410f66e4.
//
// Solidity: function getOrgUuid(uuid bytes16) constant returns(bytes16)
func (_UsersData *UsersDataCallerSession) GetOrgUuid(uuid [16]byte) ([16]byte, error) {
	return _UsersData.Contract.GetOrgUuid(&_UsersData.CallOpts, uuid)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x2e26787e.
//
// Solidity: function getPublicKey(uuid bytes16) constant returns(bytes32[2])
func (_UsersData *UsersDataCaller) GetPublicKey(opts *bind.CallOpts, uuid [16]byte) ([2][32]byte, error) {
	var (
		ret0 = new([2][32]byte)
	)
	out := ret0
	err := _UsersData.contract.Call(opts, out, "getPublicKey", uuid)
	return *ret0, err
}

// GetPublicKey is a free data retrieval call binding the contract method 0x2e26787e.
//
// Solidity: function getPublicKey(uuid bytes16) constant returns(bytes32[2])
func (_UsersData *UsersDataSession) GetPublicKey(uuid [16]byte) ([2][32]byte, error) {
	return _UsersData.Contract.GetPublicKey(&_UsersData.CallOpts, uuid)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x2e26787e.
//
// Solidity: function getPublicKey(uuid bytes16) constant returns(bytes32[2])
func (_UsersData *UsersDataCallerSession) GetPublicKey(uuid [16]byte) ([2][32]byte, error) {
	return _UsersData.Contract.GetPublicKey(&_UsersData.CallOpts, uuid)
}

// GetSuperByIndex is a free data retrieval call binding the contract method 0xab013c26.
//
// Solidity: function getSuperByIndex(index uint256) constant returns(address)
func (_UsersData *UsersDataCaller) GetSuperByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UsersData.contract.Call(opts, out, "getSuperByIndex", index)
	return *ret0, err
}

// GetSuperByIndex is a free data retrieval call binding the contract method 0xab013c26.
//
// Solidity: function getSuperByIndex(index uint256) constant returns(address)
func (_UsersData *UsersDataSession) GetSuperByIndex(index *big.Int) (common.Address, error) {
	return _UsersData.Contract.GetSuperByIndex(&_UsersData.CallOpts, index)
}

// GetSuperByIndex is a free data retrieval call binding the contract method 0xab013c26.
//
// Solidity: function getSuperByIndex(index uint256) constant returns(address)
func (_UsersData *UsersDataCallerSession) GetSuperByIndex(index *big.Int) (common.Address, error) {
	return _UsersData.Contract.GetSuperByIndex(&_UsersData.CallOpts, index)
}

// GetSuperSize is a free data retrieval call binding the contract method 0x34349621.
//
// Solidity: function getSuperSize() constant returns(uint256)
func (_UsersData *UsersDataCaller) GetSuperSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UsersData.contract.Call(opts, out, "getSuperSize")
	return *ret0, err
}

// GetSuperSize is a free data retrieval call binding the contract method 0x34349621.
//
// Solidity: function getSuperSize() constant returns(uint256)
func (_UsersData *UsersDataSession) GetSuperSize() (*big.Int, error) {
	return _UsersData.Contract.GetSuperSize(&_UsersData.CallOpts)
}

// GetSuperSize is a free data retrieval call binding the contract method 0x34349621.
//
// Solidity: function getSuperSize() constant returns(uint256)
func (_UsersData *UsersDataCallerSession) GetSuperSize() (*big.Int, error) {
	return _UsersData.Contract.GetSuperSize(&_UsersData.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0xfe07283a.
//
// Solidity: function getTime(uuid bytes16) constant returns(uint256)
func (_UsersData *UsersDataCaller) GetTime(opts *bind.CallOpts, uuid [16]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UsersData.contract.Call(opts, out, "getTime", uuid)
	return *ret0, err
}

// GetTime is a free data retrieval call binding the contract method 0xfe07283a.
//
// Solidity: function getTime(uuid bytes16) constant returns(uint256)
func (_UsersData *UsersDataSession) GetTime(uuid [16]byte) (*big.Int, error) {
	return _UsersData.Contract.GetTime(&_UsersData.CallOpts, uuid)
}

// GetTime is a free data retrieval call binding the contract method 0xfe07283a.
//
// Solidity: function getTime(uuid bytes16) constant returns(uint256)
func (_UsersData *UsersDataCallerSession) GetTime(uuid [16]byte) (*big.Int, error) {
	return _UsersData.Contract.GetTime(&_UsersData.CallOpts, uuid)
}

// GetUserAddress is a free data retrieval call binding the contract method 0x3d1ef2b7.
//
// Solidity: function getUserAddress(uuid bytes16) constant returns(address)
func (_UsersData *UsersDataCaller) GetUserAddress(opts *bind.CallOpts, uuid [16]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UsersData.contract.Call(opts, out, "getUserAddress", uuid)
	return *ret0, err
}

// GetUserAddress is a free data retrieval call binding the contract method 0x3d1ef2b7.
//
// Solidity: function getUserAddress(uuid bytes16) constant returns(address)
func (_UsersData *UsersDataSession) GetUserAddress(uuid [16]byte) (common.Address, error) {
	return _UsersData.Contract.GetUserAddress(&_UsersData.CallOpts, uuid)
}

// GetUserAddress is a free data retrieval call binding the contract method 0x3d1ef2b7.
//
// Solidity: function getUserAddress(uuid bytes16) constant returns(address)
func (_UsersData *UsersDataCallerSession) GetUserAddress(uuid [16]byte) (common.Address, error) {
	return _UsersData.Contract.GetUserAddress(&_UsersData.CallOpts, uuid)
}

// GetUuidByIdCartNoHash is a free data retrieval call binding the contract method 0x26c172b5.
//
// Solidity: function getUuidByIdCartNoHash(idCartNoHash bytes32) constant returns(bytes16)
func (_UsersData *UsersDataCaller) GetUuidByIdCartNoHash(opts *bind.CallOpts, idCartNoHash [32]byte) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _UsersData.contract.Call(opts, out, "getUuidByIdCartNoHash", idCartNoHash)
	return *ret0, err
}

// GetUuidByIdCartNoHash is a free data retrieval call binding the contract method 0x26c172b5.
//
// Solidity: function getUuidByIdCartNoHash(idCartNoHash bytes32) constant returns(bytes16)
func (_UsersData *UsersDataSession) GetUuidByIdCartNoHash(idCartNoHash [32]byte) ([16]byte, error) {
	return _UsersData.Contract.GetUuidByIdCartNoHash(&_UsersData.CallOpts, idCartNoHash)
}

// GetUuidByIdCartNoHash is a free data retrieval call binding the contract method 0x26c172b5.
//
// Solidity: function getUuidByIdCartNoHash(idCartNoHash bytes32) constant returns(bytes16)
func (_UsersData *UsersDataCallerSession) GetUuidByIdCartNoHash(idCartNoHash [32]byte) ([16]byte, error) {
	return _UsersData.Contract.GetUuidByIdCartNoHash(&_UsersData.CallOpts, idCartNoHash)
}

// GetUuidByIndex is a free data retrieval call binding the contract method 0x7c333c0a.
//
// Solidity: function getUuidByIndex(index uint256) constant returns(bytes16)
func (_UsersData *UsersDataCaller) GetUuidByIndex(opts *bind.CallOpts, index *big.Int) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _UsersData.contract.Call(opts, out, "getUuidByIndex", index)
	return *ret0, err
}

// GetUuidByIndex is a free data retrieval call binding the contract method 0x7c333c0a.
//
// Solidity: function getUuidByIndex(index uint256) constant returns(bytes16)
func (_UsersData *UsersDataSession) GetUuidByIndex(index *big.Int) ([16]byte, error) {
	return _UsersData.Contract.GetUuidByIndex(&_UsersData.CallOpts, index)
}

// GetUuidByIndex is a free data retrieval call binding the contract method 0x7c333c0a.
//
// Solidity: function getUuidByIndex(index uint256) constant returns(bytes16)
func (_UsersData *UsersDataCallerSession) GetUuidByIndex(index *big.Int) ([16]byte, error) {
	return _UsersData.Contract.GetUuidByIndex(&_UsersData.CallOpts, index)
}

// GetUuidByUserAddress is a free data retrieval call binding the contract method 0x9bd027ce.
//
// Solidity: function getUuidByUserAddress(userAddress address) constant returns(bytes16)
func (_UsersData *UsersDataCaller) GetUuidByUserAddress(opts *bind.CallOpts, userAddress common.Address) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _UsersData.contract.Call(opts, out, "getUuidByUserAddress", userAddress)
	return *ret0, err
}

// GetUuidByUserAddress is a free data retrieval call binding the contract method 0x9bd027ce.
//
// Solidity: function getUuidByUserAddress(userAddress address) constant returns(bytes16)
func (_UsersData *UsersDataSession) GetUuidByUserAddress(userAddress common.Address) ([16]byte, error) {
	return _UsersData.Contract.GetUuidByUserAddress(&_UsersData.CallOpts, userAddress)
}

// GetUuidByUserAddress is a free data retrieval call binding the contract method 0x9bd027ce.
//
// Solidity: function getUuidByUserAddress(userAddress address) constant returns(bytes16)
func (_UsersData *UsersDataCallerSession) GetUuidByUserAddress(userAddress common.Address) ([16]byte, error) {
	return _UsersData.Contract.GetUuidByUserAddress(&_UsersData.CallOpts, userAddress)
}

// GetUuidListSize is a free data retrieval call binding the contract method 0x6b303682.
//
// Solidity: function getUuidListSize() constant returns(uint256)
func (_UsersData *UsersDataCaller) GetUuidListSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UsersData.contract.Call(opts, out, "getUuidListSize")
	return *ret0, err
}

// GetUuidListSize is a free data retrieval call binding the contract method 0x6b303682.
//
// Solidity: function getUuidListSize() constant returns(uint256)
func (_UsersData *UsersDataSession) GetUuidListSize() (*big.Int, error) {
	return _UsersData.Contract.GetUuidListSize(&_UsersData.CallOpts)
}

// GetUuidListSize is a free data retrieval call binding the contract method 0x6b303682.
//
// Solidity: function getUuidListSize() constant returns(uint256)
func (_UsersData *UsersDataCallerSession) GetUuidListSize() (*big.Int, error) {
	return _UsersData.Contract.GetUuidListSize(&_UsersData.CallOpts)
}

// IsUuidExist is a free data retrieval call binding the contract method 0x8eeded90.
//
// Solidity: function isUuidExist(uuid bytes16) constant returns(bool)
func (_UsersData *UsersDataCaller) IsUuidExist(opts *bind.CallOpts, uuid [16]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UsersData.contract.Call(opts, out, "isUuidExist", uuid)
	return *ret0, err
}

// IsUuidExist is a free data retrieval call binding the contract method 0x8eeded90.
//
// Solidity: function isUuidExist(uuid bytes16) constant returns(bool)
func (_UsersData *UsersDataSession) IsUuidExist(uuid [16]byte) (bool, error) {
	return _UsersData.Contract.IsUuidExist(&_UsersData.CallOpts, uuid)
}

// IsUuidExist is a free data retrieval call binding the contract method 0x8eeded90.
//
// Solidity: function isUuidExist(uuid bytes16) constant returns(bool)
func (_UsersData *UsersDataCallerSession) IsUuidExist(uuid [16]byte) (bool, error) {
	return _UsersData.Contract.IsUuidExist(&_UsersData.CallOpts, uuid)
}

// AddSuper is a paid mutator transaction binding the contract method 0xe1e8637b.
//
// Solidity: function addSuper(addr address) returns()
func (_UsersData *UsersDataTransactor) AddSuper(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _UsersData.contract.Transact(opts, "addSuper", addr)
}

// AddSuper is a paid mutator transaction binding the contract method 0xe1e8637b.
//
// Solidity: function addSuper(addr address) returns()
func (_UsersData *UsersDataSession) AddSuper(addr common.Address) (*types.Transaction, error) {
	return _UsersData.Contract.AddSuper(&_UsersData.TransactOpts, addr)
}

// AddSuper is a paid mutator transaction binding the contract method 0xe1e8637b.
//
// Solidity: function addSuper(addr address) returns()
func (_UsersData *UsersDataTransactorSession) AddSuper(addr common.Address) (*types.Transaction, error) {
	return _UsersData.Contract.AddSuper(&_UsersData.TransactOpts, addr)
}

// AddUser is a paid mutator transaction binding the contract method 0x88fa6b71.
//
// Solidity: function addUser(uuid bytes16, userAddress address, orgUuid bytes16, publicKey bytes32[2], idCartNoHash bytes32, time uint256) returns()
func (_UsersData *UsersDataTransactor) AddUser(opts *bind.TransactOpts, uuid [16]byte, userAddress common.Address, orgUuid [16]byte, publicKey [2][32]byte, idCartNoHash [32]byte, time *big.Int) (*types.Transaction, error) {
	return _UsersData.contract.Transact(opts, "addUser", uuid, userAddress, orgUuid, publicKey, idCartNoHash, time)
}

// AddUser is a paid mutator transaction binding the contract method 0x88fa6b71.
//
// Solidity: function addUser(uuid bytes16, userAddress address, orgUuid bytes16, publicKey bytes32[2], idCartNoHash bytes32, time uint256) returns()
func (_UsersData *UsersDataSession) AddUser(uuid [16]byte, userAddress common.Address, orgUuid [16]byte, publicKey [2][32]byte, idCartNoHash [32]byte, time *big.Int) (*types.Transaction, error) {
	return _UsersData.Contract.AddUser(&_UsersData.TransactOpts, uuid, userAddress, orgUuid, publicKey, idCartNoHash, time)
}

// AddUser is a paid mutator transaction binding the contract method 0x88fa6b71.
//
// Solidity: function addUser(uuid bytes16, userAddress address, orgUuid bytes16, publicKey bytes32[2], idCartNoHash bytes32, time uint256) returns()
func (_UsersData *UsersDataTransactorSession) AddUser(uuid [16]byte, userAddress common.Address, orgUuid [16]byte, publicKey [2][32]byte, idCartNoHash [32]byte, time *big.Int) (*types.Transaction, error) {
	return _UsersData.Contract.AddUser(&_UsersData.TransactOpts, uuid, userAddress, orgUuid, publicKey, idCartNoHash, time)
}

// DelSuper is a paid mutator transaction binding the contract method 0x65f337f5.
//
// Solidity: function delSuper(addr address) returns()
func (_UsersData *UsersDataTransactor) DelSuper(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _UsersData.contract.Transact(opts, "delSuper", addr)
}

// DelSuper is a paid mutator transaction binding the contract method 0x65f337f5.
//
// Solidity: function delSuper(addr address) returns()
func (_UsersData *UsersDataSession) DelSuper(addr common.Address) (*types.Transaction, error) {
	return _UsersData.Contract.DelSuper(&_UsersData.TransactOpts, addr)
}

// DelSuper is a paid mutator transaction binding the contract method 0x65f337f5.
//
// Solidity: function delSuper(addr address) returns()
func (_UsersData *UsersDataTransactorSession) DelSuper(addr common.Address) (*types.Transaction, error) {
	return _UsersData.Contract.DelSuper(&_UsersData.TransactOpts, addr)
}

// DelUser is a paid mutator transaction binding the contract method 0xfc75c96a.
//
// Solidity: function delUser(uuid bytes16) returns()
func (_UsersData *UsersDataTransactor) DelUser(opts *bind.TransactOpts, uuid [16]byte) (*types.Transaction, error) {
	return _UsersData.contract.Transact(opts, "delUser", uuid)
}

// DelUser is a paid mutator transaction binding the contract method 0xfc75c96a.
//
// Solidity: function delUser(uuid bytes16) returns()
func (_UsersData *UsersDataSession) DelUser(uuid [16]byte) (*types.Transaction, error) {
	return _UsersData.Contract.DelUser(&_UsersData.TransactOpts, uuid)
}

// DelUser is a paid mutator transaction binding the contract method 0xfc75c96a.
//
// Solidity: function delUser(uuid bytes16) returns()
func (_UsersData *UsersDataTransactorSession) DelUser(uuid [16]byte) (*types.Transaction, error) {
	return _UsersData.Contract.DelUser(&_UsersData.TransactOpts, uuid)
}

// SetActive is a paid mutator transaction binding the contract method 0x24716b55.
//
// Solidity: function setActive(uuid bytes16, active bool) returns()
func (_UsersData *UsersDataTransactor) SetActive(opts *bind.TransactOpts, uuid [16]byte, active bool) (*types.Transaction, error) {
	return _UsersData.contract.Transact(opts, "setActive", uuid, active)
}

// SetActive is a paid mutator transaction binding the contract method 0x24716b55.
//
// Solidity: function setActive(uuid bytes16, active bool) returns()
func (_UsersData *UsersDataSession) SetActive(uuid [16]byte, active bool) (*types.Transaction, error) {
	return _UsersData.Contract.SetActive(&_UsersData.TransactOpts, uuid, active)
}

// SetActive is a paid mutator transaction binding the contract method 0x24716b55.
//
// Solidity: function setActive(uuid bytes16, active bool) returns()
func (_UsersData *UsersDataTransactorSession) SetActive(uuid [16]byte, active bool) (*types.Transaction, error) {
	return _UsersData.Contract.SetActive(&_UsersData.TransactOpts, uuid, active)
}

// SetIdCartNoHash is a paid mutator transaction binding the contract method 0x2a63069e.
//
// Solidity: function setIdCartNoHash(uuid bytes16, idCartNoHash bytes32) returns()
func (_UsersData *UsersDataTransactor) SetIdCartNoHash(opts *bind.TransactOpts, uuid [16]byte, idCartNoHash [32]byte) (*types.Transaction, error) {
	return _UsersData.contract.Transact(opts, "setIdCartNoHash", uuid, idCartNoHash)
}

// SetIdCartNoHash is a paid mutator transaction binding the contract method 0x2a63069e.
//
// Solidity: function setIdCartNoHash(uuid bytes16, idCartNoHash bytes32) returns()
func (_UsersData *UsersDataSession) SetIdCartNoHash(uuid [16]byte, idCartNoHash [32]byte) (*types.Transaction, error) {
	return _UsersData.Contract.SetIdCartNoHash(&_UsersData.TransactOpts, uuid, idCartNoHash)
}

// SetIdCartNoHash is a paid mutator transaction binding the contract method 0x2a63069e.
//
// Solidity: function setIdCartNoHash(uuid bytes16, idCartNoHash bytes32) returns()
func (_UsersData *UsersDataTransactorSession) SetIdCartNoHash(uuid [16]byte, idCartNoHash [32]byte) (*types.Transaction, error) {
	return _UsersData.Contract.SetIdCartNoHash(&_UsersData.TransactOpts, uuid, idCartNoHash)
}

// SetOrgUuid is a paid mutator transaction binding the contract method 0xc13c7424.
//
// Solidity: function setOrgUuid(uuid bytes16, orgUuid bytes16) returns()
func (_UsersData *UsersDataTransactor) SetOrgUuid(opts *bind.TransactOpts, uuid [16]byte, orgUuid [16]byte) (*types.Transaction, error) {
	return _UsersData.contract.Transact(opts, "setOrgUuid", uuid, orgUuid)
}

// SetOrgUuid is a paid mutator transaction binding the contract method 0xc13c7424.
//
// Solidity: function setOrgUuid(uuid bytes16, orgUuid bytes16) returns()
func (_UsersData *UsersDataSession) SetOrgUuid(uuid [16]byte, orgUuid [16]byte) (*types.Transaction, error) {
	return _UsersData.Contract.SetOrgUuid(&_UsersData.TransactOpts, uuid, orgUuid)
}

// SetOrgUuid is a paid mutator transaction binding the contract method 0xc13c7424.
//
// Solidity: function setOrgUuid(uuid bytes16, orgUuid bytes16) returns()
func (_UsersData *UsersDataTransactorSession) SetOrgUuid(uuid [16]byte, orgUuid [16]byte) (*types.Transaction, error) {
	return _UsersData.Contract.SetOrgUuid(&_UsersData.TransactOpts, uuid, orgUuid)
}

// SetTime is a paid mutator transaction binding the contract method 0x3f7a52b6.
//
// Solidity: function setTime(uuid bytes16, time uint256) returns()
func (_UsersData *UsersDataTransactor) SetTime(opts *bind.TransactOpts, uuid [16]byte, time *big.Int) (*types.Transaction, error) {
	return _UsersData.contract.Transact(opts, "setTime", uuid, time)
}

// SetTime is a paid mutator transaction binding the contract method 0x3f7a52b6.
//
// Solidity: function setTime(uuid bytes16, time uint256) returns()
func (_UsersData *UsersDataSession) SetTime(uuid [16]byte, time *big.Int) (*types.Transaction, error) {
	return _UsersData.Contract.SetTime(&_UsersData.TransactOpts, uuid, time)
}

// SetTime is a paid mutator transaction binding the contract method 0x3f7a52b6.
//
// Solidity: function setTime(uuid bytes16, time uint256) returns()
func (_UsersData *UsersDataTransactorSession) SetTime(uuid [16]byte, time *big.Int) (*types.Transaction, error) {
	return _UsersData.Contract.SetTime(&_UsersData.TransactOpts, uuid, time)
}

// SetUserAddressAndPublicKey is a paid mutator transaction binding the contract method 0x5249f18f.
//
// Solidity: function setUserAddressAndPublicKey(uuid bytes16, userAddress address, publicKey bytes32[2]) returns()
func (_UsersData *UsersDataTransactor) SetUserAddressAndPublicKey(opts *bind.TransactOpts, uuid [16]byte, userAddress common.Address, publicKey [2][32]byte) (*types.Transaction, error) {
	return _UsersData.contract.Transact(opts, "setUserAddressAndPublicKey", uuid, userAddress, publicKey)
}

// SetUserAddressAndPublicKey is a paid mutator transaction binding the contract method 0x5249f18f.
//
// Solidity: function setUserAddressAndPublicKey(uuid bytes16, userAddress address, publicKey bytes32[2]) returns()
func (_UsersData *UsersDataSession) SetUserAddressAndPublicKey(uuid [16]byte, userAddress common.Address, publicKey [2][32]byte) (*types.Transaction, error) {
	return _UsersData.Contract.SetUserAddressAndPublicKey(&_UsersData.TransactOpts, uuid, userAddress, publicKey)
}

// SetUserAddressAndPublicKey is a paid mutator transaction binding the contract method 0x5249f18f.
//
// Solidity: function setUserAddressAndPublicKey(uuid bytes16, userAddress address, publicKey bytes32[2]) returns()
func (_UsersData *UsersDataTransactorSession) SetUserAddressAndPublicKey(uuid [16]byte, userAddress common.Address, publicKey [2][32]byte) (*types.Transaction, error) {
	return _UsersData.Contract.SetUserAddressAndPublicKey(&_UsersData.TransactOpts, uuid, userAddress, publicKey)
}

// UsersDataOnAddSuperIterator is returned from FilterOnAddSuper and is used to iterate over the raw logs and unpacked data for OnAddSuper events raised by the UsersData contract.
type UsersDataOnAddSuperIterator struct {
	Event *UsersDataOnAddSuper // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsersDataOnAddSuperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersDataOnAddSuper)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsersDataOnAddSuper)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersDataOnAddSuperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersDataOnAddSuperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersDataOnAddSuper represents a OnAddSuper event raised by the UsersData contract.
type UsersDataOnAddSuper struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnAddSuper is a free log retrieval operation binding the contract event 0x920999c03a7fd19172e4c657c22aa62a1bfb5e027a8cf6a2a0ee9daf4c5a3474.
//
// Solidity: e onAddSuper(addr address)
func (_UsersData *UsersDataFilterer) FilterOnAddSuper(opts *bind.FilterOpts) (*UsersDataOnAddSuperIterator, error) {

	logs, sub, err := _UsersData.contract.FilterLogs(opts, "onAddSuper")
	if err != nil {
		return nil, err
	}
	return &UsersDataOnAddSuperIterator{contract: _UsersData.contract, event: "onAddSuper", logs: logs, sub: sub}, nil
}

// WatchOnAddSuper is a free log subscription operation binding the contract event 0x920999c03a7fd19172e4c657c22aa62a1bfb5e027a8cf6a2a0ee9daf4c5a3474.
//
// Solidity: e onAddSuper(addr address)
func (_UsersData *UsersDataFilterer) WatchOnAddSuper(opts *bind.WatchOpts, sink chan<- *UsersDataOnAddSuper) (event.Subscription, error) {

	logs, sub, err := _UsersData.contract.WatchLogs(opts, "onAddSuper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersDataOnAddSuper)
				if err := _UsersData.contract.UnpackLog(event, "onAddSuper", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UsersDataOnAddUserIterator is returned from FilterOnAddUser and is used to iterate over the raw logs and unpacked data for OnAddUser events raised by the UsersData contract.
type UsersDataOnAddUserIterator struct {
	Event *UsersDataOnAddUser // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsersDataOnAddUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersDataOnAddUser)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsersDataOnAddUser)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersDataOnAddUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersDataOnAddUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersDataOnAddUser represents a OnAddUser event raised by the UsersData contract.
type UsersDataOnAddUser struct {
	Uuid         [16]byte
	UserAddress  common.Address
	OrgUuid      [16]byte
	PublicKey    [2][32]byte
	IdCartNoHash [32]byte
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOnAddUser is a free log retrieval operation binding the contract event 0xfc1ca23f6abe5b0a15b609b2202d84994ac33c4913c0a2644bb44f9dee400aa5.
//
// Solidity: e onAddUser(uuid bytes16, userAddress address, orgUuid bytes16, publicKey bytes32[2], idCartNoHash bytes32, time uint256)
func (_UsersData *UsersDataFilterer) FilterOnAddUser(opts *bind.FilterOpts) (*UsersDataOnAddUserIterator, error) {

	logs, sub, err := _UsersData.contract.FilterLogs(opts, "onAddUser")
	if err != nil {
		return nil, err
	}
	return &UsersDataOnAddUserIterator{contract: _UsersData.contract, event: "onAddUser", logs: logs, sub: sub}, nil
}

// WatchOnAddUser is a free log subscription operation binding the contract event 0xfc1ca23f6abe5b0a15b609b2202d84994ac33c4913c0a2644bb44f9dee400aa5.
//
// Solidity: e onAddUser(uuid bytes16, userAddress address, orgUuid bytes16, publicKey bytes32[2], idCartNoHash bytes32, time uint256)
func (_UsersData *UsersDataFilterer) WatchOnAddUser(opts *bind.WatchOpts, sink chan<- *UsersDataOnAddUser) (event.Subscription, error) {

	logs, sub, err := _UsersData.contract.WatchLogs(opts, "onAddUser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersDataOnAddUser)
				if err := _UsersData.contract.UnpackLog(event, "onAddUser", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UsersDataOnDelSuperIterator is returned from FilterOnDelSuper and is used to iterate over the raw logs and unpacked data for OnDelSuper events raised by the UsersData contract.
type UsersDataOnDelSuperIterator struct {
	Event *UsersDataOnDelSuper // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsersDataOnDelSuperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersDataOnDelSuper)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsersDataOnDelSuper)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersDataOnDelSuperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersDataOnDelSuperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersDataOnDelSuper represents a OnDelSuper event raised by the UsersData contract.
type UsersDataOnDelSuper struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnDelSuper is a free log retrieval operation binding the contract event 0x5561c1cda602083d14af2f0773152f273f6f9bf8feb17ce6265f94aec25bde00.
//
// Solidity: e onDelSuper(addr address)
func (_UsersData *UsersDataFilterer) FilterOnDelSuper(opts *bind.FilterOpts) (*UsersDataOnDelSuperIterator, error) {

	logs, sub, err := _UsersData.contract.FilterLogs(opts, "onDelSuper")
	if err != nil {
		return nil, err
	}
	return &UsersDataOnDelSuperIterator{contract: _UsersData.contract, event: "onDelSuper", logs: logs, sub: sub}, nil
}

// WatchOnDelSuper is a free log subscription operation binding the contract event 0x5561c1cda602083d14af2f0773152f273f6f9bf8feb17ce6265f94aec25bde00.
//
// Solidity: e onDelSuper(addr address)
func (_UsersData *UsersDataFilterer) WatchOnDelSuper(opts *bind.WatchOpts, sink chan<- *UsersDataOnDelSuper) (event.Subscription, error) {

	logs, sub, err := _UsersData.contract.WatchLogs(opts, "onDelSuper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersDataOnDelSuper)
				if err := _UsersData.contract.UnpackLog(event, "onDelSuper", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UsersDataOnDelUserIterator is returned from FilterOnDelUser and is used to iterate over the raw logs and unpacked data for OnDelUser events raised by the UsersData contract.
type UsersDataOnDelUserIterator struct {
	Event *UsersDataOnDelUser // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsersDataOnDelUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersDataOnDelUser)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsersDataOnDelUser)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersDataOnDelUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersDataOnDelUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersDataOnDelUser represents a OnDelUser event raised by the UsersData contract.
type UsersDataOnDelUser struct {
	Uuid [16]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnDelUser is a free log retrieval operation binding the contract event 0x7bc4ecd86fa11d1e4eb1e479a0fc84e9e48b8b799ff51ffa6375f7fac922cf15.
//
// Solidity: e onDelUser(uuid bytes16)
func (_UsersData *UsersDataFilterer) FilterOnDelUser(opts *bind.FilterOpts) (*UsersDataOnDelUserIterator, error) {

	logs, sub, err := _UsersData.contract.FilterLogs(opts, "onDelUser")
	if err != nil {
		return nil, err
	}
	return &UsersDataOnDelUserIterator{contract: _UsersData.contract, event: "onDelUser", logs: logs, sub: sub}, nil
}

// WatchOnDelUser is a free log subscription operation binding the contract event 0x7bc4ecd86fa11d1e4eb1e479a0fc84e9e48b8b799ff51ffa6375f7fac922cf15.
//
// Solidity: e onDelUser(uuid bytes16)
func (_UsersData *UsersDataFilterer) WatchOnDelUser(opts *bind.WatchOpts, sink chan<- *UsersDataOnDelUser) (event.Subscription, error) {

	logs, sub, err := _UsersData.contract.WatchLogs(opts, "onDelUser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersDataOnDelUser)
				if err := _UsersData.contract.UnpackLog(event, "onDelUser", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UsersDataOnSetActiveIterator is returned from FilterOnSetActive and is used to iterate over the raw logs and unpacked data for OnSetActive events raised by the UsersData contract.
type UsersDataOnSetActiveIterator struct {
	Event *UsersDataOnSetActive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsersDataOnSetActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersDataOnSetActive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsersDataOnSetActive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersDataOnSetActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersDataOnSetActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersDataOnSetActive represents a OnSetActive event raised by the UsersData contract.
type UsersDataOnSetActive struct {
	Uuid   [16]byte
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOnSetActive is a free log retrieval operation binding the contract event 0x713e572bb47b301d3993c6d748aa8f43912b3e2864d1fd701e16e586fbb53905.
//
// Solidity: e onSetActive(uuid bytes16, active bool)
func (_UsersData *UsersDataFilterer) FilterOnSetActive(opts *bind.FilterOpts) (*UsersDataOnSetActiveIterator, error) {

	logs, sub, err := _UsersData.contract.FilterLogs(opts, "onSetActive")
	if err != nil {
		return nil, err
	}
	return &UsersDataOnSetActiveIterator{contract: _UsersData.contract, event: "onSetActive", logs: logs, sub: sub}, nil
}

// WatchOnSetActive is a free log subscription operation binding the contract event 0x713e572bb47b301d3993c6d748aa8f43912b3e2864d1fd701e16e586fbb53905.
//
// Solidity: e onSetActive(uuid bytes16, active bool)
func (_UsersData *UsersDataFilterer) WatchOnSetActive(opts *bind.WatchOpts, sink chan<- *UsersDataOnSetActive) (event.Subscription, error) {

	logs, sub, err := _UsersData.contract.WatchLogs(opts, "onSetActive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersDataOnSetActive)
				if err := _UsersData.contract.UnpackLog(event, "onSetActive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UsersDataOnSetIdCartNoHashIterator is returned from FilterOnSetIdCartNoHash and is used to iterate over the raw logs and unpacked data for OnSetIdCartNoHash events raised by the UsersData contract.
type UsersDataOnSetIdCartNoHashIterator struct {
	Event *UsersDataOnSetIdCartNoHash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsersDataOnSetIdCartNoHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersDataOnSetIdCartNoHash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsersDataOnSetIdCartNoHash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersDataOnSetIdCartNoHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersDataOnSetIdCartNoHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersDataOnSetIdCartNoHash represents a OnSetIdCartNoHash event raised by the UsersData contract.
type UsersDataOnSetIdCartNoHash struct {
	Uuid         [16]byte
	IdCartNoHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOnSetIdCartNoHash is a free log retrieval operation binding the contract event 0x7811c8a53358a5d74aa42672dd20237aa4f645d7f260bd174b35851fc382ce55.
//
// Solidity: e onSetIdCartNoHash(uuid bytes16, idCartNoHash bytes32)
func (_UsersData *UsersDataFilterer) FilterOnSetIdCartNoHash(opts *bind.FilterOpts) (*UsersDataOnSetIdCartNoHashIterator, error) {

	logs, sub, err := _UsersData.contract.FilterLogs(opts, "onSetIdCartNoHash")
	if err != nil {
		return nil, err
	}
	return &UsersDataOnSetIdCartNoHashIterator{contract: _UsersData.contract, event: "onSetIdCartNoHash", logs: logs, sub: sub}, nil
}

// WatchOnSetIdCartNoHash is a free log subscription operation binding the contract event 0x7811c8a53358a5d74aa42672dd20237aa4f645d7f260bd174b35851fc382ce55.
//
// Solidity: e onSetIdCartNoHash(uuid bytes16, idCartNoHash bytes32)
func (_UsersData *UsersDataFilterer) WatchOnSetIdCartNoHash(opts *bind.WatchOpts, sink chan<- *UsersDataOnSetIdCartNoHash) (event.Subscription, error) {

	logs, sub, err := _UsersData.contract.WatchLogs(opts, "onSetIdCartNoHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersDataOnSetIdCartNoHash)
				if err := _UsersData.contract.UnpackLog(event, "onSetIdCartNoHash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UsersDataOnSetOrgUuidIterator is returned from FilterOnSetOrgUuid and is used to iterate over the raw logs and unpacked data for OnSetOrgUuid events raised by the UsersData contract.
type UsersDataOnSetOrgUuidIterator struct {
	Event *UsersDataOnSetOrgUuid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsersDataOnSetOrgUuidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersDataOnSetOrgUuid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsersDataOnSetOrgUuid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersDataOnSetOrgUuidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersDataOnSetOrgUuidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersDataOnSetOrgUuid represents a OnSetOrgUuid event raised by the UsersData contract.
type UsersDataOnSetOrgUuid struct {
	Uuid    [16]byte
	OrgUuid [16]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOnSetOrgUuid is a free log retrieval operation binding the contract event 0xc3f343d86c0f17b6b536a18799f6c0f89d37202680752b8e438d3f14d48291f2.
//
// Solidity: e onSetOrgUuid(uuid bytes16, orgUuid bytes16)
func (_UsersData *UsersDataFilterer) FilterOnSetOrgUuid(opts *bind.FilterOpts) (*UsersDataOnSetOrgUuidIterator, error) {

	logs, sub, err := _UsersData.contract.FilterLogs(opts, "onSetOrgUuid")
	if err != nil {
		return nil, err
	}
	return &UsersDataOnSetOrgUuidIterator{contract: _UsersData.contract, event: "onSetOrgUuid", logs: logs, sub: sub}, nil
}

// WatchOnSetOrgUuid is a free log subscription operation binding the contract event 0xc3f343d86c0f17b6b536a18799f6c0f89d37202680752b8e438d3f14d48291f2.
//
// Solidity: e onSetOrgUuid(uuid bytes16, orgUuid bytes16)
func (_UsersData *UsersDataFilterer) WatchOnSetOrgUuid(opts *bind.WatchOpts, sink chan<- *UsersDataOnSetOrgUuid) (event.Subscription, error) {

	logs, sub, err := _UsersData.contract.WatchLogs(opts, "onSetOrgUuid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersDataOnSetOrgUuid)
				if err := _UsersData.contract.UnpackLog(event, "onSetOrgUuid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UsersDataOnSetTimeIterator is returned from FilterOnSetTime and is used to iterate over the raw logs and unpacked data for OnSetTime events raised by the UsersData contract.
type UsersDataOnSetTimeIterator struct {
	Event *UsersDataOnSetTime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsersDataOnSetTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersDataOnSetTime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsersDataOnSetTime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersDataOnSetTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersDataOnSetTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersDataOnSetTime represents a OnSetTime event raised by the UsersData contract.
type UsersDataOnSetTime struct {
	Uuid [16]byte
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnSetTime is a free log retrieval operation binding the contract event 0x50843b091ea0bec4b6ead7f725b4fcb769671b6948c77b9cd98b25ff2440d84c.
//
// Solidity: e onSetTime(uuid bytes16, time uint256)
func (_UsersData *UsersDataFilterer) FilterOnSetTime(opts *bind.FilterOpts) (*UsersDataOnSetTimeIterator, error) {

	logs, sub, err := _UsersData.contract.FilterLogs(opts, "onSetTime")
	if err != nil {
		return nil, err
	}
	return &UsersDataOnSetTimeIterator{contract: _UsersData.contract, event: "onSetTime", logs: logs, sub: sub}, nil
}

// WatchOnSetTime is a free log subscription operation binding the contract event 0x50843b091ea0bec4b6ead7f725b4fcb769671b6948c77b9cd98b25ff2440d84c.
//
// Solidity: e onSetTime(uuid bytes16, time uint256)
func (_UsersData *UsersDataFilterer) WatchOnSetTime(opts *bind.WatchOpts, sink chan<- *UsersDataOnSetTime) (event.Subscription, error) {

	logs, sub, err := _UsersData.contract.WatchLogs(opts, "onSetTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersDataOnSetTime)
				if err := _UsersData.contract.UnpackLog(event, "onSetTime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UsersDataOnSetUserAddressAndPublicKeyIterator is returned from FilterOnSetUserAddressAndPublicKey and is used to iterate over the raw logs and unpacked data for OnSetUserAddressAndPublicKey events raised by the UsersData contract.
type UsersDataOnSetUserAddressAndPublicKeyIterator struct {
	Event *UsersDataOnSetUserAddressAndPublicKey // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UsersDataOnSetUserAddressAndPublicKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersDataOnSetUserAddressAndPublicKey)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UsersDataOnSetUserAddressAndPublicKey)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UsersDataOnSetUserAddressAndPublicKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersDataOnSetUserAddressAndPublicKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersDataOnSetUserAddressAndPublicKey represents a OnSetUserAddressAndPublicKey event raised by the UsersData contract.
type UsersDataOnSetUserAddressAndPublicKey struct {
	Uuid        [16]byte
	UserAddress common.Address
	PublicKey   [2][32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOnSetUserAddressAndPublicKey is a free log retrieval operation binding the contract event 0xd64e6361ad680cd08aec52b6c8e50b6dda9a9c07b585c9ccc4375d984a488e03.
//
// Solidity: e onSetUserAddressAndPublicKey(uuid bytes16, userAddress address, publicKey bytes32[2])
func (_UsersData *UsersDataFilterer) FilterOnSetUserAddressAndPublicKey(opts *bind.FilterOpts) (*UsersDataOnSetUserAddressAndPublicKeyIterator, error) {

	logs, sub, err := _UsersData.contract.FilterLogs(opts, "onSetUserAddressAndPublicKey")
	if err != nil {
		return nil, err
	}
	return &UsersDataOnSetUserAddressAndPublicKeyIterator{contract: _UsersData.contract, event: "onSetUserAddressAndPublicKey", logs: logs, sub: sub}, nil
}

// WatchOnSetUserAddressAndPublicKey is a free log subscription operation binding the contract event 0xd64e6361ad680cd08aec52b6c8e50b6dda9a9c07b585c9ccc4375d984a488e03.
//
// Solidity: e onSetUserAddressAndPublicKey(uuid bytes16, userAddress address, publicKey bytes32[2])
func (_UsersData *UsersDataFilterer) WatchOnSetUserAddressAndPublicKey(opts *bind.WatchOpts, sink chan<- *UsersDataOnSetUserAddressAndPublicKey) (event.Subscription, error) {

	logs, sub, err := _UsersData.contract.WatchLogs(opts, "onSetUserAddressAndPublicKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersDataOnSetUserAddressAndPublicKey)
				if err := _UsersData.contract.UnpackLog(event, "onSetUserAddressAndPublicKey", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

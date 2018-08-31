// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package medi

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ControllerABI is the input ABI used to generate the binding from.
const ControllerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"start\",\"type\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getFileSignersAndDataByUuid\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"uint8[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"orgUuid\",\"type\":\"bytes16\"},{\"name\":\"publicKey\",\"type\":\"bytes32[2]\"},{\"name\":\"idCartNoHash\",\"type\":\"bytes32\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"addUser\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"publicKey\",\"type\":\"bytes32[2]\"},{\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32[4]\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"addOrg\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getOrgByUuid\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32[2]\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32[4]\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getUserByUuid\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes16\"},{\"name\":\"\",\"type\":\"bytes32[2]\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getFileByUuid\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"},{\"name\":\"\",\"type\":\"bytes16\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32[4]\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"addSign\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"ownerUuid\",\"type\":\"bytes16\"},{\"name\":\"fileType\",\"type\":\"bytes32\"},{\"name\":\"fileDesc\",\"type\":\"bytes32[4]\"},{\"name\":\"keccak256Hash\",\"type\":\"bytes32\"},{\"name\":\"sha256Hash\",\"type\":\"bytes32\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"addFile\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"easyCnsAddress\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"}]"

// ControllerBin is the compiled bytecode used for deploying new contracts.
const ControllerBin = `606060405234156200000d57fe5b60405160208062002ce0833981016040528080519060200190919050505b6200004a81620000a06401000000000262002a35176401000000009004565b1515620000575760006000fd5b80600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50620000dd565b6000600060008373ffffffffffffffffffffffffffffffffffffffff1614151515620000cc5760006000fd5b823b90506000811191505b50919050565b612bf380620000ed6000396000f3006060604052361561008c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806316f6c69e1461008e5780633380c48b1461023857806349464123146102c95780635b0f231d1461036a578063652801881461046d578063cf38f49814610560578063d1fa05f71461066b578063d2dbcbfa146106c4575bfe5b341561009657fe5b6100d160048080356fffffffffffffffffffffffffffffffff1916906020019091908035906020019091908035906020019091905050610795565b604051808060200180602001806020018060200185810385528981815181526020019150805190602001906020028083836000831461012f575b80518252602083111561012f5760208201915060208101905060208303925061010b565b50505090500185810384528881815181526020019150805190602001906020028083836000831461017f575b80518252602083111561017f5760208201915060208101905060208303925061015b565b5050509050018581038352878181518152602001915080519060200190602002808383600083146101cf575b8051825260208311156101cf576020820191506020810190506020830392506101ab565b50505090500185810382528681815181526020019150805190602001906020028083836000831461021f575b80518252602083111561021f576020820191506020810190506020830392506101fb565b5050509050019850505050505050505060405180910390f35b341561024057fe5b6102c760048080356fffffffffffffffffffffffffffffffff19169060200190919080356fffffffffffffffffffffffffffffffff1916906020019091908060400190600280602002604051908101604052809291908260026020028082843782019150505050509190803560001916906020019091908035906020019091905050610c2a565b005b34156102d157fe5b61036860048080356fffffffffffffffffffffffffffffffff19169060200190919080604001906002806020026040519081016040528092919082600260200280828437820191505050505091908035600019169060200190919080608001906004806020026040519081016040528092919082600460200280828437820191505050505091908035906020019091905050610dc2565b005b341561037257fe5b61039b60048080356fffffffffffffffffffffffffffffffff1916906020019091905050610f6a565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200185600260200280838360008314610407575b805182526020831115610407576020820191506020810190506020830392506103e3565b505050905001846000191660001916815260200183600460200280838360008314610451575b8051825260208311156104515760208201915060208101905060208303925061042d565b5050509050018281526020019550505050505060405180910390f35b341561047557fe5b61049e60048080356fffffffffffffffffffffffffffffffff19169060200190919050506113d2565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200184600260200280838360008314610536575b80518252602083111561053657602082019150602081019050602083039250610512565b50505090500183600019166000191681526020018281526020019550505050505060405180910390f35b341561056857fe5b61059160048080356fffffffffffffffffffffffffffffffff1916906020019091905050611831565b60405180886fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001876fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001866000191660001916815260200185600460200280838360008314610631575b8051825260208311156106315760208201915060208101905060208303925061060d565b5050509050018460001916600019168152602001836000191660001916815260200182815260200197505050505050505060405180910390f35b341561067357fe5b6106c260048080356fffffffffffffffffffffffffffffffff1916906020019091908035600019169060200190919080356000191690602001909190803560ff16906020019091905050611e42565b005b34156106cc57fe5b61079360048080356fffffffffffffffffffffffffffffffff19169060200190919080356fffffffffffffffffffffffffffffffff19169060200190919080356000191690602001909190806080019060048060200260405190810160405280929190826004602002808284378201915050505050919080356000191690602001909190803560001916906020019091908035600019169060200190919080356000191690602001909190803560ff1690602001909190803590602001909190505061206a565b005b61079d612b1f565b6107a5612b33565b6107ad612b33565b6107b5612b47565b60006107bf612b1f565b6107c7612b33565b6107cf612b33565b6107d7612b47565b600060008c101580156107ea575060008b115b80156107fa5750600a60ff168b11155b15156108065760006000fd5b61080e61233c565b151561081a5760006000fd5b6108f8600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635cbccaea8f6000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b15156108d957fe5b6102c65a03f115156108e757fe5b505050604051805190508d8d612565565b9550856040518059106109085750595b908082528060200260200182016040525b5094508560405180591061092a5750595b908082528060200260200182016040525b5093508560405180591061094c5750595b908082528060200260200182016040525b5092508560405180591061096e5750595b908082528060200260200182016040525b5091508b90505b85811015610c0e57600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166338fc12148e836000604051602001526040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180836fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200182815260200192505050602060405180830381600087803b1515610a5257fe5b6102c65a03f11515610a6057fe5b505050604051805190508582815181101515610a7857fe5b906020019060200201906fffffffffffffffffffffffffffffffff191690816fffffffffffffffffffffffffffffffff191681525050600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a8ecd6498e836000604051606001526040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180836fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200182815260200192505050606060405180830381600087803b1515610b7257fe5b6102c65a03f11515610b8057fe5b505050604051805190602001805190602001805190508684815181101515610ba457fe5b9060200190602002018685815181101515610bbb57fe5b9060200190602002018686815181101515610bd257fe5b9060200190602002018360ff1660ff168152508360001916600019168152508360001916600019168152505050505b8080600101915050610986565b8484848499509950995099505b50505050505093509350935093565b610c3261259e565b1515610c3e5760006000fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166388fa6b718633878787876040518763ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180876fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200184600260200280838360008314610d71575b805182526020831115610d7157602082019150602081019050602083039250610d4d565b50505090500183600019166000191681526020018281526020019650505050505050600060405180830381600087803b1515610da957fe5b6102c65a03f11515610db757fe5b5050505b5050505050565b610dca6127c7565b1515610dd65760006000fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663397d02bb8633878787876040518763ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180876fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200185600260200280838360008314610edd575b805182526020831115610edd57602082019150602081019050602083039250610eb9565b505050905001846000191660001916815260200183600460200280838360008314610f27575b805182526020831115610f2757602082019150602081019050602083039250610f03565b5050509050018281526020019650505050505050600060405180830381600087803b1515610f5157fe5b6102c65a03f11515610f5f57fe5b5050505b5050505050565b6000610f74612b5b565b6000610f7e612b87565b6000610f886127c7565b1515610f945760006000fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639d8e9108876000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b151561105057fe5b6102c65a03f1151561105e57fe5b50505060405180519050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e26787e886000604051604001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050604060405180830381600087803b151561112457fe5b6102c65a03f1151561113257fe5b50505060405180604001604052600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166312d9fdaf896000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b15156111fb57fe5b6102c65a03f1151561120957fe5b50505060405180519050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a5d95ad98a6000604051608001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050608060405180830381600087803b15156112cf57fe5b6102c65a03f115156112dd57fe5b50505060405180608001604052600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fe07283a8b6000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b15156113a657fe5b6102c65a03f115156113b457fe5b50505060405180519050945094509450945094505b91939590929450565b600060006113de612b5b565b600060006113ea61259e565b15156113f65760006000fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633d1ef2b7876000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b15156114b257fe5b6102c65a03f115156114c057fe5b50505060405180519050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663410f66e4886000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b151561158657fe5b6102c65a03f1151561159457fe5b50505060405180519050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e26787e896000604051604001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050604060405180830381600087803b151561165a57fe5b6102c65a03f1151561166857fe5b50505060405180604001604052600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f0eaa2638a6000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b151561173157fe5b6102c65a03f1151561173f57fe5b50505060405180519050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fe07283a8b6000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b151561180557fe5b6102c65a03f1151561181357fe5b50505060405180519050945094509450945094505b91939590929450565b60006000600061183f612b87565b60006000600061184d61233c565b15156118595760006000fd5b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635b379418896000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b151561191557fe5b6102c65a03f1151561192357fe5b50505060405180519050600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166306e4dffd8a6000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b15156119e957fe5b6102c65a03f115156119f757fe5b50505060405180519050600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c8c459b98b6000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b1515611abd57fe5b6102c65a03f11515611acb57fe5b50505060405180519050600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eedbe9158c6000604051608001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050608060405180830381600087803b1515611b9157fe5b6102c65a03f11515611b9f57fe5b50505060405180608001604052600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b21124ef8d6000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b1515611c6857fe5b6102c65a03f11515611c7657fe5b50505060405180519050600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c66a81458e6000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b1515611d3c57fe5b6102c65a03f11515611d4a57fe5b50505060405180519050600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fe07283a8f6000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b1515611e1057fe5b6102c65a03f11515611e1e57fe5b5050506040518051905096509650965096509650965096505b919395979092949650565b611e4a61259e565b1515611e565760006000fd5b611e5e61233c565b1515611e6a5760006000fd5b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663af1e3afc85600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639bd027ce336000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1515611f6b57fe5b6102c65a03f11515611f7957fe5b505050604051805190508686866040518663ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180866fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001846000191660001916815260200183600019166000191681526020018260ff1660ff16815260200195505050505050600060405180830381600087803b151561205257fe5b6102c65a03f1151561206057fe5b5050505b50505050565b61207261259e565b151561207e5760006000fd5b61208661233c565b15156120925760006000fd5b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663818030628b8b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639bd027ce336000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561219457fe5b6102c65a03f115156121a257fe5b505050604051805190508c8c8c8c8c8c8c8c6040518c63ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808c6fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020018b6fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020018a6fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020018960001916600019168152602001886004602002808383600083146122ab575b8051825260208311156122ab57602082019150602081019050602083039250612287565b50505090500187600019166000191681526020018660001916600019168152602001856000191660001916815260200184600019166000191681526020018360ff1660ff1681526020018281526020019b505050505050505050505050600060405180830381600087803b151561231e57fe5b6102c65a03f1151561232c57fe5b5050505b50505050505050505050565b60006000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6123866129f0565b6000604051602001526040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360008314612409575b805182526020831115612409576020820191506020810190506020830392506123e5565b505050905090810190601f1680156124355780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b151561245057fe5b6102c65a03f1151561245e57fe5b50505060405180519050905061247381612a35565b15156124825760009150612561565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151561255c5780600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600191505b5090565b6000600084841015156125785760006000fd5b6125828484612a71565b90508085106125915780612593565b845b91505b509392505050565b60006000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6125e8612a95565b6000604051602001526040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200182810382528381815181526020019150805190602001908083836000831461266b575b80518252602083111561266b57602082019150602081019050602083039250612647565b505050905090810190601f1680156126975780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15156126b257fe5b6102c65a03f115156126c057fe5b5050506040518051905090506126d581612a35565b15156126e457600091506127c3565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156127be5780600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600191505b5090565b60006000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e612811612ada565b6000604051602001526040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360008314612894575b80518252602083111561289457602082019150602081019050602083039250612870565b505050905090810190601f1680156128c05780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15156128db57fe5b6102c65a03f115156128e957fe5b5050506040518051905090506128fe81612a35565b151561290d57600091506129ec565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156129e75780600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600191505b5090565b6129f8612bb3565b604060405190810160405280600981526020017f46696c657344617461000000000000000000000000000000000000000000000081525090505b90565b6000600060008373ffffffffffffffffffffffffffffffffffffffff1614151515612a605760006000fd5b823b90506000811191505b50919050565b600060008284019050838110151515612a8a5760006000fd5b8091505b5092915050565b612a9d612bb3565b604060405190810160405280600981526020017f557365727344617461000000000000000000000000000000000000000000000081525090505b90565b612ae2612bb3565b604060405190810160405280600881526020017f4f7267734461746100000000000000000000000000000000000000000000000081525090505b90565b602060405190810160405280600081525090565b602060405190810160405280600081525090565b602060405190810160405280600081525090565b6040604051908101604052806002905b600060001916815260200190600190039081612b6b5790505090565b6080604051908101604052806004905b600060001916815260200190600190039081612b975790505090565b6020604051908101604052806000815250905600a165627a7a72305820d2aae1bc9cff657fbc59e13b64e9b9344ed5cfd2dd480ff598f206fc3663eb410029`

// DeployController deploys a new Ethereum contract, binding an instance of Controller to it.
func DeployController(auth *bind.TransactOpts, backend bind.ContractBackend, easyCnsAddress common.Address) (common.Address, *types.Transaction, *Controller, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ControllerBin), backend, easyCnsAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// GetFileByUuid is a free data retrieval call binding the contract method 0xcf38f498.
//
// Solidity: function getFileByUuid(uuid bytes16) constant returns(bytes16, bytes16, bytes32, bytes32[4], bytes32, bytes32, uint256)
func (_Controller *ControllerCaller) GetFileByUuid(opts *bind.CallOpts, uuid [16]byte) ([16]byte, [16]byte, [32]byte, [4][32]byte, [32]byte, [32]byte, *big.Int, error) {
	var (
		ret0 = new([16]byte)
		ret1 = new([16]byte)
		ret2 = new([32]byte)
		ret3 = new([4][32]byte)
		ret4 = new([32]byte)
		ret5 = new([32]byte)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _Controller.contract.Call(opts, out, "getFileByUuid", uuid)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetFileByUuid is a free data retrieval call binding the contract method 0xcf38f498.
//
// Solidity: function getFileByUuid(uuid bytes16) constant returns(bytes16, bytes16, bytes32, bytes32[4], bytes32, bytes32, uint256)
func (_Controller *ControllerSession) GetFileByUuid(uuid [16]byte) ([16]byte, [16]byte, [32]byte, [4][32]byte, [32]byte, [32]byte, *big.Int, error) {
	return _Controller.Contract.GetFileByUuid(&_Controller.CallOpts, uuid)
}

// GetFileByUuid is a free data retrieval call binding the contract method 0xcf38f498.
//
// Solidity: function getFileByUuid(uuid bytes16) constant returns(bytes16, bytes16, bytes32, bytes32[4], bytes32, bytes32, uint256)
func (_Controller *ControllerCallerSession) GetFileByUuid(uuid [16]byte) ([16]byte, [16]byte, [32]byte, [4][32]byte, [32]byte, [32]byte, *big.Int, error) {
	return _Controller.Contract.GetFileByUuid(&_Controller.CallOpts, uuid)
}

// GetFileSignersAndDataByUuid is a free data retrieval call binding the contract method 0x16f6c69e.
//
// Solidity: function getFileSignersAndDataByUuid(uuid bytes16, start uint256, limit uint256) constant returns(bytes16[], bytes32[], bytes32[], uint8[])
func (_Controller *ControllerCaller) GetFileSignersAndDataByUuid(opts *bind.CallOpts, uuid [16]byte, start *big.Int, limit *big.Int) ([][16]byte, [][32]byte, [][32]byte, []uint8, error) {
	var (
		ret0 = new([][16]byte)
		ret1 = new([][32]byte)
		ret2 = new([][32]byte)
		ret3 = new([]uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Controller.contract.Call(opts, out, "getFileSignersAndDataByUuid", uuid, start, limit)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetFileSignersAndDataByUuid is a free data retrieval call binding the contract method 0x16f6c69e.
//
// Solidity: function getFileSignersAndDataByUuid(uuid bytes16, start uint256, limit uint256) constant returns(bytes16[], bytes32[], bytes32[], uint8[])
func (_Controller *ControllerSession) GetFileSignersAndDataByUuid(uuid [16]byte, start *big.Int, limit *big.Int) ([][16]byte, [][32]byte, [][32]byte, []uint8, error) {
	return _Controller.Contract.GetFileSignersAndDataByUuid(&_Controller.CallOpts, uuid, start, limit)
}

// GetFileSignersAndDataByUuid is a free data retrieval call binding the contract method 0x16f6c69e.
//
// Solidity: function getFileSignersAndDataByUuid(uuid bytes16, start uint256, limit uint256) constant returns(bytes16[], bytes32[], bytes32[], uint8[])
func (_Controller *ControllerCallerSession) GetFileSignersAndDataByUuid(uuid [16]byte, start *big.Int, limit *big.Int) ([][16]byte, [][32]byte, [][32]byte, []uint8, error) {
	return _Controller.Contract.GetFileSignersAndDataByUuid(&_Controller.CallOpts, uuid, start, limit)
}

// GetOrgByUuid is a free data retrieval call binding the contract method 0x5b0f231d.
//
// Solidity: function getOrgByUuid(uuid bytes16) constant returns(address, bytes32[2], bytes32, bytes32[4], uint256)
func (_Controller *ControllerCaller) GetOrgByUuid(opts *bind.CallOpts, uuid [16]byte) (common.Address, [2][32]byte, [32]byte, [4][32]byte, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new([2][32]byte)
		ret2 = new([32]byte)
		ret3 = new([4][32]byte)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Controller.contract.Call(opts, out, "getOrgByUuid", uuid)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetOrgByUuid is a free data retrieval call binding the contract method 0x5b0f231d.
//
// Solidity: function getOrgByUuid(uuid bytes16) constant returns(address, bytes32[2], bytes32, bytes32[4], uint256)
func (_Controller *ControllerSession) GetOrgByUuid(uuid [16]byte) (common.Address, [2][32]byte, [32]byte, [4][32]byte, *big.Int, error) {
	return _Controller.Contract.GetOrgByUuid(&_Controller.CallOpts, uuid)
}

// GetOrgByUuid is a free data retrieval call binding the contract method 0x5b0f231d.
//
// Solidity: function getOrgByUuid(uuid bytes16) constant returns(address, bytes32[2], bytes32, bytes32[4], uint256)
func (_Controller *ControllerCallerSession) GetOrgByUuid(uuid [16]byte) (common.Address, [2][32]byte, [32]byte, [4][32]byte, *big.Int, error) {
	return _Controller.Contract.GetOrgByUuid(&_Controller.CallOpts, uuid)
}

// GetUserByUuid is a free data retrieval call binding the contract method 0x65280188.
//
// Solidity: function getUserByUuid(uuid bytes16) constant returns(address, bytes16, bytes32[2], bytes32, uint256)
func (_Controller *ControllerCaller) GetUserByUuid(opts *bind.CallOpts, uuid [16]byte) (common.Address, [16]byte, [2][32]byte, [32]byte, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new([16]byte)
		ret2 = new([2][32]byte)
		ret3 = new([32]byte)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Controller.contract.Call(opts, out, "getUserByUuid", uuid)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetUserByUuid is a free data retrieval call binding the contract method 0x65280188.
//
// Solidity: function getUserByUuid(uuid bytes16) constant returns(address, bytes16, bytes32[2], bytes32, uint256)
func (_Controller *ControllerSession) GetUserByUuid(uuid [16]byte) (common.Address, [16]byte, [2][32]byte, [32]byte, *big.Int, error) {
	return _Controller.Contract.GetUserByUuid(&_Controller.CallOpts, uuid)
}

// GetUserByUuid is a free data retrieval call binding the contract method 0x65280188.
//
// Solidity: function getUserByUuid(uuid bytes16) constant returns(address, bytes16, bytes32[2], bytes32, uint256)
func (_Controller *ControllerCallerSession) GetUserByUuid(uuid [16]byte) (common.Address, [16]byte, [2][32]byte, [32]byte, *big.Int, error) {
	return _Controller.Contract.GetUserByUuid(&_Controller.CallOpts, uuid)
}

// AddFile is a paid mutator transaction binding the contract method 0xd2dbcbfa.
//
// Solidity: function addFile(uuid bytes16, ownerUuid bytes16, fileType bytes32, fileDesc bytes32[4], keccak256Hash bytes32, sha256Hash bytes32, r bytes32, s bytes32, v uint8, time uint256) returns()
func (_Controller *ControllerTransactor) AddFile(opts *bind.TransactOpts, uuid [16]byte, ownerUuid [16]byte, fileType [32]byte, fileDesc [4][32]byte, keccak256Hash [32]byte, sha256Hash [32]byte, r [32]byte, s [32]byte, v uint8, time *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addFile", uuid, ownerUuid, fileType, fileDesc, keccak256Hash, sha256Hash, r, s, v, time)
}

// AddFile is a paid mutator transaction binding the contract method 0xd2dbcbfa.
//
// Solidity: function addFile(uuid bytes16, ownerUuid bytes16, fileType bytes32, fileDesc bytes32[4], keccak256Hash bytes32, sha256Hash bytes32, r bytes32, s bytes32, v uint8, time uint256) returns()
func (_Controller *ControllerSession) AddFile(uuid [16]byte, ownerUuid [16]byte, fileType [32]byte, fileDesc [4][32]byte, keccak256Hash [32]byte, sha256Hash [32]byte, r [32]byte, s [32]byte, v uint8, time *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AddFile(&_Controller.TransactOpts, uuid, ownerUuid, fileType, fileDesc, keccak256Hash, sha256Hash, r, s, v, time)
}

// AddFile is a paid mutator transaction binding the contract method 0xd2dbcbfa.
//
// Solidity: function addFile(uuid bytes16, ownerUuid bytes16, fileType bytes32, fileDesc bytes32[4], keccak256Hash bytes32, sha256Hash bytes32, r bytes32, s bytes32, v uint8, time uint256) returns()
func (_Controller *ControllerTransactorSession) AddFile(uuid [16]byte, ownerUuid [16]byte, fileType [32]byte, fileDesc [4][32]byte, keccak256Hash [32]byte, sha256Hash [32]byte, r [32]byte, s [32]byte, v uint8, time *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AddFile(&_Controller.TransactOpts, uuid, ownerUuid, fileType, fileDesc, keccak256Hash, sha256Hash, r, s, v, time)
}

// AddOrg is a paid mutator transaction binding the contract method 0x49464123.
//
// Solidity: function addOrg(uuid bytes16, publicKey bytes32[2], nameHash bytes32, name bytes32[4], time uint256) returns()
func (_Controller *ControllerTransactor) AddOrg(opts *bind.TransactOpts, uuid [16]byte, publicKey [2][32]byte, nameHash [32]byte, name [4][32]byte, time *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addOrg", uuid, publicKey, nameHash, name, time)
}

// AddOrg is a paid mutator transaction binding the contract method 0x49464123.
//
// Solidity: function addOrg(uuid bytes16, publicKey bytes32[2], nameHash bytes32, name bytes32[4], time uint256) returns()
func (_Controller *ControllerSession) AddOrg(uuid [16]byte, publicKey [2][32]byte, nameHash [32]byte, name [4][32]byte, time *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AddOrg(&_Controller.TransactOpts, uuid, publicKey, nameHash, name, time)
}

// AddOrg is a paid mutator transaction binding the contract method 0x49464123.
//
// Solidity: function addOrg(uuid bytes16, publicKey bytes32[2], nameHash bytes32, name bytes32[4], time uint256) returns()
func (_Controller *ControllerTransactorSession) AddOrg(uuid [16]byte, publicKey [2][32]byte, nameHash [32]byte, name [4][32]byte, time *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AddOrg(&_Controller.TransactOpts, uuid, publicKey, nameHash, name, time)
}

// AddSign is a paid mutator transaction binding the contract method 0xd1fa05f7.
//
// Solidity: function addSign(uuid bytes16, r bytes32, s bytes32, v uint8) returns()
func (_Controller *ControllerTransactor) AddSign(opts *bind.TransactOpts, uuid [16]byte, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addSign", uuid, r, s, v)
}

// AddSign is a paid mutator transaction binding the contract method 0xd1fa05f7.
//
// Solidity: function addSign(uuid bytes16, r bytes32, s bytes32, v uint8) returns()
func (_Controller *ControllerSession) AddSign(uuid [16]byte, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Controller.Contract.AddSign(&_Controller.TransactOpts, uuid, r, s, v)
}

// AddSign is a paid mutator transaction binding the contract method 0xd1fa05f7.
//
// Solidity: function addSign(uuid bytes16, r bytes32, s bytes32, v uint8) returns()
func (_Controller *ControllerTransactorSession) AddSign(uuid [16]byte, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Controller.Contract.AddSign(&_Controller.TransactOpts, uuid, r, s, v)
}

// AddUser is a paid mutator transaction binding the contract method 0x3380c48b.
//
// Solidity: function addUser(uuid bytes16, orgUuid bytes16, publicKey bytes32[2], idCartNoHash bytes32, time uint256) returns()
func (_Controller *ControllerTransactor) AddUser(opts *bind.TransactOpts, uuid [16]byte, orgUuid [16]byte, publicKey [2][32]byte, idCartNoHash [32]byte, time *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addUser", uuid, orgUuid, publicKey, idCartNoHash, time)
}

// AddUser is a paid mutator transaction binding the contract method 0x3380c48b.
//
// Solidity: function addUser(uuid bytes16, orgUuid bytes16, publicKey bytes32[2], idCartNoHash bytes32, time uint256) returns()
func (_Controller *ControllerSession) AddUser(uuid [16]byte, orgUuid [16]byte, publicKey [2][32]byte, idCartNoHash [32]byte, time *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AddUser(&_Controller.TransactOpts, uuid, orgUuid, publicKey, idCartNoHash, time)
}

// AddUser is a paid mutator transaction binding the contract method 0x3380c48b.
//
// Solidity: function addUser(uuid bytes16, orgUuid bytes16, publicKey bytes32[2], idCartNoHash bytes32, time uint256) returns()
func (_Controller *ControllerTransactorSession) AddUser(uuid [16]byte, orgUuid [16]byte, publicKey [2][32]byte, idCartNoHash [32]byte, time *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AddUser(&_Controller.TransactOpts, uuid, orgUuid, publicKey, idCartNoHash, time)
}

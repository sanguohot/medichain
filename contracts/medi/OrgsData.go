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

// OrgsDataABI is the input ABI used to generate the binding from.
const OrgsDataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getNameHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"setActive\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[2]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSuperSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"orgAddress\",\"type\":\"address\"},{\"name\":\"publicKey\",\"type\":\"bytes32[2]\"}],\"name\":\"setOrgAddressAndPublicKey\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"orgAddress\",\"type\":\"address\"},{\"name\":\"publicKey\",\"type\":\"bytes32[2]\"},{\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32[4]\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"addOrg\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"setTime\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"delSuper\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUuidListSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getUuidByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"delOrg\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"},{\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32[4]\"}],\"name\":\"setNameHashAndName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"isUuidExist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getOrgAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"getUuidByNameHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[4]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSuperByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"orgAddress\",\"type\":\"address\"}],\"name\":\"getUuidByOrgAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addSuper\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"getTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"easyCnsAddress\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"orgAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"bytes32[2]\"},{\"indexed\":false,\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes32[4]\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"onAddOrg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"}],\"name\":\"onDelOrg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"onSetActive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"orgAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"bytes32[2]\"}],\"name\":\"onSetOrgAddressAndPublicKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes32[4]\"}],\"name\":\"onSetNameHashAndName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"onSetTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"onAddSuper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"onDelSuper\",\"type\":\"event\"}]"

// OrgsDataBin is the compiled bytecode used for deploying new contracts.
const OrgsDataBin = `606060405234156200000d57fe5b60405160208062002b13833981016040528080519060200190919050505b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b6200008d81620000e3640100000000026200297b176401000000009004565b15156200009a5760006000fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5062000120565b6000600060008373ffffffffffffffffffffffffffffffffffffffff16141515156200010f5760006000fd5b823b90506000811191505b50919050565b6129e380620001306000396000f30060606040523615610110576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806312d9fdaf1461011257806324716b55146101615780632e26787e1461019f578063343496211461021c5780633931cc9f14610242578063397d02bb146102c05780633f7a52b61461038057806365f337f5146103bc5780636b303682146103f25780637c333c0a146104185780637f67b908146104725780638a4a2c61146104a55780638eeded90146105115780639d8e91081461055c578063a177cbb4146105cf578063a5d95ad91461062d578063ab013c26146106aa578063c16433281461070a578063e1e8637b1461077a578063fe07283a146107b0575bfe5b341561011a57fe5b61014360048080356fffffffffffffffffffffffffffffffff19169060200190919050506107f7565b60405180826000191660001916815260200191505060405180910390f35b341561016957fe5b61019d60048080356fffffffffffffffffffffffffffffffff19169060200190919080351515906020019091905050610897565b005b34156101a757fe5b6101d060048080356fffffffffffffffffffffffffffffffff19169060200190919050506109c7565b604051808260026020028083836000831461020a575b80518252602083111561020a576020820191506020810190506020830392506101e6565b50505090500191505060405180910390f35b341561022457fe5b61022c610aac565b6040518082815260200191505060405180910390f35b341561024a57fe5b6102be60048080356fffffffffffffffffffffffffffffffff191690602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919080604001906002806020026040519081016040528092919082600260200280828437820191505050505091905050610aba565b005b34156102c857fe5b61037e60048080356fffffffffffffffffffffffffffffffff191690602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919080604001906002806020026040519081016040528092919082600260200280828437820191505050505091908035600019169060200190919080608001906004806020026040519081016040528092919082600460200280828437820191505050505091908035906020019091905050610e2b565b005b341561038857fe5b6103ba60048080356fffffffffffffffffffffffffffffffff19169060200190919080359060200190919050506113bb565b005b34156103c457fe5b6103f0600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611571565b005b34156103fa57fe5b610402611783565b6040518082815260200191505060405180910390f35b341561042057fe5b6104366004808035906020019091905050611791565b60405180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b341561047a57fe5b6104a360048080356fffffffffffffffffffffffffffffffff1916906020019091905050611801565b005b34156104ad57fe5b61050f60048080356fffffffffffffffffffffffffffffffff1916906020019091908035600019169060200190919080608001906004806020026040519081016040528092919082600460200280828437820191505050505091905050611aae565b005b341561051957fe5b61054260048080356fffffffffffffffffffffffffffffffff1916906020019091905050611e2e565b604051808215151515815260200191505060405180910390f35b341561056457fe5b61058d60048080356fffffffffffffffffffffffffffffffff1916906020019091905050611e92565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156105d757fe5b6105f1600480803560001916906020019091905050611f52565b60405180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b341561063557fe5b61065e60048080356fffffffffffffffffffffffffffffffff1916906020019091905050612023565b6040518082600460200280838360008314610698575b80518252602083111561069857602082019150602081019050602083039250610674565b50505090500191505060405180910390f35b34156106b257fe5b6106c86004808035906020019091905050612108565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561071257fe5b61073e600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612163565b60405180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b341561078257fe5b6107ae600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061227c565b005b34156107b857fe5b6107e160048080356fffffffffffffffffffffffffffffffff191690602001909190505061243a565b6040518082815260200191505060405180910390f35b60008160046000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff1615156108505760006000fd5b60046000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020016000206003015491505b5b50919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806108f857506108f7336124da565b5b15156109045760006000fd5b8060046000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160006101000a81548160ff0219169083151502179055507f713e572bb47b301d3993c6d748aa8f43912b3e2864d1fd701e16e586fbb53905828260405180836fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001821515151581526020019250505060405180910390a15b5b5050565b6109cf61275b565b8160046000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff161515610a265760006000fd5b60046000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020600101600280602002604051908101604052809291908260028015610a9d576020028201915b81546000191681526020019060010190808311610a85575b505050505091505b5b50919050565b600060018054905090505b90565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610b1b5750610b1a336124da565b5b1515610b275760006000fd5b8260046000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff161515610b7e5760006000fd5b600070010000000000000000000000000000000002846fffffffffffffffffffffffffffffffff191614158015610bcc575060008373ffffffffffffffffffffffffffffffffffffffff1614155b8015610bdd5750610bdc82612554565b5b1515610be95760006000fd5b610bf2836125a4565b1515610bfe5760006000fd5b610c088383612634565b1515610c145760006000fd5b8260046000866fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160046000866fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020600101906002610cda929190612787565b5083600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff0219169083700100000000000000000000000000000000900402179055507fd83647f29e9b03c927a4fa98e046f8931a1fa2f097acca63953dab74e81f8a5584848460405180846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182600260200280838360008314610e0f575b805182526020831115610e0f57602082019150602081019050602083039250610deb565b505050905001935050505060405180910390a15b5b505b505050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610e8c5750610e8b336124da565b5b1515610e985760006000fd5b8560046000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff16151515610ef05760006000fd5b600070010000000000000000000000000000000002876fffffffffffffffffffffffffffffffff191614158015610f3e575060008673ffffffffffffffffffffffffffffffffffffffff1614155b8015610f4f5750610f4e85612554565b5b8015610f6357506000600102846000191614155b8015610f70575060008214155b1515610f7c5760006000fd5b610f8584612676565b1515610f915760006000fd5b610f9a866125a4565b1515610fa65760006000fd5b610fb08686612634565b1515610fbc5760006000fd5b826040518082600460200280838360008314610ff7575b805182526020831115610ff757602082019150602081019050602083039250610fd3565b505050905001915050604051809103902060001916846000191614151561101e5760006000fd5b60c0604051908101604052806001151581526020018773ffffffffffffffffffffffffffffffffffffffff168152602001868152602001856000191681526020018481526020018381525060046000896fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816001019060026111209291906127cd565b5060608201518160030190600019169055608082015181600401906004611148929190612813565b5060a082015181600801559050508660056000866000191660001916815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff02191690837001000000000000000000000000000000009004021790555086600660008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff021916908370010000000000000000000000000000000090040217905550600780548060010182816112369190612859565b91600052602060002090600291828204019190066010025b89909190916101000a8154816fffffffffffffffffffffffffffffffff021916908370010000000000000000000000000000000090040217905550507f72015815f96db3c437b25806f64ca98e0fcf191b8391b37b8f1fa8920989722787878787878760405180876fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200185600260200280838360008314611349575b80518252602083111561134957602082019150602081019050602083039250611325565b505050905001846000191660001916815260200183600460200280838360008314611393575b8051825260208311156113935760208201915060208101905060208303925061136f565b505050905001828152602001965050505050505060405180910390a15b5b505b505050505050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061141c575061141b336124da565b5b15156114285760006000fd5b8160046000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff16151561147f5760006000fd5b600070010000000000000000000000000000000002836fffffffffffffffffffffffffffffffff1916141580156114b7575060008214155b15156114c35760006000fd5b8160046000856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020600801819055507f50843b091ea0bec4b6ead7f725b4fcb769671b6948c77b9cd98b25ff2440d84c838360405180836fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020018281526020019250505060405180910390a15b5b505b5050565b60006000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156115d25760006000fd5b6115db836124da565b15156115e75760006000fd5b60009150600090505b6001805490508110156116835760018181548110151561160c57fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156116755760019150611683565b5b80806001019150506115f0565b81151561168c57fe5b60018181548110151561169b57fe5b906000526020600020900160005b6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690557f5561c1cda602083d14af2f0773152f273f6f9bf8feb17ce6265f94aec25bde0083604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a15b5b505050565b600060078054905090505b90565b6000600082101580156117aa57506117a7611783565b82105b15156117b65760006000fd5b6007828154811015156117c557fe5b90600052602060002090600291828204019190066010025b9054906101000a90047001000000000000000000000000000000000290505b919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806118625750611861336124da565b5b151561186e5760006000fd5b8060046000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff1615156118c55760006000fd5b6005600060046000856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020600301546000191660001916815260200190815260200160002060006101000a8154906fffffffffffffffffffffffffffffffff02191690556006600060046000856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a8154906fffffffffffffffffffffffffffffffff0219169055600060046000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160006101000a81548160ff0219169083151502179055507fd534b834b48c1cc340c4ea4cd46a15979d94896dfa8b58062437b765afb2ee4f8260405180826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200191505060405180910390a15b5b505b50565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480611b0f5750611b0e336124da565b5b1515611b1b5760006000fd5b8260046000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff161515611b725760006000fd5b600070010000000000000000000000000000000002846fffffffffffffffffffffffffffffffff191614158015611bb157506000600102836000191614155b1515611bbd5760006000fd5b816040518082600460200280838360008314611bf8575b805182526020831115611bf857602082019150602081019050602083039250611bd4565b5050509050019150506040518091039020600019168360001916141515611c1f5760006000fd5b6005600060046000876fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020600301546000191660001916815260200190815260200160002060006101000a8154906fffffffffffffffffffffffffffffffff02191690558260046000866fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060030181600019169055508160046000866fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020600401906004611d25929190612893565b508360056000856000191660001916815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff0219169083700100000000000000000000000000000000900402179055507fafb81b72f45edce9ee21fd3f57e8904776d25048eb0603f40a508b85a62adc4b84848460405180846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001836000191660001916815260200182600460200280838360008314611e12575b805182526020831115611e1257602082019150602081019050602083039250611dee565b505050905001935050505060405180910390a15b5b505b505050565b600060046000836fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff1615611e885760019050611e8d565b600090505b919050565b60008160046000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff161515611eeb5760006000fd5b60046000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505b5b50919050565b60006004600060056000856000191660001916815260200190815260200160002060009054906101000a9004700100000000000000000000000000000000026fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff161515611fe25760006000fd5b60056000836000191660001916815260200190815260200160002060009054906101000a90047001000000000000000000000000000000000290505b919050565b61202b6128d9565b8160046000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff1615156120825760006000fd5b60046000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020016000206004016004806020026040519081016040528092919082600480156120f9576020028201915b815460001916815260200190600101908083116120e1575b505050505091505b5b50919050565b6000612113826126e2565b151561211f5760006000fd5b60018281548110151561212e57fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b919050565b600060046000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a9004700100000000000000000000000000000000026fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff1615156122175760006000fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a90047001000000000000000000000000000000000290505b919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156122d95760006000fd5b60008173ffffffffffffffffffffffffffffffffffffffff16141515156123005760006000fd5b612309816124da565b1515156123165760006000fd5b6001805480600101828161232a9190612905565b916000526020600020900160005b83909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f920999c03a7fd19172e4c657c22aa62a1bfb5e027a8cf6a2a0ee9daf4c5a347481604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a15b5b50565b60008160046000826fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060000160009054906101000a900460ff1615156124935760006000fd5b60046000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020016000206008015491505b5b50919050565b600060008273ffffffffffffffffffffffffffffffffffffffff161415801561254c5750600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b90505b919050565b6000600060010282600060028110151561256a57fe5b6020020151600019161415801561259c5750600060010282600160028110151561259057fe5b60200201516000191614155b90505b919050565b6000600070010000000000000000000000000000000002600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a9004700100000000000000000000000000000000026fffffffffffffffffffffffffffffffff19161490505b919050565b60008273ffffffffffffffffffffffffffffffffffffffff1661265683612701565b73ffffffffffffffffffffffffffffffffffffffff161490505b92915050565b600060007001000000000000000000000000000000000260056000846000191660001916815260200190815260200160002060009054906101000a9004700100000000000000000000000000000000026fffffffffffffffffffffffffffffffff19161490505b919050565b6000600082101580156126f9575060018054905082105b90505b919050565b600081604051808260026020028083836000831461273e575b80518252602083111561273e5760208201915060208101905060208303925061271a565b50505090500191505060405180910390206001900490505b919050565b6040604051908101604052806002905b60006000191681526020019060019003908161276b5790505090565b82600281019282156127bc579160200282015b828111156127bb57825182906000191690559160200191906001019061279a565b5b5090506127c99190612931565b5090565b8260028101928215612802579160200282015b828111156128015782518290600019169055916020019190600101906127e0565b5b50905061280f9190612931565b5090565b8260048101928215612848579160200282015b82811115612847578251829060001916905591602001919060010190612826565b5b5090506128559190612931565b5090565b81548183558181151161288e57600101600290048160010160029004836000526020600020918201910161288d9190612956565b5b505050565b82600481019282156128c8579160200282015b828111156128c75782518290600019169055916020019190600101906128a6565b5b5090506128d59190612931565b5090565b6080604051908101604052806004905b6000600019168152602001906001900390816128e95790505090565b81548183558181151161292c5781836000526020600020918201910161292b9190612956565b5b505050565b61295391905b8082111561294f576000816000905550600101612937565b5090565b90565b61297891905b8082111561297457600081600090555060010161295c565b5090565b90565b6000600060008373ffffffffffffffffffffffffffffffffffffffff16141515156129a65760006000fd5b823b90506000811191505b509190505600a165627a7a7230582010927d638af02653a55cdf74be671b77251a881da0eaea4e627f39ae0ae13bf10029`

// DeployOrgsData deploys a new Ethereum contract, binding an instance of OrgsData to it.
func DeployOrgsData(auth *bind.TransactOpts, backend bind.ContractBackend, easyCnsAddress common.Address) (common.Address, *types.Transaction, *OrgsData, error) {
	parsed, err := abi.JSON(strings.NewReader(OrgsDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrgsDataBin), backend, easyCnsAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrgsData{OrgsDataCaller: OrgsDataCaller{contract: contract}, OrgsDataTransactor: OrgsDataTransactor{contract: contract}, OrgsDataFilterer: OrgsDataFilterer{contract: contract}}, nil
}

// OrgsData is an auto generated Go binding around an Ethereum contract.
type OrgsData struct {
	OrgsDataCaller     // Read-only binding to the contract
	OrgsDataTransactor // Write-only binding to the contract
	OrgsDataFilterer   // Log filterer for contract events
}

// OrgsDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrgsDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgsDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrgsDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgsDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrgsDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgsDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrgsDataSession struct {
	Contract     *OrgsData         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrgsDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrgsDataCallerSession struct {
	Contract *OrgsDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OrgsDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrgsDataTransactorSession struct {
	Contract     *OrgsDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OrgsDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrgsDataRaw struct {
	Contract *OrgsData // Generic contract binding to access the raw methods on
}

// OrgsDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrgsDataCallerRaw struct {
	Contract *OrgsDataCaller // Generic read-only contract binding to access the raw methods on
}

// OrgsDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrgsDataTransactorRaw struct {
	Contract *OrgsDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrgsData creates a new instance of OrgsData, bound to a specific deployed contract.
func NewOrgsData(address common.Address, backend bind.ContractBackend) (*OrgsData, error) {
	contract, err := bindOrgsData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrgsData{OrgsDataCaller: OrgsDataCaller{contract: contract}, OrgsDataTransactor: OrgsDataTransactor{contract: contract}, OrgsDataFilterer: OrgsDataFilterer{contract: contract}}, nil
}

// NewOrgsDataCaller creates a new read-only instance of OrgsData, bound to a specific deployed contract.
func NewOrgsDataCaller(address common.Address, caller bind.ContractCaller) (*OrgsDataCaller, error) {
	contract, err := bindOrgsData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrgsDataCaller{contract: contract}, nil
}

// NewOrgsDataTransactor creates a new write-only instance of OrgsData, bound to a specific deployed contract.
func NewOrgsDataTransactor(address common.Address, transactor bind.ContractTransactor) (*OrgsDataTransactor, error) {
	contract, err := bindOrgsData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrgsDataTransactor{contract: contract}, nil
}

// NewOrgsDataFilterer creates a new log filterer instance of OrgsData, bound to a specific deployed contract.
func NewOrgsDataFilterer(address common.Address, filterer bind.ContractFilterer) (*OrgsDataFilterer, error) {
	contract, err := bindOrgsData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrgsDataFilterer{contract: contract}, nil
}

// bindOrgsData binds a generic wrapper to an already deployed contract.
func bindOrgsData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrgsDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrgsData *OrgsDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrgsData.Contract.OrgsDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrgsData *OrgsDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrgsData.Contract.OrgsDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrgsData *OrgsDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrgsData.Contract.OrgsDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrgsData *OrgsDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrgsData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrgsData *OrgsDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrgsData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrgsData *OrgsDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrgsData.Contract.contract.Transact(opts, method, params...)
}

// GetName is a free data retrieval call binding the contract method 0xa5d95ad9.
//
// Solidity: function getName(uuid bytes16) constant returns(bytes32[4])
func (_OrgsData *OrgsDataCaller) GetName(opts *bind.CallOpts, uuid [16]byte) ([4][32]byte, error) {
	var (
		ret0 = new([4][32]byte)
	)
	out := ret0
	err := _OrgsData.contract.Call(opts, out, "getName", uuid)
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0xa5d95ad9.
//
// Solidity: function getName(uuid bytes16) constant returns(bytes32[4])
func (_OrgsData *OrgsDataSession) GetName(uuid [16]byte) ([4][32]byte, error) {
	return _OrgsData.Contract.GetName(&_OrgsData.CallOpts, uuid)
}

// GetName is a free data retrieval call binding the contract method 0xa5d95ad9.
//
// Solidity: function getName(uuid bytes16) constant returns(bytes32[4])
func (_OrgsData *OrgsDataCallerSession) GetName(uuid [16]byte) ([4][32]byte, error) {
	return _OrgsData.Contract.GetName(&_OrgsData.CallOpts, uuid)
}

// GetNameHash is a free data retrieval call binding the contract method 0x12d9fdaf.
//
// Solidity: function getNameHash(uuid bytes16) constant returns(bytes32)
func (_OrgsData *OrgsDataCaller) GetNameHash(opts *bind.CallOpts, uuid [16]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrgsData.contract.Call(opts, out, "getNameHash", uuid)
	return *ret0, err
}

// GetNameHash is a free data retrieval call binding the contract method 0x12d9fdaf.
//
// Solidity: function getNameHash(uuid bytes16) constant returns(bytes32)
func (_OrgsData *OrgsDataSession) GetNameHash(uuid [16]byte) ([32]byte, error) {
	return _OrgsData.Contract.GetNameHash(&_OrgsData.CallOpts, uuid)
}

// GetNameHash is a free data retrieval call binding the contract method 0x12d9fdaf.
//
// Solidity: function getNameHash(uuid bytes16) constant returns(bytes32)
func (_OrgsData *OrgsDataCallerSession) GetNameHash(uuid [16]byte) ([32]byte, error) {
	return _OrgsData.Contract.GetNameHash(&_OrgsData.CallOpts, uuid)
}

// GetOrgAddress is a free data retrieval call binding the contract method 0x9d8e9108.
//
// Solidity: function getOrgAddress(uuid bytes16) constant returns(address)
func (_OrgsData *OrgsDataCaller) GetOrgAddress(opts *bind.CallOpts, uuid [16]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OrgsData.contract.Call(opts, out, "getOrgAddress", uuid)
	return *ret0, err
}

// GetOrgAddress is a free data retrieval call binding the contract method 0x9d8e9108.
//
// Solidity: function getOrgAddress(uuid bytes16) constant returns(address)
func (_OrgsData *OrgsDataSession) GetOrgAddress(uuid [16]byte) (common.Address, error) {
	return _OrgsData.Contract.GetOrgAddress(&_OrgsData.CallOpts, uuid)
}

// GetOrgAddress is a free data retrieval call binding the contract method 0x9d8e9108.
//
// Solidity: function getOrgAddress(uuid bytes16) constant returns(address)
func (_OrgsData *OrgsDataCallerSession) GetOrgAddress(uuid [16]byte) (common.Address, error) {
	return _OrgsData.Contract.GetOrgAddress(&_OrgsData.CallOpts, uuid)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x2e26787e.
//
// Solidity: function getPublicKey(uuid bytes16) constant returns(bytes32[2])
func (_OrgsData *OrgsDataCaller) GetPublicKey(opts *bind.CallOpts, uuid [16]byte) ([2][32]byte, error) {
	var (
		ret0 = new([2][32]byte)
	)
	out := ret0
	err := _OrgsData.contract.Call(opts, out, "getPublicKey", uuid)
	return *ret0, err
}

// GetPublicKey is a free data retrieval call binding the contract method 0x2e26787e.
//
// Solidity: function getPublicKey(uuid bytes16) constant returns(bytes32[2])
func (_OrgsData *OrgsDataSession) GetPublicKey(uuid [16]byte) ([2][32]byte, error) {
	return _OrgsData.Contract.GetPublicKey(&_OrgsData.CallOpts, uuid)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x2e26787e.
//
// Solidity: function getPublicKey(uuid bytes16) constant returns(bytes32[2])
func (_OrgsData *OrgsDataCallerSession) GetPublicKey(uuid [16]byte) ([2][32]byte, error) {
	return _OrgsData.Contract.GetPublicKey(&_OrgsData.CallOpts, uuid)
}

// GetSuperByIndex is a free data retrieval call binding the contract method 0xab013c26.
//
// Solidity: function getSuperByIndex(index uint256) constant returns(address)
func (_OrgsData *OrgsDataCaller) GetSuperByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OrgsData.contract.Call(opts, out, "getSuperByIndex", index)
	return *ret0, err
}

// GetSuperByIndex is a free data retrieval call binding the contract method 0xab013c26.
//
// Solidity: function getSuperByIndex(index uint256) constant returns(address)
func (_OrgsData *OrgsDataSession) GetSuperByIndex(index *big.Int) (common.Address, error) {
	return _OrgsData.Contract.GetSuperByIndex(&_OrgsData.CallOpts, index)
}

// GetSuperByIndex is a free data retrieval call binding the contract method 0xab013c26.
//
// Solidity: function getSuperByIndex(index uint256) constant returns(address)
func (_OrgsData *OrgsDataCallerSession) GetSuperByIndex(index *big.Int) (common.Address, error) {
	return _OrgsData.Contract.GetSuperByIndex(&_OrgsData.CallOpts, index)
}

// GetSuperSize is a free data retrieval call binding the contract method 0x34349621.
//
// Solidity: function getSuperSize() constant returns(uint256)
func (_OrgsData *OrgsDataCaller) GetSuperSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OrgsData.contract.Call(opts, out, "getSuperSize")
	return *ret0, err
}

// GetSuperSize is a free data retrieval call binding the contract method 0x34349621.
//
// Solidity: function getSuperSize() constant returns(uint256)
func (_OrgsData *OrgsDataSession) GetSuperSize() (*big.Int, error) {
	return _OrgsData.Contract.GetSuperSize(&_OrgsData.CallOpts)
}

// GetSuperSize is a free data retrieval call binding the contract method 0x34349621.
//
// Solidity: function getSuperSize() constant returns(uint256)
func (_OrgsData *OrgsDataCallerSession) GetSuperSize() (*big.Int, error) {
	return _OrgsData.Contract.GetSuperSize(&_OrgsData.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0xfe07283a.
//
// Solidity: function getTime(uuid bytes16) constant returns(uint256)
func (_OrgsData *OrgsDataCaller) GetTime(opts *bind.CallOpts, uuid [16]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OrgsData.contract.Call(opts, out, "getTime", uuid)
	return *ret0, err
}

// GetTime is a free data retrieval call binding the contract method 0xfe07283a.
//
// Solidity: function getTime(uuid bytes16) constant returns(uint256)
func (_OrgsData *OrgsDataSession) GetTime(uuid [16]byte) (*big.Int, error) {
	return _OrgsData.Contract.GetTime(&_OrgsData.CallOpts, uuid)
}

// GetTime is a free data retrieval call binding the contract method 0xfe07283a.
//
// Solidity: function getTime(uuid bytes16) constant returns(uint256)
func (_OrgsData *OrgsDataCallerSession) GetTime(uuid [16]byte) (*big.Int, error) {
	return _OrgsData.Contract.GetTime(&_OrgsData.CallOpts, uuid)
}

// GetUuidByIndex is a free data retrieval call binding the contract method 0x7c333c0a.
//
// Solidity: function getUuidByIndex(index uint256) constant returns(bytes16)
func (_OrgsData *OrgsDataCaller) GetUuidByIndex(opts *bind.CallOpts, index *big.Int) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _OrgsData.contract.Call(opts, out, "getUuidByIndex", index)
	return *ret0, err
}

// GetUuidByIndex is a free data retrieval call binding the contract method 0x7c333c0a.
//
// Solidity: function getUuidByIndex(index uint256) constant returns(bytes16)
func (_OrgsData *OrgsDataSession) GetUuidByIndex(index *big.Int) ([16]byte, error) {
	return _OrgsData.Contract.GetUuidByIndex(&_OrgsData.CallOpts, index)
}

// GetUuidByIndex is a free data retrieval call binding the contract method 0x7c333c0a.
//
// Solidity: function getUuidByIndex(index uint256) constant returns(bytes16)
func (_OrgsData *OrgsDataCallerSession) GetUuidByIndex(index *big.Int) ([16]byte, error) {
	return _OrgsData.Contract.GetUuidByIndex(&_OrgsData.CallOpts, index)
}

// GetUuidByNameHash is a free data retrieval call binding the contract method 0xa177cbb4.
//
// Solidity: function getUuidByNameHash(nameHash bytes32) constant returns(bytes16)
func (_OrgsData *OrgsDataCaller) GetUuidByNameHash(opts *bind.CallOpts, nameHash [32]byte) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _OrgsData.contract.Call(opts, out, "getUuidByNameHash", nameHash)
	return *ret0, err
}

// GetUuidByNameHash is a free data retrieval call binding the contract method 0xa177cbb4.
//
// Solidity: function getUuidByNameHash(nameHash bytes32) constant returns(bytes16)
func (_OrgsData *OrgsDataSession) GetUuidByNameHash(nameHash [32]byte) ([16]byte, error) {
	return _OrgsData.Contract.GetUuidByNameHash(&_OrgsData.CallOpts, nameHash)
}

// GetUuidByNameHash is a free data retrieval call binding the contract method 0xa177cbb4.
//
// Solidity: function getUuidByNameHash(nameHash bytes32) constant returns(bytes16)
func (_OrgsData *OrgsDataCallerSession) GetUuidByNameHash(nameHash [32]byte) ([16]byte, error) {
	return _OrgsData.Contract.GetUuidByNameHash(&_OrgsData.CallOpts, nameHash)
}

// GetUuidByOrgAddress is a free data retrieval call binding the contract method 0xc1643328.
//
// Solidity: function getUuidByOrgAddress(orgAddress address) constant returns(bytes16)
func (_OrgsData *OrgsDataCaller) GetUuidByOrgAddress(opts *bind.CallOpts, orgAddress common.Address) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _OrgsData.contract.Call(opts, out, "getUuidByOrgAddress", orgAddress)
	return *ret0, err
}

// GetUuidByOrgAddress is a free data retrieval call binding the contract method 0xc1643328.
//
// Solidity: function getUuidByOrgAddress(orgAddress address) constant returns(bytes16)
func (_OrgsData *OrgsDataSession) GetUuidByOrgAddress(orgAddress common.Address) ([16]byte, error) {
	return _OrgsData.Contract.GetUuidByOrgAddress(&_OrgsData.CallOpts, orgAddress)
}

// GetUuidByOrgAddress is a free data retrieval call binding the contract method 0xc1643328.
//
// Solidity: function getUuidByOrgAddress(orgAddress address) constant returns(bytes16)
func (_OrgsData *OrgsDataCallerSession) GetUuidByOrgAddress(orgAddress common.Address) ([16]byte, error) {
	return _OrgsData.Contract.GetUuidByOrgAddress(&_OrgsData.CallOpts, orgAddress)
}

// GetUuidListSize is a free data retrieval call binding the contract method 0x6b303682.
//
// Solidity: function getUuidListSize() constant returns(uint256)
func (_OrgsData *OrgsDataCaller) GetUuidListSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OrgsData.contract.Call(opts, out, "getUuidListSize")
	return *ret0, err
}

// GetUuidListSize is a free data retrieval call binding the contract method 0x6b303682.
//
// Solidity: function getUuidListSize() constant returns(uint256)
func (_OrgsData *OrgsDataSession) GetUuidListSize() (*big.Int, error) {
	return _OrgsData.Contract.GetUuidListSize(&_OrgsData.CallOpts)
}

// GetUuidListSize is a free data retrieval call binding the contract method 0x6b303682.
//
// Solidity: function getUuidListSize() constant returns(uint256)
func (_OrgsData *OrgsDataCallerSession) GetUuidListSize() (*big.Int, error) {
	return _OrgsData.Contract.GetUuidListSize(&_OrgsData.CallOpts)
}

// IsUuidExist is a free data retrieval call binding the contract method 0x8eeded90.
//
// Solidity: function isUuidExist(uuid bytes16) constant returns(bool)
func (_OrgsData *OrgsDataCaller) IsUuidExist(opts *bind.CallOpts, uuid [16]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OrgsData.contract.Call(opts, out, "isUuidExist", uuid)
	return *ret0, err
}

// IsUuidExist is a free data retrieval call binding the contract method 0x8eeded90.
//
// Solidity: function isUuidExist(uuid bytes16) constant returns(bool)
func (_OrgsData *OrgsDataSession) IsUuidExist(uuid [16]byte) (bool, error) {
	return _OrgsData.Contract.IsUuidExist(&_OrgsData.CallOpts, uuid)
}

// IsUuidExist is a free data retrieval call binding the contract method 0x8eeded90.
//
// Solidity: function isUuidExist(uuid bytes16) constant returns(bool)
func (_OrgsData *OrgsDataCallerSession) IsUuidExist(uuid [16]byte) (bool, error) {
	return _OrgsData.Contract.IsUuidExist(&_OrgsData.CallOpts, uuid)
}

// AddOrg is a paid mutator transaction binding the contract method 0x397d02bb.
//
// Solidity: function addOrg(uuid bytes16, orgAddress address, publicKey bytes32[2], nameHash bytes32, name bytes32[4], time uint256) returns()
func (_OrgsData *OrgsDataTransactor) AddOrg(opts *bind.TransactOpts, uuid [16]byte, orgAddress common.Address, publicKey [2][32]byte, nameHash [32]byte, name [4][32]byte, time *big.Int) (*types.Transaction, error) {
	return _OrgsData.contract.Transact(opts, "addOrg", uuid, orgAddress, publicKey, nameHash, name, time)
}

// AddOrg is a paid mutator transaction binding the contract method 0x397d02bb.
//
// Solidity: function addOrg(uuid bytes16, orgAddress address, publicKey bytes32[2], nameHash bytes32, name bytes32[4], time uint256) returns()
func (_OrgsData *OrgsDataSession) AddOrg(uuid [16]byte, orgAddress common.Address, publicKey [2][32]byte, nameHash [32]byte, name [4][32]byte, time *big.Int) (*types.Transaction, error) {
	return _OrgsData.Contract.AddOrg(&_OrgsData.TransactOpts, uuid, orgAddress, publicKey, nameHash, name, time)
}

// AddOrg is a paid mutator transaction binding the contract method 0x397d02bb.
//
// Solidity: function addOrg(uuid bytes16, orgAddress address, publicKey bytes32[2], nameHash bytes32, name bytes32[4], time uint256) returns()
func (_OrgsData *OrgsDataTransactorSession) AddOrg(uuid [16]byte, orgAddress common.Address, publicKey [2][32]byte, nameHash [32]byte, name [4][32]byte, time *big.Int) (*types.Transaction, error) {
	return _OrgsData.Contract.AddOrg(&_OrgsData.TransactOpts, uuid, orgAddress, publicKey, nameHash, name, time)
}

// AddSuper is a paid mutator transaction binding the contract method 0xe1e8637b.
//
// Solidity: function addSuper(addr address) returns()
func (_OrgsData *OrgsDataTransactor) AddSuper(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _OrgsData.contract.Transact(opts, "addSuper", addr)
}

// AddSuper is a paid mutator transaction binding the contract method 0xe1e8637b.
//
// Solidity: function addSuper(addr address) returns()
func (_OrgsData *OrgsDataSession) AddSuper(addr common.Address) (*types.Transaction, error) {
	return _OrgsData.Contract.AddSuper(&_OrgsData.TransactOpts, addr)
}

// AddSuper is a paid mutator transaction binding the contract method 0xe1e8637b.
//
// Solidity: function addSuper(addr address) returns()
func (_OrgsData *OrgsDataTransactorSession) AddSuper(addr common.Address) (*types.Transaction, error) {
	return _OrgsData.Contract.AddSuper(&_OrgsData.TransactOpts, addr)
}

// DelOrg is a paid mutator transaction binding the contract method 0x7f67b908.
//
// Solidity: function delOrg(uuid bytes16) returns()
func (_OrgsData *OrgsDataTransactor) DelOrg(opts *bind.TransactOpts, uuid [16]byte) (*types.Transaction, error) {
	return _OrgsData.contract.Transact(opts, "delOrg", uuid)
}

// DelOrg is a paid mutator transaction binding the contract method 0x7f67b908.
//
// Solidity: function delOrg(uuid bytes16) returns()
func (_OrgsData *OrgsDataSession) DelOrg(uuid [16]byte) (*types.Transaction, error) {
	return _OrgsData.Contract.DelOrg(&_OrgsData.TransactOpts, uuid)
}

// DelOrg is a paid mutator transaction binding the contract method 0x7f67b908.
//
// Solidity: function delOrg(uuid bytes16) returns()
func (_OrgsData *OrgsDataTransactorSession) DelOrg(uuid [16]byte) (*types.Transaction, error) {
	return _OrgsData.Contract.DelOrg(&_OrgsData.TransactOpts, uuid)
}

// DelSuper is a paid mutator transaction binding the contract method 0x65f337f5.
//
// Solidity: function delSuper(addr address) returns()
func (_OrgsData *OrgsDataTransactor) DelSuper(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _OrgsData.contract.Transact(opts, "delSuper", addr)
}

// DelSuper is a paid mutator transaction binding the contract method 0x65f337f5.
//
// Solidity: function delSuper(addr address) returns()
func (_OrgsData *OrgsDataSession) DelSuper(addr common.Address) (*types.Transaction, error) {
	return _OrgsData.Contract.DelSuper(&_OrgsData.TransactOpts, addr)
}

// DelSuper is a paid mutator transaction binding the contract method 0x65f337f5.
//
// Solidity: function delSuper(addr address) returns()
func (_OrgsData *OrgsDataTransactorSession) DelSuper(addr common.Address) (*types.Transaction, error) {
	return _OrgsData.Contract.DelSuper(&_OrgsData.TransactOpts, addr)
}

// SetActive is a paid mutator transaction binding the contract method 0x24716b55.
//
// Solidity: function setActive(uuid bytes16, active bool) returns()
func (_OrgsData *OrgsDataTransactor) SetActive(opts *bind.TransactOpts, uuid [16]byte, active bool) (*types.Transaction, error) {
	return _OrgsData.contract.Transact(opts, "setActive", uuid, active)
}

// SetActive is a paid mutator transaction binding the contract method 0x24716b55.
//
// Solidity: function setActive(uuid bytes16, active bool) returns()
func (_OrgsData *OrgsDataSession) SetActive(uuid [16]byte, active bool) (*types.Transaction, error) {
	return _OrgsData.Contract.SetActive(&_OrgsData.TransactOpts, uuid, active)
}

// SetActive is a paid mutator transaction binding the contract method 0x24716b55.
//
// Solidity: function setActive(uuid bytes16, active bool) returns()
func (_OrgsData *OrgsDataTransactorSession) SetActive(uuid [16]byte, active bool) (*types.Transaction, error) {
	return _OrgsData.Contract.SetActive(&_OrgsData.TransactOpts, uuid, active)
}

// SetNameHashAndName is a paid mutator transaction binding the contract method 0x8a4a2c61.
//
// Solidity: function setNameHashAndName(uuid bytes16, nameHash bytes32, name bytes32[4]) returns()
func (_OrgsData *OrgsDataTransactor) SetNameHashAndName(opts *bind.TransactOpts, uuid [16]byte, nameHash [32]byte, name [4][32]byte) (*types.Transaction, error) {
	return _OrgsData.contract.Transact(opts, "setNameHashAndName", uuid, nameHash, name)
}

// SetNameHashAndName is a paid mutator transaction binding the contract method 0x8a4a2c61.
//
// Solidity: function setNameHashAndName(uuid bytes16, nameHash bytes32, name bytes32[4]) returns()
func (_OrgsData *OrgsDataSession) SetNameHashAndName(uuid [16]byte, nameHash [32]byte, name [4][32]byte) (*types.Transaction, error) {
	return _OrgsData.Contract.SetNameHashAndName(&_OrgsData.TransactOpts, uuid, nameHash, name)
}

// SetNameHashAndName is a paid mutator transaction binding the contract method 0x8a4a2c61.
//
// Solidity: function setNameHashAndName(uuid bytes16, nameHash bytes32, name bytes32[4]) returns()
func (_OrgsData *OrgsDataTransactorSession) SetNameHashAndName(uuid [16]byte, nameHash [32]byte, name [4][32]byte) (*types.Transaction, error) {
	return _OrgsData.Contract.SetNameHashAndName(&_OrgsData.TransactOpts, uuid, nameHash, name)
}

// SetOrgAddressAndPublicKey is a paid mutator transaction binding the contract method 0x3931cc9f.
//
// Solidity: function setOrgAddressAndPublicKey(uuid bytes16, orgAddress address, publicKey bytes32[2]) returns()
func (_OrgsData *OrgsDataTransactor) SetOrgAddressAndPublicKey(opts *bind.TransactOpts, uuid [16]byte, orgAddress common.Address, publicKey [2][32]byte) (*types.Transaction, error) {
	return _OrgsData.contract.Transact(opts, "setOrgAddressAndPublicKey", uuid, orgAddress, publicKey)
}

// SetOrgAddressAndPublicKey is a paid mutator transaction binding the contract method 0x3931cc9f.
//
// Solidity: function setOrgAddressAndPublicKey(uuid bytes16, orgAddress address, publicKey bytes32[2]) returns()
func (_OrgsData *OrgsDataSession) SetOrgAddressAndPublicKey(uuid [16]byte, orgAddress common.Address, publicKey [2][32]byte) (*types.Transaction, error) {
	return _OrgsData.Contract.SetOrgAddressAndPublicKey(&_OrgsData.TransactOpts, uuid, orgAddress, publicKey)
}

// SetOrgAddressAndPublicKey is a paid mutator transaction binding the contract method 0x3931cc9f.
//
// Solidity: function setOrgAddressAndPublicKey(uuid bytes16, orgAddress address, publicKey bytes32[2]) returns()
func (_OrgsData *OrgsDataTransactorSession) SetOrgAddressAndPublicKey(uuid [16]byte, orgAddress common.Address, publicKey [2][32]byte) (*types.Transaction, error) {
	return _OrgsData.Contract.SetOrgAddressAndPublicKey(&_OrgsData.TransactOpts, uuid, orgAddress, publicKey)
}

// SetTime is a paid mutator transaction binding the contract method 0x3f7a52b6.
//
// Solidity: function setTime(uuid bytes16, time uint256) returns()
func (_OrgsData *OrgsDataTransactor) SetTime(opts *bind.TransactOpts, uuid [16]byte, time *big.Int) (*types.Transaction, error) {
	return _OrgsData.contract.Transact(opts, "setTime", uuid, time)
}

// SetTime is a paid mutator transaction binding the contract method 0x3f7a52b6.
//
// Solidity: function setTime(uuid bytes16, time uint256) returns()
func (_OrgsData *OrgsDataSession) SetTime(uuid [16]byte, time *big.Int) (*types.Transaction, error) {
	return _OrgsData.Contract.SetTime(&_OrgsData.TransactOpts, uuid, time)
}

// SetTime is a paid mutator transaction binding the contract method 0x3f7a52b6.
//
// Solidity: function setTime(uuid bytes16, time uint256) returns()
func (_OrgsData *OrgsDataTransactorSession) SetTime(uuid [16]byte, time *big.Int) (*types.Transaction, error) {
	return _OrgsData.Contract.SetTime(&_OrgsData.TransactOpts, uuid, time)
}

// OrgsDataOnAddOrgIterator is returned from FilterOnAddOrg and is used to iterate over the raw logs and unpacked data for OnAddOrg events raised by the OrgsData contract.
type OrgsDataOnAddOrgIterator struct {
	Event *OrgsDataOnAddOrg // Event containing the contract specifics and raw log

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
func (it *OrgsDataOnAddOrgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgsDataOnAddOrg)
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
		it.Event = new(OrgsDataOnAddOrg)
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
func (it *OrgsDataOnAddOrgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgsDataOnAddOrgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgsDataOnAddOrg represents a OnAddOrg event raised by the OrgsData contract.
type OrgsDataOnAddOrg struct {
	Uuid       [16]byte
	OrgAddress common.Address
	PublicKey  [2][32]byte
	NameHash   [32]byte
	Name       [4][32]byte
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOnAddOrg is a free log retrieval operation binding the contract event 0x72015815f96db3c437b25806f64ca98e0fcf191b8391b37b8f1fa89209897227.
//
// Solidity: e onAddOrg(uuid bytes16, orgAddress address, publicKey bytes32[2], nameHash bytes32, name bytes32[4], time uint256)
func (_OrgsData *OrgsDataFilterer) FilterOnAddOrg(opts *bind.FilterOpts) (*OrgsDataOnAddOrgIterator, error) {

	logs, sub, err := _OrgsData.contract.FilterLogs(opts, "onAddOrg")
	if err != nil {
		return nil, err
	}
	return &OrgsDataOnAddOrgIterator{contract: _OrgsData.contract, event: "onAddOrg", logs: logs, sub: sub}, nil
}

// WatchOnAddOrg is a free log subscription operation binding the contract event 0x72015815f96db3c437b25806f64ca98e0fcf191b8391b37b8f1fa89209897227.
//
// Solidity: e onAddOrg(uuid bytes16, orgAddress address, publicKey bytes32[2], nameHash bytes32, name bytes32[4], time uint256)
func (_OrgsData *OrgsDataFilterer) WatchOnAddOrg(opts *bind.WatchOpts, sink chan<- *OrgsDataOnAddOrg) (event.Subscription, error) {

	logs, sub, err := _OrgsData.contract.WatchLogs(opts, "onAddOrg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgsDataOnAddOrg)
				if err := _OrgsData.contract.UnpackLog(event, "onAddOrg", log); err != nil {
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

// OrgsDataOnAddSuperIterator is returned from FilterOnAddSuper and is used to iterate over the raw logs and unpacked data for OnAddSuper events raised by the OrgsData contract.
type OrgsDataOnAddSuperIterator struct {
	Event *OrgsDataOnAddSuper // Event containing the contract specifics and raw log

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
func (it *OrgsDataOnAddSuperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgsDataOnAddSuper)
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
		it.Event = new(OrgsDataOnAddSuper)
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
func (it *OrgsDataOnAddSuperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgsDataOnAddSuperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgsDataOnAddSuper represents a OnAddSuper event raised by the OrgsData contract.
type OrgsDataOnAddSuper struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnAddSuper is a free log retrieval operation binding the contract event 0x920999c03a7fd19172e4c657c22aa62a1bfb5e027a8cf6a2a0ee9daf4c5a3474.
//
// Solidity: e onAddSuper(addr address)
func (_OrgsData *OrgsDataFilterer) FilterOnAddSuper(opts *bind.FilterOpts) (*OrgsDataOnAddSuperIterator, error) {

	logs, sub, err := _OrgsData.contract.FilterLogs(opts, "onAddSuper")
	if err != nil {
		return nil, err
	}
	return &OrgsDataOnAddSuperIterator{contract: _OrgsData.contract, event: "onAddSuper", logs: logs, sub: sub}, nil
}

// WatchOnAddSuper is a free log subscription operation binding the contract event 0x920999c03a7fd19172e4c657c22aa62a1bfb5e027a8cf6a2a0ee9daf4c5a3474.
//
// Solidity: e onAddSuper(addr address)
func (_OrgsData *OrgsDataFilterer) WatchOnAddSuper(opts *bind.WatchOpts, sink chan<- *OrgsDataOnAddSuper) (event.Subscription, error) {

	logs, sub, err := _OrgsData.contract.WatchLogs(opts, "onAddSuper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgsDataOnAddSuper)
				if err := _OrgsData.contract.UnpackLog(event, "onAddSuper", log); err != nil {
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

// OrgsDataOnDelOrgIterator is returned from FilterOnDelOrg and is used to iterate over the raw logs and unpacked data for OnDelOrg events raised by the OrgsData contract.
type OrgsDataOnDelOrgIterator struct {
	Event *OrgsDataOnDelOrg // Event containing the contract specifics and raw log

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
func (it *OrgsDataOnDelOrgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgsDataOnDelOrg)
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
		it.Event = new(OrgsDataOnDelOrg)
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
func (it *OrgsDataOnDelOrgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgsDataOnDelOrgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgsDataOnDelOrg represents a OnDelOrg event raised by the OrgsData contract.
type OrgsDataOnDelOrg struct {
	Uuid [16]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnDelOrg is a free log retrieval operation binding the contract event 0xd534b834b48c1cc340c4ea4cd46a15979d94896dfa8b58062437b765afb2ee4f.
//
// Solidity: e onDelOrg(uuid bytes16)
func (_OrgsData *OrgsDataFilterer) FilterOnDelOrg(opts *bind.FilterOpts) (*OrgsDataOnDelOrgIterator, error) {

	logs, sub, err := _OrgsData.contract.FilterLogs(opts, "onDelOrg")
	if err != nil {
		return nil, err
	}
	return &OrgsDataOnDelOrgIterator{contract: _OrgsData.contract, event: "onDelOrg", logs: logs, sub: sub}, nil
}

// WatchOnDelOrg is a free log subscription operation binding the contract event 0xd534b834b48c1cc340c4ea4cd46a15979d94896dfa8b58062437b765afb2ee4f.
//
// Solidity: e onDelOrg(uuid bytes16)
func (_OrgsData *OrgsDataFilterer) WatchOnDelOrg(opts *bind.WatchOpts, sink chan<- *OrgsDataOnDelOrg) (event.Subscription, error) {

	logs, sub, err := _OrgsData.contract.WatchLogs(opts, "onDelOrg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgsDataOnDelOrg)
				if err := _OrgsData.contract.UnpackLog(event, "onDelOrg", log); err != nil {
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

// OrgsDataOnDelSuperIterator is returned from FilterOnDelSuper and is used to iterate over the raw logs and unpacked data for OnDelSuper events raised by the OrgsData contract.
type OrgsDataOnDelSuperIterator struct {
	Event *OrgsDataOnDelSuper // Event containing the contract specifics and raw log

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
func (it *OrgsDataOnDelSuperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgsDataOnDelSuper)
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
		it.Event = new(OrgsDataOnDelSuper)
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
func (it *OrgsDataOnDelSuperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgsDataOnDelSuperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgsDataOnDelSuper represents a OnDelSuper event raised by the OrgsData contract.
type OrgsDataOnDelSuper struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnDelSuper is a free log retrieval operation binding the contract event 0x5561c1cda602083d14af2f0773152f273f6f9bf8feb17ce6265f94aec25bde00.
//
// Solidity: e onDelSuper(addr address)
func (_OrgsData *OrgsDataFilterer) FilterOnDelSuper(opts *bind.FilterOpts) (*OrgsDataOnDelSuperIterator, error) {

	logs, sub, err := _OrgsData.contract.FilterLogs(opts, "onDelSuper")
	if err != nil {
		return nil, err
	}
	return &OrgsDataOnDelSuperIterator{contract: _OrgsData.contract, event: "onDelSuper", logs: logs, sub: sub}, nil
}

// WatchOnDelSuper is a free log subscription operation binding the contract event 0x5561c1cda602083d14af2f0773152f273f6f9bf8feb17ce6265f94aec25bde00.
//
// Solidity: e onDelSuper(addr address)
func (_OrgsData *OrgsDataFilterer) WatchOnDelSuper(opts *bind.WatchOpts, sink chan<- *OrgsDataOnDelSuper) (event.Subscription, error) {

	logs, sub, err := _OrgsData.contract.WatchLogs(opts, "onDelSuper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgsDataOnDelSuper)
				if err := _OrgsData.contract.UnpackLog(event, "onDelSuper", log); err != nil {
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

// OrgsDataOnSetActiveIterator is returned from FilterOnSetActive and is used to iterate over the raw logs and unpacked data for OnSetActive events raised by the OrgsData contract.
type OrgsDataOnSetActiveIterator struct {
	Event *OrgsDataOnSetActive // Event containing the contract specifics and raw log

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
func (it *OrgsDataOnSetActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgsDataOnSetActive)
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
		it.Event = new(OrgsDataOnSetActive)
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
func (it *OrgsDataOnSetActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgsDataOnSetActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgsDataOnSetActive represents a OnSetActive event raised by the OrgsData contract.
type OrgsDataOnSetActive struct {
	Uuid   [16]byte
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOnSetActive is a free log retrieval operation binding the contract event 0x713e572bb47b301d3993c6d748aa8f43912b3e2864d1fd701e16e586fbb53905.
//
// Solidity: e onSetActive(uuid bytes16, active bool)
func (_OrgsData *OrgsDataFilterer) FilterOnSetActive(opts *bind.FilterOpts) (*OrgsDataOnSetActiveIterator, error) {

	logs, sub, err := _OrgsData.contract.FilterLogs(opts, "onSetActive")
	if err != nil {
		return nil, err
	}
	return &OrgsDataOnSetActiveIterator{contract: _OrgsData.contract, event: "onSetActive", logs: logs, sub: sub}, nil
}

// WatchOnSetActive is a free log subscription operation binding the contract event 0x713e572bb47b301d3993c6d748aa8f43912b3e2864d1fd701e16e586fbb53905.
//
// Solidity: e onSetActive(uuid bytes16, active bool)
func (_OrgsData *OrgsDataFilterer) WatchOnSetActive(opts *bind.WatchOpts, sink chan<- *OrgsDataOnSetActive) (event.Subscription, error) {

	logs, sub, err := _OrgsData.contract.WatchLogs(opts, "onSetActive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgsDataOnSetActive)
				if err := _OrgsData.contract.UnpackLog(event, "onSetActive", log); err != nil {
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

// OrgsDataOnSetNameHashAndNameIterator is returned from FilterOnSetNameHashAndName and is used to iterate over the raw logs and unpacked data for OnSetNameHashAndName events raised by the OrgsData contract.
type OrgsDataOnSetNameHashAndNameIterator struct {
	Event *OrgsDataOnSetNameHashAndName // Event containing the contract specifics and raw log

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
func (it *OrgsDataOnSetNameHashAndNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgsDataOnSetNameHashAndName)
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
		it.Event = new(OrgsDataOnSetNameHashAndName)
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
func (it *OrgsDataOnSetNameHashAndNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgsDataOnSetNameHashAndNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgsDataOnSetNameHashAndName represents a OnSetNameHashAndName event raised by the OrgsData contract.
type OrgsDataOnSetNameHashAndName struct {
	Uuid     [16]byte
	NameHash [32]byte
	Name     [4][32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOnSetNameHashAndName is a free log retrieval operation binding the contract event 0xafb81b72f45edce9ee21fd3f57e8904776d25048eb0603f40a508b85a62adc4b.
//
// Solidity: e onSetNameHashAndName(uuid bytes16, nameHash bytes32, name bytes32[4])
func (_OrgsData *OrgsDataFilterer) FilterOnSetNameHashAndName(opts *bind.FilterOpts) (*OrgsDataOnSetNameHashAndNameIterator, error) {

	logs, sub, err := _OrgsData.contract.FilterLogs(opts, "onSetNameHashAndName")
	if err != nil {
		return nil, err
	}
	return &OrgsDataOnSetNameHashAndNameIterator{contract: _OrgsData.contract, event: "onSetNameHashAndName", logs: logs, sub: sub}, nil
}

// WatchOnSetNameHashAndName is a free log subscription operation binding the contract event 0xafb81b72f45edce9ee21fd3f57e8904776d25048eb0603f40a508b85a62adc4b.
//
// Solidity: e onSetNameHashAndName(uuid bytes16, nameHash bytes32, name bytes32[4])
func (_OrgsData *OrgsDataFilterer) WatchOnSetNameHashAndName(opts *bind.WatchOpts, sink chan<- *OrgsDataOnSetNameHashAndName) (event.Subscription, error) {

	logs, sub, err := _OrgsData.contract.WatchLogs(opts, "onSetNameHashAndName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgsDataOnSetNameHashAndName)
				if err := _OrgsData.contract.UnpackLog(event, "onSetNameHashAndName", log); err != nil {
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

// OrgsDataOnSetOrgAddressAndPublicKeyIterator is returned from FilterOnSetOrgAddressAndPublicKey and is used to iterate over the raw logs and unpacked data for OnSetOrgAddressAndPublicKey events raised by the OrgsData contract.
type OrgsDataOnSetOrgAddressAndPublicKeyIterator struct {
	Event *OrgsDataOnSetOrgAddressAndPublicKey // Event containing the contract specifics and raw log

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
func (it *OrgsDataOnSetOrgAddressAndPublicKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgsDataOnSetOrgAddressAndPublicKey)
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
		it.Event = new(OrgsDataOnSetOrgAddressAndPublicKey)
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
func (it *OrgsDataOnSetOrgAddressAndPublicKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgsDataOnSetOrgAddressAndPublicKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgsDataOnSetOrgAddressAndPublicKey represents a OnSetOrgAddressAndPublicKey event raised by the OrgsData contract.
type OrgsDataOnSetOrgAddressAndPublicKey struct {
	Uuid       [16]byte
	OrgAddress common.Address
	PublicKey  [2][32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOnSetOrgAddressAndPublicKey is a free log retrieval operation binding the contract event 0xd83647f29e9b03c927a4fa98e046f8931a1fa2f097acca63953dab74e81f8a55.
//
// Solidity: e onSetOrgAddressAndPublicKey(uuid bytes16, orgAddress address, publicKey bytes32[2])
func (_OrgsData *OrgsDataFilterer) FilterOnSetOrgAddressAndPublicKey(opts *bind.FilterOpts) (*OrgsDataOnSetOrgAddressAndPublicKeyIterator, error) {

	logs, sub, err := _OrgsData.contract.FilterLogs(opts, "onSetOrgAddressAndPublicKey")
	if err != nil {
		return nil, err
	}
	return &OrgsDataOnSetOrgAddressAndPublicKeyIterator{contract: _OrgsData.contract, event: "onSetOrgAddressAndPublicKey", logs: logs, sub: sub}, nil
}

// WatchOnSetOrgAddressAndPublicKey is a free log subscription operation binding the contract event 0xd83647f29e9b03c927a4fa98e046f8931a1fa2f097acca63953dab74e81f8a55.
//
// Solidity: e onSetOrgAddressAndPublicKey(uuid bytes16, orgAddress address, publicKey bytes32[2])
func (_OrgsData *OrgsDataFilterer) WatchOnSetOrgAddressAndPublicKey(opts *bind.WatchOpts, sink chan<- *OrgsDataOnSetOrgAddressAndPublicKey) (event.Subscription, error) {

	logs, sub, err := _OrgsData.contract.WatchLogs(opts, "onSetOrgAddressAndPublicKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgsDataOnSetOrgAddressAndPublicKey)
				if err := _OrgsData.contract.UnpackLog(event, "onSetOrgAddressAndPublicKey", log); err != nil {
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

// OrgsDataOnSetTimeIterator is returned from FilterOnSetTime and is used to iterate over the raw logs and unpacked data for OnSetTime events raised by the OrgsData contract.
type OrgsDataOnSetTimeIterator struct {
	Event *OrgsDataOnSetTime // Event containing the contract specifics and raw log

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
func (it *OrgsDataOnSetTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgsDataOnSetTime)
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
		it.Event = new(OrgsDataOnSetTime)
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
func (it *OrgsDataOnSetTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgsDataOnSetTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgsDataOnSetTime represents a OnSetTime event raised by the OrgsData contract.
type OrgsDataOnSetTime struct {
	Uuid [16]byte
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnSetTime is a free log retrieval operation binding the contract event 0x50843b091ea0bec4b6ead7f725b4fcb769671b6948c77b9cd98b25ff2440d84c.
//
// Solidity: e onSetTime(uuid bytes16, time uint256)
func (_OrgsData *OrgsDataFilterer) FilterOnSetTime(opts *bind.FilterOpts) (*OrgsDataOnSetTimeIterator, error) {

	logs, sub, err := _OrgsData.contract.FilterLogs(opts, "onSetTime")
	if err != nil {
		return nil, err
	}
	return &OrgsDataOnSetTimeIterator{contract: _OrgsData.contract, event: "onSetTime", logs: logs, sub: sub}, nil
}

// WatchOnSetTime is a free log subscription operation binding the contract event 0x50843b091ea0bec4b6ead7f725b4fcb769671b6948c77b9cd98b25ff2440d84c.
//
// Solidity: e onSetTime(uuid bytes16, time uint256)
func (_OrgsData *OrgsDataFilterer) WatchOnSetTime(opts *bind.WatchOpts, sink chan<- *OrgsDataOnSetTime) (event.Subscription, error) {

	logs, sub, err := _OrgsData.contract.WatchLogs(opts, "onSetTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgsDataOnSetTime)
				if err := _OrgsData.contract.UnpackLog(event, "onSetTime", log); err != nil {
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

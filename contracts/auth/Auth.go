// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auth

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AuthABI is the input ABI used to generate the binding from.
const AuthABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getSha3FromHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getKeccak256FromHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"verifySignatureWithPrefix\",\"outputs\":[{\"name\":\"retAddr\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"verifySignatureWithoutPrefix\",\"outputs\":[{\"name\":\"retAddr\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"}]"

// AuthBin is the compiled bytecode used for deploying new contracts.
const AuthBin = `6060604052604060405190810160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152506000908051906020019061005092919061005e565b50341561005957fe5b610103565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061009f57805160ff19168380011785556100cd565b828001600101855582156100cd579182015b828111156100cc5782518255916020019190600101906100b1565b5b5090506100da91906100de565b5090565b61010091905b808211156100fc5760008160009055506001016100e4565b5090565b90565b6104c5806101126000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806329b41eb01461005c5780635faa442c1461009c578063b9639488146100dc578063bc8326c514610166575bfe5b341561006457fe5b61007e6004808035600019169060200190919050506101f0565b60405180826000191660001916815260200191505060405180910390f35b34156100a457fe5b6100be600480803560001916906020019091905050610274565b60405180826000191660001916815260200191505060405180910390f35b34156100e457fe5b61012460048080356000191690602001909190803560ff1690602001909190803560001916906020019091908035600019169060200190919050506102f8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561016e57fe5b6101ae60048080356000191690602001909190803560ff169060200190919080356000191690602001909190803560001916906020019091905050610408565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600060008260405180838054600181600116156101000203166002900480156102505780601f1061022e576101008083540402835291820191610250565b820191906000526020600020905b81548152906001019060200180831161023c575b5050826000191660001916815260200192505050604051809103902090505b919050565b600060008260405180838054600181600116156101000203166002900480156102d45780601f106102b25761010080835404028352918201916102d4565b820191906000526020600020905b8154815290600101906020018083116102c0575b5050826000191660001916815260200192505050604051809103902090505b919050565b60006000600086604051808380546001816001161561010002031660029004801561035a5780601f1061033857610100808354040283529182019161035a565b820191906000526020600020905b815481529060010190602001808311610346575b505082600019166000191681526020019250505060405180910390209050600181868686604051806000526020016040526000604051602001526040518085600019166000191681526020018460ff1660ff16815260200183600019166000191681526020018260001916600019168152602001945050505050602060405160208103908084039060008661646e5a03f115156103f357fe5b50506020604051035191505b50949350505050565b6000600185858585604051806000526020016040526000604051602001526040518085600019166000191681526020018460ff1660ff16815260200183600019166000191681526020018260001916600019168152602001945050505050602060405160208103908084039060008661646e5a03f1151561048557fe5b50506020604051035190505b9493505050505600a165627a7a723058204ac5f40d5eaded769e249f17b8539d57435562a278992b50624f7bd8330c42810029`

// DeployAuth deploys a new Ethereum contract, binding an instance of Auth to it.
func DeployAuth(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Auth, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuthBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Auth{AuthCaller: AuthCaller{contract: contract}, AuthTransactor: AuthTransactor{contract: contract}, AuthFilterer: AuthFilterer{contract: contract}}, nil
}

// Auth is an auto generated Go binding around an Ethereum contract.
type Auth struct {
	AuthCaller     // Read-only binding to the contract
	AuthTransactor // Write-only binding to the contract
	AuthFilterer   // Log filterer for contract events
}

// AuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthSession struct {
	Contract     *Auth             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthCallerSession struct {
	Contract *AuthCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthTransactorSession struct {
	Contract     *AuthTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthRaw struct {
	Contract *Auth // Generic contract binding to access the raw methods on
}

// AuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthCallerRaw struct {
	Contract *AuthCaller // Generic read-only contract binding to access the raw methods on
}

// AuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthTransactorRaw struct {
	Contract *AuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuth creates a new instance of Auth, bound to a specific deployed contract.
func NewAuth(address common.Address, backend bind.ContractBackend) (*Auth, error) {
	contract, err := bindAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auth{AuthCaller: AuthCaller{contract: contract}, AuthTransactor: AuthTransactor{contract: contract}, AuthFilterer: AuthFilterer{contract: contract}}, nil
}

// NewAuthCaller creates a new read-only instance of Auth, bound to a specific deployed contract.
func NewAuthCaller(address common.Address, caller bind.ContractCaller) (*AuthCaller, error) {
	contract, err := bindAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthCaller{contract: contract}, nil
}

// NewAuthTransactor creates a new write-only instance of Auth, bound to a specific deployed contract.
func NewAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthTransactor, error) {
	contract, err := bindAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthTransactor{contract: contract}, nil
}

// NewAuthFilterer creates a new log filterer instance of Auth, bound to a specific deployed contract.
func NewAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthFilterer, error) {
	contract, err := bindAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthFilterer{contract: contract}, nil
}

// bindAuth binds a generic wrapper to an already deployed contract.
func bindAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auth *AuthRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Auth.Contract.AuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auth *AuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auth.Contract.AuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auth *AuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auth.Contract.AuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auth *AuthCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Auth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auth *AuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auth *AuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auth.Contract.contract.Transact(opts, method, params...)
}

// GetKeccak256FromHash is a free data retrieval call binding the contract method 0x5faa442c.
//
// Solidity: function getKeccak256FromHash(hash bytes32) constant returns(bytes32)
func (_Auth *AuthCaller) GetKeccak256FromHash(opts *bind.CallOpts, hash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "getKeccak256FromHash", hash)
	return *ret0, err
}

// GetKeccak256FromHash is a free data retrieval call binding the contract method 0x5faa442c.
//
// Solidity: function getKeccak256FromHash(hash bytes32) constant returns(bytes32)
func (_Auth *AuthSession) GetKeccak256FromHash(hash [32]byte) ([32]byte, error) {
	return _Auth.Contract.GetKeccak256FromHash(&_Auth.CallOpts, hash)
}

// GetKeccak256FromHash is a free data retrieval call binding the contract method 0x5faa442c.
//
// Solidity: function getKeccak256FromHash(hash bytes32) constant returns(bytes32)
func (_Auth *AuthCallerSession) GetKeccak256FromHash(hash [32]byte) ([32]byte, error) {
	return _Auth.Contract.GetKeccak256FromHash(&_Auth.CallOpts, hash)
}

// GetSha3FromHash is a free data retrieval call binding the contract method 0x29b41eb0.
//
// Solidity: function getSha3FromHash(hash bytes32) constant returns(bytes32)
func (_Auth *AuthCaller) GetSha3FromHash(opts *bind.CallOpts, hash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "getSha3FromHash", hash)
	return *ret0, err
}

// GetSha3FromHash is a free data retrieval call binding the contract method 0x29b41eb0.
//
// Solidity: function getSha3FromHash(hash bytes32) constant returns(bytes32)
func (_Auth *AuthSession) GetSha3FromHash(hash [32]byte) ([32]byte, error) {
	return _Auth.Contract.GetSha3FromHash(&_Auth.CallOpts, hash)
}

// GetSha3FromHash is a free data retrieval call binding the contract method 0x29b41eb0.
//
// Solidity: function getSha3FromHash(hash bytes32) constant returns(bytes32)
func (_Auth *AuthCallerSession) GetSha3FromHash(hash [32]byte) ([32]byte, error) {
	return _Auth.Contract.GetSha3FromHash(&_Auth.CallOpts, hash)
}

// VerifySignatureWithPrefix is a free data retrieval call binding the contract method 0xb9639488.
//
// Solidity: function verifySignatureWithPrefix(hash bytes32, v uint8, r bytes32, s bytes32) constant returns(retAddr address)
func (_Auth *AuthCaller) VerifySignatureWithPrefix(opts *bind.CallOpts, hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "verifySignatureWithPrefix", hash, v, r, s)
	return *ret0, err
}

// VerifySignatureWithPrefix is a free data retrieval call binding the contract method 0xb9639488.
//
// Solidity: function verifySignatureWithPrefix(hash bytes32, v uint8, r bytes32, s bytes32) constant returns(retAddr address)
func (_Auth *AuthSession) VerifySignatureWithPrefix(hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Auth.Contract.VerifySignatureWithPrefix(&_Auth.CallOpts, hash, v, r, s)
}

// VerifySignatureWithPrefix is a free data retrieval call binding the contract method 0xb9639488.
//
// Solidity: function verifySignatureWithPrefix(hash bytes32, v uint8, r bytes32, s bytes32) constant returns(retAddr address)
func (_Auth *AuthCallerSession) VerifySignatureWithPrefix(hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Auth.Contract.VerifySignatureWithPrefix(&_Auth.CallOpts, hash, v, r, s)
}

// VerifySignatureWithoutPrefix is a free data retrieval call binding the contract method 0xbc8326c5.
//
// Solidity: function verifySignatureWithoutPrefix(hash bytes32, v uint8, r bytes32, s bytes32) constant returns(retAddr address)
func (_Auth *AuthCaller) VerifySignatureWithoutPrefix(opts *bind.CallOpts, hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "verifySignatureWithoutPrefix", hash, v, r, s)
	return *ret0, err
}

// VerifySignatureWithoutPrefix is a free data retrieval call binding the contract method 0xbc8326c5.
//
// Solidity: function verifySignatureWithoutPrefix(hash bytes32, v uint8, r bytes32, s bytes32) constant returns(retAddr address)
func (_Auth *AuthSession) VerifySignatureWithoutPrefix(hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Auth.Contract.VerifySignatureWithoutPrefix(&_Auth.CallOpts, hash, v, r, s)
}

// VerifySignatureWithoutPrefix is a free data retrieval call binding the contract method 0xbc8326c5.
//
// Solidity: function verifySignatureWithoutPrefix(hash bytes32, v uint8, r bytes32, s bytes32) constant returns(retAddr address)
func (_Auth *AuthCallerSession) VerifySignatureWithoutPrefix(hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Auth.Contract.VerifySignatureWithoutPrefix(&_Auth.CallOpts, hash, v, r, s)
}

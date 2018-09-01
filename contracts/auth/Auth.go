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
const AuthABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"verifyWithPrefix\",\"outputs\":[{\"name\":\"retAddr\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"verifyWithoutPrefix\",\"outputs\":[{\"name\":\"retAddr\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"}]"

// AuthBin is the compiled bytecode used for deploying new contracts.
const AuthBin = `6060604052341561000c57fe5b5b6103748061001c6000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680634d73b75714610046578063abdaa9d4146100d0575bfe5b341561004e57fe5b61008e60048080356000191690602001909190803560ff16906020019091908035600019169060200190919080356000191690602001909190505061015a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156100d857fe5b61011860048080356000191690602001909190803560ff1690602001909190803560001916906020019091908035600019169060200190919050506102a3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000610164610334565b6000604060405190810160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250915081876040518083805190602001908083835b602083106101d457805182526020820191506020810190506020830392506101b1565b6001836020036101000a03801982511681845116808217855250505050505090500182600019166000191681526020019250505060405180910390209050600181878787604051806000526020016040526000604051602001526040518085600019166000191681526020018460ff1660ff16815260200183600019166000191681526020018260001916600019168152602001945050505050602060405160208103908084039060008661646e5a03f1151561028d57fe5b50506020604051035192505b5050949350505050565b6000600185858585604051806000526020016040526000604051602001526040518085600019166000191681526020018460ff1660ff16815260200183600019166000191681526020018260001916600019168152602001945050505050602060405160208103908084039060008661646e5a03f1151561032057fe5b50506020604051035190505b949350505050565b6020604051908101604052806000815250905600a165627a7a723058202796880cd6a41373e7a935b269e1def7ad2bc997fbbb4b8a60dc9435b5a6d4620029`

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

// VerifyWithPrefix is a free data retrieval call binding the contract method 0x4d73b757.
//
// Solidity: function verifyWithPrefix(hash bytes32, v uint8, r bytes32, s bytes32) constant returns(retAddr address)
func (_Auth *AuthCaller) VerifyWithPrefix(opts *bind.CallOpts, hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "verifyWithPrefix", hash, v, r, s)
	return *ret0, err
}

// VerifyWithPrefix is a free data retrieval call binding the contract method 0x4d73b757.
//
// Solidity: function verifyWithPrefix(hash bytes32, v uint8, r bytes32, s bytes32) constant returns(retAddr address)
func (_Auth *AuthSession) VerifyWithPrefix(hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Auth.Contract.VerifyWithPrefix(&_Auth.CallOpts, hash, v, r, s)
}

// VerifyWithPrefix is a free data retrieval call binding the contract method 0x4d73b757.
//
// Solidity: function verifyWithPrefix(hash bytes32, v uint8, r bytes32, s bytes32) constant returns(retAddr address)
func (_Auth *AuthCallerSession) VerifyWithPrefix(hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Auth.Contract.VerifyWithPrefix(&_Auth.CallOpts, hash, v, r, s)
}

// VerifyWithoutPrefix is a free data retrieval call binding the contract method 0xabdaa9d4.
//
// Solidity: function verifyWithoutPrefix(hash bytes32, v uint8, r bytes32, s bytes32) constant returns(retAddr address)
func (_Auth *AuthCaller) VerifyWithoutPrefix(opts *bind.CallOpts, hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "verifyWithoutPrefix", hash, v, r, s)
	return *ret0, err
}

// VerifyWithoutPrefix is a free data retrieval call binding the contract method 0xabdaa9d4.
//
// Solidity: function verifyWithoutPrefix(hash bytes32, v uint8, r bytes32, s bytes32) constant returns(retAddr address)
func (_Auth *AuthSession) VerifyWithoutPrefix(hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Auth.Contract.VerifyWithoutPrefix(&_Auth.CallOpts, hash, v, r, s)
}

// VerifyWithoutPrefix is a free data retrieval call binding the contract method 0xabdaa9d4.
//
// Solidity: function verifyWithoutPrefix(hash bytes32, v uint8, r bytes32, s bytes32) constant returns(retAddr address)
func (_Auth *AuthCallerSession) VerifyWithoutPrefix(hash [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Auth.Contract.VerifyWithoutPrefix(&_Auth.CallOpts, hash, v, r, s)
}

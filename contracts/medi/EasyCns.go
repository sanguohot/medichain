// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package medi

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// EasyCnsABI is the input ABI used to generate the binding from.
const EasyCnsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"onSet\",\"type\":\"event\"}]"

// EasyCnsBin is the compiled bytecode used for deploying new contracts.
const EasyCnsBin = `6060604052341561000c57fe5b5b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b6104988061005f6000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063693ec85e14610046578063a815ff15146100e0575bfe5b341561004e57fe5b61009e600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610159565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156100e857fe5b610157600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506101ed565b005b60006001826040518082805190602001908083835b60208310610191578051825260208201915060208101905060208303925061016e565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561024a5760006000fd5b61025382610400565b151561025f5760006000fd5b6102688161041c565b15156102745760006000fd5b806001836040518082805190602001908083835b602083106102ab5780518252602082019150602081019050602083039250610288565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507ffb61c28e210c57f5ca20724afa35c1dbda662f286bd8d214606c6cfbdc41f439828260405180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281038252848181518152602001915080519060200190808383600083146103c0575b8051825260208311156103c05760208201915060208101905060208303925061039c565b505050905090810190601f1680156103ec5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15b5b5050565b600061040a610458565b82905060008151141591505b50919050565b6000600060008373ffffffffffffffffffffffffffffffffffffffff16141515156104475760006000fd5b823b90506000811191505b50919050565b6020604051908101604052806000815250905600a165627a7a7230582007f6413302e1da5bb69f245552b02417265bafe690c0f88c9e8229cb34a6fe230029`

// DeployEasyCns deploys a new Ethereum contract, binding an instance of EasyCns to it.
func DeployEasyCns(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EasyCns, error) {
	parsed, err := abi.JSON(strings.NewReader(EasyCnsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EasyCnsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EasyCns{EasyCnsCaller: EasyCnsCaller{contract: contract}, EasyCnsTransactor: EasyCnsTransactor{contract: contract}, EasyCnsFilterer: EasyCnsFilterer{contract: contract}}, nil
}

// EasyCns is an auto generated Go binding around an Ethereum contract.
type EasyCns struct {
	EasyCnsCaller     // Read-only binding to the contract
	EasyCnsTransactor // Write-only binding to the contract
	EasyCnsFilterer   // Log filterer for contract events
}

// EasyCnsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EasyCnsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EasyCnsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EasyCnsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EasyCnsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EasyCnsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EasyCnsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EasyCnsSession struct {
	Contract     *EasyCns          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EasyCnsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EasyCnsCallerSession struct {
	Contract *EasyCnsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EasyCnsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EasyCnsTransactorSession struct {
	Contract     *EasyCnsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EasyCnsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EasyCnsRaw struct {
	Contract *EasyCns // Generic contract binding to access the raw methods on
}

// EasyCnsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EasyCnsCallerRaw struct {
	Contract *EasyCnsCaller // Generic read-only contract binding to access the raw methods on
}

// EasyCnsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EasyCnsTransactorRaw struct {
	Contract *EasyCnsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEasyCns creates a new instance of EasyCns, bound to a specific deployed contract.
func NewEasyCns(address common.Address, backend bind.ContractBackend) (*EasyCns, error) {
	contract, err := bindEasyCns(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EasyCns{EasyCnsCaller: EasyCnsCaller{contract: contract}, EasyCnsTransactor: EasyCnsTransactor{contract: contract}, EasyCnsFilterer: EasyCnsFilterer{contract: contract}}, nil
}

// NewEasyCnsCaller creates a new read-only instance of EasyCns, bound to a specific deployed contract.
func NewEasyCnsCaller(address common.Address, caller bind.ContractCaller) (*EasyCnsCaller, error) {
	contract, err := bindEasyCns(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EasyCnsCaller{contract: contract}, nil
}

// NewEasyCnsTransactor creates a new write-only instance of EasyCns, bound to a specific deployed contract.
func NewEasyCnsTransactor(address common.Address, transactor bind.ContractTransactor) (*EasyCnsTransactor, error) {
	contract, err := bindEasyCns(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EasyCnsTransactor{contract: contract}, nil
}

// NewEasyCnsFilterer creates a new log filterer instance of EasyCns, bound to a specific deployed contract.
func NewEasyCnsFilterer(address common.Address, filterer bind.ContractFilterer) (*EasyCnsFilterer, error) {
	contract, err := bindEasyCns(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EasyCnsFilterer{contract: contract}, nil
}

// bindEasyCns binds a generic wrapper to an already deployed contract.
func bindEasyCns(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EasyCnsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EasyCns *EasyCnsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EasyCns.Contract.EasyCnsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EasyCns *EasyCnsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EasyCns.Contract.EasyCnsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EasyCns *EasyCnsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EasyCns.Contract.EasyCnsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EasyCns *EasyCnsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EasyCns.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EasyCns *EasyCnsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EasyCns.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EasyCns *EasyCnsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EasyCns.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(name string) constant returns(address)
func (_EasyCns *EasyCnsCaller) Get(opts *bind.CallOpts, name string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EasyCns.contract.Call(opts, out, "get", name)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(name string) constant returns(address)
func (_EasyCns *EasyCnsSession) Get(name string) (common.Address, error) {
	return _EasyCns.Contract.Get(&_EasyCns.CallOpts, name)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(name string) constant returns(address)
func (_EasyCns *EasyCnsCallerSession) Get(name string) (common.Address, error) {
	return _EasyCns.Contract.Get(&_EasyCns.CallOpts, name)
}

// Set is a paid mutator transaction binding the contract method 0xa815ff15.
//
// Solidity: function set(name string, addr address) returns()
func (_EasyCns *EasyCnsTransactor) Set(opts *bind.TransactOpts, name string, addr common.Address) (*types.Transaction, error) {
	return _EasyCns.contract.Transact(opts, "set", name, addr)
}

// Set is a paid mutator transaction binding the contract method 0xa815ff15.
//
// Solidity: function set(name string, addr address) returns()
func (_EasyCns *EasyCnsSession) Set(name string, addr common.Address) (*types.Transaction, error) {
	return _EasyCns.Contract.Set(&_EasyCns.TransactOpts, name, addr)
}

// Set is a paid mutator transaction binding the contract method 0xa815ff15.
//
// Solidity: function set(name string, addr address) returns()
func (_EasyCns *EasyCnsTransactorSession) Set(name string, addr common.Address) (*types.Transaction, error) {
	return _EasyCns.Contract.Set(&_EasyCns.TransactOpts, name, addr)
}

// EasyCnsOnSetIterator is returned from FilterOnSet and is used to iterate over the raw logs and unpacked data for OnSet events raised by the EasyCns contract.
type EasyCnsOnSetIterator struct {
	Event *EasyCnsOnSet // Event containing the contract specifics and raw log

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
func (it *EasyCnsOnSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EasyCnsOnSet)
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
		it.Event = new(EasyCnsOnSet)
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
func (it *EasyCnsOnSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EasyCnsOnSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EasyCnsOnSet represents a OnSet event raised by the EasyCns contract.
type EasyCnsOnSet struct {
	Name string
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnSet is a free log retrieval operation binding the contract event 0xfb61c28e210c57f5ca20724afa35c1dbda662f286bd8d214606c6cfbdc41f439.
//
// Solidity: e onSet(name string, addr address)
func (_EasyCns *EasyCnsFilterer) FilterOnSet(opts *bind.FilterOpts) (*EasyCnsOnSetIterator, error) {

	logs, sub, err := _EasyCns.contract.FilterLogs(opts, "onSet")
	if err != nil {
		return nil, err
	}
	return &EasyCnsOnSetIterator{contract: _EasyCns.contract, event: "onSet", logs: logs, sub: sub}, nil
}

// WatchOnSet is a free log subscription operation binding the contract event 0xfb61c28e210c57f5ca20724afa35c1dbda662f286bd8d214606c6cfbdc41f439.
//
// Solidity: e onSet(name string, addr address)
func (_EasyCns *EasyCnsFilterer) WatchOnSet(opts *bind.WatchOpts, sink chan<- *EasyCnsOnSet) (event.Subscription, error) {

	logs, sub, err := _EasyCns.contract.WatchLogs(opts, "onSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EasyCnsOnSet)
				if err := _EasyCns.contract.UnpackLog(event, "onSet", log); err != nil {
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

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// HelloWorldABI is the input ABI used to generate the binding from.
const HelloWorldABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"speak\",\"outputs\":[{\"name\":\"itSays\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newSaying\",\"type\":\"string\"}],\"name\":\"saySomethingElse\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newSaying\",\"type\":\"string\"}],\"name\":\"onSaySomethingElse\",\"type\":\"event\"}]"

// HelloWorldBin is the compiled bytecode used for deploying new contracts.
const HelloWorldBin = `0x6060604052341561000c57fe5b5b60408051808201909152600c8082527f48656c6c6f20576f726c64210000000000000000000000000000000000000000602090920191825261005191600091610058565b505b6100f8565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061009957805160ff19168380011785556100c6565b828001600101855582156100c6579182015b828111156100c65782518255916020019190600101906100ab565b5b506100d39291506100d7565b5090565b6100f591905b808211156100d357600081556001016100dd565b5090565b90565b6102d6806101076000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166350d853158114610045578063b07d7e39146100d5575bfe5b341561004d57fe5b61005561013f565b60408051602080825283518183015283519192839290830191850190808383821561009b575b80518252602083111561009b57601f19909201916020918201910161007b565b505050905090810190601f1680156100c75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156100dd57fe5b61012b600480803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437509496506101d895505050505050565b604080519115158252519081900360200190f35b6101476101f8565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101cd5780601f106101a2576101008083540402835291602001916101cd565b820191906000526020600020905b8154815290600101906020018083116101b057829003601f168201915b505050505090505b90565b80516000906101ed908290602085019061020a565b50600190505b919050565b60408051602081019091526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061024b57805160ff1916838001178555610278565b82800160010185558215610278579182015b8281111561027857825182559160200191906001019061025d565b5b50610285929150610289565b5090565b6101d591905b80821115610285576000815560010161028f565b5090565b905600a165627a7a72305820906e3bc2e30e73c18756e9386c694d735307d7076006862dd265483dca1b605c0029`

// DeployHelloWorld deploys a new Ethereum contract, binding an instance of HelloWorld to it.
func DeployHelloWorld(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HelloWorld, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloWorldABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HelloWorldBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HelloWorld{HelloWorldCaller: HelloWorldCaller{contract: contract}, HelloWorldTransactor: HelloWorldTransactor{contract: contract}, HelloWorldFilterer: HelloWorldFilterer{contract: contract}}, nil
}

// HelloWorld is an auto generated Go binding around an Ethereum contract.
type HelloWorld struct {
	HelloWorldCaller     // Read-only binding to the contract
	HelloWorldTransactor // Write-only binding to the contract
	HelloWorldFilterer   // Log filterer for contract events
}

// HelloWorldCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloWorldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloWorldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloWorldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloWorldSession struct {
	Contract     *HelloWorld       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloWorldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloWorldCallerSession struct {
	Contract *HelloWorldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HelloWorldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloWorldTransactorSession struct {
	Contract     *HelloWorldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HelloWorldRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloWorldRaw struct {
	Contract *HelloWorld // Generic contract binding to access the raw methods on
}

// HelloWorldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloWorldCallerRaw struct {
	Contract *HelloWorldCaller // Generic read-only contract binding to access the raw methods on
}

// HelloWorldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloWorldTransactorRaw struct {
	Contract *HelloWorldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloWorld creates a new instance of HelloWorld, bound to a specific deployed contract.
func NewHelloWorld(address common.Address, backend bind.ContractBackend) (*HelloWorld, error) {
	contract, err := bindHelloWorld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HelloWorld{HelloWorldCaller: HelloWorldCaller{contract: contract}, HelloWorldTransactor: HelloWorldTransactor{contract: contract}, HelloWorldFilterer: HelloWorldFilterer{contract: contract}}, nil
}

// NewHelloWorldCaller creates a new read-only instance of HelloWorld, bound to a specific deployed contract.
func NewHelloWorldCaller(address common.Address, caller bind.ContractCaller) (*HelloWorldCaller, error) {
	contract, err := bindHelloWorld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloWorldCaller{contract: contract}, nil
}

// NewHelloWorldTransactor creates a new write-only instance of HelloWorld, bound to a specific deployed contract.
func NewHelloWorldTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloWorldTransactor, error) {
	contract, err := bindHelloWorld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloWorldTransactor{contract: contract}, nil
}

// NewHelloWorldFilterer creates a new log filterer instance of HelloWorld, bound to a specific deployed contract.
func NewHelloWorldFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloWorldFilterer, error) {
	contract, err := bindHelloWorld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloWorldFilterer{contract: contract}, nil
}

// bindHelloWorld binds a generic wrapper to an already deployed contract.
func bindHelloWorld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloWorldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloWorld *HelloWorldRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HelloWorld.Contract.HelloWorldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloWorld *HelloWorldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloWorld.Contract.HelloWorldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloWorld *HelloWorldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloWorld.Contract.HelloWorldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloWorld *HelloWorldCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HelloWorld.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloWorld *HelloWorldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloWorld.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloWorld *HelloWorldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloWorld.Contract.contract.Transact(opts, method, params...)
}

// Speak is a free data retrieval call binding the contract method 0x50d85315.
//
// Solidity: function speak() constant returns(itSays string)
func (_HelloWorld *HelloWorldCaller) Speak(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _HelloWorld.contract.Call(opts, out, "speak")
	return *ret0, err
}

// Speak is a free data retrieval call binding the contract method 0x50d85315.
//
// Solidity: function speak() constant returns(itSays string)
func (_HelloWorld *HelloWorldSession) Speak() (string, error) {
	return _HelloWorld.Contract.Speak(&_HelloWorld.CallOpts)
}

// Speak is a free data retrieval call binding the contract method 0x50d85315.
//
// Solidity: function speak() constant returns(itSays string)
func (_HelloWorld *HelloWorldCallerSession) Speak() (string, error) {
	return _HelloWorld.Contract.Speak(&_HelloWorld.CallOpts)
}

// SaySomethingElse is a paid mutator transaction binding the contract method 0xb07d7e39.
//
// Solidity: function saySomethingElse(newSaying string) returns(success bool)
func (_HelloWorld *HelloWorldTransactor) SaySomethingElse(opts *bind.TransactOpts, newSaying string) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "saySomethingElse", newSaying)
}

// SaySomethingElse is a paid mutator transaction binding the contract method 0xb07d7e39.
//
// Solidity: function saySomethingElse(newSaying string) returns(success bool)
func (_HelloWorld *HelloWorldSession) SaySomethingElse(newSaying string) (*types.Transaction, error) {
	return _HelloWorld.Contract.SaySomethingElse(&_HelloWorld.TransactOpts, newSaying)
}

// SaySomethingElse is a paid mutator transaction binding the contract method 0xb07d7e39.
//
// Solidity: function saySomethingElse(newSaying string) returns(success bool)
func (_HelloWorld *HelloWorldTransactorSession) SaySomethingElse(newSaying string) (*types.Transaction, error) {
	return _HelloWorld.Contract.SaySomethingElse(&_HelloWorld.TransactOpts, newSaying)
}

// HelloWorldOnSaySomethingElseIterator is returned from FilterOnSaySomethingElse and is used to iterate over the raw logs and unpacked data for OnSaySomethingElse events raised by the HelloWorld contract.
type HelloWorldOnSaySomethingElseIterator struct {
	Event *HelloWorldOnSaySomethingElse // Event containing the contract specifics and raw log

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
func (it *HelloWorldOnSaySomethingElseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloWorldOnSaySomethingElse)
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
		it.Event = new(HelloWorldOnSaySomethingElse)
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
func (it *HelloWorldOnSaySomethingElseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloWorldOnSaySomethingElseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloWorldOnSaySomethingElse represents a OnSaySomethingElse event raised by the HelloWorld contract.
type HelloWorldOnSaySomethingElse struct {
	NewSaying string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnSaySomethingElse is a free log retrieval operation binding the contract event 0x87a4687b43c451ab438a8f4ac187d795809dd4e395a6a9ff46b9120ad40848cf.
//
// Solidity: e onSaySomethingElse(newSaying string)
func (_HelloWorld *HelloWorldFilterer) FilterOnSaySomethingElse(opts *bind.FilterOpts) (*HelloWorldOnSaySomethingElseIterator, error) {

	logs, sub, err := _HelloWorld.contract.FilterLogs(opts, "onSaySomethingElse")
	if err != nil {
		return nil, err
	}
	return &HelloWorldOnSaySomethingElseIterator{contract: _HelloWorld.contract, event: "onSaySomethingElse", logs: logs, sub: sub}, nil
}

// WatchOnSaySomethingElse is a free log subscription operation binding the contract event 0x87a4687b43c451ab438a8f4ac187d795809dd4e395a6a9ff46b9120ad40848cf.
//
// Solidity: e onSaySomethingElse(newSaying string)
func (_HelloWorld *HelloWorldFilterer) WatchOnSaySomethingElse(opts *bind.WatchOpts, sink chan<- *HelloWorldOnSaySomethingElse) (event.Subscription, error) {

	logs, sub, err := _HelloWorld.contract.WatchLogs(opts, "onSaySomethingElse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloWorldOnSaySomethingElse)
				if err := _HelloWorld.contract.UnpackLog(event, "onSaySomethingElse", log); err != nil {
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

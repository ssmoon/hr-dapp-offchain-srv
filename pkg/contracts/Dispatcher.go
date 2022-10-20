// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DispatcherMetaData contains all meta data concerning the Dispatcher contract.
var DispatcherMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"AddressImported\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"repository\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"importAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"removeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"names\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"destinations\",\"type\":\"address[]\"}],\"name\":\"areAddressesImported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"tryGetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"getExistedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DispatcherABI is the input ABI used to generate the binding from.
// Deprecated: Use DispatcherMetaData.ABI instead.
var DispatcherABI = DispatcherMetaData.ABI

// Dispatcher is an auto generated Go binding around an Ethereum contract.
type Dispatcher struct {
	DispatcherCaller     // Read-only binding to the contract
	DispatcherTransactor // Write-only binding to the contract
	DispatcherFilterer   // Log filterer for contract events
}

// DispatcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type DispatcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DispatcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DispatcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DispatcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DispatcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DispatcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DispatcherSession struct {
	Contract     *Dispatcher       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DispatcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DispatcherCallerSession struct {
	Contract *DispatcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DispatcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DispatcherTransactorSession struct {
	Contract     *DispatcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DispatcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type DispatcherRaw struct {
	Contract *Dispatcher // Generic contract binding to access the raw methods on
}

// DispatcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DispatcherCallerRaw struct {
	Contract *DispatcherCaller // Generic read-only contract binding to access the raw methods on
}

// DispatcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DispatcherTransactorRaw struct {
	Contract *DispatcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDispatcher creates a new instance of Dispatcher, bound to a specific deployed contract.
func NewDispatcher(address common.Address, backend bind.ContractBackend) (*Dispatcher, error) {
	contract, err := bindDispatcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dispatcher{DispatcherCaller: DispatcherCaller{contract: contract}, DispatcherTransactor: DispatcherTransactor{contract: contract}, DispatcherFilterer: DispatcherFilterer{contract: contract}}, nil
}

// NewDispatcherCaller creates a new read-only instance of Dispatcher, bound to a specific deployed contract.
func NewDispatcherCaller(address common.Address, caller bind.ContractCaller) (*DispatcherCaller, error) {
	contract, err := bindDispatcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DispatcherCaller{contract: contract}, nil
}

// NewDispatcherTransactor creates a new write-only instance of Dispatcher, bound to a specific deployed contract.
func NewDispatcherTransactor(address common.Address, transactor bind.ContractTransactor) (*DispatcherTransactor, error) {
	contract, err := bindDispatcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DispatcherTransactor{contract: contract}, nil
}

// NewDispatcherFilterer creates a new log filterer instance of Dispatcher, bound to a specific deployed contract.
func NewDispatcherFilterer(address common.Address, filterer bind.ContractFilterer) (*DispatcherFilterer, error) {
	contract, err := bindDispatcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DispatcherFilterer{contract: contract}, nil
}

// bindDispatcher binds a generic wrapper to an already deployed contract.
func bindDispatcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DispatcherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dispatcher *DispatcherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dispatcher.Contract.DispatcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dispatcher *DispatcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dispatcher.Contract.DispatcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dispatcher *DispatcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dispatcher.Contract.DispatcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dispatcher *DispatcherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dispatcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dispatcher *DispatcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dispatcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dispatcher *DispatcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dispatcher.Contract.contract.Transact(opts, method, params...)
}

// AreAddressesImported is a free data retrieval call binding the contract method 0x9f42102f.
//
// Solidity: function areAddressesImported(bytes32[] names, address[] destinations) view returns(bool)
func (_Dispatcher *DispatcherCaller) AreAddressesImported(opts *bind.CallOpts, names [][32]byte, destinations []common.Address) (bool, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "areAddressesImported", names, destinations)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreAddressesImported is a free data retrieval call binding the contract method 0x9f42102f.
//
// Solidity: function areAddressesImported(bytes32[] names, address[] destinations) view returns(bool)
func (_Dispatcher *DispatcherSession) AreAddressesImported(names [][32]byte, destinations []common.Address) (bool, error) {
	return _Dispatcher.Contract.AreAddressesImported(&_Dispatcher.CallOpts, names, destinations)
}

// AreAddressesImported is a free data retrieval call binding the contract method 0x9f42102f.
//
// Solidity: function areAddressesImported(bytes32[] names, address[] destinations) view returns(bool)
func (_Dispatcher *DispatcherCallerSession) AreAddressesImported(names [][32]byte, destinations []common.Address) (bool, error) {
	return _Dispatcher.Contract.AreAddressesImported(&_Dispatcher.CallOpts, names, destinations)
}

// GetExistedAddress is a free data retrieval call binding the contract method 0x6e1425fd.
//
// Solidity: function getExistedAddress(bytes32 name, string reason) view returns(address)
func (_Dispatcher *DispatcherCaller) GetExistedAddress(opts *bind.CallOpts, name [32]byte, reason string) (common.Address, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "getExistedAddress", name, reason)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExistedAddress is a free data retrieval call binding the contract method 0x6e1425fd.
//
// Solidity: function getExistedAddress(bytes32 name, string reason) view returns(address)
func (_Dispatcher *DispatcherSession) GetExistedAddress(name [32]byte, reason string) (common.Address, error) {
	return _Dispatcher.Contract.GetExistedAddress(&_Dispatcher.CallOpts, name, reason)
}

// GetExistedAddress is a free data retrieval call binding the contract method 0x6e1425fd.
//
// Solidity: function getExistedAddress(bytes32 name, string reason) view returns(address)
func (_Dispatcher *DispatcherCallerSession) GetExistedAddress(name [32]byte, reason string) (common.Address, error) {
	return _Dispatcher.Contract.GetExistedAddress(&_Dispatcher.CallOpts, name, reason)
}

// Repository is a free data retrieval call binding the contract method 0x187f7935.
//
// Solidity: function repository(bytes32 ) view returns(address)
func (_Dispatcher *DispatcherCaller) Repository(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "repository", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Repository is a free data retrieval call binding the contract method 0x187f7935.
//
// Solidity: function repository(bytes32 ) view returns(address)
func (_Dispatcher *DispatcherSession) Repository(arg0 [32]byte) (common.Address, error) {
	return _Dispatcher.Contract.Repository(&_Dispatcher.CallOpts, arg0)
}

// Repository is a free data retrieval call binding the contract method 0x187f7935.
//
// Solidity: function repository(bytes32 ) view returns(address)
func (_Dispatcher *DispatcherCallerSession) Repository(arg0 [32]byte) (common.Address, error) {
	return _Dispatcher.Contract.Repository(&_Dispatcher.CallOpts, arg0)
}

// TryGetAddress is a free data retrieval call binding the contract method 0x0e01b557.
//
// Solidity: function tryGetAddress(bytes32 name) view returns(address)
func (_Dispatcher *DispatcherCaller) TryGetAddress(opts *bind.CallOpts, name [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "tryGetAddress", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TryGetAddress is a free data retrieval call binding the contract method 0x0e01b557.
//
// Solidity: function tryGetAddress(bytes32 name) view returns(address)
func (_Dispatcher *DispatcherSession) TryGetAddress(name [32]byte) (common.Address, error) {
	return _Dispatcher.Contract.TryGetAddress(&_Dispatcher.CallOpts, name)
}

// TryGetAddress is a free data retrieval call binding the contract method 0x0e01b557.
//
// Solidity: function tryGetAddress(bytes32 name) view returns(address)
func (_Dispatcher *DispatcherCallerSession) TryGetAddress(name [32]byte) (common.Address, error) {
	return _Dispatcher.Contract.TryGetAddress(&_Dispatcher.CallOpts, name)
}

// ImportAddress is a paid mutator transaction binding the contract method 0xcfcd591c.
//
// Solidity: function importAddress(bytes32 name, address destination) returns()
func (_Dispatcher *DispatcherTransactor) ImportAddress(opts *bind.TransactOpts, name [32]byte, destination common.Address) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "importAddress", name, destination)
}

// ImportAddress is a paid mutator transaction binding the contract method 0xcfcd591c.
//
// Solidity: function importAddress(bytes32 name, address destination) returns()
func (_Dispatcher *DispatcherSession) ImportAddress(name [32]byte, destination common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.ImportAddress(&_Dispatcher.TransactOpts, name, destination)
}

// ImportAddress is a paid mutator transaction binding the contract method 0xcfcd591c.
//
// Solidity: function importAddress(bytes32 name, address destination) returns()
func (_Dispatcher *DispatcherTransactorSession) ImportAddress(name [32]byte, destination common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.ImportAddress(&_Dispatcher.TransactOpts, name, destination)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x9faf6fb6.
//
// Solidity: function removeAddress(bytes32 name) returns()
func (_Dispatcher *DispatcherTransactor) RemoveAddress(opts *bind.TransactOpts, name [32]byte) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "removeAddress", name)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x9faf6fb6.
//
// Solidity: function removeAddress(bytes32 name) returns()
func (_Dispatcher *DispatcherSession) RemoveAddress(name [32]byte) (*types.Transaction, error) {
	return _Dispatcher.Contract.RemoveAddress(&_Dispatcher.TransactOpts, name)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x9faf6fb6.
//
// Solidity: function removeAddress(bytes32 name) returns()
func (_Dispatcher *DispatcherTransactorSession) RemoveAddress(name [32]byte) (*types.Transaction, error) {
	return _Dispatcher.Contract.RemoveAddress(&_Dispatcher.TransactOpts, name)
}

// DispatcherAddressImportedIterator is returned from FilterAddressImported and is used to iterate over the raw logs and unpacked data for AddressImported events raised by the Dispatcher contract.
type DispatcherAddressImportedIterator struct {
	Event *DispatcherAddressImported // Event containing the contract specifics and raw log

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
func (it *DispatcherAddressImportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherAddressImported)
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
		it.Event = new(DispatcherAddressImported)
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
func (it *DispatcherAddressImportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherAddressImportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherAddressImported represents a AddressImported event raised by the Dispatcher contract.
type DispatcherAddressImported struct {
	Name        [32]byte
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddressImported is a free log retrieval operation binding the contract event 0xefe884cc7f82a6cf3cf68f64221519dcf96b5cae9048e1bb008ee32cd05aaa91.
//
// Solidity: event AddressImported(bytes32 name, address destination)
func (_Dispatcher *DispatcherFilterer) FilterAddressImported(opts *bind.FilterOpts) (*DispatcherAddressImportedIterator, error) {

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "AddressImported")
	if err != nil {
		return nil, err
	}
	return &DispatcherAddressImportedIterator{contract: _Dispatcher.contract, event: "AddressImported", logs: logs, sub: sub}, nil
}

// WatchAddressImported is a free log subscription operation binding the contract event 0xefe884cc7f82a6cf3cf68f64221519dcf96b5cae9048e1bb008ee32cd05aaa91.
//
// Solidity: event AddressImported(bytes32 name, address destination)
func (_Dispatcher *DispatcherFilterer) WatchAddressImported(opts *bind.WatchOpts, sink chan<- *DispatcherAddressImported) (event.Subscription, error) {

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "AddressImported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherAddressImported)
				if err := _Dispatcher.contract.UnpackLog(event, "AddressImported", log); err != nil {
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

// ParseAddressImported is a log parse operation binding the contract event 0xefe884cc7f82a6cf3cf68f64221519dcf96b5cae9048e1bb008ee32cd05aaa91.
//
// Solidity: event AddressImported(bytes32 name, address destination)
func (_Dispatcher *DispatcherFilterer) ParseAddressImported(log types.Log) (*DispatcherAddressImported, error) {
	event := new(DispatcherAddressImported)
	if err := _Dispatcher.contract.UnpackLog(event, "AddressImported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

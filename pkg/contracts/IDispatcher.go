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

// IDispatcherMetaData contains all meta data concerning the IDispatcher contract.
var IDispatcherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"importAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"names\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"destinations\",\"type\":\"address[]\"}],\"name\":\"areAddressesImported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"tryGetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"getExistedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IDispatcherABI is the input ABI used to generate the binding from.
// Deprecated: Use IDispatcherMetaData.ABI instead.
var IDispatcherABI = IDispatcherMetaData.ABI

// IDispatcher is an auto generated Go binding around an Ethereum contract.
type IDispatcher struct {
	IDispatcherCaller     // Read-only binding to the contract
	IDispatcherTransactor // Write-only binding to the contract
	IDispatcherFilterer   // Log filterer for contract events
}

// IDispatcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDispatcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDispatcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDispatcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDispatcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDispatcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDispatcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDispatcherSession struct {
	Contract     *IDispatcher      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDispatcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDispatcherCallerSession struct {
	Contract *IDispatcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IDispatcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDispatcherTransactorSession struct {
	Contract     *IDispatcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IDispatcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDispatcherRaw struct {
	Contract *IDispatcher // Generic contract binding to access the raw methods on
}

// IDispatcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDispatcherCallerRaw struct {
	Contract *IDispatcherCaller // Generic read-only contract binding to access the raw methods on
}

// IDispatcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDispatcherTransactorRaw struct {
	Contract *IDispatcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDispatcher creates a new instance of IDispatcher, bound to a specific deployed contract.
func NewIDispatcher(address common.Address, backend bind.ContractBackend) (*IDispatcher, error) {
	contract, err := bindIDispatcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDispatcher{IDispatcherCaller: IDispatcherCaller{contract: contract}, IDispatcherTransactor: IDispatcherTransactor{contract: contract}, IDispatcherFilterer: IDispatcherFilterer{contract: contract}}, nil
}

// NewIDispatcherCaller creates a new read-only instance of IDispatcher, bound to a specific deployed contract.
func NewIDispatcherCaller(address common.Address, caller bind.ContractCaller) (*IDispatcherCaller, error) {
	contract, err := bindIDispatcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDispatcherCaller{contract: contract}, nil
}

// NewIDispatcherTransactor creates a new write-only instance of IDispatcher, bound to a specific deployed contract.
func NewIDispatcherTransactor(address common.Address, transactor bind.ContractTransactor) (*IDispatcherTransactor, error) {
	contract, err := bindIDispatcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDispatcherTransactor{contract: contract}, nil
}

// NewIDispatcherFilterer creates a new log filterer instance of IDispatcher, bound to a specific deployed contract.
func NewIDispatcherFilterer(address common.Address, filterer bind.ContractFilterer) (*IDispatcherFilterer, error) {
	contract, err := bindIDispatcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDispatcherFilterer{contract: contract}, nil
}

// bindIDispatcher binds a generic wrapper to an already deployed contract.
func bindIDispatcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDispatcherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDispatcher *IDispatcherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDispatcher.Contract.IDispatcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDispatcher *IDispatcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDispatcher.Contract.IDispatcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDispatcher *IDispatcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDispatcher.Contract.IDispatcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDispatcher *IDispatcherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDispatcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDispatcher *IDispatcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDispatcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDispatcher *IDispatcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDispatcher.Contract.contract.Transact(opts, method, params...)
}

// AreAddressesImported is a free data retrieval call binding the contract method 0x9f42102f.
//
// Solidity: function areAddressesImported(bytes32[] names, address[] destinations) view returns(bool)
func (_IDispatcher *IDispatcherCaller) AreAddressesImported(opts *bind.CallOpts, names [][32]byte, destinations []common.Address) (bool, error) {
	var out []interface{}
	err := _IDispatcher.contract.Call(opts, &out, "areAddressesImported", names, destinations)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreAddressesImported is a free data retrieval call binding the contract method 0x9f42102f.
//
// Solidity: function areAddressesImported(bytes32[] names, address[] destinations) view returns(bool)
func (_IDispatcher *IDispatcherSession) AreAddressesImported(names [][32]byte, destinations []common.Address) (bool, error) {
	return _IDispatcher.Contract.AreAddressesImported(&_IDispatcher.CallOpts, names, destinations)
}

// AreAddressesImported is a free data retrieval call binding the contract method 0x9f42102f.
//
// Solidity: function areAddressesImported(bytes32[] names, address[] destinations) view returns(bool)
func (_IDispatcher *IDispatcherCallerSession) AreAddressesImported(names [][32]byte, destinations []common.Address) (bool, error) {
	return _IDispatcher.Contract.AreAddressesImported(&_IDispatcher.CallOpts, names, destinations)
}

// GetExistedAddress is a free data retrieval call binding the contract method 0x6e1425fd.
//
// Solidity: function getExistedAddress(bytes32 name, string reason) view returns(address)
func (_IDispatcher *IDispatcherCaller) GetExistedAddress(opts *bind.CallOpts, name [32]byte, reason string) (common.Address, error) {
	var out []interface{}
	err := _IDispatcher.contract.Call(opts, &out, "getExistedAddress", name, reason)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExistedAddress is a free data retrieval call binding the contract method 0x6e1425fd.
//
// Solidity: function getExistedAddress(bytes32 name, string reason) view returns(address)
func (_IDispatcher *IDispatcherSession) GetExistedAddress(name [32]byte, reason string) (common.Address, error) {
	return _IDispatcher.Contract.GetExistedAddress(&_IDispatcher.CallOpts, name, reason)
}

// GetExistedAddress is a free data retrieval call binding the contract method 0x6e1425fd.
//
// Solidity: function getExistedAddress(bytes32 name, string reason) view returns(address)
func (_IDispatcher *IDispatcherCallerSession) GetExistedAddress(name [32]byte, reason string) (common.Address, error) {
	return _IDispatcher.Contract.GetExistedAddress(&_IDispatcher.CallOpts, name, reason)
}

// TryGetAddress is a free data retrieval call binding the contract method 0x0e01b557.
//
// Solidity: function tryGetAddress(bytes32 name) view returns(address)
func (_IDispatcher *IDispatcherCaller) TryGetAddress(opts *bind.CallOpts, name [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IDispatcher.contract.Call(opts, &out, "tryGetAddress", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TryGetAddress is a free data retrieval call binding the contract method 0x0e01b557.
//
// Solidity: function tryGetAddress(bytes32 name) view returns(address)
func (_IDispatcher *IDispatcherSession) TryGetAddress(name [32]byte) (common.Address, error) {
	return _IDispatcher.Contract.TryGetAddress(&_IDispatcher.CallOpts, name)
}

// TryGetAddress is a free data retrieval call binding the contract method 0x0e01b557.
//
// Solidity: function tryGetAddress(bytes32 name) view returns(address)
func (_IDispatcher *IDispatcherCallerSession) TryGetAddress(name [32]byte) (common.Address, error) {
	return _IDispatcher.Contract.TryGetAddress(&_IDispatcher.CallOpts, name)
}

// ImportAddress is a paid mutator transaction binding the contract method 0xcfcd591c.
//
// Solidity: function importAddress(bytes32 name, address destination) returns()
func (_IDispatcher *IDispatcherTransactor) ImportAddress(opts *bind.TransactOpts, name [32]byte, destination common.Address) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "importAddress", name, destination)
}

// ImportAddress is a paid mutator transaction binding the contract method 0xcfcd591c.
//
// Solidity: function importAddress(bytes32 name, address destination) returns()
func (_IDispatcher *IDispatcherSession) ImportAddress(name [32]byte, destination common.Address) (*types.Transaction, error) {
	return _IDispatcher.Contract.ImportAddress(&_IDispatcher.TransactOpts, name, destination)
}

// ImportAddress is a paid mutator transaction binding the contract method 0xcfcd591c.
//
// Solidity: function importAddress(bytes32 name, address destination) returns()
func (_IDispatcher *IDispatcherTransactorSession) ImportAddress(name [32]byte, destination common.Address) (*types.Transaction, error) {
	return _IDispatcher.Contract.ImportAddress(&_IDispatcher.TransactOpts, name, destination)
}

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

// IWorkerStorageMetaData contains all meta data concerning the IWorkerStorage contract.
var IWorkerStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"graduatedAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"birthAt\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"collegeCode\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structWorkerDefine.Worker\",\"name\":\"worker\",\"type\":\"tuple\"}],\"name\":\"createWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getWorkerBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"graduatedAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"birthAt\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"collegeCode\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structWorkerDefine.Worker\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"checkWorkerExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IWorkerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use IWorkerStorageMetaData.ABI instead.
var IWorkerStorageABI = IWorkerStorageMetaData.ABI

// IWorkerStorage is an auto generated Go binding around an Ethereum contract.
type IWorkerStorage struct {
	IWorkerStorageCaller     // Read-only binding to the contract
	IWorkerStorageTransactor // Write-only binding to the contract
	IWorkerStorageFilterer   // Log filterer for contract events
}

// IWorkerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWorkerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWorkerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWorkerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWorkerStorageSession struct {
	Contract     *IWorkerStorage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWorkerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWorkerStorageCallerSession struct {
	Contract *IWorkerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IWorkerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWorkerStorageTransactorSession struct {
	Contract     *IWorkerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IWorkerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWorkerStorageRaw struct {
	Contract *IWorkerStorage // Generic contract binding to access the raw methods on
}

// IWorkerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWorkerStorageCallerRaw struct {
	Contract *IWorkerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// IWorkerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWorkerStorageTransactorRaw struct {
	Contract *IWorkerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWorkerStorage creates a new instance of IWorkerStorage, bound to a specific deployed contract.
func NewIWorkerStorage(address common.Address, backend bind.ContractBackend) (*IWorkerStorage, error) {
	contract, err := bindIWorkerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWorkerStorage{IWorkerStorageCaller: IWorkerStorageCaller{contract: contract}, IWorkerStorageTransactor: IWorkerStorageTransactor{contract: contract}, IWorkerStorageFilterer: IWorkerStorageFilterer{contract: contract}}, nil
}

// NewIWorkerStorageCaller creates a new read-only instance of IWorkerStorage, bound to a specific deployed contract.
func NewIWorkerStorageCaller(address common.Address, caller bind.ContractCaller) (*IWorkerStorageCaller, error) {
	contract, err := bindIWorkerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWorkerStorageCaller{contract: contract}, nil
}

// NewIWorkerStorageTransactor creates a new write-only instance of IWorkerStorage, bound to a specific deployed contract.
func NewIWorkerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*IWorkerStorageTransactor, error) {
	contract, err := bindIWorkerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWorkerStorageTransactor{contract: contract}, nil
}

// NewIWorkerStorageFilterer creates a new log filterer instance of IWorkerStorage, bound to a specific deployed contract.
func NewIWorkerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*IWorkerStorageFilterer, error) {
	contract, err := bindIWorkerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWorkerStorageFilterer{contract: contract}, nil
}

// bindIWorkerStorage binds a generic wrapper to an already deployed contract.
func bindIWorkerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWorkerStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWorkerStorage *IWorkerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWorkerStorage.Contract.IWorkerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWorkerStorage *IWorkerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerStorage.Contract.IWorkerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWorkerStorage *IWorkerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWorkerStorage.Contract.IWorkerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWorkerStorage *IWorkerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWorkerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWorkerStorage *IWorkerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWorkerStorage *IWorkerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWorkerStorage.Contract.contract.Transact(opts, method, params...)
}

// CheckWorkerExist is a free data retrieval call binding the contract method 0x8e3f594b.
//
// Solidity: function checkWorkerExist(bytes32 securityNo) view returns(bool)
func (_IWorkerStorage *IWorkerStorageCaller) CheckWorkerExist(opts *bind.CallOpts, securityNo [32]byte) (bool, error) {
	var out []interface{}
	err := _IWorkerStorage.contract.Call(opts, &out, "checkWorkerExist", securityNo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckWorkerExist is a free data retrieval call binding the contract method 0x8e3f594b.
//
// Solidity: function checkWorkerExist(bytes32 securityNo) view returns(bool)
func (_IWorkerStorage *IWorkerStorageSession) CheckWorkerExist(securityNo [32]byte) (bool, error) {
	return _IWorkerStorage.Contract.CheckWorkerExist(&_IWorkerStorage.CallOpts, securityNo)
}

// CheckWorkerExist is a free data retrieval call binding the contract method 0x8e3f594b.
//
// Solidity: function checkWorkerExist(bytes32 securityNo) view returns(bool)
func (_IWorkerStorage *IWorkerStorageCallerSession) CheckWorkerExist(securityNo [32]byte) (bool, error) {
	return _IWorkerStorage.Contract.CheckWorkerExist(&_IWorkerStorage.CallOpts, securityNo)
}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_IWorkerStorage *IWorkerStorageCaller) GetWorkerBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) (WorkerDefineWorker, error) {
	var out []interface{}
	err := _IWorkerStorage.contract.Call(opts, &out, "getWorkerBySecurityNo", securityNo)

	if err != nil {
		return *new(WorkerDefineWorker), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkerDefineWorker)).(*WorkerDefineWorker)

	return out0, err

}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_IWorkerStorage *IWorkerStorageSession) GetWorkerBySecurityNo(securityNo [32]byte) (WorkerDefineWorker, error) {
	return _IWorkerStorage.Contract.GetWorkerBySecurityNo(&_IWorkerStorage.CallOpts, securityNo)
}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_IWorkerStorage *IWorkerStorageCallerSession) GetWorkerBySecurityNo(securityNo [32]byte) (WorkerDefineWorker, error) {
	return _IWorkerStorage.Contract.GetWorkerBySecurityNo(&_IWorkerStorage.CallOpts, securityNo)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_IWorkerStorage *IWorkerStorageTransactor) CreateWorker(opts *bind.TransactOpts, worker WorkerDefineWorker) (*types.Transaction, error) {
	return _IWorkerStorage.contract.Transact(opts, "createWorker", worker)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_IWorkerStorage *IWorkerStorageSession) CreateWorker(worker WorkerDefineWorker) (*types.Transaction, error) {
	return _IWorkerStorage.Contract.CreateWorker(&_IWorkerStorage.TransactOpts, worker)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_IWorkerStorage *IWorkerStorageTransactorSession) CreateWorker(worker WorkerDefineWorker) (*types.Transaction, error) {
	return _IWorkerStorage.Contract.CreateWorker(&_IWorkerStorage.TransactOpts, worker)
}

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


// WorkerStorageMetaData contains all meta data concerning the WorkerStorage contract.
var WorkerStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dispatcher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"graduatedAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"birthAt\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"collegeCode\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structWorkerDefine.Worker\",\"name\":\"worker\",\"type\":\"tuple\"}],\"name\":\"createWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getWorkerBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"graduatedAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"birthAt\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"collegeCode\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structWorkerDefine.Worker\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"checkWorkerExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WorkerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use WorkerStorageMetaData.ABI instead.
var WorkerStorageABI = WorkerStorageMetaData.ABI

// WorkerStorage is an auto generated Go binding around an Ethereum contract.
type WorkerStorage struct {
	WorkerStorageCaller     // Read-only binding to the contract
	WorkerStorageTransactor // Write-only binding to the contract
	WorkerStorageFilterer   // Log filterer for contract events
}

// WorkerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorkerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorkerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorkerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorkerStorageSession struct {
	Contract     *WorkerStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorkerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorkerStorageCallerSession struct {
	Contract *WorkerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WorkerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorkerStorageTransactorSession struct {
	Contract     *WorkerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WorkerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorkerStorageRaw struct {
	Contract *WorkerStorage // Generic contract binding to access the raw methods on
}

// WorkerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorkerStorageCallerRaw struct {
	Contract *WorkerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// WorkerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorkerStorageTransactorRaw struct {
	Contract *WorkerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorkerStorage creates a new instance of WorkerStorage, bound to a specific deployed contract.
func NewWorkerStorage(address common.Address, backend bind.ContractBackend) (*WorkerStorage, error) {
	contract, err := bindWorkerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkerStorage{WorkerStorageCaller: WorkerStorageCaller{contract: contract}, WorkerStorageTransactor: WorkerStorageTransactor{contract: contract}, WorkerStorageFilterer: WorkerStorageFilterer{contract: contract}}, nil
}

// NewWorkerStorageCaller creates a new read-only instance of WorkerStorage, bound to a specific deployed contract.
func NewWorkerStorageCaller(address common.Address, caller bind.ContractCaller) (*WorkerStorageCaller, error) {
	contract, err := bindWorkerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerStorageCaller{contract: contract}, nil
}

// NewWorkerStorageTransactor creates a new write-only instance of WorkerStorage, bound to a specific deployed contract.
func NewWorkerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkerStorageTransactor, error) {
	contract, err := bindWorkerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerStorageTransactor{contract: contract}, nil
}

// NewWorkerStorageFilterer creates a new log filterer instance of WorkerStorage, bound to a specific deployed contract.
func NewWorkerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkerStorageFilterer, error) {
	contract, err := bindWorkerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkerStorageFilterer{contract: contract}, nil
}

// bindWorkerStorage binds a generic wrapper to an already deployed contract.
func bindWorkerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WorkerStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerStorage *WorkerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerStorage.Contract.WorkerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerStorage *WorkerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerStorage.Contract.WorkerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerStorage *WorkerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerStorage.Contract.WorkerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerStorage *WorkerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerStorage *WorkerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerStorage *WorkerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerStorage.Contract.contract.Transact(opts, method, params...)
}

// CheckWorkerExist is a free data retrieval call binding the contract method 0x8e3f594b.
//
// Solidity: function checkWorkerExist(bytes32 securityNo) view returns(bool)
func (_WorkerStorage *WorkerStorageCaller) CheckWorkerExist(opts *bind.CallOpts, securityNo [32]byte) (bool, error) {
	var out []interface{}
	err := _WorkerStorage.contract.Call(opts, &out, "checkWorkerExist", securityNo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckWorkerExist is a free data retrieval call binding the contract method 0x8e3f594b.
//
// Solidity: function checkWorkerExist(bytes32 securityNo) view returns(bool)
func (_WorkerStorage *WorkerStorageSession) CheckWorkerExist(securityNo [32]byte) (bool, error) {
	return _WorkerStorage.Contract.CheckWorkerExist(&_WorkerStorage.CallOpts, securityNo)
}

// CheckWorkerExist is a free data retrieval call binding the contract method 0x8e3f594b.
//
// Solidity: function checkWorkerExist(bytes32 securityNo) view returns(bool)
func (_WorkerStorage *WorkerStorageCallerSession) CheckWorkerExist(securityNo [32]byte) (bool, error) {
	return _WorkerStorage.Contract.CheckWorkerExist(&_WorkerStorage.CallOpts, securityNo)
}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_WorkerStorage *WorkerStorageCaller) GetWorkerBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) (WorkerDefineWorker, error) {
	var out []interface{}
	err := _WorkerStorage.contract.Call(opts, &out, "getWorkerBySecurityNo", securityNo)

	if err != nil {
		return *new(WorkerDefineWorker), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkerDefineWorker)).(*WorkerDefineWorker)

	return out0, err

}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_WorkerStorage *WorkerStorageSession) GetWorkerBySecurityNo(securityNo [32]byte) (WorkerDefineWorker, error) {
	return _WorkerStorage.Contract.GetWorkerBySecurityNo(&_WorkerStorage.CallOpts, securityNo)
}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_WorkerStorage *WorkerStorageCallerSession) GetWorkerBySecurityNo(securityNo [32]byte) (WorkerDefineWorker, error) {
	return _WorkerStorage.Contract.GetWorkerBySecurityNo(&_WorkerStorage.CallOpts, securityNo)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_WorkerStorage *WorkerStorageTransactor) CreateWorker(opts *bind.TransactOpts, worker WorkerDefineWorker) (*types.Transaction, error) {
	return _WorkerStorage.contract.Transact(opts, "createWorker", worker)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_WorkerStorage *WorkerStorageSession) CreateWorker(worker WorkerDefineWorker) (*types.Transaction, error) {
	return _WorkerStorage.Contract.CreateWorker(&_WorkerStorage.TransactOpts, worker)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_WorkerStorage *WorkerStorageTransactorSession) CreateWorker(worker WorkerDefineWorker) (*types.Transaction, error) {
	return _WorkerStorage.Contract.CreateWorker(&_WorkerStorage.TransactOpts, worker)
}

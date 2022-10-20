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

// IWorkerServiceMetaData contains all meta data concerning the IWorkerService contract.
var IWorkerServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"graduatedAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"birthAt\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"collegeCode\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structWorkerDefine.Worker\",\"name\":\"worker\",\"type\":\"tuple\"}],\"name\":\"createWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getWorkerBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"graduatedAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"birthAt\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"collegeCode\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structWorkerDefine.Worker\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IWorkerServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use IWorkerServiceMetaData.ABI instead.
var IWorkerServiceABI = IWorkerServiceMetaData.ABI

// IWorkerService is an auto generated Go binding around an Ethereum contract.
type IWorkerService struct {
	IWorkerServiceCaller     // Read-only binding to the contract
	IWorkerServiceTransactor // Write-only binding to the contract
	IWorkerServiceFilterer   // Log filterer for contract events
}

// IWorkerServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWorkerServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkerServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWorkerServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkerServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWorkerServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkerServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWorkerServiceSession struct {
	Contract     *IWorkerService   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWorkerServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWorkerServiceCallerSession struct {
	Contract *IWorkerServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IWorkerServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWorkerServiceTransactorSession struct {
	Contract     *IWorkerServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IWorkerServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWorkerServiceRaw struct {
	Contract *IWorkerService // Generic contract binding to access the raw methods on
}

// IWorkerServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWorkerServiceCallerRaw struct {
	Contract *IWorkerServiceCaller // Generic read-only contract binding to access the raw methods on
}

// IWorkerServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWorkerServiceTransactorRaw struct {
	Contract *IWorkerServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWorkerService creates a new instance of IWorkerService, bound to a specific deployed contract.
func NewIWorkerService(address common.Address, backend bind.ContractBackend) (*IWorkerService, error) {
	contract, err := bindIWorkerService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWorkerService{IWorkerServiceCaller: IWorkerServiceCaller{contract: contract}, IWorkerServiceTransactor: IWorkerServiceTransactor{contract: contract}, IWorkerServiceFilterer: IWorkerServiceFilterer{contract: contract}}, nil
}

// NewIWorkerServiceCaller creates a new read-only instance of IWorkerService, bound to a specific deployed contract.
func NewIWorkerServiceCaller(address common.Address, caller bind.ContractCaller) (*IWorkerServiceCaller, error) {
	contract, err := bindIWorkerService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWorkerServiceCaller{contract: contract}, nil
}

// NewIWorkerServiceTransactor creates a new write-only instance of IWorkerService, bound to a specific deployed contract.
func NewIWorkerServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*IWorkerServiceTransactor, error) {
	contract, err := bindIWorkerService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWorkerServiceTransactor{contract: contract}, nil
}

// NewIWorkerServiceFilterer creates a new log filterer instance of IWorkerService, bound to a specific deployed contract.
func NewIWorkerServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*IWorkerServiceFilterer, error) {
	contract, err := bindIWorkerService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWorkerServiceFilterer{contract: contract}, nil
}

// bindIWorkerService binds a generic wrapper to an already deployed contract.
func bindIWorkerService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWorkerServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWorkerService *IWorkerServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWorkerService.Contract.IWorkerServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWorkerService *IWorkerServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerService.Contract.IWorkerServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWorkerService *IWorkerServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWorkerService.Contract.IWorkerServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWorkerService *IWorkerServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWorkerService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWorkerService *IWorkerServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWorkerService *IWorkerServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWorkerService.Contract.contract.Transact(opts, method, params...)
}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_IWorkerService *IWorkerServiceCaller) GetWorkerBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) (WorkerDefineWorker, error) {
	var out []interface{}
	err := _IWorkerService.contract.Call(opts, &out, "getWorkerBySecurityNo", securityNo)

	if err != nil {
		return *new(WorkerDefineWorker), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkerDefineWorker)).(*WorkerDefineWorker)

	return out0, err

}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_IWorkerService *IWorkerServiceSession) GetWorkerBySecurityNo(securityNo [32]byte) (WorkerDefineWorker, error) {
	return _IWorkerService.Contract.GetWorkerBySecurityNo(&_IWorkerService.CallOpts, securityNo)
}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_IWorkerService *IWorkerServiceCallerSession) GetWorkerBySecurityNo(securityNo [32]byte) (WorkerDefineWorker, error) {
	return _IWorkerService.Contract.GetWorkerBySecurityNo(&_IWorkerService.CallOpts, securityNo)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_IWorkerService *IWorkerServiceTransactor) CreateWorker(opts *bind.TransactOpts, worker WorkerDefineWorker) (*types.Transaction, error) {
	return _IWorkerService.contract.Transact(opts, "createWorker", worker)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_IWorkerService *IWorkerServiceSession) CreateWorker(worker WorkerDefineWorker) (*types.Transaction, error) {
	return _IWorkerService.Contract.CreateWorker(&_IWorkerService.TransactOpts, worker)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_IWorkerService *IWorkerServiceTransactorSession) CreateWorker(worker WorkerDefineWorker) (*types.Transaction, error) {
	return _IWorkerService.Contract.CreateWorker(&_IWorkerService.TransactOpts, worker)
}

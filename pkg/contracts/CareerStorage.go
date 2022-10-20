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


// CareerStorageMetaData contains all meta data concerning the CareerStorage contract.
var CareerStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dispatcher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"internalType\":\"structWorkExperienceDefine.WorkExperience\",\"name\":\"newExperience\",\"type\":\"tuple\"}],\"name\":\"createCareer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"endYear\",\"type\":\"uint16\"}],\"name\":\"finishLastCareer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getLastCareer\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"internalType\":\"structWorkExperienceDefine.WorkExperience\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getCareerBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"internalType\":\"structWorkExperienceDefine.WorkExperience[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CareerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use CareerStorageMetaData.ABI instead.
var CareerStorageABI = CareerStorageMetaData.ABI

// CareerStorage is an auto generated Go binding around an Ethereum contract.
type CareerStorage struct {
	CareerStorageCaller     // Read-only binding to the contract
	CareerStorageTransactor // Write-only binding to the contract
	CareerStorageFilterer   // Log filterer for contract events
}

// CareerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type CareerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CareerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CareerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CareerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CareerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CareerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CareerStorageSession struct {
	Contract     *CareerStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CareerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CareerStorageCallerSession struct {
	Contract *CareerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CareerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CareerStorageTransactorSession struct {
	Contract     *CareerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CareerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type CareerStorageRaw struct {
	Contract *CareerStorage // Generic contract binding to access the raw methods on
}

// CareerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CareerStorageCallerRaw struct {
	Contract *CareerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// CareerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CareerStorageTransactorRaw struct {
	Contract *CareerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCareerStorage creates a new instance of CareerStorage, bound to a specific deployed contract.
func NewCareerStorage(address common.Address, backend bind.ContractBackend) (*CareerStorage, error) {
	contract, err := bindCareerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CareerStorage{CareerStorageCaller: CareerStorageCaller{contract: contract}, CareerStorageTransactor: CareerStorageTransactor{contract: contract}, CareerStorageFilterer: CareerStorageFilterer{contract: contract}}, nil
}

// NewCareerStorageCaller creates a new read-only instance of CareerStorage, bound to a specific deployed contract.
func NewCareerStorageCaller(address common.Address, caller bind.ContractCaller) (*CareerStorageCaller, error) {
	contract, err := bindCareerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CareerStorageCaller{contract: contract}, nil
}

// NewCareerStorageTransactor creates a new write-only instance of CareerStorage, bound to a specific deployed contract.
func NewCareerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*CareerStorageTransactor, error) {
	contract, err := bindCareerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CareerStorageTransactor{contract: contract}, nil
}

// NewCareerStorageFilterer creates a new log filterer instance of CareerStorage, bound to a specific deployed contract.
func NewCareerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*CareerStorageFilterer, error) {
	contract, err := bindCareerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CareerStorageFilterer{contract: contract}, nil
}

// bindCareerStorage binds a generic wrapper to an already deployed contract.
func bindCareerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CareerStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CareerStorage *CareerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CareerStorage.Contract.CareerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CareerStorage *CareerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CareerStorage.Contract.CareerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CareerStorage *CareerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CareerStorage.Contract.CareerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CareerStorage *CareerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CareerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CareerStorage *CareerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CareerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CareerStorage *CareerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CareerStorage.Contract.contract.Transact(opts, method, params...)
}

// GetCareerBySecurityNo is a free data retrieval call binding the contract method 0x5fa3636d.
//
// Solidity: function getCareerBySecurityNo(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32)[])
func (_CareerStorage *CareerStorageCaller) GetCareerBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) ([]WorkExperienceDefineWorkExperience, error) {
	var out []interface{}
	err := _CareerStorage.contract.Call(opts, &out, "getCareerBySecurityNo", securityNo)

	if err != nil {
		return *new([]WorkExperienceDefineWorkExperience), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkExperienceDefineWorkExperience)).(*[]WorkExperienceDefineWorkExperience)

	return out0, err

}

// GetCareerBySecurityNo is a free data retrieval call binding the contract method 0x5fa3636d.
//
// Solidity: function getCareerBySecurityNo(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32)[])
func (_CareerStorage *CareerStorageSession) GetCareerBySecurityNo(securityNo [32]byte) ([]WorkExperienceDefineWorkExperience, error) {
	return _CareerStorage.Contract.GetCareerBySecurityNo(&_CareerStorage.CallOpts, securityNo)
}

// GetCareerBySecurityNo is a free data retrieval call binding the contract method 0x5fa3636d.
//
// Solidity: function getCareerBySecurityNo(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32)[])
func (_CareerStorage *CareerStorageCallerSession) GetCareerBySecurityNo(securityNo [32]byte) ([]WorkExperienceDefineWorkExperience, error) {
	return _CareerStorage.Contract.GetCareerBySecurityNo(&_CareerStorage.CallOpts, securityNo)
}

// GetLastCareer is a free data retrieval call binding the contract method 0xd3dfe65b.
//
// Solidity: function getLastCareer(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32))
func (_CareerStorage *CareerStorageCaller) GetLastCareer(opts *bind.CallOpts, securityNo [32]byte) (WorkExperienceDefineWorkExperience, error) {
	var out []interface{}
	err := _CareerStorage.contract.Call(opts, &out, "getLastCareer", securityNo)

	if err != nil {
		return *new(WorkExperienceDefineWorkExperience), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkExperienceDefineWorkExperience)).(*WorkExperienceDefineWorkExperience)

	return out0, err

}

// GetLastCareer is a free data retrieval call binding the contract method 0xd3dfe65b.
//
// Solidity: function getLastCareer(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32))
func (_CareerStorage *CareerStorageSession) GetLastCareer(securityNo [32]byte) (WorkExperienceDefineWorkExperience, error) {
	return _CareerStorage.Contract.GetLastCareer(&_CareerStorage.CallOpts, securityNo)
}

// GetLastCareer is a free data retrieval call binding the contract method 0xd3dfe65b.
//
// Solidity: function getLastCareer(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32))
func (_CareerStorage *CareerStorageCallerSession) GetLastCareer(securityNo [32]byte) (WorkExperienceDefineWorkExperience, error) {
	return _CareerStorage.Contract.GetLastCareer(&_CareerStorage.CallOpts, securityNo)
}

// CreateCareer is a paid mutator transaction binding the contract method 0x8ab95b51.
//
// Solidity: function createCareer(bytes32 securityNo, (uint16,uint16,bool,bytes32) newExperience) returns()
func (_CareerStorage *CareerStorageTransactor) CreateCareer(opts *bind.TransactOpts, securityNo [32]byte, newExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _CareerStorage.contract.Transact(opts, "createCareer", securityNo, newExperience)
}

// CreateCareer is a paid mutator transaction binding the contract method 0x8ab95b51.
//
// Solidity: function createCareer(bytes32 securityNo, (uint16,uint16,bool,bytes32) newExperience) returns()
func (_CareerStorage *CareerStorageSession) CreateCareer(securityNo [32]byte, newExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _CareerStorage.Contract.CreateCareer(&_CareerStorage.TransactOpts, securityNo, newExperience)
}

// CreateCareer is a paid mutator transaction binding the contract method 0x8ab95b51.
//
// Solidity: function createCareer(bytes32 securityNo, (uint16,uint16,bool,bytes32) newExperience) returns()
func (_CareerStorage *CareerStorageTransactorSession) CreateCareer(securityNo [32]byte, newExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _CareerStorage.Contract.CreateCareer(&_CareerStorage.TransactOpts, securityNo, newExperience)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_CareerStorage *CareerStorageTransactor) FinishLastCareer(opts *bind.TransactOpts, securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _CareerStorage.contract.Transact(opts, "finishLastCareer", securityNo, endYear)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_CareerStorage *CareerStorageSession) FinishLastCareer(securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _CareerStorage.Contract.FinishLastCareer(&_CareerStorage.TransactOpts, securityNo, endYear)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_CareerStorage *CareerStorageTransactorSession) FinishLastCareer(securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _CareerStorage.Contract.FinishLastCareer(&_CareerStorage.TransactOpts, securityNo, endYear)
}

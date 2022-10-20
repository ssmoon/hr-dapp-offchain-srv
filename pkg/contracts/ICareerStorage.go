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


// ICareerStorageMetaData contains all meta data concerning the ICareerStorage contract.
var ICareerStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"internalType\":\"structWorkExperienceDefine.WorkExperience\",\"name\":\"newExperience\",\"type\":\"tuple\"}],\"name\":\"createCareer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"endYear\",\"type\":\"uint16\"}],\"name\":\"finishLastCareer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getLastCareer\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"internalType\":\"structWorkExperienceDefine.WorkExperience\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getCareerBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"internalType\":\"structWorkExperienceDefine.WorkExperience[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ICareerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use ICareerStorageMetaData.ABI instead.
var ICareerStorageABI = ICareerStorageMetaData.ABI

// ICareerStorage is an auto generated Go binding around an Ethereum contract.
type ICareerStorage struct {
	ICareerStorageCaller     // Read-only binding to the contract
	ICareerStorageTransactor // Write-only binding to the contract
	ICareerStorageFilterer   // Log filterer for contract events
}

// ICareerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICareerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICareerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICareerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICareerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICareerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICareerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICareerStorageSession struct {
	Contract     *ICareerStorage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICareerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICareerStorageCallerSession struct {
	Contract *ICareerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ICareerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICareerStorageTransactorSession struct {
	Contract     *ICareerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ICareerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICareerStorageRaw struct {
	Contract *ICareerStorage // Generic contract binding to access the raw methods on
}

// ICareerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICareerStorageCallerRaw struct {
	Contract *ICareerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ICareerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICareerStorageTransactorRaw struct {
	Contract *ICareerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICareerStorage creates a new instance of ICareerStorage, bound to a specific deployed contract.
func NewICareerStorage(address common.Address, backend bind.ContractBackend) (*ICareerStorage, error) {
	contract, err := bindICareerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICareerStorage{ICareerStorageCaller: ICareerStorageCaller{contract: contract}, ICareerStorageTransactor: ICareerStorageTransactor{contract: contract}, ICareerStorageFilterer: ICareerStorageFilterer{contract: contract}}, nil
}

// NewICareerStorageCaller creates a new read-only instance of ICareerStorage, bound to a specific deployed contract.
func NewICareerStorageCaller(address common.Address, caller bind.ContractCaller) (*ICareerStorageCaller, error) {
	contract, err := bindICareerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICareerStorageCaller{contract: contract}, nil
}

// NewICareerStorageTransactor creates a new write-only instance of ICareerStorage, bound to a specific deployed contract.
func NewICareerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ICareerStorageTransactor, error) {
	contract, err := bindICareerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICareerStorageTransactor{contract: contract}, nil
}

// NewICareerStorageFilterer creates a new log filterer instance of ICareerStorage, bound to a specific deployed contract.
func NewICareerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ICareerStorageFilterer, error) {
	contract, err := bindICareerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICareerStorageFilterer{contract: contract}, nil
}

// bindICareerStorage binds a generic wrapper to an already deployed contract.
func bindICareerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICareerStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICareerStorage *ICareerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICareerStorage.Contract.ICareerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICareerStorage *ICareerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICareerStorage.Contract.ICareerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICareerStorage *ICareerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICareerStorage.Contract.ICareerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICareerStorage *ICareerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICareerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICareerStorage *ICareerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICareerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICareerStorage *ICareerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICareerStorage.Contract.contract.Transact(opts, method, params...)
}

// GetCareerBySecurityNo is a free data retrieval call binding the contract method 0x5fa3636d.
//
// Solidity: function getCareerBySecurityNo(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32)[])
func (_ICareerStorage *ICareerStorageCaller) GetCareerBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) ([]WorkExperienceDefineWorkExperience, error) {
	var out []interface{}
	err := _ICareerStorage.contract.Call(opts, &out, "getCareerBySecurityNo", securityNo)

	if err != nil {
		return *new([]WorkExperienceDefineWorkExperience), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkExperienceDefineWorkExperience)).(*[]WorkExperienceDefineWorkExperience)

	return out0, err

}

// GetCareerBySecurityNo is a free data retrieval call binding the contract method 0x5fa3636d.
//
// Solidity: function getCareerBySecurityNo(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32)[])
func (_ICareerStorage *ICareerStorageSession) GetCareerBySecurityNo(securityNo [32]byte) ([]WorkExperienceDefineWorkExperience, error) {
	return _ICareerStorage.Contract.GetCareerBySecurityNo(&_ICareerStorage.CallOpts, securityNo)
}

// GetCareerBySecurityNo is a free data retrieval call binding the contract method 0x5fa3636d.
//
// Solidity: function getCareerBySecurityNo(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32)[])
func (_ICareerStorage *ICareerStorageCallerSession) GetCareerBySecurityNo(securityNo [32]byte) ([]WorkExperienceDefineWorkExperience, error) {
	return _ICareerStorage.Contract.GetCareerBySecurityNo(&_ICareerStorage.CallOpts, securityNo)
}

// GetLastCareer is a free data retrieval call binding the contract method 0xd3dfe65b.
//
// Solidity: function getLastCareer(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32))
func (_ICareerStorage *ICareerStorageCaller) GetLastCareer(opts *bind.CallOpts, securityNo [32]byte) (WorkExperienceDefineWorkExperience, error) {
	var out []interface{}
	err := _ICareerStorage.contract.Call(opts, &out, "getLastCareer", securityNo)

	if err != nil {
		return *new(WorkExperienceDefineWorkExperience), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkExperienceDefineWorkExperience)).(*WorkExperienceDefineWorkExperience)

	return out0, err

}

// GetLastCareer is a free data retrieval call binding the contract method 0xd3dfe65b.
//
// Solidity: function getLastCareer(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32))
func (_ICareerStorage *ICareerStorageSession) GetLastCareer(securityNo [32]byte) (WorkExperienceDefineWorkExperience, error) {
	return _ICareerStorage.Contract.GetLastCareer(&_ICareerStorage.CallOpts, securityNo)
}

// GetLastCareer is a free data retrieval call binding the contract method 0xd3dfe65b.
//
// Solidity: function getLastCareer(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32))
func (_ICareerStorage *ICareerStorageCallerSession) GetLastCareer(securityNo [32]byte) (WorkExperienceDefineWorkExperience, error) {
	return _ICareerStorage.Contract.GetLastCareer(&_ICareerStorage.CallOpts, securityNo)
}

// CreateCareer is a paid mutator transaction binding the contract method 0x8ab95b51.
//
// Solidity: function createCareer(bytes32 securityNo, (uint16,uint16,bool,bytes32) newExperience) returns()
func (_ICareerStorage *ICareerStorageTransactor) CreateCareer(opts *bind.TransactOpts, securityNo [32]byte, newExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _ICareerStorage.contract.Transact(opts, "createCareer", securityNo, newExperience)
}

// CreateCareer is a paid mutator transaction binding the contract method 0x8ab95b51.
//
// Solidity: function createCareer(bytes32 securityNo, (uint16,uint16,bool,bytes32) newExperience) returns()
func (_ICareerStorage *ICareerStorageSession) CreateCareer(securityNo [32]byte, newExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _ICareerStorage.Contract.CreateCareer(&_ICareerStorage.TransactOpts, securityNo, newExperience)
}

// CreateCareer is a paid mutator transaction binding the contract method 0x8ab95b51.
//
// Solidity: function createCareer(bytes32 securityNo, (uint16,uint16,bool,bytes32) newExperience) returns()
func (_ICareerStorage *ICareerStorageTransactorSession) CreateCareer(securityNo [32]byte, newExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _ICareerStorage.Contract.CreateCareer(&_ICareerStorage.TransactOpts, securityNo, newExperience)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_ICareerStorage *ICareerStorageTransactor) FinishLastCareer(opts *bind.TransactOpts, securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _ICareerStorage.contract.Transact(opts, "finishLastCareer", securityNo, endYear)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_ICareerStorage *ICareerStorageSession) FinishLastCareer(securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _ICareerStorage.Contract.FinishLastCareer(&_ICareerStorage.TransactOpts, securityNo, endYear)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_ICareerStorage *ICareerStorageTransactorSession) FinishLastCareer(securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _ICareerStorage.Contract.FinishLastCareer(&_ICareerStorage.TransactOpts, securityNo, endYear)
}

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


// ICareerServiceMetaData contains all meta data concerning the ICareerService contract.
var ICareerServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"internalType\":\"structWorkExperienceDefine.WorkExperience\",\"name\":\"workExperience\",\"type\":\"tuple\"}],\"name\":\"addWorkExperience\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"endYear\",\"type\":\"uint16\"}],\"name\":\"finishLastCareer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getWorkExperienceBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"internalType\":\"structWorkExperienceDefine.WorkExperience[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ICareerServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use ICareerServiceMetaData.ABI instead.
var ICareerServiceABI = ICareerServiceMetaData.ABI

// ICareerService is an auto generated Go binding around an Ethereum contract.
type ICareerService struct {
	ICareerServiceCaller     // Read-only binding to the contract
	ICareerServiceTransactor // Write-only binding to the contract
	ICareerServiceFilterer   // Log filterer for contract events
}

// ICareerServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICareerServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICareerServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICareerServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICareerServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICareerServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICareerServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICareerServiceSession struct {
	Contract     *ICareerService   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICareerServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICareerServiceCallerSession struct {
	Contract *ICareerServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ICareerServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICareerServiceTransactorSession struct {
	Contract     *ICareerServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ICareerServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICareerServiceRaw struct {
	Contract *ICareerService // Generic contract binding to access the raw methods on
}

// ICareerServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICareerServiceCallerRaw struct {
	Contract *ICareerServiceCaller // Generic read-only contract binding to access the raw methods on
}

// ICareerServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICareerServiceTransactorRaw struct {
	Contract *ICareerServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICareerService creates a new instance of ICareerService, bound to a specific deployed contract.
func NewICareerService(address common.Address, backend bind.ContractBackend) (*ICareerService, error) {
	contract, err := bindICareerService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICareerService{ICareerServiceCaller: ICareerServiceCaller{contract: contract}, ICareerServiceTransactor: ICareerServiceTransactor{contract: contract}, ICareerServiceFilterer: ICareerServiceFilterer{contract: contract}}, nil
}

// NewICareerServiceCaller creates a new read-only instance of ICareerService, bound to a specific deployed contract.
func NewICareerServiceCaller(address common.Address, caller bind.ContractCaller) (*ICareerServiceCaller, error) {
	contract, err := bindICareerService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICareerServiceCaller{contract: contract}, nil
}

// NewICareerServiceTransactor creates a new write-only instance of ICareerService, bound to a specific deployed contract.
func NewICareerServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*ICareerServiceTransactor, error) {
	contract, err := bindICareerService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICareerServiceTransactor{contract: contract}, nil
}

// NewICareerServiceFilterer creates a new log filterer instance of ICareerService, bound to a specific deployed contract.
func NewICareerServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*ICareerServiceFilterer, error) {
	contract, err := bindICareerService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICareerServiceFilterer{contract: contract}, nil
}

// bindICareerService binds a generic wrapper to an already deployed contract.
func bindICareerService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICareerServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICareerService *ICareerServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICareerService.Contract.ICareerServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICareerService *ICareerServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICareerService.Contract.ICareerServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICareerService *ICareerServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICareerService.Contract.ICareerServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICareerService *ICareerServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICareerService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICareerService *ICareerServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICareerService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICareerService *ICareerServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICareerService.Contract.contract.Transact(opts, method, params...)
}

// AddWorkExperience is a paid mutator transaction binding the contract method 0x36efe55e.
//
// Solidity: function addWorkExperience(bytes32 securityNo, (uint16,uint16,bool,bytes32) workExperience) returns()
func (_ICareerService *ICareerServiceTransactor) AddWorkExperience(opts *bind.TransactOpts, securityNo [32]byte, workExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _ICareerService.contract.Transact(opts, "addWorkExperience", securityNo, workExperience)
}

// AddWorkExperience is a paid mutator transaction binding the contract method 0x36efe55e.
//
// Solidity: function addWorkExperience(bytes32 securityNo, (uint16,uint16,bool,bytes32) workExperience) returns()
func (_ICareerService *ICareerServiceSession) AddWorkExperience(securityNo [32]byte, workExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _ICareerService.Contract.AddWorkExperience(&_ICareerService.TransactOpts, securityNo, workExperience)
}

// AddWorkExperience is a paid mutator transaction binding the contract method 0x36efe55e.
//
// Solidity: function addWorkExperience(bytes32 securityNo, (uint16,uint16,bool,bytes32) workExperience) returns()
func (_ICareerService *ICareerServiceTransactorSession) AddWorkExperience(securityNo [32]byte, workExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _ICareerService.Contract.AddWorkExperience(&_ICareerService.TransactOpts, securityNo, workExperience)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_ICareerService *ICareerServiceTransactor) FinishLastCareer(opts *bind.TransactOpts, securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _ICareerService.contract.Transact(opts, "finishLastCareer", securityNo, endYear)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_ICareerService *ICareerServiceSession) FinishLastCareer(securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _ICareerService.Contract.FinishLastCareer(&_ICareerService.TransactOpts, securityNo, endYear)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_ICareerService *ICareerServiceTransactorSession) FinishLastCareer(securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _ICareerService.Contract.FinishLastCareer(&_ICareerService.TransactOpts, securityNo, endYear)
}

// GetWorkExperienceBySecurityNo is a paid mutator transaction binding the contract method 0xabaed258.
//
// Solidity: function getWorkExperienceBySecurityNo(bytes32 securityNo) returns((uint16,uint16,bool,bytes32)[])
func (_ICareerService *ICareerServiceTransactor) GetWorkExperienceBySecurityNo(opts *bind.TransactOpts, securityNo [32]byte) (*types.Transaction, error) {
	return _ICareerService.contract.Transact(opts, "getWorkExperienceBySecurityNo", securityNo)
}

// GetWorkExperienceBySecurityNo is a paid mutator transaction binding the contract method 0xabaed258.
//
// Solidity: function getWorkExperienceBySecurityNo(bytes32 securityNo) returns((uint16,uint16,bool,bytes32)[])
func (_ICareerService *ICareerServiceSession) GetWorkExperienceBySecurityNo(securityNo [32]byte) (*types.Transaction, error) {
	return _ICareerService.Contract.GetWorkExperienceBySecurityNo(&_ICareerService.TransactOpts, securityNo)
}

// GetWorkExperienceBySecurityNo is a paid mutator transaction binding the contract method 0xabaed258.
//
// Solidity: function getWorkExperienceBySecurityNo(bytes32 securityNo) returns((uint16,uint16,bool,bytes32)[])
func (_ICareerService *ICareerServiceTransactorSession) GetWorkExperienceBySecurityNo(securityNo [32]byte) (*types.Transaction, error) {
	return _ICareerService.Contract.GetWorkExperienceBySecurityNo(&_ICareerService.TransactOpts, securityNo)
}

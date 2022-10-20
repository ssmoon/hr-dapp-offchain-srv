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

// CertificateStorageMetaData contains all meta data concerning the CertificateStorage contract.
var CertificateStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dispatcher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"certCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"acquiredAt\",\"type\":\"uint16\"}],\"internalType\":\"structCertificateDefine.Certificate\",\"name\":\"certifcate\",\"type\":\"tuple\"}],\"name\":\"createCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getCertificateBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"certCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"acquiredAt\",\"type\":\"uint16\"}],\"internalType\":\"structCertificateDefine.Certificate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"certificateCode\",\"type\":\"bytes32\"}],\"name\":\"checkUserHasCertificate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CertificateStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use CertificateStorageMetaData.ABI instead.
var CertificateStorageABI = CertificateStorageMetaData.ABI

// CertificateStorage is an auto generated Go binding around an Ethereum contract.
type CertificateStorage struct {
	CertificateStorageCaller     // Read-only binding to the contract
	CertificateStorageTransactor // Write-only binding to the contract
	CertificateStorageFilterer   // Log filterer for contract events
}

// CertificateStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertificateStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertificateStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertificateStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertificateStorageSession struct {
	Contract     *CertificateStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CertificateStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertificateStorageCallerSession struct {
	Contract *CertificateStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CertificateStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertificateStorageTransactorSession struct {
	Contract     *CertificateStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CertificateStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertificateStorageRaw struct {
	Contract *CertificateStorage // Generic contract binding to access the raw methods on
}

// CertificateStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertificateStorageCallerRaw struct {
	Contract *CertificateStorageCaller // Generic read-only contract binding to access the raw methods on
}

// CertificateStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertificateStorageTransactorRaw struct {
	Contract *CertificateStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertificateStorage creates a new instance of CertificateStorage, bound to a specific deployed contract.
func NewCertificateStorage(address common.Address, backend bind.ContractBackend) (*CertificateStorage, error) {
	contract, err := bindCertificateStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CertificateStorage{CertificateStorageCaller: CertificateStorageCaller{contract: contract}, CertificateStorageTransactor: CertificateStorageTransactor{contract: contract}, CertificateStorageFilterer: CertificateStorageFilterer{contract: contract}}, nil
}

// NewCertificateStorageCaller creates a new read-only instance of CertificateStorage, bound to a specific deployed contract.
func NewCertificateStorageCaller(address common.Address, caller bind.ContractCaller) (*CertificateStorageCaller, error) {
	contract, err := bindCertificateStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateStorageCaller{contract: contract}, nil
}

// NewCertificateStorageTransactor creates a new write-only instance of CertificateStorage, bound to a specific deployed contract.
func NewCertificateStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*CertificateStorageTransactor, error) {
	contract, err := bindCertificateStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateStorageTransactor{contract: contract}, nil
}

// NewCertificateStorageFilterer creates a new log filterer instance of CertificateStorage, bound to a specific deployed contract.
func NewCertificateStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*CertificateStorageFilterer, error) {
	contract, err := bindCertificateStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertificateStorageFilterer{contract: contract}, nil
}

// bindCertificateStorage binds a generic wrapper to an already deployed contract.
func bindCertificateStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CertificateStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificateStorage *CertificateStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CertificateStorage.Contract.CertificateStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificateStorage *CertificateStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificateStorage.Contract.CertificateStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificateStorage *CertificateStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertificateStorage.Contract.CertificateStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificateStorage *CertificateStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CertificateStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificateStorage *CertificateStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificateStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificateStorage *CertificateStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertificateStorage.Contract.contract.Transact(opts, method, params...)
}

// CheckUserHasCertificate is a free data retrieval call binding the contract method 0xc4aa76c5.
//
// Solidity: function checkUserHasCertificate(bytes32 securityNo, bytes32 certificateCode) view returns(bool)
func (_CertificateStorage *CertificateStorageCaller) CheckUserHasCertificate(opts *bind.CallOpts, securityNo [32]byte, certificateCode [32]byte) (bool, error) {
	var out []interface{}
	err := _CertificateStorage.contract.Call(opts, &out, "checkUserHasCertificate", securityNo, certificateCode)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckUserHasCertificate is a free data retrieval call binding the contract method 0xc4aa76c5.
//
// Solidity: function checkUserHasCertificate(bytes32 securityNo, bytes32 certificateCode) view returns(bool)
func (_CertificateStorage *CertificateStorageSession) CheckUserHasCertificate(securityNo [32]byte, certificateCode [32]byte) (bool, error) {
	return _CertificateStorage.Contract.CheckUserHasCertificate(&_CertificateStorage.CallOpts, securityNo, certificateCode)
}

// CheckUserHasCertificate is a free data retrieval call binding the contract method 0xc4aa76c5.
//
// Solidity: function checkUserHasCertificate(bytes32 securityNo, bytes32 certificateCode) view returns(bool)
func (_CertificateStorage *CertificateStorageCallerSession) CheckUserHasCertificate(securityNo [32]byte, certificateCode [32]byte) (bool, error) {
	return _CertificateStorage.Contract.CheckUserHasCertificate(&_CertificateStorage.CallOpts, securityNo, certificateCode)
}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_CertificateStorage *CertificateStorageCaller) GetCertificateBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	var out []interface{}
	err := _CertificateStorage.contract.Call(opts, &out, "getCertificateBySecurityNo", securityNo)

	if err != nil {
		return *new([]CertificateDefineCertificate), err
	}

	out0 := *abi.ConvertType(out[0], new([]CertificateDefineCertificate)).(*[]CertificateDefineCertificate)

	return out0, err

}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_CertificateStorage *CertificateStorageSession) GetCertificateBySecurityNo(securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	return _CertificateStorage.Contract.GetCertificateBySecurityNo(&_CertificateStorage.CallOpts, securityNo)
}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_CertificateStorage *CertificateStorageCallerSession) GetCertificateBySecurityNo(securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	return _CertificateStorage.Contract.GetCertificateBySecurityNo(&_CertificateStorage.CallOpts, securityNo)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_CertificateStorage *CertificateStorageTransactor) CreateCertificate(opts *bind.TransactOpts, securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _CertificateStorage.contract.Transact(opts, "createCertificate", securityNo, certifcate)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_CertificateStorage *CertificateStorageSession) CreateCertificate(securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _CertificateStorage.Contract.CreateCertificate(&_CertificateStorage.TransactOpts, securityNo, certifcate)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_CertificateStorage *CertificateStorageTransactorSession) CreateCertificate(securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _CertificateStorage.Contract.CreateCertificate(&_CertificateStorage.TransactOpts, securityNo, certifcate)
}

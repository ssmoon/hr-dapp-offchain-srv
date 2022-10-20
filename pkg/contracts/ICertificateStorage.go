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


// ICertificateStorageMetaData contains all meta data concerning the ICertificateStorage contract.
var ICertificateStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"certCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"acquiredAt\",\"type\":\"uint16\"}],\"internalType\":\"structCertificateDefine.Certificate\",\"name\":\"certifcate\",\"type\":\"tuple\"}],\"name\":\"createCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getCertificateBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"certCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"acquiredAt\",\"type\":\"uint16\"}],\"internalType\":\"structCertificateDefine.Certificate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"certificateCode\",\"type\":\"bytes32\"}],\"name\":\"checkUserHasCertificate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ICertificateStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use ICertificateStorageMetaData.ABI instead.
var ICertificateStorageABI = ICertificateStorageMetaData.ABI

// ICertificateStorage is an auto generated Go binding around an Ethereum contract.
type ICertificateStorage struct {
	ICertificateStorageCaller     // Read-only binding to the contract
	ICertificateStorageTransactor // Write-only binding to the contract
	ICertificateStorageFilterer   // Log filterer for contract events
}

// ICertificateStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICertificateStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICertificateStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICertificateStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICertificateStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICertificateStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICertificateStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICertificateStorageSession struct {
	Contract     *ICertificateStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ICertificateStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICertificateStorageCallerSession struct {
	Contract *ICertificateStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ICertificateStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICertificateStorageTransactorSession struct {
	Contract     *ICertificateStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ICertificateStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICertificateStorageRaw struct {
	Contract *ICertificateStorage // Generic contract binding to access the raw methods on
}

// ICertificateStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICertificateStorageCallerRaw struct {
	Contract *ICertificateStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ICertificateStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICertificateStorageTransactorRaw struct {
	Contract *ICertificateStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICertificateStorage creates a new instance of ICertificateStorage, bound to a specific deployed contract.
func NewICertificateStorage(address common.Address, backend bind.ContractBackend) (*ICertificateStorage, error) {
	contract, err := bindICertificateStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICertificateStorage{ICertificateStorageCaller: ICertificateStorageCaller{contract: contract}, ICertificateStorageTransactor: ICertificateStorageTransactor{contract: contract}, ICertificateStorageFilterer: ICertificateStorageFilterer{contract: contract}}, nil
}

// NewICertificateStorageCaller creates a new read-only instance of ICertificateStorage, bound to a specific deployed contract.
func NewICertificateStorageCaller(address common.Address, caller bind.ContractCaller) (*ICertificateStorageCaller, error) {
	contract, err := bindICertificateStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICertificateStorageCaller{contract: contract}, nil
}

// NewICertificateStorageTransactor creates a new write-only instance of ICertificateStorage, bound to a specific deployed contract.
func NewICertificateStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ICertificateStorageTransactor, error) {
	contract, err := bindICertificateStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICertificateStorageTransactor{contract: contract}, nil
}

// NewICertificateStorageFilterer creates a new log filterer instance of ICertificateStorage, bound to a specific deployed contract.
func NewICertificateStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ICertificateStorageFilterer, error) {
	contract, err := bindICertificateStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICertificateStorageFilterer{contract: contract}, nil
}

// bindICertificateStorage binds a generic wrapper to an already deployed contract.
func bindICertificateStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICertificateStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICertificateStorage *ICertificateStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICertificateStorage.Contract.ICertificateStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICertificateStorage *ICertificateStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICertificateStorage.Contract.ICertificateStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICertificateStorage *ICertificateStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICertificateStorage.Contract.ICertificateStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICertificateStorage *ICertificateStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICertificateStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICertificateStorage *ICertificateStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICertificateStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICertificateStorage *ICertificateStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICertificateStorage.Contract.contract.Transact(opts, method, params...)
}

// CheckUserHasCertificate is a free data retrieval call binding the contract method 0xc4aa76c5.
//
// Solidity: function checkUserHasCertificate(bytes32 securityNo, bytes32 certificateCode) view returns(bool)
func (_ICertificateStorage *ICertificateStorageCaller) CheckUserHasCertificate(opts *bind.CallOpts, securityNo [32]byte, certificateCode [32]byte) (bool, error) {
	var out []interface{}
	err := _ICertificateStorage.contract.Call(opts, &out, "checkUserHasCertificate", securityNo, certificateCode)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckUserHasCertificate is a free data retrieval call binding the contract method 0xc4aa76c5.
//
// Solidity: function checkUserHasCertificate(bytes32 securityNo, bytes32 certificateCode) view returns(bool)
func (_ICertificateStorage *ICertificateStorageSession) CheckUserHasCertificate(securityNo [32]byte, certificateCode [32]byte) (bool, error) {
	return _ICertificateStorage.Contract.CheckUserHasCertificate(&_ICertificateStorage.CallOpts, securityNo, certificateCode)
}

// CheckUserHasCertificate is a free data retrieval call binding the contract method 0xc4aa76c5.
//
// Solidity: function checkUserHasCertificate(bytes32 securityNo, bytes32 certificateCode) view returns(bool)
func (_ICertificateStorage *ICertificateStorageCallerSession) CheckUserHasCertificate(securityNo [32]byte, certificateCode [32]byte) (bool, error) {
	return _ICertificateStorage.Contract.CheckUserHasCertificate(&_ICertificateStorage.CallOpts, securityNo, certificateCode)
}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_ICertificateStorage *ICertificateStorageCaller) GetCertificateBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	var out []interface{}
	err := _ICertificateStorage.contract.Call(opts, &out, "getCertificateBySecurityNo", securityNo)

	if err != nil {
		return *new([]CertificateDefineCertificate), err
	}

	out0 := *abi.ConvertType(out[0], new([]CertificateDefineCertificate)).(*[]CertificateDefineCertificate)

	return out0, err

}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_ICertificateStorage *ICertificateStorageSession) GetCertificateBySecurityNo(securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	return _ICertificateStorage.Contract.GetCertificateBySecurityNo(&_ICertificateStorage.CallOpts, securityNo)
}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_ICertificateStorage *ICertificateStorageCallerSession) GetCertificateBySecurityNo(securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	return _ICertificateStorage.Contract.GetCertificateBySecurityNo(&_ICertificateStorage.CallOpts, securityNo)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_ICertificateStorage *ICertificateStorageTransactor) CreateCertificate(opts *bind.TransactOpts, securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _ICertificateStorage.contract.Transact(opts, "createCertificate", securityNo, certifcate)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_ICertificateStorage *ICertificateStorageSession) CreateCertificate(securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _ICertificateStorage.Contract.CreateCertificate(&_ICertificateStorage.TransactOpts, securityNo, certifcate)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_ICertificateStorage *ICertificateStorageTransactorSession) CreateCertificate(securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _ICertificateStorage.Contract.CreateCertificate(&_ICertificateStorage.TransactOpts, securityNo, certifcate)
}

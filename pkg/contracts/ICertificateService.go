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

// ICertificateServiceMetaData contains all meta data concerning the ICertificateService contract.
var ICertificateServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"certCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"acquiredAt\",\"type\":\"uint16\"}],\"internalType\":\"structCertificateDefine.Certificate\",\"name\":\"certifcate\",\"type\":\"tuple\"}],\"name\":\"createCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getCertificateBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"certCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"acquiredAt\",\"type\":\"uint16\"}],\"internalType\":\"structCertificateDefine.Certificate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ICertificateServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use ICertificateServiceMetaData.ABI instead.
var ICertificateServiceABI = ICertificateServiceMetaData.ABI

// ICertificateService is an auto generated Go binding around an Ethereum contract.
type ICertificateService struct {
	ICertificateServiceCaller     // Read-only binding to the contract
	ICertificateServiceTransactor // Write-only binding to the contract
	ICertificateServiceFilterer   // Log filterer for contract events
}

// ICertificateServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICertificateServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICertificateServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICertificateServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICertificateServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICertificateServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICertificateServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICertificateServiceSession struct {
	Contract     *ICertificateService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ICertificateServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICertificateServiceCallerSession struct {
	Contract *ICertificateServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ICertificateServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICertificateServiceTransactorSession struct {
	Contract     *ICertificateServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ICertificateServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICertificateServiceRaw struct {
	Contract *ICertificateService // Generic contract binding to access the raw methods on
}

// ICertificateServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICertificateServiceCallerRaw struct {
	Contract *ICertificateServiceCaller // Generic read-only contract binding to access the raw methods on
}

// ICertificateServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICertificateServiceTransactorRaw struct {
	Contract *ICertificateServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICertificateService creates a new instance of ICertificateService, bound to a specific deployed contract.
func NewICertificateService(address common.Address, backend bind.ContractBackend) (*ICertificateService, error) {
	contract, err := bindICertificateService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICertificateService{ICertificateServiceCaller: ICertificateServiceCaller{contract: contract}, ICertificateServiceTransactor: ICertificateServiceTransactor{contract: contract}, ICertificateServiceFilterer: ICertificateServiceFilterer{contract: contract}}, nil
}

// NewICertificateServiceCaller creates a new read-only instance of ICertificateService, bound to a specific deployed contract.
func NewICertificateServiceCaller(address common.Address, caller bind.ContractCaller) (*ICertificateServiceCaller, error) {
	contract, err := bindICertificateService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICertificateServiceCaller{contract: contract}, nil
}

// NewICertificateServiceTransactor creates a new write-only instance of ICertificateService, bound to a specific deployed contract.
func NewICertificateServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*ICertificateServiceTransactor, error) {
	contract, err := bindICertificateService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICertificateServiceTransactor{contract: contract}, nil
}

// NewICertificateServiceFilterer creates a new log filterer instance of ICertificateService, bound to a specific deployed contract.
func NewICertificateServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*ICertificateServiceFilterer, error) {
	contract, err := bindICertificateService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICertificateServiceFilterer{contract: contract}, nil
}

// bindICertificateService binds a generic wrapper to an already deployed contract.
func bindICertificateService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICertificateServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICertificateService *ICertificateServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICertificateService.Contract.ICertificateServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICertificateService *ICertificateServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICertificateService.Contract.ICertificateServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICertificateService *ICertificateServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICertificateService.Contract.ICertificateServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICertificateService *ICertificateServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICertificateService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICertificateService *ICertificateServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICertificateService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICertificateService *ICertificateServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICertificateService.Contract.contract.Transact(opts, method, params...)
}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_ICertificateService *ICertificateServiceCaller) GetCertificateBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	var out []interface{}
	err := _ICertificateService.contract.Call(opts, &out, "getCertificateBySecurityNo", securityNo)

	if err != nil {
		return *new([]CertificateDefineCertificate), err
	}

	out0 := *abi.ConvertType(out[0], new([]CertificateDefineCertificate)).(*[]CertificateDefineCertificate)

	return out0, err

}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_ICertificateService *ICertificateServiceSession) GetCertificateBySecurityNo(securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	return _ICertificateService.Contract.GetCertificateBySecurityNo(&_ICertificateService.CallOpts, securityNo)
}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_ICertificateService *ICertificateServiceCallerSession) GetCertificateBySecurityNo(securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	return _ICertificateService.Contract.GetCertificateBySecurityNo(&_ICertificateService.CallOpts, securityNo)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_ICertificateService *ICertificateServiceTransactor) CreateCertificate(opts *bind.TransactOpts, securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _ICertificateService.contract.Transact(opts, "createCertificate", securityNo, certifcate)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_ICertificateService *ICertificateServiceSession) CreateCertificate(securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _ICertificateService.Contract.CreateCertificate(&_ICertificateService.TransactOpts, securityNo, certifcate)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_ICertificateService *ICertificateServiceTransactorSession) CreateCertificate(securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _ICertificateService.Contract.CreateCertificate(&_ICertificateService.TransactOpts, securityNo, certifcate)
}

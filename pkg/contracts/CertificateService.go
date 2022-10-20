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


// CertificateServiceMetaData contains all meta data concerning the CertificateService contract.
var CertificateServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dispatcher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"certCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"acquiredAt\",\"type\":\"uint16\"}],\"internalType\":\"structCertificateDefine.Certificate\",\"name\":\"certifcate\",\"type\":\"tuple\"}],\"name\":\"createCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getCertificateBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"certCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"acquiredAt\",\"type\":\"uint16\"}],\"internalType\":\"structCertificateDefine.Certificate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CertificateServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use CertificateServiceMetaData.ABI instead.
var CertificateServiceABI = CertificateServiceMetaData.ABI

// CertificateService is an auto generated Go binding around an Ethereum contract.
type CertificateService struct {
	CertificateServiceCaller     // Read-only binding to the contract
	CertificateServiceTransactor // Write-only binding to the contract
	CertificateServiceFilterer   // Log filterer for contract events
}

// CertificateServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertificateServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertificateServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertificateServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertificateServiceSession struct {
	Contract     *CertificateService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CertificateServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertificateServiceCallerSession struct {
	Contract *CertificateServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CertificateServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertificateServiceTransactorSession struct {
	Contract     *CertificateServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CertificateServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertificateServiceRaw struct {
	Contract *CertificateService // Generic contract binding to access the raw methods on
}

// CertificateServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertificateServiceCallerRaw struct {
	Contract *CertificateServiceCaller // Generic read-only contract binding to access the raw methods on
}

// CertificateServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertificateServiceTransactorRaw struct {
	Contract *CertificateServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertificateService creates a new instance of CertificateService, bound to a specific deployed contract.
func NewCertificateService(address common.Address, backend bind.ContractBackend) (*CertificateService, error) {
	contract, err := bindCertificateService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CertificateService{CertificateServiceCaller: CertificateServiceCaller{contract: contract}, CertificateServiceTransactor: CertificateServiceTransactor{contract: contract}, CertificateServiceFilterer: CertificateServiceFilterer{contract: contract}}, nil
}

// NewCertificateServiceCaller creates a new read-only instance of CertificateService, bound to a specific deployed contract.
func NewCertificateServiceCaller(address common.Address, caller bind.ContractCaller) (*CertificateServiceCaller, error) {
	contract, err := bindCertificateService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateServiceCaller{contract: contract}, nil
}

// NewCertificateServiceTransactor creates a new write-only instance of CertificateService, bound to a specific deployed contract.
func NewCertificateServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*CertificateServiceTransactor, error) {
	contract, err := bindCertificateService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateServiceTransactor{contract: contract}, nil
}

// NewCertificateServiceFilterer creates a new log filterer instance of CertificateService, bound to a specific deployed contract.
func NewCertificateServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*CertificateServiceFilterer, error) {
	contract, err := bindCertificateService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertificateServiceFilterer{contract: contract}, nil
}

// bindCertificateService binds a generic wrapper to an already deployed contract.
func bindCertificateService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CertificateServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificateService *CertificateServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CertificateService.Contract.CertificateServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificateService *CertificateServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificateService.Contract.CertificateServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificateService *CertificateServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertificateService.Contract.CertificateServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificateService *CertificateServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CertificateService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificateService *CertificateServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificateService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificateService *CertificateServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertificateService.Contract.contract.Transact(opts, method, params...)
}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_CertificateService *CertificateServiceCaller) GetCertificateBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	var out []interface{}
	err := _CertificateService.contract.Call(opts, &out, "getCertificateBySecurityNo", securityNo)

	if err != nil {
		return *new([]CertificateDefineCertificate), err
	}

	out0 := *abi.ConvertType(out[0], new([]CertificateDefineCertificate)).(*[]CertificateDefineCertificate)

	return out0, err

}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_CertificateService *CertificateServiceSession) GetCertificateBySecurityNo(securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	return _CertificateService.Contract.GetCertificateBySecurityNo(&_CertificateService.CallOpts, securityNo)
}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_CertificateService *CertificateServiceCallerSession) GetCertificateBySecurityNo(securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	return _CertificateService.Contract.GetCertificateBySecurityNo(&_CertificateService.CallOpts, securityNo)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_CertificateService *CertificateServiceTransactor) CreateCertificate(opts *bind.TransactOpts, securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _CertificateService.contract.Transact(opts, "createCertificate", securityNo, certifcate)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_CertificateService *CertificateServiceSession) CreateCertificate(securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _CertificateService.Contract.CreateCertificate(&_CertificateService.TransactOpts, securityNo, certifcate)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_CertificateService *CertificateServiceTransactorSession) CreateCertificate(securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _CertificateService.Contract.CreateCertificate(&_CertificateService.TransactOpts, securityNo, certifcate)
}

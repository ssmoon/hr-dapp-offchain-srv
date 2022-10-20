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

// IGetContractNameMetaData contains all meta data concerning the IGetContractName contract.
var IGetContractNameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// IGetContractNameABI is the input ABI used to generate the binding from.
// Deprecated: Use IGetContractNameMetaData.ABI instead.
var IGetContractNameABI = IGetContractNameMetaData.ABI

// IGetContractName is an auto generated Go binding around an Ethereum contract.
type IGetContractName struct {
	IGetContractNameCaller     // Read-only binding to the contract
	IGetContractNameTransactor // Write-only binding to the contract
	IGetContractNameFilterer   // Log filterer for contract events
}

// IGetContractNameCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGetContractNameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGetContractNameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGetContractNameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGetContractNameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGetContractNameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGetContractNameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGetContractNameSession struct {
	Contract     *IGetContractName // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGetContractNameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGetContractNameCallerSession struct {
	Contract *IGetContractNameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IGetContractNameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGetContractNameTransactorSession struct {
	Contract     *IGetContractNameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IGetContractNameRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGetContractNameRaw struct {
	Contract *IGetContractName // Generic contract binding to access the raw methods on
}

// IGetContractNameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGetContractNameCallerRaw struct {
	Contract *IGetContractNameCaller // Generic read-only contract binding to access the raw methods on
}

// IGetContractNameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGetContractNameTransactorRaw struct {
	Contract *IGetContractNameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGetContractName creates a new instance of IGetContractName, bound to a specific deployed contract.
func NewIGetContractName(address common.Address, backend bind.ContractBackend) (*IGetContractName, error) {
	contract, err := bindIGetContractName(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGetContractName{IGetContractNameCaller: IGetContractNameCaller{contract: contract}, IGetContractNameTransactor: IGetContractNameTransactor{contract: contract}, IGetContractNameFilterer: IGetContractNameFilterer{contract: contract}}, nil
}

// NewIGetContractNameCaller creates a new read-only instance of IGetContractName, bound to a specific deployed contract.
func NewIGetContractNameCaller(address common.Address, caller bind.ContractCaller) (*IGetContractNameCaller, error) {
	contract, err := bindIGetContractName(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGetContractNameCaller{contract: contract}, nil
}

// NewIGetContractNameTransactor creates a new write-only instance of IGetContractName, bound to a specific deployed contract.
func NewIGetContractNameTransactor(address common.Address, transactor bind.ContractTransactor) (*IGetContractNameTransactor, error) {
	contract, err := bindIGetContractName(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGetContractNameTransactor{contract: contract}, nil
}

// NewIGetContractNameFilterer creates a new log filterer instance of IGetContractName, bound to a specific deployed contract.
func NewIGetContractNameFilterer(address common.Address, filterer bind.ContractFilterer) (*IGetContractNameFilterer, error) {
	contract, err := bindIGetContractName(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGetContractNameFilterer{contract: contract}, nil
}

// bindIGetContractName binds a generic wrapper to an already deployed contract.
func bindIGetContractName(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGetContractNameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGetContractName *IGetContractNameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGetContractName.Contract.IGetContractNameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGetContractName *IGetContractNameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGetContractName.Contract.IGetContractNameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGetContractName *IGetContractNameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGetContractName.Contract.IGetContractNameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGetContractName *IGetContractNameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGetContractName.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGetContractName *IGetContractNameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGetContractName.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGetContractName *IGetContractNameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGetContractName.Contract.contract.Transact(opts, method, params...)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(bytes32)
func (_IGetContractName *IGetContractNameCaller) GetContractName(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IGetContractName.contract.Call(opts, &out, "getContractName")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(bytes32)
func (_IGetContractName *IGetContractNameSession) GetContractName() ([32]byte, error) {
	return _IGetContractName.Contract.GetContractName(&_IGetContractName.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(bytes32)
func (_IGetContractName *IGetContractNameCallerSession) GetContractName() ([32]byte, error) {
	return _IGetContractName.Contract.GetContractName(&_IGetContractName.CallOpts)
}

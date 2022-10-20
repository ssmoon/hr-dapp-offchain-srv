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

// BaseResolverMetaData contains all meta data concerning the BaseResolver contract.
var BaseResolverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dispatcher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
}

// BaseResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseResolverMetaData.ABI instead.
var BaseResolverABI = BaseResolverMetaData.ABI

// BaseResolver is an auto generated Go binding around an Ethereum contract.
type BaseResolver struct {
	BaseResolverCaller     // Read-only binding to the contract
	BaseResolverTransactor // Write-only binding to the contract
	BaseResolverFilterer   // Log filterer for contract events
}

// BaseResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseResolverSession struct {
	Contract     *BaseResolver     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseResolverCallerSession struct {
	Contract *BaseResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BaseResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseResolverTransactorSession struct {
	Contract     *BaseResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BaseResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseResolverRaw struct {
	Contract *BaseResolver // Generic contract binding to access the raw methods on
}

// BaseResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseResolverCallerRaw struct {
	Contract *BaseResolverCaller // Generic read-only contract binding to access the raw methods on
}

// BaseResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseResolverTransactorRaw struct {
	Contract *BaseResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseResolver creates a new instance of BaseResolver, bound to a specific deployed contract.
func NewBaseResolver(address common.Address, backend bind.ContractBackend) (*BaseResolver, error) {
	contract, err := bindBaseResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseResolver{BaseResolverCaller: BaseResolverCaller{contract: contract}, BaseResolverTransactor: BaseResolverTransactor{contract: contract}, BaseResolverFilterer: BaseResolverFilterer{contract: contract}}, nil
}

// NewBaseResolverCaller creates a new read-only instance of BaseResolver, bound to a specific deployed contract.
func NewBaseResolverCaller(address common.Address, caller bind.ContractCaller) (*BaseResolverCaller, error) {
	contract, err := bindBaseResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseResolverCaller{contract: contract}, nil
}

// NewBaseResolverTransactor creates a new write-only instance of BaseResolver, bound to a specific deployed contract.
func NewBaseResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseResolverTransactor, error) {
	contract, err := bindBaseResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseResolverTransactor{contract: contract}, nil
}

// NewBaseResolverFilterer creates a new log filterer instance of BaseResolver, bound to a specific deployed contract.
func NewBaseResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseResolverFilterer, error) {
	contract, err := bindBaseResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseResolverFilterer{contract: contract}, nil
}

// bindBaseResolver binds a generic wrapper to an already deployed contract.
func bindBaseResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseResolver *BaseResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseResolver.Contract.BaseResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseResolver *BaseResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseResolver.Contract.BaseResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseResolver *BaseResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseResolver.Contract.BaseResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseResolver *BaseResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseResolver *BaseResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseResolver *BaseResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseResolver.Contract.contract.Transact(opts, method, params...)
}

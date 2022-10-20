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

// WorkerServiceMetaData contains all meta data concerning the WorkerService contract.
var WorkerServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dispatcher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"graduatedAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"birthAt\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"collegeCode\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structWorkerDefine.Worker\",\"name\":\"worker\",\"type\":\"tuple\"}],\"name\":\"workerCreated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"graduatedAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"birthAt\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"collegeCode\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structWorkerDefine.Worker\",\"name\":\"worker\",\"type\":\"tuple\"}],\"name\":\"createWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getWorkerBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"graduatedAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"birthAt\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"collegeCode\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structWorkerDefine.Worker\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WorkerServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use WorkerServiceMetaData.ABI instead.
var WorkerServiceABI = WorkerServiceMetaData.ABI

// WorkerService is an auto generated Go binding around an Ethereum contract.
type WorkerService struct {
	WorkerServiceCaller     // Read-only binding to the contract
	WorkerServiceTransactor // Write-only binding to the contract
	WorkerServiceFilterer   // Log filterer for contract events
}

// WorkerServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorkerServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorkerServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorkerServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorkerServiceSession struct {
	Contract     *WorkerService    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorkerServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorkerServiceCallerSession struct {
	Contract *WorkerServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WorkerServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorkerServiceTransactorSession struct {
	Contract     *WorkerServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WorkerServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorkerServiceRaw struct {
	Contract *WorkerService // Generic contract binding to access the raw methods on
}

// WorkerServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorkerServiceCallerRaw struct {
	Contract *WorkerServiceCaller // Generic read-only contract binding to access the raw methods on
}

// WorkerServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorkerServiceTransactorRaw struct {
	Contract *WorkerServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorkerService creates a new instance of WorkerService, bound to a specific deployed contract.
func NewWorkerService(address common.Address, backend bind.ContractBackend) (*WorkerService, error) {
	contract, err := bindWorkerService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkerService{WorkerServiceCaller: WorkerServiceCaller{contract: contract}, WorkerServiceTransactor: WorkerServiceTransactor{contract: contract}, WorkerServiceFilterer: WorkerServiceFilterer{contract: contract}}, nil
}

// NewWorkerServiceCaller creates a new read-only instance of WorkerService, bound to a specific deployed contract.
func NewWorkerServiceCaller(address common.Address, caller bind.ContractCaller) (*WorkerServiceCaller, error) {
	contract, err := bindWorkerService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerServiceCaller{contract: contract}, nil
}

// NewWorkerServiceTransactor creates a new write-only instance of WorkerService, bound to a specific deployed contract.
func NewWorkerServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkerServiceTransactor, error) {
	contract, err := bindWorkerService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerServiceTransactor{contract: contract}, nil
}

// NewWorkerServiceFilterer creates a new log filterer instance of WorkerService, bound to a specific deployed contract.
func NewWorkerServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkerServiceFilterer, error) {
	contract, err := bindWorkerService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkerServiceFilterer{contract: contract}, nil
}

// bindWorkerService binds a generic wrapper to an already deployed contract.
func bindWorkerService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WorkerServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerService *WorkerServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerService.Contract.WorkerServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerService *WorkerServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerService.Contract.WorkerServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerService *WorkerServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerService.Contract.WorkerServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerService *WorkerServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerService *WorkerServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerService *WorkerServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerService.Contract.contract.Transact(opts, method, params...)
}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_WorkerService *WorkerServiceCaller) GetWorkerBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) (WorkerDefineWorker, error) {
	var out []interface{}
	err := _WorkerService.contract.Call(opts, &out, "getWorkerBySecurityNo", securityNo)

	if err != nil {
		return *new(WorkerDefineWorker), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkerDefineWorker)).(*WorkerDefineWorker)

	return out0, err

}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_WorkerService *WorkerServiceSession) GetWorkerBySecurityNo(securityNo [32]byte) (WorkerDefineWorker, error) {
	return _WorkerService.Contract.GetWorkerBySecurityNo(&_WorkerService.CallOpts, securityNo)
}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_WorkerService *WorkerServiceCallerSession) GetWorkerBySecurityNo(securityNo [32]byte) (WorkerDefineWorker, error) {
	return _WorkerService.Contract.GetWorkerBySecurityNo(&_WorkerService.CallOpts, securityNo)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_WorkerService *WorkerServiceTransactor) CreateWorker(opts *bind.TransactOpts, worker WorkerDefineWorker) (*types.Transaction, error) {
	return _WorkerService.contract.Transact(opts, "createWorker", worker)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_WorkerService *WorkerServiceSession) CreateWorker(worker WorkerDefineWorker) (*types.Transaction, error) {
	return _WorkerService.Contract.CreateWorker(&_WorkerService.TransactOpts, worker)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_WorkerService *WorkerServiceTransactorSession) CreateWorker(worker WorkerDefineWorker) (*types.Transaction, error) {
	return _WorkerService.Contract.CreateWorker(&_WorkerService.TransactOpts, worker)
}

// WorkerServiceWorkerCreatedIterator is returned from FilterWorkerCreated and is used to iterate over the raw logs and unpacked data for WorkerCreated events raised by the WorkerService contract.
type WorkerServiceWorkerCreatedIterator struct {
	Event *WorkerServiceWorkerCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WorkerServiceWorkerCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerServiceWorkerCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WorkerServiceWorkerCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WorkerServiceWorkerCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerServiceWorkerCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerServiceWorkerCreated represents a WorkerCreated event raised by the WorkerService contract.
type WorkerServiceWorkerCreated struct {
	Worker WorkerDefineWorker
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWorkerCreated is a free log retrieval operation binding the contract event 0xf6862b5f094c436491d330ef12a80914dfecbcb2bb15dc775ae6848a57b32f67.
//
// Solidity: event workerCreated((bytes32,uint16,uint16,bytes32,bool) worker)
func (_WorkerService *WorkerServiceFilterer) FilterWorkerCreated(opts *bind.FilterOpts) (*WorkerServiceWorkerCreatedIterator, error) {

	logs, sub, err := _WorkerService.contract.FilterLogs(opts, "workerCreated")
	if err != nil {
		return nil, err
	}
	return &WorkerServiceWorkerCreatedIterator{contract: _WorkerService.contract, event: "workerCreated", logs: logs, sub: sub}, nil
}

// WatchWorkerCreated is a free log subscription operation binding the contract event 0xf6862b5f094c436491d330ef12a80914dfecbcb2bb15dc775ae6848a57b32f67.
//
// Solidity: event workerCreated((bytes32,uint16,uint16,bytes32,bool) worker)
func (_WorkerService *WorkerServiceFilterer) WatchWorkerCreated(opts *bind.WatchOpts, sink chan<- *WorkerServiceWorkerCreated) (event.Subscription, error) {

	logs, sub, err := _WorkerService.contract.WatchLogs(opts, "workerCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerServiceWorkerCreated)
				if err := _WorkerService.contract.UnpackLog(event, "workerCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWorkerCreated is a log parse operation binding the contract event 0xf6862b5f094c436491d330ef12a80914dfecbcb2bb15dc775ae6848a57b32f67.
//
// Solidity: event workerCreated((bytes32,uint16,uint16,bytes32,bool) worker)
func (_WorkerService *WorkerServiceFilterer) ParseWorkerCreated(log types.Log) (*WorkerServiceWorkerCreated, error) {
	event := new(WorkerServiceWorkerCreated)
	if err := _WorkerService.contract.UnpackLog(event, "workerCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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


// CareerServiceMetaData contains all meta data concerning the CareerService contract.
var CareerServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dispatcher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structWorkExperienceDefine.WorkExperience\",\"name\":\"workExperience\",\"type\":\"tuple\"}],\"name\":\"careerExperienceCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"internalType\":\"structWorkExperienceDefine.WorkExperience\",\"name\":\"workExperience\",\"type\":\"tuple\"}],\"name\":\"addWorkExperience\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"endYear\",\"type\":\"uint16\"}],\"name\":\"finishLastCareer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getWorkExperienceBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"internalType\":\"structWorkExperienceDefine.WorkExperience[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CareerServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use CareerServiceMetaData.ABI instead.
var CareerServiceABI = CareerServiceMetaData.ABI

// CareerService is an auto generated Go binding around an Ethereum contract.
type CareerService struct {
	CareerServiceCaller     // Read-only binding to the contract
	CareerServiceTransactor // Write-only binding to the contract
	CareerServiceFilterer   // Log filterer for contract events
}

// CareerServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type CareerServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CareerServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CareerServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CareerServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CareerServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CareerServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CareerServiceSession struct {
	Contract     *CareerService    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CareerServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CareerServiceCallerSession struct {
	Contract *CareerServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CareerServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CareerServiceTransactorSession struct {
	Contract     *CareerServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CareerServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type CareerServiceRaw struct {
	Contract *CareerService // Generic contract binding to access the raw methods on
}

// CareerServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CareerServiceCallerRaw struct {
	Contract *CareerServiceCaller // Generic read-only contract binding to access the raw methods on
}

// CareerServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CareerServiceTransactorRaw struct {
	Contract *CareerServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCareerService creates a new instance of CareerService, bound to a specific deployed contract.
func NewCareerService(address common.Address, backend bind.ContractBackend) (*CareerService, error) {
	contract, err := bindCareerService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CareerService{CareerServiceCaller: CareerServiceCaller{contract: contract}, CareerServiceTransactor: CareerServiceTransactor{contract: contract}, CareerServiceFilterer: CareerServiceFilterer{contract: contract}}, nil
}

// NewCareerServiceCaller creates a new read-only instance of CareerService, bound to a specific deployed contract.
func NewCareerServiceCaller(address common.Address, caller bind.ContractCaller) (*CareerServiceCaller, error) {
	contract, err := bindCareerService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CareerServiceCaller{contract: contract}, nil
}

// NewCareerServiceTransactor creates a new write-only instance of CareerService, bound to a specific deployed contract.
func NewCareerServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*CareerServiceTransactor, error) {
	contract, err := bindCareerService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CareerServiceTransactor{contract: contract}, nil
}

// NewCareerServiceFilterer creates a new log filterer instance of CareerService, bound to a specific deployed contract.
func NewCareerServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*CareerServiceFilterer, error) {
	contract, err := bindCareerService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CareerServiceFilterer{contract: contract}, nil
}

// bindCareerService binds a generic wrapper to an already deployed contract.
func bindCareerService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CareerServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CareerService *CareerServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CareerService.Contract.CareerServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CareerService *CareerServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CareerService.Contract.CareerServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CareerService *CareerServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CareerService.Contract.CareerServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CareerService *CareerServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CareerService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CareerService *CareerServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CareerService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CareerService *CareerServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CareerService.Contract.contract.Transact(opts, method, params...)
}

// GetWorkExperienceBySecurityNo is a free data retrieval call binding the contract method 0xabaed258.
//
// Solidity: function getWorkExperienceBySecurityNo(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32)[])
func (_CareerService *CareerServiceCaller) GetWorkExperienceBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) ([]WorkExperienceDefineWorkExperience, error) {
	var out []interface{}
	err := _CareerService.contract.Call(opts, &out, "getWorkExperienceBySecurityNo", securityNo)

	if err != nil {
		return *new([]WorkExperienceDefineWorkExperience), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkExperienceDefineWorkExperience)).(*[]WorkExperienceDefineWorkExperience)

	return out0, err

}

// GetWorkExperienceBySecurityNo is a free data retrieval call binding the contract method 0xabaed258.
//
// Solidity: function getWorkExperienceBySecurityNo(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32)[])
func (_CareerService *CareerServiceSession) GetWorkExperienceBySecurityNo(securityNo [32]byte) ([]WorkExperienceDefineWorkExperience, error) {
	return _CareerService.Contract.GetWorkExperienceBySecurityNo(&_CareerService.CallOpts, securityNo)
}

// GetWorkExperienceBySecurityNo is a free data retrieval call binding the contract method 0xabaed258.
//
// Solidity: function getWorkExperienceBySecurityNo(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32)[])
func (_CareerService *CareerServiceCallerSession) GetWorkExperienceBySecurityNo(securityNo [32]byte) ([]WorkExperienceDefineWorkExperience, error) {
	return _CareerService.Contract.GetWorkExperienceBySecurityNo(&_CareerService.CallOpts, securityNo)
}

// AddWorkExperience is a paid mutator transaction binding the contract method 0x36efe55e.
//
// Solidity: function addWorkExperience(bytes32 securityNo, (uint16,uint16,bool,bytes32) workExperience) returns()
func (_CareerService *CareerServiceTransactor) AddWorkExperience(opts *bind.TransactOpts, securityNo [32]byte, workExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _CareerService.contract.Transact(opts, "addWorkExperience", securityNo, workExperience)
}

// AddWorkExperience is a paid mutator transaction binding the contract method 0x36efe55e.
//
// Solidity: function addWorkExperience(bytes32 securityNo, (uint16,uint16,bool,bytes32) workExperience) returns()
func (_CareerService *CareerServiceSession) AddWorkExperience(securityNo [32]byte, workExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _CareerService.Contract.AddWorkExperience(&_CareerService.TransactOpts, securityNo, workExperience)
}

// AddWorkExperience is a paid mutator transaction binding the contract method 0x36efe55e.
//
// Solidity: function addWorkExperience(bytes32 securityNo, (uint16,uint16,bool,bytes32) workExperience) returns()
func (_CareerService *CareerServiceTransactorSession) AddWorkExperience(securityNo [32]byte, workExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _CareerService.Contract.AddWorkExperience(&_CareerService.TransactOpts, securityNo, workExperience)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_CareerService *CareerServiceTransactor) FinishLastCareer(opts *bind.TransactOpts, securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _CareerService.contract.Transact(opts, "finishLastCareer", securityNo, endYear)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_CareerService *CareerServiceSession) FinishLastCareer(securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _CareerService.Contract.FinishLastCareer(&_CareerService.TransactOpts, securityNo, endYear)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_CareerService *CareerServiceTransactorSession) FinishLastCareer(securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _CareerService.Contract.FinishLastCareer(&_CareerService.TransactOpts, securityNo, endYear)
}

// CareerServiceCareerExperienceCreatedIterator is returned from FilterCareerExperienceCreated and is used to iterate over the raw logs and unpacked data for CareerExperienceCreated events raised by the CareerService contract.
type CareerServiceCareerExperienceCreatedIterator struct {
	Event *CareerServiceCareerExperienceCreated // Event containing the contract specifics and raw log

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
func (it *CareerServiceCareerExperienceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CareerServiceCareerExperienceCreated)
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
		it.Event = new(CareerServiceCareerExperienceCreated)
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
func (it *CareerServiceCareerExperienceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CareerServiceCareerExperienceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CareerServiceCareerExperienceCreated represents a CareerExperienceCreated event raised by the CareerService contract.
type CareerServiceCareerExperienceCreated struct {
	WorkExperience WorkExperienceDefineWorkExperience
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCareerExperienceCreated is a free log retrieval operation binding the contract event 0xc6a63ac1bc54eec3c544951574780d90ed14bb22cdc5b328d352d79fe92af36c.
//
// Solidity: event careerExperienceCreated((uint16,uint16,bool,bytes32) workExperience)
func (_CareerService *CareerServiceFilterer) FilterCareerExperienceCreated(opts *bind.FilterOpts) (*CareerServiceCareerExperienceCreatedIterator, error) {

	logs, sub, err := _CareerService.contract.FilterLogs(opts, "careerExperienceCreated")
	if err != nil {
		return nil, err
	}
	return &CareerServiceCareerExperienceCreatedIterator{contract: _CareerService.contract, event: "careerExperienceCreated", logs: logs, sub: sub}, nil
}

// WatchCareerExperienceCreated is a free log subscription operation binding the contract event 0xc6a63ac1bc54eec3c544951574780d90ed14bb22cdc5b328d352d79fe92af36c.
//
// Solidity: event careerExperienceCreated((uint16,uint16,bool,bytes32) workExperience)
func (_CareerService *CareerServiceFilterer) WatchCareerExperienceCreated(opts *bind.WatchOpts, sink chan<- *CareerServiceCareerExperienceCreated) (event.Subscription, error) {

	logs, sub, err := _CareerService.contract.WatchLogs(opts, "careerExperienceCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CareerServiceCareerExperienceCreated)
				if err := _CareerService.contract.UnpackLog(event, "careerExperienceCreated", log); err != nil {
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

// ParseCareerExperienceCreated is a log parse operation binding the contract event 0xc6a63ac1bc54eec3c544951574780d90ed14bb22cdc5b328d352d79fe92af36c.
//
// Solidity: event careerExperienceCreated((uint16,uint16,bool,bytes32) workExperience)
func (_CareerService *CareerServiceFilterer) ParseCareerExperienceCreated(log types.Log) (*CareerServiceCareerExperienceCreated, error) {
	event := new(CareerServiceCareerExperienceCreated)
	if err := _CareerService.contract.UnpackLog(event, "careerExperienceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

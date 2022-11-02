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

// CertificateDefineCertificate is an auto generated low-level Go binding around an user-defined struct.
type CertificateDefineCertificate struct {
	CertCode   [32]byte
	AcquiredAt uint16
}

// WorkExperienceDefineWorkExperience is an auto generated low-level Go binding around an user-defined struct.
type WorkExperienceDefineWorkExperience struct {
	StartAt     uint16
	EndAt       uint16
	HasEnded    bool
	CompanyCode [32]byte
}

// WorkerDefineWorker is an auto generated low-level Go binding around an user-defined struct.
type WorkerDefineWorker struct {
	SecurityNo  [32]byte
	GraduatedAt uint16
	BirthAt     uint16
	CollegeCode [32]byte
	IsValue     bool
}

// FacadeMetaData contains all meta data concerning the Facade contract.
var FacadeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRIVILEGED_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dispatcher\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"certCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"acquiredAt\",\"type\":\"uint16\"}],\"internalType\":\"structCertificateDefine.Certificate\",\"name\":\"certifcate\",\"type\":\"tuple\"}],\"name\":\"createCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getCertificateBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"certCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"acquiredAt\",\"type\":\"uint16\"}],\"internalType\":\"structCertificateDefine.Certificate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"internalType\":\"structWorkExperienceDefine.WorkExperience\",\"name\":\"workExperience\",\"type\":\"tuple\"}],\"name\":\"addWorkExperience\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"endYear\",\"type\":\"uint16\"}],\"name\":\"finishLastCareer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getWorkExperienceBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"startAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endAt\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"hasEnded\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"companyCode\",\"type\":\"bytes32\"}],\"internalType\":\"structWorkExperienceDefine.WorkExperience[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"graduatedAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"birthAt\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"collegeCode\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structWorkerDefine.Worker\",\"name\":\"worker\",\"type\":\"tuple\"}],\"name\":\"createWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"}],\"name\":\"getWorkerBySecurityNo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"securityNo\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"graduatedAt\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"birthAt\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"collegeCode\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structWorkerDefine.Worker\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"createUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"checkUserExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ping\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pingByOwner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pingByPrivilegedUser\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FacadeABI is the input ABI used to generate the binding from.
// Deprecated: Use FacadeMetaData.ABI instead.
var FacadeABI = FacadeMetaData.ABI

// Facade is an auto generated Go binding around an Ethereum contract.
type Facade struct {
	FacadeCaller     // Read-only binding to the contract
	FacadeTransactor // Write-only binding to the contract
	FacadeFilterer   // Log filterer for contract events
}

// FacadeCaller is an auto generated read-only Go binding around an Ethereum contract.
type FacadeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FacadeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FacadeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FacadeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FacadeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FacadeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FacadeSession struct {
	Contract     *Facade           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FacadeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FacadeCallerSession struct {
	Contract *FacadeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FacadeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FacadeTransactorSession struct {
	Contract     *FacadeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FacadeRaw is an auto generated low-level Go binding around an Ethereum contract.
type FacadeRaw struct {
	Contract *Facade // Generic contract binding to access the raw methods on
}

// FacadeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FacadeCallerRaw struct {
	Contract *FacadeCaller // Generic read-only contract binding to access the raw methods on
}

// FacadeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FacadeTransactorRaw struct {
	Contract *FacadeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFacade creates a new instance of Facade, bound to a specific deployed contract.
func NewFacade(address common.Address, backend bind.ContractBackend) (*Facade, error) {
	contract, err := bindFacade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Facade{FacadeCaller: FacadeCaller{contract: contract}, FacadeTransactor: FacadeTransactor{contract: contract}, FacadeFilterer: FacadeFilterer{contract: contract}}, nil
}

// NewFacadeCaller creates a new read-only instance of Facade, bound to a specific deployed contract.
func NewFacadeCaller(address common.Address, caller bind.ContractCaller) (*FacadeCaller, error) {
	contract, err := bindFacade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FacadeCaller{contract: contract}, nil
}

// NewFacadeTransactor creates a new write-only instance of Facade, bound to a specific deployed contract.
func NewFacadeTransactor(address common.Address, transactor bind.ContractTransactor) (*FacadeTransactor, error) {
	contract, err := bindFacade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FacadeTransactor{contract: contract}, nil
}

// NewFacadeFilterer creates a new log filterer instance of Facade, bound to a specific deployed contract.
func NewFacadeFilterer(address common.Address, filterer bind.ContractFilterer) (*FacadeFilterer, error) {
	contract, err := bindFacade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FacadeFilterer{contract: contract}, nil
}

// bindFacade binds a generic wrapper to an already deployed contract.
func bindFacade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FacadeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Facade *FacadeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Facade.Contract.FacadeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Facade *FacadeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Facade.Contract.FacadeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Facade *FacadeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Facade.Contract.FacadeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Facade *FacadeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Facade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Facade *FacadeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Facade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Facade *FacadeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Facade.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Facade *FacadeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Facade.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Facade *FacadeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Facade.Contract.DEFAULTADMINROLE(&_Facade.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Facade *FacadeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Facade.Contract.DEFAULTADMINROLE(&_Facade.CallOpts)
}

// PRIVILEGEDROLE is a free data retrieval call binding the contract method 0x93fb8f76.
//
// Solidity: function PRIVILEGED_ROLE() view returns(bytes32)
func (_Facade *FacadeCaller) PRIVILEGEDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Facade.contract.Call(opts, &out, "PRIVILEGED_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PRIVILEGEDROLE is a free data retrieval call binding the contract method 0x93fb8f76.
//
// Solidity: function PRIVILEGED_ROLE() view returns(bytes32)
func (_Facade *FacadeSession) PRIVILEGEDROLE() ([32]byte, error) {
	return _Facade.Contract.PRIVILEGEDROLE(&_Facade.CallOpts)
}

// PRIVILEGEDROLE is a free data retrieval call binding the contract method 0x93fb8f76.
//
// Solidity: function PRIVILEGED_ROLE() view returns(bytes32)
func (_Facade *FacadeCallerSession) PRIVILEGEDROLE() ([32]byte, error) {
	return _Facade.Contract.PRIVILEGEDROLE(&_Facade.CallOpts)
}

// CheckUserExist is a free data retrieval call binding the contract method 0x926bef63.
//
// Solidity: function checkUserExist(address addr) view returns(bool)
func (_Facade *FacadeCaller) CheckUserExist(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Facade.contract.Call(opts, &out, "checkUserExist", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckUserExist is a free data retrieval call binding the contract method 0x926bef63.
//
// Solidity: function checkUserExist(address addr) view returns(bool)
func (_Facade *FacadeSession) CheckUserExist(addr common.Address) (bool, error) {
	return _Facade.Contract.CheckUserExist(&_Facade.CallOpts, addr)
}

// CheckUserExist is a free data retrieval call binding the contract method 0x926bef63.
//
// Solidity: function checkUserExist(address addr) view returns(bool)
func (_Facade *FacadeCallerSession) CheckUserExist(addr common.Address) (bool, error) {
	return _Facade.Contract.CheckUserExist(&_Facade.CallOpts, addr)
}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_Facade *FacadeCaller) GetCertificateBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	var out []interface{}
	err := _Facade.contract.Call(opts, &out, "getCertificateBySecurityNo", securityNo)

	if err != nil {
		return *new([]CertificateDefineCertificate), err
	}

	out0 := *abi.ConvertType(out[0], new([]CertificateDefineCertificate)).(*[]CertificateDefineCertificate)

	return out0, err

}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_Facade *FacadeSession) GetCertificateBySecurityNo(securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	return _Facade.Contract.GetCertificateBySecurityNo(&_Facade.CallOpts, securityNo)
}

// GetCertificateBySecurityNo is a free data retrieval call binding the contract method 0x47a76876.
//
// Solidity: function getCertificateBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16)[])
func (_Facade *FacadeCallerSession) GetCertificateBySecurityNo(securityNo [32]byte) ([]CertificateDefineCertificate, error) {
	return _Facade.Contract.GetCertificateBySecurityNo(&_Facade.CallOpts, securityNo)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Facade *FacadeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Facade.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Facade *FacadeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Facade.Contract.GetRoleAdmin(&_Facade.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Facade *FacadeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Facade.Contract.GetRoleAdmin(&_Facade.CallOpts, role)
}

// GetWorkExperienceBySecurityNo is a free data retrieval call binding the contract method 0xabaed258.
//
// Solidity: function getWorkExperienceBySecurityNo(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32)[])
func (_Facade *FacadeCaller) GetWorkExperienceBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) ([]WorkExperienceDefineWorkExperience, error) {
	var out []interface{}
	err := _Facade.contract.Call(opts, &out, "getWorkExperienceBySecurityNo", securityNo)

	if err != nil {
		return *new([]WorkExperienceDefineWorkExperience), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkExperienceDefineWorkExperience)).(*[]WorkExperienceDefineWorkExperience)

	return out0, err

}

// GetWorkExperienceBySecurityNo is a free data retrieval call binding the contract method 0xabaed258.
//
// Solidity: function getWorkExperienceBySecurityNo(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32)[])
func (_Facade *FacadeSession) GetWorkExperienceBySecurityNo(securityNo [32]byte) ([]WorkExperienceDefineWorkExperience, error) {
	return _Facade.Contract.GetWorkExperienceBySecurityNo(&_Facade.CallOpts, securityNo)
}

// GetWorkExperienceBySecurityNo is a free data retrieval call binding the contract method 0xabaed258.
//
// Solidity: function getWorkExperienceBySecurityNo(bytes32 securityNo) view returns((uint16,uint16,bool,bytes32)[])
func (_Facade *FacadeCallerSession) GetWorkExperienceBySecurityNo(securityNo [32]byte) ([]WorkExperienceDefineWorkExperience, error) {
	return _Facade.Contract.GetWorkExperienceBySecurityNo(&_Facade.CallOpts, securityNo)
}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_Facade *FacadeCaller) GetWorkerBySecurityNo(opts *bind.CallOpts, securityNo [32]byte) (WorkerDefineWorker, error) {
	var out []interface{}
	err := _Facade.contract.Call(opts, &out, "getWorkerBySecurityNo", securityNo)

	if err != nil {
		return *new(WorkerDefineWorker), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkerDefineWorker)).(*WorkerDefineWorker)

	return out0, err

}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_Facade *FacadeSession) GetWorkerBySecurityNo(securityNo [32]byte) (WorkerDefineWorker, error) {
	return _Facade.Contract.GetWorkerBySecurityNo(&_Facade.CallOpts, securityNo)
}

// GetWorkerBySecurityNo is a free data retrieval call binding the contract method 0xe831ec86.
//
// Solidity: function getWorkerBySecurityNo(bytes32 securityNo) view returns((bytes32,uint16,uint16,bytes32,bool))
func (_Facade *FacadeCallerSession) GetWorkerBySecurityNo(securityNo [32]byte) (WorkerDefineWorker, error) {
	return _Facade.Contract.GetWorkerBySecurityNo(&_Facade.CallOpts, securityNo)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Facade *FacadeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Facade.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Facade *FacadeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Facade.Contract.HasRole(&_Facade.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Facade *FacadeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Facade.Contract.HasRole(&_Facade.CallOpts, role, account)
}

// Ping is a free data retrieval call binding the contract method 0x5c36b186.
//
// Solidity: function ping() pure returns(bytes32)
func (_Facade *FacadeCaller) Ping(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Facade.contract.Call(opts, &out, "ping")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Ping is a free data retrieval call binding the contract method 0x5c36b186.
//
// Solidity: function ping() pure returns(bytes32)
func (_Facade *FacadeSession) Ping() ([32]byte, error) {
	return _Facade.Contract.Ping(&_Facade.CallOpts)
}

// Ping is a free data retrieval call binding the contract method 0x5c36b186.
//
// Solidity: function ping() pure returns(bytes32)
func (_Facade *FacadeCallerSession) Ping() ([32]byte, error) {
	return _Facade.Contract.Ping(&_Facade.CallOpts)
}

// PingByOwner is a free data retrieval call binding the contract method 0xb718e316.
//
// Solidity: function pingByOwner() view returns(bytes32)
func (_Facade *FacadeCaller) PingByOwner(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Facade.contract.Call(opts, &out, "pingByOwner")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PingByOwner is a free data retrieval call binding the contract method 0xb718e316.
//
// Solidity: function pingByOwner() view returns(bytes32)
func (_Facade *FacadeSession) PingByOwner() ([32]byte, error) {
	return _Facade.Contract.PingByOwner(&_Facade.CallOpts)
}

// PingByOwner is a free data retrieval call binding the contract method 0xb718e316.
//
// Solidity: function pingByOwner() view returns(bytes32)
func (_Facade *FacadeCallerSession) PingByOwner() ([32]byte, error) {
	return _Facade.Contract.PingByOwner(&_Facade.CallOpts)
}

// PingByPrivilegedUser is a free data retrieval call binding the contract method 0x8be65c32.
//
// Solidity: function pingByPrivilegedUser() view returns(bytes32)
func (_Facade *FacadeCaller) PingByPrivilegedUser(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Facade.contract.Call(opts, &out, "pingByPrivilegedUser")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PingByPrivilegedUser is a free data retrieval call binding the contract method 0x8be65c32.
//
// Solidity: function pingByPrivilegedUser() view returns(bytes32)
func (_Facade *FacadeSession) PingByPrivilegedUser() ([32]byte, error) {
	return _Facade.Contract.PingByPrivilegedUser(&_Facade.CallOpts)
}

// PingByPrivilegedUser is a free data retrieval call binding the contract method 0x8be65c32.
//
// Solidity: function pingByPrivilegedUser() view returns(bytes32)
func (_Facade *FacadeCallerSession) PingByPrivilegedUser() ([32]byte, error) {
	return _Facade.Contract.PingByPrivilegedUser(&_Facade.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Facade *FacadeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Facade.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Facade *FacadeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Facade.Contract.SupportsInterface(&_Facade.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Facade *FacadeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Facade.Contract.SupportsInterface(&_Facade.CallOpts, interfaceId)
}

// AddWorkExperience is a paid mutator transaction binding the contract method 0x36efe55e.
//
// Solidity: function addWorkExperience(bytes32 securityNo, (uint16,uint16,bool,bytes32) workExperience) returns()
func (_Facade *FacadeTransactor) AddWorkExperience(opts *bind.TransactOpts, securityNo [32]byte, workExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _Facade.contract.Transact(opts, "addWorkExperience", securityNo, workExperience)
}

// AddWorkExperience is a paid mutator transaction binding the contract method 0x36efe55e.
//
// Solidity: function addWorkExperience(bytes32 securityNo, (uint16,uint16,bool,bytes32) workExperience) returns()
func (_Facade *FacadeSession) AddWorkExperience(securityNo [32]byte, workExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _Facade.Contract.AddWorkExperience(&_Facade.TransactOpts, securityNo, workExperience)
}

// AddWorkExperience is a paid mutator transaction binding the contract method 0x36efe55e.
//
// Solidity: function addWorkExperience(bytes32 securityNo, (uint16,uint16,bool,bytes32) workExperience) returns()
func (_Facade *FacadeTransactorSession) AddWorkExperience(securityNo [32]byte, workExperience WorkExperienceDefineWorkExperience) (*types.Transaction, error) {
	return _Facade.Contract.AddWorkExperience(&_Facade.TransactOpts, securityNo, workExperience)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_Facade *FacadeTransactor) CreateCertificate(opts *bind.TransactOpts, securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _Facade.contract.Transact(opts, "createCertificate", securityNo, certifcate)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_Facade *FacadeSession) CreateCertificate(securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _Facade.Contract.CreateCertificate(&_Facade.TransactOpts, securityNo, certifcate)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xfca8af7c.
//
// Solidity: function createCertificate(bytes32 securityNo, (bytes32,uint16) certifcate) returns()
func (_Facade *FacadeTransactorSession) CreateCertificate(securityNo [32]byte, certifcate CertificateDefineCertificate) (*types.Transaction, error) {
	return _Facade.Contract.CreateCertificate(&_Facade.TransactOpts, securityNo, certifcate)
}

// CreateUser is a paid mutator transaction binding the contract method 0xcdd87618.
//
// Solidity: function createUser(address addr) returns()
func (_Facade *FacadeTransactor) CreateUser(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Facade.contract.Transact(opts, "createUser", addr)
}

// CreateUser is a paid mutator transaction binding the contract method 0xcdd87618.
//
// Solidity: function createUser(address addr) returns()
func (_Facade *FacadeSession) CreateUser(addr common.Address) (*types.Transaction, error) {
	return _Facade.Contract.CreateUser(&_Facade.TransactOpts, addr)
}

// CreateUser is a paid mutator transaction binding the contract method 0xcdd87618.
//
// Solidity: function createUser(address addr) returns()
func (_Facade *FacadeTransactorSession) CreateUser(addr common.Address) (*types.Transaction, error) {
	return _Facade.Contract.CreateUser(&_Facade.TransactOpts, addr)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_Facade *FacadeTransactor) CreateWorker(opts *bind.TransactOpts, worker WorkerDefineWorker) (*types.Transaction, error) {
	return _Facade.contract.Transact(opts, "createWorker", worker)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_Facade *FacadeSession) CreateWorker(worker WorkerDefineWorker) (*types.Transaction, error) {
	return _Facade.Contract.CreateWorker(&_Facade.TransactOpts, worker)
}

// CreateWorker is a paid mutator transaction binding the contract method 0x8d7b1607.
//
// Solidity: function createWorker((bytes32,uint16,uint16,bytes32,bool) worker) returns()
func (_Facade *FacadeTransactorSession) CreateWorker(worker WorkerDefineWorker) (*types.Transaction, error) {
	return _Facade.Contract.CreateWorker(&_Facade.TransactOpts, worker)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_Facade *FacadeTransactor) FinishLastCareer(opts *bind.TransactOpts, securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _Facade.contract.Transact(opts, "finishLastCareer", securityNo, endYear)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_Facade *FacadeSession) FinishLastCareer(securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _Facade.Contract.FinishLastCareer(&_Facade.TransactOpts, securityNo, endYear)
}

// FinishLastCareer is a paid mutator transaction binding the contract method 0x6781e4a1.
//
// Solidity: function finishLastCareer(bytes32 securityNo, uint16 endYear) returns()
func (_Facade *FacadeTransactorSession) FinishLastCareer(securityNo [32]byte, endYear uint16) (*types.Transaction, error) {
	return _Facade.Contract.FinishLastCareer(&_Facade.TransactOpts, securityNo, endYear)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Facade *FacadeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Facade.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Facade *FacadeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Facade.Contract.GrantRole(&_Facade.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Facade *FacadeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Facade.Contract.GrantRole(&_Facade.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dispatcher) returns()
func (_Facade *FacadeTransactor) Initialize(opts *bind.TransactOpts, _dispatcher common.Address) (*types.Transaction, error) {
	return _Facade.contract.Transact(opts, "initialize", _dispatcher)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dispatcher) returns()
func (_Facade *FacadeSession) Initialize(_dispatcher common.Address) (*types.Transaction, error) {
	return _Facade.Contract.Initialize(&_Facade.TransactOpts, _dispatcher)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dispatcher) returns()
func (_Facade *FacadeTransactorSession) Initialize(_dispatcher common.Address) (*types.Transaction, error) {
	return _Facade.Contract.Initialize(&_Facade.TransactOpts, _dispatcher)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address addr) returns()
func (_Facade *FacadeTransactor) RemoveUser(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Facade.contract.Transact(opts, "removeUser", addr)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address addr) returns()
func (_Facade *FacadeSession) RemoveUser(addr common.Address) (*types.Transaction, error) {
	return _Facade.Contract.RemoveUser(&_Facade.TransactOpts, addr)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address addr) returns()
func (_Facade *FacadeTransactorSession) RemoveUser(addr common.Address) (*types.Transaction, error) {
	return _Facade.Contract.RemoveUser(&_Facade.TransactOpts, addr)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Facade *FacadeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Facade.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Facade *FacadeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Facade.Contract.RenounceRole(&_Facade.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Facade *FacadeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Facade.Contract.RenounceRole(&_Facade.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Facade *FacadeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Facade.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Facade *FacadeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Facade.Contract.RevokeRole(&_Facade.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Facade *FacadeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Facade.Contract.RevokeRole(&_Facade.TransactOpts, role, account)
}

// FacadeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Facade contract.
type FacadeInitializedIterator struct {
	Event *FacadeInitialized // Event containing the contract specifics and raw log

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
func (it *FacadeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FacadeInitialized)
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
		it.Event = new(FacadeInitialized)
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
func (it *FacadeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FacadeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FacadeInitialized represents a Initialized event raised by the Facade contract.
type FacadeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Facade *FacadeFilterer) FilterInitialized(opts *bind.FilterOpts) (*FacadeInitializedIterator, error) {

	logs, sub, err := _Facade.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FacadeInitializedIterator{contract: _Facade.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Facade *FacadeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FacadeInitialized) (event.Subscription, error) {

	logs, sub, err := _Facade.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FacadeInitialized)
				if err := _Facade.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Facade *FacadeFilterer) ParseInitialized(log types.Log) (*FacadeInitialized, error) {
	event := new(FacadeInitialized)
	if err := _Facade.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FacadeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Facade contract.
type FacadeRoleAdminChangedIterator struct {
	Event *FacadeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FacadeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FacadeRoleAdminChanged)
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
		it.Event = new(FacadeRoleAdminChanged)
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
func (it *FacadeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FacadeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FacadeRoleAdminChanged represents a RoleAdminChanged event raised by the Facade contract.
type FacadeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Facade *FacadeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FacadeRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Facade.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FacadeRoleAdminChangedIterator{contract: _Facade.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Facade *FacadeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FacadeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Facade.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FacadeRoleAdminChanged)
				if err := _Facade.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Facade *FacadeFilterer) ParseRoleAdminChanged(log types.Log) (*FacadeRoleAdminChanged, error) {
	event := new(FacadeRoleAdminChanged)
	if err := _Facade.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FacadeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Facade contract.
type FacadeRoleGrantedIterator struct {
	Event *FacadeRoleGranted // Event containing the contract specifics and raw log

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
func (it *FacadeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FacadeRoleGranted)
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
		it.Event = new(FacadeRoleGranted)
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
func (it *FacadeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FacadeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FacadeRoleGranted represents a RoleGranted event raised by the Facade contract.
type FacadeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Facade *FacadeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FacadeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Facade.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FacadeRoleGrantedIterator{contract: _Facade.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Facade *FacadeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FacadeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Facade.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FacadeRoleGranted)
				if err := _Facade.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Facade *FacadeFilterer) ParseRoleGranted(log types.Log) (*FacadeRoleGranted, error) {
	event := new(FacadeRoleGranted)
	if err := _Facade.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FacadeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Facade contract.
type FacadeRoleRevokedIterator struct {
	Event *FacadeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FacadeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FacadeRoleRevoked)
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
		it.Event = new(FacadeRoleRevoked)
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
func (it *FacadeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FacadeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FacadeRoleRevoked represents a RoleRevoked event raised by the Facade contract.
type FacadeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Facade *FacadeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FacadeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Facade.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FacadeRoleRevokedIterator{contract: _Facade.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Facade *FacadeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FacadeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Facade.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FacadeRoleRevoked)
				if err := _Facade.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Facade *FacadeFilterer) ParseRoleRevoked(log types.Log) (*FacadeRoleRevoked, error) {
	event := new(FacadeRoleRevoked)
	if err := _Facade.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

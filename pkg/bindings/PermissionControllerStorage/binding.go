// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PermissionControllerStorage

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
	_ = abi.ConvertType
)

// PermissionControllerStorageMetaData contains all meta data concerning the PermissionControllerStorage contract.
var PermissionControllerStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canCall\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdmins\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppointeePermissions\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAppointees\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getPendingAdmins\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pendingAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAppointee\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removePendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAppointee\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AdminRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AdminSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppointeeRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppointeeSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingAdminAdded\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingAdminRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AdminAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminNotPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppointeeAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppointeeNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotHaveZeroAdmins\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAdmin\",\"inputs\":[]}]",
}

// PermissionControllerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use PermissionControllerStorageMetaData.ABI instead.
var PermissionControllerStorageABI = PermissionControllerStorageMetaData.ABI

// PermissionControllerStorage is an auto generated Go binding around an Ethereum contract.
type PermissionControllerStorage struct {
	PermissionControllerStorageCaller     // Read-only binding to the contract
	PermissionControllerStorageTransactor // Write-only binding to the contract
	PermissionControllerStorageFilterer   // Log filterer for contract events
}

// PermissionControllerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionControllerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionControllerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionControllerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionControllerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionControllerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionControllerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionControllerStorageSession struct {
	Contract     *PermissionControllerStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PermissionControllerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionControllerStorageCallerSession struct {
	Contract *PermissionControllerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// PermissionControllerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionControllerStorageTransactorSession struct {
	Contract     *PermissionControllerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// PermissionControllerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionControllerStorageRaw struct {
	Contract *PermissionControllerStorage // Generic contract binding to access the raw methods on
}

// PermissionControllerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionControllerStorageCallerRaw struct {
	Contract *PermissionControllerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionControllerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionControllerStorageTransactorRaw struct {
	Contract *PermissionControllerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissionControllerStorage creates a new instance of PermissionControllerStorage, bound to a specific deployed contract.
func NewPermissionControllerStorage(address common.Address, backend bind.ContractBackend) (*PermissionControllerStorage, error) {
	contract, err := bindPermissionControllerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerStorage{PermissionControllerStorageCaller: PermissionControllerStorageCaller{contract: contract}, PermissionControllerStorageTransactor: PermissionControllerStorageTransactor{contract: contract}, PermissionControllerStorageFilterer: PermissionControllerStorageFilterer{contract: contract}}, nil
}

// NewPermissionControllerStorageCaller creates a new read-only instance of PermissionControllerStorage, bound to a specific deployed contract.
func NewPermissionControllerStorageCaller(address common.Address, caller bind.ContractCaller) (*PermissionControllerStorageCaller, error) {
	contract, err := bindPermissionControllerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerStorageCaller{contract: contract}, nil
}

// NewPermissionControllerStorageTransactor creates a new write-only instance of PermissionControllerStorage, bound to a specific deployed contract.
func NewPermissionControllerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionControllerStorageTransactor, error) {
	contract, err := bindPermissionControllerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerStorageTransactor{contract: contract}, nil
}

// NewPermissionControllerStorageFilterer creates a new log filterer instance of PermissionControllerStorage, bound to a specific deployed contract.
func NewPermissionControllerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionControllerStorageFilterer, error) {
	contract, err := bindPermissionControllerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerStorageFilterer{contract: contract}, nil
}

// bindPermissionControllerStorage binds a generic wrapper to an already deployed contract.
func bindPermissionControllerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PermissionControllerStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionControllerStorage *PermissionControllerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionControllerStorage.Contract.PermissionControllerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionControllerStorage *PermissionControllerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.PermissionControllerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionControllerStorage *PermissionControllerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.PermissionControllerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionControllerStorage *PermissionControllerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionControllerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionControllerStorage *PermissionControllerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionControllerStorage *PermissionControllerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.contract.Transact(opts, method, params...)
}

// GetAdmins is a free data retrieval call binding the contract method 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (_PermissionControllerStorage *PermissionControllerStorageCaller) GetAdmins(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _PermissionControllerStorage.contract.Call(opts, &out, "getAdmins", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (_PermissionControllerStorage *PermissionControllerStorageSession) GetAdmins(account common.Address) ([]common.Address, error) {
	return _PermissionControllerStorage.Contract.GetAdmins(&_PermissionControllerStorage.CallOpts, account)
}

// GetAdmins is a free data retrieval call binding the contract method 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (_PermissionControllerStorage *PermissionControllerStorageCallerSession) GetAdmins(account common.Address) ([]common.Address, error) {
	return _PermissionControllerStorage.Contract.GetAdmins(&_PermissionControllerStorage.CallOpts, account)
}

// GetPendingAdmins is a free data retrieval call binding the contract method 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (_PermissionControllerStorage *PermissionControllerStorageCaller) GetPendingAdmins(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _PermissionControllerStorage.contract.Call(opts, &out, "getPendingAdmins", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPendingAdmins is a free data retrieval call binding the contract method 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (_PermissionControllerStorage *PermissionControllerStorageSession) GetPendingAdmins(account common.Address) ([]common.Address, error) {
	return _PermissionControllerStorage.Contract.GetPendingAdmins(&_PermissionControllerStorage.CallOpts, account)
}

// GetPendingAdmins is a free data retrieval call binding the contract method 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (_PermissionControllerStorage *PermissionControllerStorageCallerSession) GetPendingAdmins(account common.Address) ([]common.Address, error) {
	return _PermissionControllerStorage.Contract.GetPendingAdmins(&_PermissionControllerStorage.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (_PermissionControllerStorage *PermissionControllerStorageCaller) IsAdmin(opts *bind.CallOpts, account common.Address, caller common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionControllerStorage.contract.Call(opts, &out, "isAdmin", account, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (_PermissionControllerStorage *PermissionControllerStorageSession) IsAdmin(account common.Address, caller common.Address) (bool, error) {
	return _PermissionControllerStorage.Contract.IsAdmin(&_PermissionControllerStorage.CallOpts, account, caller)
}

// IsAdmin is a free data retrieval call binding the contract method 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (_PermissionControllerStorage *PermissionControllerStorageCallerSession) IsAdmin(account common.Address, caller common.Address) (bool, error) {
	return _PermissionControllerStorage.Contract.IsAdmin(&_PermissionControllerStorage.CallOpts, account, caller)
}

// IsPendingAdmin is a free data retrieval call binding the contract method 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (_PermissionControllerStorage *PermissionControllerStorageCaller) IsPendingAdmin(opts *bind.CallOpts, account common.Address, pendingAdmin common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionControllerStorage.contract.Call(opts, &out, "isPendingAdmin", account, pendingAdmin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingAdmin is a free data retrieval call binding the contract method 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (_PermissionControllerStorage *PermissionControllerStorageSession) IsPendingAdmin(account common.Address, pendingAdmin common.Address) (bool, error) {
	return _PermissionControllerStorage.Contract.IsPendingAdmin(&_PermissionControllerStorage.CallOpts, account, pendingAdmin)
}

// IsPendingAdmin is a free data retrieval call binding the contract method 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (_PermissionControllerStorage *PermissionControllerStorageCallerSession) IsPendingAdmin(account common.Address, pendingAdmin common.Address) (bool, error) {
	return _PermissionControllerStorage.Contract.IsPendingAdmin(&_PermissionControllerStorage.CallOpts, account, pendingAdmin)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_PermissionControllerStorage *PermissionControllerStorageCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PermissionControllerStorage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_PermissionControllerStorage *PermissionControllerStorageSession) Version() (string, error) {
	return _PermissionControllerStorage.Contract.Version(&_PermissionControllerStorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_PermissionControllerStorage *PermissionControllerStorageCallerSession) Version() (string, error) {
	return _PermissionControllerStorage.Contract.Version(&_PermissionControllerStorage.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address account) returns()
func (_PermissionControllerStorage *PermissionControllerStorageTransactor) AcceptAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.contract.Transact(opts, "acceptAdmin", account)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address account) returns()
func (_PermissionControllerStorage *PermissionControllerStorageSession) AcceptAdmin(account common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.AcceptAdmin(&_PermissionControllerStorage.TransactOpts, account)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address account) returns()
func (_PermissionControllerStorage *PermissionControllerStorageTransactorSession) AcceptAdmin(account common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.AcceptAdmin(&_PermissionControllerStorage.TransactOpts, account)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0xeb5a4e87.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (_PermissionControllerStorage *PermissionControllerStorageTransactor) AddPendingAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.contract.Transact(opts, "addPendingAdmin", account, admin)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0xeb5a4e87.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (_PermissionControllerStorage *PermissionControllerStorageSession) AddPendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.AddPendingAdmin(&_PermissionControllerStorage.TransactOpts, account, admin)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0xeb5a4e87.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (_PermissionControllerStorage *PermissionControllerStorageTransactorSession) AddPendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.AddPendingAdmin(&_PermissionControllerStorage.TransactOpts, account, admin)
}

// CanCall is a paid mutator transaction binding the contract method 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) returns(bool)
func (_PermissionControllerStorage *PermissionControllerStorageTransactor) CanCall(opts *bind.TransactOpts, account common.Address, caller common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionControllerStorage.contract.Transact(opts, "canCall", account, caller, target, selector)
}

// CanCall is a paid mutator transaction binding the contract method 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) returns(bool)
func (_PermissionControllerStorage *PermissionControllerStorageSession) CanCall(account common.Address, caller common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.CanCall(&_PermissionControllerStorage.TransactOpts, account, caller, target, selector)
}

// CanCall is a paid mutator transaction binding the contract method 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) returns(bool)
func (_PermissionControllerStorage *PermissionControllerStorageTransactorSession) CanCall(account common.Address, caller common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.CanCall(&_PermissionControllerStorage.TransactOpts, account, caller, target, selector)
}

// GetAppointeePermissions is a paid mutator transaction binding the contract method 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) returns(address[], bytes4[])
func (_PermissionControllerStorage *PermissionControllerStorageTransactor) GetAppointeePermissions(opts *bind.TransactOpts, account common.Address, appointee common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.contract.Transact(opts, "getAppointeePermissions", account, appointee)
}

// GetAppointeePermissions is a paid mutator transaction binding the contract method 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) returns(address[], bytes4[])
func (_PermissionControllerStorage *PermissionControllerStorageSession) GetAppointeePermissions(account common.Address, appointee common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.GetAppointeePermissions(&_PermissionControllerStorage.TransactOpts, account, appointee)
}

// GetAppointeePermissions is a paid mutator transaction binding the contract method 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) returns(address[], bytes4[])
func (_PermissionControllerStorage *PermissionControllerStorageTransactorSession) GetAppointeePermissions(account common.Address, appointee common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.GetAppointeePermissions(&_PermissionControllerStorage.TransactOpts, account, appointee)
}

// GetAppointees is a paid mutator transaction binding the contract method 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) returns(address[])
func (_PermissionControllerStorage *PermissionControllerStorageTransactor) GetAppointees(opts *bind.TransactOpts, account common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionControllerStorage.contract.Transact(opts, "getAppointees", account, target, selector)
}

// GetAppointees is a paid mutator transaction binding the contract method 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) returns(address[])
func (_PermissionControllerStorage *PermissionControllerStorageSession) GetAppointees(account common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.GetAppointees(&_PermissionControllerStorage.TransactOpts, account, target, selector)
}

// GetAppointees is a paid mutator transaction binding the contract method 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) returns(address[])
func (_PermissionControllerStorage *PermissionControllerStorageTransactorSession) GetAppointees(account common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.GetAppointees(&_PermissionControllerStorage.TransactOpts, account, target, selector)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (_PermissionControllerStorage *PermissionControllerStorageTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.contract.Transact(opts, "removeAdmin", account, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (_PermissionControllerStorage *PermissionControllerStorageSession) RemoveAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.RemoveAdmin(&_PermissionControllerStorage.TransactOpts, account, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (_PermissionControllerStorage *PermissionControllerStorageTransactorSession) RemoveAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.RemoveAdmin(&_PermissionControllerStorage.TransactOpts, account, admin)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0x06641201.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionControllerStorage *PermissionControllerStorageTransactor) RemoveAppointee(opts *bind.TransactOpts, account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionControllerStorage.contract.Transact(opts, "removeAppointee", account, appointee, target, selector)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0x06641201.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionControllerStorage *PermissionControllerStorageSession) RemoveAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.RemoveAppointee(&_PermissionControllerStorage.TransactOpts, account, appointee, target, selector)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0x06641201.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionControllerStorage *PermissionControllerStorageTransactorSession) RemoveAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.RemoveAppointee(&_PermissionControllerStorage.TransactOpts, account, appointee, target, selector)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x4f906cf9.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (_PermissionControllerStorage *PermissionControllerStorageTransactor) RemovePendingAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.contract.Transact(opts, "removePendingAdmin", account, admin)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x4f906cf9.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (_PermissionControllerStorage *PermissionControllerStorageSession) RemovePendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.RemovePendingAdmin(&_PermissionControllerStorage.TransactOpts, account, admin)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x4f906cf9.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (_PermissionControllerStorage *PermissionControllerStorageTransactorSession) RemovePendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.RemovePendingAdmin(&_PermissionControllerStorage.TransactOpts, account, admin)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x950d806e.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionControllerStorage *PermissionControllerStorageTransactor) SetAppointee(opts *bind.TransactOpts, account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionControllerStorage.contract.Transact(opts, "setAppointee", account, appointee, target, selector)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x950d806e.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionControllerStorage *PermissionControllerStorageSession) SetAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.SetAppointee(&_PermissionControllerStorage.TransactOpts, account, appointee, target, selector)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x950d806e.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_PermissionControllerStorage *PermissionControllerStorageTransactorSession) SetAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _PermissionControllerStorage.Contract.SetAppointee(&_PermissionControllerStorage.TransactOpts, account, appointee, target, selector)
}

// PermissionControllerStorageAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the PermissionControllerStorage contract.
type PermissionControllerStorageAdminRemovedIterator struct {
	Event *PermissionControllerStorageAdminRemoved // Event containing the contract specifics and raw log

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
func (it *PermissionControllerStorageAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerStorageAdminRemoved)
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
		it.Event = new(PermissionControllerStorageAdminRemoved)
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
func (it *PermissionControllerStorageAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerStorageAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerStorageAdminRemoved represents a AdminRemoved event raised by the PermissionControllerStorage contract.
type PermissionControllerStorageAdminRemoved struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*PermissionControllerStorageAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionControllerStorage.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerStorageAdminRemovedIterator{contract: _PermissionControllerStorage.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *PermissionControllerStorageAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionControllerStorage.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerStorageAdminRemoved)
				if err := _PermissionControllerStorage.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) ParseAdminRemoved(log types.Log) (*PermissionControllerStorageAdminRemoved, error) {
	event := new(PermissionControllerStorageAdminRemoved)
	if err := _PermissionControllerStorage.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerStorageAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the PermissionControllerStorage contract.
type PermissionControllerStorageAdminSetIterator struct {
	Event *PermissionControllerStorageAdminSet // Event containing the contract specifics and raw log

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
func (it *PermissionControllerStorageAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerStorageAdminSet)
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
		it.Event = new(PermissionControllerStorageAdminSet)
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
func (it *PermissionControllerStorageAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerStorageAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerStorageAdminSet represents a AdminSet event raised by the PermissionControllerStorage contract.
type PermissionControllerStorageAdminSet struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0xbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff97.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) FilterAdminSet(opts *bind.FilterOpts, account []common.Address) (*PermissionControllerStorageAdminSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionControllerStorage.contract.FilterLogs(opts, "AdminSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerStorageAdminSetIterator{contract: _PermissionControllerStorage.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0xbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff97.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *PermissionControllerStorageAdminSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionControllerStorage.contract.WatchLogs(opts, "AdminSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerStorageAdminSet)
				if err := _PermissionControllerStorage.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

// ParseAdminSet is a log parse operation binding the contract event 0xbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff97.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) ParseAdminSet(log types.Log) (*PermissionControllerStorageAdminSet, error) {
	event := new(PermissionControllerStorageAdminSet)
	if err := _PermissionControllerStorage.contract.UnpackLog(event, "AdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerStorageAppointeeRemovedIterator is returned from FilterAppointeeRemoved and is used to iterate over the raw logs and unpacked data for AppointeeRemoved events raised by the PermissionControllerStorage contract.
type PermissionControllerStorageAppointeeRemovedIterator struct {
	Event *PermissionControllerStorageAppointeeRemoved // Event containing the contract specifics and raw log

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
func (it *PermissionControllerStorageAppointeeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerStorageAppointeeRemoved)
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
		it.Event = new(PermissionControllerStorageAppointeeRemoved)
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
func (it *PermissionControllerStorageAppointeeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerStorageAppointeeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerStorageAppointeeRemoved represents a AppointeeRemoved event raised by the PermissionControllerStorage contract.
type PermissionControllerStorageAppointeeRemoved struct {
	Account   common.Address
	Appointee common.Address
	Target    common.Address
	Selector  [4]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppointeeRemoved is a free log retrieval operation binding the contract event 0x18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) FilterAppointeeRemoved(opts *bind.FilterOpts, account []common.Address, appointee []common.Address) (*PermissionControllerStorageAppointeeRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _PermissionControllerStorage.contract.FilterLogs(opts, "AppointeeRemoved", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerStorageAppointeeRemovedIterator{contract: _PermissionControllerStorage.contract, event: "AppointeeRemoved", logs: logs, sub: sub}, nil
}

// WatchAppointeeRemoved is a free log subscription operation binding the contract event 0x18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) WatchAppointeeRemoved(opts *bind.WatchOpts, sink chan<- *PermissionControllerStorageAppointeeRemoved, account []common.Address, appointee []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _PermissionControllerStorage.contract.WatchLogs(opts, "AppointeeRemoved", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerStorageAppointeeRemoved)
				if err := _PermissionControllerStorage.contract.UnpackLog(event, "AppointeeRemoved", log); err != nil {
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

// ParseAppointeeRemoved is a log parse operation binding the contract event 0x18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) ParseAppointeeRemoved(log types.Log) (*PermissionControllerStorageAppointeeRemoved, error) {
	event := new(PermissionControllerStorageAppointeeRemoved)
	if err := _PermissionControllerStorage.contract.UnpackLog(event, "AppointeeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerStorageAppointeeSetIterator is returned from FilterAppointeeSet and is used to iterate over the raw logs and unpacked data for AppointeeSet events raised by the PermissionControllerStorage contract.
type PermissionControllerStorageAppointeeSetIterator struct {
	Event *PermissionControllerStorageAppointeeSet // Event containing the contract specifics and raw log

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
func (it *PermissionControllerStorageAppointeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerStorageAppointeeSet)
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
		it.Event = new(PermissionControllerStorageAppointeeSet)
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
func (it *PermissionControllerStorageAppointeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerStorageAppointeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerStorageAppointeeSet represents a AppointeeSet event raised by the PermissionControllerStorage contract.
type PermissionControllerStorageAppointeeSet struct {
	Account   common.Address
	Appointee common.Address
	Target    common.Address
	Selector  [4]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppointeeSet is a free log retrieval operation binding the contract event 0x037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) FilterAppointeeSet(opts *bind.FilterOpts, account []common.Address, appointee []common.Address) (*PermissionControllerStorageAppointeeSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _PermissionControllerStorage.contract.FilterLogs(opts, "AppointeeSet", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerStorageAppointeeSetIterator{contract: _PermissionControllerStorage.contract, event: "AppointeeSet", logs: logs, sub: sub}, nil
}

// WatchAppointeeSet is a free log subscription operation binding the contract event 0x037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) WatchAppointeeSet(opts *bind.WatchOpts, sink chan<- *PermissionControllerStorageAppointeeSet, account []common.Address, appointee []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _PermissionControllerStorage.contract.WatchLogs(opts, "AppointeeSet", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerStorageAppointeeSet)
				if err := _PermissionControllerStorage.contract.UnpackLog(event, "AppointeeSet", log); err != nil {
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

// ParseAppointeeSet is a log parse operation binding the contract event 0x037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) ParseAppointeeSet(log types.Log) (*PermissionControllerStorageAppointeeSet, error) {
	event := new(PermissionControllerStorageAppointeeSet)
	if err := _PermissionControllerStorage.contract.UnpackLog(event, "AppointeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerStoragePendingAdminAddedIterator is returned from FilterPendingAdminAdded and is used to iterate over the raw logs and unpacked data for PendingAdminAdded events raised by the PermissionControllerStorage contract.
type PermissionControllerStoragePendingAdminAddedIterator struct {
	Event *PermissionControllerStoragePendingAdminAdded // Event containing the contract specifics and raw log

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
func (it *PermissionControllerStoragePendingAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerStoragePendingAdminAdded)
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
		it.Event = new(PermissionControllerStoragePendingAdminAdded)
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
func (it *PermissionControllerStoragePendingAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerStoragePendingAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerStoragePendingAdminAdded represents a PendingAdminAdded event raised by the PermissionControllerStorage contract.
type PermissionControllerStoragePendingAdminAdded struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPendingAdminAdded is a free log retrieval operation binding the contract event 0xb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) FilterPendingAdminAdded(opts *bind.FilterOpts, account []common.Address) (*PermissionControllerStoragePendingAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionControllerStorage.contract.FilterLogs(opts, "PendingAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerStoragePendingAdminAddedIterator{contract: _PermissionControllerStorage.contract, event: "PendingAdminAdded", logs: logs, sub: sub}, nil
}

// WatchPendingAdminAdded is a free log subscription operation binding the contract event 0xb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) WatchPendingAdminAdded(opts *bind.WatchOpts, sink chan<- *PermissionControllerStoragePendingAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionControllerStorage.contract.WatchLogs(opts, "PendingAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerStoragePendingAdminAdded)
				if err := _PermissionControllerStorage.contract.UnpackLog(event, "PendingAdminAdded", log); err != nil {
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

// ParsePendingAdminAdded is a log parse operation binding the contract event 0xb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) ParsePendingAdminAdded(log types.Log) (*PermissionControllerStoragePendingAdminAdded, error) {
	event := new(PermissionControllerStoragePendingAdminAdded)
	if err := _PermissionControllerStorage.contract.UnpackLog(event, "PendingAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionControllerStoragePendingAdminRemovedIterator is returned from FilterPendingAdminRemoved and is used to iterate over the raw logs and unpacked data for PendingAdminRemoved events raised by the PermissionControllerStorage contract.
type PermissionControllerStoragePendingAdminRemovedIterator struct {
	Event *PermissionControllerStoragePendingAdminRemoved // Event containing the contract specifics and raw log

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
func (it *PermissionControllerStoragePendingAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionControllerStoragePendingAdminRemoved)
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
		it.Event = new(PermissionControllerStoragePendingAdminRemoved)
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
func (it *PermissionControllerStoragePendingAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionControllerStoragePendingAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionControllerStoragePendingAdminRemoved represents a PendingAdminRemoved event raised by the PermissionControllerStorage contract.
type PermissionControllerStoragePendingAdminRemoved struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPendingAdminRemoved is a free log retrieval operation binding the contract event 0xd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) FilterPendingAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*PermissionControllerStoragePendingAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionControllerStorage.contract.FilterLogs(opts, "PendingAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PermissionControllerStoragePendingAdminRemovedIterator{contract: _PermissionControllerStorage.contract, event: "PendingAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchPendingAdminRemoved is a free log subscription operation binding the contract event 0xd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) WatchPendingAdminRemoved(opts *bind.WatchOpts, sink chan<- *PermissionControllerStoragePendingAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PermissionControllerStorage.contract.WatchLogs(opts, "PendingAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionControllerStoragePendingAdminRemoved)
				if err := _PermissionControllerStorage.contract.UnpackLog(event, "PendingAdminRemoved", log); err != nil {
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

// ParsePendingAdminRemoved is a log parse operation binding the contract event 0xd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (_PermissionControllerStorage *PermissionControllerStorageFilterer) ParsePendingAdminRemoved(log types.Log) (*PermissionControllerStoragePendingAdminRemoved, error) {
	event := new(PermissionControllerStoragePendingAdminRemoved)
	if err := _PermissionControllerStorage.contract.UnpackLog(event, "PendingAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

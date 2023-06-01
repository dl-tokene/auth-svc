// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package master_access_management

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

// IRBACResourceWithPermissions is an auto generated low-level Go binding around an user-defined struct.
type IRBACResourceWithPermissions struct {
	Resource    string
	Permissions []string
}

// MasterAccessManagementMetaData contains all meta data concerning the MasterAccessManagement contract.
var MasterAccessManagementMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"permissionsToAdd\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AddedPermissions\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"rolesToGrant\",\"type\":\"string[]\"}],\"name\":\"GrantedRoles\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"permissionsToRemove\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"RemovedPermissions\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"rolesToRevoke\",\"type\":\"string[]\"}],\"name\":\"RevokedRoles\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ALL_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ALL_RESOURCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSTANTS_REGISTRY_RESOURCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CREATE_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELETE_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTE_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_REGISTRY_RESOURCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_ROLE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RBAC_RESOURCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"READ_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REVIEWABLE_REQUESTS_RESOURCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"master_\",\"type\":\"address\"}],\"name\":\"__MasterAccessManagement_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"permissions\",\"type\":\"string[]\"}],\"internalType\":\"structIRBAC.ResourceWithPermissions[]\",\"name\":\"permissionsToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"addPermissionsToRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"name\":\"getRolePermissions\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"permissions\",\"type\":\"string[]\"}],\"internalType\":\"structIRBAC.ResourceWithPermissions[]\",\"name\":\"allowed\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"permissions\",\"type\":\"string[]\"}],\"internalType\":\"structIRBAC.ResourceWithPermissions[]\",\"name\":\"disallowed\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"getUserRoles\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"roles\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"rolesToGrant\",\"type\":\"string[]\"}],\"name\":\"grantRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasConstantsRegistryCreatePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasConstantsRegistryDeletePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasMasterContractsRegistryCreatePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasMasterContractsRegistryDeletePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasMasterContractsRegistryUpdatePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"permission\",\"type\":\"string\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasReviewableRequestsCreatePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasReviewableRequestsDeletePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasReviewableRequestsExecutePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"permissions\",\"type\":\"string[]\"}],\"internalType\":\"structIRBAC.ResourceWithPermissions[]\",\"name\":\"permissionsToRemove\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"removePermissionsFromRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"rolesToRevoke\",\"type\":\"string[]\"}],\"name\":\"revokeRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MasterAccessManagementABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterAccessManagementMetaData.ABI instead.
var MasterAccessManagementABI = MasterAccessManagementMetaData.ABI

// MasterAccessManagement is an auto generated Go binding around an Ethereum contract.
type MasterAccessManagement struct {
	MasterAccessManagementCaller     // Read-only binding to the contract
	MasterAccessManagementTransactor // Write-only binding to the contract
	MasterAccessManagementFilterer   // Log filterer for contract events
}

// MasterAccessManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type MasterAccessManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterAccessManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterAccessManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterAccessManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterAccessManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterAccessManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterAccessManagementSession struct {
	Contract     *MasterAccessManagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MasterAccessManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterAccessManagementCallerSession struct {
	Contract *MasterAccessManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MasterAccessManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterAccessManagementTransactorSession struct {
	Contract     *MasterAccessManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MasterAccessManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type MasterAccessManagementRaw struct {
	Contract *MasterAccessManagement // Generic contract binding to access the raw methods on
}

// MasterAccessManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterAccessManagementCallerRaw struct {
	Contract *MasterAccessManagementCaller // Generic read-only contract binding to access the raw methods on
}

// MasterAccessManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterAccessManagementTransactorRaw struct {
	Contract *MasterAccessManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterAccessManagement creates a new instance of MasterAccessManagement, bound to a specific deployed contract.
func NewMasterAccessManagement(address common.Address, backend bind.ContractBackend) (*MasterAccessManagement, error) {
	contract, err := bindMasterAccessManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasterAccessManagement{MasterAccessManagementCaller: MasterAccessManagementCaller{contract: contract}, MasterAccessManagementTransactor: MasterAccessManagementTransactor{contract: contract}, MasterAccessManagementFilterer: MasterAccessManagementFilterer{contract: contract}}, nil
}

// NewMasterAccessManagementCaller creates a new read-only instance of MasterAccessManagement, bound to a specific deployed contract.
func NewMasterAccessManagementCaller(address common.Address, caller bind.ContractCaller) (*MasterAccessManagementCaller, error) {
	contract, err := bindMasterAccessManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterAccessManagementCaller{contract: contract}, nil
}

// NewMasterAccessManagementTransactor creates a new write-only instance of MasterAccessManagement, bound to a specific deployed contract.
func NewMasterAccessManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*MasterAccessManagementTransactor, error) {
	contract, err := bindMasterAccessManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterAccessManagementTransactor{contract: contract}, nil
}

// NewMasterAccessManagementFilterer creates a new log filterer instance of MasterAccessManagement, bound to a specific deployed contract.
func NewMasterAccessManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*MasterAccessManagementFilterer, error) {
	contract, err := bindMasterAccessManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterAccessManagementFilterer{contract: contract}, nil
}

// bindMasterAccessManagement binds a generic wrapper to an already deployed contract.
func bindMasterAccessManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MasterAccessManagementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterAccessManagement *MasterAccessManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterAccessManagement.Contract.MasterAccessManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterAccessManagement *MasterAccessManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.MasterAccessManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterAccessManagement *MasterAccessManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.MasterAccessManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterAccessManagement *MasterAccessManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterAccessManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterAccessManagement *MasterAccessManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterAccessManagement *MasterAccessManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.contract.Transact(opts, method, params...)
}

// ALLPERMISSION is a free data retrieval call binding the contract method 0xff846fb9.
//
// Solidity: function ALL_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCaller) ALLPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "ALL_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ALLPERMISSION is a free data retrieval call binding the contract method 0xff846fb9.
//
// Solidity: function ALL_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementSession) ALLPERMISSION() (string, error) {
	return _MasterAccessManagement.Contract.ALLPERMISSION(&_MasterAccessManagement.CallOpts)
}

// ALLPERMISSION is a free data retrieval call binding the contract method 0xff846fb9.
//
// Solidity: function ALL_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) ALLPERMISSION() (string, error) {
	return _MasterAccessManagement.Contract.ALLPERMISSION(&_MasterAccessManagement.CallOpts)
}

// ALLRESOURCE is a free data retrieval call binding the contract method 0x4a9e12c5.
//
// Solidity: function ALL_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCaller) ALLRESOURCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "ALL_RESOURCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ALLRESOURCE is a free data retrieval call binding the contract method 0x4a9e12c5.
//
// Solidity: function ALL_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementSession) ALLRESOURCE() (string, error) {
	return _MasterAccessManagement.Contract.ALLRESOURCE(&_MasterAccessManagement.CallOpts)
}

// ALLRESOURCE is a free data retrieval call binding the contract method 0x4a9e12c5.
//
// Solidity: function ALL_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) ALLRESOURCE() (string, error) {
	return _MasterAccessManagement.Contract.ALLRESOURCE(&_MasterAccessManagement.CallOpts)
}

// CONSTANTSREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x6a5efd26.
//
// Solidity: function CONSTANTS_REGISTRY_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCaller) CONSTANTSREGISTRYRESOURCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "CONSTANTS_REGISTRY_RESOURCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CONSTANTSREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x6a5efd26.
//
// Solidity: function CONSTANTS_REGISTRY_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementSession) CONSTANTSREGISTRYRESOURCE() (string, error) {
	return _MasterAccessManagement.Contract.CONSTANTSREGISTRYRESOURCE(&_MasterAccessManagement.CallOpts)
}

// CONSTANTSREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x6a5efd26.
//
// Solidity: function CONSTANTS_REGISTRY_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) CONSTANTSREGISTRYRESOURCE() (string, error) {
	return _MasterAccessManagement.Contract.CONSTANTSREGISTRYRESOURCE(&_MasterAccessManagement.CallOpts)
}

// CREATEPERMISSION is a free data retrieval call binding the contract method 0xb3e657fb.
//
// Solidity: function CREATE_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCaller) CREATEPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "CREATE_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CREATEPERMISSION is a free data retrieval call binding the contract method 0xb3e657fb.
//
// Solidity: function CREATE_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementSession) CREATEPERMISSION() (string, error) {
	return _MasterAccessManagement.Contract.CREATEPERMISSION(&_MasterAccessManagement.CallOpts)
}

// CREATEPERMISSION is a free data retrieval call binding the contract method 0xb3e657fb.
//
// Solidity: function CREATE_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) CREATEPERMISSION() (string, error) {
	return _MasterAccessManagement.Contract.CREATEPERMISSION(&_MasterAccessManagement.CallOpts)
}

// DELETEPERMISSION is a free data retrieval call binding the contract method 0xb832a5a2.
//
// Solidity: function DELETE_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCaller) DELETEPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "DELETE_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DELETEPERMISSION is a free data retrieval call binding the contract method 0xb832a5a2.
//
// Solidity: function DELETE_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementSession) DELETEPERMISSION() (string, error) {
	return _MasterAccessManagement.Contract.DELETEPERMISSION(&_MasterAccessManagement.CallOpts)
}

// DELETEPERMISSION is a free data retrieval call binding the contract method 0xb832a5a2.
//
// Solidity: function DELETE_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) DELETEPERMISSION() (string, error) {
	return _MasterAccessManagement.Contract.DELETEPERMISSION(&_MasterAccessManagement.CallOpts)
}

// EXECUTEPERMISSION is a free data retrieval call binding the contract method 0xbcfa4784.
//
// Solidity: function EXECUTE_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCaller) EXECUTEPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "EXECUTE_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EXECUTEPERMISSION is a free data retrieval call binding the contract method 0xbcfa4784.
//
// Solidity: function EXECUTE_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementSession) EXECUTEPERMISSION() (string, error) {
	return _MasterAccessManagement.Contract.EXECUTEPERMISSION(&_MasterAccessManagement.CallOpts)
}

// EXECUTEPERMISSION is a free data retrieval call binding the contract method 0xbcfa4784.
//
// Solidity: function EXECUTE_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) EXECUTEPERMISSION() (string, error) {
	return _MasterAccessManagement.Contract.EXECUTEPERMISSION(&_MasterAccessManagement.CallOpts)
}

// MASTERREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x5db2bad0.
//
// Solidity: function MASTER_REGISTRY_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCaller) MASTERREGISTRYRESOURCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "MASTER_REGISTRY_RESOURCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MASTERREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x5db2bad0.
//
// Solidity: function MASTER_REGISTRY_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementSession) MASTERREGISTRYRESOURCE() (string, error) {
	return _MasterAccessManagement.Contract.MASTERREGISTRYRESOURCE(&_MasterAccessManagement.CallOpts)
}

// MASTERREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x5db2bad0.
//
// Solidity: function MASTER_REGISTRY_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) MASTERREGISTRYRESOURCE() (string, error) {
	return _MasterAccessManagement.Contract.MASTERREGISTRYRESOURCE(&_MasterAccessManagement.CallOpts)
}

// MASTERROLE is a free data retrieval call binding the contract method 0xdc224863.
//
// Solidity: function MASTER_ROLE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCaller) MASTERROLE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "MASTER_ROLE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MASTERROLE is a free data retrieval call binding the contract method 0xdc224863.
//
// Solidity: function MASTER_ROLE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementSession) MASTERROLE() (string, error) {
	return _MasterAccessManagement.Contract.MASTERROLE(&_MasterAccessManagement.CallOpts)
}

// MASTERROLE is a free data retrieval call binding the contract method 0xdc224863.
//
// Solidity: function MASTER_ROLE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) MASTERROLE() (string, error) {
	return _MasterAccessManagement.Contract.MASTERROLE(&_MasterAccessManagement.CallOpts)
}

// RBACRESOURCE is a free data retrieval call binding the contract method 0x733352b3.
//
// Solidity: function RBAC_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCaller) RBACRESOURCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "RBAC_RESOURCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RBACRESOURCE is a free data retrieval call binding the contract method 0x733352b3.
//
// Solidity: function RBAC_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementSession) RBACRESOURCE() (string, error) {
	return _MasterAccessManagement.Contract.RBACRESOURCE(&_MasterAccessManagement.CallOpts)
}

// RBACRESOURCE is a free data retrieval call binding the contract method 0x733352b3.
//
// Solidity: function RBAC_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) RBACRESOURCE() (string, error) {
	return _MasterAccessManagement.Contract.RBACRESOURCE(&_MasterAccessManagement.CallOpts)
}

// READPERMISSION is a free data retrieval call binding the contract method 0x03bc0b3e.
//
// Solidity: function READ_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCaller) READPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "READ_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// READPERMISSION is a free data retrieval call binding the contract method 0x03bc0b3e.
//
// Solidity: function READ_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementSession) READPERMISSION() (string, error) {
	return _MasterAccessManagement.Contract.READPERMISSION(&_MasterAccessManagement.CallOpts)
}

// READPERMISSION is a free data retrieval call binding the contract method 0x03bc0b3e.
//
// Solidity: function READ_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) READPERMISSION() (string, error) {
	return _MasterAccessManagement.Contract.READPERMISSION(&_MasterAccessManagement.CallOpts)
}

// REVIEWABLEREQUESTSRESOURCE is a free data retrieval call binding the contract method 0xf55303bd.
//
// Solidity: function REVIEWABLE_REQUESTS_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCaller) REVIEWABLEREQUESTSRESOURCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "REVIEWABLE_REQUESTS_RESOURCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// REVIEWABLEREQUESTSRESOURCE is a free data retrieval call binding the contract method 0xf55303bd.
//
// Solidity: function REVIEWABLE_REQUESTS_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementSession) REVIEWABLEREQUESTSRESOURCE() (string, error) {
	return _MasterAccessManagement.Contract.REVIEWABLEREQUESTSRESOURCE(&_MasterAccessManagement.CallOpts)
}

// REVIEWABLEREQUESTSRESOURCE is a free data retrieval call binding the contract method 0xf55303bd.
//
// Solidity: function REVIEWABLE_REQUESTS_RESOURCE() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) REVIEWABLEREQUESTSRESOURCE() (string, error) {
	return _MasterAccessManagement.Contract.REVIEWABLEREQUESTSRESOURCE(&_MasterAccessManagement.CallOpts)
}

// UPDATEPERMISSION is a free data retrieval call binding the contract method 0x0ead6f1e.
//
// Solidity: function UPDATE_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCaller) UPDATEPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "UPDATE_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPDATEPERMISSION is a free data retrieval call binding the contract method 0x0ead6f1e.
//
// Solidity: function UPDATE_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementSession) UPDATEPERMISSION() (string, error) {
	return _MasterAccessManagement.Contract.UPDATEPERMISSION(&_MasterAccessManagement.CallOpts)
}

// UPDATEPERMISSION is a free data retrieval call binding the contract method 0x0ead6f1e.
//
// Solidity: function UPDATE_PERMISSION() view returns(string)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) UPDATEPERMISSION() (string, error) {
	return _MasterAccessManagement.Contract.UPDATEPERMISSION(&_MasterAccessManagement.CallOpts)
}

// GetRolePermissions is a free data retrieval call binding the contract method 0x002f5bc0.
//
// Solidity: function getRolePermissions(string role) view returns((string,string[])[] allowed, (string,string[])[] disallowed)
func (_MasterAccessManagement *MasterAccessManagementCaller) GetRolePermissions(opts *bind.CallOpts, role string) (struct {
	Allowed    []IRBACResourceWithPermissions
	Disallowed []IRBACResourceWithPermissions
}, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "getRolePermissions", role)

	outstruct := new(struct {
		Allowed    []IRBACResourceWithPermissions
		Disallowed []IRBACResourceWithPermissions
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Allowed = *abi.ConvertType(out[0], new([]IRBACResourceWithPermissions)).(*[]IRBACResourceWithPermissions)
	outstruct.Disallowed = *abi.ConvertType(out[1], new([]IRBACResourceWithPermissions)).(*[]IRBACResourceWithPermissions)

	return *outstruct, err

}

// GetRolePermissions is a free data retrieval call binding the contract method 0x002f5bc0.
//
// Solidity: function getRolePermissions(string role) view returns((string,string[])[] allowed, (string,string[])[] disallowed)
func (_MasterAccessManagement *MasterAccessManagementSession) GetRolePermissions(role string) (struct {
	Allowed    []IRBACResourceWithPermissions
	Disallowed []IRBACResourceWithPermissions
}, error) {
	return _MasterAccessManagement.Contract.GetRolePermissions(&_MasterAccessManagement.CallOpts, role)
}

// GetRolePermissions is a free data retrieval call binding the contract method 0x002f5bc0.
//
// Solidity: function getRolePermissions(string role) view returns((string,string[])[] allowed, (string,string[])[] disallowed)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) GetRolePermissions(role string) (struct {
	Allowed    []IRBACResourceWithPermissions
	Disallowed []IRBACResourceWithPermissions
}, error) {
	return _MasterAccessManagement.Contract.GetRolePermissions(&_MasterAccessManagement.CallOpts, role)
}

// GetUserRoles is a free data retrieval call binding the contract method 0x06a36aee.
//
// Solidity: function getUserRoles(address who) view returns(string[] roles)
func (_MasterAccessManagement *MasterAccessManagementCaller) GetUserRoles(opts *bind.CallOpts, who common.Address) ([]string, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "getUserRoles", who)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetUserRoles is a free data retrieval call binding the contract method 0x06a36aee.
//
// Solidity: function getUserRoles(address who) view returns(string[] roles)
func (_MasterAccessManagement *MasterAccessManagementSession) GetUserRoles(who common.Address) ([]string, error) {
	return _MasterAccessManagement.Contract.GetUserRoles(&_MasterAccessManagement.CallOpts, who)
}

// GetUserRoles is a free data retrieval call binding the contract method 0x06a36aee.
//
// Solidity: function getUserRoles(address who) view returns(string[] roles)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) GetUserRoles(who common.Address) ([]string, error) {
	return _MasterAccessManagement.Contract.GetUserRoles(&_MasterAccessManagement.CallOpts, who)
}

// HasConstantsRegistryCreatePermission is a free data retrieval call binding the contract method 0x5b44a636.
//
// Solidity: function hasConstantsRegistryCreatePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCaller) HasConstantsRegistryCreatePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "hasConstantsRegistryCreatePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasConstantsRegistryCreatePermission is a free data retrieval call binding the contract method 0x5b44a636.
//
// Solidity: function hasConstantsRegistryCreatePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementSession) HasConstantsRegistryCreatePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasConstantsRegistryCreatePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasConstantsRegistryCreatePermission is a free data retrieval call binding the contract method 0x5b44a636.
//
// Solidity: function hasConstantsRegistryCreatePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) HasConstantsRegistryCreatePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasConstantsRegistryCreatePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasConstantsRegistryDeletePermission is a free data retrieval call binding the contract method 0x05e10d32.
//
// Solidity: function hasConstantsRegistryDeletePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCaller) HasConstantsRegistryDeletePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "hasConstantsRegistryDeletePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasConstantsRegistryDeletePermission is a free data retrieval call binding the contract method 0x05e10d32.
//
// Solidity: function hasConstantsRegistryDeletePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementSession) HasConstantsRegistryDeletePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasConstantsRegistryDeletePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasConstantsRegistryDeletePermission is a free data retrieval call binding the contract method 0x05e10d32.
//
// Solidity: function hasConstantsRegistryDeletePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) HasConstantsRegistryDeletePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasConstantsRegistryDeletePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasMasterContractsRegistryCreatePermission is a free data retrieval call binding the contract method 0x25536582.
//
// Solidity: function hasMasterContractsRegistryCreatePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCaller) HasMasterContractsRegistryCreatePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "hasMasterContractsRegistryCreatePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMasterContractsRegistryCreatePermission is a free data retrieval call binding the contract method 0x25536582.
//
// Solidity: function hasMasterContractsRegistryCreatePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementSession) HasMasterContractsRegistryCreatePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasMasterContractsRegistryCreatePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasMasterContractsRegistryCreatePermission is a free data retrieval call binding the contract method 0x25536582.
//
// Solidity: function hasMasterContractsRegistryCreatePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) HasMasterContractsRegistryCreatePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasMasterContractsRegistryCreatePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasMasterContractsRegistryDeletePermission is a free data retrieval call binding the contract method 0x7043f0d8.
//
// Solidity: function hasMasterContractsRegistryDeletePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCaller) HasMasterContractsRegistryDeletePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "hasMasterContractsRegistryDeletePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMasterContractsRegistryDeletePermission is a free data retrieval call binding the contract method 0x7043f0d8.
//
// Solidity: function hasMasterContractsRegistryDeletePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementSession) HasMasterContractsRegistryDeletePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasMasterContractsRegistryDeletePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasMasterContractsRegistryDeletePermission is a free data retrieval call binding the contract method 0x7043f0d8.
//
// Solidity: function hasMasterContractsRegistryDeletePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) HasMasterContractsRegistryDeletePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasMasterContractsRegistryDeletePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasMasterContractsRegistryUpdatePermission is a free data retrieval call binding the contract method 0x0aefca57.
//
// Solidity: function hasMasterContractsRegistryUpdatePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCaller) HasMasterContractsRegistryUpdatePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "hasMasterContractsRegistryUpdatePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMasterContractsRegistryUpdatePermission is a free data retrieval call binding the contract method 0x0aefca57.
//
// Solidity: function hasMasterContractsRegistryUpdatePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementSession) HasMasterContractsRegistryUpdatePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasMasterContractsRegistryUpdatePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasMasterContractsRegistryUpdatePermission is a free data retrieval call binding the contract method 0x0aefca57.
//
// Solidity: function hasMasterContractsRegistryUpdatePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) HasMasterContractsRegistryUpdatePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasMasterContractsRegistryUpdatePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasPermission is a free data retrieval call binding the contract method 0x7951c6da.
//
// Solidity: function hasPermission(address who, string resource, string permission) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCaller) HasPermission(opts *bind.CallOpts, who common.Address, resource string, permission string) (bool, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "hasPermission", who, resource, permission)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0x7951c6da.
//
// Solidity: function hasPermission(address who, string resource, string permission) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementSession) HasPermission(who common.Address, resource string, permission string) (bool, error) {
	return _MasterAccessManagement.Contract.HasPermission(&_MasterAccessManagement.CallOpts, who, resource, permission)
}

// HasPermission is a free data retrieval call binding the contract method 0x7951c6da.
//
// Solidity: function hasPermission(address who, string resource, string permission) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) HasPermission(who common.Address, resource string, permission string) (bool, error) {
	return _MasterAccessManagement.Contract.HasPermission(&_MasterAccessManagement.CallOpts, who, resource, permission)
}

// HasReviewableRequestsCreatePermission is a free data retrieval call binding the contract method 0xe6cdf5d0.
//
// Solidity: function hasReviewableRequestsCreatePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCaller) HasReviewableRequestsCreatePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "hasReviewableRequestsCreatePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasReviewableRequestsCreatePermission is a free data retrieval call binding the contract method 0xe6cdf5d0.
//
// Solidity: function hasReviewableRequestsCreatePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementSession) HasReviewableRequestsCreatePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasReviewableRequestsCreatePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasReviewableRequestsCreatePermission is a free data retrieval call binding the contract method 0xe6cdf5d0.
//
// Solidity: function hasReviewableRequestsCreatePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) HasReviewableRequestsCreatePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasReviewableRequestsCreatePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasReviewableRequestsDeletePermission is a free data retrieval call binding the contract method 0x6d19225e.
//
// Solidity: function hasReviewableRequestsDeletePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCaller) HasReviewableRequestsDeletePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "hasReviewableRequestsDeletePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasReviewableRequestsDeletePermission is a free data retrieval call binding the contract method 0x6d19225e.
//
// Solidity: function hasReviewableRequestsDeletePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementSession) HasReviewableRequestsDeletePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasReviewableRequestsDeletePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasReviewableRequestsDeletePermission is a free data retrieval call binding the contract method 0x6d19225e.
//
// Solidity: function hasReviewableRequestsDeletePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) HasReviewableRequestsDeletePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasReviewableRequestsDeletePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasReviewableRequestsExecutePermission is a free data retrieval call binding the contract method 0x20b55d81.
//
// Solidity: function hasReviewableRequestsExecutePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCaller) HasReviewableRequestsExecutePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _MasterAccessManagement.contract.Call(opts, &out, "hasReviewableRequestsExecutePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasReviewableRequestsExecutePermission is a free data retrieval call binding the contract method 0x20b55d81.
//
// Solidity: function hasReviewableRequestsExecutePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementSession) HasReviewableRequestsExecutePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasReviewableRequestsExecutePermission(&_MasterAccessManagement.CallOpts, account_)
}

// HasReviewableRequestsExecutePermission is a free data retrieval call binding the contract method 0x20b55d81.
//
// Solidity: function hasReviewableRequestsExecutePermission(address account_) view returns(bool)
func (_MasterAccessManagement *MasterAccessManagementCallerSession) HasReviewableRequestsExecutePermission(account_ common.Address) (bool, error) {
	return _MasterAccessManagement.Contract.HasReviewableRequestsExecutePermission(&_MasterAccessManagement.CallOpts, account_)
}

// MasterAccessManagementInit is a paid mutator transaction binding the contract method 0x305e4c6c.
//
// Solidity: function __MasterAccessManagement_init(address master_) returns()
func (_MasterAccessManagement *MasterAccessManagementTransactor) MasterAccessManagementInit(opts *bind.TransactOpts, master_ common.Address) (*types.Transaction, error) {
	return _MasterAccessManagement.contract.Transact(opts, "__MasterAccessManagement_init", master_)
}

// MasterAccessManagementInit is a paid mutator transaction binding the contract method 0x305e4c6c.
//
// Solidity: function __MasterAccessManagement_init(address master_) returns()
func (_MasterAccessManagement *MasterAccessManagementSession) MasterAccessManagementInit(master_ common.Address) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.MasterAccessManagementInit(&_MasterAccessManagement.TransactOpts, master_)
}

// MasterAccessManagementInit is a paid mutator transaction binding the contract method 0x305e4c6c.
//
// Solidity: function __MasterAccessManagement_init(address master_) returns()
func (_MasterAccessManagement *MasterAccessManagementTransactorSession) MasterAccessManagementInit(master_ common.Address) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.MasterAccessManagementInit(&_MasterAccessManagement.TransactOpts, master_)
}

// AddPermissionsToRole is a paid mutator transaction binding the contract method 0x37ff630d.
//
// Solidity: function addPermissionsToRole(string role, (string,string[])[] permissionsToAdd, bool allowed) returns()
func (_MasterAccessManagement *MasterAccessManagementTransactor) AddPermissionsToRole(opts *bind.TransactOpts, role string, permissionsToAdd []IRBACResourceWithPermissions, allowed bool) (*types.Transaction, error) {
	return _MasterAccessManagement.contract.Transact(opts, "addPermissionsToRole", role, permissionsToAdd, allowed)
}

// AddPermissionsToRole is a paid mutator transaction binding the contract method 0x37ff630d.
//
// Solidity: function addPermissionsToRole(string role, (string,string[])[] permissionsToAdd, bool allowed) returns()
func (_MasterAccessManagement *MasterAccessManagementSession) AddPermissionsToRole(role string, permissionsToAdd []IRBACResourceWithPermissions, allowed bool) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.AddPermissionsToRole(&_MasterAccessManagement.TransactOpts, role, permissionsToAdd, allowed)
}

// AddPermissionsToRole is a paid mutator transaction binding the contract method 0x37ff630d.
//
// Solidity: function addPermissionsToRole(string role, (string,string[])[] permissionsToAdd, bool allowed) returns()
func (_MasterAccessManagement *MasterAccessManagementTransactorSession) AddPermissionsToRole(role string, permissionsToAdd []IRBACResourceWithPermissions, allowed bool) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.AddPermissionsToRole(&_MasterAccessManagement.TransactOpts, role, permissionsToAdd, allowed)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xee2f6ce5.
//
// Solidity: function grantRoles(address to, string[] rolesToGrant) returns()
func (_MasterAccessManagement *MasterAccessManagementTransactor) GrantRoles(opts *bind.TransactOpts, to common.Address, rolesToGrant []string) (*types.Transaction, error) {
	return _MasterAccessManagement.contract.Transact(opts, "grantRoles", to, rolesToGrant)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xee2f6ce5.
//
// Solidity: function grantRoles(address to, string[] rolesToGrant) returns()
func (_MasterAccessManagement *MasterAccessManagementSession) GrantRoles(to common.Address, rolesToGrant []string) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.GrantRoles(&_MasterAccessManagement.TransactOpts, to, rolesToGrant)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xee2f6ce5.
//
// Solidity: function grantRoles(address to, string[] rolesToGrant) returns()
func (_MasterAccessManagement *MasterAccessManagementTransactorSession) GrantRoles(to common.Address, rolesToGrant []string) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.GrantRoles(&_MasterAccessManagement.TransactOpts, to, rolesToGrant)
}

// RemovePermissionsFromRole is a paid mutator transaction binding the contract method 0x75e025e7.
//
// Solidity: function removePermissionsFromRole(string role, (string,string[])[] permissionsToRemove, bool allowed) returns()
func (_MasterAccessManagement *MasterAccessManagementTransactor) RemovePermissionsFromRole(opts *bind.TransactOpts, role string, permissionsToRemove []IRBACResourceWithPermissions, allowed bool) (*types.Transaction, error) {
	return _MasterAccessManagement.contract.Transact(opts, "removePermissionsFromRole", role, permissionsToRemove, allowed)
}

// RemovePermissionsFromRole is a paid mutator transaction binding the contract method 0x75e025e7.
//
// Solidity: function removePermissionsFromRole(string role, (string,string[])[] permissionsToRemove, bool allowed) returns()
func (_MasterAccessManagement *MasterAccessManagementSession) RemovePermissionsFromRole(role string, permissionsToRemove []IRBACResourceWithPermissions, allowed bool) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.RemovePermissionsFromRole(&_MasterAccessManagement.TransactOpts, role, permissionsToRemove, allowed)
}

// RemovePermissionsFromRole is a paid mutator transaction binding the contract method 0x75e025e7.
//
// Solidity: function removePermissionsFromRole(string role, (string,string[])[] permissionsToRemove, bool allowed) returns()
func (_MasterAccessManagement *MasterAccessManagementTransactorSession) RemovePermissionsFromRole(role string, permissionsToRemove []IRBACResourceWithPermissions, allowed bool) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.RemovePermissionsFromRole(&_MasterAccessManagement.TransactOpts, role, permissionsToRemove, allowed)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4f0d84e3.
//
// Solidity: function revokeRoles(address from, string[] rolesToRevoke) returns()
func (_MasterAccessManagement *MasterAccessManagementTransactor) RevokeRoles(opts *bind.TransactOpts, from common.Address, rolesToRevoke []string) (*types.Transaction, error) {
	return _MasterAccessManagement.contract.Transact(opts, "revokeRoles", from, rolesToRevoke)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4f0d84e3.
//
// Solidity: function revokeRoles(address from, string[] rolesToRevoke) returns()
func (_MasterAccessManagement *MasterAccessManagementSession) RevokeRoles(from common.Address, rolesToRevoke []string) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.RevokeRoles(&_MasterAccessManagement.TransactOpts, from, rolesToRevoke)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4f0d84e3.
//
// Solidity: function revokeRoles(address from, string[] rolesToRevoke) returns()
func (_MasterAccessManagement *MasterAccessManagementTransactorSession) RevokeRoles(from common.Address, rolesToRevoke []string) (*types.Transaction, error) {
	return _MasterAccessManagement.Contract.RevokeRoles(&_MasterAccessManagement.TransactOpts, from, rolesToRevoke)
}

// MasterAccessManagementAddedPermissionsIterator is returned from FilterAddedPermissions and is used to iterate over the raw logs and unpacked data for AddedPermissions events raised by the MasterAccessManagement contract.
type MasterAccessManagementAddedPermissionsIterator struct {
	Event *MasterAccessManagementAddedPermissions // Event containing the contract specifics and raw log

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
func (it *MasterAccessManagementAddedPermissionsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterAccessManagementAddedPermissions)
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
		it.Event = new(MasterAccessManagementAddedPermissions)
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
func (it *MasterAccessManagementAddedPermissionsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterAccessManagementAddedPermissionsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterAccessManagementAddedPermissions represents a AddedPermissions event raised by the MasterAccessManagement contract.
type MasterAccessManagementAddedPermissions struct {
	Role             string
	Resource         string
	PermissionsToAdd []string
	Allowed          bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAddedPermissions is a free log retrieval operation binding the contract event 0x06ea5bb05be257dd30f31cc76f4077a3df230e09f7e33e9b52fb09c40a8f695e.
//
// Solidity: event AddedPermissions(string role, string resource, string[] permissionsToAdd, bool allowed)
func (_MasterAccessManagement *MasterAccessManagementFilterer) FilterAddedPermissions(opts *bind.FilterOpts) (*MasterAccessManagementAddedPermissionsIterator, error) {

	logs, sub, err := _MasterAccessManagement.contract.FilterLogs(opts, "AddedPermissions")
	if err != nil {
		return nil, err
	}
	return &MasterAccessManagementAddedPermissionsIterator{contract: _MasterAccessManagement.contract, event: "AddedPermissions", logs: logs, sub: sub}, nil
}

// WatchAddedPermissions is a free log subscription operation binding the contract event 0x06ea5bb05be257dd30f31cc76f4077a3df230e09f7e33e9b52fb09c40a8f695e.
//
// Solidity: event AddedPermissions(string role, string resource, string[] permissionsToAdd, bool allowed)
func (_MasterAccessManagement *MasterAccessManagementFilterer) WatchAddedPermissions(opts *bind.WatchOpts, sink chan<- *MasterAccessManagementAddedPermissions) (event.Subscription, error) {

	logs, sub, err := _MasterAccessManagement.contract.WatchLogs(opts, "AddedPermissions")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterAccessManagementAddedPermissions)
				if err := _MasterAccessManagement.contract.UnpackLog(event, "AddedPermissions", log); err != nil {
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

// ParseAddedPermissions is a log parse operation binding the contract event 0x06ea5bb05be257dd30f31cc76f4077a3df230e09f7e33e9b52fb09c40a8f695e.
//
// Solidity: event AddedPermissions(string role, string resource, string[] permissionsToAdd, bool allowed)
func (_MasterAccessManagement *MasterAccessManagementFilterer) ParseAddedPermissions(log types.Log) (*MasterAccessManagementAddedPermissions, error) {
	event := new(MasterAccessManagementAddedPermissions)
	if err := _MasterAccessManagement.contract.UnpackLog(event, "AddedPermissions", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterAccessManagementGrantedRolesIterator is returned from FilterGrantedRoles and is used to iterate over the raw logs and unpacked data for GrantedRoles events raised by the MasterAccessManagement contract.
type MasterAccessManagementGrantedRolesIterator struct {
	Event *MasterAccessManagementGrantedRoles // Event containing the contract specifics and raw log

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
func (it *MasterAccessManagementGrantedRolesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterAccessManagementGrantedRoles)
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
		it.Event = new(MasterAccessManagementGrantedRoles)
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
func (it *MasterAccessManagementGrantedRolesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterAccessManagementGrantedRolesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterAccessManagementGrantedRoles represents a GrantedRoles event raised by the MasterAccessManagement contract.
type MasterAccessManagementGrantedRoles struct {
	To           common.Address
	RolesToGrant []string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGrantedRoles is a free log retrieval operation binding the contract event 0x44240f5b60cedf44a65c3717503d91c46a899ef33c5348880e4c29131ac87311.
//
// Solidity: event GrantedRoles(address to, string[] rolesToGrant)
func (_MasterAccessManagement *MasterAccessManagementFilterer) FilterGrantedRoles(opts *bind.FilterOpts) (*MasterAccessManagementGrantedRolesIterator, error) {

	logs, sub, err := _MasterAccessManagement.contract.FilterLogs(opts, "GrantedRoles")
	if err != nil {
		return nil, err
	}
	return &MasterAccessManagementGrantedRolesIterator{contract: _MasterAccessManagement.contract, event: "GrantedRoles", logs: logs, sub: sub}, nil
}

// WatchGrantedRoles is a free log subscription operation binding the contract event 0x44240f5b60cedf44a65c3717503d91c46a899ef33c5348880e4c29131ac87311.
//
// Solidity: event GrantedRoles(address to, string[] rolesToGrant)
func (_MasterAccessManagement *MasterAccessManagementFilterer) WatchGrantedRoles(opts *bind.WatchOpts, sink chan<- *MasterAccessManagementGrantedRoles) (event.Subscription, error) {

	logs, sub, err := _MasterAccessManagement.contract.WatchLogs(opts, "GrantedRoles")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterAccessManagementGrantedRoles)
				if err := _MasterAccessManagement.contract.UnpackLog(event, "GrantedRoles", log); err != nil {
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

// ParseGrantedRoles is a log parse operation binding the contract event 0x44240f5b60cedf44a65c3717503d91c46a899ef33c5348880e4c29131ac87311.
//
// Solidity: event GrantedRoles(address to, string[] rolesToGrant)
func (_MasterAccessManagement *MasterAccessManagementFilterer) ParseGrantedRoles(log types.Log) (*MasterAccessManagementGrantedRoles, error) {
	event := new(MasterAccessManagementGrantedRoles)
	if err := _MasterAccessManagement.contract.UnpackLog(event, "GrantedRoles", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterAccessManagementRemovedPermissionsIterator is returned from FilterRemovedPermissions and is used to iterate over the raw logs and unpacked data for RemovedPermissions events raised by the MasterAccessManagement contract.
type MasterAccessManagementRemovedPermissionsIterator struct {
	Event *MasterAccessManagementRemovedPermissions // Event containing the contract specifics and raw log

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
func (it *MasterAccessManagementRemovedPermissionsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterAccessManagementRemovedPermissions)
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
		it.Event = new(MasterAccessManagementRemovedPermissions)
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
func (it *MasterAccessManagementRemovedPermissionsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterAccessManagementRemovedPermissionsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterAccessManagementRemovedPermissions represents a RemovedPermissions event raised by the MasterAccessManagement contract.
type MasterAccessManagementRemovedPermissions struct {
	Role                string
	Resource            string
	PermissionsToRemove []string
	Allowed             bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRemovedPermissions is a free log retrieval operation binding the contract event 0x27f963dfdedb973db8acb94be0be26fbd4c55498bc54cfd0733b9da9c0ab8296.
//
// Solidity: event RemovedPermissions(string role, string resource, string[] permissionsToRemove, bool allowed)
func (_MasterAccessManagement *MasterAccessManagementFilterer) FilterRemovedPermissions(opts *bind.FilterOpts) (*MasterAccessManagementRemovedPermissionsIterator, error) {

	logs, sub, err := _MasterAccessManagement.contract.FilterLogs(opts, "RemovedPermissions")
	if err != nil {
		return nil, err
	}
	return &MasterAccessManagementRemovedPermissionsIterator{contract: _MasterAccessManagement.contract, event: "RemovedPermissions", logs: logs, sub: sub}, nil
}

// WatchRemovedPermissions is a free log subscription operation binding the contract event 0x27f963dfdedb973db8acb94be0be26fbd4c55498bc54cfd0733b9da9c0ab8296.
//
// Solidity: event RemovedPermissions(string role, string resource, string[] permissionsToRemove, bool allowed)
func (_MasterAccessManagement *MasterAccessManagementFilterer) WatchRemovedPermissions(opts *bind.WatchOpts, sink chan<- *MasterAccessManagementRemovedPermissions) (event.Subscription, error) {

	logs, sub, err := _MasterAccessManagement.contract.WatchLogs(opts, "RemovedPermissions")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterAccessManagementRemovedPermissions)
				if err := _MasterAccessManagement.contract.UnpackLog(event, "RemovedPermissions", log); err != nil {
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

// ParseRemovedPermissions is a log parse operation binding the contract event 0x27f963dfdedb973db8acb94be0be26fbd4c55498bc54cfd0733b9da9c0ab8296.
//
// Solidity: event RemovedPermissions(string role, string resource, string[] permissionsToRemove, bool allowed)
func (_MasterAccessManagement *MasterAccessManagementFilterer) ParseRemovedPermissions(log types.Log) (*MasterAccessManagementRemovedPermissions, error) {
	event := new(MasterAccessManagementRemovedPermissions)
	if err := _MasterAccessManagement.contract.UnpackLog(event, "RemovedPermissions", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterAccessManagementRevokedRolesIterator is returned from FilterRevokedRoles and is used to iterate over the raw logs and unpacked data for RevokedRoles events raised by the MasterAccessManagement contract.
type MasterAccessManagementRevokedRolesIterator struct {
	Event *MasterAccessManagementRevokedRoles // Event containing the contract specifics and raw log

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
func (it *MasterAccessManagementRevokedRolesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterAccessManagementRevokedRoles)
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
		it.Event = new(MasterAccessManagementRevokedRoles)
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
func (it *MasterAccessManagementRevokedRolesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterAccessManagementRevokedRolesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterAccessManagementRevokedRoles represents a RevokedRoles event raised by the MasterAccessManagement contract.
type MasterAccessManagementRevokedRoles struct {
	From          common.Address
	RolesToRevoke []string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevokedRoles is a free log retrieval operation binding the contract event 0x037c273ac6ee1105154063d4b014a5afeec9981076c152f47a0899c6e2854740.
//
// Solidity: event RevokedRoles(address from, string[] rolesToRevoke)
func (_MasterAccessManagement *MasterAccessManagementFilterer) FilterRevokedRoles(opts *bind.FilterOpts) (*MasterAccessManagementRevokedRolesIterator, error) {

	logs, sub, err := _MasterAccessManagement.contract.FilterLogs(opts, "RevokedRoles")
	if err != nil {
		return nil, err
	}
	return &MasterAccessManagementRevokedRolesIterator{contract: _MasterAccessManagement.contract, event: "RevokedRoles", logs: logs, sub: sub}, nil
}

// WatchRevokedRoles is a free log subscription operation binding the contract event 0x037c273ac6ee1105154063d4b014a5afeec9981076c152f47a0899c6e2854740.
//
// Solidity: event RevokedRoles(address from, string[] rolesToRevoke)
func (_MasterAccessManagement *MasterAccessManagementFilterer) WatchRevokedRoles(opts *bind.WatchOpts, sink chan<- *MasterAccessManagementRevokedRoles) (event.Subscription, error) {

	logs, sub, err := _MasterAccessManagement.contract.WatchLogs(opts, "RevokedRoles")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterAccessManagementRevokedRoles)
				if err := _MasterAccessManagement.contract.UnpackLog(event, "RevokedRoles", log); err != nil {
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

// ParseRevokedRoles is a log parse operation binding the contract event 0x037c273ac6ee1105154063d4b014a5afeec9981076c152f47a0899c6e2854740.
//
// Solidity: event RevokedRoles(address from, string[] rolesToRevoke)
func (_MasterAccessManagement *MasterAccessManagementFilterer) ParseRevokedRoles(log types.Log) (*MasterAccessManagementRevokedRoles, error) {
	event := new(MasterAccessManagementRevokedRoles)
	if err := _MasterAccessManagement.contract.UnpackLog(event, "RevokedRoles", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

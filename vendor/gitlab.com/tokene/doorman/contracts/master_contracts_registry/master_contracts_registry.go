// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package master_contracts_registry

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

// MasterContractsRegistryMetaData contains all meta data concerning the MasterContractsRegistry contract.
var MasterContractsRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isProxy\",\"type\":\"bool\"}],\"name\":\"AddedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"RemovedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONSTANTS_REGISTRY_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_ACCESS_MANAGEMENT_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REVIEWABLE_REQUESTS_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterAccess_\",\"type\":\"address\"}],\"name\":\"__MasterContractsRegistry_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"addContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"addProxyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConstantsRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMasterAccessManagement\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyUpgrader\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReviewableRequests\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"hasContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"injectDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"justAddProxyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"removeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"}],\"name\":\"upgradeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"upgradeContractAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// MasterContractsRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterContractsRegistryMetaData.ABI instead.
var MasterContractsRegistryABI = MasterContractsRegistryMetaData.ABI

// MasterContractsRegistry is an auto generated Go binding around an Ethereum contract.
type MasterContractsRegistry struct {
	MasterContractsRegistryCaller     // Read-only binding to the contract
	MasterContractsRegistryTransactor // Write-only binding to the contract
	MasterContractsRegistryFilterer   // Log filterer for contract events
}

// MasterContractsRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MasterContractsRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterContractsRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterContractsRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterContractsRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterContractsRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterContractsRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterContractsRegistrySession struct {
	Contract     *MasterContractsRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MasterContractsRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterContractsRegistryCallerSession struct {
	Contract *MasterContractsRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// MasterContractsRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterContractsRegistryTransactorSession struct {
	Contract     *MasterContractsRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// MasterContractsRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MasterContractsRegistryRaw struct {
	Contract *MasterContractsRegistry // Generic contract binding to access the raw methods on
}

// MasterContractsRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterContractsRegistryCallerRaw struct {
	Contract *MasterContractsRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MasterContractsRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterContractsRegistryTransactorRaw struct {
	Contract *MasterContractsRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterContractsRegistry creates a new instance of MasterContractsRegistry, bound to a specific deployed contract.
func NewMasterContractsRegistry(address common.Address, backend bind.ContractBackend) (*MasterContractsRegistry, error) {
	contract, err := bindMasterContractsRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasterContractsRegistry{MasterContractsRegistryCaller: MasterContractsRegistryCaller{contract: contract}, MasterContractsRegistryTransactor: MasterContractsRegistryTransactor{contract: contract}, MasterContractsRegistryFilterer: MasterContractsRegistryFilterer{contract: contract}}, nil
}

// NewMasterContractsRegistryCaller creates a new read-only instance of MasterContractsRegistry, bound to a specific deployed contract.
func NewMasterContractsRegistryCaller(address common.Address, caller bind.ContractCaller) (*MasterContractsRegistryCaller, error) {
	contract, err := bindMasterContractsRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterContractsRegistryCaller{contract: contract}, nil
}

// NewMasterContractsRegistryTransactor creates a new write-only instance of MasterContractsRegistry, bound to a specific deployed contract.
func NewMasterContractsRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MasterContractsRegistryTransactor, error) {
	contract, err := bindMasterContractsRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterContractsRegistryTransactor{contract: contract}, nil
}

// NewMasterContractsRegistryFilterer creates a new log filterer instance of MasterContractsRegistry, bound to a specific deployed contract.
func NewMasterContractsRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MasterContractsRegistryFilterer, error) {
	contract, err := bindMasterContractsRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterContractsRegistryFilterer{contract: contract}, nil
}

// bindMasterContractsRegistry binds a generic wrapper to an already deployed contract.
func bindMasterContractsRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MasterContractsRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterContractsRegistry *MasterContractsRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterContractsRegistry.Contract.MasterContractsRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterContractsRegistry *MasterContractsRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.MasterContractsRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterContractsRegistry *MasterContractsRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.MasterContractsRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterContractsRegistry *MasterContractsRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterContractsRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterContractsRegistry *MasterContractsRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterContractsRegistry *MasterContractsRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.contract.Transact(opts, method, params...)
}

// CONSTANTSREGISTRYNAME is a free data retrieval call binding the contract method 0x26a31c6b.
//
// Solidity: function CONSTANTS_REGISTRY_NAME() view returns(string)
func (_MasterContractsRegistry *MasterContractsRegistryCaller) CONSTANTSREGISTRYNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterContractsRegistry.contract.Call(opts, &out, "CONSTANTS_REGISTRY_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CONSTANTSREGISTRYNAME is a free data retrieval call binding the contract method 0x26a31c6b.
//
// Solidity: function CONSTANTS_REGISTRY_NAME() view returns(string)
func (_MasterContractsRegistry *MasterContractsRegistrySession) CONSTANTSREGISTRYNAME() (string, error) {
	return _MasterContractsRegistry.Contract.CONSTANTSREGISTRYNAME(&_MasterContractsRegistry.CallOpts)
}

// CONSTANTSREGISTRYNAME is a free data retrieval call binding the contract method 0x26a31c6b.
//
// Solidity: function CONSTANTS_REGISTRY_NAME() view returns(string)
func (_MasterContractsRegistry *MasterContractsRegistryCallerSession) CONSTANTSREGISTRYNAME() (string, error) {
	return _MasterContractsRegistry.Contract.CONSTANTSREGISTRYNAME(&_MasterContractsRegistry.CallOpts)
}

// MASTERACCESSMANAGEMENTNAME is a free data retrieval call binding the contract method 0xa4ce3439.
//
// Solidity: function MASTER_ACCESS_MANAGEMENT_NAME() view returns(string)
func (_MasterContractsRegistry *MasterContractsRegistryCaller) MASTERACCESSMANAGEMENTNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterContractsRegistry.contract.Call(opts, &out, "MASTER_ACCESS_MANAGEMENT_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MASTERACCESSMANAGEMENTNAME is a free data retrieval call binding the contract method 0xa4ce3439.
//
// Solidity: function MASTER_ACCESS_MANAGEMENT_NAME() view returns(string)
func (_MasterContractsRegistry *MasterContractsRegistrySession) MASTERACCESSMANAGEMENTNAME() (string, error) {
	return _MasterContractsRegistry.Contract.MASTERACCESSMANAGEMENTNAME(&_MasterContractsRegistry.CallOpts)
}

// MASTERACCESSMANAGEMENTNAME is a free data retrieval call binding the contract method 0xa4ce3439.
//
// Solidity: function MASTER_ACCESS_MANAGEMENT_NAME() view returns(string)
func (_MasterContractsRegistry *MasterContractsRegistryCallerSession) MASTERACCESSMANAGEMENTNAME() (string, error) {
	return _MasterContractsRegistry.Contract.MASTERACCESSMANAGEMENTNAME(&_MasterContractsRegistry.CallOpts)
}

// REVIEWABLEREQUESTSNAME is a free data retrieval call binding the contract method 0x1b236de3.
//
// Solidity: function REVIEWABLE_REQUESTS_NAME() view returns(string)
func (_MasterContractsRegistry *MasterContractsRegistryCaller) REVIEWABLEREQUESTSNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterContractsRegistry.contract.Call(opts, &out, "REVIEWABLE_REQUESTS_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// REVIEWABLEREQUESTSNAME is a free data retrieval call binding the contract method 0x1b236de3.
//
// Solidity: function REVIEWABLE_REQUESTS_NAME() view returns(string)
func (_MasterContractsRegistry *MasterContractsRegistrySession) REVIEWABLEREQUESTSNAME() (string, error) {
	return _MasterContractsRegistry.Contract.REVIEWABLEREQUESTSNAME(&_MasterContractsRegistry.CallOpts)
}

// REVIEWABLEREQUESTSNAME is a free data retrieval call binding the contract method 0x1b236de3.
//
// Solidity: function REVIEWABLE_REQUESTS_NAME() view returns(string)
func (_MasterContractsRegistry *MasterContractsRegistryCallerSession) REVIEWABLEREQUESTSNAME() (string, error) {
	return _MasterContractsRegistry.Contract.REVIEWABLEREQUESTSNAME(&_MasterContractsRegistry.CallOpts)
}

// GetConstantsRegistry is a free data retrieval call binding the contract method 0xf448f730.
//
// Solidity: function getConstantsRegistry() view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistryCaller) GetConstantsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterContractsRegistry.contract.Call(opts, &out, "getConstantsRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConstantsRegistry is a free data retrieval call binding the contract method 0xf448f730.
//
// Solidity: function getConstantsRegistry() view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistrySession) GetConstantsRegistry() (common.Address, error) {
	return _MasterContractsRegistry.Contract.GetConstantsRegistry(&_MasterContractsRegistry.CallOpts)
}

// GetConstantsRegistry is a free data retrieval call binding the contract method 0xf448f730.
//
// Solidity: function getConstantsRegistry() view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistryCallerSession) GetConstantsRegistry() (common.Address, error) {
	return _MasterContractsRegistry.Contract.GetConstantsRegistry(&_MasterContractsRegistry.CallOpts)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string name) view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistryCaller) GetContract(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _MasterContractsRegistry.contract.Call(opts, &out, "getContract", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string name) view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistrySession) GetContract(name string) (common.Address, error) {
	return _MasterContractsRegistry.Contract.GetContract(&_MasterContractsRegistry.CallOpts, name)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string name) view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistryCallerSession) GetContract(name string) (common.Address, error) {
	return _MasterContractsRegistry.Contract.GetContract(&_MasterContractsRegistry.CallOpts, name)
}

// GetImplementation is a free data retrieval call binding the contract method 0x6b683896.
//
// Solidity: function getImplementation(string name) view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistryCaller) GetImplementation(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _MasterContractsRegistry.contract.Call(opts, &out, "getImplementation", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0x6b683896.
//
// Solidity: function getImplementation(string name) view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistrySession) GetImplementation(name string) (common.Address, error) {
	return _MasterContractsRegistry.Contract.GetImplementation(&_MasterContractsRegistry.CallOpts, name)
}

// GetImplementation is a free data retrieval call binding the contract method 0x6b683896.
//
// Solidity: function getImplementation(string name) view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistryCallerSession) GetImplementation(name string) (common.Address, error) {
	return _MasterContractsRegistry.Contract.GetImplementation(&_MasterContractsRegistry.CallOpts, name)
}

// GetMasterAccessManagement is a free data retrieval call binding the contract method 0x085ec297.
//
// Solidity: function getMasterAccessManagement() view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistryCaller) GetMasterAccessManagement(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterContractsRegistry.contract.Call(opts, &out, "getMasterAccessManagement")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMasterAccessManagement is a free data retrieval call binding the contract method 0x085ec297.
//
// Solidity: function getMasterAccessManagement() view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistrySession) GetMasterAccessManagement() (common.Address, error) {
	return _MasterContractsRegistry.Contract.GetMasterAccessManagement(&_MasterContractsRegistry.CallOpts)
}

// GetMasterAccessManagement is a free data retrieval call binding the contract method 0x085ec297.
//
// Solidity: function getMasterAccessManagement() view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistryCallerSession) GetMasterAccessManagement() (common.Address, error) {
	return _MasterContractsRegistry.Contract.GetMasterAccessManagement(&_MasterContractsRegistry.CallOpts)
}

// GetProxyUpgrader is a free data retrieval call binding the contract method 0xd10611fc.
//
// Solidity: function getProxyUpgrader() view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistryCaller) GetProxyUpgrader(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterContractsRegistry.contract.Call(opts, &out, "getProxyUpgrader")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyUpgrader is a free data retrieval call binding the contract method 0xd10611fc.
//
// Solidity: function getProxyUpgrader() view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistrySession) GetProxyUpgrader() (common.Address, error) {
	return _MasterContractsRegistry.Contract.GetProxyUpgrader(&_MasterContractsRegistry.CallOpts)
}

// GetProxyUpgrader is a free data retrieval call binding the contract method 0xd10611fc.
//
// Solidity: function getProxyUpgrader() view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistryCallerSession) GetProxyUpgrader() (common.Address, error) {
	return _MasterContractsRegistry.Contract.GetProxyUpgrader(&_MasterContractsRegistry.CallOpts)
}

// GetReviewableRequests is a free data retrieval call binding the contract method 0x70371abb.
//
// Solidity: function getReviewableRequests() view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistryCaller) GetReviewableRequests(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterContractsRegistry.contract.Call(opts, &out, "getReviewableRequests")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReviewableRequests is a free data retrieval call binding the contract method 0x70371abb.
//
// Solidity: function getReviewableRequests() view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistrySession) GetReviewableRequests() (common.Address, error) {
	return _MasterContractsRegistry.Contract.GetReviewableRequests(&_MasterContractsRegistry.CallOpts)
}

// GetReviewableRequests is a free data retrieval call binding the contract method 0x70371abb.
//
// Solidity: function getReviewableRequests() view returns(address)
func (_MasterContractsRegistry *MasterContractsRegistryCallerSession) GetReviewableRequests() (common.Address, error) {
	return _MasterContractsRegistry.Contract.GetReviewableRequests(&_MasterContractsRegistry.CallOpts)
}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(string name) view returns(bool)
func (_MasterContractsRegistry *MasterContractsRegistryCaller) HasContract(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _MasterContractsRegistry.contract.Call(opts, &out, "hasContract", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(string name) view returns(bool)
func (_MasterContractsRegistry *MasterContractsRegistrySession) HasContract(name string) (bool, error) {
	return _MasterContractsRegistry.Contract.HasContract(&_MasterContractsRegistry.CallOpts, name)
}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(string name) view returns(bool)
func (_MasterContractsRegistry *MasterContractsRegistryCallerSession) HasContract(name string) (bool, error) {
	return _MasterContractsRegistry.Contract.HasContract(&_MasterContractsRegistry.CallOpts, name)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MasterContractsRegistry *MasterContractsRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MasterContractsRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MasterContractsRegistry *MasterContractsRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _MasterContractsRegistry.Contract.ProxiableUUID(&_MasterContractsRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MasterContractsRegistry *MasterContractsRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _MasterContractsRegistry.Contract.ProxiableUUID(&_MasterContractsRegistry.CallOpts)
}

// MasterContractsRegistryInit is a paid mutator transaction binding the contract method 0x3e6decde.
//
// Solidity: function __MasterContractsRegistry_init(address masterAccess_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactor) MasterContractsRegistryInit(opts *bind.TransactOpts, masterAccess_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.contract.Transact(opts, "__MasterContractsRegistry_init", masterAccess_)
}

// MasterContractsRegistryInit is a paid mutator transaction binding the contract method 0x3e6decde.
//
// Solidity: function __MasterContractsRegistry_init(address masterAccess_) returns()
func (_MasterContractsRegistry *MasterContractsRegistrySession) MasterContractsRegistryInit(masterAccess_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.MasterContractsRegistryInit(&_MasterContractsRegistry.TransactOpts, masterAccess_)
}

// MasterContractsRegistryInit is a paid mutator transaction binding the contract method 0x3e6decde.
//
// Solidity: function __MasterContractsRegistry_init(address masterAccess_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactorSession) MasterContractsRegistryInit(masterAccess_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.MasterContractsRegistryInit(&_MasterContractsRegistry.TransactOpts, masterAccess_)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string name_, address contractAddress_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactor) AddContract(opts *bind.TransactOpts, name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.contract.Transact(opts, "addContract", name_, contractAddress_)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string name_, address contractAddress_) returns()
func (_MasterContractsRegistry *MasterContractsRegistrySession) AddContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.AddContract(&_MasterContractsRegistry.TransactOpts, name_, contractAddress_)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string name_, address contractAddress_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactorSession) AddContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.AddContract(&_MasterContractsRegistry.TransactOpts, name_, contractAddress_)
}

// AddProxyContract is a paid mutator transaction binding the contract method 0xe0e084f8.
//
// Solidity: function addProxyContract(string name_, address contractAddress_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactor) AddProxyContract(opts *bind.TransactOpts, name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.contract.Transact(opts, "addProxyContract", name_, contractAddress_)
}

// AddProxyContract is a paid mutator transaction binding the contract method 0xe0e084f8.
//
// Solidity: function addProxyContract(string name_, address contractAddress_) returns()
func (_MasterContractsRegistry *MasterContractsRegistrySession) AddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.AddProxyContract(&_MasterContractsRegistry.TransactOpts, name_, contractAddress_)
}

// AddProxyContract is a paid mutator transaction binding the contract method 0xe0e084f8.
//
// Solidity: function addProxyContract(string name_, address contractAddress_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactorSession) AddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.AddProxyContract(&_MasterContractsRegistry.TransactOpts, name_, contractAddress_)
}

// InjectDependencies is a paid mutator transaction binding the contract method 0x1adad8cf.
//
// Solidity: function injectDependencies(string name_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactor) InjectDependencies(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _MasterContractsRegistry.contract.Transact(opts, "injectDependencies", name_)
}

// InjectDependencies is a paid mutator transaction binding the contract method 0x1adad8cf.
//
// Solidity: function injectDependencies(string name_) returns()
func (_MasterContractsRegistry *MasterContractsRegistrySession) InjectDependencies(name_ string) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.InjectDependencies(&_MasterContractsRegistry.TransactOpts, name_)
}

// InjectDependencies is a paid mutator transaction binding the contract method 0x1adad8cf.
//
// Solidity: function injectDependencies(string name_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactorSession) InjectDependencies(name_ string) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.InjectDependencies(&_MasterContractsRegistry.TransactOpts, name_)
}

// JustAddProxyContract is a paid mutator transaction binding the contract method 0x51dad82c.
//
// Solidity: function justAddProxyContract(string name_, address contractAddress_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactor) JustAddProxyContract(opts *bind.TransactOpts, name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.contract.Transact(opts, "justAddProxyContract", name_, contractAddress_)
}

// JustAddProxyContract is a paid mutator transaction binding the contract method 0x51dad82c.
//
// Solidity: function justAddProxyContract(string name_, address contractAddress_) returns()
func (_MasterContractsRegistry *MasterContractsRegistrySession) JustAddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.JustAddProxyContract(&_MasterContractsRegistry.TransactOpts, name_, contractAddress_)
}

// JustAddProxyContract is a paid mutator transaction binding the contract method 0x51dad82c.
//
// Solidity: function justAddProxyContract(string name_, address contractAddress_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactorSession) JustAddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.JustAddProxyContract(&_MasterContractsRegistry.TransactOpts, name_, contractAddress_)
}

// RemoveContract is a paid mutator transaction binding the contract method 0x97623b58.
//
// Solidity: function removeContract(string name_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactor) RemoveContract(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _MasterContractsRegistry.contract.Transact(opts, "removeContract", name_)
}

// RemoveContract is a paid mutator transaction binding the contract method 0x97623b58.
//
// Solidity: function removeContract(string name_) returns()
func (_MasterContractsRegistry *MasterContractsRegistrySession) RemoveContract(name_ string) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.RemoveContract(&_MasterContractsRegistry.TransactOpts, name_)
}

// RemoveContract is a paid mutator transaction binding the contract method 0x97623b58.
//
// Solidity: function removeContract(string name_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactorSession) RemoveContract(name_ string) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.RemoveContract(&_MasterContractsRegistry.TransactOpts, name_)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x1271bd53.
//
// Solidity: function upgradeContract(string name_, address newImplementation_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactor) UpgradeContract(opts *bind.TransactOpts, name_ string, newImplementation_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.contract.Transact(opts, "upgradeContract", name_, newImplementation_)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x1271bd53.
//
// Solidity: function upgradeContract(string name_, address newImplementation_) returns()
func (_MasterContractsRegistry *MasterContractsRegistrySession) UpgradeContract(name_ string, newImplementation_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.UpgradeContract(&_MasterContractsRegistry.TransactOpts, name_, newImplementation_)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x1271bd53.
//
// Solidity: function upgradeContract(string name_, address newImplementation_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactorSession) UpgradeContract(name_ string, newImplementation_ common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.UpgradeContract(&_MasterContractsRegistry.TransactOpts, name_, newImplementation_)
}

// UpgradeContractAndCall is a paid mutator transaction binding the contract method 0x6bbe8694.
//
// Solidity: function upgradeContractAndCall(string name_, address newImplementation_, bytes data_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactor) UpgradeContractAndCall(opts *bind.TransactOpts, name_ string, newImplementation_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _MasterContractsRegistry.contract.Transact(opts, "upgradeContractAndCall", name_, newImplementation_, data_)
}

// UpgradeContractAndCall is a paid mutator transaction binding the contract method 0x6bbe8694.
//
// Solidity: function upgradeContractAndCall(string name_, address newImplementation_, bytes data_) returns()
func (_MasterContractsRegistry *MasterContractsRegistrySession) UpgradeContractAndCall(name_ string, newImplementation_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.UpgradeContractAndCall(&_MasterContractsRegistry.TransactOpts, name_, newImplementation_, data_)
}

// UpgradeContractAndCall is a paid mutator transaction binding the contract method 0x6bbe8694.
//
// Solidity: function upgradeContractAndCall(string name_, address newImplementation_, bytes data_) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactorSession) UpgradeContractAndCall(name_ string, newImplementation_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.UpgradeContractAndCall(&_MasterContractsRegistry.TransactOpts, name_, newImplementation_, data_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MasterContractsRegistry *MasterContractsRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.UpgradeTo(&_MasterContractsRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.UpgradeTo(&_MasterContractsRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MasterContractsRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MasterContractsRegistry *MasterContractsRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.UpgradeToAndCall(&_MasterContractsRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MasterContractsRegistry *MasterContractsRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MasterContractsRegistry.Contract.UpgradeToAndCall(&_MasterContractsRegistry.TransactOpts, newImplementation, data)
}

// MasterContractsRegistryAddedContractIterator is returned from FilterAddedContract and is used to iterate over the raw logs and unpacked data for AddedContract events raised by the MasterContractsRegistry contract.
type MasterContractsRegistryAddedContractIterator struct {
	Event *MasterContractsRegistryAddedContract // Event containing the contract specifics and raw log

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
func (it *MasterContractsRegistryAddedContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterContractsRegistryAddedContract)
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
		it.Event = new(MasterContractsRegistryAddedContract)
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
func (it *MasterContractsRegistryAddedContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterContractsRegistryAddedContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterContractsRegistryAddedContract represents a AddedContract event raised by the MasterContractsRegistry contract.
type MasterContractsRegistryAddedContract struct {
	Name            string
	ContractAddress common.Address
	IsProxy         bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddedContract is a free log retrieval operation binding the contract event 0xd8607592920f7ac4d1e9d198396c7f7cdaee8e675bc3ffb12924a71979867625.
//
// Solidity: event AddedContract(string name, address contractAddress, bool isProxy)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) FilterAddedContract(opts *bind.FilterOpts) (*MasterContractsRegistryAddedContractIterator, error) {

	logs, sub, err := _MasterContractsRegistry.contract.FilterLogs(opts, "AddedContract")
	if err != nil {
		return nil, err
	}
	return &MasterContractsRegistryAddedContractIterator{contract: _MasterContractsRegistry.contract, event: "AddedContract", logs: logs, sub: sub}, nil
}

// WatchAddedContract is a free log subscription operation binding the contract event 0xd8607592920f7ac4d1e9d198396c7f7cdaee8e675bc3ffb12924a71979867625.
//
// Solidity: event AddedContract(string name, address contractAddress, bool isProxy)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) WatchAddedContract(opts *bind.WatchOpts, sink chan<- *MasterContractsRegistryAddedContract) (event.Subscription, error) {

	logs, sub, err := _MasterContractsRegistry.contract.WatchLogs(opts, "AddedContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterContractsRegistryAddedContract)
				if err := _MasterContractsRegistry.contract.UnpackLog(event, "AddedContract", log); err != nil {
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

// ParseAddedContract is a log parse operation binding the contract event 0xd8607592920f7ac4d1e9d198396c7f7cdaee8e675bc3ffb12924a71979867625.
//
// Solidity: event AddedContract(string name, address contractAddress, bool isProxy)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) ParseAddedContract(log types.Log) (*MasterContractsRegistryAddedContract, error) {
	event := new(MasterContractsRegistryAddedContract)
	if err := _MasterContractsRegistry.contract.UnpackLog(event, "AddedContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterContractsRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the MasterContractsRegistry contract.
type MasterContractsRegistryAdminChangedIterator struct {
	Event *MasterContractsRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *MasterContractsRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterContractsRegistryAdminChanged)
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
		it.Event = new(MasterContractsRegistryAdminChanged)
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
func (it *MasterContractsRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterContractsRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterContractsRegistryAdminChanged represents a AdminChanged event raised by the MasterContractsRegistry contract.
type MasterContractsRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*MasterContractsRegistryAdminChangedIterator, error) {

	logs, sub, err := _MasterContractsRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &MasterContractsRegistryAdminChangedIterator{contract: _MasterContractsRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *MasterContractsRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _MasterContractsRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterContractsRegistryAdminChanged)
				if err := _MasterContractsRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) ParseAdminChanged(log types.Log) (*MasterContractsRegistryAdminChanged, error) {
	event := new(MasterContractsRegistryAdminChanged)
	if err := _MasterContractsRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterContractsRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the MasterContractsRegistry contract.
type MasterContractsRegistryBeaconUpgradedIterator struct {
	Event *MasterContractsRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *MasterContractsRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterContractsRegistryBeaconUpgraded)
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
		it.Event = new(MasterContractsRegistryBeaconUpgraded)
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
func (it *MasterContractsRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterContractsRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterContractsRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the MasterContractsRegistry contract.
type MasterContractsRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*MasterContractsRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _MasterContractsRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &MasterContractsRegistryBeaconUpgradedIterator{contract: _MasterContractsRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *MasterContractsRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _MasterContractsRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterContractsRegistryBeaconUpgraded)
				if err := _MasterContractsRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*MasterContractsRegistryBeaconUpgraded, error) {
	event := new(MasterContractsRegistryBeaconUpgraded)
	if err := _MasterContractsRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterContractsRegistryRemovedContractIterator is returned from FilterRemovedContract and is used to iterate over the raw logs and unpacked data for RemovedContract events raised by the MasterContractsRegistry contract.
type MasterContractsRegistryRemovedContractIterator struct {
	Event *MasterContractsRegistryRemovedContract // Event containing the contract specifics and raw log

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
func (it *MasterContractsRegistryRemovedContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterContractsRegistryRemovedContract)
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
		it.Event = new(MasterContractsRegistryRemovedContract)
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
func (it *MasterContractsRegistryRemovedContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterContractsRegistryRemovedContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterContractsRegistryRemovedContract represents a RemovedContract event raised by the MasterContractsRegistry contract.
type MasterContractsRegistryRemovedContract struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedContract is a free log retrieval operation binding the contract event 0x2216f135d7c337bc5b25f8bf96227730d40241c50adf6fde4d180deca4b25664.
//
// Solidity: event RemovedContract(string name)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) FilterRemovedContract(opts *bind.FilterOpts) (*MasterContractsRegistryRemovedContractIterator, error) {

	logs, sub, err := _MasterContractsRegistry.contract.FilterLogs(opts, "RemovedContract")
	if err != nil {
		return nil, err
	}
	return &MasterContractsRegistryRemovedContractIterator{contract: _MasterContractsRegistry.contract, event: "RemovedContract", logs: logs, sub: sub}, nil
}

// WatchRemovedContract is a free log subscription operation binding the contract event 0x2216f135d7c337bc5b25f8bf96227730d40241c50adf6fde4d180deca4b25664.
//
// Solidity: event RemovedContract(string name)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) WatchRemovedContract(opts *bind.WatchOpts, sink chan<- *MasterContractsRegistryRemovedContract) (event.Subscription, error) {

	logs, sub, err := _MasterContractsRegistry.contract.WatchLogs(opts, "RemovedContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterContractsRegistryRemovedContract)
				if err := _MasterContractsRegistry.contract.UnpackLog(event, "RemovedContract", log); err != nil {
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

// ParseRemovedContract is a log parse operation binding the contract event 0x2216f135d7c337bc5b25f8bf96227730d40241c50adf6fde4d180deca4b25664.
//
// Solidity: event RemovedContract(string name)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) ParseRemovedContract(log types.Log) (*MasterContractsRegistryRemovedContract, error) {
	event := new(MasterContractsRegistryRemovedContract)
	if err := _MasterContractsRegistry.contract.UnpackLog(event, "RemovedContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterContractsRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MasterContractsRegistry contract.
type MasterContractsRegistryUpgradedIterator struct {
	Event *MasterContractsRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *MasterContractsRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterContractsRegistryUpgraded)
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
		it.Event = new(MasterContractsRegistryUpgraded)
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
func (it *MasterContractsRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterContractsRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterContractsRegistryUpgraded represents a Upgraded event raised by the MasterContractsRegistry contract.
type MasterContractsRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MasterContractsRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MasterContractsRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MasterContractsRegistryUpgradedIterator{contract: _MasterContractsRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MasterContractsRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MasterContractsRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterContractsRegistryUpgraded)
				if err := _MasterContractsRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MasterContractsRegistry *MasterContractsRegistryFilterer) ParseUpgraded(log types.Log) (*MasterContractsRegistryUpgraded, error) {
	event := new(MasterContractsRegistryUpgraded)
	if err := _MasterContractsRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

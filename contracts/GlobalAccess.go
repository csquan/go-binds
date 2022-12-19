// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GlobalAccess

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

// GlobalAccessV1StorageUserInfo is an auto generated low-level Go binding around an user-defined struct.
type GlobalAccessV1StorageUserInfo struct {
	Idx   *big.Int
	Typ   *big.Int
	Level *big.Int
}

// GlobalAccessV1StorageUserStatus is an auto generated low-level Go binding around an user-defined struct.
type GlobalAccessV1StorageUserStatus struct {
	Status    uint8
	PauseFrom *big.Int
	PausedTo  *big.Int
}

// GlobalAccessMetaData contains all meta data concerning the GlobalAccess contract.
var GlobalAccessMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SetQuota\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SetUserInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumGlobalAccessV1Storage.Status\",\"name\":\"\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SetUserStatus\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"administrators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractGlobalAccessProxy\",\"name\":\"_globalAccessProxy\",\"type\":\"address\"}],\"name\":\"become\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"typ\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"internalType\":\"structGlobalAccessV1Storage.UserInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getUserStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumGlobalAccessV1Storage.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pauseFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pausedTo\",\"type\":\"uint256\"}],\"internalType\":\"structGlobalAccessV1Storage.UserStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalAccessImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGlobalAccessImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"quotas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quota\",\"type\":\"uint256\"}],\"name\":\"setQuota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"}],\"name\":\"setUserInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"enumGlobalAccessV1Storage.Status\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_pausedFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pausedTo\",\"type\":\"uint256\"}],\"name\":\"setUserStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userStatus\",\"outputs\":[{\"internalType\":\"enumGlobalAccessV1Storage.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pauseFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pausedTo\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"typ\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usersList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GlobalAccessABI is the input ABI used to generate the binding from.
// Deprecated: Use GlobalAccessMetaData.ABI instead.
var GlobalAccessABI = GlobalAccessMetaData.ABI

// GlobalAccess is an auto generated Go binding around an Ethereum contract.
type GlobalAccess struct {
	GlobalAccessCaller     // Read-only binding to the contract
	GlobalAccessTransactor // Write-only binding to the contract
	GlobalAccessFilterer   // Log filterer for contract events
}

// GlobalAccessCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalAccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalAccessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalAccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalAccessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalAccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalAccessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalAccessSession struct {
	Contract     *GlobalAccess     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlobalAccessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalAccessCallerSession struct {
	Contract *GlobalAccessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GlobalAccessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalAccessTransactorSession struct {
	Contract     *GlobalAccessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GlobalAccessRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalAccessRaw struct {
	Contract *GlobalAccess // Generic contract binding to access the raw methods on
}

// GlobalAccessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalAccessCallerRaw struct {
	Contract *GlobalAccessCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalAccessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalAccessTransactorRaw struct {
	Contract *GlobalAccessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalAccess creates a new instance of GlobalAccess, bound to a specific deployed contract.
func NewGlobalAccess(address common.Address, backend bind.ContractBackend) (*GlobalAccess, error) {
	contract, err := bindGlobalAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalAccess{GlobalAccessCaller: GlobalAccessCaller{contract: contract}, GlobalAccessTransactor: GlobalAccessTransactor{contract: contract}, GlobalAccessFilterer: GlobalAccessFilterer{contract: contract}}, nil
}

// NewGlobalAccessCaller creates a new read-only instance of GlobalAccess, bound to a specific deployed contract.
func NewGlobalAccessCaller(address common.Address, caller bind.ContractCaller) (*GlobalAccessCaller, error) {
	contract, err := bindGlobalAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalAccessCaller{contract: contract}, nil
}

// NewGlobalAccessTransactor creates a new write-only instance of GlobalAccess, bound to a specific deployed contract.
func NewGlobalAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalAccessTransactor, error) {
	contract, err := bindGlobalAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalAccessTransactor{contract: contract}, nil
}

// NewGlobalAccessFilterer creates a new log filterer instance of GlobalAccess, bound to a specific deployed contract.
func NewGlobalAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalAccessFilterer, error) {
	contract, err := bindGlobalAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalAccessFilterer{contract: contract}, nil
}

// bindGlobalAccess binds a generic wrapper to an already deployed contract.
func bindGlobalAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalAccessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalAccess *GlobalAccessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlobalAccess.Contract.GlobalAccessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalAccess *GlobalAccessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalAccess.Contract.GlobalAccessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalAccess *GlobalAccessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalAccess.Contract.GlobalAccessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalAccess *GlobalAccessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlobalAccess.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalAccess *GlobalAccessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalAccess.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalAccess *GlobalAccessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalAccess.Contract.contract.Transact(opts, method, params...)
}

// Administrators is a free data retrieval call binding the contract method 0x76be1585.
//
// Solidity: function administrators(address ) view returns(uint256)
func (_GlobalAccess *GlobalAccessCaller) Administrators(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GlobalAccess.contract.Call(opts, &out, "administrators", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Administrators is a free data retrieval call binding the contract method 0x76be1585.
//
// Solidity: function administrators(address ) view returns(uint256)
func (_GlobalAccess *GlobalAccessSession) Administrators(arg0 common.Address) (*big.Int, error) {
	return _GlobalAccess.Contract.Administrators(&_GlobalAccess.CallOpts, arg0)
}

// Administrators is a free data retrieval call binding the contract method 0x76be1585.
//
// Solidity: function administrators(address ) view returns(uint256)
func (_GlobalAccess *GlobalAccessCallerSession) Administrators(arg0 common.Address) (*big.Int, error) {
	return _GlobalAccess.Contract.Administrators(&_GlobalAccess.CallOpts, arg0)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address _account) view returns((uint256,uint256,uint256))
func (_GlobalAccess *GlobalAccessCaller) GetUserInfo(opts *bind.CallOpts, _account common.Address) (GlobalAccessV1StorageUserInfo, error) {
	var out []interface{}
	err := _GlobalAccess.contract.Call(opts, &out, "getUserInfo", _account)

	if err != nil {
		return *new(GlobalAccessV1StorageUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(GlobalAccessV1StorageUserInfo)).(*GlobalAccessV1StorageUserInfo)

	return out0, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address _account) view returns((uint256,uint256,uint256))
func (_GlobalAccess *GlobalAccessSession) GetUserInfo(_account common.Address) (GlobalAccessV1StorageUserInfo, error) {
	return _GlobalAccess.Contract.GetUserInfo(&_GlobalAccess.CallOpts, _account)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address _account) view returns((uint256,uint256,uint256))
func (_GlobalAccess *GlobalAccessCallerSession) GetUserInfo(_account common.Address) (GlobalAccessV1StorageUserInfo, error) {
	return _GlobalAccess.Contract.GetUserInfo(&_GlobalAccess.CallOpts, _account)
}

// GetUserNumber is a free data retrieval call binding the contract method 0x20c9ad2c.
//
// Solidity: function getUserNumber() view returns(uint256)
func (_GlobalAccess *GlobalAccessCaller) GetUserNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlobalAccess.contract.Call(opts, &out, "getUserNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserNumber is a free data retrieval call binding the contract method 0x20c9ad2c.
//
// Solidity: function getUserNumber() view returns(uint256)
func (_GlobalAccess *GlobalAccessSession) GetUserNumber() (*big.Int, error) {
	return _GlobalAccess.Contract.GetUserNumber(&_GlobalAccess.CallOpts)
}

// GetUserNumber is a free data retrieval call binding the contract method 0x20c9ad2c.
//
// Solidity: function getUserNumber() view returns(uint256)
func (_GlobalAccess *GlobalAccessCallerSession) GetUserNumber() (*big.Int, error) {
	return _GlobalAccess.Contract.GetUserNumber(&_GlobalAccess.CallOpts)
}

// GetUserStatus is a free data retrieval call binding the contract method 0xea0d5dcd.
//
// Solidity: function getUserStatus(address _account) view returns((uint8,uint256,uint256))
func (_GlobalAccess *GlobalAccessCaller) GetUserStatus(opts *bind.CallOpts, _account common.Address) (GlobalAccessV1StorageUserStatus, error) {
	var out []interface{}
	err := _GlobalAccess.contract.Call(opts, &out, "getUserStatus", _account)

	if err != nil {
		return *new(GlobalAccessV1StorageUserStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(GlobalAccessV1StorageUserStatus)).(*GlobalAccessV1StorageUserStatus)

	return out0, err

}

// GetUserStatus is a free data retrieval call binding the contract method 0xea0d5dcd.
//
// Solidity: function getUserStatus(address _account) view returns((uint8,uint256,uint256))
func (_GlobalAccess *GlobalAccessSession) GetUserStatus(_account common.Address) (GlobalAccessV1StorageUserStatus, error) {
	return _GlobalAccess.Contract.GetUserStatus(&_GlobalAccess.CallOpts, _account)
}

// GetUserStatus is a free data retrieval call binding the contract method 0xea0d5dcd.
//
// Solidity: function getUserStatus(address _account) view returns((uint8,uint256,uint256))
func (_GlobalAccess *GlobalAccessCallerSession) GetUserStatus(_account common.Address) (GlobalAccessV1StorageUserStatus, error) {
	return _GlobalAccess.Contract.GetUserStatus(&_GlobalAccess.CallOpts, _account)
}

// GlobalAccessImplementation is a free data retrieval call binding the contract method 0x296efb68.
//
// Solidity: function globalAccessImplementation() view returns(address)
func (_GlobalAccess *GlobalAccessCaller) GlobalAccessImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlobalAccess.contract.Call(opts, &out, "globalAccessImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalAccessImplementation is a free data retrieval call binding the contract method 0x296efb68.
//
// Solidity: function globalAccessImplementation() view returns(address)
func (_GlobalAccess *GlobalAccessSession) GlobalAccessImplementation() (common.Address, error) {
	return _GlobalAccess.Contract.GlobalAccessImplementation(&_GlobalAccess.CallOpts)
}

// GlobalAccessImplementation is a free data retrieval call binding the contract method 0x296efb68.
//
// Solidity: function globalAccessImplementation() view returns(address)
func (_GlobalAccess *GlobalAccessCallerSession) GlobalAccessImplementation() (common.Address, error) {
	return _GlobalAccess.Contract.GlobalAccessImplementation(&_GlobalAccess.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_GlobalAccess *GlobalAccessCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GlobalAccess.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_GlobalAccess *GlobalAccessSession) Operators(arg0 common.Address) (*big.Int, error) {
	return _GlobalAccess.Contract.Operators(&_GlobalAccess.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_GlobalAccess *GlobalAccessCallerSession) Operators(arg0 common.Address) (*big.Int, error) {
	return _GlobalAccess.Contract.Operators(&_GlobalAccess.CallOpts, arg0)
}

// PendingGlobalAccessImplementation is a free data retrieval call binding the contract method 0x2dc84d22.
//
// Solidity: function pendingGlobalAccessImplementation() view returns(address)
func (_GlobalAccess *GlobalAccessCaller) PendingGlobalAccessImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlobalAccess.contract.Call(opts, &out, "pendingGlobalAccessImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGlobalAccessImplementation is a free data retrieval call binding the contract method 0x2dc84d22.
//
// Solidity: function pendingGlobalAccessImplementation() view returns(address)
func (_GlobalAccess *GlobalAccessSession) PendingGlobalAccessImplementation() (common.Address, error) {
	return _GlobalAccess.Contract.PendingGlobalAccessImplementation(&_GlobalAccess.CallOpts)
}

// PendingGlobalAccessImplementation is a free data retrieval call binding the contract method 0x2dc84d22.
//
// Solidity: function pendingGlobalAccessImplementation() view returns(address)
func (_GlobalAccess *GlobalAccessCallerSession) PendingGlobalAccessImplementation() (common.Address, error) {
	return _GlobalAccess.Contract.PendingGlobalAccessImplementation(&_GlobalAccess.CallOpts)
}

// Quotas is a free data retrieval call binding the contract method 0xb6dea308.
//
// Solidity: function quotas(uint256 , uint256 ) view returns(uint256)
func (_GlobalAccess *GlobalAccessCaller) Quotas(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GlobalAccess.contract.Call(opts, &out, "quotas", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quotas is a free data retrieval call binding the contract method 0xb6dea308.
//
// Solidity: function quotas(uint256 , uint256 ) view returns(uint256)
func (_GlobalAccess *GlobalAccessSession) Quotas(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _GlobalAccess.Contract.Quotas(&_GlobalAccess.CallOpts, arg0, arg1)
}

// Quotas is a free data retrieval call binding the contract method 0xb6dea308.
//
// Solidity: function quotas(uint256 , uint256 ) view returns(uint256)
func (_GlobalAccess *GlobalAccessCallerSession) Quotas(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _GlobalAccess.Contract.Quotas(&_GlobalAccess.CallOpts, arg0, arg1)
}

// UserStatus is a free data retrieval call binding the contract method 0x225d29a1.
//
// Solidity: function userStatus(address ) view returns(uint8 status, uint256 pauseFrom, uint256 pausedTo)
func (_GlobalAccess *GlobalAccessCaller) UserStatus(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status    uint8
	PauseFrom *big.Int
	PausedTo  *big.Int
}, error) {
	var out []interface{}
	err := _GlobalAccess.contract.Call(opts, &out, "userStatus", arg0)

	outstruct := new(struct {
		Status    uint8
		PauseFrom *big.Int
		PausedTo  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.PauseFrom = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PausedTo = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserStatus is a free data retrieval call binding the contract method 0x225d29a1.
//
// Solidity: function userStatus(address ) view returns(uint8 status, uint256 pauseFrom, uint256 pausedTo)
func (_GlobalAccess *GlobalAccessSession) UserStatus(arg0 common.Address) (struct {
	Status    uint8
	PauseFrom *big.Int
	PausedTo  *big.Int
}, error) {
	return _GlobalAccess.Contract.UserStatus(&_GlobalAccess.CallOpts, arg0)
}

// UserStatus is a free data retrieval call binding the contract method 0x225d29a1.
//
// Solidity: function userStatus(address ) view returns(uint8 status, uint256 pauseFrom, uint256 pausedTo)
func (_GlobalAccess *GlobalAccessCallerSession) UserStatus(arg0 common.Address) (struct {
	Status    uint8
	PauseFrom *big.Int
	PausedTo  *big.Int
}, error) {
	return _GlobalAccess.Contract.UserStatus(&_GlobalAccess.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(uint256 idx, uint256 typ, uint256 level)
func (_GlobalAccess *GlobalAccessCaller) Users(opts *bind.CallOpts, arg0 common.Address) (struct {
	Idx   *big.Int
	Typ   *big.Int
	Level *big.Int
}, error) {
	var out []interface{}
	err := _GlobalAccess.contract.Call(opts, &out, "users", arg0)

	outstruct := new(struct {
		Idx   *big.Int
		Typ   *big.Int
		Level *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Idx = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Typ = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Level = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(uint256 idx, uint256 typ, uint256 level)
func (_GlobalAccess *GlobalAccessSession) Users(arg0 common.Address) (struct {
	Idx   *big.Int
	Typ   *big.Int
	Level *big.Int
}, error) {
	return _GlobalAccess.Contract.Users(&_GlobalAccess.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(uint256 idx, uint256 typ, uint256 level)
func (_GlobalAccess *GlobalAccessCallerSession) Users(arg0 common.Address) (struct {
	Idx   *big.Int
	Typ   *big.Int
	Level *big.Int
}, error) {
	return _GlobalAccess.Contract.Users(&_GlobalAccess.CallOpts, arg0)
}

// UsersList is a free data retrieval call binding the contract method 0x502aa3b5.
//
// Solidity: function usersList(uint256 ) view returns(address)
func (_GlobalAccess *GlobalAccessCaller) UsersList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GlobalAccess.contract.Call(opts, &out, "usersList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsersList is a free data retrieval call binding the contract method 0x502aa3b5.
//
// Solidity: function usersList(uint256 ) view returns(address)
func (_GlobalAccess *GlobalAccessSession) UsersList(arg0 *big.Int) (common.Address, error) {
	return _GlobalAccess.Contract.UsersList(&_GlobalAccess.CallOpts, arg0)
}

// UsersList is a free data retrieval call binding the contract method 0x502aa3b5.
//
// Solidity: function usersList(uint256 ) view returns(address)
func (_GlobalAccess *GlobalAccessCallerSession) UsersList(arg0 *big.Int) (common.Address, error) {
	return _GlobalAccess.Contract.UsersList(&_GlobalAccess.CallOpts, arg0)
}

// Become is a paid mutator transaction binding the contract method 0xa1194c8e.
//
// Solidity: function become(address _globalAccessProxy) returns()
func (_GlobalAccess *GlobalAccessTransactor) Become(opts *bind.TransactOpts, _globalAccessProxy common.Address) (*types.Transaction, error) {
	return _GlobalAccess.contract.Transact(opts, "become", _globalAccessProxy)
}

// Become is a paid mutator transaction binding the contract method 0xa1194c8e.
//
// Solidity: function become(address _globalAccessProxy) returns()
func (_GlobalAccess *GlobalAccessSession) Become(_globalAccessProxy common.Address) (*types.Transaction, error) {
	return _GlobalAccess.Contract.Become(&_GlobalAccess.TransactOpts, _globalAccessProxy)
}

// Become is a paid mutator transaction binding the contract method 0xa1194c8e.
//
// Solidity: function become(address _globalAccessProxy) returns()
func (_GlobalAccess *GlobalAccessTransactorSession) Become(_globalAccessProxy common.Address) (*types.Transaction, error) {
	return _GlobalAccess.Contract.Become(&_GlobalAccess.TransactOpts, _globalAccessProxy)
}

// SetQuota is a paid mutator transaction binding the contract method 0x42a9b01b.
//
// Solidity: function setQuota(uint256 _type, uint256 _level, uint256 _quota) returns()
func (_GlobalAccess *GlobalAccessTransactor) SetQuota(opts *bind.TransactOpts, _type *big.Int, _level *big.Int, _quota *big.Int) (*types.Transaction, error) {
	return _GlobalAccess.contract.Transact(opts, "setQuota", _type, _level, _quota)
}

// SetQuota is a paid mutator transaction binding the contract method 0x42a9b01b.
//
// Solidity: function setQuota(uint256 _type, uint256 _level, uint256 _quota) returns()
func (_GlobalAccess *GlobalAccessSession) SetQuota(_type *big.Int, _level *big.Int, _quota *big.Int) (*types.Transaction, error) {
	return _GlobalAccess.Contract.SetQuota(&_GlobalAccess.TransactOpts, _type, _level, _quota)
}

// SetQuota is a paid mutator transaction binding the contract method 0x42a9b01b.
//
// Solidity: function setQuota(uint256 _type, uint256 _level, uint256 _quota) returns()
func (_GlobalAccess *GlobalAccessTransactorSession) SetQuota(_type *big.Int, _level *big.Int, _quota *big.Int) (*types.Transaction, error) {
	return _GlobalAccess.Contract.SetQuota(&_GlobalAccess.TransactOpts, _type, _level, _quota)
}

// SetUserInfo is a paid mutator transaction binding the contract method 0xf5c972ac.
//
// Solidity: function setUserInfo(address _account, uint256 _type, uint256 _level) returns()
func (_GlobalAccess *GlobalAccessTransactor) SetUserInfo(opts *bind.TransactOpts, _account common.Address, _type *big.Int, _level *big.Int) (*types.Transaction, error) {
	return _GlobalAccess.contract.Transact(opts, "setUserInfo", _account, _type, _level)
}

// SetUserInfo is a paid mutator transaction binding the contract method 0xf5c972ac.
//
// Solidity: function setUserInfo(address _account, uint256 _type, uint256 _level) returns()
func (_GlobalAccess *GlobalAccessSession) SetUserInfo(_account common.Address, _type *big.Int, _level *big.Int) (*types.Transaction, error) {
	return _GlobalAccess.Contract.SetUserInfo(&_GlobalAccess.TransactOpts, _account, _type, _level)
}

// SetUserInfo is a paid mutator transaction binding the contract method 0xf5c972ac.
//
// Solidity: function setUserInfo(address _account, uint256 _type, uint256 _level) returns()
func (_GlobalAccess *GlobalAccessTransactorSession) SetUserInfo(_account common.Address, _type *big.Int, _level *big.Int) (*types.Transaction, error) {
	return _GlobalAccess.Contract.SetUserInfo(&_GlobalAccess.TransactOpts, _account, _type, _level)
}

// SetUserStatus is a paid mutator transaction binding the contract method 0xf376f731.
//
// Solidity: function setUserStatus(address _account, uint8 _status, uint256 _pausedFrom, uint256 _pausedTo) returns()
func (_GlobalAccess *GlobalAccessTransactor) SetUserStatus(opts *bind.TransactOpts, _account common.Address, _status uint8, _pausedFrom *big.Int, _pausedTo *big.Int) (*types.Transaction, error) {
	return _GlobalAccess.contract.Transact(opts, "setUserStatus", _account, _status, _pausedFrom, _pausedTo)
}

// SetUserStatus is a paid mutator transaction binding the contract method 0xf376f731.
//
// Solidity: function setUserStatus(address _account, uint8 _status, uint256 _pausedFrom, uint256 _pausedTo) returns()
func (_GlobalAccess *GlobalAccessSession) SetUserStatus(_account common.Address, _status uint8, _pausedFrom *big.Int, _pausedTo *big.Int) (*types.Transaction, error) {
	return _GlobalAccess.Contract.SetUserStatus(&_GlobalAccess.TransactOpts, _account, _status, _pausedFrom, _pausedTo)
}

// SetUserStatus is a paid mutator transaction binding the contract method 0xf376f731.
//
// Solidity: function setUserStatus(address _account, uint8 _status, uint256 _pausedFrom, uint256 _pausedTo) returns()
func (_GlobalAccess *GlobalAccessTransactorSession) SetUserStatus(_account common.Address, _status uint8, _pausedFrom *big.Int, _pausedTo *big.Int) (*types.Transaction, error) {
	return _GlobalAccess.Contract.SetUserStatus(&_GlobalAccess.TransactOpts, _account, _status, _pausedFrom, _pausedTo)
}

// GlobalAccessSetQuotaIterator is returned from FilterSetQuota and is used to iterate over the raw logs and unpacked data for SetQuota events raised by the GlobalAccess contract.
type GlobalAccessSetQuotaIterator struct {
	Event *GlobalAccessSetQuota // Event containing the contract specifics and raw log

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
func (it *GlobalAccessSetQuotaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalAccessSetQuota)
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
		it.Event = new(GlobalAccessSetQuota)
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
func (it *GlobalAccessSetQuotaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalAccessSetQuotaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalAccessSetQuota represents a SetQuota event raised by the GlobalAccess contract.
type GlobalAccessSetQuota struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetQuota is a free log retrieval operation binding the contract event 0xcca1bcbfa9c1a4a4a084c27886d999b229ab8ad838524621224f3d76318c6006.
//
// Solidity: event SetQuota(uint256 indexed arg0, uint256 indexed arg1, uint256 arg2)
func (_GlobalAccess *GlobalAccessFilterer) FilterSetQuota(opts *bind.FilterOpts, arg0 []*big.Int, arg1 []*big.Int) (*GlobalAccessSetQuotaIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _GlobalAccess.contract.FilterLogs(opts, "SetQuota", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return &GlobalAccessSetQuotaIterator{contract: _GlobalAccess.contract, event: "SetQuota", logs: logs, sub: sub}, nil
}

// WatchSetQuota is a free log subscription operation binding the contract event 0xcca1bcbfa9c1a4a4a084c27886d999b229ab8ad838524621224f3d76318c6006.
//
// Solidity: event SetQuota(uint256 indexed arg0, uint256 indexed arg1, uint256 arg2)
func (_GlobalAccess *GlobalAccessFilterer) WatchSetQuota(opts *bind.WatchOpts, sink chan<- *GlobalAccessSetQuota, arg0 []*big.Int, arg1 []*big.Int) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _GlobalAccess.contract.WatchLogs(opts, "SetQuota", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalAccessSetQuota)
				if err := _GlobalAccess.contract.UnpackLog(event, "SetQuota", log); err != nil {
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

// ParseSetQuota is a log parse operation binding the contract event 0xcca1bcbfa9c1a4a4a084c27886d999b229ab8ad838524621224f3d76318c6006.
//
// Solidity: event SetQuota(uint256 indexed arg0, uint256 indexed arg1, uint256 arg2)
func (_GlobalAccess *GlobalAccessFilterer) ParseSetQuota(log types.Log) (*GlobalAccessSetQuota, error) {
	event := new(GlobalAccessSetQuota)
	if err := _GlobalAccess.contract.UnpackLog(event, "SetQuota", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlobalAccessSetUserInfoIterator is returned from FilterSetUserInfo and is used to iterate over the raw logs and unpacked data for SetUserInfo events raised by the GlobalAccess contract.
type GlobalAccessSetUserInfoIterator struct {
	Event *GlobalAccessSetUserInfo // Event containing the contract specifics and raw log

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
func (it *GlobalAccessSetUserInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalAccessSetUserInfo)
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
		it.Event = new(GlobalAccessSetUserInfo)
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
func (it *GlobalAccessSetUserInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalAccessSetUserInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalAccessSetUserInfo represents a SetUserInfo event raised by the GlobalAccess contract.
type GlobalAccessSetUserInfo struct {
	Arg0 common.Address
	Arg1 *big.Int
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetUserInfo is a free log retrieval operation binding the contract event 0x93e871b638c7ad2419868ec623d797c1cd3f3f2f5b1af3c982f9d4a4a3902f7f.
//
// Solidity: event SetUserInfo(address indexed arg0, uint256 arg1, uint256 arg2)
func (_GlobalAccess *GlobalAccessFilterer) FilterSetUserInfo(opts *bind.FilterOpts, arg0 []common.Address) (*GlobalAccessSetUserInfoIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _GlobalAccess.contract.FilterLogs(opts, "SetUserInfo", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &GlobalAccessSetUserInfoIterator{contract: _GlobalAccess.contract, event: "SetUserInfo", logs: logs, sub: sub}, nil
}

// WatchSetUserInfo is a free log subscription operation binding the contract event 0x93e871b638c7ad2419868ec623d797c1cd3f3f2f5b1af3c982f9d4a4a3902f7f.
//
// Solidity: event SetUserInfo(address indexed arg0, uint256 arg1, uint256 arg2)
func (_GlobalAccess *GlobalAccessFilterer) WatchSetUserInfo(opts *bind.WatchOpts, sink chan<- *GlobalAccessSetUserInfo, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _GlobalAccess.contract.WatchLogs(opts, "SetUserInfo", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalAccessSetUserInfo)
				if err := _GlobalAccess.contract.UnpackLog(event, "SetUserInfo", log); err != nil {
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

// ParseSetUserInfo is a log parse operation binding the contract event 0x93e871b638c7ad2419868ec623d797c1cd3f3f2f5b1af3c982f9d4a4a3902f7f.
//
// Solidity: event SetUserInfo(address indexed arg0, uint256 arg1, uint256 arg2)
func (_GlobalAccess *GlobalAccessFilterer) ParseSetUserInfo(log types.Log) (*GlobalAccessSetUserInfo, error) {
	event := new(GlobalAccessSetUserInfo)
	if err := _GlobalAccess.contract.UnpackLog(event, "SetUserInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlobalAccessSetUserStatusIterator is returned from FilterSetUserStatus and is used to iterate over the raw logs and unpacked data for SetUserStatus events raised by the GlobalAccess contract.
type GlobalAccessSetUserStatusIterator struct {
	Event *GlobalAccessSetUserStatus // Event containing the contract specifics and raw log

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
func (it *GlobalAccessSetUserStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalAccessSetUserStatus)
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
		it.Event = new(GlobalAccessSetUserStatus)
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
func (it *GlobalAccessSetUserStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalAccessSetUserStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalAccessSetUserStatus represents a SetUserStatus event raised by the GlobalAccess contract.
type GlobalAccessSetUserStatus struct {
	Arg0 common.Address
	Arg1 uint8
	Arg2 *big.Int
	Arg3 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetUserStatus is a free log retrieval operation binding the contract event 0x209a9837686318dc2f386b0a03f6ce3237e62807daf38231e1295a2747bf1aa5.
//
// Solidity: event SetUserStatus(address indexed arg0, uint8 indexed arg1, uint256 arg2, uint256 arg3)
func (_GlobalAccess *GlobalAccessFilterer) FilterSetUserStatus(opts *bind.FilterOpts, arg0 []common.Address, arg1 []uint8) (*GlobalAccessSetUserStatusIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _GlobalAccess.contract.FilterLogs(opts, "SetUserStatus", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return &GlobalAccessSetUserStatusIterator{contract: _GlobalAccess.contract, event: "SetUserStatus", logs: logs, sub: sub}, nil
}

// WatchSetUserStatus is a free log subscription operation binding the contract event 0x209a9837686318dc2f386b0a03f6ce3237e62807daf38231e1295a2747bf1aa5.
//
// Solidity: event SetUserStatus(address indexed arg0, uint8 indexed arg1, uint256 arg2, uint256 arg3)
func (_GlobalAccess *GlobalAccessFilterer) WatchSetUserStatus(opts *bind.WatchOpts, sink chan<- *GlobalAccessSetUserStatus, arg0 []common.Address, arg1 []uint8) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _GlobalAccess.contract.WatchLogs(opts, "SetUserStatus", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalAccessSetUserStatus)
				if err := _GlobalAccess.contract.UnpackLog(event, "SetUserStatus", log); err != nil {
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

// ParseSetUserStatus is a log parse operation binding the contract event 0x209a9837686318dc2f386b0a03f6ce3237e62807daf38231e1295a2747bf1aa5.
//
// Solidity: event SetUserStatus(address indexed arg0, uint8 indexed arg1, uint256 arg2, uint256 arg3)
func (_GlobalAccess *GlobalAccessFilterer) ParseSetUserStatus(log types.Log) (*GlobalAccessSetUserStatus, error) {
	event := new(GlobalAccessSetUserStatus)
	if err := _GlobalAccess.contract.UnpackLog(event, "SetUserStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

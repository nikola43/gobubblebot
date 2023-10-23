// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pulsedoge

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

// PulsedogeMetaData contains all meta data concerning the Pulsedoge contract.
var PulsedogeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_managerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"addTokenHolders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnUnclaimedAmountDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnUnclaimedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getClaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"oldTokenHolderAlreadyClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_managerAddress\",\"type\":\"address\"}],\"name\":\"setManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalClaimedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PulsedogeABI is the input ABI used to generate the binding from.
// Deprecated: Use PulsedogeMetaData.ABI instead.
var PulsedogeABI = PulsedogeMetaData.ABI

// Pulsedoge is an auto generated Go binding around an Ethereum contract.
type Pulsedoge struct {
	PulsedogeCaller     // Read-only binding to the contract
	PulsedogeTransactor // Write-only binding to the contract
	PulsedogeFilterer   // Log filterer for contract events
}

// PulsedogeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PulsedogeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PulsedogeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PulsedogeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PulsedogeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PulsedogeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PulsedogeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PulsedogeSession struct {
	Contract     *Pulsedoge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PulsedogeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PulsedogeCallerSession struct {
	Contract *PulsedogeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PulsedogeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PulsedogeTransactorSession struct {
	Contract     *PulsedogeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PulsedogeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PulsedogeRaw struct {
	Contract *Pulsedoge // Generic contract binding to access the raw methods on
}

// PulsedogeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PulsedogeCallerRaw struct {
	Contract *PulsedogeCaller // Generic read-only contract binding to access the raw methods on
}

// PulsedogeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PulsedogeTransactorRaw struct {
	Contract *PulsedogeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPulsedoge creates a new instance of Pulsedoge, bound to a specific deployed contract.
func NewPulsedoge(address common.Address, backend bind.ContractBackend) (*Pulsedoge, error) {
	contract, err := bindPulsedoge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pulsedoge{PulsedogeCaller: PulsedogeCaller{contract: contract}, PulsedogeTransactor: PulsedogeTransactor{contract: contract}, PulsedogeFilterer: PulsedogeFilterer{contract: contract}}, nil
}

// NewPulsedogeCaller creates a new read-only instance of Pulsedoge, bound to a specific deployed contract.
func NewPulsedogeCaller(address common.Address, caller bind.ContractCaller) (*PulsedogeCaller, error) {
	contract, err := bindPulsedoge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PulsedogeCaller{contract: contract}, nil
}

// NewPulsedogeTransactor creates a new write-only instance of Pulsedoge, bound to a specific deployed contract.
func NewPulsedogeTransactor(address common.Address, transactor bind.ContractTransactor) (*PulsedogeTransactor, error) {
	contract, err := bindPulsedoge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PulsedogeTransactor{contract: contract}, nil
}

// NewPulsedogeFilterer creates a new log filterer instance of Pulsedoge, bound to a specific deployed contract.
func NewPulsedogeFilterer(address common.Address, filterer bind.ContractFilterer) (*PulsedogeFilterer, error) {
	contract, err := bindPulsedoge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PulsedogeFilterer{contract: contract}, nil
}

// bindPulsedoge binds a generic wrapper to an already deployed contract.
func bindPulsedoge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PulsedogeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pulsedoge *PulsedogeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pulsedoge.Contract.PulsedogeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pulsedoge *PulsedogeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pulsedoge.Contract.PulsedogeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pulsedoge *PulsedogeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pulsedoge.Contract.PulsedogeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pulsedoge *PulsedogeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pulsedoge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pulsedoge *PulsedogeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pulsedoge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pulsedoge *PulsedogeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pulsedoge.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pulsedoge *PulsedogeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pulsedoge.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pulsedoge *PulsedogeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Pulsedoge.Contract.Allowance(&_Pulsedoge.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pulsedoge *PulsedogeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Pulsedoge.Contract.Allowance(&_Pulsedoge.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Pulsedoge *PulsedogeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pulsedoge.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Pulsedoge *PulsedogeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Pulsedoge.Contract.BalanceOf(&_Pulsedoge.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Pulsedoge *PulsedogeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Pulsedoge.Contract.BalanceOf(&_Pulsedoge.CallOpts, account)
}

// BurnUnclaimedAmountDate is a free data retrieval call binding the contract method 0x58408eb6.
//
// Solidity: function burnUnclaimedAmountDate() view returns(uint256)
func (_Pulsedoge *PulsedogeCaller) BurnUnclaimedAmountDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pulsedoge.contract.Call(opts, &out, "burnUnclaimedAmountDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnUnclaimedAmountDate is a free data retrieval call binding the contract method 0x58408eb6.
//
// Solidity: function burnUnclaimedAmountDate() view returns(uint256)
func (_Pulsedoge *PulsedogeSession) BurnUnclaimedAmountDate() (*big.Int, error) {
	return _Pulsedoge.Contract.BurnUnclaimedAmountDate(&_Pulsedoge.CallOpts)
}

// BurnUnclaimedAmountDate is a free data retrieval call binding the contract method 0x58408eb6.
//
// Solidity: function burnUnclaimedAmountDate() view returns(uint256)
func (_Pulsedoge *PulsedogeCallerSession) BurnUnclaimedAmountDate() (*big.Int, error) {
	return _Pulsedoge.Contract.BurnUnclaimedAmountDate(&_Pulsedoge.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pulsedoge *PulsedogeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pulsedoge.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pulsedoge *PulsedogeSession) Decimals() (uint8, error) {
	return _Pulsedoge.Contract.Decimals(&_Pulsedoge.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pulsedoge *PulsedogeCallerSession) Decimals() (uint8, error) {
	return _Pulsedoge.Contract.Decimals(&_Pulsedoge.CallOpts)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address account) view returns(uint256)
func (_Pulsedoge *PulsedogeCaller) GetClaimableAmount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pulsedoge.contract.Call(opts, &out, "getClaimableAmount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address account) view returns(uint256)
func (_Pulsedoge *PulsedogeSession) GetClaimableAmount(account common.Address) (*big.Int, error) {
	return _Pulsedoge.Contract.GetClaimableAmount(&_Pulsedoge.CallOpts, account)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address account) view returns(uint256)
func (_Pulsedoge *PulsedogeCallerSession) GetClaimableAmount(account common.Address) (*big.Int, error) {
	return _Pulsedoge.Contract.GetClaimableAmount(&_Pulsedoge.CallOpts, account)
}

// ManagerAddress is a free data retrieval call binding the contract method 0xcf73a1bc.
//
// Solidity: function managerAddress() view returns(address)
func (_Pulsedoge *PulsedogeCaller) ManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pulsedoge.contract.Call(opts, &out, "managerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ManagerAddress is a free data retrieval call binding the contract method 0xcf73a1bc.
//
// Solidity: function managerAddress() view returns(address)
func (_Pulsedoge *PulsedogeSession) ManagerAddress() (common.Address, error) {
	return _Pulsedoge.Contract.ManagerAddress(&_Pulsedoge.CallOpts)
}

// ManagerAddress is a free data retrieval call binding the contract method 0xcf73a1bc.
//
// Solidity: function managerAddress() view returns(address)
func (_Pulsedoge *PulsedogeCallerSession) ManagerAddress() (common.Address, error) {
	return _Pulsedoge.Contract.ManagerAddress(&_Pulsedoge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pulsedoge *PulsedogeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pulsedoge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pulsedoge *PulsedogeSession) Name() (string, error) {
	return _Pulsedoge.Contract.Name(&_Pulsedoge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pulsedoge *PulsedogeCallerSession) Name() (string, error) {
	return _Pulsedoge.Contract.Name(&_Pulsedoge.CallOpts)
}

// OldTokenHolderAlreadyClaim is a free data retrieval call binding the contract method 0x965ba6bc.
//
// Solidity: function oldTokenHolderAlreadyClaim(uint256 id, address account) view returns(bool)
func (_Pulsedoge *PulsedogeCaller) OldTokenHolderAlreadyClaim(opts *bind.CallOpts, id *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _Pulsedoge.contract.Call(opts, &out, "oldTokenHolderAlreadyClaim", id, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OldTokenHolderAlreadyClaim is a free data retrieval call binding the contract method 0x965ba6bc.
//
// Solidity: function oldTokenHolderAlreadyClaim(uint256 id, address account) view returns(bool)
func (_Pulsedoge *PulsedogeSession) OldTokenHolderAlreadyClaim(id *big.Int, account common.Address) (bool, error) {
	return _Pulsedoge.Contract.OldTokenHolderAlreadyClaim(&_Pulsedoge.CallOpts, id, account)
}

// OldTokenHolderAlreadyClaim is a free data retrieval call binding the contract method 0x965ba6bc.
//
// Solidity: function oldTokenHolderAlreadyClaim(uint256 id, address account) view returns(bool)
func (_Pulsedoge *PulsedogeCallerSession) OldTokenHolderAlreadyClaim(id *big.Int, account common.Address) (bool, error) {
	return _Pulsedoge.Contract.OldTokenHolderAlreadyClaim(&_Pulsedoge.CallOpts, id, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pulsedoge *PulsedogeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pulsedoge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pulsedoge *PulsedogeSession) Owner() (common.Address, error) {
	return _Pulsedoge.Contract.Owner(&_Pulsedoge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pulsedoge *PulsedogeCallerSession) Owner() (common.Address, error) {
	return _Pulsedoge.Contract.Owner(&_Pulsedoge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pulsedoge *PulsedogeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pulsedoge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pulsedoge *PulsedogeSession) Symbol() (string, error) {
	return _Pulsedoge.Contract.Symbol(&_Pulsedoge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pulsedoge *PulsedogeCallerSession) Symbol() (string, error) {
	return _Pulsedoge.Contract.Symbol(&_Pulsedoge.CallOpts)
}

// TotalClaimedAmount is a free data retrieval call binding the contract method 0x9661cb0d.
//
// Solidity: function totalClaimedAmount() view returns(uint256)
func (_Pulsedoge *PulsedogeCaller) TotalClaimedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pulsedoge.contract.Call(opts, &out, "totalClaimedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimedAmount is a free data retrieval call binding the contract method 0x9661cb0d.
//
// Solidity: function totalClaimedAmount() view returns(uint256)
func (_Pulsedoge *PulsedogeSession) TotalClaimedAmount() (*big.Int, error) {
	return _Pulsedoge.Contract.TotalClaimedAmount(&_Pulsedoge.CallOpts)
}

// TotalClaimedAmount is a free data retrieval call binding the contract method 0x9661cb0d.
//
// Solidity: function totalClaimedAmount() view returns(uint256)
func (_Pulsedoge *PulsedogeCallerSession) TotalClaimedAmount() (*big.Int, error) {
	return _Pulsedoge.Contract.TotalClaimedAmount(&_Pulsedoge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pulsedoge *PulsedogeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pulsedoge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pulsedoge *PulsedogeSession) TotalSupply() (*big.Int, error) {
	return _Pulsedoge.Contract.TotalSupply(&_Pulsedoge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pulsedoge *PulsedogeCallerSession) TotalSupply() (*big.Int, error) {
	return _Pulsedoge.Contract.TotalSupply(&_Pulsedoge.CallOpts)
}

// AddTokenHolders is a paid mutator transaction binding the contract method 0x91a4965e.
//
// Solidity: function addTokenHolders(address[] accounts, uint256[] amounts) returns()
func (_Pulsedoge *PulsedogeTransactor) AddTokenHolders(opts *bind.TransactOpts, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Pulsedoge.contract.Transact(opts, "addTokenHolders", accounts, amounts)
}

// AddTokenHolders is a paid mutator transaction binding the contract method 0x91a4965e.
//
// Solidity: function addTokenHolders(address[] accounts, uint256[] amounts) returns()
func (_Pulsedoge *PulsedogeSession) AddTokenHolders(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.AddTokenHolders(&_Pulsedoge.TransactOpts, accounts, amounts)
}

// AddTokenHolders is a paid mutator transaction binding the contract method 0x91a4965e.
//
// Solidity: function addTokenHolders(address[] accounts, uint256[] amounts) returns()
func (_Pulsedoge *PulsedogeTransactorSession) AddTokenHolders(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.AddTokenHolders(&_Pulsedoge.TransactOpts, accounts, amounts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pulsedoge *PulsedogeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pulsedoge *PulsedogeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.Approve(&_Pulsedoge.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pulsedoge *PulsedogeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.Approve(&_Pulsedoge.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Pulsedoge *PulsedogeTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Pulsedoge *PulsedogeSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.Burn(&_Pulsedoge.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Pulsedoge *PulsedogeTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.Burn(&_Pulsedoge.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Pulsedoge *PulsedogeTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Pulsedoge *PulsedogeSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.BurnFrom(&_Pulsedoge.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Pulsedoge *PulsedogeTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.BurnFrom(&_Pulsedoge.TransactOpts, account, amount)
}

// BurnUnclaimedTokens is a paid mutator transaction binding the contract method 0x70a16aa7.
//
// Solidity: function burnUnclaimedTokens() returns()
func (_Pulsedoge *PulsedogeTransactor) BurnUnclaimedTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pulsedoge.contract.Transact(opts, "burnUnclaimedTokens")
}

// BurnUnclaimedTokens is a paid mutator transaction binding the contract method 0x70a16aa7.
//
// Solidity: function burnUnclaimedTokens() returns()
func (_Pulsedoge *PulsedogeSession) BurnUnclaimedTokens() (*types.Transaction, error) {
	return _Pulsedoge.Contract.BurnUnclaimedTokens(&_Pulsedoge.TransactOpts)
}

// BurnUnclaimedTokens is a paid mutator transaction binding the contract method 0x70a16aa7.
//
// Solidity: function burnUnclaimedTokens() returns()
func (_Pulsedoge *PulsedogeTransactorSession) BurnUnclaimedTokens() (*types.Transaction, error) {
	return _Pulsedoge.Contract.BurnUnclaimedTokens(&_Pulsedoge.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Pulsedoge *PulsedogeTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pulsedoge.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Pulsedoge *PulsedogeSession) Claim() (*types.Transaction, error) {
	return _Pulsedoge.Contract.Claim(&_Pulsedoge.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Pulsedoge *PulsedogeTransactorSession) Claim() (*types.Transaction, error) {
	return _Pulsedoge.Contract.Claim(&_Pulsedoge.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Pulsedoge *PulsedogeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Pulsedoge *PulsedogeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.DecreaseAllowance(&_Pulsedoge.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Pulsedoge *PulsedogeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.DecreaseAllowance(&_Pulsedoge.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Pulsedoge *PulsedogeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Pulsedoge *PulsedogeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.IncreaseAllowance(&_Pulsedoge.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Pulsedoge *PulsedogeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.IncreaseAllowance(&_Pulsedoge.TransactOpts, spender, addedValue)
}

// SetManagerAddress is a paid mutator transaction binding the contract method 0x41431908.
//
// Solidity: function setManagerAddress(address _managerAddress) returns()
func (_Pulsedoge *PulsedogeTransactor) SetManagerAddress(opts *bind.TransactOpts, _managerAddress common.Address) (*types.Transaction, error) {
	return _Pulsedoge.contract.Transact(opts, "setManagerAddress", _managerAddress)
}

// SetManagerAddress is a paid mutator transaction binding the contract method 0x41431908.
//
// Solidity: function setManagerAddress(address _managerAddress) returns()
func (_Pulsedoge *PulsedogeSession) SetManagerAddress(_managerAddress common.Address) (*types.Transaction, error) {
	return _Pulsedoge.Contract.SetManagerAddress(&_Pulsedoge.TransactOpts, _managerAddress)
}

// SetManagerAddress is a paid mutator transaction binding the contract method 0x41431908.
//
// Solidity: function setManagerAddress(address _managerAddress) returns()
func (_Pulsedoge *PulsedogeTransactorSession) SetManagerAddress(_managerAddress common.Address) (*types.Transaction, error) {
	return _Pulsedoge.Contract.SetManagerAddress(&_Pulsedoge.TransactOpts, _managerAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Pulsedoge *PulsedogeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Pulsedoge *PulsedogeSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.Transfer(&_Pulsedoge.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Pulsedoge *PulsedogeTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.Transfer(&_Pulsedoge.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Pulsedoge *PulsedogeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Pulsedoge *PulsedogeSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.TransferFrom(&_Pulsedoge.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Pulsedoge *PulsedogeTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pulsedoge.Contract.TransferFrom(&_Pulsedoge.TransactOpts, from, to, amount)
}

// PulsedogeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pulsedoge contract.
type PulsedogeApprovalIterator struct {
	Event *PulsedogeApproval // Event containing the contract specifics and raw log

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
func (it *PulsedogeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PulsedogeApproval)
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
		it.Event = new(PulsedogeApproval)
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
func (it *PulsedogeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PulsedogeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PulsedogeApproval represents a Approval event raised by the Pulsedoge contract.
type PulsedogeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pulsedoge *PulsedogeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PulsedogeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pulsedoge.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PulsedogeApprovalIterator{contract: _Pulsedoge.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pulsedoge *PulsedogeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PulsedogeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pulsedoge.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PulsedogeApproval)
				if err := _Pulsedoge.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pulsedoge *PulsedogeFilterer) ParseApproval(log types.Log) (*PulsedogeApproval, error) {
	event := new(PulsedogeApproval)
	if err := _Pulsedoge.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PulsedogeClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Pulsedoge contract.
type PulsedogeClaimedIterator struct {
	Event *PulsedogeClaimed // Event containing the contract specifics and raw log

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
func (it *PulsedogeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PulsedogeClaimed)
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
		it.Event = new(PulsedogeClaimed)
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
func (it *PulsedogeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PulsedogeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PulsedogeClaimed represents a Claimed event raised by the Pulsedoge contract.
type PulsedogeClaimed struct {
	Holder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed holder, uint256 amount)
func (_Pulsedoge *PulsedogeFilterer) FilterClaimed(opts *bind.FilterOpts, holder []common.Address) (*PulsedogeClaimedIterator, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}

	logs, sub, err := _Pulsedoge.contract.FilterLogs(opts, "Claimed", holderRule)
	if err != nil {
		return nil, err
	}
	return &PulsedogeClaimedIterator{contract: _Pulsedoge.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed holder, uint256 amount)
func (_Pulsedoge *PulsedogeFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *PulsedogeClaimed, holder []common.Address) (event.Subscription, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}

	logs, sub, err := _Pulsedoge.contract.WatchLogs(opts, "Claimed", holderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PulsedogeClaimed)
				if err := _Pulsedoge.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed holder, uint256 amount)
func (_Pulsedoge *PulsedogeFilterer) ParseClaimed(log types.Log) (*PulsedogeClaimed, error) {
	event := new(PulsedogeClaimed)
	if err := _Pulsedoge.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PulsedogeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pulsedoge contract.
type PulsedogeTransferIterator struct {
	Event *PulsedogeTransfer // Event containing the contract specifics and raw log

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
func (it *PulsedogeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PulsedogeTransfer)
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
		it.Event = new(PulsedogeTransfer)
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
func (it *PulsedogeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PulsedogeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PulsedogeTransfer represents a Transfer event raised by the Pulsedoge contract.
type PulsedogeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pulsedoge *PulsedogeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PulsedogeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pulsedoge.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PulsedogeTransferIterator{contract: _Pulsedoge.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pulsedoge *PulsedogeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PulsedogeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pulsedoge.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PulsedogeTransfer)
				if err := _Pulsedoge.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pulsedoge *PulsedogeFilterer) ParseTransfer(log types.Log) (*PulsedogeTransfer, error) {
	event := new(PulsedogeTransfer)
	if err := _Pulsedoge.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

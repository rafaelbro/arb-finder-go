// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitas

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ArbitasABI is the input ABI used to generate the binding from.
const ArbitasABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"routers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inPairAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountBorrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"routerPath\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPath\",\"type\":\"address[]\"}],\"name\":\"startArbitrage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"pancakeCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"waultSwapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"routerAddress\",\"type\":\"address\"}],\"name\":\"addRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRouterIn\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"routerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Arbitas is an auto generated Go binding around an Ethereum contract.
type Arbitas struct {
	ArbitasCaller     // Read-only binding to the contract
	ArbitasTransactor // Write-only binding to the contract
	ArbitasFilterer   // Log filterer for contract events
}

// ArbitasCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbitasCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitasTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbitasTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitasFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbitasFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbitasSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbitasSession struct {
	Contract     *Arbitas          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbitasCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbitasCallerSession struct {
	Contract *ArbitasCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ArbitasTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbitasTransactorSession struct {
	Contract     *ArbitasTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ArbitasRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbitasRaw struct {
	Contract *Arbitas // Generic contract binding to access the raw methods on
}

// ArbitasCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbitasCallerRaw struct {
	Contract *ArbitasCaller // Generic read-only contract binding to access the raw methods on
}

// ArbitasTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbitasTransactorRaw struct {
	Contract *ArbitasTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbitas creates a new instance of Arbitas, bound to a specific deployed contract.
func NewArbitas(address common.Address, backend bind.ContractBackend) (*Arbitas, error) {
	contract, err := bindArbitas(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arbitas{ArbitasCaller: ArbitasCaller{contract: contract}, ArbitasTransactor: ArbitasTransactor{contract: contract}, ArbitasFilterer: ArbitasFilterer{contract: contract}}, nil
}

// NewArbitasCaller creates a new read-only instance of Arbitas, bound to a specific deployed contract.
func NewArbitasCaller(address common.Address, caller bind.ContractCaller) (*ArbitasCaller, error) {
	contract, err := bindArbitas(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitasCaller{contract: contract}, nil
}

// NewArbitasTransactor creates a new write-only instance of Arbitas, bound to a specific deployed contract.
func NewArbitasTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitasTransactor, error) {
	contract, err := bindArbitas(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitasTransactor{contract: contract}, nil
}

// NewArbitasFilterer creates a new log filterer instance of Arbitas, bound to a specific deployed contract.
func NewArbitasFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitasFilterer, error) {
	contract, err := bindArbitas(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitasFilterer{contract: contract}, nil
}

// bindArbitas binds a generic wrapper to an already deployed contract.
func bindArbitas(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbitasABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arbitas *ArbitasRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arbitas.Contract.ArbitasCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arbitas *ArbitasRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitas.Contract.ArbitasTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arbitas *ArbitasRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arbitas.Contract.ArbitasTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arbitas *ArbitasCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arbitas.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arbitas *ArbitasTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arbitas.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arbitas *ArbitasTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arbitas.Contract.contract.Transact(opts, method, params...)
}

// GetRouterIn is a free data retrieval call binding the contract method 0x40853a54.
//
// Solidity: function getRouterIn(uint256 index) view returns(address routerAddress)
func (_Arbitas *ArbitasCaller) GetRouterIn(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Arbitas.contract.Call(opts, &out, "getRouterIn", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRouterIn is a free data retrieval call binding the contract method 0x40853a54.
//
// Solidity: function getRouterIn(uint256 index) view returns(address routerAddress)
func (_Arbitas *ArbitasSession) GetRouterIn(index *big.Int) (common.Address, error) {
	return _Arbitas.Contract.GetRouterIn(&_Arbitas.CallOpts, index)
}

// GetRouterIn is a free data retrieval call binding the contract method 0x40853a54.
//
// Solidity: function getRouterIn(uint256 index) view returns(address routerAddress)
func (_Arbitas *ArbitasCallerSession) GetRouterIn(index *big.Int) (common.Address, error) {
	return _Arbitas.Contract.GetRouterIn(&_Arbitas.CallOpts, index)
}

// AddRouter is a paid mutator transaction binding the contract method 0x176e87df.
//
// Solidity: function addRouter(uint256 index, address routerAddress) returns()
func (_Arbitas *ArbitasTransactor) AddRouter(opts *bind.TransactOpts, index *big.Int, routerAddress common.Address) (*types.Transaction, error) {
	return _Arbitas.contract.Transact(opts, "addRouter", index, routerAddress)
}

// AddRouter is a paid mutator transaction binding the contract method 0x176e87df.
//
// Solidity: function addRouter(uint256 index, address routerAddress) returns()
func (_Arbitas *ArbitasSession) AddRouter(index *big.Int, routerAddress common.Address) (*types.Transaction, error) {
	return _Arbitas.Contract.AddRouter(&_Arbitas.TransactOpts, index, routerAddress)
}

// AddRouter is a paid mutator transaction binding the contract method 0x176e87df.
//
// Solidity: function addRouter(uint256 index, address routerAddress) returns()
func (_Arbitas *ArbitasTransactorSession) AddRouter(index *big.Int, routerAddress common.Address) (*types.Transaction, error) {
	return _Arbitas.Contract.AddRouter(&_Arbitas.TransactOpts, index, routerAddress)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_Arbitas *ArbitasTransactor) PancakeCall(opts *bind.TransactOpts, _sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Arbitas.contract.Transact(opts, "pancakeCall", _sender, _amount0, _amount1, _data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_Arbitas *ArbitasSession) PancakeCall(_sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Arbitas.Contract.PancakeCall(&_Arbitas.TransactOpts, _sender, _amount0, _amount1, _data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_Arbitas *ArbitasTransactorSession) PancakeCall(_sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Arbitas.Contract.PancakeCall(&_Arbitas.TransactOpts, _sender, _amount0, _amount1, _data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x38bf3cfa.
//
// Solidity: function renounceOwnership(address newOwner) returns()
func (_Arbitas *ArbitasTransactor) RenounceOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Arbitas.contract.Transact(opts, "renounceOwnership", newOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x38bf3cfa.
//
// Solidity: function renounceOwnership(address newOwner) returns()
func (_Arbitas *ArbitasSession) RenounceOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Arbitas.Contract.RenounceOwnership(&_Arbitas.TransactOpts, newOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x38bf3cfa.
//
// Solidity: function renounceOwnership(address newOwner) returns()
func (_Arbitas *ArbitasTransactorSession) RenounceOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Arbitas.Contract.RenounceOwnership(&_Arbitas.TransactOpts, newOwner)
}

// StartArbitrage is a paid mutator transaction binding the contract method 0x9653a422.
//
// Solidity: function startArbitrage(address inPairAddress, uint256 amountBorrowed, uint256[] routerPath, address[] tokenPath) returns()
func (_Arbitas *ArbitasTransactor) StartArbitrage(opts *bind.TransactOpts, inPairAddress common.Address, amountBorrowed *big.Int, routerPath []*big.Int, tokenPath []common.Address) (*types.Transaction, error) {
	return _Arbitas.contract.Transact(opts, "startArbitrage", inPairAddress, amountBorrowed, routerPath, tokenPath)
}

// StartArbitrage is a paid mutator transaction binding the contract method 0x9653a422.
//
// Solidity: function startArbitrage(address inPairAddress, uint256 amountBorrowed, uint256[] routerPath, address[] tokenPath) returns()
func (_Arbitas *ArbitasSession) StartArbitrage(inPairAddress common.Address, amountBorrowed *big.Int, routerPath []*big.Int, tokenPath []common.Address) (*types.Transaction, error) {
	return _Arbitas.Contract.StartArbitrage(&_Arbitas.TransactOpts, inPairAddress, amountBorrowed, routerPath, tokenPath)
}

// StartArbitrage is a paid mutator transaction binding the contract method 0x9653a422.
//
// Solidity: function startArbitrage(address inPairAddress, uint256 amountBorrowed, uint256[] routerPath, address[] tokenPath) returns()
func (_Arbitas *ArbitasTransactorSession) StartArbitrage(inPairAddress common.Address, amountBorrowed *big.Int, routerPath []*big.Int, tokenPath []common.Address) (*types.Transaction, error) {
	return _Arbitas.Contract.StartArbitrage(&_Arbitas.TransactOpts, inPairAddress, amountBorrowed, routerPath, tokenPath)
}

// WaultSwapCall is a paid mutator transaction binding the contract method 0x485f3994.
//
// Solidity: function waultSwapCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_Arbitas *ArbitasTransactor) WaultSwapCall(opts *bind.TransactOpts, _sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Arbitas.contract.Transact(opts, "waultSwapCall", _sender, _amount0, _amount1, _data)
}

// WaultSwapCall is a paid mutator transaction binding the contract method 0x485f3994.
//
// Solidity: function waultSwapCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_Arbitas *ArbitasSession) WaultSwapCall(_sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Arbitas.Contract.WaultSwapCall(&_Arbitas.TransactOpts, _sender, _amount0, _amount1, _data)
}

// WaultSwapCall is a paid mutator transaction binding the contract method 0x485f3994.
//
// Solidity: function waultSwapCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_Arbitas *ArbitasTransactorSession) WaultSwapCall(_sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Arbitas.Contract.WaultSwapCall(&_Arbitas.TransactOpts, _sender, _amount0, _amount1, _data)
}

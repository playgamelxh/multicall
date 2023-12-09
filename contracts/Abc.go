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
	_ = abi.ConvertType
)

// AbcMetaData contains all meta data concerning the Abc contract.
var AbcMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"executeUSDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paydb\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenToDo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AbcABI is the input ABI used to generate the binding from.
// Deprecated: Use AbcMetaData.ABI instead.
var AbcABI = AbcMetaData.ABI

// Abc is an auto generated Go binding around an Ethereum contract.
type Abc struct {
	AbcCaller     // Read-only binding to the contract
	AbcTransactor // Write-only binding to the contract
	AbcFilterer   // Log filterer for contract events
}

// AbcCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbcSession struct {
	Contract     *Abc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbcCallerSession struct {
	Contract *AbcCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbcTransactorSession struct {
	Contract     *AbcTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbcRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbcRaw struct {
	Contract *Abc // Generic contract binding to access the raw methods on
}

// AbcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbcCallerRaw struct {
	Contract *AbcCaller // Generic read-only contract binding to access the raw methods on
}

// AbcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbcTransactorRaw struct {
	Contract *AbcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbc creates a new instance of Abc, bound to a specific deployed contract.
func NewAbc(address common.Address, backend bind.ContractBackend) (*Abc, error) {
	contract, err := bindAbc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abc{AbcCaller: AbcCaller{contract: contract}, AbcTransactor: AbcTransactor{contract: contract}, AbcFilterer: AbcFilterer{contract: contract}}, nil
}

// NewAbcCaller creates a new read-only instance of Abc, bound to a specific deployed contract.
func NewAbcCaller(address common.Address, caller bind.ContractCaller) (*AbcCaller, error) {
	contract, err := bindAbc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbcCaller{contract: contract}, nil
}

// NewAbcTransactor creates a new write-only instance of Abc, bound to a specific deployed contract.
func NewAbcTransactor(address common.Address, transactor bind.ContractTransactor) (*AbcTransactor, error) {
	contract, err := bindAbc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbcTransactor{contract: contract}, nil
}

// NewAbcFilterer creates a new log filterer instance of Abc, bound to a specific deployed contract.
func NewAbcFilterer(address common.Address, filterer bind.ContractFilterer) (*AbcFilterer, error) {
	contract, err := bindAbc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbcFilterer{contract: contract}, nil
}

// bindAbc binds a generic wrapper to an already deployed contract.
func bindAbc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abc *AbcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abc.Contract.AbcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abc *AbcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abc.Contract.AbcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abc *AbcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abc.Contract.AbcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abc *AbcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abc *AbcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abc *AbcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abc.Contract.contract.Transact(opts, method, params...)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Abc *AbcCaller) Limit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abc.contract.Call(opts, &out, "limit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Abc *AbcSession) Limit() (*big.Int, error) {
	return _Abc.Contract.Limit(&_Abc.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Abc *AbcCallerSession) Limit() (*big.Int, error) {
	return _Abc.Contract.Limit(&_Abc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abc *AbcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abc *AbcSession) Owner() (common.Address, error) {
	return _Abc.Contract.Owner(&_Abc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abc *AbcCallerSession) Owner() (common.Address, error) {
	return _Abc.Contract.Owner(&_Abc.CallOpts)
}

// Paydb is a free data retrieval call binding the contract method 0x5606d9de.
//
// Solidity: function paydb() view returns(address)
func (_Abc *AbcCaller) Paydb(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abc.contract.Call(opts, &out, "paydb")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Paydb is a free data retrieval call binding the contract method 0x5606d9de.
//
// Solidity: function paydb() view returns(address)
func (_Abc *AbcSession) Paydb() (common.Address, error) {
	return _Abc.Contract.Paydb(&_Abc.CallOpts)
}

// Paydb is a free data retrieval call binding the contract method 0x5606d9de.
//
// Solidity: function paydb() view returns(address)
func (_Abc *AbcCallerSession) Paydb() (common.Address, error) {
	return _Abc.Contract.Paydb(&_Abc.CallOpts)
}

// TokenToDo is a free data retrieval call binding the contract method 0x0869820b.
//
// Solidity: function tokenToDo() view returns(address)
func (_Abc *AbcCaller) TokenToDo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abc.contract.Call(opts, &out, "tokenToDo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenToDo is a free data retrieval call binding the contract method 0x0869820b.
//
// Solidity: function tokenToDo() view returns(address)
func (_Abc *AbcSession) TokenToDo() (common.Address, error) {
	return _Abc.Contract.TokenToDo(&_Abc.CallOpts)
}

// TokenToDo is a free data retrieval call binding the contract method 0x0869820b.
//
// Solidity: function tokenToDo() view returns(address)
func (_Abc *AbcCallerSession) TokenToDo() (common.Address, error) {
	return _Abc.Contract.TokenToDo(&_Abc.CallOpts)
}

// ExecuteUSDC is a paid mutator transaction binding the contract method 0x8af8ba6f.
//
// Solidity: function executeUSDC(address[] users) returns()
func (_Abc *AbcTransactor) ExecuteUSDC(opts *bind.TransactOpts, users []common.Address) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "executeUSDC", users)
}

// ExecuteUSDC is a paid mutator transaction binding the contract method 0x8af8ba6f.
//
// Solidity: function executeUSDC(address[] users) returns()
func (_Abc *AbcSession) ExecuteUSDC(users []common.Address) (*types.Transaction, error) {
	return _Abc.Contract.ExecuteUSDC(&_Abc.TransactOpts, users)
}

// ExecuteUSDC is a paid mutator transaction binding the contract method 0x8af8ba6f.
//
// Solidity: function executeUSDC(address[] users) returns()
func (_Abc *AbcTransactorSession) ExecuteUSDC(users []common.Address) (*types.Transaction, error) {
	return _Abc.Contract.ExecuteUSDC(&_Abc.TransactOpts, users)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abc *AbcTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abc *AbcSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abc.Contract.RenounceOwnership(&_Abc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abc *AbcTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abc.Contract.RenounceOwnership(&_Abc.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x3825d828.
//
// Solidity: function set(address _token, uint256 _limit) returns()
func (_Abc *AbcTransactor) Set(opts *bind.TransactOpts, _token common.Address, _limit *big.Int) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "set", _token, _limit)
}

// Set is a paid mutator transaction binding the contract method 0x3825d828.
//
// Solidity: function set(address _token, uint256 _limit) returns()
func (_Abc *AbcSession) Set(_token common.Address, _limit *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.Set(&_Abc.TransactOpts, _token, _limit)
}

// Set is a paid mutator transaction binding the contract method 0x3825d828.
//
// Solidity: function set(address _token, uint256 _limit) returns()
func (_Abc *AbcTransactorSession) Set(_token common.Address, _limit *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.Set(&_Abc.TransactOpts, _token, _limit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abc *AbcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abc *AbcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abc.Contract.TransferOwnership(&_Abc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abc *AbcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abc.Contract.TransferOwnership(&_Abc.TransactOpts, newOwner)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address account, uint256 amount) returns()
func (_Abc *AbcTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "withdrawToken", token, account, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address account, uint256 amount) returns()
func (_Abc *AbcSession) WithdrawToken(token common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.WithdrawToken(&_Abc.TransactOpts, token, account, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address account, uint256 amount) returns()
func (_Abc *AbcTransactorSession) WithdrawToken(token common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.WithdrawToken(&_Abc.TransactOpts, token, account, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abc *AbcTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Abc.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abc *AbcSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Abc.Contract.Fallback(&_Abc.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abc *AbcTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Abc.Contract.Fallback(&_Abc.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abc *AbcTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abc.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abc *AbcSession) Receive() (*types.Transaction, error) {
	return _Abc.Contract.Receive(&_Abc.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abc *AbcTransactorSession) Receive() (*types.Transaction, error) {
	return _Abc.Contract.Receive(&_Abc.TransactOpts)
}

// AbcOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Abc contract.
type AbcOwnershipTransferredIterator struct {
	Event *AbcOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AbcOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbcOwnershipTransferred)
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
		it.Event = new(AbcOwnershipTransferred)
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
func (it *AbcOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbcOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbcOwnershipTransferred represents a OwnershipTransferred event raised by the Abc contract.
type AbcOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abc *AbcFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AbcOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AbcOwnershipTransferredIterator{contract: _Abc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abc *AbcFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AbcOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbcOwnershipTransferred)
				if err := _Abc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abc *AbcFilterer) ParseOwnershipTransferred(log types.Log) (*AbcOwnershipTransferred, error) {
	event := new(AbcOwnershipTransferred)
	if err := _Abc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TxState

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

// TxStateTransactionInfo is an auto generated low-level Go binding around an user-defined struct.
type TxStateTransactionInfo struct {
	ChainId  uint64
	From     common.Address
	SeqNum   uint64
	Receiver common.Address
	Amount   *big.Int
	State    uint8
	Data     []byte
}

// TxStateMetaData contains all meta data concerning the TxState contract.
var TxStateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"seqNum\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"seqNum\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumTxState.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTxState.TransactionInfo\",\"name\":\"txInfo\",\"type\":\"tuple\"}],\"name\":\"L1transferEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"L1Txhash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"updateTxStateSuccess\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSITOR_ACCOUNT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"StoreL1Txhash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_seqNum\",\"type\":\"uint64\"}],\"name\":\"getL1Txhash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_seqNum\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"proposeTxToL1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_txHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"l2Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_seqNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_state\",\"type\":\"uint256\"}],\"name\":\"setL1TxState\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TxStateABI is the input ABI used to generate the binding from.
// Deprecated: Use TxStateMetaData.ABI instead.
var TxStateABI = TxStateMetaData.ABI

// TxState is an auto generated Go binding around an Ethereum contract.
type TxState struct {
	TxStateCaller     // Read-only binding to the contract
	TxStateTransactor // Write-only binding to the contract
	TxStateFilterer   // Log filterer for contract events
}

// TxStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type TxStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TxStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TxStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TxStateSession struct {
	Contract     *TxState          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TxStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TxStateCallerSession struct {
	Contract *TxStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TxStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TxStateTransactorSession struct {
	Contract     *TxStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TxStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type TxStateRaw struct {
	Contract *TxState // Generic contract binding to access the raw methods on
}

// TxStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TxStateCallerRaw struct {
	Contract *TxStateCaller // Generic read-only contract binding to access the raw methods on
}

// TxStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TxStateTransactorRaw struct {
	Contract *TxStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTxState creates a new instance of TxState, bound to a specific deployed contract.
func NewTxState(address common.Address, backend bind.ContractBackend) (*TxState, error) {
	contract, err := bindTxState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TxState{TxStateCaller: TxStateCaller{contract: contract}, TxStateTransactor: TxStateTransactor{contract: contract}, TxStateFilterer: TxStateFilterer{contract: contract}}, nil
}

// NewTxStateCaller creates a new read-only instance of TxState, bound to a specific deployed contract.
func NewTxStateCaller(address common.Address, caller bind.ContractCaller) (*TxStateCaller, error) {
	contract, err := bindTxState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TxStateCaller{contract: contract}, nil
}

// NewTxStateTransactor creates a new write-only instance of TxState, bound to a specific deployed contract.
func NewTxStateTransactor(address common.Address, transactor bind.ContractTransactor) (*TxStateTransactor, error) {
	contract, err := bindTxState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TxStateTransactor{contract: contract}, nil
}

// NewTxStateFilterer creates a new log filterer instance of TxState, bound to a specific deployed contract.
func NewTxStateFilterer(address common.Address, filterer bind.ContractFilterer) (*TxStateFilterer, error) {
	contract, err := bindTxState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TxStateFilterer{contract: contract}, nil
}

// bindTxState binds a generic wrapper to an already deployed contract.
func bindTxState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TxStateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxState *TxStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TxState.Contract.TxStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxState *TxStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxState.Contract.TxStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxState *TxStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxState.Contract.TxStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxState *TxStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TxState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxState *TxStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxState *TxStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxState.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_TxState *TxStateCaller) DEPOSITORACCOUNT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TxState.contract.Call(opts, &out, "DEPOSITOR_ACCOUNT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_TxState *TxStateSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _TxState.Contract.DEPOSITORACCOUNT(&_TxState.CallOpts)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_TxState *TxStateCallerSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _TxState.Contract.DEPOSITORACCOUNT(&_TxState.CallOpts)
}

// StoreL1Txhash is a free data retrieval call binding the contract method 0x996a627d.
//
// Solidity: function StoreL1Txhash(address , uint64 ) view returns(bytes)
func (_TxState *TxStateCaller) StoreL1Txhash(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) ([]byte, error) {
	var out []interface{}
	err := _TxState.contract.Call(opts, &out, "StoreL1Txhash", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// StoreL1Txhash is a free data retrieval call binding the contract method 0x996a627d.
//
// Solidity: function StoreL1Txhash(address , uint64 ) view returns(bytes)
func (_TxState *TxStateSession) StoreL1Txhash(arg0 common.Address, arg1 uint64) ([]byte, error) {
	return _TxState.Contract.StoreL1Txhash(&_TxState.CallOpts, arg0, arg1)
}

// StoreL1Txhash is a free data retrieval call binding the contract method 0x996a627d.
//
// Solidity: function StoreL1Txhash(address , uint64 ) view returns(bytes)
func (_TxState *TxStateCallerSession) StoreL1Txhash(arg0 common.Address, arg1 uint64) ([]byte, error) {
	return _TxState.Contract.StoreL1Txhash(&_TxState.CallOpts, arg0, arg1)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TxState *TxStateCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TxState.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TxState *TxStateSession) EntryPoint() (common.Address, error) {
	return _TxState.Contract.EntryPoint(&_TxState.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_TxState *TxStateCallerSession) EntryPoint() (common.Address, error) {
	return _TxState.Contract.EntryPoint(&_TxState.CallOpts)
}

// GetL1Txhash is a free data retrieval call binding the contract method 0x2867aa5f.
//
// Solidity: function getL1Txhash(address _from, uint64 _seqNum) view returns(bytes)
func (_TxState *TxStateCaller) GetL1Txhash(opts *bind.CallOpts, _from common.Address, _seqNum uint64) ([]byte, error) {
	var out []interface{}
	err := _TxState.contract.Call(opts, &out, "getL1Txhash", _from, _seqNum)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetL1Txhash is a free data retrieval call binding the contract method 0x2867aa5f.
//
// Solidity: function getL1Txhash(address _from, uint64 _seqNum) view returns(bytes)
func (_TxState *TxStateSession) GetL1Txhash(_from common.Address, _seqNum uint64) ([]byte, error) {
	return _TxState.Contract.GetL1Txhash(&_TxState.CallOpts, _from, _seqNum)
}

// GetL1Txhash is a free data retrieval call binding the contract method 0x2867aa5f.
//
// Solidity: function getL1Txhash(address _from, uint64 _seqNum) view returns(bytes)
func (_TxState *TxStateCallerSession) GetL1Txhash(_from common.Address, _seqNum uint64) ([]byte, error) {
	return _TxState.Contract.GetL1Txhash(&_TxState.CallOpts, _from, _seqNum)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TxState *TxStateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TxState.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TxState *TxStateSession) Owner() (common.Address, error) {
	return _TxState.Contract.Owner(&_TxState.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TxState *TxStateCallerSession) Owner() (common.Address, error) {
	return _TxState.Contract.Owner(&_TxState.CallOpts)
}

// ProposeTxToL1 is a paid mutator transaction binding the contract method 0x589200a3.
//
// Solidity: function proposeTxToL1(uint64 _chainId, address _from, uint64 _seqNum, address _receiver, uint256 _value, bytes _data) returns()
func (_TxState *TxStateTransactor) ProposeTxToL1(opts *bind.TransactOpts, _chainId uint64, _from common.Address, _seqNum uint64, _receiver common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _TxState.contract.Transact(opts, "proposeTxToL1", _chainId, _from, _seqNum, _receiver, _value, _data)
}

// ProposeTxToL1 is a paid mutator transaction binding the contract method 0x589200a3.
//
// Solidity: function proposeTxToL1(uint64 _chainId, address _from, uint64 _seqNum, address _receiver, uint256 _value, bytes _data) returns()
func (_TxState *TxStateSession) ProposeTxToL1(_chainId uint64, _from common.Address, _seqNum uint64, _receiver common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _TxState.Contract.ProposeTxToL1(&_TxState.TransactOpts, _chainId, _from, _seqNum, _receiver, _value, _data)
}

// ProposeTxToL1 is a paid mutator transaction binding the contract method 0x589200a3.
//
// Solidity: function proposeTxToL1(uint64 _chainId, address _from, uint64 _seqNum, address _receiver, uint256 _value, bytes _data) returns()
func (_TxState *TxStateTransactorSession) ProposeTxToL1(_chainId uint64, _from common.Address, _seqNum uint64, _receiver common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _TxState.Contract.ProposeTxToL1(&_TxState.TransactOpts, _chainId, _from, _seqNum, _receiver, _value, _data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TxState *TxStateTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxState.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TxState *TxStateSession) RenounceOwnership() (*types.Transaction, error) {
	return _TxState.Contract.RenounceOwnership(&_TxState.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TxState *TxStateTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TxState.Contract.RenounceOwnership(&_TxState.TransactOpts)
}

// SetL1TxState is a paid mutator transaction binding the contract method 0x23203762.
//
// Solidity: function setL1TxState(bytes _txHash, address l2Account, address _from, uint64 _seqNum, uint256 _state) payable returns()
func (_TxState *TxStateTransactor) SetL1TxState(opts *bind.TransactOpts, _txHash []byte, l2Account common.Address, _from common.Address, _seqNum uint64, _state *big.Int) (*types.Transaction, error) {
	return _TxState.contract.Transact(opts, "setL1TxState", _txHash, l2Account, _from, _seqNum, _state)
}

// SetL1TxState is a paid mutator transaction binding the contract method 0x23203762.
//
// Solidity: function setL1TxState(bytes _txHash, address l2Account, address _from, uint64 _seqNum, uint256 _state) payable returns()
func (_TxState *TxStateSession) SetL1TxState(_txHash []byte, l2Account common.Address, _from common.Address, _seqNum uint64, _state *big.Int) (*types.Transaction, error) {
	return _TxState.Contract.SetL1TxState(&_TxState.TransactOpts, _txHash, l2Account, _from, _seqNum, _state)
}

// SetL1TxState is a paid mutator transaction binding the contract method 0x23203762.
//
// Solidity: function setL1TxState(bytes _txHash, address l2Account, address _from, uint64 _seqNum, uint256 _state) payable returns()
func (_TxState *TxStateTransactorSession) SetL1TxState(_txHash []byte, l2Account common.Address, _from common.Address, _seqNum uint64, _state *big.Int) (*types.Transaction, error) {
	return _TxState.Contract.SetL1TxState(&_TxState.TransactOpts, _txHash, l2Account, _from, _seqNum, _state)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TxState *TxStateTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TxState.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TxState *TxStateSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TxState.Contract.TransferOwnership(&_TxState.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TxState *TxStateTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TxState.Contract.TransferOwnership(&_TxState.TransactOpts, newOwner)
}

// TxStateL1transferEventIterator is returned from FilterL1transferEvent and is used to iterate over the raw logs and unpacked data for L1transferEvent events raised by the TxState contract.
type TxStateL1transferEventIterator struct {
	Event *TxStateL1transferEvent // Event containing the contract specifics and raw log

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
func (it *TxStateL1transferEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TxStateL1transferEvent)
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
		it.Event = new(TxStateL1transferEvent)
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
func (it *TxStateL1transferEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TxStateL1transferEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TxStateL1transferEvent represents a L1transferEvent event raised by the TxState contract.
type TxStateL1transferEvent struct {
	Account common.Address
	From    common.Address
	SeqNum  uint64
	TxInfo  TxStateTransactionInfo
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterL1transferEvent is a free log retrieval operation binding the contract event 0x066387f39b13a8203bdbdc8377a678ce2489d9e4c1798f6b93cce980d7ab6502.
//
// Solidity: event L1transferEvent(address indexed account, address indexed from, uint64 indexed seqNum, (uint64,address,uint64,address,uint256,uint8,bytes) txInfo)
func (_TxState *TxStateFilterer) FilterL1transferEvent(opts *bind.FilterOpts, account []common.Address, from []common.Address, seqNum []uint64) (*TxStateL1transferEventIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var seqNumRule []interface{}
	for _, seqNumItem := range seqNum {
		seqNumRule = append(seqNumRule, seqNumItem)
	}

	logs, sub, err := _TxState.contract.FilterLogs(opts, "L1transferEvent", accountRule, fromRule, seqNumRule)
	if err != nil {
		return nil, err
	}
	return &TxStateL1transferEventIterator{contract: _TxState.contract, event: "L1transferEvent", logs: logs, sub: sub}, nil
}

// WatchL1transferEvent is a free log subscription operation binding the contract event 0x066387f39b13a8203bdbdc8377a678ce2489d9e4c1798f6b93cce980d7ab6502.
//
// Solidity: event L1transferEvent(address indexed account, address indexed from, uint64 indexed seqNum, (uint64,address,uint64,address,uint256,uint8,bytes) txInfo)
func (_TxState *TxStateFilterer) WatchL1transferEvent(opts *bind.WatchOpts, sink chan<- *TxStateL1transferEvent, account []common.Address, from []common.Address, seqNum []uint64) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var seqNumRule []interface{}
	for _, seqNumItem := range seqNum {
		seqNumRule = append(seqNumRule, seqNumItem)
	}

	logs, sub, err := _TxState.contract.WatchLogs(opts, "L1transferEvent", accountRule, fromRule, seqNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TxStateL1transferEvent)
				if err := _TxState.contract.UnpackLog(event, "L1transferEvent", log); err != nil {
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

// ParseL1transferEvent is a log parse operation binding the contract event 0x066387f39b13a8203bdbdc8377a678ce2489d9e4c1798f6b93cce980d7ab6502.
//
// Solidity: event L1transferEvent(address indexed account, address indexed from, uint64 indexed seqNum, (uint64,address,uint64,address,uint256,uint8,bytes) txInfo)
func (_TxState *TxStateFilterer) ParseL1transferEvent(log types.Log) (*TxStateL1transferEvent, error) {
	event := new(TxStateL1transferEvent)
	if err := _TxState.contract.UnpackLog(event, "L1transferEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TxStateOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TxState contract.
type TxStateOwnershipTransferredIterator struct {
	Event *TxStateOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TxStateOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TxStateOwnershipTransferred)
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
		it.Event = new(TxStateOwnershipTransferred)
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
func (it *TxStateOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TxStateOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TxStateOwnershipTransferred represents a OwnershipTransferred event raised by the TxState contract.
type TxStateOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TxState *TxStateFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TxStateOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TxState.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TxStateOwnershipTransferredIterator{contract: _TxState.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TxState *TxStateFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TxStateOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TxState.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TxStateOwnershipTransferred)
				if err := _TxState.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TxState *TxStateFilterer) ParseOwnershipTransferred(log types.Log) (*TxStateOwnershipTransferred, error) {
	event := new(TxStateOwnershipTransferred)
	if err := _TxState.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TxStateUpdateTxStateSuccessIterator is returned from FilterUpdateTxStateSuccess and is used to iterate over the raw logs and unpacked data for UpdateTxStateSuccess events raised by the TxState contract.
type TxStateUpdateTxStateSuccessIterator struct {
	Event *TxStateUpdateTxStateSuccess // Event containing the contract specifics and raw log

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
func (it *TxStateUpdateTxStateSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TxStateUpdateTxStateSuccess)
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
		it.Event = new(TxStateUpdateTxStateSuccess)
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
func (it *TxStateUpdateTxStateSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TxStateUpdateTxStateSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TxStateUpdateTxStateSuccess represents a UpdateTxStateSuccess event raised by the TxState contract.
type TxStateUpdateTxStateSuccess struct {
	L1Txhash common.Hash
	State    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateTxStateSuccess is a free log retrieval operation binding the contract event 0x8cfe54525125ba5a24f036d46ec84408796bc51fe7880e56db11467e5a9ca228.
//
// Solidity: event updateTxStateSuccess(bytes indexed L1Txhash, uint256 state)
func (_TxState *TxStateFilterer) FilterUpdateTxStateSuccess(opts *bind.FilterOpts, L1Txhash [][]byte) (*TxStateUpdateTxStateSuccessIterator, error) {

	var L1TxhashRule []interface{}
	for _, L1TxhashItem := range L1Txhash {
		L1TxhashRule = append(L1TxhashRule, L1TxhashItem)
	}

	logs, sub, err := _TxState.contract.FilterLogs(opts, "updateTxStateSuccess", L1TxhashRule)
	if err != nil {
		return nil, err
	}
	return &TxStateUpdateTxStateSuccessIterator{contract: _TxState.contract, event: "updateTxStateSuccess", logs: logs, sub: sub}, nil
}

// WatchUpdateTxStateSuccess is a free log subscription operation binding the contract event 0x8cfe54525125ba5a24f036d46ec84408796bc51fe7880e56db11467e5a9ca228.
//
// Solidity: event updateTxStateSuccess(bytes indexed L1Txhash, uint256 state)
func (_TxState *TxStateFilterer) WatchUpdateTxStateSuccess(opts *bind.WatchOpts, sink chan<- *TxStateUpdateTxStateSuccess, L1Txhash [][]byte) (event.Subscription, error) {

	var L1TxhashRule []interface{}
	for _, L1TxhashItem := range L1Txhash {
		L1TxhashRule = append(L1TxhashRule, L1TxhashItem)
	}

	logs, sub, err := _TxState.contract.WatchLogs(opts, "updateTxStateSuccess", L1TxhashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TxStateUpdateTxStateSuccess)
				if err := _TxState.contract.UnpackLog(event, "updateTxStateSuccess", log); err != nil {
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

// ParseUpdateTxStateSuccess is a log parse operation binding the contract event 0x8cfe54525125ba5a24f036d46ec84408796bc51fe7880e56db11467e5a9ca228.
//
// Solidity: event updateTxStateSuccess(bytes indexed L1Txhash, uint256 state)
func (_TxState *TxStateFilterer) ParseUpdateTxStateSuccess(log types.Log) (*TxStateUpdateTxStateSuccess, error) {
	event := new(TxStateUpdateTxStateSuccess)
	if err := _TxState.contract.UnpackLog(event, "updateTxStateSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

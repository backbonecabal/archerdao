// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package archerdao

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

// ArcherdaoABI is the input ABI used to generate the binding from.
const ArcherdaoABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"script\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"inputLocations\",\"type\":\"uint256[]\"}],\"name\":\"query\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"script\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"inputLocations\",\"type\":\"uint256[]\"}],\"name\":\"queryAllPrices\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Archerdao is an auto generated Go binding around an Ethereum contract.
type Archerdao struct {
	ArcherdaoCaller     // Read-only binding to the contract
	ArcherdaoTransactor // Write-only binding to the contract
	ArcherdaoFilterer   // Log filterer for contract events
}

// ArcherdaoCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArcherdaoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArcherdaoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArcherdaoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArcherdaoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArcherdaoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArcherdaoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArcherdaoSession struct {
	Contract     *Archerdao        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArcherdaoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArcherdaoCallerSession struct {
	Contract *ArcherdaoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ArcherdaoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArcherdaoTransactorSession struct {
	Contract     *ArcherdaoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ArcherdaoRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArcherdaoRaw struct {
	Contract *Archerdao // Generic contract binding to access the raw methods on
}

// ArcherdaoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArcherdaoCallerRaw struct {
	Contract *ArcherdaoCaller // Generic read-only contract binding to access the raw methods on
}

// ArcherdaoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArcherdaoTransactorRaw struct {
	Contract *ArcherdaoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArcherdao creates a new instance of Archerdao, bound to a specific deployed contract.
func NewArcherdao(address common.Address, backend bind.ContractBackend) (*Archerdao, error) {
	contract, err := bindArcherdao(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Archerdao{ArcherdaoCaller: ArcherdaoCaller{contract: contract}, ArcherdaoTransactor: ArcherdaoTransactor{contract: contract}, ArcherdaoFilterer: ArcherdaoFilterer{contract: contract}}, nil
}

// NewArcherdaoCaller creates a new read-only instance of Archerdao, bound to a specific deployed contract.
func NewArcherdaoCaller(address common.Address, caller bind.ContractCaller) (*ArcherdaoCaller, error) {
	contract, err := bindArcherdao(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArcherdaoCaller{contract: contract}, nil
}

// NewArcherdaoTransactor creates a new write-only instance of Archerdao, bound to a specific deployed contract.
func NewArcherdaoTransactor(address common.Address, transactor bind.ContractTransactor) (*ArcherdaoTransactor, error) {
	contract, err := bindArcherdao(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArcherdaoTransactor{contract: contract}, nil
}

// NewArcherdaoFilterer creates a new log filterer instance of Archerdao, bound to a specific deployed contract.
func NewArcherdaoFilterer(address common.Address, filterer bind.ContractFilterer) (*ArcherdaoFilterer, error) {
	contract, err := bindArcherdao(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArcherdaoFilterer{contract: contract}, nil
}

// bindArcherdao binds a generic wrapper to an already deployed contract.
func bindArcherdao(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArcherdaoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Archerdao *ArcherdaoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Archerdao.Contract.ArcherdaoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Archerdao *ArcherdaoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Archerdao.Contract.ArcherdaoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Archerdao *ArcherdaoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Archerdao.Contract.ArcherdaoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Archerdao *ArcherdaoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Archerdao.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Archerdao *ArcherdaoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Archerdao.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Archerdao *ArcherdaoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Archerdao.Contract.contract.Transact(opts, method, params...)
}

// GetPrice is a free data retrieval call binding the contract method 0x8ac06f50.
//
// Solidity: function getPrice(address contractAddress, bytes data) view returns(bytes)
func (_Archerdao *ArcherdaoCaller) GetPrice(opts *bind.CallOpts, contractAddress common.Address, data []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Archerdao.contract.Call(opts, out, "getPrice", contractAddress, data)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x8ac06f50.
//
// Solidity: function getPrice(address contractAddress, bytes data) view returns(bytes)
func (_Archerdao *ArcherdaoSession) GetPrice(contractAddress common.Address, data []byte) ([]byte, error) {
	return _Archerdao.Contract.GetPrice(&_Archerdao.CallOpts, contractAddress, data)
}

// GetPrice is a free data retrieval call binding the contract method 0x8ac06f50.
//
// Solidity: function getPrice(address contractAddress, bytes data) view returns(bytes)
func (_Archerdao *ArcherdaoCallerSession) GetPrice(contractAddress common.Address, data []byte) ([]byte, error) {
	return _Archerdao.Contract.GetPrice(&_Archerdao.CallOpts, contractAddress, data)
}

// Query is a free data retrieval call binding the contract method 0xe39e64de.
//
// Solidity: function query(bytes script, uint256[] inputLocations) view returns(uint256)
func (_Archerdao *ArcherdaoCaller) Query(opts *bind.CallOpts, script []byte, inputLocations []*big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Archerdao.contract.Call(opts, out, "query", script, inputLocations)
	return *ret0, err
}

// Query is a free data retrieval call binding the contract method 0xe39e64de.
//
// Solidity: function query(bytes script, uint256[] inputLocations) view returns(uint256)
func (_Archerdao *ArcherdaoSession) Query(script []byte, inputLocations []*big.Int) (*big.Int, error) {
	return _Archerdao.Contract.Query(&_Archerdao.CallOpts, script, inputLocations)
}

// Query is a free data retrieval call binding the contract method 0xe39e64de.
//
// Solidity: function query(bytes script, uint256[] inputLocations) view returns(uint256)
func (_Archerdao *ArcherdaoCallerSession) Query(script []byte, inputLocations []*big.Int) (*big.Int, error) {
	return _Archerdao.Contract.Query(&_Archerdao.CallOpts, script, inputLocations)
}

// QueryAllPrices is a free data retrieval call binding the contract method 0xa363a0f3.
//
// Solidity: function queryAllPrices(bytes script, uint256[] inputLocations) view returns(bytes)
func (_Archerdao *ArcherdaoCaller) QueryAllPrices(opts *bind.CallOpts, script []byte, inputLocations []*big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Archerdao.contract.Call(opts, out, "queryAllPrices", script, inputLocations)
	return *ret0, err
}

// QueryAllPrices is a free data retrieval call binding the contract method 0xa363a0f3.
//
// Solidity: function queryAllPrices(bytes script, uint256[] inputLocations) view returns(bytes)
func (_Archerdao *ArcherdaoSession) QueryAllPrices(script []byte, inputLocations []*big.Int) ([]byte, error) {
	return _Archerdao.Contract.QueryAllPrices(&_Archerdao.CallOpts, script, inputLocations)
}

// QueryAllPrices is a free data retrieval call binding the contract method 0xa363a0f3.
//
// Solidity: function queryAllPrices(bytes script, uint256[] inputLocations) view returns(bytes)
func (_Archerdao *ArcherdaoCallerSession) QueryAllPrices(script []byte, inputLocations []*big.Int) ([]byte, error) {
	return _Archerdao.Contract.QueryAllPrices(&_Archerdao.CallOpts, script, inputLocations)
}


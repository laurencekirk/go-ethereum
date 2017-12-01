// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AuthorisedMinersWhitelistABI is the input ABI used to generate the binding from.
const AuthorisedMinersWhitelistABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"isAuthorisedMiner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"removeMinersAuthorisation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"authoriseMiner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"AddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"RemovedFromWhitelist\",\"type\":\"event\"}]"

// AuthorisedMinersWhitelistBin is the compiled bytecode used for deploying new contracts.
const AuthorisedMinersWhitelistBin = `0x6060604052341561000f57600080fd5b6102a88061001e6000396000f3006060604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663949d225d8114610066578063a0cc5ebe14610092578063b9a7548b146100c5578063f045bebb146100e6575b600080fd5b341561007157600080fd5b610079610105565b60405163ffffffff909116815260200160405180910390f35b341561009d57600080fd5b6100b1600160a060020a0360043516610111565b604051901515815260200160405180910390f35b34156100d057600080fd5b6100e4600160a060020a036004351661012f565b005b34156100f157600080fd5b6100e4600160a060020a03600435166101d4565b60005463ffffffff1681565b600160a060020a031660009081526001602052604090205460ff1690565b600160a060020a03811660009081526001602052604090205460ff16151561015657600080fd5b600160a060020a038116600090815260016020526040808220805460ff19169055815463ffffffff19811663ffffffff91821660001901909116179091557fcdd2e9b91a56913d370075169cefa1602ba36be5301664f752192bb1709df75790829051600160a060020a03909116815260200160405180910390a150565b600160a060020a03811660009081526001602052604090205460ff16156101fa57600080fd5b600160a060020a0381166000908152600160208190526040808320805460ff191683179055825463ffffffff19811663ffffffff91821690930116919091179091557fa850ae9193f515cbae8d35e8925bd2be26627fc91bce650b8652ed254e9cab0390829051600160a060020a03909116815260200160405180910390a1505600a165627a7a72305820cb572b105465b4eaef6b4529eb5a5b21b983bf2d02b320829188dde7694ef0bc0029`

// DeployAuthorisedMinersWhitelist deploys a new Ethereum contract, binding an instance of AuthorisedMinersWhitelist to it.
func DeployAuthorisedMinersWhitelist(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AuthorisedMinersWhitelist, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorisedMinersWhitelistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuthorisedMinersWhitelistBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuthorisedMinersWhitelist{AuthorisedMinersWhitelistCaller: AuthorisedMinersWhitelistCaller{contract: contract}, AuthorisedMinersWhitelistTransactor: AuthorisedMinersWhitelistTransactor{contract: contract}}, nil
}

// AuthorisedMinersWhitelist is an auto generated Go binding around an Ethereum contract.
type AuthorisedMinersWhitelist struct {
	AuthorisedMinersWhitelistCaller     // Read-only binding to the contract
	AuthorisedMinersWhitelistTransactor // Write-only binding to the contract
}

// AuthorisedMinersWhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthorisedMinersWhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorisedMinersWhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthorisedMinersWhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorisedMinersWhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthorisedMinersWhitelistSession struct {
	Contract     *AuthorisedMinersWhitelist // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AuthorisedMinersWhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthorisedMinersWhitelistCallerSession struct {
	Contract *AuthorisedMinersWhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// AuthorisedMinersWhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthorisedMinersWhitelistTransactorSession struct {
	Contract     *AuthorisedMinersWhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// AuthorisedMinersWhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthorisedMinersWhitelistRaw struct {
	Contract *AuthorisedMinersWhitelist // Generic contract binding to access the raw methods on
}

// AuthorisedMinersWhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthorisedMinersWhitelistCallerRaw struct {
	Contract *AuthorisedMinersWhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// AuthorisedMinersWhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthorisedMinersWhitelistTransactorRaw struct {
	Contract *AuthorisedMinersWhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthorisedMinersWhitelist creates a new instance of AuthorisedMinersWhitelist, bound to a specific deployed contract.
func NewAuthorisedMinersWhitelist(address common.Address, backend bind.ContractBackend) (*AuthorisedMinersWhitelist, error) {
	contract, err := bindAuthorisedMinersWhitelist(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuthorisedMinersWhitelist{AuthorisedMinersWhitelistCaller: AuthorisedMinersWhitelistCaller{contract: contract}, AuthorisedMinersWhitelistTransactor: AuthorisedMinersWhitelistTransactor{contract: contract}}, nil
}

// NewAuthorisedMinersWhitelistCaller creates a new read-only instance of AuthorisedMinersWhitelist, bound to a specific deployed contract.
func NewAuthorisedMinersWhitelistCaller(address common.Address, caller bind.ContractCaller) (*AuthorisedMinersWhitelistCaller, error) {
	contract, err := bindAuthorisedMinersWhitelist(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorisedMinersWhitelistCaller{contract: contract}, nil
}

// NewAuthorisedMinersWhitelistTransactor creates a new write-only instance of AuthorisedMinersWhitelist, bound to a specific deployed contract.
func NewAuthorisedMinersWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthorisedMinersWhitelistTransactor, error) {
	contract, err := bindAuthorisedMinersWhitelist(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AuthorisedMinersWhitelistTransactor{contract: contract}, nil
}

// bindAuthorisedMinersWhitelist binds a generic wrapper to an already deployed contract.
func bindAuthorisedMinersWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorisedMinersWhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuthorisedMinersWhitelist.Contract.AuthorisedMinersWhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthorisedMinersWhitelist.Contract.AuthorisedMinersWhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthorisedMinersWhitelist.Contract.AuthorisedMinersWhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuthorisedMinersWhitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthorisedMinersWhitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthorisedMinersWhitelist.Contract.contract.Transact(opts, method, params...)
}

// IsAuthorisedMiner is a free data retrieval call binding the contract method 0xa0cc5ebe.
//
// Solidity: function isAuthorisedMiner(miner address) constant returns(bool)
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistCaller) IsAuthorisedMiner(opts *bind.CallOpts, miner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AuthorisedMinersWhitelist.contract.Call(opts, out, "isAuthorisedMiner", miner)
	return *ret0, err
}

// IsAuthorisedMiner is a free data retrieval call binding the contract method 0xa0cc5ebe.
//
// Solidity: function isAuthorisedMiner(miner address) constant returns(bool)
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistSession) IsAuthorisedMiner(miner common.Address) (bool, error) {
	return _AuthorisedMinersWhitelist.Contract.IsAuthorisedMiner(&_AuthorisedMinersWhitelist.CallOpts, miner)
}

// IsAuthorisedMiner is a free data retrieval call binding the contract method 0xa0cc5ebe.
//
// Solidity: function isAuthorisedMiner(miner address) constant returns(bool)
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistCallerSession) IsAuthorisedMiner(miner common.Address) (bool, error) {
	return _AuthorisedMinersWhitelist.Contract.IsAuthorisedMiner(&_AuthorisedMinersWhitelist.CallOpts, miner)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() constant returns(uint32)
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistCaller) Size(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _AuthorisedMinersWhitelist.contract.Call(opts, out, "size")
	return *ret0, err
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() constant returns(uint32)
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistSession) Size() (uint32, error) {
	return _AuthorisedMinersWhitelist.Contract.Size(&_AuthorisedMinersWhitelist.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() constant returns(uint32)
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistCallerSession) Size() (uint32, error) {
	return _AuthorisedMinersWhitelist.Contract.Size(&_AuthorisedMinersWhitelist.CallOpts)
}

// AuthoriseMiner is a paid mutator transaction binding the contract method 0xf045bebb.
//
// Solidity: function authoriseMiner(miner address) returns()
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistTransactor) AuthoriseMiner(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _AuthorisedMinersWhitelist.contract.Transact(opts, "authoriseMiner", miner)
}

// AuthoriseMiner is a paid mutator transaction binding the contract method 0xf045bebb.
//
// Solidity: function authoriseMiner(miner address) returns()
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistSession) AuthoriseMiner(miner common.Address) (*types.Transaction, error) {
	return _AuthorisedMinersWhitelist.Contract.AuthoriseMiner(&_AuthorisedMinersWhitelist.TransactOpts, miner)
}

// AuthoriseMiner is a paid mutator transaction binding the contract method 0xf045bebb.
//
// Solidity: function authoriseMiner(miner address) returns()
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistTransactorSession) AuthoriseMiner(miner common.Address) (*types.Transaction, error) {
	return _AuthorisedMinersWhitelist.Contract.AuthoriseMiner(&_AuthorisedMinersWhitelist.TransactOpts, miner)
}

// RemoveMinersAuthorisation is a paid mutator transaction binding the contract method 0xb9a7548b.
//
// Solidity: function removeMinersAuthorisation(miner address) returns()
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistTransactor) RemoveMinersAuthorisation(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _AuthorisedMinersWhitelist.contract.Transact(opts, "removeMinersAuthorisation", miner)
}

// RemoveMinersAuthorisation is a paid mutator transaction binding the contract method 0xb9a7548b.
//
// Solidity: function removeMinersAuthorisation(miner address) returns()
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistSession) RemoveMinersAuthorisation(miner common.Address) (*types.Transaction, error) {
	return _AuthorisedMinersWhitelist.Contract.RemoveMinersAuthorisation(&_AuthorisedMinersWhitelist.TransactOpts, miner)
}

// RemoveMinersAuthorisation is a paid mutator transaction binding the contract method 0xb9a7548b.
//
// Solidity: function removeMinersAuthorisation(miner address) returns()
func (_AuthorisedMinersWhitelist *AuthorisedMinersWhitelistTransactorSession) RemoveMinersAuthorisation(miner common.Address) (*types.Transaction, error) {
	return _AuthorisedMinersWhitelist.Contract.RemoveMinersAuthorisation(&_AuthorisedMinersWhitelist.TransactOpts, miner)
}

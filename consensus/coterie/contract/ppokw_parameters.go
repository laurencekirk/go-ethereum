// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PpokwParametersABI is the input ABI used to generate the binding from.
const PpokwParametersABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"name\":\"difficulty\",\"type\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCommitteeSize\",\"type\":\"uint32\"}],\"name\":\"setCommitteeSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newDifficulty\",\"type\":\"uint256\"}],\"name\":\"setDifficulty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"committeeSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sizeBefore\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"sizeAfter\",\"type\":\"uint32\"}],\"name\":\"CommitteeSizeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousDifficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newDifficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"DifficultyAdjusted\",\"type\":\"event\"}]"

// PpokwParametersBin is the compiled bytecode used for deploying new contracts.
const PpokwParametersBin = `0x6060604052341561000f57600080fd5b6102368061001e6000396000f3006060604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166319cae462811461006657806336890ec214610091578063602512e1146100af5780639cf4364b146100c5575b600080fd5b341561007157600080fd5b6100796100f1565b60405191825260208201526040908101905180910390f35b341561009c57600080fd5b6100ad63ffffffff600435166100fa565b005b34156100ba57600080fd5b6100ad600435610162565b34156100d057600080fd5b6100d86101fe565b60405163ffffffff909116815260200160405180910390f35b60015460025482565b6000547ff49b1bc7c3c86f8c510b5f0307a4f47c7553d95d5c4d7b0447f0889e1f9675409063ffffffff168260405163ffffffff9283168152911660208201526040908101905180910390a16000805463ffffffff191663ffffffff92909216919091179055565b806000811161017057600080fd5b7fc3045499c8ff4110de4029bd9ba27c1c54cef08823dfc8d37dc93c9051c46c6a44834333604051938452602084019290925260408084019190915273ffffffffffffffffffffffffffffffffffffffff90911660608301526080909101905180910390a1604080519081016040528281524360208201526001815181556020820151600190910155505050565b60005463ffffffff16815600a165627a7a72305820cd1b01bbb7a226fd02642aed0e5f73cca566c1846ea5dbbcd4b178bf635530b70029`

// DeployPpokwParameters deploys a new Ethereum contract, binding an instance of PpokwParameters to it.
func DeployPpokwParameters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PpokwParameters, error) {
	parsed, err := abi.JSON(strings.NewReader(PpokwParametersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PpokwParametersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PpokwParameters{PpokwParametersCaller: PpokwParametersCaller{contract: contract}, PpokwParametersTransactor: PpokwParametersTransactor{contract: contract}}, nil
}

// PpokwParameters is an auto generated Go binding around an Ethereum contract.
type PpokwParameters struct {
	PpokwParametersCaller     // Read-only binding to the contract
	PpokwParametersTransactor // Write-only binding to the contract
}

// PpokwParametersCaller is an auto generated read-only Go binding around an Ethereum contract.
type PpokwParametersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PpokwParametersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PpokwParametersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PpokwParametersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PpokwParametersSession struct {
	Contract     *PpokwParameters  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PpokwParametersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PpokwParametersCallerSession struct {
	Contract *PpokwParametersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PpokwParametersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PpokwParametersTransactorSession struct {
	Contract     *PpokwParametersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PpokwParametersRaw is an auto generated low-level Go binding around an Ethereum contract.
type PpokwParametersRaw struct {
	Contract *PpokwParameters // Generic contract binding to access the raw methods on
}

// PpokwParametersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PpokwParametersCallerRaw struct {
	Contract *PpokwParametersCaller // Generic read-only contract binding to access the raw methods on
}

// PpokwParametersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PpokwParametersTransactorRaw struct {
	Contract *PpokwParametersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPpokwParameters creates a new instance of PpokwParameters, bound to a specific deployed contract.
func NewPpokwParameters(address common.Address, backend bind.ContractBackend) (*PpokwParameters, error) {
	contract, err := bindPpokwParameters(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PpokwParameters{PpokwParametersCaller: PpokwParametersCaller{contract: contract}, PpokwParametersTransactor: PpokwParametersTransactor{contract: contract}}, nil
}

// NewPpokwParametersCaller creates a new read-only instance of PpokwParameters, bound to a specific deployed contract.
func NewPpokwParametersCaller(address common.Address, caller bind.ContractCaller) (*PpokwParametersCaller, error) {
	contract, err := bindPpokwParameters(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PpokwParametersCaller{contract: contract}, nil
}

// NewPpokwParametersTransactor creates a new write-only instance of PpokwParameters, bound to a specific deployed contract.
func NewPpokwParametersTransactor(address common.Address, transactor bind.ContractTransactor) (*PpokwParametersTransactor, error) {
	contract, err := bindPpokwParameters(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PpokwParametersTransactor{contract: contract}, nil
}

// bindPpokwParameters binds a generic wrapper to an already deployed contract.
func bindPpokwParameters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PpokwParametersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PpokwParameters *PpokwParametersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PpokwParameters.Contract.PpokwParametersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PpokwParameters *PpokwParametersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PpokwParameters.Contract.PpokwParametersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PpokwParameters *PpokwParametersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PpokwParameters.Contract.PpokwParametersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PpokwParameters *PpokwParametersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PpokwParameters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PpokwParameters *PpokwParametersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PpokwParameters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PpokwParameters *PpokwParametersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PpokwParameters.Contract.contract.Transact(opts, method, params...)
}

// CommitteeSize is a free data retrieval call binding the contract method 0x9cf4364b.
//
// Solidity: function committeeSize() constant returns(uint32)
func (_PpokwParameters *PpokwParametersCaller) CommitteeSize(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _PpokwParameters.contract.Call(opts, out, "committeeSize")
	return *ret0, err
}

// CommitteeSize is a free data retrieval call binding the contract method 0x9cf4364b.
//
// Solidity: function committeeSize() constant returns(uint32)
func (_PpokwParameters *PpokwParametersSession) CommitteeSize() (uint32, error) {
	return _PpokwParameters.Contract.CommitteeSize(&_PpokwParameters.CallOpts)
}

// CommitteeSize is a free data retrieval call binding the contract method 0x9cf4364b.
//
// Solidity: function committeeSize() constant returns(uint32)
func (_PpokwParameters *PpokwParametersCallerSession) CommitteeSize() (uint32, error) {
	return _PpokwParameters.Contract.CommitteeSize(&_PpokwParameters.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() constant returns(difficulty uint256, blockNumber uint256)
func (_PpokwParameters *PpokwParametersCaller) Difficulty(opts *bind.CallOpts) (struct {
	Difficulty  *big.Int
	BlockNumber *big.Int
}, error) {
	ret := new(struct {
		Difficulty  *big.Int
		BlockNumber *big.Int
	})
	out := ret
	err := _PpokwParameters.contract.Call(opts, out, "difficulty")
	return *ret, err
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() constant returns(difficulty uint256, blockNumber uint256)
func (_PpokwParameters *PpokwParametersSession) Difficulty() (struct {
	Difficulty  *big.Int
	BlockNumber *big.Int
}, error) {
	return _PpokwParameters.Contract.Difficulty(&_PpokwParameters.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() constant returns(difficulty uint256, blockNumber uint256)
func (_PpokwParameters *PpokwParametersCallerSession) Difficulty() (struct {
	Difficulty  *big.Int
	BlockNumber *big.Int
}, error) {
	return _PpokwParameters.Contract.Difficulty(&_PpokwParameters.CallOpts)
}

// SetCommitteeSize is a paid mutator transaction binding the contract method 0x36890ec2.
//
// Solidity: function setCommitteeSize(newCommitteeSize uint32) returns()
func (_PpokwParameters *PpokwParametersTransactor) SetCommitteeSize(opts *bind.TransactOpts, newCommitteeSize uint32) (*types.Transaction, error) {
	return _PpokwParameters.contract.Transact(opts, "setCommitteeSize", newCommitteeSize)
}

// SetCommitteeSize is a paid mutator transaction binding the contract method 0x36890ec2.
//
// Solidity: function setCommitteeSize(newCommitteeSize uint32) returns()
func (_PpokwParameters *PpokwParametersSession) SetCommitteeSize(newCommitteeSize uint32) (*types.Transaction, error) {
	return _PpokwParameters.Contract.SetCommitteeSize(&_PpokwParameters.TransactOpts, newCommitteeSize)
}

// SetCommitteeSize is a paid mutator transaction binding the contract method 0x36890ec2.
//
// Solidity: function setCommitteeSize(newCommitteeSize uint32) returns()
func (_PpokwParameters *PpokwParametersTransactorSession) SetCommitteeSize(newCommitteeSize uint32) (*types.Transaction, error) {
	return _PpokwParameters.Contract.SetCommitteeSize(&_PpokwParameters.TransactOpts, newCommitteeSize)
}

// SetDifficulty is a paid mutator transaction binding the contract method 0x602512e1.
//
// Solidity: function setDifficulty(newDifficulty uint256) returns()
func (_PpokwParameters *PpokwParametersTransactor) SetDifficulty(opts *bind.TransactOpts, newDifficulty *big.Int) (*types.Transaction, error) {
	return _PpokwParameters.contract.Transact(opts, "setDifficulty", newDifficulty)
}

// SetDifficulty is a paid mutator transaction binding the contract method 0x602512e1.
//
// Solidity: function setDifficulty(newDifficulty uint256) returns()
func (_PpokwParameters *PpokwParametersSession) SetDifficulty(newDifficulty *big.Int) (*types.Transaction, error) {
	return _PpokwParameters.Contract.SetDifficulty(&_PpokwParameters.TransactOpts, newDifficulty)
}

// SetDifficulty is a paid mutator transaction binding the contract method 0x602512e1.
//
// Solidity: function setDifficulty(newDifficulty uint256) returns()
func (_PpokwParameters *PpokwParametersTransactorSession) SetDifficulty(newDifficulty *big.Int) (*types.Transaction, error) {
	return _PpokwParameters.Contract.SetDifficulty(&_PpokwParameters.TransactOpts, newDifficulty)
}

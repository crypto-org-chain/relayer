// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relayer

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

// RelayerCallerMetaData contains all meta data concerning the RelayerCaller contract.
var RelayerCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[][]\",\"name\":\"batches\",\"type\":\"bytes[][]\"}],\"name\":\"batchCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506107078061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c8063839a7f0b1461002d575b5f80fd5b6100476004803603810190610042919061045a565b610049565b005b5f5b8151811015610149575f828281518110610068576100676104a1565b5b602002602001015160405160200161008091906105e9565b60405160208183030381529060405290505f606573ffffffffffffffffffffffffffffffffffffffff16826040516100b89190610643565b5f604051808303815f865af19150503d805f81146100f1576040519150601f19603f3d011682016040523d82523d5f602084013e6100f6565b606091505b505090508061013a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610131906106b3565b60405180910390fd5b5050808060010191505061004b565b5050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101a882610162565b810181811067ffffffffffffffff821117156101c7576101c6610172565b5b80604052505050565b5f6101d961014d565b90506101e5828261019f565b919050565b5f67ffffffffffffffff82111561020457610203610172565b5b602082029050602081019050919050565b5f80fd5b5f67ffffffffffffffff82111561023357610232610172565b5b602082029050602081019050919050565b5f80fd5b5f67ffffffffffffffff82111561026257610261610172565b5b61026b82610162565b9050602081019050919050565b828183375f83830152505050565b5f61029861029384610248565b6101d0565b9050828152602081018484840111156102b4576102b3610244565b5b6102bf848285610278565b509392505050565b5f82601f8301126102db576102da61015e565b5b81356102eb848260208601610286565b91505092915050565b5f61030661030184610219565b6101d0565b9050808382526020820190506020840283018581111561032957610328610215565b5b835b8181101561037057803567ffffffffffffffff81111561034e5761034d61015e565b5b80860161035b89826102c7565b8552602085019450505060208101905061032b565b5050509392505050565b5f82601f83011261038e5761038d61015e565b5b813561039e8482602086016102f4565b91505092915050565b5f6103b96103b4846101ea565b6101d0565b905080838252602082019050602084028301858111156103dc576103db610215565b5b835b8181101561042357803567ffffffffffffffff8111156104015761040061015e565b5b80860161040e898261037a565b855260208501945050506020810190506103de565b5050509392505050565b5f82601f8301126104415761044061015e565b5b81356104518482602086016103a7565b91505092915050565b5f6020828403121561046f5761046e610156565b5b5f82013567ffffffffffffffff81111561048c5761048b61015a565b5b6104988482850161042d565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610529826104f7565b6105338185610501565b9350610543818560208601610511565b61054c81610162565b840191505092915050565b5f610562838361051f565b905092915050565b5f602082019050919050565b5f610580826104ce565b61058a81856104d8565b93508360208202850161059c856104e8565b805f5b858110156105d757848403895281516105b88582610557565b94506105c38361056a565b925060208a0199505060018101905061059f565b50829750879550505050505092915050565b5f6020820190508181035f8301526106018184610576565b905092915050565b5f81905092915050565b5f61061d826104f7565b6106278185610609565b9350610637818560208601610511565b80840191505092915050565b5f61064e8284610613565b915081905092915050565b5f82825260208201905092915050565b7f52656c617965722063616c6c206661696c6564000000000000000000000000005f82015250565b5f61069d601383610659565b91506106a882610669565b602082019050919050565b5f6020820190508181035f8301526106ca81610691565b905091905056fea2646970667358221220c76950a4c8fd50c52db2cbc7313e97cebf2b0827fc28c1606d1fd7b86be5523864736f6c63430008190033",
}

// RelayerCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use RelayerCallerMetaData.ABI instead.
var RelayerCallerABI = RelayerCallerMetaData.ABI

// RelayerCallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RelayerCallerMetaData.Bin instead.
var RelayerCallerBin = RelayerCallerMetaData.Bin

// DeployRelayerCaller deploys a new Ethereum contract, binding an instance of RelayerCaller to it.
func DeployRelayerCaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RelayerCaller, error) {
	parsed, err := RelayerCallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RelayerCallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RelayerCaller{RelayerCallerCaller: RelayerCallerCaller{contract: contract}, RelayerCallerTransactor: RelayerCallerTransactor{contract: contract}, RelayerCallerFilterer: RelayerCallerFilterer{contract: contract}}, nil
}

// RelayerCaller is an auto generated Go binding around an Ethereum contract.
type RelayerCaller struct {
	RelayerCallerCaller     // Read-only binding to the contract
	RelayerCallerTransactor // Write-only binding to the contract
	RelayerCallerFilterer   // Log filterer for contract events
}

// RelayerCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayerCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayerCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayerCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayerCallerSession struct {
	Contract     *RelayerCaller    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayerCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayerCallerCallerSession struct {
	Contract *RelayerCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RelayerCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayerCallerTransactorSession struct {
	Contract     *RelayerCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RelayerCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayerCallerRaw struct {
	Contract *RelayerCaller // Generic contract binding to access the raw methods on
}

// RelayerCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayerCallerCallerRaw struct {
	Contract *RelayerCallerCaller // Generic read-only contract binding to access the raw methods on
}

// RelayerCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayerCallerTransactorRaw struct {
	Contract *RelayerCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayerCaller creates a new instance of RelayerCaller, bound to a specific deployed contract.
func NewRelayerCaller(address common.Address, backend bind.ContractBackend) (*RelayerCaller, error) {
	contract, err := bindRelayerCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayerCaller{RelayerCallerCaller: RelayerCallerCaller{contract: contract}, RelayerCallerTransactor: RelayerCallerTransactor{contract: contract}, RelayerCallerFilterer: RelayerCallerFilterer{contract: contract}}, nil
}

// NewRelayerCallerCaller creates a new read-only instance of RelayerCaller, bound to a specific deployed contract.
func NewRelayerCallerCaller(address common.Address, caller bind.ContractCaller) (*RelayerCallerCaller, error) {
	contract, err := bindRelayerCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerCallerCaller{contract: contract}, nil
}

// NewRelayerCallerTransactor creates a new write-only instance of RelayerCaller, bound to a specific deployed contract.
func NewRelayerCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayerCallerTransactor, error) {
	contract, err := bindRelayerCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerCallerTransactor{contract: contract}, nil
}

// NewRelayerCallerFilterer creates a new log filterer instance of RelayerCaller, bound to a specific deployed contract.
func NewRelayerCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayerCallerFilterer, error) {
	contract, err := bindRelayerCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayerCallerFilterer{contract: contract}, nil
}

// bindRelayerCaller binds a generic wrapper to an already deployed contract.
func bindRelayerCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RelayerCallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayerCaller *RelayerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayerCaller.Contract.RelayerCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayerCaller *RelayerCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayerCaller.Contract.RelayerCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayerCaller *RelayerCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayerCaller.Contract.RelayerCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayerCaller *RelayerCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayerCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayerCaller *RelayerCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayerCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayerCaller *RelayerCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayerCaller.Contract.contract.Transact(opts, method, params...)
}

// BatchCall is a paid mutator transaction binding the contract method 0x839a7f0b.
//
// Solidity: function batchCall(bytes[][] batches) returns()
func (_RelayerCaller *RelayerCallerTransactor) BatchCall(opts *bind.TransactOpts, batches [][][]byte) (*types.Transaction, error) {
	return _RelayerCaller.contract.Transact(opts, "batchCall", batches)
}

// BatchCall is a paid mutator transaction binding the contract method 0x839a7f0b.
//
// Solidity: function batchCall(bytes[][] batches) returns()
func (_RelayerCaller *RelayerCallerSession) BatchCall(batches [][][]byte) (*types.Transaction, error) {
	return _RelayerCaller.Contract.BatchCall(&_RelayerCaller.TransactOpts, batches)
}

// BatchCall is a paid mutator transaction binding the contract method 0x839a7f0b.
//
// Solidity: function batchCall(bytes[][] batches) returns()
func (_RelayerCaller *RelayerCallerTransactorSession) BatchCall(batches [][][]byte) (*types.Transaction, error) {
	return _RelayerCaller.Contract.BatchCall(&_RelayerCaller.TransactOpts, batches)
}

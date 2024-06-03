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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"}],\"name\":\"batchCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506105c98061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c806368be3cf21461002d575b5f80fd5b61004760048036038101906100429190610349565b610049565b005b5f8160405160200161005b91906104ab565b60405160208183030381529060405290505f606573ffffffffffffffffffffffffffffffffffffffff16826040516100939190610505565b5f604051808303815f865af19150503d805f81146100cc576040519150601f19603f3d011682016040523d82523d5f602084013e6100d1565b606091505b5050905080610115576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161010c90610575565b60405180910390fd5b505050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101758261012f565b810181811067ffffffffffffffff821117156101945761019361013f565b5b80604052505050565b5f6101a661011a565b90506101b2828261016c565b919050565b5f67ffffffffffffffff8211156101d1576101d061013f565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff8211156102045761020361013f565b5b61020d8261012f565b9050602081019050919050565b828183375f83830152505050565b5f61023a610235846101ea565b61019d565b905082815260208101848484011115610256576102556101e6565b5b61026184828561021a565b509392505050565b5f82601f83011261027d5761027c61012b565b5b813561028d848260208601610228565b91505092915050565b5f6102a86102a3846101b7565b61019d565b905080838252602082019050602084028301858111156102cb576102ca6101e2565b5b835b8181101561031257803567ffffffffffffffff8111156102f0576102ef61012b565b5b8086016102fd8982610269565b855260208501945050506020810190506102cd565b5050509392505050565b5f82601f8301126103305761032f61012b565b5b8135610340848260208601610296565b91505092915050565b5f6020828403121561035e5761035d610123565b5b5f82013567ffffffffffffffff81111561037b5761037a610127565b5b6103878482850161031c565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6103eb826103b9565b6103f581856103c3565b93506104058185602086016103d3565b61040e8161012f565b840191505092915050565b5f61042483836103e1565b905092915050565b5f602082019050919050565b5f61044282610390565b61044c818561039a565b93508360208202850161045e856103aa565b805f5b85811015610499578484038952815161047a8582610419565b94506104858361042c565b925060208a01995050600181019050610461565b50829750879550505050505092915050565b5f6020820190508181035f8301526104c38184610438565b905092915050565b5f81905092915050565b5f6104df826103b9565b6104e981856104cb565b93506104f98185602086016103d3565b80840191505092915050565b5f61051082846104d5565b915081905092915050565b5f82825260208201905092915050565b7f52656c617965722063616c6c206661696c6564000000000000000000000000005f82015250565b5f61055f60138361051b565b915061056a8261052b565b602082019050919050565b5f6020820190508181035f83015261058c81610553565b905091905056fea2646970667358221220cb286c4273269a9d1788ad788aa6485942bc6813c4c79b52fed4c52a1357361764736f6c63430008190033",
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

// BatchCall is a paid mutator transaction binding the contract method 0x68be3cf2.
//
// Solidity: function batchCall(bytes[] payloads) returns()
func (_RelayerCaller *RelayerCallerTransactor) BatchCall(opts *bind.TransactOpts, payloads [][]byte) (*types.Transaction, error) {
	return _RelayerCaller.contract.Transact(opts, "batchCall", payloads)
}

// BatchCall is a paid mutator transaction binding the contract method 0x68be3cf2.
//
// Solidity: function batchCall(bytes[] payloads) returns()
func (_RelayerCaller *RelayerCallerSession) BatchCall(payloads [][]byte) (*types.Transaction, error) {
	return _RelayerCaller.Contract.BatchCall(&_RelayerCaller.TransactOpts, payloads)
}

// BatchCall is a paid mutator transaction binding the contract method 0x68be3cf2.
//
// Solidity: function batchCall(bytes[] payloads) returns()
func (_RelayerCaller *RelayerCallerTransactorSession) BatchCall(payloads [][]byte) (*types.Transaction, error) {
	return _RelayerCaller.Contract.BatchCall(&_RelayerCaller.TransactOpts, payloads)
}

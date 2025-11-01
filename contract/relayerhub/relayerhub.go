// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relayerhub

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

// RelayerHubMetaData contains all meta data concerning the RelayerHub contract.
var RelayerHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardToRelayer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BUCKET_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BUCKET_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMERGENCY_OPERATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMERGENCY_UPGRADE_OPERATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC2771_FORWARDER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MOCASBT_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MOCASBT_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MOCAVC_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MOCAVC_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MOCA_EXECUTOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MOCA_EXECUTOR_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_MESSAGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_MESSAGE_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OBJECT_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OBJECT_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROXY_ADMIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_RATIO_SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100dd565b603254610100900460ff161561008a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60325460ff908116146100db576032805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b610b2c806100ec6000396000f3fe6080604052600436106102085760003560e01c806383d4433911610118578063b8e43b33116100a0578063d279c1911161006f578063d279c19114610486578063e02e86b0146104a6578063e613ae00146102cb578063eac78b33146104bb578063ed9bc82a146102cb57600080fd5b8063b8e43b33146102cb578063b9d86913146102cb578063cc12eabc14610471578063ccff6f7f146102cb57600080fd5b80639b904321116100e75780639b904321146102cb5780639feb8f5014610418578063a21d135414610438578063a9dae71c146102cb578063b287e0331461045c57600080fd5b806383d44339146103ae57806385661917146102cb5780638b30baf4146103db57806391e326001461040357600080fd5b8063618d569c1161019b5780637afffdd21161016a5780637afffdd2146102cb5780637ca1b81f1461036f5780637d2e3084146102cb5780638129fc1c1461038457806381d914801461039957600080fd5b8063618d569c1461034557806363bb23f4146102cb5780636d3358a1146102cb57806373f1e3c31461035a57600080fd5b806346934fc8116101d757806346934fc8146102cb57806346ab84e5146102cb578063557cf477146102cb578063572b6c05146102f857600080fd5b806304ea405414610267578063132f2adb146102935780632b67908d146102b6578063314c8402146102cb57600080fd5b366102625733156102605760405162461bcd60e51b815260206004820152601b60248201527f6f6e6c7920726563656976652066726f6d20746f6b656e20687562000000000060448201526064015b60405180910390fd5b005b600080fd5b34801561027357600080fd5b5061027c600a81565b60405160ff90911681526020015b60405180910390f35b34801561029f57600080fd5b506102a8606481565b60405190815260200161028a565b3480156102c257600080fd5b5061027c600781565b3480156102d757600080fd5b506102e0600081565b6040516001600160a01b03909116815260200161028a565b34801561030457600080fd5b50610335610313366004610979565b6001600160a01b031673db7d0bd38d223048b1cff39700e4c5238e346f7f1490565b604051901515815260200161028a565b34801561035157600080fd5b5061027c600281565b34801561036657600080fd5b5061027c600481565b34801561037b57600080fd5b5061027c600881565b34801561039057600080fd5b506102606104d0565b3480156103a557600080fd5b5061027c600381565b3480156103ba57600080fd5b506102a86103c9366004610979565b60656020526000908152604090205481565b3480156103e757600080fd5b506102e073db7d0bd38d223048b1cff39700e4c5238e346f7f81565b34801561040f57600080fd5b5061027c600981565b34801561042457600080fd5b5061026061043336600461099d565b6105e1565b34801561044457600080fd5b5061044d61070d565b60405161028a93929190610a0f565b34801561046857600080fd5b5061027c600b81565b34801561047d57600080fd5b5061027c600181565b34801561049257600080fd5b506102606104a1366004610979565b61075d565b3480156104b257600080fd5b5061027c600681565b3480156104c757600080fd5b5061027c600581565b603254610100900460ff16158080156104f05750603254600160ff909116105b8061050a5750303b15801561050a575060325460ff166001145b61056d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610257565b6032805460ff191660011790558015610590576032805461ff0019166101001790555b6105986108ac565b80156105de576032805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b331561062f5760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792043726f7373436861696e20636f6e747261637400000000000000006044820152606401610257565b604051639bfa764560e01b8152600481018290526000908190639bfa7645906024016020604051808303816000875af1158015610670573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106949190610a44565b6001600160a01b0384166000908152606560205260408120805492935083929091906106c1908490610a5d565b9091555050604080516001600160a01b0385168152602081018390527fcc3341048e8fd1ed288bcd99bd6231605849b6301fe5ae9170850a29d9b1c2dd910160405180910390a1505050565b60006060806207a1226040518060400160405280600a8152602001692932b630bcb2b9243ab160b11b815250604051806060016040528060278152602001610ad060279139925092509250909192565b6107656108dd565b6001600160a01b038116600090815260656020526040902054806107bf5760405162461bcd60e51b81526020600482015260116024820152701b9bc81c995b185e595c881c995dd85c99607a1b6044820152606401610257565b6001600160a01b038216600090815260656020526040812055478111156108285760405162461bcd60e51b815260206004820152601960248201527f72656c6179657220726577617264206e6f7420656e6f756768000000000000006044820152606401610257565b6040516001600160a01b0383169082156108fc029083906000818181858888f1935050505015801561085e573d6000803e3d6000fd5b50604080516001600160a01b0384168152602081018390527fd0813ff03c470dcc7baa9ce36914dc2febdfd276d639deffaac383fd3db42ba3910160405180910390a1506105de6001603355565b603254610100900460ff166108d35760405162461bcd60e51b815260040161025790610a84565b6108db61093d565b565b60026033540361092f5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610257565b6002603355565b6001603355565b603254610100900460ff166109365760405162461bcd60e51b815260040161025790610a84565b6001600160a01b03811681146105de57600080fd5b60006020828403121561098b57600080fd5b813561099681610964565b9392505050565b600080604083850312156109b057600080fd5b82356109bb81610964565b946020939093013593505050565b6000815180845260005b818110156109ef576020818501810151868301820152016109d3565b506000602082860101526020601f19601f83011685010191505092915050565b838152606060208201526000610a2860608301856109c9565b8281036040840152610a3a81856109c9565b9695505050505050565b600060208284031215610a5657600080fd5b5051919050565b80820180821115610a7e57634e487b7160e01b600052601160045260246000fd5b92915050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fe616464205f64697361626c65496e697469616c697a65727320696e20636f6e7374727563746f72a26469706673582212202d50c5930e1de8c64d68af52f5264b1cd8a878cead9d1e6433733cc1fbb273c264736f6c63430008110033",
}

// RelayerHubABI is the input ABI used to generate the binding from.
// Deprecated: Use RelayerHubMetaData.ABI instead.
var RelayerHubABI = RelayerHubMetaData.ABI

// RelayerHubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RelayerHubMetaData.Bin instead.
var RelayerHubBin = RelayerHubMetaData.Bin

// DeployRelayerHub deploys a new Ethereum contract, binding an instance of RelayerHub to it.
func DeployRelayerHub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RelayerHub, error) {
	parsed, err := RelayerHubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RelayerHubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RelayerHub{RelayerHubCaller: RelayerHubCaller{contract: contract}, RelayerHubTransactor: RelayerHubTransactor{contract: contract}, RelayerHubFilterer: RelayerHubFilterer{contract: contract}}, nil
}

// RelayerHub is an auto generated Go binding around an Ethereum contract.
type RelayerHub struct {
	RelayerHubCaller     // Read-only binding to the contract
	RelayerHubTransactor // Write-only binding to the contract
	RelayerHubFilterer   // Log filterer for contract events
}

// RelayerHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayerHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayerHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayerHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayerHubSession struct {
	Contract     *RelayerHub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayerHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayerHubCallerSession struct {
	Contract *RelayerHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RelayerHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayerHubTransactorSession struct {
	Contract     *RelayerHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RelayerHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayerHubRaw struct {
	Contract *RelayerHub // Generic contract binding to access the raw methods on
}

// RelayerHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayerHubCallerRaw struct {
	Contract *RelayerHubCaller // Generic read-only contract binding to access the raw methods on
}

// RelayerHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayerHubTransactorRaw struct {
	Contract *RelayerHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayerHub creates a new instance of RelayerHub, bound to a specific deployed contract.
func NewRelayerHub(address common.Address, backend bind.ContractBackend) (*RelayerHub, error) {
	contract, err := bindRelayerHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayerHub{RelayerHubCaller: RelayerHubCaller{contract: contract}, RelayerHubTransactor: RelayerHubTransactor{contract: contract}, RelayerHubFilterer: RelayerHubFilterer{contract: contract}}, nil
}

// NewRelayerHubCaller creates a new read-only instance of RelayerHub, bound to a specific deployed contract.
func NewRelayerHubCaller(address common.Address, caller bind.ContractCaller) (*RelayerHubCaller, error) {
	contract, err := bindRelayerHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerHubCaller{contract: contract}, nil
}

// NewRelayerHubTransactor creates a new write-only instance of RelayerHub, bound to a specific deployed contract.
func NewRelayerHubTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayerHubTransactor, error) {
	contract, err := bindRelayerHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerHubTransactor{contract: contract}, nil
}

// NewRelayerHubFilterer creates a new log filterer instance of RelayerHub, bound to a specific deployed contract.
func NewRelayerHubFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayerHubFilterer, error) {
	contract, err := bindRelayerHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayerHubFilterer{contract: contract}, nil
}

// bindRelayerHub binds a generic wrapper to an already deployed contract.
func bindRelayerHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RelayerHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayerHub *RelayerHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayerHub.Contract.RelayerHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayerHub *RelayerHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayerHub.Contract.RelayerHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayerHub *RelayerHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayerHub.Contract.RelayerHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayerHub *RelayerHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayerHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayerHub *RelayerHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayerHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayerHub *RelayerHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayerHub.Contract.contract.Transact(opts, method, params...)
}

// BUCKETCHANNELID is a free data retrieval call binding the contract method 0x73f1e3c3.
//
// Solidity: function BUCKET_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCaller) BUCKETCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "BUCKET_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BUCKETCHANNELID is a free data retrieval call binding the contract method 0x73f1e3c3.
//
// Solidity: function BUCKET_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubSession) BUCKETCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.BUCKETCHANNELID(&_RelayerHub.CallOpts)
}

// BUCKETCHANNELID is a free data retrieval call binding the contract method 0x73f1e3c3.
//
// Solidity: function BUCKET_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCallerSession) BUCKETCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.BUCKETCHANNELID(&_RelayerHub.CallOpts)
}

// BUCKETHUB is a free data retrieval call binding the contract method 0x7afffdd2.
//
// Solidity: function BUCKET_HUB() view returns(address)
func (_RelayerHub *RelayerHubCaller) BUCKETHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "BUCKET_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BUCKETHUB is a free data retrieval call binding the contract method 0x7afffdd2.
//
// Solidity: function BUCKET_HUB() view returns(address)
func (_RelayerHub *RelayerHubSession) BUCKETHUB() (common.Address, error) {
	return _RelayerHub.Contract.BUCKETHUB(&_RelayerHub.CallOpts)
}

// BUCKETHUB is a free data retrieval call binding the contract method 0x7afffdd2.
//
// Solidity: function BUCKET_HUB() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) BUCKETHUB() (common.Address, error) {
	return _RelayerHub.Contract.BUCKETHUB(&_RelayerHub.CallOpts)
}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_RelayerHub *RelayerHubCaller) CROSSCHAIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "CROSS_CHAIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_RelayerHub *RelayerHubSession) CROSSCHAIN() (common.Address, error) {
	return _RelayerHub.Contract.CROSSCHAIN(&_RelayerHub.CallOpts)
}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) CROSSCHAIN() (common.Address, error) {
	return _RelayerHub.Contract.CROSSCHAIN(&_RelayerHub.CallOpts)
}

// EMERGENCYOPERATOR is a free data retrieval call binding the contract method 0x46ab84e5.
//
// Solidity: function EMERGENCY_OPERATOR() view returns(address)
func (_RelayerHub *RelayerHubCaller) EMERGENCYOPERATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "EMERGENCY_OPERATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EMERGENCYOPERATOR is a free data retrieval call binding the contract method 0x46ab84e5.
//
// Solidity: function EMERGENCY_OPERATOR() view returns(address)
func (_RelayerHub *RelayerHubSession) EMERGENCYOPERATOR() (common.Address, error) {
	return _RelayerHub.Contract.EMERGENCYOPERATOR(&_RelayerHub.CallOpts)
}

// EMERGENCYOPERATOR is a free data retrieval call binding the contract method 0x46ab84e5.
//
// Solidity: function EMERGENCY_OPERATOR() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) EMERGENCYOPERATOR() (common.Address, error) {
	return _RelayerHub.Contract.EMERGENCYOPERATOR(&_RelayerHub.CallOpts)
}

// EMERGENCYUPGRADEOPERATOR is a free data retrieval call binding the contract method 0xb8e43b33.
//
// Solidity: function EMERGENCY_UPGRADE_OPERATOR() view returns(address)
func (_RelayerHub *RelayerHubCaller) EMERGENCYUPGRADEOPERATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "EMERGENCY_UPGRADE_OPERATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EMERGENCYUPGRADEOPERATOR is a free data retrieval call binding the contract method 0xb8e43b33.
//
// Solidity: function EMERGENCY_UPGRADE_OPERATOR() view returns(address)
func (_RelayerHub *RelayerHubSession) EMERGENCYUPGRADEOPERATOR() (common.Address, error) {
	return _RelayerHub.Contract.EMERGENCYUPGRADEOPERATOR(&_RelayerHub.CallOpts)
}

// EMERGENCYUPGRADEOPERATOR is a free data retrieval call binding the contract method 0xb8e43b33.
//
// Solidity: function EMERGENCY_UPGRADE_OPERATOR() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) EMERGENCYUPGRADEOPERATOR() (common.Address, error) {
	return _RelayerHub.Contract.EMERGENCYUPGRADEOPERATOR(&_RelayerHub.CallOpts)
}

// ERC2771FORWARDER is a free data retrieval call binding the contract method 0x8b30baf4.
//
// Solidity: function ERC2771_FORWARDER() view returns(address)
func (_RelayerHub *RelayerHubCaller) ERC2771FORWARDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "ERC2771_FORWARDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERC2771FORWARDER is a free data retrieval call binding the contract method 0x8b30baf4.
//
// Solidity: function ERC2771_FORWARDER() view returns(address)
func (_RelayerHub *RelayerHubSession) ERC2771FORWARDER() (common.Address, error) {
	return _RelayerHub.Contract.ERC2771FORWARDER(&_RelayerHub.CallOpts)
}

// ERC2771FORWARDER is a free data retrieval call binding the contract method 0x8b30baf4.
//
// Solidity: function ERC2771_FORWARDER() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) ERC2771FORWARDER() (common.Address, error) {
	return _RelayerHub.Contract.ERC2771FORWARDER(&_RelayerHub.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x81d91480.
//
// Solidity: function GOV_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "GOV_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x81d91480.
//
// Solidity: function GOV_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubSession) GOVCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.GOVCHANNELID(&_RelayerHub.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x81d91480.
//
// Solidity: function GOV_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCallerSession) GOVCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.GOVCHANNELID(&_RelayerHub.CallOpts)
}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_RelayerHub *RelayerHubCaller) GOVHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "GOV_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_RelayerHub *RelayerHubSession) GOVHUB() (common.Address, error) {
	return _RelayerHub.Contract.GOVHUB(&_RelayerHub.CallOpts)
}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) GOVHUB() (common.Address, error) {
	return _RelayerHub.Contract.GOVHUB(&_RelayerHub.CallOpts)
}

// GROUPCHANNELID is a free data retrieval call binding the contract method 0xe02e86b0.
//
// Solidity: function GROUP_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCaller) GROUPCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "GROUP_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GROUPCHANNELID is a free data retrieval call binding the contract method 0xe02e86b0.
//
// Solidity: function GROUP_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubSession) GROUPCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.GROUPCHANNELID(&_RelayerHub.CallOpts)
}

// GROUPCHANNELID is a free data retrieval call binding the contract method 0xe02e86b0.
//
// Solidity: function GROUP_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCallerSession) GROUPCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.GROUPCHANNELID(&_RelayerHub.CallOpts)
}

// GROUPHUB is a free data retrieval call binding the contract method 0x46934fc8.
//
// Solidity: function GROUP_HUB() view returns(address)
func (_RelayerHub *RelayerHubCaller) GROUPHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "GROUP_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GROUPHUB is a free data retrieval call binding the contract method 0x46934fc8.
//
// Solidity: function GROUP_HUB() view returns(address)
func (_RelayerHub *RelayerHubSession) GROUPHUB() (common.Address, error) {
	return _RelayerHub.Contract.GROUPHUB(&_RelayerHub.CallOpts)
}

// GROUPHUB is a free data retrieval call binding the contract method 0x46934fc8.
//
// Solidity: function GROUP_HUB() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) GROUPHUB() (common.Address, error) {
	return _RelayerHub.Contract.GROUPHUB(&_RelayerHub.CallOpts)
}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_RelayerHub *RelayerHubCaller) LIGHTCLIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "LIGHT_CLIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_RelayerHub *RelayerHubSession) LIGHTCLIENT() (common.Address, error) {
	return _RelayerHub.Contract.LIGHTCLIENT(&_RelayerHub.CallOpts)
}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) LIGHTCLIENT() (common.Address, error) {
	return _RelayerHub.Contract.LIGHTCLIENT(&_RelayerHub.CallOpts)
}

// MOCASBTCHANNELID is a free data retrieval call binding the contract method 0x04ea4054.
//
// Solidity: function MOCASBT_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCaller) MOCASBTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "MOCASBT_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MOCASBTCHANNELID is a free data retrieval call binding the contract method 0x04ea4054.
//
// Solidity: function MOCASBT_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubSession) MOCASBTCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.MOCASBTCHANNELID(&_RelayerHub.CallOpts)
}

// MOCASBTCHANNELID is a free data retrieval call binding the contract method 0x04ea4054.
//
// Solidity: function MOCASBT_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCallerSession) MOCASBTCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.MOCASBTCHANNELID(&_RelayerHub.CallOpts)
}

// MOCASBTHUB is a free data retrieval call binding the contract method 0x9b904321.
//
// Solidity: function MOCASBT_HUB() view returns(address)
func (_RelayerHub *RelayerHubCaller) MOCASBTHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "MOCASBT_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MOCASBTHUB is a free data retrieval call binding the contract method 0x9b904321.
//
// Solidity: function MOCASBT_HUB() view returns(address)
func (_RelayerHub *RelayerHubSession) MOCASBTHUB() (common.Address, error) {
	return _RelayerHub.Contract.MOCASBTHUB(&_RelayerHub.CallOpts)
}

// MOCASBTHUB is a free data retrieval call binding the contract method 0x9b904321.
//
// Solidity: function MOCASBT_HUB() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) MOCASBTHUB() (common.Address, error) {
	return _RelayerHub.Contract.MOCASBTHUB(&_RelayerHub.CallOpts)
}

// MOCAVCCHANNELID is a free data retrieval call binding the contract method 0xb287e033.
//
// Solidity: function MOCAVC_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCaller) MOCAVCCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "MOCAVC_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MOCAVCCHANNELID is a free data retrieval call binding the contract method 0xb287e033.
//
// Solidity: function MOCAVC_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubSession) MOCAVCCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.MOCAVCCHANNELID(&_RelayerHub.CallOpts)
}

// MOCAVCCHANNELID is a free data retrieval call binding the contract method 0xb287e033.
//
// Solidity: function MOCAVC_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCallerSession) MOCAVCCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.MOCAVCCHANNELID(&_RelayerHub.CallOpts)
}

// MOCAVCHUB is a free data retrieval call binding the contract method 0x63bb23f4.
//
// Solidity: function MOCAVC_HUB() view returns(address)
func (_RelayerHub *RelayerHubCaller) MOCAVCHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "MOCAVC_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MOCAVCHUB is a free data retrieval call binding the contract method 0x63bb23f4.
//
// Solidity: function MOCAVC_HUB() view returns(address)
func (_RelayerHub *RelayerHubSession) MOCAVCHUB() (common.Address, error) {
	return _RelayerHub.Contract.MOCAVCHUB(&_RelayerHub.CallOpts)
}

// MOCAVCHUB is a free data retrieval call binding the contract method 0x63bb23f4.
//
// Solidity: function MOCAVC_HUB() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) MOCAVCHUB() (common.Address, error) {
	return _RelayerHub.Contract.MOCAVCHUB(&_RelayerHub.CallOpts)
}

// MOCAEXECUTOR is a free data retrieval call binding the contract method 0x314c8402.
//
// Solidity: function MOCA_EXECUTOR() view returns(address)
func (_RelayerHub *RelayerHubCaller) MOCAEXECUTOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "MOCA_EXECUTOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MOCAEXECUTOR is a free data retrieval call binding the contract method 0x314c8402.
//
// Solidity: function MOCA_EXECUTOR() view returns(address)
func (_RelayerHub *RelayerHubSession) MOCAEXECUTOR() (common.Address, error) {
	return _RelayerHub.Contract.MOCAEXECUTOR(&_RelayerHub.CallOpts)
}

// MOCAEXECUTOR is a free data retrieval call binding the contract method 0x314c8402.
//
// Solidity: function MOCA_EXECUTOR() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) MOCAEXECUTOR() (common.Address, error) {
	return _RelayerHub.Contract.MOCAEXECUTOR(&_RelayerHub.CallOpts)
}

// MOCAEXECUTORCHANNELID is a free data retrieval call binding the contract method 0x91e32600.
//
// Solidity: function MOCA_EXECUTOR_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCaller) MOCAEXECUTORCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "MOCA_EXECUTOR_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MOCAEXECUTORCHANNELID is a free data retrieval call binding the contract method 0x91e32600.
//
// Solidity: function MOCA_EXECUTOR_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubSession) MOCAEXECUTORCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.MOCAEXECUTORCHANNELID(&_RelayerHub.CallOpts)
}

// MOCAEXECUTORCHANNELID is a free data retrieval call binding the contract method 0x91e32600.
//
// Solidity: function MOCA_EXECUTOR_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCallerSession) MOCAEXECUTORCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.MOCAEXECUTORCHANNELID(&_RelayerHub.CallOpts)
}

// MULTIMESSAGE is a free data retrieval call binding the contract method 0xccff6f7f.
//
// Solidity: function MULTI_MESSAGE() view returns(address)
func (_RelayerHub *RelayerHubCaller) MULTIMESSAGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "MULTI_MESSAGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MULTIMESSAGE is a free data retrieval call binding the contract method 0xccff6f7f.
//
// Solidity: function MULTI_MESSAGE() view returns(address)
func (_RelayerHub *RelayerHubSession) MULTIMESSAGE() (common.Address, error) {
	return _RelayerHub.Contract.MULTIMESSAGE(&_RelayerHub.CallOpts)
}

// MULTIMESSAGE is a free data retrieval call binding the contract method 0xccff6f7f.
//
// Solidity: function MULTI_MESSAGE() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) MULTIMESSAGE() (common.Address, error) {
	return _RelayerHub.Contract.MULTIMESSAGE(&_RelayerHub.CallOpts)
}

// MULTIMESSAGECHANNELID is a free data retrieval call binding the contract method 0x7ca1b81f.
//
// Solidity: function MULTI_MESSAGE_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCaller) MULTIMESSAGECHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "MULTI_MESSAGE_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MULTIMESSAGECHANNELID is a free data retrieval call binding the contract method 0x7ca1b81f.
//
// Solidity: function MULTI_MESSAGE_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubSession) MULTIMESSAGECHANNELID() (uint8, error) {
	return _RelayerHub.Contract.MULTIMESSAGECHANNELID(&_RelayerHub.CallOpts)
}

// MULTIMESSAGECHANNELID is a free data retrieval call binding the contract method 0x7ca1b81f.
//
// Solidity: function MULTI_MESSAGE_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCallerSession) MULTIMESSAGECHANNELID() (uint8, error) {
	return _RelayerHub.Contract.MULTIMESSAGECHANNELID(&_RelayerHub.CallOpts)
}

// OBJECTCHANNELID is a free data retrieval call binding the contract method 0xeac78b33.
//
// Solidity: function OBJECT_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCaller) OBJECTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "OBJECT_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OBJECTCHANNELID is a free data retrieval call binding the contract method 0xeac78b33.
//
// Solidity: function OBJECT_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubSession) OBJECTCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.OBJECTCHANNELID(&_RelayerHub.CallOpts)
}

// OBJECTCHANNELID is a free data retrieval call binding the contract method 0xeac78b33.
//
// Solidity: function OBJECT_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCallerSession) OBJECTCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.OBJECTCHANNELID(&_RelayerHub.CallOpts)
}

// OBJECTHUB is a free data retrieval call binding the contract method 0x7d2e3084.
//
// Solidity: function OBJECT_HUB() view returns(address)
func (_RelayerHub *RelayerHubCaller) OBJECTHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "OBJECT_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OBJECTHUB is a free data retrieval call binding the contract method 0x7d2e3084.
//
// Solidity: function OBJECT_HUB() view returns(address)
func (_RelayerHub *RelayerHubSession) OBJECTHUB() (common.Address, error) {
	return _RelayerHub.Contract.OBJECTHUB(&_RelayerHub.CallOpts)
}

// OBJECTHUB is a free data retrieval call binding the contract method 0x7d2e3084.
//
// Solidity: function OBJECT_HUB() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) OBJECTHUB() (common.Address, error) {
	return _RelayerHub.Contract.OBJECTHUB(&_RelayerHub.CallOpts)
}

// PERMISSIONCHANNELID is a free data retrieval call binding the contract method 0x2b67908d.
//
// Solidity: function PERMISSION_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCaller) PERMISSIONCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "PERMISSION_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PERMISSIONCHANNELID is a free data retrieval call binding the contract method 0x2b67908d.
//
// Solidity: function PERMISSION_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubSession) PERMISSIONCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.PERMISSIONCHANNELID(&_RelayerHub.CallOpts)
}

// PERMISSIONCHANNELID is a free data retrieval call binding the contract method 0x2b67908d.
//
// Solidity: function PERMISSION_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCallerSession) PERMISSIONCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.PERMISSIONCHANNELID(&_RelayerHub.CallOpts)
}

// PERMISSIONHUB is a free data retrieval call binding the contract method 0x85661917.
//
// Solidity: function PERMISSION_HUB() view returns(address)
func (_RelayerHub *RelayerHubCaller) PERMISSIONHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "PERMISSION_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PERMISSIONHUB is a free data retrieval call binding the contract method 0x85661917.
//
// Solidity: function PERMISSION_HUB() view returns(address)
func (_RelayerHub *RelayerHubSession) PERMISSIONHUB() (common.Address, error) {
	return _RelayerHub.Contract.PERMISSIONHUB(&_RelayerHub.CallOpts)
}

// PERMISSIONHUB is a free data retrieval call binding the contract method 0x85661917.
//
// Solidity: function PERMISSION_HUB() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) PERMISSIONHUB() (common.Address, error) {
	return _RelayerHub.Contract.PERMISSIONHUB(&_RelayerHub.CallOpts)
}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_RelayerHub *RelayerHubCaller) PROXYADMIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "PROXY_ADMIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_RelayerHub *RelayerHubSession) PROXYADMIN() (common.Address, error) {
	return _RelayerHub.Contract.PROXYADMIN(&_RelayerHub.CallOpts)
}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) PROXYADMIN() (common.Address, error) {
	return _RelayerHub.Contract.PROXYADMIN(&_RelayerHub.CallOpts)
}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_RelayerHub *RelayerHubCaller) RELAYERHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "RELAYER_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_RelayerHub *RelayerHubSession) RELAYERHUB() (common.Address, error) {
	return _RelayerHub.Contract.RELAYERHUB(&_RelayerHub.CallOpts)
}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) RELAYERHUB() (common.Address, error) {
	return _RelayerHub.Contract.RELAYERHUB(&_RelayerHub.CallOpts)
}

// REWARDRATIOSCALE is a free data retrieval call binding the contract method 0x132f2adb.
//
// Solidity: function REWARD_RATIO_SCALE() view returns(uint256)
func (_RelayerHub *RelayerHubCaller) REWARDRATIOSCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "REWARD_RATIO_SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDRATIOSCALE is a free data retrieval call binding the contract method 0x132f2adb.
//
// Solidity: function REWARD_RATIO_SCALE() view returns(uint256)
func (_RelayerHub *RelayerHubSession) REWARDRATIOSCALE() (*big.Int, error) {
	return _RelayerHub.Contract.REWARDRATIOSCALE(&_RelayerHub.CallOpts)
}

// REWARDRATIOSCALE is a free data retrieval call binding the contract method 0x132f2adb.
//
// Solidity: function REWARD_RATIO_SCALE() view returns(uint256)
func (_RelayerHub *RelayerHubCallerSession) REWARDRATIOSCALE() (*big.Int, error) {
	return _RelayerHub.Contract.REWARDRATIOSCALE(&_RelayerHub.CallOpts)
}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_RelayerHub *RelayerHubCaller) TOKENHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "TOKEN_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_RelayerHub *RelayerHubSession) TOKENHUB() (common.Address, error) {
	return _RelayerHub.Contract.TOKENHUB(&_RelayerHub.CallOpts)
}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_RelayerHub *RelayerHubCallerSession) TOKENHUB() (common.Address, error) {
	return _RelayerHub.Contract.TOKENHUB(&_RelayerHub.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0xcc12eabc.
//
// Solidity: function TRANSFER_IN_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "TRANSFER_IN_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0xcc12eabc.
//
// Solidity: function TRANSFER_IN_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubSession) TRANSFERINCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.TRANSFERINCHANNELID(&_RelayerHub.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0xcc12eabc.
//
// Solidity: function TRANSFER_IN_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.TRANSFERINCHANNELID(&_RelayerHub.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0x618d569c.
//
// Solidity: function TRANSFER_OUT_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "TRANSFER_OUT_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0x618d569c.
//
// Solidity: function TRANSFER_OUT_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.TRANSFEROUTCHANNELID(&_RelayerHub.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0x618d569c.
//
// Solidity: function TRANSFER_OUT_CHANNEL_ID() view returns(uint8)
func (_RelayerHub *RelayerHubCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _RelayerHub.Contract.TRANSFEROUTCHANNELID(&_RelayerHub.CallOpts)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) pure returns(bool)
func (_RelayerHub *RelayerHubCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) pure returns(bool)
func (_RelayerHub *RelayerHubSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _RelayerHub.Contract.IsTrustedForwarder(&_RelayerHub.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) pure returns(bool)
func (_RelayerHub *RelayerHubCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _RelayerHub.Contract.IsTrustedForwarder(&_RelayerHub.CallOpts, forwarder)
}

// RewardMap is a free data retrieval call binding the contract method 0x83d44339.
//
// Solidity: function rewardMap(address ) view returns(uint256)
func (_RelayerHub *RelayerHubCaller) RewardMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "rewardMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardMap is a free data retrieval call binding the contract method 0x83d44339.
//
// Solidity: function rewardMap(address ) view returns(uint256)
func (_RelayerHub *RelayerHubSession) RewardMap(arg0 common.Address) (*big.Int, error) {
	return _RelayerHub.Contract.RewardMap(&_RelayerHub.CallOpts, arg0)
}

// RewardMap is a free data retrieval call binding the contract method 0x83d44339.
//
// Solidity: function rewardMap(address ) view returns(uint256)
func (_RelayerHub *RelayerHubCallerSession) RewardMap(arg0 common.Address) (*big.Int, error) {
	return _RelayerHub.Contract.RewardMap(&_RelayerHub.CallOpts, arg0)
}

// VersionInfo is a free data retrieval call binding the contract method 0xa21d1354.
//
// Solidity: function versionInfo() pure returns(uint256 version, string name, string description)
func (_RelayerHub *RelayerHubCaller) VersionInfo(opts *bind.CallOpts) (struct {
	Version     *big.Int
	Name        string
	Description string
}, error) {
	var out []interface{}
	err := _RelayerHub.contract.Call(opts, &out, "versionInfo")

	outstruct := new(struct {
		Version     *big.Int
		Name        string
		Description string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Version = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// VersionInfo is a free data retrieval call binding the contract method 0xa21d1354.
//
// Solidity: function versionInfo() pure returns(uint256 version, string name, string description)
func (_RelayerHub *RelayerHubSession) VersionInfo() (struct {
	Version     *big.Int
	Name        string
	Description string
}, error) {
	return _RelayerHub.Contract.VersionInfo(&_RelayerHub.CallOpts)
}

// VersionInfo is a free data retrieval call binding the contract method 0xa21d1354.
//
// Solidity: function versionInfo() pure returns(uint256 version, string name, string description)
func (_RelayerHub *RelayerHubCallerSession) VersionInfo() (struct {
	Version     *big.Int
	Name        string
	Description string
}, error) {
	return _RelayerHub.Contract.VersionInfo(&_RelayerHub.CallOpts)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _relayer, uint256 _reward) returns()
func (_RelayerHub *RelayerHubTransactor) AddReward(opts *bind.TransactOpts, _relayer common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _RelayerHub.contract.Transact(opts, "addReward", _relayer, _reward)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _relayer, uint256 _reward) returns()
func (_RelayerHub *RelayerHubSession) AddReward(_relayer common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _RelayerHub.Contract.AddReward(&_RelayerHub.TransactOpts, _relayer, _reward)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _relayer, uint256 _reward) returns()
func (_RelayerHub *RelayerHubTransactorSession) AddReward(_relayer common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _RelayerHub.Contract.AddReward(&_RelayerHub.TransactOpts, _relayer, _reward)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _relayer) returns()
func (_RelayerHub *RelayerHubTransactor) ClaimReward(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _RelayerHub.contract.Transact(opts, "claimReward", _relayer)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _relayer) returns()
func (_RelayerHub *RelayerHubSession) ClaimReward(_relayer common.Address) (*types.Transaction, error) {
	return _RelayerHub.Contract.ClaimReward(&_RelayerHub.TransactOpts, _relayer)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _relayer) returns()
func (_RelayerHub *RelayerHubTransactorSession) ClaimReward(_relayer common.Address) (*types.Transaction, error) {
	return _RelayerHub.Contract.ClaimReward(&_RelayerHub.TransactOpts, _relayer)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RelayerHub *RelayerHubTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayerHub.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RelayerHub *RelayerHubSession) Initialize() (*types.Transaction, error) {
	return _RelayerHub.Contract.Initialize(&_RelayerHub.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RelayerHub *RelayerHubTransactorSession) Initialize() (*types.Transaction, error) {
	return _RelayerHub.Contract.Initialize(&_RelayerHub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RelayerHub *RelayerHubTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayerHub.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RelayerHub *RelayerHubSession) Receive() (*types.Transaction, error) {
	return _RelayerHub.Contract.Receive(&_RelayerHub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RelayerHub *RelayerHubTransactorSession) Receive() (*types.Transaction, error) {
	return _RelayerHub.Contract.Receive(&_RelayerHub.TransactOpts)
}

// RelayerHubClaimedRewardIterator is returned from FilterClaimedReward and is used to iterate over the raw logs and unpacked data for ClaimedReward events raised by the RelayerHub contract.
type RelayerHubClaimedRewardIterator struct {
	Event *RelayerHubClaimedReward // Event containing the contract specifics and raw log

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
func (it *RelayerHubClaimedRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerHubClaimedReward)
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
		it.Event = new(RelayerHubClaimedReward)
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
func (it *RelayerHubClaimedRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerHubClaimedRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerHubClaimedReward represents a ClaimedReward event raised by the RelayerHub contract.
type RelayerHubClaimedReward struct {
	Relayer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedReward is a free log retrieval operation binding the contract event 0xd0813ff03c470dcc7baa9ce36914dc2febdfd276d639deffaac383fd3db42ba3.
//
// Solidity: event ClaimedReward(address relayer, uint256 amount)
func (_RelayerHub *RelayerHubFilterer) FilterClaimedReward(opts *bind.FilterOpts) (*RelayerHubClaimedRewardIterator, error) {

	logs, sub, err := _RelayerHub.contract.FilterLogs(opts, "ClaimedReward")
	if err != nil {
		return nil, err
	}
	return &RelayerHubClaimedRewardIterator{contract: _RelayerHub.contract, event: "ClaimedReward", logs: logs, sub: sub}, nil
}

// WatchClaimedReward is a free log subscription operation binding the contract event 0xd0813ff03c470dcc7baa9ce36914dc2febdfd276d639deffaac383fd3db42ba3.
//
// Solidity: event ClaimedReward(address relayer, uint256 amount)
func (_RelayerHub *RelayerHubFilterer) WatchClaimedReward(opts *bind.WatchOpts, sink chan<- *RelayerHubClaimedReward) (event.Subscription, error) {

	logs, sub, err := _RelayerHub.contract.WatchLogs(opts, "ClaimedReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerHubClaimedReward)
				if err := _RelayerHub.contract.UnpackLog(event, "ClaimedReward", log); err != nil {
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

// ParseClaimedReward is a log parse operation binding the contract event 0xd0813ff03c470dcc7baa9ce36914dc2febdfd276d639deffaac383fd3db42ba3.
//
// Solidity: event ClaimedReward(address relayer, uint256 amount)
func (_RelayerHub *RelayerHubFilterer) ParseClaimedReward(log types.Log) (*RelayerHubClaimedReward, error) {
	event := new(RelayerHubClaimedReward)
	if err := _RelayerHub.contract.UnpackLog(event, "ClaimedReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayerHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RelayerHub contract.
type RelayerHubInitializedIterator struct {
	Event *RelayerHubInitialized // Event containing the contract specifics and raw log

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
func (it *RelayerHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerHubInitialized)
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
		it.Event = new(RelayerHubInitialized)
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
func (it *RelayerHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerHubInitialized represents a Initialized event raised by the RelayerHub contract.
type RelayerHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RelayerHub *RelayerHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*RelayerHubInitializedIterator, error) {

	logs, sub, err := _RelayerHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RelayerHubInitializedIterator{contract: _RelayerHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RelayerHub *RelayerHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RelayerHubInitialized) (event.Subscription, error) {

	logs, sub, err := _RelayerHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerHubInitialized)
				if err := _RelayerHub.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RelayerHub *RelayerHubFilterer) ParseInitialized(log types.Log) (*RelayerHubInitialized, error) {
	event := new(RelayerHubInitialized)
	if err := _RelayerHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayerHubRewardToRelayerIterator is returned from FilterRewardToRelayer and is used to iterate over the raw logs and unpacked data for RewardToRelayer events raised by the RelayerHub contract.
type RelayerHubRewardToRelayerIterator struct {
	Event *RelayerHubRewardToRelayer // Event containing the contract specifics and raw log

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
func (it *RelayerHubRewardToRelayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerHubRewardToRelayer)
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
		it.Event = new(RelayerHubRewardToRelayer)
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
func (it *RelayerHubRewardToRelayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerHubRewardToRelayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerHubRewardToRelayer represents a RewardToRelayer event raised by the RelayerHub contract.
type RelayerHubRewardToRelayer struct {
	Relayer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardToRelayer is a free log retrieval operation binding the contract event 0xcc3341048e8fd1ed288bcd99bd6231605849b6301fe5ae9170850a29d9b1c2dd.
//
// Solidity: event RewardToRelayer(address relayer, uint256 amount)
func (_RelayerHub *RelayerHubFilterer) FilterRewardToRelayer(opts *bind.FilterOpts) (*RelayerHubRewardToRelayerIterator, error) {

	logs, sub, err := _RelayerHub.contract.FilterLogs(opts, "RewardToRelayer")
	if err != nil {
		return nil, err
	}
	return &RelayerHubRewardToRelayerIterator{contract: _RelayerHub.contract, event: "RewardToRelayer", logs: logs, sub: sub}, nil
}

// WatchRewardToRelayer is a free log subscription operation binding the contract event 0xcc3341048e8fd1ed288bcd99bd6231605849b6301fe5ae9170850a29d9b1c2dd.
//
// Solidity: event RewardToRelayer(address relayer, uint256 amount)
func (_RelayerHub *RelayerHubFilterer) WatchRewardToRelayer(opts *bind.WatchOpts, sink chan<- *RelayerHubRewardToRelayer) (event.Subscription, error) {

	logs, sub, err := _RelayerHub.contract.WatchLogs(opts, "RewardToRelayer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerHubRewardToRelayer)
				if err := _RelayerHub.contract.UnpackLog(event, "RewardToRelayer", log); err != nil {
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

// ParseRewardToRelayer is a log parse operation binding the contract event 0xcc3341048e8fd1ed288bcd99bd6231605849b6301fe5ae9170850a29d9b1c2dd.
//
// Solidity: event RewardToRelayer(address relayer, uint256 amount)
func (_RelayerHub *RelayerHubFilterer) ParseRewardToRelayer(log types.Log) (*RelayerHubRewardToRelayer, error) {
	event := new(RelayerHubRewardToRelayer)
	if err := _RelayerHub.contract.UnpackLog(event, "RewardToRelayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

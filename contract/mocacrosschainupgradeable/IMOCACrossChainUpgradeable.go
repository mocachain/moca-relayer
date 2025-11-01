// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocacrosschainupgradeable

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

// KYCDataLibEventData is an auto generated low-level Go binding around an user-defined struct.
type KYCDataLibEventData struct {
	SrcChainId  uint32
	DestChainId uint32
	ChannelId   uint32
	Sequence    *big.Int
	Payload     []byte
}

// IMOCACrossChainUpgradeableMetaData contains all meta data concerning the IMOCACrossChainUpgradeable contract.
var IMOCACrossChainUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcChainId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"channelId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"MocaSBTCrossChainPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_GRANT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TYPES_MIRROR_FAILED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TYPES_MIRROR_PENDING\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TYPES_MIRROR_SUCCEED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TYPES_TOKEN_CREATED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ackMinted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"srcUser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destUser\",\"type\":\"address\"}],\"name\":\"forward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"}],\"name\":\"getCrossChainSequence\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getCrossChainStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sbt_contract_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"outMintCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IMOCACrossChainUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IMOCACrossChainUpgradeableMetaData.ABI instead.
var IMOCACrossChainUpgradeableABI = IMOCACrossChainUpgradeableMetaData.ABI

// IMOCACrossChainUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IMOCACrossChainUpgradeableMetaData.Bin instead.
var IMOCACrossChainUpgradeableBin = IMOCACrossChainUpgradeableMetaData.Bin

// DeployIMOCACrossChainUpgradeable deploys a new Ethereum contract, binding an instance of IMOCACrossChainUpgradeable to it.
func DeployIMOCACrossChainUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IMOCACrossChainUpgradeable, error) {
	parsed, err := IMOCACrossChainUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IMOCACrossChainUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IMOCACrossChainUpgradeable{IMOCACrossChainUpgradeableCaller: IMOCACrossChainUpgradeableCaller{contract: contract}, IMOCACrossChainUpgradeableTransactor: IMOCACrossChainUpgradeableTransactor{contract: contract}, IMOCACrossChainUpgradeableFilterer: IMOCACrossChainUpgradeableFilterer{contract: contract}}, nil
}

// IMOCACrossChainUpgradeable is an auto generated Go binding around an Ethereum contract.
type IMOCACrossChainUpgradeable struct {
	IMOCACrossChainUpgradeableCaller     // Read-only binding to the contract
	IMOCACrossChainUpgradeableTransactor // Write-only binding to the contract
	IMOCACrossChainUpgradeableFilterer   // Log filterer for contract events
}

// IMOCACrossChainUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMOCACrossChainUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMOCACrossChainUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMOCACrossChainUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMOCACrossChainUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMOCACrossChainUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMOCACrossChainUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMOCACrossChainUpgradeableSession struct {
	Contract     *IMOCACrossChainUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IMOCACrossChainUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMOCACrossChainUpgradeableCallerSession struct {
	Contract *IMOCACrossChainUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// IMOCACrossChainUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMOCACrossChainUpgradeableTransactorSession struct {
	Contract     *IMOCACrossChainUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// IMOCACrossChainUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMOCACrossChainUpgradeableRaw struct {
	Contract *IMOCACrossChainUpgradeable // Generic contract binding to access the raw methods on
}

// IMOCACrossChainUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMOCACrossChainUpgradeableCallerRaw struct {
	Contract *IMOCACrossChainUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IMOCACrossChainUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMOCACrossChainUpgradeableTransactorRaw struct {
	Contract *IMOCACrossChainUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMOCACrossChainUpgradeable creates a new instance of IMOCACrossChainUpgradeable, bound to a specific deployed contract.
func NewIMOCACrossChainUpgradeable(address common.Address, backend bind.ContractBackend) (*IMOCACrossChainUpgradeable, error) {
	contract, err := bindIMOCACrossChainUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMOCACrossChainUpgradeable{IMOCACrossChainUpgradeableCaller: IMOCACrossChainUpgradeableCaller{contract: contract}, IMOCACrossChainUpgradeableTransactor: IMOCACrossChainUpgradeableTransactor{contract: contract}, IMOCACrossChainUpgradeableFilterer: IMOCACrossChainUpgradeableFilterer{contract: contract}}, nil
}

// NewIMOCACrossChainUpgradeableCaller creates a new read-only instance of IMOCACrossChainUpgradeable, bound to a specific deployed contract.
func NewIMOCACrossChainUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IMOCACrossChainUpgradeableCaller, error) {
	contract, err := bindIMOCACrossChainUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMOCACrossChainUpgradeableCaller{contract: contract}, nil
}

// NewIMOCACrossChainUpgradeableTransactor creates a new write-only instance of IMOCACrossChainUpgradeable, bound to a specific deployed contract.
func NewIMOCACrossChainUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IMOCACrossChainUpgradeableTransactor, error) {
	contract, err := bindIMOCACrossChainUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMOCACrossChainUpgradeableTransactor{contract: contract}, nil
}

// NewIMOCACrossChainUpgradeableFilterer creates a new log filterer instance of IMOCACrossChainUpgradeable, bound to a specific deployed contract.
func NewIMOCACrossChainUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IMOCACrossChainUpgradeableFilterer, error) {
	contract, err := bindIMOCACrossChainUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMOCACrossChainUpgradeableFilterer{contract: contract}, nil
}

// bindIMOCACrossChainUpgradeable binds a generic wrapper to an already deployed contract.
func bindIMOCACrossChainUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMOCACrossChainUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMOCACrossChainUpgradeable.Contract.IMOCACrossChainUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.IMOCACrossChainUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.IMOCACrossChainUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMOCACrossChainUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IMOCACrossChainUpgradeable.Contract.DEFAULTADMINROLE(&_IMOCACrossChainUpgradeable.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IMOCACrossChainUpgradeable.Contract.DEFAULTADMINROLE(&_IMOCACrossChainUpgradeable.CallOpts)
}

// OPERATORGRANT is a free data retrieval call binding the contract method 0x02c045d9.
//
// Solidity: function OPERATOR_GRANT() view returns(uint256)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) OPERATORGRANT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "OPERATOR_GRANT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPERATORGRANT is a free data retrieval call binding the contract method 0x02c045d9.
//
// Solidity: function OPERATOR_GRANT() view returns(uint256)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) OPERATORGRANT() (*big.Int, error) {
	return _IMOCACrossChainUpgradeable.Contract.OPERATORGRANT(&_IMOCACrossChainUpgradeable.CallOpts)
}

// OPERATORGRANT is a free data retrieval call binding the contract method 0x02c045d9.
//
// Solidity: function OPERATOR_GRANT() view returns(uint256)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) OPERATORGRANT() (*big.Int, error) {
	return _IMOCACrossChainUpgradeable.Contract.OPERATORGRANT(&_IMOCACrossChainUpgradeable.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) OPERATORROLE() ([32]byte, error) {
	return _IMOCACrossChainUpgradeable.Contract.OPERATORROLE(&_IMOCACrossChainUpgradeable.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) OPERATORROLE() ([32]byte, error) {
	return _IMOCACrossChainUpgradeable.Contract.OPERATORROLE(&_IMOCACrossChainUpgradeable.CallOpts)
}

// TYPESMIRRORFAILED is a free data retrieval call binding the contract method 0x733d50c7.
//
// Solidity: function TYPES_MIRROR_FAILED() view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) TYPESMIRRORFAILED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "TYPES_MIRROR_FAILED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPESMIRRORFAILED is a free data retrieval call binding the contract method 0x733d50c7.
//
// Solidity: function TYPES_MIRROR_FAILED() view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) TYPESMIRRORFAILED() (uint8, error) {
	return _IMOCACrossChainUpgradeable.Contract.TYPESMIRRORFAILED(&_IMOCACrossChainUpgradeable.CallOpts)
}

// TYPESMIRRORFAILED is a free data retrieval call binding the contract method 0x733d50c7.
//
// Solidity: function TYPES_MIRROR_FAILED() view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) TYPESMIRRORFAILED() (uint8, error) {
	return _IMOCACrossChainUpgradeable.Contract.TYPESMIRRORFAILED(&_IMOCACrossChainUpgradeable.CallOpts)
}

// TYPESMIRRORPENDING is a free data retrieval call binding the contract method 0x63d44136.
//
// Solidity: function TYPES_MIRROR_PENDING() view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) TYPESMIRRORPENDING(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "TYPES_MIRROR_PENDING")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPESMIRRORPENDING is a free data retrieval call binding the contract method 0x63d44136.
//
// Solidity: function TYPES_MIRROR_PENDING() view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) TYPESMIRRORPENDING() (uint8, error) {
	return _IMOCACrossChainUpgradeable.Contract.TYPESMIRRORPENDING(&_IMOCACrossChainUpgradeable.CallOpts)
}

// TYPESMIRRORPENDING is a free data retrieval call binding the contract method 0x63d44136.
//
// Solidity: function TYPES_MIRROR_PENDING() view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) TYPESMIRRORPENDING() (uint8, error) {
	return _IMOCACrossChainUpgradeable.Contract.TYPESMIRRORPENDING(&_IMOCACrossChainUpgradeable.CallOpts)
}

// TYPESMIRRORSUCCEED is a free data retrieval call binding the contract method 0x8ad13e49.
//
// Solidity: function TYPES_MIRROR_SUCCEED() view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) TYPESMIRRORSUCCEED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "TYPES_MIRROR_SUCCEED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPESMIRRORSUCCEED is a free data retrieval call binding the contract method 0x8ad13e49.
//
// Solidity: function TYPES_MIRROR_SUCCEED() view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) TYPESMIRRORSUCCEED() (uint8, error) {
	return _IMOCACrossChainUpgradeable.Contract.TYPESMIRRORSUCCEED(&_IMOCACrossChainUpgradeable.CallOpts)
}

// TYPESMIRRORSUCCEED is a free data retrieval call binding the contract method 0x8ad13e49.
//
// Solidity: function TYPES_MIRROR_SUCCEED() view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) TYPESMIRRORSUCCEED() (uint8, error) {
	return _IMOCACrossChainUpgradeable.Contract.TYPESMIRRORSUCCEED(&_IMOCACrossChainUpgradeable.CallOpts)
}

// TYPESTOKENCREATED is a free data retrieval call binding the contract method 0x3f90803c.
//
// Solidity: function TYPES_TOKEN_CREATED() view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) TYPESTOKENCREATED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "TYPES_TOKEN_CREATED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPESTOKENCREATED is a free data retrieval call binding the contract method 0x3f90803c.
//
// Solidity: function TYPES_TOKEN_CREATED() view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) TYPESTOKENCREATED() (uint8, error) {
	return _IMOCACrossChainUpgradeable.Contract.TYPESTOKENCREATED(&_IMOCACrossChainUpgradeable.CallOpts)
}

// TYPESTOKENCREATED is a free data retrieval call binding the contract method 0x3f90803c.
//
// Solidity: function TYPES_TOKEN_CREATED() view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) TYPESTOKENCREATED() (uint8, error) {
	return _IMOCACrossChainUpgradeable.Contract.TYPESTOKENCREATED(&_IMOCACrossChainUpgradeable.CallOpts)
}

// GetCrossChainSequence is a free data retrieval call binding the contract method 0xd8b2a23b.
//
// Solidity: function getCrossChainSequence(uint32 chainId) view returns(uint256)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) GetCrossChainSequence(opts *bind.CallOpts, chainId uint32) (*big.Int, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "getCrossChainSequence", chainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCrossChainSequence is a free data retrieval call binding the contract method 0xd8b2a23b.
//
// Solidity: function getCrossChainSequence(uint32 chainId) view returns(uint256)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) GetCrossChainSequence(chainId uint32) (*big.Int, error) {
	return _IMOCACrossChainUpgradeable.Contract.GetCrossChainSequence(&_IMOCACrossChainUpgradeable.CallOpts, chainId)
}

// GetCrossChainSequence is a free data retrieval call binding the contract method 0xd8b2a23b.
//
// Solidity: function getCrossChainSequence(uint32 chainId) view returns(uint256)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) GetCrossChainSequence(chainId uint32) (*big.Int, error) {
	return _IMOCACrossChainUpgradeable.Contract.GetCrossChainSequence(&_IMOCACrossChainUpgradeable.CallOpts, chainId)
}

// GetCrossChainStatus is a free data retrieval call binding the contract method 0xfbdafb72.
//
// Solidity: function getCrossChainStatus(uint32 chainId, address user) view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) GetCrossChainStatus(opts *bind.CallOpts, chainId uint32, user common.Address) (uint8, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "getCrossChainStatus", chainId, user)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCrossChainStatus is a free data retrieval call binding the contract method 0xfbdafb72.
//
// Solidity: function getCrossChainStatus(uint32 chainId, address user) view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) GetCrossChainStatus(chainId uint32, user common.Address) (uint8, error) {
	return _IMOCACrossChainUpgradeable.Contract.GetCrossChainStatus(&_IMOCACrossChainUpgradeable.CallOpts, chainId, user)
}

// GetCrossChainStatus is a free data retrieval call binding the contract method 0xfbdafb72.
//
// Solidity: function getCrossChainStatus(uint32 chainId, address user) view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) GetCrossChainStatus(chainId uint32, user common.Address) (uint8, error) {
	return _IMOCACrossChainUpgradeable.Contract.GetCrossChainStatus(&_IMOCACrossChainUpgradeable.CallOpts, chainId, user)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IMOCACrossChainUpgradeable.Contract.GetRoleAdmin(&_IMOCACrossChainUpgradeable.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IMOCACrossChainUpgradeable.Contract.GetRoleAdmin(&_IMOCACrossChainUpgradeable.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IMOCACrossChainUpgradeable.Contract.HasRole(&_IMOCACrossChainUpgradeable.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IMOCACrossChainUpgradeable.Contract.HasRole(&_IMOCACrossChainUpgradeable.CallOpts, role, account)
}

// OutMintCount is a free data retrieval call binding the contract method 0x079ccd2a.
//
// Solidity: function outMintCount(uint32 chainId, address user, uint8 status) view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) OutMintCount(opts *bind.CallOpts, chainId uint32, user common.Address, status uint8) (uint8, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "outMintCount", chainId, user, status)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OutMintCount is a free data retrieval call binding the contract method 0x079ccd2a.
//
// Solidity: function outMintCount(uint32 chainId, address user, uint8 status) view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) OutMintCount(chainId uint32, user common.Address, status uint8) (uint8, error) {
	return _IMOCACrossChainUpgradeable.Contract.OutMintCount(&_IMOCACrossChainUpgradeable.CallOpts, chainId, user, status)
}

// OutMintCount is a free data retrieval call binding the contract method 0x079ccd2a.
//
// Solidity: function outMintCount(uint32 chainId, address user, uint8 status) view returns(uint8)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) OutMintCount(chainId uint32, user common.Address, status uint8) (uint8, error) {
	return _IMOCACrossChainUpgradeable.Contract.OutMintCount(&_IMOCACrossChainUpgradeable.CallOpts, chainId, user, status)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IMOCACrossChainUpgradeable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IMOCACrossChainUpgradeable.Contract.SupportsInterface(&_IMOCACrossChainUpgradeable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IMOCACrossChainUpgradeable.Contract.SupportsInterface(&_IMOCACrossChainUpgradeable.CallOpts, interfaceId)
}

// AckMinted is a paid mutator transaction binding the contract method 0xa7d7f87c.
//
// Solidity: function ackMinted(uint32 chainId, address user, uint8 status) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactor) AckMinted(opts *bind.TransactOpts, chainId uint32, user common.Address, status uint8) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.contract.Transact(opts, "ackMinted", chainId, user, status)
}

// AckMinted is a paid mutator transaction binding the contract method 0xa7d7f87c.
//
// Solidity: function ackMinted(uint32 chainId, address user, uint8 status) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) AckMinted(chainId uint32, user common.Address, status uint8) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.AckMinted(&_IMOCACrossChainUpgradeable.TransactOpts, chainId, user, status)
}

// AckMinted is a paid mutator transaction binding the contract method 0xa7d7f87c.
//
// Solidity: function ackMinted(uint32 chainId, address user, uint8 status) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactorSession) AckMinted(chainId uint32, user common.Address, status uint8) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.AckMinted(&_IMOCACrossChainUpgradeable.TransactOpts, chainId, user, status)
}

// Forward is a paid mutator transaction binding the contract method 0x04b2c8df.
//
// Solidity: function forward(uint32 chainId, address srcUser, address destUser) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactor) Forward(opts *bind.TransactOpts, chainId uint32, srcUser common.Address, destUser common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.contract.Transact(opts, "forward", chainId, srcUser, destUser)
}

// Forward is a paid mutator transaction binding the contract method 0x04b2c8df.
//
// Solidity: function forward(uint32 chainId, address srcUser, address destUser) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) Forward(chainId uint32, srcUser common.Address, destUser common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.Forward(&_IMOCACrossChainUpgradeable.TransactOpts, chainId, srcUser, destUser)
}

// Forward is a paid mutator transaction binding the contract method 0x04b2c8df.
//
// Solidity: function forward(uint32 chainId, address srcUser, address destUser) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactorSession) Forward(chainId uint32, srcUser common.Address, destUser common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.Forward(&_IMOCACrossChainUpgradeable.TransactOpts, chainId, srcUser, destUser)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.GrantRole(&_IMOCACrossChainUpgradeable.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.GrantRole(&_IMOCACrossChainUpgradeable.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin_, address sbt_contract_) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactor) Initialize(opts *bind.TransactOpts, admin_ common.Address, sbt_contract_ common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.contract.Transact(opts, "initialize", admin_, sbt_contract_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin_, address sbt_contract_) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) Initialize(admin_ common.Address, sbt_contract_ common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.Initialize(&_IMOCACrossChainUpgradeable.TransactOpts, admin_, sbt_contract_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin_, address sbt_contract_) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactorSession) Initialize(admin_ common.Address, sbt_contract_ common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.Initialize(&_IMOCACrossChainUpgradeable.TransactOpts, admin_, sbt_contract_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.RenounceRole(&_IMOCACrossChainUpgradeable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.RenounceRole(&_IMOCACrossChainUpgradeable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.RevokeRole(&_IMOCACrossChainUpgradeable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IMOCACrossChainUpgradeable.Contract.RevokeRole(&_IMOCACrossChainUpgradeable.TransactOpts, role, account)
}

// IMOCACrossChainUpgradeableMocaSBTCrossChainPackageIterator is returned from FilterMocaSBTCrossChainPackage and is used to iterate over the raw logs and unpacked data for MocaSBTCrossChainPackage events raised by the IMOCACrossChainUpgradeable contract.
type IMOCACrossChainUpgradeableMocaSBTCrossChainPackageIterator struct {
	Event *IMOCACrossChainUpgradeableMocaSBTCrossChainPackage // Event containing the contract specifics and raw log

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
func (it *IMOCACrossChainUpgradeableMocaSBTCrossChainPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMOCACrossChainUpgradeableMocaSBTCrossChainPackage)
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
		it.Event = new(IMOCACrossChainUpgradeableMocaSBTCrossChainPackage)
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
func (it *IMOCACrossChainUpgradeableMocaSBTCrossChainPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMOCACrossChainUpgradeableMocaSBTCrossChainPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMOCACrossChainUpgradeableMocaSBTCrossChainPackage represents a MocaSBTCrossChainPackage event raised by the IMOCACrossChainUpgradeable contract.
type IMOCACrossChainUpgradeableMocaSBTCrossChainPackage struct {
	SrcChainId  uint32
	DestChainId uint32
	ChannelId   uint32
	Sequence    *big.Int
	Payload     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMocaSBTCrossChainPackage is a free log retrieval operation binding the contract event 0xeae7aa948aa4486965776005e20135ce32c4e9a9bd3704ec53d108056bcba038.
//
// Solidity: event MocaSBTCrossChainPackage(uint32 srcChainId, uint32 indexed destChainId, uint32 indexed channelId, uint256 indexed sequence, bytes payload)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableFilterer) FilterMocaSBTCrossChainPackage(opts *bind.FilterOpts, destChainId []uint32, channelId []uint32, sequence []*big.Int) (*IMOCACrossChainUpgradeableMocaSBTCrossChainPackageIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _IMOCACrossChainUpgradeable.contract.FilterLogs(opts, "MocaSBTCrossChainPackage", destChainIdRule, channelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &IMOCACrossChainUpgradeableMocaSBTCrossChainPackageIterator{contract: _IMOCACrossChainUpgradeable.contract, event: "MocaSBTCrossChainPackage", logs: logs, sub: sub}, nil
}

// WatchMocaSBTCrossChainPackage is a free log subscription operation binding the contract event 0xeae7aa948aa4486965776005e20135ce32c4e9a9bd3704ec53d108056bcba038.
//
// Solidity: event MocaSBTCrossChainPackage(uint32 srcChainId, uint32 indexed destChainId, uint32 indexed channelId, uint256 indexed sequence, bytes payload)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableFilterer) WatchMocaSBTCrossChainPackage(opts *bind.WatchOpts, sink chan<- *IMOCACrossChainUpgradeableMocaSBTCrossChainPackage, destChainId []uint32, channelId []uint32, sequence []*big.Int) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _IMOCACrossChainUpgradeable.contract.WatchLogs(opts, "MocaSBTCrossChainPackage", destChainIdRule, channelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMOCACrossChainUpgradeableMocaSBTCrossChainPackage)
				if err := _IMOCACrossChainUpgradeable.contract.UnpackLog(event, "MocaSBTCrossChainPackage", log); err != nil {
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

// ParseMocaSBTCrossChainPackage is a log parse operation binding the contract event 0xeae7aa948aa4486965776005e20135ce32c4e9a9bd3704ec53d108056bcba038.
//
// Solidity: event MocaSBTCrossChainPackage(uint32 srcChainId, uint32 indexed destChainId, uint32 indexed channelId, uint256 indexed sequence, bytes payload)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableFilterer) ParseMocaSBTCrossChainPackage(log types.Log) (*IMOCACrossChainUpgradeableMocaSBTCrossChainPackage, error) {
	event := new(IMOCACrossChainUpgradeableMocaSBTCrossChainPackage)
	if err := _IMOCACrossChainUpgradeable.contract.UnpackLog(event, "MocaSBTCrossChainPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMOCACrossChainUpgradeableRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IMOCACrossChainUpgradeable contract.
type IMOCACrossChainUpgradeableRoleAdminChangedIterator struct {
	Event *IMOCACrossChainUpgradeableRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IMOCACrossChainUpgradeableRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMOCACrossChainUpgradeableRoleAdminChanged)
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
		it.Event = new(IMOCACrossChainUpgradeableRoleAdminChanged)
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
func (it *IMOCACrossChainUpgradeableRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMOCACrossChainUpgradeableRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMOCACrossChainUpgradeableRoleAdminChanged represents a RoleAdminChanged event raised by the IMOCACrossChainUpgradeable contract.
type IMOCACrossChainUpgradeableRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IMOCACrossChainUpgradeableRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IMOCACrossChainUpgradeable.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IMOCACrossChainUpgradeableRoleAdminChangedIterator{contract: _IMOCACrossChainUpgradeable.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IMOCACrossChainUpgradeableRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IMOCACrossChainUpgradeable.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMOCACrossChainUpgradeableRoleAdminChanged)
				if err := _IMOCACrossChainUpgradeable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableFilterer) ParseRoleAdminChanged(log types.Log) (*IMOCACrossChainUpgradeableRoleAdminChanged, error) {
	event := new(IMOCACrossChainUpgradeableRoleAdminChanged)
	if err := _IMOCACrossChainUpgradeable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMOCACrossChainUpgradeableRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IMOCACrossChainUpgradeable contract.
type IMOCACrossChainUpgradeableRoleGrantedIterator struct {
	Event *IMOCACrossChainUpgradeableRoleGranted // Event containing the contract specifics and raw log

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
func (it *IMOCACrossChainUpgradeableRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMOCACrossChainUpgradeableRoleGranted)
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
		it.Event = new(IMOCACrossChainUpgradeableRoleGranted)
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
func (it *IMOCACrossChainUpgradeableRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMOCACrossChainUpgradeableRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMOCACrossChainUpgradeableRoleGranted represents a RoleGranted event raised by the IMOCACrossChainUpgradeable contract.
type IMOCACrossChainUpgradeableRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IMOCACrossChainUpgradeableRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IMOCACrossChainUpgradeable.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IMOCACrossChainUpgradeableRoleGrantedIterator{contract: _IMOCACrossChainUpgradeable.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IMOCACrossChainUpgradeableRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IMOCACrossChainUpgradeable.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMOCACrossChainUpgradeableRoleGranted)
				if err := _IMOCACrossChainUpgradeable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableFilterer) ParseRoleGranted(log types.Log) (*IMOCACrossChainUpgradeableRoleGranted, error) {
	event := new(IMOCACrossChainUpgradeableRoleGranted)
	if err := _IMOCACrossChainUpgradeable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMOCACrossChainUpgradeableRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IMOCACrossChainUpgradeable contract.
type IMOCACrossChainUpgradeableRoleRevokedIterator struct {
	Event *IMOCACrossChainUpgradeableRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IMOCACrossChainUpgradeableRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMOCACrossChainUpgradeableRoleRevoked)
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
		it.Event = new(IMOCACrossChainUpgradeableRoleRevoked)
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
func (it *IMOCACrossChainUpgradeableRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMOCACrossChainUpgradeableRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMOCACrossChainUpgradeableRoleRevoked represents a RoleRevoked event raised by the IMOCACrossChainUpgradeable contract.
type IMOCACrossChainUpgradeableRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IMOCACrossChainUpgradeableRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IMOCACrossChainUpgradeable.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IMOCACrossChainUpgradeableRoleRevokedIterator{contract: _IMOCACrossChainUpgradeable.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IMOCACrossChainUpgradeableRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IMOCACrossChainUpgradeable.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMOCACrossChainUpgradeableRoleRevoked)
				if err := _IMOCACrossChainUpgradeable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IMOCACrossChainUpgradeable *IMOCACrossChainUpgradeableFilterer) ParseRoleRevoked(log types.Log) (*IMOCACrossChainUpgradeableRoleRevoked, error) {
	event := new(IMOCACrossChainUpgradeableRoleRevoked)
	if err := _IMOCACrossChainUpgradeable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

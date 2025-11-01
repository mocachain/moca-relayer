// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package universalVerifier

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

// VCDataLibEventData is an auto generated low-level Go binding around an user-defined struct.
type VCDataLibEventData struct {
	SrcChainId     uint32
	DestChainId    uint32
	ChannelId      uint32
	Sequence       *big.Int
	CrossDataBytes []byte
}

// ICircuitValidatorKeyToInputIndex is an auto generated low-level Go binding around an user-defined struct.
type ICircuitValidatorKeyToInputIndex struct {
	Key        string
	InputIndex *big.Int
}

// IZKPVerifierProofStatus is an auto generated low-level Go binding around an user-defined struct.
type IZKPVerifierProofStatus struct {
	IsVerified       bool
	ValidatorVersion string
	BlockNumber      *big.Int
	BlockTimestamp   *big.Int
}

// IZKPVerifierSubmitZkpResponseCross is an auto generated low-level Go binding around an user-defined struct.
type IZKPVerifierSubmitZkpResponseCross struct {
	Responses        []IZKPVerifierZKPResponse
	CrossChainProofs []byte
}

// IZKPVerifierSubmitZkpResponseStruct is an auto generated low-level Go binding around an user-defined struct.
type IZKPVerifierSubmitZkpResponseStruct struct {
	RequestId   uint64
	Inputs      []*big.Int
	A           [2]*big.Int
	B           [2][2]*big.Int
	C           [2]*big.Int
	UserAddress common.Address
}

// IZKPVerifierZKPRequest is an auto generated low-level Go binding around an user-defined struct.
type IZKPVerifierZKPRequest struct {
	Metadata  string
	Validator common.Address
	Data      []byte
}

// IZKPVerifierZKPResponse is an auto generated low-level Go binding around an user-defined struct.
type IZKPVerifierZKPResponse struct {
	RequestId uint64
	ZkProof   []byte
	Data      []byte
}

// UniversalVerifierMetaData contains all meta data concerning the UniversalVerifier contract.
var UniversalVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"linkID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"requestIdToCompare\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"linkIdToCompare\",\"type\":\"uint256\"}],\"name\":\"LinkedProofError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"useraddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"ErrorHandle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcChainId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"channelId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"crossDataBytes\",\"type\":\"bytes\"}],\"name\":\"VCCrossChainPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ZKPRequestSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ZKPRequestUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ZKPResponseSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REQUESTS_RETURN_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addValidatorToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"crossChainUpdateStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"crossUploadStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"disableZKPRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"srcChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"}],\"name\":\"emitCrossChainEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"enableZKPRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"getProofStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"validatorVersion\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIZKPVerifier.ProofStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getProofStorageField\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"getRequestOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"}],\"name\":\"getVCCrossChainSequence\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"getZKPRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest\",\"name\":\"zkpRequest\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getZKPRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getZKPRequestsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"grantOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIState\",\"name\":\"state\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"isProofVerified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isWhitelistedValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"isZKPRequestEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidatorFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"requestIdExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"requestOwner\",\"type\":\"address\"}],\"name\":\"setRequestOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIState\",\"name\":\"state\",\"type\":\"address\"}],\"name\":\"setState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"setZKPRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"requestIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest[]\",\"name\":\"requests\",\"type\":\"tuple[]\"}],\"name\":\"setZKPRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"uint256[]\",\"name\":\"inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"name\":\"submitZKPResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"uint256[]\",\"name\":\"inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"internalType\":\"structIZKPVerifier.SubmitZkpResponseStruct[]\",\"name\":\"response\",\"type\":\"tuple[]\"}],\"name\":\"submitZKPResponseBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"zkProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPResponse[]\",\"name\":\"responses\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"crossChainProofs\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.SubmitZkpResponseCross[]\",\"name\":\"crossResponses\",\"type\":\"tuple[]\"}],\"name\":\"submitZKPResponseBatchV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"uint256[]\",\"name\":\"inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"internalType\":\"structIZKPVerifier.SubmitZkpResponseStruct[]\",\"name\":\"response\",\"type\":\"tuple[]\"},{\"internalType\":\"uint32\",\"name\":\"srcChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"}],\"name\":\"submitZKPResponseBatchWithCrossChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"zkProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPResponse[]\",\"name\":\"responses\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"crossChainProof\",\"type\":\"bytes\"}],\"name\":\"submitZKPResponseV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"uint256[]\",\"name\":\"inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"submitZKPResponseWithUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"uint256[]\",\"name\":\"inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"srcChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"}],\"name\":\"submitZKPResponseWithUserCrossChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"updateZKPRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"requestIds\",\"type\":\"uint64[]\"}],\"name\":\"verifyLinkedProofs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"uint256[]\",\"name\":\"inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"verifyZKPResponse\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"inputIndex\",\"type\":\"uint256\"}],\"internalType\":\"structICircuitValidator.KeyToInputIndex[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50614d2d8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106103415760003560e01c806382aff29f116101bd578063b1eeab4f116100f9578063d4858051116100a2578063e348da131161007c578063e348da1314610814578063f2fde38b14610827578063f5b541a61461083a578063ffa1ad741461086157600080fd5b8063d4858051146107e6578063d547741f146107f9578063e30c39781461080c57600080fd5b8063c76d0845116100d3578063c76d0845146107a0578063c9c23ea5146107c0578063cde09fad146107d357600080fd5b8063b1eeab4f14610767578063b68967e21461077a578063c613b4081461078d57600080fd5b80639b36d03611610166578063a217fddf11610140578063a217fddf146106f5578063a527305e146106fd578063ab7bcfb714610710578063ade09fcd1461075457600080fd5b80639b36d036146106845780639f5223e0146106cf5780639f66c198146106e257600080fd5b80638c938ce4116101975780638c938ce4146106115780638da5cb5b1461062457806391d148541461062c57600080fd5b806382aff29f146105cb578063862b304c146105de5780638c1da2c9146105f157600080fd5b806349555fb11161028c5780636508e1b411610235578063715018a61161020f578063715018a6146105955780637559560e1461059d57806379ba5097146105b05780637fd7684c146105b857600080fd5b80636508e1b4146105485780636882ee951461056f5780636f2477ad1461058257600080fd5b806359f871a11161026657806359f871a1146105025780635f9e60d714610515578063617879fe1461053557600080fd5b806349555fb1146104a55780635176983b146104b857806354fd4d50146104d857600080fd5b80632f2ff15d116102ee57806334c901af116102c857806334c901af1461046c57806336568abe1461047f578063485cc9551461049257600080fd5b80632f2ff15d1461040257806330def4ac1461041557806331969e571461042857600080fd5b8063130a73ac1161031f578063130a73ac146103965780631905e7b1146103a9578063248a9ca3146103c057600080fd5b80630106a95914610346578063011f5bd61461035b57806301ffc9a71461036e575b600080fd5b6103596103543660046135fa565b610885565b005b61035961036936600461366a565b610af6565b61038161037c3660046136be565b610b78565b60405190151581526020015b60405180910390f35b6103816103a4366004613707565b610bf8565b6103b26103e881565b60405190815260200161038d565b6103b26103ce366004613722565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b61035961041036600461373b565b610cbf565b610359610423366004613707565b610d09565b7f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a03546001600160a01b03165b6040516001600160a01b03909116815260200161038d565b61035961047a366004613760565b610da5565b61035961048d36600461373b565b610ded565b6103596104a036600461377d565b610e39565b6103816104b33660046137ab565b610fb5565b6104cb6104c63660046139d1565b611093565b60405161038d9190613ab9565b604080518082019091526005815264189718971b60d91b60208201525b60405161038d9190613b32565b610359610510366004613760565b6110ae565b610528610523366004613b45565b6110bf565b60405161038d9190613bb1565b610359610543366004613760565b61110a565b7f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a02546103b2565b61035961057d366004613c0a565b61111b565b610454610590366004613707565b61117c565b610359611248565b6103596105ab3660046139d1565b61125c565b6103596112a8565b6103596105c6366004613c8d565b6112ed565b6103596105d9366004613707565b611450565b6103596105ec3660046137ab565b6114e7565b6106046105ff3660046137ab565b61151b565b60405161038d9190613d47565b61035961061f366004613d8f565b611659565b610454611757565b61038161063a36600461373b565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b610381610692366004613760565b6001600160a01b031660009081527f76aa24e3538905838cc74060b2aa4c054b1e474aacf44741879e1850715e9300602052604090205460ff1690565b6103596106dd36600461366a565b61178c565b6103b26106f0366004613de5565b6117cc565b6103b2600081565b61035961070b366004613f7b565b611811565b61038161071e366004613707565b6001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b61035961076236600461409f565b611919565b6103596107753660046137ab565b611994565b610359610788366004614106565b611a3b565b61035961079b3660046142a7565b611a85565b6107b36107ae366004613707565b611c2c565b60405161038d9190614304565b6103596107ce366004614317565b611e3f565b6103596107e1366004614333565b611ed7565b6103b26107f436600461436f565b61207a565b61035961080736600461373b565b612179565b6104546121bd565b610359610822366004613760565b6121e6565b610359610835366004613760565b61223a565b6103b27f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92981565b6104f560405180604001604052806005815260200164189718971b60d91b81525081565b600181116109005760405162461bcd60e51b815260206004820152603360248201527f4c696e6b65642070726f6f6620766572696669636174696f6e206e656564732060448201527f6d6f7265207468616e203120726571756573740000000000000000000000000060648201526084015b60405180910390fd5b60006109698484846000818110610919576109196143e2565b905060200201602081019061092e9190613707565b6040518060400160405280600681526020017f6c696e6b4944000000000000000000000000000000000000000000000000000081525061207a565b9050806000036109e15760405162461bcd60e51b815260206004820152603860248201527f43616e27742066696e64206c696e6b494420666f7220676976656e207265717560448201527f6573742049647320616e6420757365722061646472657373000000000000000060648201526084016108f7565b60015b82811015610aef576000610a0486868685818110610919576109196143e2565b9050808314610ae65784846000818110610a2057610a206143e2565b9050602002016020810190610a359190613707565b83868685818110610a4857610a486143e2565b9050602002016020810190610a5d9190613707565b6040517f0cb82c0300000000000000000000000000000000000000000000000000000000815260a06004820152601560a48201527f50726f6f667320617265206e6f74206c696e6b6564000000000000000000000060c48201526001600160401b039384166024820152604481019290925290911660648201526084810182905260e4016108f7565b506001016109e4565b5050505050565b610afe6122bf565b610b0882826122f1565b336001600160401b0383167fa2c61fcbd9637e91178d0dea7f9b5cce13f60c453603b3dc056b8f01bb3d4cb0610b3e84806143f8565b610b4e6040870160208801613760565b610b5b60408801886143f8565b604051610b6c959493929190614467565b60405180910390a35050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610bf257506301ffc9a760e01b7fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6000816001610c37826001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b610c7e5760405162461bcd60e51b81526020600482015260186024820152771c995c5d595cdd081a5908191bd95cdb89dd08195e1a5cdd60421b60448201526064016108f7565b5050506001600160401b031660009081527f70325635d67d74932012fa921ccb2f335d3b1d69e3a487f50d001cc65f531600602052604090205460ff161590565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610cf9816123c8565b610d0383836123d2565b50505050565b8033610d148261117c565b6001600160a01b0316816001600160a01b03161480610d4b5750610d36611757565b6001600160a01b0316816001600160a01b0316145b610d975760405162461bcd60e51b815260206004820152601d60248201527f4e6f7420616e206f776e6572206f722072657175657374206f776e657200000060448201526064016108f7565b610da0836124a1565b505050565b610dad6122bf565b7f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a0380546001600160a01b0319166001600160a01b03831617905550565b50565b6001600160a01b0381163314610e2f576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610da08282612579565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff1615906001600160401b0316600081158015610e835750825b90506000826001600160401b03166001148015610e9f5750303b155b905081158015610ead575080155b15610ee4576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610f1857845468ff00000000000000001916680100000000000000001785555b610f218661261f565b610f2a87612630565b610f356000876123d2565b50610f607f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929876123d2565b508315610fac57845468ff000000000000000019168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b6000816001610ff4826001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b61103b5760405162461bcd60e51b81526020600482015260186024820152771c995c5d595cdd081a5908191bd95cdb89dd08195e1a5cdd60421b60448201526064016108f7565b7f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a006001600160a01b0386166000908152602091825260408082206001600160401b0388168352909252205460ff169250505092915050565b60606110a3878787878787612641565b979650505050505050565b6110b66122bf565b610dea816126fb565b60405162461bcd60e51b815260206004820152601f60248201527f4e6f7420696d706c656d656e74656420696e20746869732076657273696f6e0060448201526060906064016108f7565b6111126122bf565b610dea81612846565b60005b83811015610aef5761117485858381811061113b5761113b6143e2565b90506020020160208101906111509190613707565b848484818110611162576111626143e2565b90506020028101906106dd919061449d565b60010161111e565b60008160016111bb826001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b6112025760405162461bcd60e51b81526020600482015260186024820152771c995c5d595cdd081a5908191bd95cdb89dd08195e1a5cdd60421b60448201526064016108f7565b5050506001600160401b031660009081527f6209bdc3799f5201408f7a7d4d471bb2a0100353e618451674b93f730b006a0060205260409020546001600160a01b031690565b6112506122bf565b61125a600061286e565b565b61126a8686868686866128a6565b60405133906001600160401b038816907f6979bc9c3e552c05dd9859285f1ed7a172e52ef39e1dce9c720e5bf8d82c9f4c90600090a3505050505050565b33806112b26121bd565b6001600160a01b0316146112e45760405163118cdaa760e01b81526001600160a01b03821660048201526024016108f7565b610dea8161286e565b6112fb8888888888886128a6565b60405133906001600160401b038a16907f6979bc9c3e552c05dd9859285f1ed7a172e52ef39e1dce9c720e5bf8d82c9f4c90600090a3600061133d848a61151b565b805190915061134c5750611446565b63ffffffff821660009081526020819052604081205480156113965790508080611375816144d3565b63ffffffff8616600090815260208190526040902081905591506113b09050565b63ffffffff84166000908152602081905260409020600190555b6040805180820182526001600160401b038d168082526001600160a01b03898116602080850191825285519081019390935251168184015282518082038401815260609091019283905290918490600b9063ffffffff8916907f2f4796132af44fdc7d0bfd7fe25cd97e2b6e8981ca28b81dfd6e22f065a01c5090611438908c9087906144ec565b60405180910390a450505050505b5050505050505050565b803361145b8261117c565b6001600160a01b0316816001600160a01b03161480611492575061147d611757565b6001600160a01b0316816001600160a01b0316145b6114de5760405162461bcd60e51b815260206004820152601d60248201527f4e6f7420616e206f776e6572206f722072657175657374206f776e657200000060448201526064016108f7565b610da0836129f9565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929611511816123c8565b610da08383611994565b61154860405180608001604052806000151581526020016060815260200160008152602001600081525090565b6001600160a01b03831660009081527f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a00602090815260408083206001600160401b0386168452825291829020825160808101909352805460ff16151583526002810180549193928301916115bb9061450b565b80601f01602080910402602001604051908101604052809291908181526020018280546115e79061450b565b80156116345780601f1061160957610100808354040283529160200191611634565b820191906000526020600020905b81548152906001019060200180831161161757829003601f168201915b5050505050815260200182600301548152602001826004015481525091505092915050565b63ffffffff811660009081526020819052604081205480156116a35790508080611682816144d3565b63ffffffff8516600090815260208190526040902081905591506116bd9050565b63ffffffff83166000908152602081905260409020600190555b6040805180820182526001600160401b0388168082526001600160a01b03888116602080850191825285519081019390935251168184015282518082038401815260609091019283905290918490600b9063ffffffff8816907f2f4796132af44fdc7d0bfd7fe25cd97e2b6e8981ca28b81dfd6e22f065a01c5090611745908b9087906144ec565b60405180910390a45050505050505050565b6000807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005b546001600160a01b031692915050565b6117968282612aa5565b336001600160401b0383167f26db2c5d21d517fa8c11af3ae555af7f4a560b9f70a5a32e73e6cd94d0cae3e8610b3e84806143f8565b63ffffffff8116600090815260208190526040812054156118095763ffffffff8216600090815260208190526040902054610bf29060019061453f565b506000919050565b60005b815181101561191557611861828281518110611832576118326143e2565b602002602001015160000151838381518110611850576118506143e2565b602002602001015160200151612b46565b60005b828281518110611876576118766143e2565b6020026020010151600001515181101561190c57336001600160a01b03168383815181106118a6576118a66143e2565b60200260200101516000015182815181106118c3576118c36143e2565b6020026020010151600001516001600160401b03167f6979bc9c3e552c05dd9859285f1ed7a172e52ef39e1dce9c720e5bf8d82c9f4c60405160405180910390a3600101611864565b50600101611814565b5050565b6119238282612b46565b60005b8251811015610da057336001600160a01b031683828151811061194b5761194b6143e2565b6020026020010151600001516001600160401b03167f6979bc9c3e552c05dd9859285f1ed7a172e52ef39e1dce9c720e5bf8d82c9f4c60405160405180910390a3600101611926565b6040517f63052f350000000000000000000000000000000000000000000000000000000081527f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a009060609073__$869f49b84e27953e40b9102dbb6d2c5983$__906363052f3590611a0f90859088908890879060040161457b565b60006040518083038186803b158015611a2757600080fd5b505af4158015611446573d6000803e3d6000fd5b611a488585858585612ddb565b60405133906001600160401b038716907f6979bc9c3e552c05dd9859285f1ed7a172e52ef39e1dce9c720e5bf8d82c9f4c90600090a35050505050565b60005b8351811015610d0357600080306001600160a01b0316868481518110611ab057611ab06143e2565b602002602001015160000151878581518110611ace57611ace6143e2565b602002602001015160200151888681518110611aec57611aec6143e2565b602002602001015160400151898781518110611b0a57611b0a6143e2565b6020026020010151606001518a8881518110611b2857611b286143e2565b6020026020010151608001518b8981518110611b4657611b466143e2565b602002602001015160a001518b8b604051602401611b6b989796959493929190614694565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f7fd7684c0000000000000000000000000000000000000000000000000000000017905251611bce9190614714565b6000604051808303816000865af19150503d8060008114611c0b576040519150601f19603f3d011682016040523d82523d6000602084013e611c10565b606091505b509150915081611c21575050611c24565b50505b600101611a88565b60408051606080820183528082526000602083015291810191909152816001611c85826001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b611ccc5760405162461bcd60e51b81526020600482015260186024820152771c995c5d595cdd081a5908191bd95cdb89dd08195e1a5cdd60421b60448201526064016108f7565b6001600160401b0384166000908152600080516020614cd8833981519152602052604090819020815160608101909252805482908290611d0b9061450b565b80601f0160208091040260200160405190810160405280929190818152602001828054611d379061450b565b8015611d845780601f10611d5957610100808354040283529160200191611d84565b820191906000526020600020905b815481529060010190602001808311611d6757829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191611db49061450b565b80601f0160208091040260200160405190810160405280929190818152602001828054611de09061450b565b8015611e2d5780601f10611e0257610100808354040283529160200191611e2d565b820191906000526020600020905b815481529060010190602001808311611e1057829003601f168201915b50505050508152505092505050919050565b8133611e4a8261117c565b6001600160a01b0316816001600160a01b03161480611e815750611e6c611757565b6001600160a01b0316816001600160a01b0316145b611ecd5760405162461bcd60e51b815260206004820152601d60248201527f4e6f7420616e206f776e6572206f722072657175657374206f776e657200000060448201526064016108f7565b610d038484612e04565b60005b815181101561191557600080306001600160a01b0316848481518110611f0257611f026143e2565b602002602001015160000151858581518110611f2057611f206143e2565b602002602001015160200151868681518110611f3e57611f3e6143e2565b602002602001015160400151878781518110611f5c57611f5c6143e2565b602002602001015160600151888881518110611f7a57611f7a6143e2565b602002602001015160800151898981518110611f9857611f986143e2565b602002602001015160a00151604051602401611fb996959493929190614726565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f7559560e000000000000000000000000000000000000000000000000000000001790525161201c9190614714565b6000604051808303816000865af19150503d8060008114612059576040519150601f19603f3d011682016040523d82523d6000602084013e61205e565b606091505b50915091508161206f575050612072565b50505b600101611eda565b60008260016120b9826001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b6121005760405162461bcd60e51b81526020600482015260186024820152771c995c5d595cdd081a5908191bd95cdb89dd08195e1a5cdd60421b60448201526064016108f7565b7f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a006001600160a01b0387166000908152602091825260408082206001600160401b0389168352909252819020905160019091019061215f908690614714565b908152602001604051809103902054925050509392505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546121b3816123c8565b610d038383612579565b6000807f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0061177c565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929612210816123c8565b610da07f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929836123d2565b6122426122bf565b7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0080546001600160a01b0319166001600160a01b0383169081178255612286611757565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a35050565b336122c8611757565b6001600160a01b03161461125a5760405163118cdaa760e01b81523360048201526024016108f7565b81600161232e826001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b6123755760405162461bcd60e51b81526020600482015260186024820152771c995c5d595cdd081a5908191bd95cdb89dd08195e1a5cdd60421b60448201526064016108f7565b6001600160401b0384166000908152600080516020614cd8833981519152602052604090207f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a00908490610fac82826148a2565b610dea8133612ee3565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612497576000848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561244d3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610bf2565b6000915050610bf2565b8060016124de826001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b6125255760405162461bcd60e51b81526020600482015260186024820152771c995c5d595cdd081a5908191bd95cdb89dd08195e1a5cdd60421b60448201526064016108f7565b60007f70325635d67d74932012fa921ccb2f335d3b1d69e3a487f50d001cc65f5316005b6001600160401b039490941660009081526020949094526040909320805460ff1916931515939093179092555050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612497576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610bf2565b612627612f70565b610dea81612fd7565b612638612f70565b610dea81613022565b6060600061264e88611c2c565b602001519050612695816001600160a01b031660009081527f76aa24e3538905838cc74060b2aa4c054b1e474aacf44741879e1850715e9300602052604090205460ff1690565b6126e15760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f742077686974656c69737465640000000060448201526064016108f7565b6126ef88888888888861302a565b98975050505050505050565b6040516301ffc9a760e01b81527f39db6a480000000000000000000000000000000000000000000000000000000060048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa15801561275f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061278391906149b4565b6127f55760405162461bcd60e51b815260206004820152602c60248201527f56616c696461746f7220646f65736e277420737570706f72742072656c65766160448201527f6e7420696e74657266616365000000000000000000000000000000000000000060648201526084016108f7565b60017f76aa24e3538905838cc74060b2aa4c054b1e474aacf44741879e1850715e93005b6001600160a01b039290921660009081526020929092526040909120805460ff1916911515919091179055565b60007f76aa24e3538905838cc74060b2aa4c054b1e474aacf44741879e1850715e9300612819565b7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0080546001600160a01b031916815561191582613090565b60007f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a005b905060006128d788613101565b60018101546040516378bbd34f60e11b81529192506000916001600160a01b039091169063f177a69e9061291c908b908b908b908b9060028a01908c90600401614a59565b6000604051808303816000875af115801561293b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526129639190810190614bad565b6040517fe48754ab00000000000000000000000000000000000000000000000000000000815290915073__$869f49b84e27953e40b9102dbb6d2c5983$__9063e48754ab906129be90869088908e9087908f90600401614bf5565b60006040518083038186803b1580156129d657600080fd5b505af41580156129ea573d6000803e3d6000fd5b50505050505050505050505050565b806001612a36826001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b612a7d5760405162461bcd60e51b81526020600482015260186024820152771c995c5d595cdd081a5908191bd95cdb89dd08195e1a5cdd60421b60448201526064016108f7565b60017f70325635d67d74932012fa921ccb2f335d3b1d69e3a487f50d001cc65f531600612549565b612ab56040820160208301613760565b6001600160a01b03811660009081527f76aa24e3538905838cc74060b2aa4c054b1e474aacf44741879e1850715e9300602052604090205460ff16612b3c5760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f742077686974656c69737465640000000060448201526064016108f7565b610da0838361316b565b7f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a03546040517fe28d24b00000000000000000000000000000000000000000000000000000000081527f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a00916001600160a01b03169063e28d24b090612bce908590600401613b32565b600060405180830381600087803b158015612be857600080fd5b505af1158015612bfc573d6000803e3d6000fd5b5050505060005b8351811015610d03576000848281518110612c2057612c206143e2565b602002602001015190506000612c333390565b90506000612c448360000151613101565b6001810154602085015160038801546040517fd6614a6b0000000000000000000000000000000000000000000000000000000081529394506000936001600160a01b039384169363d6614a6b93612ca893909260028901928a921690600401614c88565b6000604051808303816000875af1158015612cc7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612cef9190810190614bad565b84516040517fb2394c6000000000000000000000000000000000000000000000000000000000815291925073__$869f49b84e27953e40b9102dbb6d2c5983$__9163b2394c6091612d48918a918891879060040161457b565b60006040518083038186803b158015612d6057600080fd5b505af4158015612d74573d6000803e3d6000fd5b5050505060408401515115612dcb5760405162461bcd60e51b815260206004820152601a60248201527f4d65746164617461206e6f7420737570706f727465642079657400000000000060448201526064016108f7565b505060019092019150612c039050565b3360007f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a006128ca565b816001612e41826001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b612e885760405162461bcd60e51b81526020600482015260186024820152771c995c5d595cdd081a5908191bd95cdb89dd08195e1a5cdd60421b60448201526064016108f7565b50506001600160401b039190911660009081527f6209bdc3799f5201408f7a7d4d471bb2a0100353e618451674b93f730b006a006020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16611915576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044016108f7565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661125a576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612fdf612f70565b6001600160a01b0381166112e4576040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600060048201526024016108f7565b610dad612f70565b60608661303681610bf8565b6130825760405162461bcd60e51b815260206004820152601360248201527f526571756573742069732064697361626c65640000000000000000000000000060448201526064016108f7565b6126ef88888888888861317f565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b60008161310d81610bf8565b6131595760405162461bcd60e51b815260206004820152601360248201527f526571756573742069732064697361626c65640000000000000000000000000060448201526064016108f7565b613162836132bf565b91505b50919050565b6131758282613364565b6119158233612e04565b60608660016131be826001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b6132055760405162461bcd60e51b81526020600482015260186024820152771c995c5d595cdd081a5908191bd95cdb89dd08195e1a5cdd60421b60448201526064016108f7565b6001600160401b0389166000908152600080516020614cd8833981519152602052604090819020600181015491516378bbd34f60e11b815290916001600160a01b03169063f177a69e9061326a908c908c908c908c9060028901908d90600401614a59565b6000604051808303816000875af1158015613289573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526132b19190810190614bad565b9a9950505050505050505050565b60006132ca82611c2c565b6020015161330f816001600160a01b031660009081527f76aa24e3538905838cc74060b2aa4c054b1e474aacf44741879e1850715e9300602052604090205460ff1690565b61335b5760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f742077686974656c69737465640000000060448201526064016108f7565b6131628361348e565b8160006133a1826001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b156133ee5760405162461bcd60e51b815260206004820152601960248201527f7265717565737420696420616c7265616479206578697374730000000000000060448201526064016108f7565b6001600160401b0384166000908152600080516020614cd8833981519152602052604090207f512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a0090849061344182826148a2565b50506002018054600181018255600091825260209091206004820401805460039092166008026101000a6001600160401b0381810219909316969092169190910294909417909355505050565b60008161349a81610bf8565b6134e65760405162461bcd60e51b815260206004820152601360248201527f526571756573742069732064697361626c65640000000000000000000000000060448201526064016108f7565b613162836000816001613529826001600160401b03166000908152600080516020614cd883398151915260205260409020600101546001600160a01b0316151590565b6135705760405162461bcd60e51b81526020600482015260186024820152771c995c5d595cdd081a5908191bd95cdb89dd08195e1a5cdd60421b60448201526064016108f7565b5050506001600160401b03166000908152600080516020614cd88339815191526020526040902090565b6001600160a01b0381168114610dea57600080fd5b60008083601f8401126135c157600080fd5b5081356001600160401b038111156135d857600080fd5b6020830191508360208260051b85010111156135f357600080fd5b9250929050565b60008060006040848603121561360f57600080fd5b833561361a8161359a565b925060208401356001600160401b0381111561363557600080fd5b613641868287016135af565b9497909650939450505050565b80356001600160401b038116811461366557600080fd5b919050565b6000806040838503121561367d57600080fd5b6136868361364e565b915060208301356001600160401b038111156136a157600080fd5b8301606081860312156136b357600080fd5b809150509250929050565b6000602082840312156136d057600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461370057600080fd5b9392505050565b60006020828403121561371957600080fd5b6137008261364e565b60006020828403121561373457600080fd5b5035919050565b6000806040838503121561374e57600080fd5b8235915060208301356136b38161359a565b60006020828403121561377257600080fd5b81356137008161359a565b6000806040838503121561379057600080fd5b823561379b8161359a565b915060208301356136b38161359a565b600080604083850312156137be57600080fd5b82356137c98161359a565b91506137d76020840161364e565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715613818576138186137e0565b60405290565b604080519081016001600160401b0381118282101715613818576138186137e0565b60405160c081016001600160401b0381118282101715613818576138186137e0565b604051601f8201601f191681016001600160401b038111828210171561388a5761388a6137e0565b604052919050565b60006001600160401b038211156138ab576138ab6137e0565b5060051b60200190565b600082601f8301126138c657600080fd5b81356138d96138d482613892565b613862565b8082825260208201915060208360051b8601019250858311156138fb57600080fd5b602085015b83811015613918578035835260209283019201613900565b5095945050505050565b600082601f83011261393357600080fd5b61393d6040613862565b80604084018581111561394f57600080fd5b845b81811015613969578035845260209384019301613951565b509095945050505050565b600082601f83011261398557600080fd5b604061399081613862565b8060808501868111156139a257600080fd5b855b818110156139c5576139b68882613922565b845260209093019284016139a4565b50909695505050505050565b60008060008060008061016087890312156139eb57600080fd5b6139f48761364e565b955060208701356001600160401b03811115613a0f57600080fd5b613a1b89828a016138b5565b955050613a2b8860408901613922565b9350613a3a8860808901613974565b9250613a4a886101008901613922565b9150610140870135613a5b8161359a565b809150509295509295509295565b60005b83811015613a84578181015183820152602001613a6c565b50506000910152565b60008151808452613aa5816020860160208601613a69565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613b2657603f198786030184528151805160408752613b076040880182613a8d565b6020928301519783019790975250938401939190910190600101613ae1565b50929695505050505050565b6020815260006137006020830184613a8d565b60008060408385031215613b5857600080fd5b50508035926020909101359150565b6000815160608452613b7c6060850182613a8d565b90506001600160a01b03602084015116602085015260408301518482036040860152613ba88282613a8d565b95945050505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613b2657603f19878603018452613bf5858351613b67565b94506020938401939190910190600101613bd9565b60008060008060408587031215613c2057600080fd5b84356001600160401b03811115613c3657600080fd5b613c42878288016135af565b90955093505060208501356001600160401b03811115613c6157600080fd5b613c6d878288016135af565b95989497509550505050565b803563ffffffff8116811461366557600080fd5b6000806000806000806000806101a0898b031215613caa57600080fd5b613cb38961364e565b975060208901356001600160401b03811115613cce57600080fd5b613cda8b828c016138b5565b975050613cea8a60408b01613922565b9550613cf98a60808b01613974565b9450613d098a6101008b01613922565b9350610140890135613d1a8161359a565b9250613d296101608a01613c79565b9150613d386101808a01613c79565b90509295985092959890939650565b602081528151151560208201526000602083015160806040840152613d6f60a0840182613a8d565b905060408401516060840152606084015160808401528091505092915050565b60008060008060808587031215613da557600080fd5b613dae8561364e565b93506020850135613dbe8161359a565b9250613dcc60408601613c79565b9150613dda60608601613c79565b905092959194509250565b600060208284031215613df757600080fd5b61370082613c79565b60006001600160401b03821115613e1957613e196137e0565b50601f01601f191660200190565b6000613e356138d484613e00565b9050828152838383011115613e4957600080fd5b828260208301376000602084830101529392505050565b600082601f830112613e7157600080fd5b61370083833560208501613e27565b600082601f830112613e9157600080fd5b8135613e9f6138d482613892565b8082825260208201915060208360051b860101925085831115613ec157600080fd5b602085015b838110156139185780356001600160401b03811115613ee457600080fd5b86016060818903601f19011215613efa57600080fd5b613f026137f6565b613f0e6020830161364e565b815260408201356001600160401b03811115613f2957600080fd5b613f388a602083860101613e60565b60208301525060608201356001600160401b03811115613f5757600080fd5b613f668a602083860101613e60565b60408301525084525060209283019201613ec6565b600060208284031215613f8d57600080fd5b81356001600160401b03811115613fa357600080fd5b8201601f81018413613fb457600080fd5b8035613fc26138d482613892565b8082825260208201915060208360051b850101925086831115613fe457600080fd5b602084015b838110156140945780356001600160401b0381111561400757600080fd5b85016040818a03601f1901121561401d57600080fd5b61402561381e565b60208201356001600160401b0381111561403e57600080fd5b61404d8b602083860101613e80565b82525060408201356001600160401b0381111561406957600080fd5b6140788b602083860101613e60565b6020830152508085525050602083019250602081019050613fe9565b509695505050505050565b600080604083850312156140b257600080fd5b82356001600160401b038111156140c857600080fd5b6140d485828601613e80565b92505060208301356001600160401b038111156140f057600080fd5b6140fc85828601613e60565b9150509250929050565b6000806000806000610140868803121561411f57600080fd5b6141288661364e565b945060208601356001600160401b0381111561414357600080fd5b61414f888289016138b5565b94505061415f8760408801613922565b925061416e8760808801613974565b915061417e876101008801613922565b90509295509295909350565b600082601f83011261419b57600080fd5b81356141a96138d482613892565b8082825260208201915060208360051b8601019250858311156141cb57600080fd5b602085015b838110156139185780356001600160401b038111156141ee57600080fd5b8601610160818903601f1901121561420557600080fd5b61420d613840565b6142196020830161364e565b815260408201356001600160401b0381111561423457600080fd5b6142438a6020838601016138b5565b6020830152506142568960608401613922565b60408201526142688960a08401613974565b606082015261427b896101208401613922565b608082015261016082013591506142918261359a565b60a08101919091528352602092830192016141d0565b6000806000606084860312156142bc57600080fd5b83356001600160401b038111156142d257600080fd5b6142de8682870161418a565b9350506142ed60208501613c79565b91506142fb60408501613c79565b90509250925092565b6020815260006137006020830184613b67565b6000806040838503121561432a57600080fd5b61379b8361364e565b60006020828403121561434557600080fd5b81356001600160401b0381111561435b57600080fd5b6143678482850161418a565b949350505050565b60008060006060848603121561438457600080fd5b833561438f8161359a565b925061439d6020850161364e565b915060408401356001600160401b038111156143b857600080fd5b8401601f810186136143c957600080fd5b6143d886823560208401613e27565b9150509250925092565b634e487b7160e01b600052603260045260246000fd5b6000808335601e1984360301811261440f57600080fd5b8301803591506001600160401b0382111561442957600080fd5b6020019150368190038213156135f357600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60608152600061447b60608301878961443e565b6001600160a01b038616602084015282810360408401526126ef81858761443e565b60008235605e198336030181126144b357600080fd5b9190910192915050565b634e487b7160e01b600052601160045260246000fd5b6000600182016144e5576144e56144bd565b5060010190565b63ffffffff831681526040602082015260006143676040830184613a8d565b600181811c9082168061451f57607f821691505b60208210810361316557634e487b7160e01b600052602260045260246000fd5b81810381811115610bf257610bf26144bd565b60008151604084526145676040850182613a8d565b602093840151949093019390935250919050565b6000608082018683526001600160a01b03861660208401526001600160401b03851660408401526080606084015280845180835260a08501915060a08160051b86010192506020860160005b828110156145f857609f198786030184526145e3858351614552565b945060209384019391909101906001016145c7565b50929998505050505050505050565b600081518084526020840193506020830160005b8281101561463957815186526020958601959091019060010161461b565b5093949350505050565b8060005b6002811015610d03578151845260209384019390910190600101614647565b8060005b6002811015610d035761467e848351614643565b604093909301926020919091019060010161466a565b6001600160401b03891681526101a0602082015260006146b86101a083018a614607565b90506146c76040830189614643565b6146d46080830188614666565b6146e2610100830187614643565b6001600160a01b039490941661014082015263ffffffff92831661016082015291166101809091015295945050505050565b600082516144b3818460208701613a69565b6001600160401b03871681526101606020820152600061474a610160830188614607565b90506147596040830187614643565b6147666080830186614666565b614774610100830185614643565b6001600160a01b038316610140830152979650505050505050565b601f821115610da057806000526020600020601f840160051c810160208510156147b65750805b601f840160051c820191505b81811015610aef57600081556001016147c2565b60008135610bf28161359a565b6001600160401b038311156147fa576147fa6137e0565b61480e83614808835461450b565b8361478f565b6000601f841160018114614842576000851561482a5750838201355b600019600387901b1c1916600186901b178355610aef565b600083815260209020601f19861690835b828110156148735786850135825560209485019460019092019101614853565b50868210156148905760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b6148ac82836143f8565b6001600160401b038111156148c3576148c36137e0565b6148d7816148d1855461450b565b8561478f565b6000601f82116001811461490b57600083156148f35750838201355b600019600385901b1c1916600184901b178555614965565b600085815260209020601f19841690835b8281101561493c578685013582556020948501946001909201910161491c565b50848210156149595760001960f88660031b161c19848701351681555b505060018360011b0185555b50505050614999614978602084016147d6565b600183016001600160a01b0382166001600160a01b03198254161781555050565b6149a660408301836143f8565b610d038183600286016147e3565b6000602082840312156149c657600080fd5b8151801515811461370057600080fd5b600081546149e38161450b565b8085526001821680156149fd5760018114614a1957614a50565b60ff1983166020870152602082151560051b8701019350614a50565b84600052602060002060005b83811015614a475781546020828a010152600182019150602081019050614a25565b87016020019450505b50505092915050565b61016081526000614a6e610160830189614607565b614a7b6020840189614643565b614a886060840188614666565b614a9560e0840187614643565b828103610120840152614aa881866149d6565b9150506001600160a01b038316610140830152979650505050505050565b6000614ad46138d484613892565b838152905060208101600584901b830185811115614af157600080fd5b835b81811015614ba35780516001600160401b03811115614b1157600080fd5b850160006040828a031215614b24578081fd5b614b2c61381e565b905081516001600160401b03811115614b4457600080fd5b8201601f81018a13614b5557600080fd5b8051614b636138d482613e00565b8181528b6020838501011115614b7857600080fd5b614b89826020830160208601613a69565b835250506020918201518282015284529283019201614af3565b5050509392505050565b600060208284031215614bbf57600080fd5b81516001600160401b03811115614bd557600080fd5b8201601f81018413614be657600080fd5b61436784825160208401614ac6565b600060a082018783526001600160a01b03871660208401526001600160401b038616604084015260a0606084015280855180835260c08501915060c08160051b86010192506020870160005b82811015614c725760bf19878603018452614c5d858351614552565b94506020938401939190910190600101614c41565b5050505082810360808401526126ef8185614607565b608081526000614c9b6080830187613a8d565b8281036020840152614cad81876149d6565b9150506001600160a01b03841660408301526001600160a01b03831660608301529594505050505056fe512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a01a26469706673582212204efdbd24aec78688bb82d6cc80310e1eb0caa0a0b67db7f2938df5cd35627c7864736f6c634300081b0033",
}

// UniversalVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use UniversalVerifierMetaData.ABI instead.
var UniversalVerifierABI = UniversalVerifierMetaData.ABI

// UniversalVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniversalVerifierMetaData.Bin instead.
var UniversalVerifierBin = UniversalVerifierMetaData.Bin

// DeployUniversalVerifier deploys a new Ethereum contract, binding an instance of UniversalVerifier to it.
func DeployUniversalVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniversalVerifier, error) {
	parsed, err := UniversalVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniversalVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniversalVerifier{UniversalVerifierCaller: UniversalVerifierCaller{contract: contract}, UniversalVerifierTransactor: UniversalVerifierTransactor{contract: contract}, UniversalVerifierFilterer: UniversalVerifierFilterer{contract: contract}}, nil
}

// UniversalVerifier is an auto generated Go binding around an Ethereum contract.
type UniversalVerifier struct {
	UniversalVerifierCaller     // Read-only binding to the contract
	UniversalVerifierTransactor // Write-only binding to the contract
	UniversalVerifierFilterer   // Log filterer for contract events
}

// UniversalVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniversalVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniversalVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniversalVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniversalVerifierSession struct {
	Contract     *UniversalVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniversalVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniversalVerifierCallerSession struct {
	Contract *UniversalVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UniversalVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniversalVerifierTransactorSession struct {
	Contract     *UniversalVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UniversalVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniversalVerifierRaw struct {
	Contract *UniversalVerifier // Generic contract binding to access the raw methods on
}

// UniversalVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniversalVerifierCallerRaw struct {
	Contract *UniversalVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// UniversalVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniversalVerifierTransactorRaw struct {
	Contract *UniversalVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversalVerifier creates a new instance of UniversalVerifier, bound to a specific deployed contract.
func NewUniversalVerifier(address common.Address, backend bind.ContractBackend) (*UniversalVerifier, error) {
	contract, err := bindUniversalVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifier{UniversalVerifierCaller: UniversalVerifierCaller{contract: contract}, UniversalVerifierTransactor: UniversalVerifierTransactor{contract: contract}, UniversalVerifierFilterer: UniversalVerifierFilterer{contract: contract}}, nil
}

// NewUniversalVerifierCaller creates a new read-only instance of UniversalVerifier, bound to a specific deployed contract.
func NewUniversalVerifierCaller(address common.Address, caller bind.ContractCaller) (*UniversalVerifierCaller, error) {
	contract, err := bindUniversalVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierCaller{contract: contract}, nil
}

// NewUniversalVerifierTransactor creates a new write-only instance of UniversalVerifier, bound to a specific deployed contract.
func NewUniversalVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversalVerifierTransactor, error) {
	contract, err := bindUniversalVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierTransactor{contract: contract}, nil
}

// NewUniversalVerifierFilterer creates a new log filterer instance of UniversalVerifier, bound to a specific deployed contract.
func NewUniversalVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversalVerifierFilterer, error) {
	contract, err := bindUniversalVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierFilterer{contract: contract}, nil
}

// bindUniversalVerifier binds a generic wrapper to an already deployed contract.
func bindUniversalVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniversalVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalVerifier *UniversalVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalVerifier.Contract.UniversalVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalVerifier *UniversalVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.UniversalVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalVerifier *UniversalVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.UniversalVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalVerifier *UniversalVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalVerifier *UniversalVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalVerifier *UniversalVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UniversalVerifier *UniversalVerifierCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UniversalVerifier *UniversalVerifierSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UniversalVerifier.Contract.DEFAULTADMINROLE(&_UniversalVerifier.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UniversalVerifier *UniversalVerifierCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UniversalVerifier.Contract.DEFAULTADMINROLE(&_UniversalVerifier.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_UniversalVerifier *UniversalVerifierCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_UniversalVerifier *UniversalVerifierSession) OPERATORROLE() ([32]byte, error) {
	return _UniversalVerifier.Contract.OPERATORROLE(&_UniversalVerifier.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_UniversalVerifier *UniversalVerifierCallerSession) OPERATORROLE() ([32]byte, error) {
	return _UniversalVerifier.Contract.OPERATORROLE(&_UniversalVerifier.CallOpts)
}

// REQUESTSRETURNLIMIT is a free data retrieval call binding the contract method 0x1905e7b1.
//
// Solidity: function REQUESTS_RETURN_LIMIT() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCaller) REQUESTSRETURNLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "REQUESTS_RETURN_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REQUESTSRETURNLIMIT is a free data retrieval call binding the contract method 0x1905e7b1.
//
// Solidity: function REQUESTS_RETURN_LIMIT() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierSession) REQUESTSRETURNLIMIT() (*big.Int, error) {
	return _UniversalVerifier.Contract.REQUESTSRETURNLIMIT(&_UniversalVerifier.CallOpts)
}

// REQUESTSRETURNLIMIT is a free data retrieval call binding the contract method 0x1905e7b1.
//
// Solidity: function REQUESTS_RETURN_LIMIT() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCallerSession) REQUESTSRETURNLIMIT() (*big.Int, error) {
	return _UniversalVerifier.Contract.REQUESTSRETURNLIMIT(&_UniversalVerifier.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UniversalVerifier *UniversalVerifierCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UniversalVerifier *UniversalVerifierSession) VERSION() (string, error) {
	return _UniversalVerifier.Contract.VERSION(&_UniversalVerifier.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UniversalVerifier *UniversalVerifierCallerSession) VERSION() (string, error) {
	return _UniversalVerifier.Contract.VERSION(&_UniversalVerifier.CallOpts)
}

// GetProofStatus is a free data retrieval call binding the contract method 0x8c1da2c9.
//
// Solidity: function getProofStatus(address sender, uint64 requestId) view returns((bool,string,uint256,uint256))
func (_UniversalVerifier *UniversalVerifierCaller) GetProofStatus(opts *bind.CallOpts, sender common.Address, requestId uint64) (IZKPVerifierProofStatus, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getProofStatus", sender, requestId)

	if err != nil {
		return *new(IZKPVerifierProofStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IZKPVerifierProofStatus)).(*IZKPVerifierProofStatus)

	return out0, err

}

// GetProofStatus is a free data retrieval call binding the contract method 0x8c1da2c9.
//
// Solidity: function getProofStatus(address sender, uint64 requestId) view returns((bool,string,uint256,uint256))
func (_UniversalVerifier *UniversalVerifierSession) GetProofStatus(sender common.Address, requestId uint64) (IZKPVerifierProofStatus, error) {
	return _UniversalVerifier.Contract.GetProofStatus(&_UniversalVerifier.CallOpts, sender, requestId)
}

// GetProofStatus is a free data retrieval call binding the contract method 0x8c1da2c9.
//
// Solidity: function getProofStatus(address sender, uint64 requestId) view returns((bool,string,uint256,uint256))
func (_UniversalVerifier *UniversalVerifierCallerSession) GetProofStatus(sender common.Address, requestId uint64) (IZKPVerifierProofStatus, error) {
	return _UniversalVerifier.Contract.GetProofStatus(&_UniversalVerifier.CallOpts, sender, requestId)
}

// GetProofStorageField is a free data retrieval call binding the contract method 0xd4858051.
//
// Solidity: function getProofStorageField(address user, uint64 requestId, string key) view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCaller) GetProofStorageField(opts *bind.CallOpts, user common.Address, requestId uint64, key string) (*big.Int, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getProofStorageField", user, requestId, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProofStorageField is a free data retrieval call binding the contract method 0xd4858051.
//
// Solidity: function getProofStorageField(address user, uint64 requestId, string key) view returns(uint256)
func (_UniversalVerifier *UniversalVerifierSession) GetProofStorageField(user common.Address, requestId uint64, key string) (*big.Int, error) {
	return _UniversalVerifier.Contract.GetProofStorageField(&_UniversalVerifier.CallOpts, user, requestId, key)
}

// GetProofStorageField is a free data retrieval call binding the contract method 0xd4858051.
//
// Solidity: function getProofStorageField(address user, uint64 requestId, string key) view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetProofStorageField(user common.Address, requestId uint64, key string) (*big.Int, error) {
	return _UniversalVerifier.Contract.GetProofStorageField(&_UniversalVerifier.CallOpts, user, requestId, key)
}

// GetRequestOwner is a free data retrieval call binding the contract method 0x6f2477ad.
//
// Solidity: function getRequestOwner(uint64 requestId) view returns(address)
func (_UniversalVerifier *UniversalVerifierCaller) GetRequestOwner(opts *bind.CallOpts, requestId uint64) (common.Address, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getRequestOwner", requestId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRequestOwner is a free data retrieval call binding the contract method 0x6f2477ad.
//
// Solidity: function getRequestOwner(uint64 requestId) view returns(address)
func (_UniversalVerifier *UniversalVerifierSession) GetRequestOwner(requestId uint64) (common.Address, error) {
	return _UniversalVerifier.Contract.GetRequestOwner(&_UniversalVerifier.CallOpts, requestId)
}

// GetRequestOwner is a free data retrieval call binding the contract method 0x6f2477ad.
//
// Solidity: function getRequestOwner(uint64 requestId) view returns(address)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetRequestOwner(requestId uint64) (common.Address, error) {
	return _UniversalVerifier.Contract.GetRequestOwner(&_UniversalVerifier.CallOpts, requestId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UniversalVerifier *UniversalVerifierCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UniversalVerifier *UniversalVerifierSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UniversalVerifier.Contract.GetRoleAdmin(&_UniversalVerifier.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UniversalVerifier.Contract.GetRoleAdmin(&_UniversalVerifier.CallOpts, role)
}

// GetStateAddress is a free data retrieval call binding the contract method 0x31969e57.
//
// Solidity: function getStateAddress() view returns(address)
func (_UniversalVerifier *UniversalVerifierCaller) GetStateAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getStateAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStateAddress is a free data retrieval call binding the contract method 0x31969e57.
//
// Solidity: function getStateAddress() view returns(address)
func (_UniversalVerifier *UniversalVerifierSession) GetStateAddress() (common.Address, error) {
	return _UniversalVerifier.Contract.GetStateAddress(&_UniversalVerifier.CallOpts)
}

// GetStateAddress is a free data retrieval call binding the contract method 0x31969e57.
//
// Solidity: function getStateAddress() view returns(address)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetStateAddress() (common.Address, error) {
	return _UniversalVerifier.Contract.GetStateAddress(&_UniversalVerifier.CallOpts)
}

// GetVCCrossChainSequence is a free data retrieval call binding the contract method 0x9f66c198.
//
// Solidity: function getVCCrossChainSequence(uint32 chainId) view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCaller) GetVCCrossChainSequence(opts *bind.CallOpts, chainId uint32) (*big.Int, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getVCCrossChainSequence", chainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVCCrossChainSequence is a free data retrieval call binding the contract method 0x9f66c198.
//
// Solidity: function getVCCrossChainSequence(uint32 chainId) view returns(uint256)
func (_UniversalVerifier *UniversalVerifierSession) GetVCCrossChainSequence(chainId uint32) (*big.Int, error) {
	return _UniversalVerifier.Contract.GetVCCrossChainSequence(&_UniversalVerifier.CallOpts, chainId)
}

// GetVCCrossChainSequence is a free data retrieval call binding the contract method 0x9f66c198.
//
// Solidity: function getVCCrossChainSequence(uint32 chainId) view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetVCCrossChainSequence(chainId uint32) (*big.Int, error) {
	return _UniversalVerifier.Contract.GetVCCrossChainSequence(&_UniversalVerifier.CallOpts, chainId)
}

// GetZKPRequest is a free data retrieval call binding the contract method 0xc76d0845.
//
// Solidity: function getZKPRequest(uint64 requestId) view returns((string,address,bytes) zkpRequest)
func (_UniversalVerifier *UniversalVerifierCaller) GetZKPRequest(opts *bind.CallOpts, requestId uint64) (IZKPVerifierZKPRequest, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getZKPRequest", requestId)

	if err != nil {
		return *new(IZKPVerifierZKPRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IZKPVerifierZKPRequest)).(*IZKPVerifierZKPRequest)

	return out0, err

}

// GetZKPRequest is a free data retrieval call binding the contract method 0xc76d0845.
//
// Solidity: function getZKPRequest(uint64 requestId) view returns((string,address,bytes) zkpRequest)
func (_UniversalVerifier *UniversalVerifierSession) GetZKPRequest(requestId uint64) (IZKPVerifierZKPRequest, error) {
	return _UniversalVerifier.Contract.GetZKPRequest(&_UniversalVerifier.CallOpts, requestId)
}

// GetZKPRequest is a free data retrieval call binding the contract method 0xc76d0845.
//
// Solidity: function getZKPRequest(uint64 requestId) view returns((string,address,bytes) zkpRequest)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetZKPRequest(requestId uint64) (IZKPVerifierZKPRequest, error) {
	return _UniversalVerifier.Contract.GetZKPRequest(&_UniversalVerifier.CallOpts, requestId)
}

// GetZKPRequests is a free data retrieval call binding the contract method 0x5f9e60d7.
//
// Solidity: function getZKPRequests(uint256 startIndex, uint256 length) view returns((string,address,bytes)[])
func (_UniversalVerifier *UniversalVerifierCaller) GetZKPRequests(opts *bind.CallOpts, startIndex *big.Int, length *big.Int) ([]IZKPVerifierZKPRequest, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getZKPRequests", startIndex, length)

	if err != nil {
		return *new([]IZKPVerifierZKPRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]IZKPVerifierZKPRequest)).(*[]IZKPVerifierZKPRequest)

	return out0, err

}

// GetZKPRequests is a free data retrieval call binding the contract method 0x5f9e60d7.
//
// Solidity: function getZKPRequests(uint256 startIndex, uint256 length) view returns((string,address,bytes)[])
func (_UniversalVerifier *UniversalVerifierSession) GetZKPRequests(startIndex *big.Int, length *big.Int) ([]IZKPVerifierZKPRequest, error) {
	return _UniversalVerifier.Contract.GetZKPRequests(&_UniversalVerifier.CallOpts, startIndex, length)
}

// GetZKPRequests is a free data retrieval call binding the contract method 0x5f9e60d7.
//
// Solidity: function getZKPRequests(uint256 startIndex, uint256 length) view returns((string,address,bytes)[])
func (_UniversalVerifier *UniversalVerifierCallerSession) GetZKPRequests(startIndex *big.Int, length *big.Int) ([]IZKPVerifierZKPRequest, error) {
	return _UniversalVerifier.Contract.GetZKPRequests(&_UniversalVerifier.CallOpts, startIndex, length)
}

// GetZKPRequestsCount is a free data retrieval call binding the contract method 0x6508e1b4.
//
// Solidity: function getZKPRequestsCount() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCaller) GetZKPRequestsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getZKPRequestsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetZKPRequestsCount is a free data retrieval call binding the contract method 0x6508e1b4.
//
// Solidity: function getZKPRequestsCount() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierSession) GetZKPRequestsCount() (*big.Int, error) {
	return _UniversalVerifier.Contract.GetZKPRequestsCount(&_UniversalVerifier.CallOpts)
}

// GetZKPRequestsCount is a free data retrieval call binding the contract method 0x6508e1b4.
//
// Solidity: function getZKPRequestsCount() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetZKPRequestsCount() (*big.Int, error) {
	return _UniversalVerifier.Contract.GetZKPRequestsCount(&_UniversalVerifier.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UniversalVerifier.Contract.HasRole(&_UniversalVerifier.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UniversalVerifier.Contract.HasRole(&_UniversalVerifier.CallOpts, role, account)
}

// IsProofVerified is a free data retrieval call binding the contract method 0x49555fb1.
//
// Solidity: function isProofVerified(address sender, uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) IsProofVerified(opts *bind.CallOpts, sender common.Address, requestId uint64) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "isProofVerified", sender, requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProofVerified is a free data retrieval call binding the contract method 0x49555fb1.
//
// Solidity: function isProofVerified(address sender, uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) IsProofVerified(sender common.Address, requestId uint64) (bool, error) {
	return _UniversalVerifier.Contract.IsProofVerified(&_UniversalVerifier.CallOpts, sender, requestId)
}

// IsProofVerified is a free data retrieval call binding the contract method 0x49555fb1.
//
// Solidity: function isProofVerified(address sender, uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) IsProofVerified(sender common.Address, requestId uint64) (bool, error) {
	return _UniversalVerifier.Contract.IsProofVerified(&_UniversalVerifier.CallOpts, sender, requestId)
}

// IsWhitelistedValidator is a free data retrieval call binding the contract method 0x9b36d036.
//
// Solidity: function isWhitelistedValidator(address validator) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) IsWhitelistedValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "isWhitelistedValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedValidator is a free data retrieval call binding the contract method 0x9b36d036.
//
// Solidity: function isWhitelistedValidator(address validator) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) IsWhitelistedValidator(validator common.Address) (bool, error) {
	return _UniversalVerifier.Contract.IsWhitelistedValidator(&_UniversalVerifier.CallOpts, validator)
}

// IsWhitelistedValidator is a free data retrieval call binding the contract method 0x9b36d036.
//
// Solidity: function isWhitelistedValidator(address validator) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) IsWhitelistedValidator(validator common.Address) (bool, error) {
	return _UniversalVerifier.Contract.IsWhitelistedValidator(&_UniversalVerifier.CallOpts, validator)
}

// IsZKPRequestEnabled is a free data retrieval call binding the contract method 0x130a73ac.
//
// Solidity: function isZKPRequestEnabled(uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) IsZKPRequestEnabled(opts *bind.CallOpts, requestId uint64) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "isZKPRequestEnabled", requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZKPRequestEnabled is a free data retrieval call binding the contract method 0x130a73ac.
//
// Solidity: function isZKPRequestEnabled(uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) IsZKPRequestEnabled(requestId uint64) (bool, error) {
	return _UniversalVerifier.Contract.IsZKPRequestEnabled(&_UniversalVerifier.CallOpts, requestId)
}

// IsZKPRequestEnabled is a free data retrieval call binding the contract method 0x130a73ac.
//
// Solidity: function isZKPRequestEnabled(uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) IsZKPRequestEnabled(requestId uint64) (bool, error) {
	return _UniversalVerifier.Contract.IsZKPRequestEnabled(&_UniversalVerifier.CallOpts, requestId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniversalVerifier *UniversalVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniversalVerifier *UniversalVerifierSession) Owner() (common.Address, error) {
	return _UniversalVerifier.Contract.Owner(&_UniversalVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniversalVerifier *UniversalVerifierCallerSession) Owner() (common.Address, error) {
	return _UniversalVerifier.Contract.Owner(&_UniversalVerifier.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniversalVerifier *UniversalVerifierCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniversalVerifier *UniversalVerifierSession) PendingOwner() (common.Address, error) {
	return _UniversalVerifier.Contract.PendingOwner(&_UniversalVerifier.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniversalVerifier *UniversalVerifierCallerSession) PendingOwner() (common.Address, error) {
	return _UniversalVerifier.Contract.PendingOwner(&_UniversalVerifier.CallOpts)
}

// RequestIdExists is a free data retrieval call binding the contract method 0xab7bcfb7.
//
// Solidity: function requestIdExists(uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) RequestIdExists(opts *bind.CallOpts, requestId uint64) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "requestIdExists", requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestIdExists is a free data retrieval call binding the contract method 0xab7bcfb7.
//
// Solidity: function requestIdExists(uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) RequestIdExists(requestId uint64) (bool, error) {
	return _UniversalVerifier.Contract.RequestIdExists(&_UniversalVerifier.CallOpts, requestId)
}

// RequestIdExists is a free data retrieval call binding the contract method 0xab7bcfb7.
//
// Solidity: function requestIdExists(uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) RequestIdExists(requestId uint64) (bool, error) {
	return _UniversalVerifier.Contract.RequestIdExists(&_UniversalVerifier.CallOpts, requestId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UniversalVerifier.Contract.SupportsInterface(&_UniversalVerifier.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UniversalVerifier.Contract.SupportsInterface(&_UniversalVerifier.CallOpts, interfaceId)
}

// VerifyLinkedProofs is a free data retrieval call binding the contract method 0x0106a959.
//
// Solidity: function verifyLinkedProofs(address sender, uint64[] requestIds) view returns()
func (_UniversalVerifier *UniversalVerifierCaller) VerifyLinkedProofs(opts *bind.CallOpts, sender common.Address, requestIds []uint64) error {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "verifyLinkedProofs", sender, requestIds)

	if err != nil {
		return err
	}

	return err

}

// VerifyLinkedProofs is a free data retrieval call binding the contract method 0x0106a959.
//
// Solidity: function verifyLinkedProofs(address sender, uint64[] requestIds) view returns()
func (_UniversalVerifier *UniversalVerifierSession) VerifyLinkedProofs(sender common.Address, requestIds []uint64) error {
	return _UniversalVerifier.Contract.VerifyLinkedProofs(&_UniversalVerifier.CallOpts, sender, requestIds)
}

// VerifyLinkedProofs is a free data retrieval call binding the contract method 0x0106a959.
//
// Solidity: function verifyLinkedProofs(address sender, uint64[] requestIds) view returns()
func (_UniversalVerifier *UniversalVerifierCallerSession) VerifyLinkedProofs(sender common.Address, requestIds []uint64) error {
	return _UniversalVerifier.Contract.VerifyLinkedProofs(&_UniversalVerifier.CallOpts, sender, requestIds)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_UniversalVerifier *UniversalVerifierCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_UniversalVerifier *UniversalVerifierSession) Version() (string, error) {
	return _UniversalVerifier.Contract.Version(&_UniversalVerifier.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_UniversalVerifier *UniversalVerifierCallerSession) Version() (string, error) {
	return _UniversalVerifier.Contract.Version(&_UniversalVerifier.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniversalVerifier *UniversalVerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniversalVerifier *UniversalVerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _UniversalVerifier.Contract.AcceptOwnership(&_UniversalVerifier.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _UniversalVerifier.Contract.AcceptOwnership(&_UniversalVerifier.TransactOpts)
}

// AddValidatorToWhitelist is a paid mutator transaction binding the contract method 0x59f871a1.
//
// Solidity: function addValidatorToWhitelist(address validator) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) AddValidatorToWhitelist(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "addValidatorToWhitelist", validator)
}

// AddValidatorToWhitelist is a paid mutator transaction binding the contract method 0x59f871a1.
//
// Solidity: function addValidatorToWhitelist(address validator) returns()
func (_UniversalVerifier *UniversalVerifierSession) AddValidatorToWhitelist(validator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.AddValidatorToWhitelist(&_UniversalVerifier.TransactOpts, validator)
}

// AddValidatorToWhitelist is a paid mutator transaction binding the contract method 0x59f871a1.
//
// Solidity: function addValidatorToWhitelist(address validator) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) AddValidatorToWhitelist(validator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.AddValidatorToWhitelist(&_UniversalVerifier.TransactOpts, validator)
}

// CrossChainUpdateStatus is a paid mutator transaction binding the contract method 0x862b304c.
//
// Solidity: function crossChainUpdateStatus(address userAddress, uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) CrossChainUpdateStatus(opts *bind.TransactOpts, userAddress common.Address, requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "crossChainUpdateStatus", userAddress, requestId)
}

// CrossChainUpdateStatus is a paid mutator transaction binding the contract method 0x862b304c.
//
// Solidity: function crossChainUpdateStatus(address userAddress, uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierSession) CrossChainUpdateStatus(userAddress common.Address, requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.CrossChainUpdateStatus(&_UniversalVerifier.TransactOpts, userAddress, requestId)
}

// CrossChainUpdateStatus is a paid mutator transaction binding the contract method 0x862b304c.
//
// Solidity: function crossChainUpdateStatus(address userAddress, uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) CrossChainUpdateStatus(userAddress common.Address, requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.CrossChainUpdateStatus(&_UniversalVerifier.TransactOpts, userAddress, requestId)
}

// CrossUploadStatus is a paid mutator transaction binding the contract method 0xb1eeab4f.
//
// Solidity: function crossUploadStatus(address userAddress, uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) CrossUploadStatus(opts *bind.TransactOpts, userAddress common.Address, requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "crossUploadStatus", userAddress, requestId)
}

// CrossUploadStatus is a paid mutator transaction binding the contract method 0xb1eeab4f.
//
// Solidity: function crossUploadStatus(address userAddress, uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierSession) CrossUploadStatus(userAddress common.Address, requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.CrossUploadStatus(&_UniversalVerifier.TransactOpts, userAddress, requestId)
}

// CrossUploadStatus is a paid mutator transaction binding the contract method 0xb1eeab4f.
//
// Solidity: function crossUploadStatus(address userAddress, uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) CrossUploadStatus(userAddress common.Address, requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.CrossUploadStatus(&_UniversalVerifier.TransactOpts, userAddress, requestId)
}

// DisableZKPRequest is a paid mutator transaction binding the contract method 0x82aff29f.
//
// Solidity: function disableZKPRequest(uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) DisableZKPRequest(opts *bind.TransactOpts, requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "disableZKPRequest", requestId)
}

// DisableZKPRequest is a paid mutator transaction binding the contract method 0x82aff29f.
//
// Solidity: function disableZKPRequest(uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierSession) DisableZKPRequest(requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.DisableZKPRequest(&_UniversalVerifier.TransactOpts, requestId)
}

// DisableZKPRequest is a paid mutator transaction binding the contract method 0x82aff29f.
//
// Solidity: function disableZKPRequest(uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) DisableZKPRequest(requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.DisableZKPRequest(&_UniversalVerifier.TransactOpts, requestId)
}

// EmitCrossChainEvent is a paid mutator transaction binding the contract method 0x8c938ce4.
//
// Solidity: function emitCrossChainEvent(uint64 requestId, address userAddr, uint32 srcChainId, uint32 chainId) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) EmitCrossChainEvent(opts *bind.TransactOpts, requestId uint64, userAddr common.Address, srcChainId uint32, chainId uint32) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "emitCrossChainEvent", requestId, userAddr, srcChainId, chainId)
}

// EmitCrossChainEvent is a paid mutator transaction binding the contract method 0x8c938ce4.
//
// Solidity: function emitCrossChainEvent(uint64 requestId, address userAddr, uint32 srcChainId, uint32 chainId) returns()
func (_UniversalVerifier *UniversalVerifierSession) EmitCrossChainEvent(requestId uint64, userAddr common.Address, srcChainId uint32, chainId uint32) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.EmitCrossChainEvent(&_UniversalVerifier.TransactOpts, requestId, userAddr, srcChainId, chainId)
}

// EmitCrossChainEvent is a paid mutator transaction binding the contract method 0x8c938ce4.
//
// Solidity: function emitCrossChainEvent(uint64 requestId, address userAddr, uint32 srcChainId, uint32 chainId) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) EmitCrossChainEvent(requestId uint64, userAddr common.Address, srcChainId uint32, chainId uint32) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.EmitCrossChainEvent(&_UniversalVerifier.TransactOpts, requestId, userAddr, srcChainId, chainId)
}

// EnableZKPRequest is a paid mutator transaction binding the contract method 0x30def4ac.
//
// Solidity: function enableZKPRequest(uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) EnableZKPRequest(opts *bind.TransactOpts, requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "enableZKPRequest", requestId)
}

// EnableZKPRequest is a paid mutator transaction binding the contract method 0x30def4ac.
//
// Solidity: function enableZKPRequest(uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierSession) EnableZKPRequest(requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.EnableZKPRequest(&_UniversalVerifier.TransactOpts, requestId)
}

// EnableZKPRequest is a paid mutator transaction binding the contract method 0x30def4ac.
//
// Solidity: function enableZKPRequest(uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) EnableZKPRequest(requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.EnableZKPRequest(&_UniversalVerifier.TransactOpts, requestId)
}

// GrantOperator is a paid mutator transaction binding the contract method 0xe348da13.
//
// Solidity: function grantOperator(address operator) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) GrantOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "grantOperator", operator)
}

// GrantOperator is a paid mutator transaction binding the contract method 0xe348da13.
//
// Solidity: function grantOperator(address operator) returns()
func (_UniversalVerifier *UniversalVerifierSession) GrantOperator(operator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.GrantOperator(&_UniversalVerifier.TransactOpts, operator)
}

// GrantOperator is a paid mutator transaction binding the contract method 0xe348da13.
//
// Solidity: function grantOperator(address operator) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) GrantOperator(operator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.GrantOperator(&_UniversalVerifier.TransactOpts, operator)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UniversalVerifier *UniversalVerifierSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.GrantRole(&_UniversalVerifier.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.GrantRole(&_UniversalVerifier.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address state, address owner) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) Initialize(opts *bind.TransactOpts, state common.Address, owner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "initialize", state, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address state, address owner) returns()
func (_UniversalVerifier *UniversalVerifierSession) Initialize(state common.Address, owner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.Initialize(&_UniversalVerifier.TransactOpts, state, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address state, address owner) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) Initialize(state common.Address, owner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.Initialize(&_UniversalVerifier.TransactOpts, state, owner)
}

// RemoveValidatorFromWhitelist is a paid mutator transaction binding the contract method 0x617879fe.
//
// Solidity: function removeValidatorFromWhitelist(address validator) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) RemoveValidatorFromWhitelist(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "removeValidatorFromWhitelist", validator)
}

// RemoveValidatorFromWhitelist is a paid mutator transaction binding the contract method 0x617879fe.
//
// Solidity: function removeValidatorFromWhitelist(address validator) returns()
func (_UniversalVerifier *UniversalVerifierSession) RemoveValidatorFromWhitelist(validator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.RemoveValidatorFromWhitelist(&_UniversalVerifier.TransactOpts, validator)
}

// RemoveValidatorFromWhitelist is a paid mutator transaction binding the contract method 0x617879fe.
//
// Solidity: function removeValidatorFromWhitelist(address validator) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) RemoveValidatorFromWhitelist(validator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.RemoveValidatorFromWhitelist(&_UniversalVerifier.TransactOpts, validator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniversalVerifier *UniversalVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniversalVerifier *UniversalVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniversalVerifier.Contract.RenounceOwnership(&_UniversalVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniversalVerifier.Contract.RenounceOwnership(&_UniversalVerifier.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UniversalVerifier *UniversalVerifierSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.RenounceRole(&_UniversalVerifier.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.RenounceRole(&_UniversalVerifier.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UniversalVerifier *UniversalVerifierSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.RevokeRole(&_UniversalVerifier.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.RevokeRole(&_UniversalVerifier.TransactOpts, role, account)
}

// SetRequestOwner is a paid mutator transaction binding the contract method 0xc9c23ea5.
//
// Solidity: function setRequestOwner(uint64 requestId, address requestOwner) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetRequestOwner(opts *bind.TransactOpts, requestId uint64, requestOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setRequestOwner", requestId, requestOwner)
}

// SetRequestOwner is a paid mutator transaction binding the contract method 0xc9c23ea5.
//
// Solidity: function setRequestOwner(uint64 requestId, address requestOwner) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetRequestOwner(requestId uint64, requestOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetRequestOwner(&_UniversalVerifier.TransactOpts, requestId, requestOwner)
}

// SetRequestOwner is a paid mutator transaction binding the contract method 0xc9c23ea5.
//
// Solidity: function setRequestOwner(uint64 requestId, address requestOwner) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetRequestOwner(requestId uint64, requestOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetRequestOwner(&_UniversalVerifier.TransactOpts, requestId, requestOwner)
}

// SetState is a paid mutator transaction binding the contract method 0x34c901af.
//
// Solidity: function setState(address state) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetState(opts *bind.TransactOpts, state common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setState", state)
}

// SetState is a paid mutator transaction binding the contract method 0x34c901af.
//
// Solidity: function setState(address state) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetState(state common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetState(&_UniversalVerifier.TransactOpts, state)
}

// SetState is a paid mutator transaction binding the contract method 0x34c901af.
//
// Solidity: function setState(address state) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetState(state common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetState(&_UniversalVerifier.TransactOpts, state)
}

// SetZKPRequest is a paid mutator transaction binding the contract method 0x9f5223e0.
//
// Solidity: function setZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetZKPRequest(opts *bind.TransactOpts, requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setZKPRequest", requestId, request)
}

// SetZKPRequest is a paid mutator transaction binding the contract method 0x9f5223e0.
//
// Solidity: function setZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetZKPRequest(requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetZKPRequest(&_UniversalVerifier.TransactOpts, requestId, request)
}

// SetZKPRequest is a paid mutator transaction binding the contract method 0x9f5223e0.
//
// Solidity: function setZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetZKPRequest(requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetZKPRequest(&_UniversalVerifier.TransactOpts, requestId, request)
}

// SetZKPRequests is a paid mutator transaction binding the contract method 0x6882ee95.
//
// Solidity: function setZKPRequests(uint64[] requestIds, (string,address,bytes)[] requests) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetZKPRequests(opts *bind.TransactOpts, requestIds []uint64, requests []IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setZKPRequests", requestIds, requests)
}

// SetZKPRequests is a paid mutator transaction binding the contract method 0x6882ee95.
//
// Solidity: function setZKPRequests(uint64[] requestIds, (string,address,bytes)[] requests) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetZKPRequests(requestIds []uint64, requests []IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetZKPRequests(&_UniversalVerifier.TransactOpts, requestIds, requests)
}

// SetZKPRequests is a paid mutator transaction binding the contract method 0x6882ee95.
//
// Solidity: function setZKPRequests(uint64[] requestIds, (string,address,bytes)[] requests) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetZKPRequests(requestIds []uint64, requests []IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetZKPRequests(&_UniversalVerifier.TransactOpts, requestIds, requests)
}

// SubmitZKPResponse is a paid mutator transaction binding the contract method 0xb68967e2.
//
// Solidity: function submitZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SubmitZKPResponse(opts *bind.TransactOpts, requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "submitZKPResponse", requestId, inputs, a, b, c)
}

// SubmitZKPResponse is a paid mutator transaction binding the contract method 0xb68967e2.
//
// Solidity: function submitZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_UniversalVerifier *UniversalVerifierSession) SubmitZKPResponse(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponse(&_UniversalVerifier.TransactOpts, requestId, inputs, a, b, c)
}

// SubmitZKPResponse is a paid mutator transaction binding the contract method 0xb68967e2.
//
// Solidity: function submitZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SubmitZKPResponse(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponse(&_UniversalVerifier.TransactOpts, requestId, inputs, a, b, c)
}

// SubmitZKPResponseBatch is a paid mutator transaction binding the contract method 0xcde09fad.
//
// Solidity: function submitZKPResponseBatch((uint64,uint256[],uint256[2],uint256[2][2],uint256[2],address)[] response) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SubmitZKPResponseBatch(opts *bind.TransactOpts, response []IZKPVerifierSubmitZkpResponseStruct) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "submitZKPResponseBatch", response)
}

// SubmitZKPResponseBatch is a paid mutator transaction binding the contract method 0xcde09fad.
//
// Solidity: function submitZKPResponseBatch((uint64,uint256[],uint256[2],uint256[2][2],uint256[2],address)[] response) returns()
func (_UniversalVerifier *UniversalVerifierSession) SubmitZKPResponseBatch(response []IZKPVerifierSubmitZkpResponseStruct) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseBatch(&_UniversalVerifier.TransactOpts, response)
}

// SubmitZKPResponseBatch is a paid mutator transaction binding the contract method 0xcde09fad.
//
// Solidity: function submitZKPResponseBatch((uint64,uint256[],uint256[2],uint256[2][2],uint256[2],address)[] response) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SubmitZKPResponseBatch(response []IZKPVerifierSubmitZkpResponseStruct) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseBatch(&_UniversalVerifier.TransactOpts, response)
}

// SubmitZKPResponseBatchV2 is a paid mutator transaction binding the contract method 0xa527305e.
//
// Solidity: function submitZKPResponseBatchV2(((uint64,bytes,bytes)[],bytes)[] crossResponses) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SubmitZKPResponseBatchV2(opts *bind.TransactOpts, crossResponses []IZKPVerifierSubmitZkpResponseCross) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "submitZKPResponseBatchV2", crossResponses)
}

// SubmitZKPResponseBatchV2 is a paid mutator transaction binding the contract method 0xa527305e.
//
// Solidity: function submitZKPResponseBatchV2(((uint64,bytes,bytes)[],bytes)[] crossResponses) returns()
func (_UniversalVerifier *UniversalVerifierSession) SubmitZKPResponseBatchV2(crossResponses []IZKPVerifierSubmitZkpResponseCross) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseBatchV2(&_UniversalVerifier.TransactOpts, crossResponses)
}

// SubmitZKPResponseBatchV2 is a paid mutator transaction binding the contract method 0xa527305e.
//
// Solidity: function submitZKPResponseBatchV2(((uint64,bytes,bytes)[],bytes)[] crossResponses) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SubmitZKPResponseBatchV2(crossResponses []IZKPVerifierSubmitZkpResponseCross) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseBatchV2(&_UniversalVerifier.TransactOpts, crossResponses)
}

// SubmitZKPResponseBatchWithCrossChain is a paid mutator transaction binding the contract method 0xc613b408.
//
// Solidity: function submitZKPResponseBatchWithCrossChain((uint64,uint256[],uint256[2],uint256[2][2],uint256[2],address)[] response, uint32 srcChainId, uint32 chainId) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SubmitZKPResponseBatchWithCrossChain(opts *bind.TransactOpts, response []IZKPVerifierSubmitZkpResponseStruct, srcChainId uint32, chainId uint32) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "submitZKPResponseBatchWithCrossChain", response, srcChainId, chainId)
}

// SubmitZKPResponseBatchWithCrossChain is a paid mutator transaction binding the contract method 0xc613b408.
//
// Solidity: function submitZKPResponseBatchWithCrossChain((uint64,uint256[],uint256[2],uint256[2][2],uint256[2],address)[] response, uint32 srcChainId, uint32 chainId) returns()
func (_UniversalVerifier *UniversalVerifierSession) SubmitZKPResponseBatchWithCrossChain(response []IZKPVerifierSubmitZkpResponseStruct, srcChainId uint32, chainId uint32) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseBatchWithCrossChain(&_UniversalVerifier.TransactOpts, response, srcChainId, chainId)
}

// SubmitZKPResponseBatchWithCrossChain is a paid mutator transaction binding the contract method 0xc613b408.
//
// Solidity: function submitZKPResponseBatchWithCrossChain((uint64,uint256[],uint256[2],uint256[2][2],uint256[2],address)[] response, uint32 srcChainId, uint32 chainId) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SubmitZKPResponseBatchWithCrossChain(response []IZKPVerifierSubmitZkpResponseStruct, srcChainId uint32, chainId uint32) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseBatchWithCrossChain(&_UniversalVerifier.TransactOpts, response, srcChainId, chainId)
}

// SubmitZKPResponseV2 is a paid mutator transaction binding the contract method 0xade09fcd.
//
// Solidity: function submitZKPResponseV2((uint64,bytes,bytes)[] responses, bytes crossChainProof) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SubmitZKPResponseV2(opts *bind.TransactOpts, responses []IZKPVerifierZKPResponse, crossChainProof []byte) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "submitZKPResponseV2", responses, crossChainProof)
}

// SubmitZKPResponseV2 is a paid mutator transaction binding the contract method 0xade09fcd.
//
// Solidity: function submitZKPResponseV2((uint64,bytes,bytes)[] responses, bytes crossChainProof) returns()
func (_UniversalVerifier *UniversalVerifierSession) SubmitZKPResponseV2(responses []IZKPVerifierZKPResponse, crossChainProof []byte) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseV2(&_UniversalVerifier.TransactOpts, responses, crossChainProof)
}

// SubmitZKPResponseV2 is a paid mutator transaction binding the contract method 0xade09fcd.
//
// Solidity: function submitZKPResponseV2((uint64,bytes,bytes)[] responses, bytes crossChainProof) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SubmitZKPResponseV2(responses []IZKPVerifierZKPResponse, crossChainProof []byte) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseV2(&_UniversalVerifier.TransactOpts, responses, crossChainProof)
}

// SubmitZKPResponseWithUser is a paid mutator transaction binding the contract method 0x7559560e.
//
// Solidity: function submitZKPResponseWithUser(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c, address userAddr) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SubmitZKPResponseWithUser(opts *bind.TransactOpts, requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, userAddr common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "submitZKPResponseWithUser", requestId, inputs, a, b, c, userAddr)
}

// SubmitZKPResponseWithUser is a paid mutator transaction binding the contract method 0x7559560e.
//
// Solidity: function submitZKPResponseWithUser(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c, address userAddr) returns()
func (_UniversalVerifier *UniversalVerifierSession) SubmitZKPResponseWithUser(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, userAddr common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseWithUser(&_UniversalVerifier.TransactOpts, requestId, inputs, a, b, c, userAddr)
}

// SubmitZKPResponseWithUser is a paid mutator transaction binding the contract method 0x7559560e.
//
// Solidity: function submitZKPResponseWithUser(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c, address userAddr) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SubmitZKPResponseWithUser(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, userAddr common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseWithUser(&_UniversalVerifier.TransactOpts, requestId, inputs, a, b, c, userAddr)
}

// SubmitZKPResponseWithUserCrossChain is a paid mutator transaction binding the contract method 0x7fd7684c.
//
// Solidity: function submitZKPResponseWithUserCrossChain(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c, address userAddr, uint32 srcChainId, uint32 chainId) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SubmitZKPResponseWithUserCrossChain(opts *bind.TransactOpts, requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, userAddr common.Address, srcChainId uint32, chainId uint32) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "submitZKPResponseWithUserCrossChain", requestId, inputs, a, b, c, userAddr, srcChainId, chainId)
}

// SubmitZKPResponseWithUserCrossChain is a paid mutator transaction binding the contract method 0x7fd7684c.
//
// Solidity: function submitZKPResponseWithUserCrossChain(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c, address userAddr, uint32 srcChainId, uint32 chainId) returns()
func (_UniversalVerifier *UniversalVerifierSession) SubmitZKPResponseWithUserCrossChain(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, userAddr common.Address, srcChainId uint32, chainId uint32) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseWithUserCrossChain(&_UniversalVerifier.TransactOpts, requestId, inputs, a, b, c, userAddr, srcChainId, chainId)
}

// SubmitZKPResponseWithUserCrossChain is a paid mutator transaction binding the contract method 0x7fd7684c.
//
// Solidity: function submitZKPResponseWithUserCrossChain(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c, address userAddr, uint32 srcChainId, uint32 chainId) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SubmitZKPResponseWithUserCrossChain(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, userAddr common.Address, srcChainId uint32, chainId uint32) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseWithUserCrossChain(&_UniversalVerifier.TransactOpts, requestId, inputs, a, b, c, userAddr, srcChainId, chainId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniversalVerifier *UniversalVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.TransferOwnership(&_UniversalVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.TransferOwnership(&_UniversalVerifier.TransactOpts, newOwner)
}

// UpdateZKPRequest is a paid mutator transaction binding the contract method 0x011f5bd6.
//
// Solidity: function updateZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) UpdateZKPRequest(opts *bind.TransactOpts, requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "updateZKPRequest", requestId, request)
}

// UpdateZKPRequest is a paid mutator transaction binding the contract method 0x011f5bd6.
//
// Solidity: function updateZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_UniversalVerifier *UniversalVerifierSession) UpdateZKPRequest(requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.UpdateZKPRequest(&_UniversalVerifier.TransactOpts, requestId, request)
}

// UpdateZKPRequest is a paid mutator transaction binding the contract method 0x011f5bd6.
//
// Solidity: function updateZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) UpdateZKPRequest(requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.UpdateZKPRequest(&_UniversalVerifier.TransactOpts, requestId, request)
}

// VerifyZKPResponse is a paid mutator transaction binding the contract method 0x5176983b.
//
// Solidity: function verifyZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c, address sender) returns((string,uint256)[])
func (_UniversalVerifier *UniversalVerifierTransactor) VerifyZKPResponse(opts *bind.TransactOpts, requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, sender common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "verifyZKPResponse", requestId, inputs, a, b, c, sender)
}

// VerifyZKPResponse is a paid mutator transaction binding the contract method 0x5176983b.
//
// Solidity: function verifyZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c, address sender) returns((string,uint256)[])
func (_UniversalVerifier *UniversalVerifierSession) VerifyZKPResponse(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, sender common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.VerifyZKPResponse(&_UniversalVerifier.TransactOpts, requestId, inputs, a, b, c, sender)
}

// VerifyZKPResponse is a paid mutator transaction binding the contract method 0x5176983b.
//
// Solidity: function verifyZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c, address sender) returns((string,uint256)[])
func (_UniversalVerifier *UniversalVerifierTransactorSession) VerifyZKPResponse(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, sender common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.VerifyZKPResponse(&_UniversalVerifier.TransactOpts, requestId, inputs, a, b, c, sender)
}

// UniversalVerifierErrorHandleIterator is returned from FilterErrorHandle and is used to iterate over the raw logs and unpacked data for ErrorHandle events raised by the UniversalVerifier contract.
type UniversalVerifierErrorHandleIterator struct {
	Event *UniversalVerifierErrorHandle // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierErrorHandleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierErrorHandle)
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
		it.Event = new(UniversalVerifierErrorHandle)
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
func (it *UniversalVerifierErrorHandleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierErrorHandleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierErrorHandle represents a ErrorHandle event raised by the UniversalVerifier contract.
type UniversalVerifierErrorHandle struct {
	RequestId   uint64
	Useraddress common.Address
	Err         common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterErrorHandle is a free log retrieval operation binding the contract event 0x718b7cf68a5d91420c89b2a1db5b652dc19d8595087d9aeb080d235ef14a57c8.
//
// Solidity: event ErrorHandle(uint64 indexed requestId, address indexed useraddress, bytes indexed err)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterErrorHandle(opts *bind.FilterOpts, requestId []uint64, useraddress []common.Address, err [][]byte) (*UniversalVerifierErrorHandleIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var useraddressRule []interface{}
	for _, useraddressItem := range useraddress {
		useraddressRule = append(useraddressRule, useraddressItem)
	}
	var errRule []interface{}
	for _, errItem := range err {
		errRule = append(errRule, errItem)
	}

	logs, sub, errs := _UniversalVerifier.contract.FilterLogs(opts, "ErrorHandle", requestIdRule, useraddressRule, errRule)
	if errs != nil {
		return nil, errs
	}
	return &UniversalVerifierErrorHandleIterator{contract: _UniversalVerifier.contract, event: "ErrorHandle", logs: logs, sub: sub}, nil
}

// WatchErrorHandle is a free log subscription operation binding the contract event 0x718b7cf68a5d91420c89b2a1db5b652dc19d8595087d9aeb080d235ef14a57c8.
//
// Solidity: event ErrorHandle(uint64 indexed requestId, address indexed useraddress, bytes indexed err)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchErrorHandle(opts *bind.WatchOpts, sink chan<- *UniversalVerifierErrorHandle, requestId []uint64, useraddress []common.Address, err [][]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var useraddressRule []interface{}
	for _, useraddressItem := range useraddress {
		useraddressRule = append(useraddressRule, useraddressItem)
	}
	var errRule []interface{}
	for _, errItem := range err {
		errRule = append(errRule, errItem)
	}

	logs, sub, errs := _UniversalVerifier.contract.WatchLogs(opts, "ErrorHandle", requestIdRule, useraddressRule, errRule)
	if errs != nil {
		return nil, errs
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierErrorHandle)
				if err := _UniversalVerifier.contract.UnpackLog(event, "ErrorHandle", log); err != nil {
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

// ParseErrorHandle is a log parse operation binding the contract event 0x718b7cf68a5d91420c89b2a1db5b652dc19d8595087d9aeb080d235ef14a57c8.
//
// Solidity: event ErrorHandle(uint64 indexed requestId, address indexed useraddress, bytes indexed err)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseErrorHandle(log types.Log) (*UniversalVerifierErrorHandle, error) {
	event := new(UniversalVerifierErrorHandle)
	if err := _UniversalVerifier.contract.UnpackLog(event, "ErrorHandle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UniversalVerifier contract.
type UniversalVerifierInitializedIterator struct {
	Event *UniversalVerifierInitialized // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierInitialized)
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
		it.Event = new(UniversalVerifierInitialized)
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
func (it *UniversalVerifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierInitialized represents a Initialized event raised by the UniversalVerifier contract.
type UniversalVerifierInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*UniversalVerifierInitializedIterator, error) {

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierInitializedIterator{contract: _UniversalVerifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UniversalVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierInitialized)
				if err := _UniversalVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseInitialized(log types.Log) (*UniversalVerifierInitialized, error) {
	event := new(UniversalVerifierInitialized)
	if err := _UniversalVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the UniversalVerifier contract.
type UniversalVerifierOwnershipTransferStartedIterator struct {
	Event *UniversalVerifierOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierOwnershipTransferStarted)
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
		it.Event = new(UniversalVerifierOwnershipTransferStarted)
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
func (it *UniversalVerifierOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the UniversalVerifier contract.
type UniversalVerifierOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniversalVerifierOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierOwnershipTransferStartedIterator{contract: _UniversalVerifier.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *UniversalVerifierOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierOwnershipTransferStarted)
				if err := _UniversalVerifier.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseOwnershipTransferStarted(log types.Log) (*UniversalVerifierOwnershipTransferStarted, error) {
	event := new(UniversalVerifierOwnershipTransferStarted)
	if err := _UniversalVerifier.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UniversalVerifier contract.
type UniversalVerifierOwnershipTransferredIterator struct {
	Event *UniversalVerifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierOwnershipTransferred)
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
		it.Event = new(UniversalVerifierOwnershipTransferred)
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
func (it *UniversalVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the UniversalVerifier contract.
type UniversalVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniversalVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierOwnershipTransferredIterator{contract: _UniversalVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UniversalVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierOwnershipTransferred)
				if err := _UniversalVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UniversalVerifier *UniversalVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*UniversalVerifierOwnershipTransferred, error) {
	event := new(UniversalVerifierOwnershipTransferred)
	if err := _UniversalVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the UniversalVerifier contract.
type UniversalVerifierRoleAdminChangedIterator struct {
	Event *UniversalVerifierRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierRoleAdminChanged)
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
		it.Event = new(UniversalVerifierRoleAdminChanged)
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
func (it *UniversalVerifierRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierRoleAdminChanged represents a RoleAdminChanged event raised by the UniversalVerifier contract.
type UniversalVerifierRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*UniversalVerifierRoleAdminChangedIterator, error) {

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

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierRoleAdminChangedIterator{contract: _UniversalVerifier.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *UniversalVerifierRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierRoleAdminChanged)
				if err := _UniversalVerifier.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_UniversalVerifier *UniversalVerifierFilterer) ParseRoleAdminChanged(log types.Log) (*UniversalVerifierRoleAdminChanged, error) {
	event := new(UniversalVerifierRoleAdminChanged)
	if err := _UniversalVerifier.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the UniversalVerifier contract.
type UniversalVerifierRoleGrantedIterator struct {
	Event *UniversalVerifierRoleGranted // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierRoleGranted)
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
		it.Event = new(UniversalVerifierRoleGranted)
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
func (it *UniversalVerifierRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierRoleGranted represents a RoleGranted event raised by the UniversalVerifier contract.
type UniversalVerifierRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UniversalVerifierRoleGrantedIterator, error) {

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

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierRoleGrantedIterator{contract: _UniversalVerifier.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *UniversalVerifierRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierRoleGranted)
				if err := _UniversalVerifier.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_UniversalVerifier *UniversalVerifierFilterer) ParseRoleGranted(log types.Log) (*UniversalVerifierRoleGranted, error) {
	event := new(UniversalVerifierRoleGranted)
	if err := _UniversalVerifier.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the UniversalVerifier contract.
type UniversalVerifierRoleRevokedIterator struct {
	Event *UniversalVerifierRoleRevoked // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierRoleRevoked)
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
		it.Event = new(UniversalVerifierRoleRevoked)
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
func (it *UniversalVerifierRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierRoleRevoked represents a RoleRevoked event raised by the UniversalVerifier contract.
type UniversalVerifierRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UniversalVerifierRoleRevokedIterator, error) {

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

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierRoleRevokedIterator{contract: _UniversalVerifier.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *UniversalVerifierRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierRoleRevoked)
				if err := _UniversalVerifier.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_UniversalVerifier *UniversalVerifierFilterer) ParseRoleRevoked(log types.Log) (*UniversalVerifierRoleRevoked, error) {
	event := new(UniversalVerifierRoleRevoked)
	if err := _UniversalVerifier.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierVCCrossChainPackageIterator is returned from FilterVCCrossChainPackage and is used to iterate over the raw logs and unpacked data for VCCrossChainPackage events raised by the UniversalVerifier contract.
type UniversalVerifierVCCrossChainPackageIterator struct {
	Event *UniversalVerifierVCCrossChainPackage // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierVCCrossChainPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierVCCrossChainPackage)
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
		it.Event = new(UniversalVerifierVCCrossChainPackage)
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
func (it *UniversalVerifierVCCrossChainPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierVCCrossChainPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierVCCrossChainPackage represents a VCCrossChainPackage event raised by the UniversalVerifier contract.
type UniversalVerifierVCCrossChainPackage struct {
	SrcChainId     uint32
	DestChainId    uint32
	ChannelId      uint32
	Sequence       *big.Int
	CrossDataBytes []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVCCrossChainPackage is a free log retrieval operation binding the contract event 0x2f4796132af44fdc7d0bfd7fe25cd97e2b6e8981ca28b81dfd6e22f065a01c50.
//
// Solidity: event VCCrossChainPackage(uint32 srcChainId, uint32 indexed destChainId, uint32 indexed channelId, uint256 indexed sequence, bytes crossDataBytes)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterVCCrossChainPackage(opts *bind.FilterOpts, destChainId []uint32, channelId []uint32, sequence []*big.Int) (*UniversalVerifierVCCrossChainPackageIterator, error) {

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

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "VCCrossChainPackage", destChainIdRule, channelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierVCCrossChainPackageIterator{contract: _UniversalVerifier.contract, event: "VCCrossChainPackage", logs: logs, sub: sub}, nil
}

// WatchVCCrossChainPackage is a free log subscription operation binding the contract event 0x2f4796132af44fdc7d0bfd7fe25cd97e2b6e8981ca28b81dfd6e22f065a01c50.
//
// Solidity: event VCCrossChainPackage(uint32 srcChainId, uint32 indexed destChainId, uint32 indexed channelId, uint256 indexed sequence, bytes crossDataBytes)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchVCCrossChainPackage(opts *bind.WatchOpts, sink chan<- *UniversalVerifierVCCrossChainPackage, destChainId []uint32, channelId []uint32, sequence []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "VCCrossChainPackage", destChainIdRule, channelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierVCCrossChainPackage)
				if err := _UniversalVerifier.contract.UnpackLog(event, "VCCrossChainPackage", log); err != nil {
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

// ParseVCCrossChainPackage is a log parse operation binding the contract event 0x2f4796132af44fdc7d0bfd7fe25cd97e2b6e8981ca28b81dfd6e22f065a01c50.
//
// Solidity: event VCCrossChainPackage(uint32 srcChainId, uint32 indexed destChainId, uint32 indexed channelId, uint256 indexed sequence, bytes crossDataBytes)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseVCCrossChainPackage(log types.Log) (*UniversalVerifierVCCrossChainPackage, error) {
	event := new(UniversalVerifierVCCrossChainPackage)
	if err := _UniversalVerifier.contract.UnpackLog(event, "VCCrossChainPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierZKPRequestSetIterator is returned from FilterZKPRequestSet and is used to iterate over the raw logs and unpacked data for ZKPRequestSet events raised by the UniversalVerifier contract.
type UniversalVerifierZKPRequestSetIterator struct {
	Event *UniversalVerifierZKPRequestSet // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierZKPRequestSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierZKPRequestSet)
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
		it.Event = new(UniversalVerifierZKPRequestSet)
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
func (it *UniversalVerifierZKPRequestSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierZKPRequestSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierZKPRequestSet represents a ZKPRequestSet event raised by the UniversalVerifier contract.
type UniversalVerifierZKPRequestSet struct {
	RequestId    uint64
	RequestOwner common.Address
	Metadata     string
	Validator    common.Address
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterZKPRequestSet is a free log retrieval operation binding the contract event 0x26db2c5d21d517fa8c11af3ae555af7f4a560b9f70a5a32e73e6cd94d0cae3e8.
//
// Solidity: event ZKPRequestSet(uint64 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes data)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterZKPRequestSet(opts *bind.FilterOpts, requestId []uint64, requestOwner []common.Address) (*UniversalVerifierZKPRequestSetIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestOwnerRule []interface{}
	for _, requestOwnerItem := range requestOwner {
		requestOwnerRule = append(requestOwnerRule, requestOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "ZKPRequestSet", requestIdRule, requestOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierZKPRequestSetIterator{contract: _UniversalVerifier.contract, event: "ZKPRequestSet", logs: logs, sub: sub}, nil
}

// WatchZKPRequestSet is a free log subscription operation binding the contract event 0x26db2c5d21d517fa8c11af3ae555af7f4a560b9f70a5a32e73e6cd94d0cae3e8.
//
// Solidity: event ZKPRequestSet(uint64 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes data)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchZKPRequestSet(opts *bind.WatchOpts, sink chan<- *UniversalVerifierZKPRequestSet, requestId []uint64, requestOwner []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestOwnerRule []interface{}
	for _, requestOwnerItem := range requestOwner {
		requestOwnerRule = append(requestOwnerRule, requestOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "ZKPRequestSet", requestIdRule, requestOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierZKPRequestSet)
				if err := _UniversalVerifier.contract.UnpackLog(event, "ZKPRequestSet", log); err != nil {
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

// ParseZKPRequestSet is a log parse operation binding the contract event 0x26db2c5d21d517fa8c11af3ae555af7f4a560b9f70a5a32e73e6cd94d0cae3e8.
//
// Solidity: event ZKPRequestSet(uint64 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes data)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseZKPRequestSet(log types.Log) (*UniversalVerifierZKPRequestSet, error) {
	event := new(UniversalVerifierZKPRequestSet)
	if err := _UniversalVerifier.contract.UnpackLog(event, "ZKPRequestSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierZKPRequestUpdateIterator is returned from FilterZKPRequestUpdate and is used to iterate over the raw logs and unpacked data for ZKPRequestUpdate events raised by the UniversalVerifier contract.
type UniversalVerifierZKPRequestUpdateIterator struct {
	Event *UniversalVerifierZKPRequestUpdate // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierZKPRequestUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierZKPRequestUpdate)
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
		it.Event = new(UniversalVerifierZKPRequestUpdate)
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
func (it *UniversalVerifierZKPRequestUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierZKPRequestUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierZKPRequestUpdate represents a ZKPRequestUpdate event raised by the UniversalVerifier contract.
type UniversalVerifierZKPRequestUpdate struct {
	RequestId    uint64
	RequestOwner common.Address
	Metadata     string
	Validator    common.Address
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterZKPRequestUpdate is a free log retrieval operation binding the contract event 0xa2c61fcbd9637e91178d0dea7f9b5cce13f60c453603b3dc056b8f01bb3d4cb0.
//
// Solidity: event ZKPRequestUpdate(uint64 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes data)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterZKPRequestUpdate(opts *bind.FilterOpts, requestId []uint64, requestOwner []common.Address) (*UniversalVerifierZKPRequestUpdateIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestOwnerRule []interface{}
	for _, requestOwnerItem := range requestOwner {
		requestOwnerRule = append(requestOwnerRule, requestOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "ZKPRequestUpdate", requestIdRule, requestOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierZKPRequestUpdateIterator{contract: _UniversalVerifier.contract, event: "ZKPRequestUpdate", logs: logs, sub: sub}, nil
}

// WatchZKPRequestUpdate is a free log subscription operation binding the contract event 0xa2c61fcbd9637e91178d0dea7f9b5cce13f60c453603b3dc056b8f01bb3d4cb0.
//
// Solidity: event ZKPRequestUpdate(uint64 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes data)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchZKPRequestUpdate(opts *bind.WatchOpts, sink chan<- *UniversalVerifierZKPRequestUpdate, requestId []uint64, requestOwner []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestOwnerRule []interface{}
	for _, requestOwnerItem := range requestOwner {
		requestOwnerRule = append(requestOwnerRule, requestOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "ZKPRequestUpdate", requestIdRule, requestOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierZKPRequestUpdate)
				if err := _UniversalVerifier.contract.UnpackLog(event, "ZKPRequestUpdate", log); err != nil {
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

// ParseZKPRequestUpdate is a log parse operation binding the contract event 0xa2c61fcbd9637e91178d0dea7f9b5cce13f60c453603b3dc056b8f01bb3d4cb0.
//
// Solidity: event ZKPRequestUpdate(uint64 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes data)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseZKPRequestUpdate(log types.Log) (*UniversalVerifierZKPRequestUpdate, error) {
	event := new(UniversalVerifierZKPRequestUpdate)
	if err := _UniversalVerifier.contract.UnpackLog(event, "ZKPRequestUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierZKPResponseSubmittedIterator is returned from FilterZKPResponseSubmitted and is used to iterate over the raw logs and unpacked data for ZKPResponseSubmitted events raised by the UniversalVerifier contract.
type UniversalVerifierZKPResponseSubmittedIterator struct {
	Event *UniversalVerifierZKPResponseSubmitted // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierZKPResponseSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierZKPResponseSubmitted)
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
		it.Event = new(UniversalVerifierZKPResponseSubmitted)
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
func (it *UniversalVerifierZKPResponseSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierZKPResponseSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierZKPResponseSubmitted represents a ZKPResponseSubmitted event raised by the UniversalVerifier contract.
type UniversalVerifierZKPResponseSubmitted struct {
	RequestId uint64
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZKPResponseSubmitted is a free log retrieval operation binding the contract event 0x6979bc9c3e552c05dd9859285f1ed7a172e52ef39e1dce9c720e5bf8d82c9f4c.
//
// Solidity: event ZKPResponseSubmitted(uint64 indexed requestId, address indexed caller)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterZKPResponseSubmitted(opts *bind.FilterOpts, requestId []uint64, caller []common.Address) (*UniversalVerifierZKPResponseSubmittedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "ZKPResponseSubmitted", requestIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierZKPResponseSubmittedIterator{contract: _UniversalVerifier.contract, event: "ZKPResponseSubmitted", logs: logs, sub: sub}, nil
}

// WatchZKPResponseSubmitted is a free log subscription operation binding the contract event 0x6979bc9c3e552c05dd9859285f1ed7a172e52ef39e1dce9c720e5bf8d82c9f4c.
//
// Solidity: event ZKPResponseSubmitted(uint64 indexed requestId, address indexed caller)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchZKPResponseSubmitted(opts *bind.WatchOpts, sink chan<- *UniversalVerifierZKPResponseSubmitted, requestId []uint64, caller []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "ZKPResponseSubmitted", requestIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierZKPResponseSubmitted)
				if err := _UniversalVerifier.contract.UnpackLog(event, "ZKPResponseSubmitted", log); err != nil {
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

// ParseZKPResponseSubmitted is a log parse operation binding the contract event 0x6979bc9c3e552c05dd9859285f1ed7a172e52ef39e1dce9c720e5bf8d82c9f4c.
//
// Solidity: event ZKPResponseSubmitted(uint64 indexed requestId, address indexed caller)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseZKPResponseSubmitted(log types.Log) (*UniversalVerifierZKPResponseSubmitted, error) {
	event := new(UniversalVerifierZKPResponseSubmitted)
	if err := _UniversalVerifier.contract.UnpackLog(event, "ZKPResponseSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

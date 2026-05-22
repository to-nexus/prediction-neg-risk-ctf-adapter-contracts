// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// Order is an auto generated low-level Go binding around an user-defined struct.

// OrderStatus is an auto generated low-level Go binding around an user-defined struct.
type OrderStatus struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}

// NegRiskCtfExchangeMetaData contains all meta data concerning the NegRiskCtfExchange contract.
var NegRiskCtfExchangeMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ctf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_negRiskAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxyFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_safeFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fillAmounts\",\"type\":\"uint256[]\"}],\"name\":\"fillOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"getComplement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"getConditionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCtf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isFilledOrCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"internalType\":\"structOrderStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPolyProxyFactoryImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPolyProxyWalletAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getSafeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSafeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSafeFactoryImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"isValidNonce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"takerOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"makerOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"takerFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"makerFillAmounts\",\"type\":\"uint256[]\"}],\"name\":\"matchOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFilledOrCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parentCollectionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complement\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"complement\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOperatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"safeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newProxyFactory\",\"type\":\"address\"}],\"name\":\"setProxyFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSafeFactory\",\"type\":\"address\"}],\"name\":\"setSafeFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complement\",\"type\":\"uint256\"}],\"name\":\"validateComplement\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"validateOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"validateOrderSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"validateTokenId\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOperatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"takerOrderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"takerOrderMaker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmountFilled\",\"type\":\"uint256\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProxyFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProxyFactory\",\"type\":\"address\"}],\"name\":\"ProxyFactoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"RemovedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSafeFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSafeFactory\",\"type\":\"address\"}],\"name\":\"SafeFactoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"token0\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"token1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"TradingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"TradingUnpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidComplement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingGtRemaining\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedTokenIds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCrossing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderFilledOrCancelled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLittleTokensReceived\",\"type\":\"error\"}]",
	ID:  "NegRiskCtfExchange",
	Bin: "0x6101a060405260016000556003805460ff191690553480156200002157600080fd5b5060405162004868380380620048688339810160408190526200004491620003aa565b604080518082018252601781527f506f6c796d61726b6574204354462045786368616e67650000000000000000006020808301918252835180850185526001808252603160f81b82840190815233600090815282855287812083905560028552879020919091558451909320815190932060e08490526101008190524660a081815287517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818701819052818a0188905260608201859052608082019390935230818301528851808203909201825260c001909752865196909301959095208a9589958995899587958795879587959492938d938d938793879390916080523060c05261012052505050506001600160a01b0382811661014081905290821661016081905260405163095ea7b360e01b81526004810191909152600019602482015263095ea7b3906044016020604051808303816000875af1158015620001af573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001d591906200041a565b50620001e391505062000339565b610180525050600680546001600160a01b03199081166001600160a01b03948516179091556007805490911691831691909117905560405163a22cb46560e01b81528a8216600482015260016024820152908b16965063a22cb465955060440193506200024f92505050565b600060405180830381600087803b1580156200026a57600080fd5b505af11580156200027f573d6000803e3d6000fd5b505060405163a22cb46560e01b8152306004820152600160248201526001600160a01b038716925063a22cb4659150604401600060405180830381600087803b158015620002cc57600080fd5b505af1158015620002e1573d6000803e3d6000fd5b50505050505050505062000445565b6040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b600060c0516001600160a01b0316306001600160a01b031614801562000360575060a05146145b156200036d575060805190565b620003886101205160e05161010051620002f060201b60201c565b905090565b80516001600160a01b0381168114620003a557600080fd5b919050565b600080600080600060a08688031215620003c357600080fd5b620003ce866200038d565b9450620003de602087016200038d565b9350620003ee604087016200038d565b9250620003fe606087016200038d565b91506200040e608087016200038d565b90509295509295909350565b6000602082840312156200042d57600080fd5b815180151581146200043e57600080fd5b9392505050565b60805160a05160c05160e051610100516101205161014051610160516101805161437a620004ee60003960006108970152600081816104c801528181612698015281816129450152818161359401526136c40152600081816105eb015281816125e3015281816128ed015281816135d0015261370001526000612258015260006122a701526000612282015260006121db015260006122050152600061222f015261437a6000f3fe608060405234801561001057600080fd5b50600436106103365760003560e01c806370480275116101b2578063d798eff6116100f9578063e60f0c05116100a2578063f698da251161007c578063f698da2514610892578063fa950b48146108b9578063fbddd751146108cc578063fe729aaf146108df57600080fd5b8063e60f0c0514610834578063edef7d8e14610847578063f23a6e611461085a57600080fd5b8063e03ac3d0116100d3578063e03ac3d014610806578063e2eec4051461080e578063e50e4f971461082157600080fd5b8063d798eff6146107bd578063d7fb272f146107d0578063d82da838146107f357600080fd5b8063a287bdf11161015b578063b28c51c011610135578063b28c51c01461073b578063bc197c8114610759578063c10f1a751461079d57600080fd5b8063a287bdf114610702578063a6dfcf8614610715578063ac8a584a1461072857600080fd5b806383b8a5ae1161018c57806383b8a5ae146106d45780639870d7fe146106dc578063a10f3dce146106ef57600080fd5b8063704802751461068357806375d7370a146106965780637ecebe00146106b457600080fd5b8063429b62e5116102815780635893253c1161022a578063627cdcb911610204578063627cdcb91461061c578063654f0ce41461062457806368c7450f146106375780636d70f7ae1461064a57600080fd5b80635893253c146105ad5780635c1548fb146105e95780635c975abb1461060f57600080fd5b8063456068d21161025b578063456068d21461052f57806346423aa7146105375780634a2a11f5146105a557600080fd5b8063429b62e5146104f457806344bea37e146105145780634544f0551461051c57600080fd5b80631785f53c116102e357806334600901116102bd57806334600901146104b35780633b521d78146104c65780633d6d3598146104ec57600080fd5b80631785f53c1461042257806324d7806c146104355780632dff692d1461046f57600080fd5b80631031e36e116103145780631031e36e146103ca578063131e7e1c146103d457806313e7c9d8146103f457600080fd5b806301ffc9a71461033b5780630647ee201461036357806306b9d6911461039d575b600080fd5b61034e610349366004613724565b6108f2565b60405190151581526020015b60405180910390f35b61034e610371366004613798565b73ffffffffffffffffffffffffffffffffffffffff919091166000908152600460205260409020541490565b6103a561098b565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161035a565b6103d2610a24565b005b6007546103a59073ffffffffffffffffffffffffffffffffffffffff1681565b6104146104023660046137c4565b60026020526000908152604090205481565b60405190815260200161035a565b6103d26104303660046137c4565b610a78565b61034e6104433660046137c4565b73ffffffffffffffffffffffffffffffffffffffff166000908152600160208190526040909120541490565b61049c61047d3660046137e1565b6008602052600090815260409020805460019091015460ff9091169082565b60408051921515835260208301919091520161035a565b6103d26104c13660046137e1565b610b15565b7f00000000000000000000000000000000000000000000000000000000000000006103a5565b6103d2610b5f565b6104146105023660046137c4565b60016020526000908152604090205481565b610414600081565b6103d261052a3660046137c4565b610be3565b6103d2610c36565b6105886105453660046137e1565b6040805180820190915260008082526020820152506000908152600860209081526040918290208251808401909352805460ff1615158352600101549082015290565b60408051825115158152602092830151928101929092520161035a565b6103e8610414565b6105d46105bb3660046137e1565b6005602052600090815260409020805460019091015482565b6040805192835260208301919091520161035a565b7f00000000000000000000000000000000000000000000000000000000000000006103a5565b60035461034e9060ff1681565b6103d2610c88565b6103d2610632366004613a3a565b610c92565b6103d2610645366004613a6f565b610cad565b61034e6106583660046137c4565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460011490565b6103d26106913660046137c4565b610d07565b60075473ffffffffffffffffffffffffffffffffffffffff166103a5565b6104146106c23660046137c4565b60046020526000908152604090205481565b6103d2610da7565b6103d26106ea3660046137c4565b610e2c565b6104146106fd3660046137e1565b610eca565b6103a56107103660046137c4565b610ee8565b6103d2610723366004613a3a565b610f14565b6103d26107363660046137c4565b610f1d565b60065473ffffffffffffffffffffffffffffffffffffffff166103a5565b61076c610767366004613b2a565b610fba565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200161035a565b6006546103a59073ffffffffffffffffffffffffffffffffffffffff1681565b6103d26107cb366004613c58565b610fe5565b6104146107de3660046137e1565b60009081526005602052604090206001015490565b6103d2610801366004613cbc565b6110f5565b6103a5611136565b6103d261081c366004613cde565b6111a6565b61041461082f366004613a3a565b6111fb565b6103d2610842366004613d1b565b611298565b6103a56108553660046137c4565b6113a6565b61076c610868366004613dad565b7ff23a6e610000000000000000000000000000000000000000000000000000000095945050505050565b6104147f000000000000000000000000000000000000000000000000000000000000000081565b6103d26108c7366004613e16565b6113d2565b6103d26108da3660046137c4565b611409565b6103d26108ed366004613e4b565b61145c565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f4e2312e000000000000000000000000000000000000000000000000000000000148061098557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b600654604080517faaf10f42000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163aaf10f429160048083019260209291908290030181865afa1580156109fb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1f9190613e90565b905090565b3360009081526001602081905260409091205414610a6e576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a7661155e565b565b3360009081526001602081905260409091205414610ac2576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116600081815260016020526040808220829055513392917f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e91a350565b6000818152600560205260408120549003610b5c576040517f3f6cc76800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b33600090815260026020526040902054600114610ba8576040517f7c214f0400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336000818152600260205260408082208290555182917ff7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c91a3565b3360009081526001602081905260409091205414610c2d576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b5c816115b6565b3360009081526001602081905260409091205414610c80576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a76611644565b610a766001611699565b6000610c9d826111fb565b9050610ca981836116c7565b5050565b3360009081526001602081905260409091205414610cf7576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d02838383611826565b505050565b3360009081526001602081905260409091205414610d51576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116600081815260016020819052604080832091909155513392917ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc91a350565b3360009081526001602081905260409091205414610df1576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336000818152600160205260408082208290555182917f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e91a3565b3360009081526001602081905260409091205414610e76576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811660008181526002602052604080822060019055513392917ff1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c91a350565b6000610ed582610b15565b5060009081526005602052604090205490565b600061098582610ef6611136565b60075473ffffffffffffffffffffffffffffffffffffffff16611982565b610b5c81611a80565b3360009081526001602081905260409091205414610f67576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116600081815260026020526040808220829055513392917ff7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c91a350565b7fbc197c81000000000000000000000000000000000000000000000000000000005b95945050505050565b600054600203611056576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f5245454e5452414e43590000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b6002600081815533815260209190915260409020546001146110a4576040517f7c214f0400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035460ff16156110e1576040517f9e87fac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110ec828233611b85565b50506001600055565b806110ff83610eca565b14610ca9576040517f66f8620a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600754604080517fa619486e000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163a619486e9160048083019260209291908290030181865afa1580156109fb573d6000803e3d6000fd5b6111c58160400151826020015184846101800151856101600151611bde565b610ca9576040517f8baa579f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006109857fa852566c4e14d00869b6db0220888a9090a13eccdaea03713ff0a3d27bf9767c836000015184602001518560400151866060015187608001518860a001518960c001518a60e001518b61010001518c61012001518d61014001518e610160015160405160200161127d9d9c9b9a99989796959493929190613ef0565b60405160208183030381529060405280519060200120611c3c565b600054600203611304576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f5245454e5452414e435900000000000000000000000000000000000000000000604482015260640161104d565b600260008181553381526020919091526040902054600114611352576040517f7c214f0400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035460ff161561138f576040517f9e87fac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61139b84848484611ca5565b505060016000555050565b6000610985826113b461098b565b60065473ffffffffffffffffffffffffffffffffffffffff16611e5c565b805160005b81811015610d02576114018382815181106113f4576113f4613f8e565b6020026020010151611a80565b6001016113d7565b3360009081526001602081905260409091205414611453576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b5c81611ebe565b6000546002036114c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f5245454e5452414e435900000000000000000000000000000000000000000000604482015260640161104d565b600260008181553381526020919091526040902054600114611516576040517f7c214f0400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035460ff1615611553576040517f9e87fac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110ec828233611f4c565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905560405133907f203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d53690600090a2565b60075460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f9726d7faf7429d6b059560dc858ed769377ccdf8b7541eabe12b22548719831f90600090a3600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905560405133907fa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b90600090a2565b336000908152600460205260409020546116b4908290613fec565b3360009081526004602052604090205550565b60008160e001511180156116de5750428160e00151105b15611715576040517fc56873ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61171f82826111a6565b6103e88161012001511115611760576040517fcd4e616700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61176d8160800151610b15565b60008281526008602052604090205460ff16156117b6576040517f7b38b76e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117f0816020015182610100015173ffffffffffffffffffffffffffffffffffffffff919091166000908152600460205260409020541490565b610ca9576040517f756688fe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8183148061183a575082158061183a575081155b15611871576040517f3f6cc76800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008381526005602052604090205415158061189a575060008281526005602052604090205415155b156118d1576040517f3a81d6fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182528381526020808201848152600087815260058084528582209451855591516001948501558451808601865288815280840187815288835292909352848120925183559051919092015590518291849186917fbc9a2432e8aeb48327246cddd6e872ef452812b4243c04e6bfb786a2cd8faf0d91a48083837fbc9a2432e8aeb48327246cddd6e872ef452812b4243c04e6bfb786a2cd8faf0d60405160405180910390a4505050565b60008061198e8461205a565b8051906020012090506000856040516020016119c6919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815282825280516020918201207fff000000000000000000000000000000000000000000000000000000000000008285015260609790971b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166021840152603583019690965260558083019490945280518083039094018452607590910190525080519201919091209392505050565b602081015173ffffffffffffffffffffffffffffffffffffffff163314611ad3576040517f30cd747100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611ade826111fb565b600081815260086020526040902080549192509060ff1615611b2c576040517f7b38b76e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117815560405182907f5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d90600090a2505050565b825160005b81811015611bd757611bcf858281518110611ba757611ba7613f8e565b6020026020010151858381518110611bc157611bc1613f8e565b602002602001015185611f4c565b600101611b8a565b5050505050565b600080826002811115611bf357611bf3613ead565b03611c0b57611c04868686866120eb565b9050610fdc565b6002826002811115611c1f57611c1f613ead565b03611c3057611c0486868686612139565b611c048686868661218d565b6000610985611c496121c1565b836040517f19010000000000000000000000000000000000000000000000000000000000006020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b81600080611cb387846122f5565b91509150600080611cc389612342565b91509150611cd78960200151308488612379565b611ce28989886123a3565b611cec84826123f5565b6101208a0151909450600090611d2e90828c61014001516001811115611d1457611d14613ead565b14611d1f5787611d21565b865b88888e610140015161243d565b9050611d4b308b60200151848489611d469190614004565b612379565b611d573033848461252d565b6000611d6284612596565b90508015611d7a57611d7a308c602001518684612379565b60208b8101516040805187815292830186905282018990526060820188905260808201849052309173ffffffffffffffffffffffffffffffffffffffff9091169087907fd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f69060a00160405180910390a46020808c01516040805187815292830186905282018990526060820188905273ffffffffffffffffffffffffffffffffffffffff169086907f63bf4d16b7fa898ef4c4b2b6d90fd201e9c56313b65638af6088d149d2ce956c9060800160405180910390a35050505050505050505050565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166020820152600090611eb49083908590603401604051602081830303815290604052805190602001206126c6565b90505b9392505050565b60065460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f3053c6252a932554235c173caffc1913604dba3a41cee89516f631c4a1a50a3790600090a3600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b81600080611f5a86846122f5565b6101208801519193509150600090611fa790825b8961014001516001811115611f8557611f85613ead565b14611f905785611f92565b845b8960a001518a60c001518b610140015161243d565b9050600080611fb589612342565b91509150611fcf338a60200151838689611d469190614004565b611fdf8960200151888489612379565b6020898101516040805185815292830184905282018890526060820187905260808201859052339173ffffffffffffffffffffffffffffffffffffffff9091169086907fd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f69060a00160405180910390a4505050505050505050565b6060604051806101a0016040528061017181526020016141d461017191396040805173ffffffffffffffffffffffffffffffffffffffff8516602082015201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526120d59291602001614047565b6040516020818303038152906040529050919050565b60008373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614801561212e575061212e858484612763565b90505b949350505050565b6000612146858484612763565b801561212e57508373ffffffffffffffffffffffffffffffffffffffff1661216d86610ee8565b73ffffffffffffffffffffffffffffffffffffffff161495945050505050565b600061219a858484612763565b801561212e57508373ffffffffffffffffffffffffffffffffffffffff1661216d866113a6565b60003073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561222757507f000000000000000000000000000000000000000000000000000000000000000046145b1561225157507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b60008061230584606001516127a5565b61230e846111fb565b905061231a81856116c7565b61232d838560a001518660c00151612817565b915061233a81858561283e565b509250929050565b60008080836101400151600181111561235d5761235d613ead565b0361236d57505060800151600091565b50506080015190600090565b816000036123915761238c8484836128eb565b61239d565b61239d84848484612940565b50505050565b815160005b81811015611bd7576123ed858583815181106123c6576123c6613f8e565b60200260200101518584815181106123e0576123e0613f8e565b602002602001015161296d565b6001016123a8565b60008061240183612596565b905083811015611eb7576040517fdf4d808000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008515610fdc576000612452858585612a52565b905060008111801561246c5750670de0b6b3a76400008111155b1561252357600083600181111561248557612485613ead565b036124d75761249661271082614076565b866124b2836124ad81670de0b6b3a7640000614004565b612ac1565b6124bc908a614076565b6124c69190614076565b6124d091906140b3565b9150612523565b6124eb670de0b6b3a7640000612710614076565b86612502836124ad81670de0b6b3a7640000614004565b61250c908a614076565b6125169190614076565b61252091906140b3565b91505b5095945050505050565b801561239d5761253f84848484612379565b604080518381526020810183905273ffffffffffffffffffffffffffffffffffffffff8516917facffcc86834d0f1a64b0d5a675798deed6ff0bcfc2231edd3480e7288dba7ff4910160405180910390a250505050565b60008160000361264f576040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a08231906024015b602060405180830381865afa15801561262b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098591906140ee565b6040517efdd58e0000000000000000000000000000000000000000000000000000000081523060048201526024810183905273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169062fdd58e9060440161260e565b6000806126d38585612ad7565b8051602091820120604080517fff000000000000000000000000000000000000000000000000000000000000008185015260609890981b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166021890152603588019590955260558088019190915284518088039091018152607590960190935250508251920191909120919050565b60008373ffffffffffffffffffffffffffffffffffffffff166127868484612c5a565b73ffffffffffffffffffffffffffffffffffffffff1614949350505050565b73ffffffffffffffffffffffffffffffffffffffff8116158015906127e0575073ffffffffffffffffffffffffffffffffffffffff81163314155b15610b5c576040517f5211a07900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260000361282957506000611eb7565b826128348386614076565b611eb491906140b3565b6000838152600860205260409020600181015490811561285e5781612864565b8360a001515b9150818311156128a0576040517fe2cc6ad600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6128aa8383614004565b9150816000036128de5780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781555b6001018190559392505050565b7f00000000000000000000000000000000000000000000000000000000000000003073ffffffffffffffffffffffffffffffffffffffff8516036129345761238c818484612c7e565b61239d81858585612c89565b61239d7f000000000000000000000000000000000000000000000000000000000000000085858585612c95565b60006129798484612d41565b9050612986848483612ddd565b8160008061299486846122f5565b61012088015191935091506000906129ac9082611f6e565b90506000806129ba89612342565b915091506129d186868b6020015185858c89612e89565b6020808b01518a820151604080518681529384018590528301899052606083018890526080830186905273ffffffffffffffffffffffffffffffffffffffff9182169291169086907fd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f69060a00160405180910390a450505050505050505050565b600080826001811115612a6757612a67613ead565b03612a9f5782600003612a7b576000612a98565b82612a8e670de0b6b3a764000086614076565b612a9891906140b3565b9050611eb7565b83600003612aae576000611eb4565b83612834670de0b6b3a764000085614076565b6000818310612ad05781611eb7565b5090919050565b6040805160008082526020820190925260609190612af89060448101614107565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f52e831dd000000000000000000000000000000000000000000000000000000001790528151606380825260a082019093529293506000929190820181803683370190505090507f3d3d606380380380913d393d73bebebebebebebebebebebebebebebebebebebe60208201526c010000000000000000000000008502602d8201527f5af4602a57600080fd5b602d8060366000396000f3363d3d373d3d3d363d73be60418201526c01000000000000000000000000840260608201527f5af43d82803e903d91602b57fd5bf3000000000000000000000000000000000060748201528082604051602001612c41929190614047565b6040516020818303038152906040529250505092915050565b6000806000612c698585612f09565b91509150612c7681612f4e565b509392505050565b610d02838383613101565b61239d848484846131d0565b6040517ff242432a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301528481166024830152604482018490526064820183905260a06084830152600060a483015286169063f242432a9060c401600060405180830381600087803b158015612d2257600080fd5b505af1158015612d36573d6000803e3d6000fd5b505050505050505050565b6000808361014001516001811115612d5b57612d5b613ead565b148015612d7e575060008261014001516001811115612d7c57612d7c613ead565b145b15612d8b57506001610985565b60018361014001516001811115612da457612da4613ead565b148015612dc7575060018261014001516001811115612dc557612dc5613ead565b145b15612dd457506002610985565b50600092915050565b612de783836132bb565b612e1d576040517f7f9a6f4600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000816002811115612e3157612e31613ead565b03612e77578160800151836080015114610d02576040517fa0b9446500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d02836080015183608001516110f5565b612e958530868a612379565b612ea28787868686613305565b85612eac84612596565b1015612ee4576040517fdf4d808000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612ef4308685611d46858b614004565b612f003033858461252d565b50505050505050565b6000808251604103612f3f5760208301516040840151606085015160001a612f338782858561338d565b94509450505050612f47565b506000905060025b9250929050565b6000816004811115612f6257612f62613ead565b03612f6a5750565b6001816004811115612f7e57612f7e613ead565b03612fe5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161104d565b6002816004811115612ff957612ff9613ead565b03613060576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161104d565b600381600481111561307457613074613ead565b03610b5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161104d565b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152826024820152602060006044836000895af13d15601f3d116001600051141617169150508061239d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f5452414e534645525f4641494c45440000000000000000000000000000000000604482015260640161104d565b60006040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff841660248201528260448201526020600060648360008a5af13d15601f3d1160016000511416171691505080611bd7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5452414e534645525f46524f4d5f4641494c4544000000000000000000000000604482015260640161104d565b60008260c00151600014806132d2575060c0820151155b156132df57506001610985565b611eb76132eb8461347c565b6132f48461347c565b856101400151856101400151613496565b600081600281111561331957613319613ead565b14611bd757600181600281111561333257613332613ead565b03613358576000828152600560205260409020600101546133539085613530565b611bd7565b600281600281111561336c5761336c613ead565b03611bd7576000838152600560205260409020600101546133539086613660565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156133c45750600090506003613473565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613418573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661346c57600060019250925050613473565b9150600090505b94509492505050565b60006109858260a001518360c00151846101400151612a52565b6000808360018111156134ab576134ab613ead565b036134ef5760008260018111156134c4576134c4613ead565b036134e557670de0b6b3a76400006134dc8587613fec565b10159050612131565b5082841015612131565b600082600181111561350357613503613ead565b03613512575083831015612131565b670de0b6b3a76400006135258587613fec565b111595945050505050565b60408051600280825260608201835260009260208301908036833701905050905060018160008151811061356657613566613f8e565b60200260200101818152505060028160018151811061358757613587613f8e565b60209081029190910101527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166372ce42757f00000000000000000000000000000000000000000000000000000000000000005b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526136329190600090889087908990600401614158565b600060405180830381600087803b15801561364c57600080fd5b505af1158015612f00573d6000803e3d6000fd5b60408051600280825260608201835260009260208301908036833701905050905060018160008151811061369657613696613f8e565b6020026020010181815250506002816001815181106136b7576136b7613f8e565b60209081029190910101527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639e7212ad7f00000000000000000000000000000000000000000000000000000000000000006135f0565b60006020828403121561373657600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611eb757600080fd5b73ffffffffffffffffffffffffffffffffffffffff81168114610b5c57600080fd5b803561379381613766565b919050565b600080604083850312156137ab57600080fd5b82356137b681613766565b946020939093013593505050565b6000602082840312156137d657600080fd5b8135611eb781613766565b6000602082840312156137f357600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516101a0810167ffffffffffffffff8111828210171561384d5761384d6137fa565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561389a5761389a6137fa565b604052919050565b80356002811061379357600080fd5b80356003811061379357600080fd5b600082601f8301126138d157600080fd5b813567ffffffffffffffff8111156138eb576138eb6137fa565b61391c60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613853565b81815284602083860101111561393157600080fd5b816020850160208301376000918101602001919091529392505050565b60006101a0828403121561396157600080fd5b613969613829565b90508135815261397b60208301613788565b602082015261398c60408301613788565b604082015261399d60608301613788565b60608201526080820135608082015260a082013560a082015260c082013560c082015260e082013560e08201526101008083013581830152506101208083013581830152506101406139f08184016138a2565b90820152610160613a028382016138b1565b908201526101808281013567ffffffffffffffff811115613a2257600080fd5b613a2e858286016138c0565b82840152505092915050565b600060208284031215613a4c57600080fd5b813567ffffffffffffffff811115613a6357600080fd5b6121318482850161394e565b600080600060608486031215613a8457600080fd5b505081359360208301359350604090920135919050565b600067ffffffffffffffff821115613ab557613ab56137fa565b5060051b60200190565b600082601f830112613ad057600080fd5b81356020613ae5613ae083613a9b565b613853565b82815260059290921b84018101918181019086841115613b0457600080fd5b8286015b84811015613b1f5780358352918301918301613b08565b509695505050505050565b600080600080600060a08688031215613b4257600080fd5b8535613b4d81613766565b94506020860135613b5d81613766565b9350604086013567ffffffffffffffff80821115613b7a57600080fd5b613b8689838a01613abf565b94506060880135915080821115613b9c57600080fd5b613ba889838a01613abf565b93506080880135915080821115613bbe57600080fd5b50613bcb888289016138c0565b9150509295509295909350565b600082601f830112613be957600080fd5b81356020613bf9613ae083613a9b565b82815260059290921b84018101918181019086841115613c1857600080fd5b8286015b84811015613b1f57803567ffffffffffffffff811115613c3c5760008081fd5b613c4a8986838b010161394e565b845250918301918301613c1c565b60008060408385031215613c6b57600080fd5b823567ffffffffffffffff80821115613c8357600080fd5b613c8f86838701613bd8565b93506020850135915080821115613ca557600080fd5b50613cb285828601613abf565b9150509250929050565b60008060408385031215613ccf57600080fd5b50508035926020909101359150565b60008060408385031215613cf157600080fd5b82359150602083013567ffffffffffffffff811115613d0f57600080fd5b613cb28582860161394e565b60008060008060808587031215613d3157600080fd5b843567ffffffffffffffff80821115613d4957600080fd5b613d558883890161394e565b95506020870135915080821115613d6b57600080fd5b613d7788838901613bd8565b9450604087013593506060870135915080821115613d9457600080fd5b50613da187828801613abf565b91505092959194509250565b600080600080600060a08688031215613dc557600080fd5b8535613dd081613766565b94506020860135613de081613766565b93506040860135925060608601359150608086013567ffffffffffffffff811115613e0a57600080fd5b613bcb888289016138c0565b600060208284031215613e2857600080fd5b813567ffffffffffffffff811115613e3f57600080fd5b61213184828501613bd8565b60008060408385031215613e5e57600080fd5b823567ffffffffffffffff811115613e7557600080fd5b613e818582860161394e565b95602094909401359450505050565b600060208284031215613ea257600080fd5b8151611eb781613766565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110613eec57613eec613ead565b9052565b60006101a0820190508e82528d602083015273ffffffffffffffffffffffffffffffffffffffff808e166040840152808d166060840152808c166080840152508960a08301528860c08301528760e083015286610100830152856101208301528461014083015260028410613f6757613f67613ead565b83610160830152613f7c610180830184613edc565b9e9d5050505050505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115613fff57613fff613fbd565b500190565b60008282101561401657614016613fbd565b500390565b60005b8381101561403657818101518382015260200161401e565b8381111561239d5750506000910152565b6000835161405981846020880161401b565b83519083019061406d81836020880161401b565b01949350505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156140ae576140ae613fbd565b500290565b6000826140e9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60006020828403121561410057600080fd5b5051919050565b602081526000825180602084015261412681604085016020870161401b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600060a0820173ffffffffffffffffffffffffffffffffffffffff881683526020878185015286604085015260a0606085015281865180845260c086019150828801935060005b818110156141bb5784518352938301939183019160010161419f565b5050809350505050826080830152969550505050505056fe608060405234801561001057600080fd5b5060405161017138038061017183398101604081905261002f916100b9565b6001600160a01b0381166100945760405162461bcd60e51b815260206004820152602260248201527f496e76616c69642073696e676c65746f6e20616464726573732070726f766964604482015261195960f21b606482015260840160405180910390fd5b600080546001600160a01b0319166001600160a01b03929092169190911790556100e7565b6000602082840312156100ca578081fd5b81516001600160a01b03811681146100e0578182fd5b9392505050565b607c806100f56000396000f3fe6080604052600080546001600160a01b0316813563530ca43760e11b1415602857808252602082f35b3682833781823684845af490503d82833e806041573d82fd5b503d81f3fea264697066735822122015938e3bf2c49f5df5c1b7f9569fa85cc5d6f3074bb258a2dc0c7e299bc9e33664736f6c63430008040033a2646970667358221220fcb5c2e27e576ebe35c2aaad79fd81b4c2d84de9e748412eaa9933ecb6e9818c64736f6c634300080f0033",
}

// NegRiskCtfExchange is an auto generated Go binding around an Ethereum contract.
type NegRiskCtfExchange struct {
	abi abi.ABI
}

// NewNegRiskCtfExchange creates a new instance of NegRiskCtfExchange.
func NewNegRiskCtfExchange() *NegRiskCtfExchange {
	parsed, err := NegRiskCtfExchangeMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &NegRiskCtfExchange{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *NegRiskCtfExchange) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _collateral, address _ctf, address _negRiskAdapter, address _proxyFactory, address _safeFactory) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackConstructor(_collateral common.Address, _ctf common.Address, _negRiskAdapter common.Address, _proxyFactory common.Address, _safeFactory common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("", _collateral, _ctf, _negRiskAdapter, _proxyFactory, _safeFactory)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAddAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70480275.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addAdmin(address admin_) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackAddAdmin(admin common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("addAdmin", admin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70480275.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addAdmin(address admin_) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackAddAdmin(admin common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("addAdmin", admin)
}

// PackAddOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9870d7fe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addOperator(address operator_) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackAddOperator(operator common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("addOperator", operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9870d7fe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addOperator(address operator_) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackAddOperator(operator common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("addOperator", operator)
}

// PackAdmins is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x429b62e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function admins(address ) view returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) PackAdmins(arg0 common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("admins", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAdmins is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x429b62e5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function admins(address ) view returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackAdmins(arg0 common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("admins", arg0)
}

// UnpackAdmins is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackAdmins(data []byte) (*big.Int, error) {
	out, err := negRiskCtfExchange.abi.Unpack("admins", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCancelOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa6dfcf86.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackCancelOrder(order Order) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("cancelOrder", order)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa6dfcf86.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackCancelOrder(order Order) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("cancelOrder", order)
}

// PackCancelOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa950b48.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackCancelOrders(orders []Order) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("cancelOrders", orders)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa950b48.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackCancelOrders(orders []Order) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("cancelOrders", orders)
}

// PackDomainSeparator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf698da25.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (negRiskCtfExchange *NegRiskCtfExchange) PackDomainSeparator() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("domainSeparator")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDomainSeparator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf698da25.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackDomainSeparator() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("domainSeparator")
}

// UnpackDomainSeparator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackDomainSeparator(data []byte) ([32]byte, error) {
	out, err := negRiskCtfExchange.abi.Unpack("domainSeparator", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackFillOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe729aaf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackFillOrder(order Order, fillAmount *big.Int) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("fillOrder", order, fillAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe729aaf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackFillOrder(order Order, fillAmount *big.Int) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("fillOrder", order, fillAmount)
}

// PackFillOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd798eff6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackFillOrders(orders []Order, fillAmounts []*big.Int) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("fillOrders", orders, fillAmounts)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd798eff6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackFillOrders(orders []Order, fillAmounts []*big.Int) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("fillOrders", orders, fillAmounts)
}

// PackGetCollateral is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c1548fb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCollateral() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) PackGetCollateral() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("getCollateral")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCollateral is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c1548fb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCollateral() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackGetCollateral() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("getCollateral")
}

// UnpackGetCollateral is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackGetCollateral(data []byte) (common.Address, error) {
	out, err := negRiskCtfExchange.abi.Unpack("getCollateral", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetComplement is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa10f3dce.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getComplement(uint256 token) view returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) PackGetComplement(token *big.Int) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("getComplement", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetComplement is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa10f3dce.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getComplement(uint256 token) view returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackGetComplement(token *big.Int) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("getComplement", token)
}

// UnpackGetComplement is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa10f3dce.
//
// Solidity: function getComplement(uint256 token) view returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackGetComplement(data []byte) (*big.Int, error) {
	out, err := negRiskCtfExchange.abi.Unpack("getComplement", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetConditionId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd7fb272f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getConditionId(uint256 token) view returns(bytes32)
func (negRiskCtfExchange *NegRiskCtfExchange) PackGetConditionId(token *big.Int) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("getConditionId", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetConditionId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd7fb272f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getConditionId(uint256 token) view returns(bytes32)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackGetConditionId(token *big.Int) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("getConditionId", token)
}

// UnpackGetConditionId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd7fb272f.
//
// Solidity: function getConditionId(uint256 token) view returns(bytes32)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackGetConditionId(data []byte) ([32]byte, error) {
	out, err := negRiskCtfExchange.abi.Unpack("getConditionId", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetCtf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3b521d78.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCtf() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) PackGetCtf() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("getCtf")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCtf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3b521d78.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCtf() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackGetCtf() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("getCtf")
}

// UnpackGetCtf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackGetCtf(data []byte) (common.Address, error) {
	out, err := negRiskCtfExchange.abi.Unpack("getCtf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetMaxFeeRate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4a2a11f5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMaxFeeRate() pure returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) PackGetMaxFeeRate() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("getMaxFeeRate")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMaxFeeRate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4a2a11f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMaxFeeRate() pure returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackGetMaxFeeRate() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("getMaxFeeRate")
}

// UnpackGetMaxFeeRate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() pure returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackGetMaxFeeRate(data []byte) (*big.Int, error) {
	out, err := negRiskCtfExchange.abi.Unpack("getMaxFeeRate", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetOrderStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46423aa7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint256))
func (negRiskCtfExchange *NegRiskCtfExchange) PackGetOrderStatus(orderHash [32]byte) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("getOrderStatus", orderHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetOrderStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46423aa7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint256))
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackGetOrderStatus(orderHash [32]byte) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("getOrderStatus", orderHash)
}

// UnpackGetOrderStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint256))
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackGetOrderStatus(data []byte) (OrderStatus, error) {
	out, err := negRiskCtfExchange.abi.Unpack("getOrderStatus", data)
	if err != nil {
		return *new(OrderStatus), err
	}
	out0 := *abi.ConvertType(out[0], new(OrderStatus)).(*OrderStatus)
	return out0, nil
}

// PackGetPolyProxyFactoryImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06b9d691.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPolyProxyFactoryImplementation() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) PackGetPolyProxyFactoryImplementation() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("getPolyProxyFactoryImplementation")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPolyProxyFactoryImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06b9d691.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPolyProxyFactoryImplementation() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackGetPolyProxyFactoryImplementation() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("getPolyProxyFactoryImplementation")
}

// UnpackGetPolyProxyFactoryImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06b9d691.
//
// Solidity: function getPolyProxyFactoryImplementation() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackGetPolyProxyFactoryImplementation(data []byte) (common.Address, error) {
	out, err := negRiskCtfExchange.abi.Unpack("getPolyProxyFactoryImplementation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetPolyProxyWalletAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xedef7d8e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPolyProxyWalletAddress(address _addr) view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) PackGetPolyProxyWalletAddress(addr common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("getPolyProxyWalletAddress", addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPolyProxyWalletAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xedef7d8e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPolyProxyWalletAddress(address _addr) view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackGetPolyProxyWalletAddress(addr common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("getPolyProxyWalletAddress", addr)
}

// UnpackGetPolyProxyWalletAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xedef7d8e.
//
// Solidity: function getPolyProxyWalletAddress(address _addr) view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackGetPolyProxyWalletAddress(data []byte) (common.Address, error) {
	out, err := negRiskCtfExchange.abi.Unpack("getPolyProxyWalletAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetProxyFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb28c51c0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getProxyFactory() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) PackGetProxyFactory() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("getProxyFactory")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetProxyFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb28c51c0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getProxyFactory() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackGetProxyFactory() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("getProxyFactory")
}

// UnpackGetProxyFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackGetProxyFactory(data []byte) (common.Address, error) {
	out, err := negRiskCtfExchange.abi.Unpack("getProxyFactory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetSafeAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa287bdf1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSafeAddress(address _addr) view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) PackGetSafeAddress(addr common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("getSafeAddress", addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSafeAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa287bdf1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSafeAddress(address _addr) view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackGetSafeAddress(addr common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("getSafeAddress", addr)
}

// UnpackGetSafeAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa287bdf1.
//
// Solidity: function getSafeAddress(address _addr) view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackGetSafeAddress(data []byte) (common.Address, error) {
	out, err := negRiskCtfExchange.abi.Unpack("getSafeAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetSafeFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75d7370a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSafeFactory() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) PackGetSafeFactory() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("getSafeFactory")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSafeFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75d7370a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSafeFactory() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackGetSafeFactory() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("getSafeFactory")
}

// UnpackGetSafeFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackGetSafeFactory(data []byte) (common.Address, error) {
	out, err := negRiskCtfExchange.abi.Unpack("getSafeFactory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetSafeFactoryImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe03ac3d0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSafeFactoryImplementation() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) PackGetSafeFactoryImplementation() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("getSafeFactoryImplementation")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSafeFactoryImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe03ac3d0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSafeFactoryImplementation() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackGetSafeFactoryImplementation() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("getSafeFactoryImplementation")
}

// UnpackGetSafeFactoryImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe03ac3d0.
//
// Solidity: function getSafeFactoryImplementation() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackGetSafeFactoryImplementation(data []byte) (common.Address, error) {
	out, err := negRiskCtfExchange.abi.Unpack("getSafeFactoryImplementation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackHashOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe50e4f97.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hashOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns(bytes32)
func (negRiskCtfExchange *NegRiskCtfExchange) PackHashOrder(order Order) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("hashOrder", order)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHashOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe50e4f97.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hashOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns(bytes32)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackHashOrder(order Order) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("hashOrder", order)
}

// UnpackHashOrder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe50e4f97.
//
// Solidity: function hashOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns(bytes32)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackHashOrder(data []byte) ([32]byte, error) {
	out, err := negRiskCtfExchange.abi.Unpack("hashOrder", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackIncrementNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x627cdcb9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function incrementNonce() returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackIncrementNonce() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("incrementNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIncrementNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x627cdcb9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function incrementNonce() returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackIncrementNonce() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("incrementNonce")
}

// PackIsAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24d7806c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) PackIsAdmin(usr common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("isAdmin", usr)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24d7806c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackIsAdmin(usr common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("isAdmin", usr)
}

// UnpackIsAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackIsAdmin(data []byte) (bool, error) {
	out, err := negRiskCtfExchange.abi.Unpack("isAdmin", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d70f7ae.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) PackIsOperator(usr common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("isOperator", usr)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d70f7ae.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackIsOperator(usr common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("isOperator", usr)
}

// UnpackIsOperator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackIsOperator(data []byte) (bool, error) {
	out, err := negRiskCtfExchange.abi.Unpack("isOperator", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsValidNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0647ee20.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isValidNonce(address usr, uint256 nonce) view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) PackIsValidNonce(usr common.Address, nonce *big.Int) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("isValidNonce", usr, nonce)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsValidNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0647ee20.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isValidNonce(address usr, uint256 nonce) view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackIsValidNonce(usr common.Address, nonce *big.Int) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("isValidNonce", usr, nonce)
}

// UnpackIsValidNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0647ee20.
//
// Solidity: function isValidNonce(address usr, uint256 nonce) view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackIsValidNonce(data []byte) (bool, error) {
	out, err := negRiskCtfExchange.abi.Unpack("isValidNonce", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackMatchOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe60f0c05.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackMatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("matchOrders", takerOrder, makerOrders, takerFillAmount, makerFillAmounts)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMatchOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe60f0c05.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackMatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("matchOrders", takerOrder, makerOrders, takerFillAmount, makerFillAmounts)
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nonces(address ) view returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) PackNonces(arg0 common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("nonces", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function nonces(address ) view returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackNonces(arg0 common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("nonces", arg0)
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := negRiskCtfExchange.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (negRiskCtfExchange *NegRiskCtfExchange) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := negRiskCtfExchange.abi.Unpack("onERC1155BatchReceived", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (negRiskCtfExchange *NegRiskCtfExchange) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := negRiskCtfExchange.abi.Unpack("onERC1155Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackOperators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13e7c9d8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function operators(address ) view returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) PackOperators(arg0 common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("operators", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOperators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13e7c9d8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function operators(address ) view returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackOperators(arg0 common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("operators", arg0)
}

// UnpackOperators is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackOperators(data []byte) (*big.Int, error) {
	out, err := negRiskCtfExchange.abi.Unpack("operators", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOrderStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2dff692d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool isFilledOrCancelled, uint256 remaining)
func (negRiskCtfExchange *NegRiskCtfExchange) PackOrderStatus(arg0 [32]byte) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("orderStatus", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOrderStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2dff692d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool isFilledOrCancelled, uint256 remaining)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackOrderStatus(arg0 [32]byte) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("orderStatus", arg0)
}

// OrderStatusOutput serves as a container for the return parameters of contract
// method OrderStatus.
type OrderStatusOutput struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}

// UnpackOrderStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool isFilledOrCancelled, uint256 remaining)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackOrderStatus(data []byte) (OrderStatusOutput, error) {
	out, err := negRiskCtfExchange.abi.Unpack("orderStatus", data)
	outstruct := new(OrderStatusOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.IsFilledOrCancelled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Remaining = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackParentCollectionId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44bea37e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function parentCollectionId() view returns(bytes32)
func (negRiskCtfExchange *NegRiskCtfExchange) PackParentCollectionId() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("parentCollectionId")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackParentCollectionId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44bea37e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function parentCollectionId() view returns(bytes32)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackParentCollectionId() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("parentCollectionId")
}

// UnpackParentCollectionId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x44bea37e.
//
// Solidity: function parentCollectionId() view returns(bytes32)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackParentCollectionId(data []byte) ([32]byte, error) {
	out, err := negRiskCtfExchange.abi.Unpack("parentCollectionId", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackPauseTrading is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1031e36e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pauseTrading() returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackPauseTrading() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("pauseTrading")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPauseTrading is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1031e36e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pauseTrading() returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackPauseTrading() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("pauseTrading")
}

// PackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function paused() view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) PackPaused() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("paused")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function paused() view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackPaused() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("paused")
}

// UnpackPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackPaused(data []byte) (bool, error) {
	out, err := negRiskCtfExchange.abi.Unpack("paused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackProxyFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc10f1a75.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proxyFactory() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) PackProxyFactory() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("proxyFactory")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProxyFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc10f1a75.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proxyFactory() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackProxyFactory() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("proxyFactory")
}

// UnpackProxyFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackProxyFactory(data []byte) (common.Address, error) {
	out, err := negRiskCtfExchange.abi.Unpack("proxyFactory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRegisterToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x68c7450f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registerToken(uint256 token, uint256 complement, bytes32 conditionId) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackRegisterToken(token *big.Int, complement *big.Int, conditionId [32]byte) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("registerToken", token, complement, conditionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegisterToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x68c7450f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registerToken(uint256 token, uint256 complement, bytes32 conditionId) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackRegisterToken(token *big.Int, complement *big.Int, conditionId [32]byte) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("registerToken", token, complement, conditionId)
}

// PackRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5893253c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registry(uint256 ) view returns(uint256 complement, bytes32 conditionId)
func (negRiskCtfExchange *NegRiskCtfExchange) PackRegistry(arg0 *big.Int) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("registry", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5893253c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registry(uint256 ) view returns(uint256 complement, bytes32 conditionId)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackRegistry(arg0 *big.Int) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("registry", arg0)
}

// RegistryOutput serves as a container for the return parameters of contract
// method Registry.
type RegistryOutput struct {
	Complement  *big.Int
	ConditionId [32]byte
}

// UnpackRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5893253c.
//
// Solidity: function registry(uint256 ) view returns(uint256 complement, bytes32 conditionId)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackRegistry(data []byte) (RegistryOutput, error) {
	out, err := negRiskCtfExchange.abi.Unpack("registry", data)
	outstruct := new(RegistryOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Complement = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.ConditionId = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackRemoveAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1785f53c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeAdmin(address admin) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackRemoveAdmin(admin common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("removeAdmin", admin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1785f53c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeAdmin(address admin) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackRemoveAdmin(admin common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("removeAdmin", admin)
}

// PackRemoveOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac8a584a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeOperator(address operator) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackRemoveOperator(operator common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("removeOperator", operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac8a584a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeOperator(address operator) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackRemoveOperator(operator common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("removeOperator", operator)
}

// PackRenounceAdminRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83b8a5ae.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceAdminRole() returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackRenounceAdminRole() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("renounceAdminRole")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceAdminRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83b8a5ae.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceAdminRole() returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackRenounceAdminRole() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("renounceAdminRole")
}

// PackRenounceOperatorRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3d6d3598.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOperatorRole() returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackRenounceOperatorRole() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("renounceOperatorRole")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceOperatorRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3d6d3598.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceOperatorRole() returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackRenounceOperatorRole() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("renounceOperatorRole")
}

// PackSafeFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x131e7e1c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function safeFactory() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) PackSafeFactory() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("safeFactory")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSafeFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x131e7e1c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function safeFactory() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackSafeFactory() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("safeFactory")
}

// UnpackSafeFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x131e7e1c.
//
// Solidity: function safeFactory() view returns(address)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackSafeFactory(data []byte) (common.Address, error) {
	out, err := negRiskCtfExchange.abi.Unpack("safeFactory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetProxyFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfbddd751.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setProxyFactory(address _newProxyFactory) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackSetProxyFactory(newProxyFactory common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("setProxyFactory", newProxyFactory)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetProxyFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfbddd751.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setProxyFactory(address _newProxyFactory) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackSetProxyFactory(newProxyFactory common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("setProxyFactory", newProxyFactory)
}

// PackSetSafeFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4544f055.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSafeFactory(address _newSafeFactory) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackSetSafeFactory(newSafeFactory common.Address) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("setSafeFactory", newSafeFactory)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSafeFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4544f055.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSafeFactory(address _newSafeFactory) returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackSetSafeFactory(newSafeFactory common.Address) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("setSafeFactory", newSafeFactory)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := negRiskCtfExchange.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackUnpauseTrading is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x456068d2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unpauseTrading() returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackUnpauseTrading() []byte {
	enc, err := negRiskCtfExchange.abi.Pack("unpauseTrading")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnpauseTrading is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x456068d2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unpauseTrading() returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackUnpauseTrading() ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("unpauseTrading")
}

// PackValidateComplement is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd82da838.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function validateComplement(uint256 token, uint256 complement) view returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackValidateComplement(token *big.Int, complement *big.Int) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("validateComplement", token, complement)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackValidateComplement is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd82da838.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function validateComplement(uint256 token, uint256 complement) view returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackValidateComplement(token *big.Int, complement *big.Int) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("validateComplement", token, complement)
}

// PackValidateOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x654f0ce4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function validateOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackValidateOrder(order Order) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("validateOrder", order)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackValidateOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x654f0ce4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function validateOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackValidateOrder(order Order) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("validateOrder", order)
}

// PackValidateOrderSignature is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2eec405.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackValidateOrderSignature(orderHash [32]byte, order Order) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("validateOrderSignature", orderHash, order)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackValidateOrderSignature is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2eec405.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackValidateOrderSignature(orderHash [32]byte, order Order) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("validateOrderSignature", orderHash, order)
}

// PackValidateTokenId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x34600901.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function validateTokenId(uint256 tokenId) view returns()
func (negRiskCtfExchange *NegRiskCtfExchange) PackValidateTokenId(tokenId *big.Int) []byte {
	enc, err := negRiskCtfExchange.abi.Pack("validateTokenId", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackValidateTokenId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x34600901.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function validateTokenId(uint256 tokenId) view returns()
func (negRiskCtfExchange *NegRiskCtfExchange) TryPackValidateTokenId(tokenId *big.Int) ([]byte, error) {
	return negRiskCtfExchange.abi.Pack("validateTokenId", tokenId)
}

// NegRiskCtfExchangeFeeCharged represents a FeeCharged event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeFeeCharged struct {
	Receiver common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeFeeChargedEventName = "FeeCharged"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeFeeCharged) ContractEventName() string {
	return NegRiskCtfExchangeFeeChargedEventName
}

// UnpackFeeChargedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 tokenId, uint256 amount)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackFeeChargedEvent(log *types.Log) (*NegRiskCtfExchangeFeeCharged, error) {
	event := "FeeCharged"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeFeeCharged)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// NegRiskCtfExchangeNewAdmin represents a NewAdmin event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeNewAdmin struct {
	NewAdminAddress common.Address
	Admin           common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeNewAdminEventName = "NewAdmin"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeNewAdmin) ContractEventName() string {
	return NegRiskCtfExchangeNewAdminEventName
}

// UnpackNewAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackNewAdminEvent(log *types.Log) (*NegRiskCtfExchangeNewAdmin, error) {
	event := "NewAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeNewAdmin)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// NegRiskCtfExchangeNewOperator represents a NewOperator event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeNewOperator struct {
	NewOperatorAddress common.Address
	Admin              common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeNewOperatorEventName = "NewOperator"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeNewOperator) ContractEventName() string {
	return NegRiskCtfExchangeNewOperatorEventName
}

// UnpackNewOperatorEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackNewOperatorEvent(log *types.Log) (*NegRiskCtfExchangeNewOperator, error) {
	event := "NewOperator"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeNewOperator)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// NegRiskCtfExchangeOrderCancelled represents a OrderCancelled event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeOrderCancelled struct {
	OrderHash [32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeOrderCancelledEventName = "OrderCancelled"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeOrderCancelled) ContractEventName() string {
	return NegRiskCtfExchangeOrderCancelledEventName
}

// UnpackOrderCancelledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackOrderCancelledEvent(log *types.Log) (*NegRiskCtfExchangeOrderCancelled, error) {
	event := "OrderCancelled"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeOrderCancelled)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// NegRiskCtfExchangeOrderFilled represents a OrderFilled event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeOrderFilled struct {
	OrderHash         [32]byte
	Maker             common.Address
	Taker             common.Address
	MakerAssetId      *big.Int
	TakerAssetId      *big.Int
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Fee               *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeOrderFilledEventName = "OrderFilled"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeOrderFilled) ContractEventName() string {
	return NegRiskCtfExchangeOrderFilledEventName
}

// UnpackOrderFilledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackOrderFilledEvent(log *types.Log) (*NegRiskCtfExchangeOrderFilled, error) {
	event := "OrderFilled"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeOrderFilled)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// NegRiskCtfExchangeOrdersMatched represents a OrdersMatched event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeOrdersMatched struct {
	TakerOrderHash    [32]byte
	TakerOrderMaker   common.Address
	MakerAssetId      *big.Int
	TakerAssetId      *big.Int
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeOrdersMatchedEventName = "OrdersMatched"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeOrdersMatched) ContractEventName() string {
	return NegRiskCtfExchangeOrdersMatchedEventName
}

// UnpackOrdersMatchedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackOrdersMatchedEvent(log *types.Log) (*NegRiskCtfExchangeOrdersMatched, error) {
	event := "OrdersMatched"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeOrdersMatched)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// NegRiskCtfExchangeProxyFactoryUpdated represents a ProxyFactoryUpdated event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeProxyFactoryUpdated struct {
	OldProxyFactory common.Address
	NewProxyFactory common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeProxyFactoryUpdatedEventName = "ProxyFactoryUpdated"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeProxyFactoryUpdated) ContractEventName() string {
	return NegRiskCtfExchangeProxyFactoryUpdatedEventName
}

// UnpackProxyFactoryUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProxyFactoryUpdated(address indexed oldProxyFactory, address indexed newProxyFactory)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackProxyFactoryUpdatedEvent(log *types.Log) (*NegRiskCtfExchangeProxyFactoryUpdated, error) {
	event := "ProxyFactoryUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeProxyFactoryUpdated)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// NegRiskCtfExchangeRemovedAdmin represents a RemovedAdmin event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeRemovedAdmin struct {
	RemovedAdmin common.Address
	Admin        common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeRemovedAdminEventName = "RemovedAdmin"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeRemovedAdmin) ContractEventName() string {
	return NegRiskCtfExchangeRemovedAdminEventName
}

// UnpackRemovedAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackRemovedAdminEvent(log *types.Log) (*NegRiskCtfExchangeRemovedAdmin, error) {
	event := "RemovedAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeRemovedAdmin)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// NegRiskCtfExchangeRemovedOperator represents a RemovedOperator event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeRemovedOperator struct {
	RemovedOperator common.Address
	Admin           common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeRemovedOperatorEventName = "RemovedOperator"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeRemovedOperator) ContractEventName() string {
	return NegRiskCtfExchangeRemovedOperatorEventName
}

// UnpackRemovedOperatorEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackRemovedOperatorEvent(log *types.Log) (*NegRiskCtfExchangeRemovedOperator, error) {
	event := "RemovedOperator"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeRemovedOperator)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// NegRiskCtfExchangeSafeFactoryUpdated represents a SafeFactoryUpdated event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeSafeFactoryUpdated struct {
	OldSafeFactory common.Address
	NewSafeFactory common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeSafeFactoryUpdatedEventName = "SafeFactoryUpdated"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeSafeFactoryUpdated) ContractEventName() string {
	return NegRiskCtfExchangeSafeFactoryUpdatedEventName
}

// UnpackSafeFactoryUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SafeFactoryUpdated(address indexed oldSafeFactory, address indexed newSafeFactory)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackSafeFactoryUpdatedEvent(log *types.Log) (*NegRiskCtfExchangeSafeFactoryUpdated, error) {
	event := "SafeFactoryUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeSafeFactoryUpdated)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// NegRiskCtfExchangeTokenRegistered represents a TokenRegistered event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeTokenRegistered struct {
	Token0      *big.Int
	Token1      *big.Int
	ConditionId [32]byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeTokenRegisteredEventName = "TokenRegistered"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeTokenRegistered) ContractEventName() string {
	return NegRiskCtfExchangeTokenRegisteredEventName
}

// UnpackTokenRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TokenRegistered(uint256 indexed token0, uint256 indexed token1, bytes32 indexed conditionId)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackTokenRegisteredEvent(log *types.Log) (*NegRiskCtfExchangeTokenRegistered, error) {
	event := "TokenRegistered"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeTokenRegistered)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// NegRiskCtfExchangeTradingPaused represents a TradingPaused event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeTradingPaused struct {
	Pauser common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeTradingPausedEventName = "TradingPaused"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeTradingPaused) ContractEventName() string {
	return NegRiskCtfExchangeTradingPausedEventName
}

// UnpackTradingPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TradingPaused(address indexed pauser)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackTradingPausedEvent(log *types.Log) (*NegRiskCtfExchangeTradingPaused, error) {
	event := "TradingPaused"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeTradingPaused)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// NegRiskCtfExchangeTradingUnpaused represents a TradingUnpaused event raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeTradingUnpaused struct {
	Pauser common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const NegRiskCtfExchangeTradingUnpausedEventName = "TradingUnpaused"

// ContractEventName returns the user-defined event name.
func (NegRiskCtfExchangeTradingUnpaused) ContractEventName() string {
	return NegRiskCtfExchangeTradingUnpausedEventName
}

// UnpackTradingUnpausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackTradingUnpausedEvent(log *types.Log) (*NegRiskCtfExchangeTradingUnpaused, error) {
	event := "TradingUnpaused"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskCtfExchange.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskCtfExchangeTradingUnpaused)
	if len(log.Data) > 0 {
		if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskCtfExchange.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["AlreadyRegistered"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackAlreadyRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["FeeTooHigh"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackFeeTooHighError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["InvalidComplement"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackInvalidComplementError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["InvalidSignature"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["InvalidTokenId"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackInvalidTokenIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["MakingGtRemaining"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackMakingGtRemainingError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["MismatchedTokenIds"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackMismatchedTokenIdsError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["NotAdmin"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackNotAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["NotCrossing"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackNotCrossingError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["NotOperator"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackNotOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["NotOwner"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackNotOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["NotTaker"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackNotTakerError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["OrderExpired"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackOrderExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["OrderFilledOrCancelled"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackOrderFilledOrCancelledError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["Paused"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskCtfExchange.abi.Errors["TooLittleTokensReceived"].ID.Bytes()[:4]) {
		return negRiskCtfExchange.UnpackTooLittleTokensReceivedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// NegRiskCtfExchangeAlreadyRegistered represents a AlreadyRegistered error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeAlreadyRegistered struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyRegistered()
func NegRiskCtfExchangeAlreadyRegisteredErrorID() common.Hash {
	return common.HexToHash("0x3a81d6fc79a5635b1f06f5006b0b65f8bc72a09a41893acedbabc02c660943ca")
}

// UnpackAlreadyRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyRegistered()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackAlreadyRegisteredError(raw []byte) (*NegRiskCtfExchangeAlreadyRegistered, error) {
	out := new(NegRiskCtfExchangeAlreadyRegistered)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "AlreadyRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeFeeTooHigh represents a FeeTooHigh error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeFeeTooHigh struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeTooHigh()
func NegRiskCtfExchangeFeeTooHighErrorID() common.Hash {
	return common.HexToHash("0xcd4e6167a0147beade9e7daca0e52cd42e992cd9c3dc1dd3ce8a2b6956f53601")
}

// UnpackFeeTooHighError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeTooHigh()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackFeeTooHighError(raw []byte) (*NegRiskCtfExchangeFeeTooHigh, error) {
	out := new(NegRiskCtfExchangeFeeTooHigh)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "FeeTooHigh", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeInvalidComplement represents a InvalidComplement error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeInvalidComplement struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidComplement()
func NegRiskCtfExchangeInvalidComplementErrorID() common.Hash {
	return common.HexToHash("0x66f8620ad59ed0a1bf1246a37b03df0eaaa3d8f1c27f74046d5ed7aca26b4741")
}

// UnpackInvalidComplementError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidComplement()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackInvalidComplementError(raw []byte) (*NegRiskCtfExchangeInvalidComplement, error) {
	out := new(NegRiskCtfExchangeInvalidComplement)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "InvalidComplement", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeInvalidNonce represents a InvalidNonce error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func NegRiskCtfExchangeInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackInvalidNonceError(raw []byte) (*NegRiskCtfExchangeInvalidNonce, error) {
	out := new(NegRiskCtfExchangeInvalidNonce)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeInvalidSignature represents a InvalidSignature error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSignature()
func NegRiskCtfExchangeInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0x8baa579fce362245063d36f11747a89dd489c54795634fc673cc0e0db51fedc5")
}

// UnpackInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSignature()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackInvalidSignatureError(raw []byte) (*NegRiskCtfExchangeInvalidSignature, error) {
	out := new(NegRiskCtfExchangeInvalidSignature)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "InvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeInvalidTokenId represents a InvalidTokenId error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeInvalidTokenId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidTokenId()
func NegRiskCtfExchangeInvalidTokenIdErrorID() common.Hash {
	return common.HexToHash("0x3f6cc7688a7a1035429f3e76ee7f3966a043563e74eb2909996229362fa2b66a")
}

// UnpackInvalidTokenIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidTokenId()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackInvalidTokenIdError(raw []byte) (*NegRiskCtfExchangeInvalidTokenId, error) {
	out := new(NegRiskCtfExchangeInvalidTokenId)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "InvalidTokenId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeMakingGtRemaining represents a MakingGtRemaining error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeMakingGtRemaining struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MakingGtRemaining()
func NegRiskCtfExchangeMakingGtRemainingErrorID() common.Hash {
	return common.HexToHash("0xe2cc6ad68a4645963bc882cfe24d414e14bb57144138e0edd88f2522e9a4d78e")
}

// UnpackMakingGtRemainingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MakingGtRemaining()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackMakingGtRemainingError(raw []byte) (*NegRiskCtfExchangeMakingGtRemaining, error) {
	out := new(NegRiskCtfExchangeMakingGtRemaining)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "MakingGtRemaining", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeMismatchedTokenIds represents a MismatchedTokenIds error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeMismatchedTokenIds struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MismatchedTokenIds()
func NegRiskCtfExchangeMismatchedTokenIdsErrorID() common.Hash {
	return common.HexToHash("0xa0b9446586ee4d8ac4c2552cac1d13d4921e8bad2ae26868289a972e4bfc28ad")
}

// UnpackMismatchedTokenIdsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MismatchedTokenIds()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackMismatchedTokenIdsError(raw []byte) (*NegRiskCtfExchangeMismatchedTokenIds, error) {
	out := new(NegRiskCtfExchangeMismatchedTokenIds)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "MismatchedTokenIds", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeNotAdmin represents a NotAdmin error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeNotAdmin struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAdmin()
func NegRiskCtfExchangeNotAdminErrorID() common.Hash {
	return common.HexToHash("0x7bfa4b9fb0cd3687c1f539f384b3f3f258f2c9aa9186353d0815413b508ed97d")
}

// UnpackNotAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAdmin()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackNotAdminError(raw []byte) (*NegRiskCtfExchangeNotAdmin, error) {
	out := new(NegRiskCtfExchangeNotAdmin)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "NotAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeNotCrossing represents a NotCrossing error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeNotCrossing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotCrossing()
func NegRiskCtfExchangeNotCrossingErrorID() common.Hash {
	return common.HexToHash("0x7f9a6f46e6433ed6106ae4041009fc892baea5caf8b41cb0b98399829566499b")
}

// UnpackNotCrossingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotCrossing()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackNotCrossingError(raw []byte) (*NegRiskCtfExchangeNotCrossing, error) {
	out := new(NegRiskCtfExchangeNotCrossing)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "NotCrossing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeNotOperator represents a NotOperator error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeNotOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOperator()
func NegRiskCtfExchangeNotOperatorErrorID() common.Hash {
	return common.HexToHash("0x7c214f0474e418b49066b523b4b4fdaaeed28f6f33b3f87920010fb4856b27c9")
}

// UnpackNotOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOperator()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackNotOperatorError(raw []byte) (*NegRiskCtfExchangeNotOperator, error) {
	out := new(NegRiskCtfExchangeNotOperator)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "NotOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeNotOwner represents a NotOwner error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeNotOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwner()
func NegRiskCtfExchangeNotOwnerErrorID() common.Hash {
	return common.HexToHash("0x30cd74712f59d478562d48e2d35de830db72c60a63dd08ae59199eec990b5bc4")
}

// UnpackNotOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwner()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackNotOwnerError(raw []byte) (*NegRiskCtfExchangeNotOwner, error) {
	out := new(NegRiskCtfExchangeNotOwner)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "NotOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeNotTaker represents a NotTaker error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeNotTaker struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotTaker()
func NegRiskCtfExchangeNotTakerErrorID() common.Hash {
	return common.HexToHash("0x5211a0797747df48df7596e18a47517c59d077b8bae8a2bef7dabd4a828e8071")
}

// UnpackNotTakerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotTaker()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackNotTakerError(raw []byte) (*NegRiskCtfExchangeNotTaker, error) {
	out := new(NegRiskCtfExchangeNotTaker)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "NotTaker", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeOrderExpired represents a OrderExpired error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeOrderExpired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OrderExpired()
func NegRiskCtfExchangeOrderExpiredErrorID() common.Hash {
	return common.HexToHash("0xc56873bac2ec2d1dd6d40a2a9b432692a2b77a4253c07d6f7ef4929f60ced89c")
}

// UnpackOrderExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OrderExpired()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackOrderExpiredError(raw []byte) (*NegRiskCtfExchangeOrderExpired, error) {
	out := new(NegRiskCtfExchangeOrderExpired)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "OrderExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeOrderFilledOrCancelled represents a OrderFilledOrCancelled error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeOrderFilledOrCancelled struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OrderFilledOrCancelled()
func NegRiskCtfExchangeOrderFilledOrCancelledErrorID() common.Hash {
	return common.HexToHash("0x7b38b76e3ca626669502f5b0d7687400e4377e3c8ac1bcca660cf6c4e81ed306")
}

// UnpackOrderFilledOrCancelledError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OrderFilledOrCancelled()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackOrderFilledOrCancelledError(raw []byte) (*NegRiskCtfExchangeOrderFilledOrCancelled, error) {
	out := new(NegRiskCtfExchangeOrderFilledOrCancelled)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "OrderFilledOrCancelled", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangePaused represents a Paused error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangePaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Paused()
func NegRiskCtfExchangePausedErrorID() common.Hash {
	return common.HexToHash("0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752")
}

// UnpackPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Paused()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackPausedError(raw []byte) (*NegRiskCtfExchangePaused, error) {
	out := new(NegRiskCtfExchangePaused)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "Paused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskCtfExchangeTooLittleTokensReceived represents a TooLittleTokensReceived error raised by the NegRiskCtfExchange contract.
type NegRiskCtfExchangeTooLittleTokensReceived struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TooLittleTokensReceived()
func NegRiskCtfExchangeTooLittleTokensReceivedErrorID() common.Hash {
	return common.HexToHash("0xdf4d80806ef3806591828d768fc0beb8a2f7fb5d5e9bab7b2b3afab395389c65")
}

// UnpackTooLittleTokensReceivedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TooLittleTokensReceived()
func (negRiskCtfExchange *NegRiskCtfExchange) UnpackTooLittleTokensReceivedError(raw []byte) (*NegRiskCtfExchangeTooLittleTokensReceived, error) {
	out := new(NegRiskCtfExchangeTooLittleTokensReceived)
	if err := negRiskCtfExchange.abi.UnpackIntoInterface(out, "TooLittleTokensReceived", raw); err != nil {
		return nil, err
	}
	return out, nil
}

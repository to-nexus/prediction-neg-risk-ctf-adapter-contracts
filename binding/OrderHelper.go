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

// OrderHelperMetaData contains all meta data concerning the OrderHelper contract.
var OrderHelperMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IS_SCRIPT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "OrderHelper",
	Bin: "0x6080604052600c805462ff00ff191662010001179055348015601f575f5ffd5b50608680602b5f395ff3fe6080604052348015600e575f5ffd5b50600436106026575f3560e01c8063f8ccbf4714602a575b5f5ffd5b600c54603c9062010000900460ff1681565b604051901515815260200160405180910390f3fea264697066735822122052ce93845a7475210f31efe0102dfdfacea6731f1d2b0281a5ee8d236d97a9a764736f6c634300081e0033",
}

// OrderHelper is an auto generated Go binding around an Ethereum contract.
type OrderHelper struct {
	abi abi.ABI
}

// NewOrderHelper creates a new instance of OrderHelper.
func NewOrderHelper() *OrderHelper {
	parsed, err := OrderHelperMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &OrderHelper{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *OrderHelper) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackISSCRIPT is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8ccbf47.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (orderHelper *OrderHelper) PackISSCRIPT() []byte {
	enc, err := orderHelper.abi.Pack("IS_SCRIPT")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackISSCRIPT is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8ccbf47.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (orderHelper *OrderHelper) TryPackISSCRIPT() ([]byte, error) {
	return orderHelper.abi.Pack("IS_SCRIPT")
}

// UnpackISSCRIPT is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (orderHelper *OrderHelper) UnpackISSCRIPT(data []byte) (bool, error) {
	out, err := orderHelper.abi.Unpack("IS_SCRIPT", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

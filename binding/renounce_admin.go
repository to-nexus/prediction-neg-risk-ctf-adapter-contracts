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

// RenounceAdminMetaData contains all meta data concerning the RenounceAdmin contract.
var RenounceAdminMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"run\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "RenounceAdmin",
	Bin: "0x6080604052348015600e575f5ffd5b506101848061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063522bb7041461002d575b5f5ffd5b61004061003b366004610114565b610042565b005b5f819050737109709ecfa91a80626ff3989d68f67f5b1dd12d73ffffffffffffffffffffffffffffffffffffffff1663afc980406040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561009f575f5ffd5b505af11580156100b1573d5f5f3e3d5ffd5b505050508073ffffffffffffffffffffffffffffffffffffffff16638bad0c0a6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156100fa575f5ffd5b505af115801561010c573d5f5f3e3d5ffd5b505050505050565b5f60208284031215610124575f5ffd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610147575f5ffd5b939250505056fea26469706673582212209fb43c74d49bb636503232bdfe07c1fb4144c11b350339be84c3c3c47d6ff3b364736f6c634300081e0033",
}

// RenounceAdmin is an auto generated Go binding around an Ethereum contract.
type RenounceAdmin struct {
	abi abi.ABI
}

// NewRenounceAdmin creates a new instance of RenounceAdmin.
func NewRenounceAdmin() *RenounceAdmin {
	parsed, err := RenounceAdminMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RenounceAdmin{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RenounceAdmin) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackRun is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x522bb704.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function run(address _target) returns()
func (renounceAdmin *RenounceAdmin) PackRun(target common.Address) []byte {
	enc, err := renounceAdmin.abi.Pack("run", target)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRun is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x522bb704.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function run(address _target) returns()
func (renounceAdmin *RenounceAdmin) TryPackRun(target common.Address) ([]byte, error) {
	return renounceAdmin.abi.Pack("run", target)
}

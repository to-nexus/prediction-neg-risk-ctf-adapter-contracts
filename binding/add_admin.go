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

// AddAdminMetaData contains all meta data concerning the AddAdmin contract.
var AddAdminMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"run\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "AddAdmin",
	Bin: "0x6080604052348015600e575f5ffd5b506101c58061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063fc4dcacb1461002d575b5f5ffd5b61004061003b36600461015e565b610042565b005b5f829050737109709ecfa91a80626ff3989d68f67f5b1dd12d73ffffffffffffffffffffffffffffffffffffffff1663afc980406040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561009f575f5ffd5b505af11580156100b1573d5f5f3e3d5ffd5b50506040517f7048027500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015284169250637048027591506024015f604051808303815f87803b15801561011b575f5ffd5b505af115801561012d573d5f5f3e3d5ffd5b50505050505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610159575f5ffd5b919050565b5f5f6040838503121561016f575f5ffd5b61017883610136565b915061018660208401610136565b9050925092905056fea2646970667358221220960a44f593fff74ec2439d853b632a216234430315a2ca2f27a26fb6f889401364736f6c634300081e0033",
}

// AddAdmin is an auto generated Go binding around an Ethereum contract.
type AddAdmin struct {
	abi abi.ABI
}

// NewAddAdmin creates a new instance of AddAdmin.
func NewAddAdmin() *AddAdmin {
	parsed, err := AddAdminMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AddAdmin{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AddAdmin) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackRun is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc4dcacb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function run(address _target, address _newAdmin) returns()
func (addAdmin *AddAdmin) PackRun(target common.Address, newAdmin common.Address) []byte {
	enc, err := addAdmin.abi.Pack("run", target, newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRun is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc4dcacb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function run(address _target, address _newAdmin) returns()
func (addAdmin *AddAdmin) TryPackRun(target common.Address, newAdmin common.Address) ([]byte, error) {
	return addAdmin.abi.Pack("run", target, newAdmin)
}

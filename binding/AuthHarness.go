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

// AuthHarnessMetaData contains all meta data concerning the AuthHarness contract.
var AuthHarnessMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useOnlyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"}]",
	ID:  "AuthHarness",
	Bin: "0x6080604052348015600e575f5ffd5b50335f908152602081905260409020600190556103948061002e5f395ff3fe608060405234801561000f575f5ffd5b506004361061006f575f3560e01c8063704802751161004d57806370480275146101025780637ff56f22146101155780638bad0c0a1461011d575f5ffd5b80631785f53c1461007357806324d7806c14610088578063429b62e5146100d5575b5f5ffd5b610086610081366004610324565b610125565b005b6100c0610096366004610324565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205460011490565b60405190151581526020015b60405180910390f35b6100f46100e3366004610324565b5f6020819052908152604090205481565b6040519081526020016100cc565b610086610110366004610324565b6101be565b610086610258565b6100866102a2565b335f9081526020819052604090205460011461016d576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f818152602081905260408082208290555133917f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e91a350565b335f90815260208190526040902054600114610206576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f81815260208190526040808220600190555133917ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc91a350565b335f908152602081905260409020546001146102a0576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b335f908152602081905260409020546001146102ea576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f818152602081905260408082208290555182917f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e91a3565b5f60208284031215610334575f5ffd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610357575f5ffd5b939250505056fea2646970667358221220ec0bca1549a2127e06eb2d7978132d5a027bb3f87df474f759379389646dec3464736f6c634300081e0033",
}

// AuthHarness is an auto generated Go binding around an Ethereum contract.
type AuthHarness struct {
	abi abi.ABI
}

// NewAuthHarness creates a new instance of AuthHarness.
func NewAuthHarness() *AuthHarness {
	parsed, err := AuthHarnessMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AuthHarness{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AuthHarness) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70480275.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addAdmin(address admin) returns()
func (authHarness *AuthHarness) PackAddAdmin(admin common.Address) []byte {
	enc, err := authHarness.abi.Pack("addAdmin", admin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70480275.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addAdmin(address admin) returns()
func (authHarness *AuthHarness) TryPackAddAdmin(admin common.Address) ([]byte, error) {
	return authHarness.abi.Pack("addAdmin", admin)
}

// PackAdmins is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x429b62e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function admins(address ) view returns(uint256)
func (authHarness *AuthHarness) PackAdmins(arg0 common.Address) []byte {
	enc, err := authHarness.abi.Pack("admins", arg0)
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
func (authHarness *AuthHarness) TryPackAdmins(arg0 common.Address) ([]byte, error) {
	return authHarness.abi.Pack("admins", arg0)
}

// UnpackAdmins is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (authHarness *AuthHarness) UnpackAdmins(data []byte) (*big.Int, error) {
	out, err := authHarness.abi.Unpack("admins", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackIsAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24d7806c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (authHarness *AuthHarness) PackIsAdmin(addr common.Address) []byte {
	enc, err := authHarness.abi.Pack("isAdmin", addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24d7806c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (authHarness *AuthHarness) TryPackIsAdmin(addr common.Address) ([]byte, error) {
	return authHarness.abi.Pack("isAdmin", addr)
}

// UnpackIsAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (authHarness *AuthHarness) UnpackIsAdmin(data []byte) (bool, error) {
	out, err := authHarness.abi.Unpack("isAdmin", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRemoveAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1785f53c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeAdmin(address admin) returns()
func (authHarness *AuthHarness) PackRemoveAdmin(admin common.Address) []byte {
	enc, err := authHarness.abi.Pack("removeAdmin", admin)
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
func (authHarness *AuthHarness) TryPackRemoveAdmin(admin common.Address) ([]byte, error) {
	return authHarness.abi.Pack("removeAdmin", admin)
}

// PackRenounceAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8bad0c0a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceAdmin() returns()
func (authHarness *AuthHarness) PackRenounceAdmin() []byte {
	enc, err := authHarness.abi.Pack("renounceAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8bad0c0a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceAdmin() returns()
func (authHarness *AuthHarness) TryPackRenounceAdmin() ([]byte, error) {
	return authHarness.abi.Pack("renounceAdmin")
}

// PackUseOnlyAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ff56f22.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function useOnlyAdmin() returns()
func (authHarness *AuthHarness) PackUseOnlyAdmin() []byte {
	enc, err := authHarness.abi.Pack("useOnlyAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUseOnlyAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ff56f22.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function useOnlyAdmin() returns()
func (authHarness *AuthHarness) TryPackUseOnlyAdmin() ([]byte, error) {
	return authHarness.abi.Pack("useOnlyAdmin")
}

// AuthHarnessNewAdmin represents a NewAdmin event raised by the AuthHarness contract.
type AuthHarnessNewAdmin struct {
	Admin           common.Address
	NewAdminAddress common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const AuthHarnessNewAdminEventName = "NewAdmin"

// ContractEventName returns the user-defined event name.
func (AuthHarnessNewAdmin) ContractEventName() string {
	return AuthHarnessNewAdminEventName
}

// UnpackNewAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (authHarness *AuthHarness) UnpackNewAdminEvent(log *types.Log) (*AuthHarnessNewAdmin, error) {
	event := "NewAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != authHarness.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthHarnessNewAdmin)
	if len(log.Data) > 0 {
		if err := authHarness.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authHarness.abi.Events[event].Inputs {
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

// AuthHarnessRemovedAdmin represents a RemovedAdmin event raised by the AuthHarness contract.
type AuthHarnessRemovedAdmin struct {
	Admin        common.Address
	RemovedAdmin common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const AuthHarnessRemovedAdminEventName = "RemovedAdmin"

// ContractEventName returns the user-defined event name.
func (AuthHarnessRemovedAdmin) ContractEventName() string {
	return AuthHarnessRemovedAdminEventName
}

// UnpackRemovedAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (authHarness *AuthHarness) UnpackRemovedAdminEvent(log *types.Log) (*AuthHarnessRemovedAdmin, error) {
	event := "RemovedAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != authHarness.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthHarnessRemovedAdmin)
	if len(log.Data) > 0 {
		if err := authHarness.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authHarness.abi.Events[event].Inputs {
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
func (authHarness *AuthHarness) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], authHarness.abi.Errors["NotAdmin"].ID.Bytes()[:4]) {
		return authHarness.UnpackNotAdminError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// AuthHarnessNotAdmin represents a NotAdmin error raised by the AuthHarness contract.
type AuthHarnessNotAdmin struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAdmin()
func AuthHarnessNotAdminErrorID() common.Hash {
	return common.HexToHash("0x7bfa4b9fb0cd3687c1f539f384b3f3f258f2c9aa9186353d0815413b508ed97d")
}

// UnpackNotAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAdmin()
func (authHarness *AuthHarness) UnpackNotAdminError(raw []byte) (*AuthHarnessNotAdmin, error) {
	out := new(AuthHarnessNotAdmin)
	if err := authHarness.abi.UnpackIntoInterface(out, "NotAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

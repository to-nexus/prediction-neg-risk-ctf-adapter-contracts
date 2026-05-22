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

// VaultMetaData contains all meta data concerning the Vault contract.
var VaultMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc1155\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc1155\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"}]",
	ID:  "Vault",
	Bin: "0x6080604052348015600e575f5ffd5b50335f90815260208190526040902060019055610af38061002e5f395ff3fe608060405234801561000f575f5ffd5b50600436106100b9575f3560e01c80638bad0c0a11610072578063a9412a5911610058578063a9412a591461018d578063bc197c81146101a0578063f23a6e611461020c575f5ffd5b80638bad0c0a146101725780639db5dbe41461017a575f5ffd5b806324d7806c116100a257806324d7806c146100e5578063429b62e514610132578063704802751461015f575f5ffd5b80630a7e880c146100bd5780631785f53c146100d2575b5f5ffd5b6100d06100cb366004610729565b610245565b005b6100d06100e0366004610768565b610330565b61011d6100f3366004610768565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205460011490565b60405190151581526020015b60405180910390f35b610151610140366004610768565b5f6020819052908152604090205481565b604051908152602001610129565b6100d061016d366004610768565b6103c9565b6100d0610463565b6100d0610188366004610788565b6104e5565b6100d061019b36600461080a565b610553565b6101db6101ae3660046108d8565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091168152602001610129565b6101db61021a366004610997565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b335f9081526020819052604090205460011461028d576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517ff242432a00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8481166024830152604482018490526064820183905260a060848301525f60a483015285169063f242432a9060c4015f604051808303815f87803b158015610314575f5ffd5b505af1158015610326573d5f5f3e3d5ffd5b5050505050505050565b335f90815260208190526040902054600114610378576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f818152602081905260408082208290555133917f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e91a350565b335f90815260208190526040902054600114610411576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f81815260208190526040808220600190555133917ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc91a350565b335f908152602081905260409020546001146104ab576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f818152602081905260408082208290555182917f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e91a3565b335f9081526020819052604090205460011461052d576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61054e73ffffffffffffffffffffffffffffffffffffffff8416838361062c565b505050565b335f9081526020819052604090205460011461059b576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f2eb2c2d600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff871690632eb2c2d6906105f790309089908990899089908990600401610a41565b5f604051808303815f87803b15801561060e575f5ffd5b505af1158015610620573d5f5f3e3d5ffd5b50505050505050505050565b5f6040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015282602482015260205f6044835f895af13d15601f3d1160015f5114161716915050806106fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f5452414e534645525f4641494c45440000000000000000000000000000000000604482015260640160405180910390fd5b50505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610724575f5ffd5b919050565b5f5f5f5f6080858703121561073c575f5ffd5b61074585610701565b935061075360208601610701565b93969395505050506040820135916060013590565b5f60208284031215610778575f5ffd5b61078182610701565b9392505050565b5f5f5f6060848603121561079a575f5ffd5b6107a384610701565b92506107b160208501610701565b929592945050506040919091013590565b5f5f83601f8401126107d2575f5ffd5b50813567ffffffffffffffff8111156107e9575f5ffd5b6020830191508360208260051b8501011115610803575f5ffd5b9250929050565b5f5f5f5f5f5f6080878903121561081f575f5ffd5b61082887610701565b955061083660208801610701565b9450604087013567ffffffffffffffff811115610851575f5ffd5b61085d89828a016107c2565b909550935050606087013567ffffffffffffffff81111561087c575f5ffd5b61088889828a016107c2565b979a9699509497509295939492505050565b5f5f83601f8401126108aa575f5ffd5b50813567ffffffffffffffff8111156108c1575f5ffd5b602083019150836020828501011115610803575f5ffd5b5f5f5f5f5f5f5f5f60a0898b0312156108ef575f5ffd5b6108f889610701565b975061090660208a01610701565b9650604089013567ffffffffffffffff811115610921575f5ffd5b61092d8b828c016107c2565b909750955050606089013567ffffffffffffffff81111561094c575f5ffd5b6109588b828c016107c2565b909550935050608089013567ffffffffffffffff811115610977575f5ffd5b6109838b828c0161089a565b999c989b5096995094979396929594505050565b5f5f5f5f5f5f60a087890312156109ac575f5ffd5b6109b587610701565b95506109c360208801610701565b94506040870135935060608701359250608087013567ffffffffffffffff8111156109ec575f5ffd5b61088889828a0161089a565b8183525f7f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115610a28575f5ffd5b8260051b80836020870137939093016020019392505050565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015260a060408201525f610a8c60a0830186886109f8565b8281036060840152610a9f8185876109f8565b83810360809094019390935250505f8152602001969550505050505056fea2646970667358221220eddcc16a87310725a9d62bc7385accea64d208941baea9929c0132bb116699e164736f6c634300081e0033",
}

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	abi abi.ABI
}

// NewVault creates a new instance of Vault.
func NewVault() *Vault {
	parsed, err := VaultMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Vault{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Vault) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70480275.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addAdmin(address admin) returns()
func (vault *Vault) PackAddAdmin(admin common.Address) []byte {
	enc, err := vault.abi.Pack("addAdmin", admin)
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
func (vault *Vault) TryPackAddAdmin(admin common.Address) ([]byte, error) {
	return vault.abi.Pack("addAdmin", admin)
}

// PackAdmins is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x429b62e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function admins(address ) view returns(uint256)
func (vault *Vault) PackAdmins(arg0 common.Address) []byte {
	enc, err := vault.abi.Pack("admins", arg0)
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
func (vault *Vault) TryPackAdmins(arg0 common.Address) ([]byte, error) {
	return vault.abi.Pack("admins", arg0)
}

// UnpackAdmins is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (vault *Vault) UnpackAdmins(data []byte) (*big.Int, error) {
	out, err := vault.abi.Unpack("admins", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackBatchTransferERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9412a59.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function batchTransferERC1155(address _erc1155, address _to, uint256[] _ids, uint256[] _values) returns()
func (vault *Vault) PackBatchTransferERC1155(erc1155 common.Address, to common.Address, ids []*big.Int, values []*big.Int) []byte {
	enc, err := vault.abi.Pack("batchTransferERC1155", erc1155, to, ids, values)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBatchTransferERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9412a59.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function batchTransferERC1155(address _erc1155, address _to, uint256[] _ids, uint256[] _values) returns()
func (vault *Vault) TryPackBatchTransferERC1155(erc1155 common.Address, to common.Address, ids []*big.Int, values []*big.Int) ([]byte, error) {
	return vault.abi.Pack("batchTransferERC1155", erc1155, to, ids, values)
}

// PackIsAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24d7806c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (vault *Vault) PackIsAdmin(addr common.Address) []byte {
	enc, err := vault.abi.Pack("isAdmin", addr)
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
func (vault *Vault) TryPackIsAdmin(addr common.Address) ([]byte, error) {
	return vault.abi.Pack("isAdmin", addr)
}

// UnpackIsAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (vault *Vault) UnpackIsAdmin(data []byte) (bool, error) {
	out, err := vault.abi.Unpack("isAdmin", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (vault *Vault) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) []byte {
	enc, err := vault.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
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
func (vault *Vault) TryPackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([]byte, error) {
	return vault.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (vault *Vault) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := vault.abi.Unpack("onERC1155BatchReceived", data)
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
func (vault *Vault) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) []byte {
	enc, err := vault.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
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
func (vault *Vault) TryPackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return vault.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (vault *Vault) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := vault.abi.Unpack("onERC1155Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackRemoveAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1785f53c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeAdmin(address admin) returns()
func (vault *Vault) PackRemoveAdmin(admin common.Address) []byte {
	enc, err := vault.abi.Pack("removeAdmin", admin)
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
func (vault *Vault) TryPackRemoveAdmin(admin common.Address) ([]byte, error) {
	return vault.abi.Pack("removeAdmin", admin)
}

// PackRenounceAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8bad0c0a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceAdmin() returns()
func (vault *Vault) PackRenounceAdmin() []byte {
	enc, err := vault.abi.Pack("renounceAdmin")
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
func (vault *Vault) TryPackRenounceAdmin() ([]byte, error) {
	return vault.abi.Pack("renounceAdmin")
}

// PackTransferERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a7e880c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferERC1155(address _erc1155, address _to, uint256 _id, uint256 _value) returns()
func (vault *Vault) PackTransferERC1155(erc1155 common.Address, to common.Address, id *big.Int, value *big.Int) []byte {
	enc, err := vault.abi.Pack("transferERC1155", erc1155, to, id, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferERC1155 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a7e880c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferERC1155(address _erc1155, address _to, uint256 _id, uint256 _value) returns()
func (vault *Vault) TryPackTransferERC1155(erc1155 common.Address, to common.Address, id *big.Int, value *big.Int) ([]byte, error) {
	return vault.abi.Pack("transferERC1155", erc1155, to, id, value)
}

// PackTransferERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9db5dbe4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferERC20(address _erc20, address _to, uint256 _amount) returns()
func (vault *Vault) PackTransferERC20(erc20 common.Address, to common.Address, amount *big.Int) []byte {
	enc, err := vault.abi.Pack("transferERC20", erc20, to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9db5dbe4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferERC20(address _erc20, address _to, uint256 _amount) returns()
func (vault *Vault) TryPackTransferERC20(erc20 common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	return vault.abi.Pack("transferERC20", erc20, to, amount)
}

// VaultNewAdmin represents a NewAdmin event raised by the Vault contract.
type VaultNewAdmin struct {
	Admin           common.Address
	NewAdminAddress common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const VaultNewAdminEventName = "NewAdmin"

// ContractEventName returns the user-defined event name.
func (VaultNewAdmin) ContractEventName() string {
	return VaultNewAdminEventName
}

// UnpackNewAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (vault *Vault) UnpackNewAdminEvent(log *types.Log) (*VaultNewAdmin, error) {
	event := "NewAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != vault.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VaultNewAdmin)
	if len(log.Data) > 0 {
		if err := vault.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vault.abi.Events[event].Inputs {
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

// VaultRemovedAdmin represents a RemovedAdmin event raised by the Vault contract.
type VaultRemovedAdmin struct {
	Admin        common.Address
	RemovedAdmin common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const VaultRemovedAdminEventName = "RemovedAdmin"

// ContractEventName returns the user-defined event name.
func (VaultRemovedAdmin) ContractEventName() string {
	return VaultRemovedAdminEventName
}

// UnpackRemovedAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (vault *Vault) UnpackRemovedAdminEvent(log *types.Log) (*VaultRemovedAdmin, error) {
	event := "RemovedAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != vault.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VaultRemovedAdmin)
	if len(log.Data) > 0 {
		if err := vault.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vault.abi.Events[event].Inputs {
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
func (vault *Vault) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], vault.abi.Errors["NotAdmin"].ID.Bytes()[:4]) {
		return vault.UnpackNotAdminError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// VaultNotAdmin represents a NotAdmin error raised by the Vault contract.
type VaultNotAdmin struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAdmin()
func VaultNotAdminErrorID() common.Hash {
	return common.HexToHash("0x7bfa4b9fb0cd3687c1f539f384b3f3f258f2c9aa9186353d0815413b508ed97d")
}

// UnpackNotAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAdmin()
func (vault *Vault) UnpackNotAdminError(raw []byte) (*VaultNotAdmin, error) {
	out := new(VaultNotAdmin)
	if err := vault.abi.UnpackIntoInterface(out, "NotAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

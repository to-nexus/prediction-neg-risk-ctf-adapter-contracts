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

// USDCMetaData contains all meta data concerning the USDC contract.
var USDCMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]",
	ID:  "USDC",
	Bin: "0x60e060405234801561000f575f5ffd5b506040805180820182526004808252635553444360e01b60208084018290528451808601909552918452908301529060065f61004b84826101a8565b50600161005883826101a8565b5060ff81166080524660a05261006c610078565b60c052506102d3915050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f6040516100a89190610262565b6040805191829003822060208301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061013857607f821691505b60208210810361015657634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101a357805f5260205f20601f840160051c810160208510156101815750805b601f840160051c820191505b818110156101a0575f815560010161018d565b50505b505050565b81516001600160401b038111156101c1576101c1610110565b6101d5816101cf8454610124565b8461015c565b6020601f821160018114610207575f83156101f05750848201515b5f19600385901b1c1916600184901b1784556101a0565b5f84815260208120601f198516915b828110156102365787850151825560209485019460019092019101610216565b508482101561025357868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f5f835461026f81610124565b600182168015610286576001811461029b576102c8565b60ff19831686528115158202860193506102c8565b865f5260205f205f5b838110156102c0578154888201526001909101906020016102a4565b505081860193505b509195945050505050565b60805160a05160c051610d9a6102fd5f395f6104cf01525f61049a01525f6101590152610d9a5ff3fe608060405234801561000f575f5ffd5b50600436106100e5575f3560e01c806370a08231116100885780639dc29fac116100635780639dc29fac146101f0578063a9059cbb14610203578063d505accf14610216578063dd62ed3e14610229575f5ffd5b806370a08231146101aa5780637ecebe00146101c957806395d89b41146101e8575f5ffd5b806323b872dd116100c357806323b872dd14610141578063313ce567146101545780633644e5151461018d57806340c10f1914610195575f5ffd5b806306fdde03146100e9578063095ea7b31461010757806318160ddd1461012a575b5f5ffd5b6100f1610253565b6040516100fe9190610a52565b60405180910390f35b61011a610115366004610acd565b6102de565b60405190151581526020016100fe565b61013360025481565b6040519081526020016100fe565b61011a61014f366004610af5565b610357565b61017b7f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016100fe565b610133610497565b6101a86101a3366004610acd565b6104f1565b005b6101336101b8366004610b2f565b60036020525f908152604090205481565b6101336101d7366004610b2f565b60056020525f908152604090205481565b6100f16104ff565b6101a86101fe366004610acd565b61050c565b61011a610211366004610acd565b610516565b6101a8610224366004610b4f565b610599565b610133610237366004610bbc565b600460209081525f928352604080842090915290825290205481565b5f805461025f90610bed565b80601f016020809104026020016040519081016040528092919081815260200182805461028b90610bed565b80156102d65780601f106102ad576101008083540402835291602001916102d6565b820191905f5260205f20905b8154815290600101906020018083116102b957829003601f168201915b505050505081565b335f81815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103459086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526004602090815260408083203384529091528120547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103e9576103b88382610c6b565b73ffffffffffffffffffffffffffffffffffffffff86165f9081526004602090815260408083203384529091529020555b73ffffffffffffffffffffffffffffffffffffffff85165f908152600360205260408120805485929061041d908490610c6b565b909155505073ffffffffffffffffffffffffffffffffffffffff8085165f81815260036020526040908190208054870190555190918716907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906104849087815260200190565b60405180910390a3506001949350505050565b5f7f000000000000000000000000000000000000000000000000000000000000000046146104cc576104c76108b7565b905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b6104fb828261094f565b5050565b6001805461025f90610bed565b6104fb82826109c6565b335f90815260036020526040812080548391908390610536908490610c6b565b909155505073ffffffffffffffffffffffffffffffffffffffff83165f81815260036020526040908190208054850190555133907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906103459086815260200190565b42841015610608576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f5045524d49545f444541444c494e455f4558504952454400000000000000000060448201526064015b60405180910390fd5b5f6001610613610497565b73ffffffffffffffffffffffffffffffffffffffff8a81165f8181526005602090815260409182902080546001810190915582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98184015280840194909452938d166060840152608083018c905260a083019390935260c08083018b90528151808403909101815260e0830190915280519201919091207f190100000000000000000000000000000000000000000000000000000000000061010083015261010282019290925261012281019190915261014201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201205f84529083018083525260ff871690820152606081018590526080810184905260a0016020604051602081039080840390855afa158015610761573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116158015906107dc57508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610842576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f494e56414c49445f5349474e455200000000000000000000000000000000000060448201526064016105ff565b73ffffffffffffffffffffffffffffffffffffffff9081165f9081526004602090815260408083208a8516808552908352928190208990555188815291928a16917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a350505050505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f6040516108e79190610c7e565b6040805191829003822060208301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b8060025f8282546109609190610d51565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f818152600360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91015b60405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff82165f90815260036020526040812080548392906109fa908490610c6b565b90915550506002805482900390556040518181525f9073ffffffffffffffffffffffffffffffffffffffff8416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016109ba565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610ac8575f5ffd5b919050565b5f5f60408385031215610ade575f5ffd5b610ae783610aa5565b946020939093013593505050565b5f5f5f60608486031215610b07575f5ffd5b610b1084610aa5565b9250610b1e60208501610aa5565b929592945050506040919091013590565b5f60208284031215610b3f575f5ffd5b610b4882610aa5565b9392505050565b5f5f5f5f5f5f5f60e0888a031215610b65575f5ffd5b610b6e88610aa5565b9650610b7c60208901610aa5565b95506040880135945060608801359350608088013560ff81168114610b9f575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215610bcd575f5ffd5b610bd683610aa5565b9150610be460208401610aa5565b90509250929050565b600181811c90821680610c0157607f821691505b602082108103610c38577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561035157610351610c3e565b5f5f83545f8160011c90506001821680610c9957607f821691505b602082108103610cd0577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b808015610ce45760018114610d1757610d45565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0084168752821515830287019450610d45565b5f888152602090205f5b84811015610d3d57815489820152600190910190602001610d21565b505082870194505b50929695505050505050565b8082018082111561035157610351610c3e56fea2646970667358221220284dc38f3018f884bc0694d5809140cd0644f392049d0a1685d2a9db862cb5a964736f6c634300081e0033",
}

// USDC is an auto generated Go binding around an Ethereum contract.
type USDC struct {
	abi abi.ABI
}

// NewUSDC creates a new instance of USDC.
func NewUSDC() *USDC {
	parsed, err := USDCMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &USDC{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *USDC) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (uSDC *USDC) PackDOMAINSEPARATOR() []byte {
	enc, err := uSDC.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (uSDC *USDC) TryPackDOMAINSEPARATOR() ([]byte, error) {
	return uSDC.abi.Pack("DOMAIN_SEPARATOR")
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (uSDC *USDC) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := uSDC.abi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (uSDC *USDC) PackAllowance(arg0 common.Address, arg1 common.Address) []byte {
	enc, err := uSDC.abi.Pack("allowance", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (uSDC *USDC) TryPackAllowance(arg0 common.Address, arg1 common.Address) ([]byte, error) {
	return uSDC.abi.Pack("allowance", arg0, arg1)
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (uSDC *USDC) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := uSDC.abi.Unpack("allowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (uSDC *USDC) PackApprove(spender common.Address, amount *big.Int) []byte {
	enc, err := uSDC.abi.Pack("approve", spender, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (uSDC *USDC) TryPackApprove(spender common.Address, amount *big.Int) ([]byte, error) {
	return uSDC.abi.Pack("approve", spender, amount)
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (uSDC *USDC) UnpackApprove(data []byte) (bool, error) {
	out, err := uSDC.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (uSDC *USDC) PackBalanceOf(arg0 common.Address) []byte {
	enc, err := uSDC.abi.Pack("balanceOf", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (uSDC *USDC) TryPackBalanceOf(arg0 common.Address) ([]byte, error) {
	return uSDC.abi.Pack("balanceOf", arg0)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (uSDC *USDC) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := uSDC.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dc29fac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (uSDC *USDC) PackBurn(from common.Address, amount *big.Int) []byte {
	enc, err := uSDC.abi.Pack("burn", from, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dc29fac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (uSDC *USDC) TryPackBurn(from common.Address, amount *big.Int) ([]byte, error) {
	return uSDC.abi.Pack("burn", from, amount)
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function decimals() view returns(uint8)
func (uSDC *USDC) PackDecimals() []byte {
	enc, err := uSDC.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function decimals() view returns(uint8)
func (uSDC *USDC) TryPackDecimals() ([]byte, error) {
	return uSDC.abi.Pack("decimals")
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (uSDC *USDC) UnpackDecimals(data []byte) (uint8, error) {
	out, err := uSDC.abi.Unpack("decimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (uSDC *USDC) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := uSDC.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (uSDC *USDC) TryPackMint(to common.Address, amount *big.Int) ([]byte, error) {
	return uSDC.abi.Pack("mint", to, amount)
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function name() view returns(string)
func (uSDC *USDC) PackName() []byte {
	enc, err := uSDC.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function name() view returns(string)
func (uSDC *USDC) TryPackName() ([]byte, error) {
	return uSDC.abi.Pack("name")
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (uSDC *USDC) UnpackName(data []byte) (string, error) {
	out, err := uSDC.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nonces(address ) view returns(uint256)
func (uSDC *USDC) PackNonces(arg0 common.Address) []byte {
	enc, err := uSDC.abi.Pack("nonces", arg0)
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
func (uSDC *USDC) TryPackNonces(arg0 common.Address) ([]byte, error) {
	return uSDC.abi.Pack("nonces", arg0)
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (uSDC *USDC) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := uSDC.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd505accf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (uSDC *USDC) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := uSDC.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd505accf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (uSDC *USDC) TryPackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return uSDC.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function symbol() view returns(string)
func (uSDC *USDC) PackSymbol() []byte {
	enc, err := uSDC.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function symbol() view returns(string)
func (uSDC *USDC) TryPackSymbol() ([]byte, error) {
	return uSDC.abi.Pack("symbol")
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (uSDC *USDC) UnpackSymbol(data []byte) (string, error) {
	out, err := uSDC.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function totalSupply() view returns(uint256)
func (uSDC *USDC) PackTotalSupply() []byte {
	enc, err := uSDC.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function totalSupply() view returns(uint256)
func (uSDC *USDC) TryPackTotalSupply() ([]byte, error) {
	return uSDC.abi.Pack("totalSupply")
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (uSDC *USDC) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := uSDC.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (uSDC *USDC) PackTransfer(to common.Address, amount *big.Int) []byte {
	enc, err := uSDC.abi.Pack("transfer", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (uSDC *USDC) TryPackTransfer(to common.Address, amount *big.Int) ([]byte, error) {
	return uSDC.abi.Pack("transfer", to, amount)
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (uSDC *USDC) UnpackTransfer(data []byte) (bool, error) {
	out, err := uSDC.abi.Unpack("transfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (uSDC *USDC) PackTransferFrom(from common.Address, to common.Address, amount *big.Int) []byte {
	enc, err := uSDC.abi.Pack("transferFrom", from, to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (uSDC *USDC) TryPackTransferFrom(from common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	return uSDC.abi.Pack("transferFrom", from, to, amount)
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (uSDC *USDC) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := uSDC.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// USDCApproval represents a Approval event raised by the USDC contract.
type USDCApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const USDCApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (USDCApproval) ContractEventName() string {
	return USDCApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (uSDC *USDC) UnpackApprovalEvent(log *types.Log) (*USDCApproval, error) {
	event := "Approval"
	if len(log.Topics) == 0 || log.Topics[0] != uSDC.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(USDCApproval)
	if len(log.Data) > 0 {
		if err := uSDC.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range uSDC.abi.Events[event].Inputs {
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

// USDCTransfer represents a Transfer event raised by the USDC contract.
type USDCTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const USDCTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (USDCTransfer) ContractEventName() string {
	return USDCTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (uSDC *USDC) UnpackTransferEvent(log *types.Log) (*USDCTransfer, error) {
	event := "Transfer"
	if len(log.Topics) == 0 || log.Topics[0] != uSDC.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(USDCTransfer)
	if len(log.Data) > 0 {
		if err := uSDC.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range uSDC.abi.Events[event].Inputs {
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

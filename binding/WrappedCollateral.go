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

// WrappedCollateralMetaData contains all meta data concerning the WrappedCollateral contract.
var WrappedCollateralMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_underlying\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"}]",
	ID:  "WrappedCollateral",
	Bin: "0x610120604052348015610010575f5ffd5b506040516116a43803806116a483398101604081905261002f9161015c565b6040518060400160405280601281526020017115dc985c1c19590810dbdb1b185d195c985b60721b8152506040518060400160405280600481526020016315d0d3d360e21b81525082825f9081610086919061023c565b506001610093838261023c565b5060ff81166080524660a0526100a76100c4565b60c05250503360e05250506001600160a01b031661010052610367565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f6040516100f491906102f6565b6040805191829003822060208301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f5f6040838503121561016d575f5ffd5b82516001600160a01b0381168114610183575f5ffd5b602084015190925060ff81168114610199575f5ffd5b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806101cc57607f821691505b6020821081036101ea57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561023757805f5260205f20601f840160051c810160208510156102155750805b601f840160051c820191505b81811015610234575f8155600101610221565b50505b505050565b81516001600160401b03811115610255576102556101a4565b6102698161026384546101b8565b846101f0565b6020601f82116001811461029b575f83156102845750848201515b5f19600385901b1c1916600184901b178455610234565b5f84815260208120601f198516915b828110156102ca57878501518255602094850194600190920191016102aa565b50848210156102e757868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f5f8354610303816101b8565b60018216801561031a576001811461032f5761035c565b60ff198316865281151582028601935061035c565b865f5260205f205f5b8381101561035457815488820152600190910190602001610338565b505081860193505b509195945050505050565b60805160a05160c05160e051610100516112d66103ce5f395f8181610239015281816103ec01526108cd01525f81816102c30152818161037b015281816106d701528181610760015261085c01525f61069301525f61065e01525f6101d201526112d65ff3fe608060405234801561000f575f5ffd5b5060043610610149575f3560e01c80636f307dc3116100c7578063a0712d681161007d578063bf376c7a11610063578063bf376c7a14610313578063d505accf14610326578063dd62ed3e14610339575f5ffd5b8063a0712d68146102ed578063a9059cbb14610300575f5ffd5b80637ecebe00116100ad5780637ecebe001461029f5780638da5cb5b146102be57806395d89b41146102e5575f5ffd5b80636f307dc31461023457806370a0823114610280575f5ffd5b806323b872dd1161011c5780633644e515116101025780633644e5151461020657806339f476931461020e57806342966c6814610221575f5ffd5b806323b872dd146101ba578063313ce567146101cd575f5ffd5b80630357371d1461014d57806306fdde0314610162578063095ea7b31461018057806318160ddd146101a3575b5f5ffd5b61016061015b366004610f9f565b610363565b005b61016a610417565b6040516101779190610fc7565b60405180910390f35b61019361018e366004610f9f565b6104a2565b6040519015158152602001610177565b6101ac60025481565b604051908152602001610177565b6101936101c836600461101a565b61051b565b6101f47f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff9091168152602001610177565b6101ac61065b565b61016061021c366004610f9f565b6106b5565b61016061022f366004611054565b6106bf565b61025b7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610177565b6101ac61028e36600461106b565b60036020525f908152604090205481565b6101ac6102ad36600461106b565b60056020525f908152604090205481565b61025b7f000000000000000000000000000000000000000000000000000000000000000081565b61016a61073b565b6101606102fb366004611054565b610748565b61019361030e366004610f9f565b6107c1565b610160610321366004610f9f565b610844565b61016061033436600461108b565b6108ff565b6101ac6103473660046110f8565b600460209081525f928352604080842090915290825290205481565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146103d2576040517f5fc483c500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61041373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168383610c1d565b5050565b5f805461042390611129565b80601f016020809104026020016040519081016040528092919081815260200182805461044f90611129565b801561049a5780601f106104715761010080835404028352916020019161049a565b820191905f5260205f20905b81548152906001019060200180831161047d57829003601f168201915b505050505081565b335f81815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906105099086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526004602090815260408083203384529091528120547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146105ad5761057c83826111a7565b73ffffffffffffffffffffffffffffffffffffffff86165f9081526004602090815260408083203384529091529020555b73ffffffffffffffffffffffffffffffffffffffff85165f90815260036020526040812080548592906105e19084906111a7565b909155505073ffffffffffffffffffffffffffffffffffffffff8085165f81815260036020526040908190208054870190555190918716907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906106489087815260200190565b60405180910390a3506001949350505050565b5f7f000000000000000000000000000000000000000000000000000000000000000046146106905761068b610cee565b905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b6103d23382610d86565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461072e576040517f5fc483c500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107383382610d86565b50565b6001805461042390611129565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146107b7576040517f5fc483c500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107383382610e1a565b335f908152600360205260408120805483919083906107e19084906111a7565b909155505073ffffffffffffffffffffffffffffffffffffffff83165f81815260036020526040908190208054850190555133907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906105099086815260200190565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146108b3576040517f5fc483c500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108f573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084610e89565b6104138282610e1a565b4284101561096e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f5045524d49545f444541444c494e455f4558504952454400000000000000000060448201526064015b60405180910390fd5b5f600161097961065b565b73ffffffffffffffffffffffffffffffffffffffff8a81165f8181526005602090815260409182902080546001810190915582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98184015280840194909452938d166060840152608083018c905260a083019390935260c08083018b90528151808403909101815260e0830190915280519201919091207f190100000000000000000000000000000000000000000000000000000000000061010083015261010282019290925261012281019190915261014201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201205f84529083018083525260ff871690820152606081018590526080810184905260a0016020604051602081039080840390855afa158015610ac7573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811615801590610b4257508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610ba8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f494e56414c49445f5349474e45520000000000000000000000000000000000006044820152606401610965565b73ffffffffffffffffffffffffffffffffffffffff9081165f9081526004602090815260408083208a8516808552908352928190208990555188815291928a16917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a350505050505050565b5f6040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015282602482015260205f6044835f895af13d15601f3d1160015f511416171691505080610ce8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f5452414e534645525f4641494c454400000000000000000000000000000000006044820152606401610965565b50505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f604051610d1e91906111ba565b6040805191829003822060208301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526003602052604081208054839290610dba9084906111a7565b90915550506002805482900390556040518181525f9073ffffffffffffffffffffffffffffffffffffffff8416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020015b60405180910390a35050565b8060025f828254610e2b919061128d565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f818152600360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610e0e565b5f6040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8416602482015282604482015260205f6064835f8a5af13d15601f3d1160015f511416171691505080610f70576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5452414e534645525f46524f4d5f4641494c45440000000000000000000000006044820152606401610965565b5050505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610f9a575f5ffd5b919050565b5f5f60408385031215610fb0575f5ffd5b610fb983610f77565b946020939093013593505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f5f6060848603121561102c575f5ffd5b61103584610f77565b925061104360208501610f77565b929592945050506040919091013590565b5f60208284031215611064575f5ffd5b5035919050565b5f6020828403121561107b575f5ffd5b61108482610f77565b9392505050565b5f5f5f5f5f5f5f60e0888a0312156110a1575f5ffd5b6110aa88610f77565b96506110b860208901610f77565b95506040880135945060608801359350608088013560ff811681146110db575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215611109575f5ffd5b61111283610f77565b915061112060208401610f77565b90509250929050565b600181811c9082168061113d57607f821691505b602082108103611174577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156105155761051561117a565b5f5f83545f8160011c905060018216806111d557607f821691505b60208210810361120c577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b808015611220576001811461125357611281565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0084168752821515830287019450611281565b5f888152602090205f5b848110156112795781548982015260019091019060200161125d565b505082870194505b50929695505050505050565b808201808211156105155761051561117a56fea2646970667358221220a546b2937d4c2dc7e92647d917bd7ff82a139f878d1d30a3ca27bfb5690f78fc64736f6c634300081e0033",
}

// WrappedCollateral is an auto generated Go binding around an Ethereum contract.
type WrappedCollateral struct {
	abi abi.ABI
}

// NewWrappedCollateral creates a new instance of WrappedCollateral.
func NewWrappedCollateral() *WrappedCollateral {
	parsed, err := WrappedCollateralMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &WrappedCollateral{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *WrappedCollateral) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _underlying, uint8 _decimals) returns()
func (wrappedCollateral *WrappedCollateral) PackConstructor(_underlying common.Address, _decimals uint8) []byte {
	enc, err := wrappedCollateral.abi.Pack("", _underlying, _decimals)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (wrappedCollateral *WrappedCollateral) PackDOMAINSEPARATOR() []byte {
	enc, err := wrappedCollateral.abi.Pack("DOMAIN_SEPARATOR")
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
func (wrappedCollateral *WrappedCollateral) TryPackDOMAINSEPARATOR() ([]byte, error) {
	return wrappedCollateral.abi.Pack("DOMAIN_SEPARATOR")
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (wrappedCollateral *WrappedCollateral) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := wrappedCollateral.abi.Unpack("DOMAIN_SEPARATOR", data)
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
func (wrappedCollateral *WrappedCollateral) PackAllowance(arg0 common.Address, arg1 common.Address) []byte {
	enc, err := wrappedCollateral.abi.Pack("allowance", arg0, arg1)
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
func (wrappedCollateral *WrappedCollateral) TryPackAllowance(arg0 common.Address, arg1 common.Address) ([]byte, error) {
	return wrappedCollateral.abi.Pack("allowance", arg0, arg1)
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (wrappedCollateral *WrappedCollateral) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := wrappedCollateral.abi.Unpack("allowance", data)
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
func (wrappedCollateral *WrappedCollateral) PackApprove(spender common.Address, amount *big.Int) []byte {
	enc, err := wrappedCollateral.abi.Pack("approve", spender, amount)
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
func (wrappedCollateral *WrappedCollateral) TryPackApprove(spender common.Address, amount *big.Int) ([]byte, error) {
	return wrappedCollateral.abi.Pack("approve", spender, amount)
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (wrappedCollateral *WrappedCollateral) UnpackApprove(data []byte) (bool, error) {
	out, err := wrappedCollateral.abi.Unpack("approve", data)
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
func (wrappedCollateral *WrappedCollateral) PackBalanceOf(arg0 common.Address) []byte {
	enc, err := wrappedCollateral.abi.Pack("balanceOf", arg0)
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
func (wrappedCollateral *WrappedCollateral) TryPackBalanceOf(arg0 common.Address) ([]byte, error) {
	return wrappedCollateral.abi.Pack("balanceOf", arg0)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (wrappedCollateral *WrappedCollateral) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := wrappedCollateral.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42966c68.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burn(uint256 _amount) returns()
func (wrappedCollateral *WrappedCollateral) PackBurn(amount *big.Int) []byte {
	enc, err := wrappedCollateral.abi.Pack("burn", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42966c68.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burn(uint256 _amount) returns()
func (wrappedCollateral *WrappedCollateral) TryPackBurn(amount *big.Int) ([]byte, error) {
	return wrappedCollateral.abi.Pack("burn", amount)
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function decimals() view returns(uint8)
func (wrappedCollateral *WrappedCollateral) PackDecimals() []byte {
	enc, err := wrappedCollateral.abi.Pack("decimals")
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
func (wrappedCollateral *WrappedCollateral) TryPackDecimals() ([]byte, error) {
	return wrappedCollateral.abi.Pack("decimals")
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (wrappedCollateral *WrappedCollateral) UnpackDecimals(data []byte) (uint8, error) {
	out, err := wrappedCollateral.abi.Unpack("decimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa0712d68.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mint(uint256 _amount) returns()
func (wrappedCollateral *WrappedCollateral) PackMint(amount *big.Int) []byte {
	enc, err := wrappedCollateral.abi.Pack("mint", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa0712d68.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mint(uint256 _amount) returns()
func (wrappedCollateral *WrappedCollateral) TryPackMint(amount *big.Int) ([]byte, error) {
	return wrappedCollateral.abi.Pack("mint", amount)
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function name() view returns(string)
func (wrappedCollateral *WrappedCollateral) PackName() []byte {
	enc, err := wrappedCollateral.abi.Pack("name")
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
func (wrappedCollateral *WrappedCollateral) TryPackName() ([]byte, error) {
	return wrappedCollateral.abi.Pack("name")
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (wrappedCollateral *WrappedCollateral) UnpackName(data []byte) (string, error) {
	out, err := wrappedCollateral.abi.Unpack("name", data)
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
func (wrappedCollateral *WrappedCollateral) PackNonces(arg0 common.Address) []byte {
	enc, err := wrappedCollateral.abi.Pack("nonces", arg0)
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
func (wrappedCollateral *WrappedCollateral) TryPackNonces(arg0 common.Address) ([]byte, error) {
	return wrappedCollateral.abi.Pack("nonces", arg0)
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (wrappedCollateral *WrappedCollateral) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := wrappedCollateral.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (wrappedCollateral *WrappedCollateral) PackOwner() []byte {
	enc, err := wrappedCollateral.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (wrappedCollateral *WrappedCollateral) TryPackOwner() ([]byte, error) {
	return wrappedCollateral.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (wrappedCollateral *WrappedCollateral) UnpackOwner(data []byte) (common.Address, error) {
	out, err := wrappedCollateral.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd505accf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (wrappedCollateral *WrappedCollateral) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := wrappedCollateral.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
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
func (wrappedCollateral *WrappedCollateral) TryPackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return wrappedCollateral.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
}

// PackRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0357371d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function release(address _to, uint256 _amount) returns()
func (wrappedCollateral *WrappedCollateral) PackRelease(to common.Address, amount *big.Int) []byte {
	enc, err := wrappedCollateral.abi.Pack("release", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0357371d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function release(address _to, uint256 _amount) returns()
func (wrappedCollateral *WrappedCollateral) TryPackRelease(to common.Address, amount *big.Int) ([]byte, error) {
	return wrappedCollateral.abi.Pack("release", to, amount)
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function symbol() view returns(string)
func (wrappedCollateral *WrappedCollateral) PackSymbol() []byte {
	enc, err := wrappedCollateral.abi.Pack("symbol")
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
func (wrappedCollateral *WrappedCollateral) TryPackSymbol() ([]byte, error) {
	return wrappedCollateral.abi.Pack("symbol")
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (wrappedCollateral *WrappedCollateral) UnpackSymbol(data []byte) (string, error) {
	out, err := wrappedCollateral.abi.Unpack("symbol", data)
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
func (wrappedCollateral *WrappedCollateral) PackTotalSupply() []byte {
	enc, err := wrappedCollateral.abi.Pack("totalSupply")
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
func (wrappedCollateral *WrappedCollateral) TryPackTotalSupply() ([]byte, error) {
	return wrappedCollateral.abi.Pack("totalSupply")
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (wrappedCollateral *WrappedCollateral) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := wrappedCollateral.abi.Unpack("totalSupply", data)
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
func (wrappedCollateral *WrappedCollateral) PackTransfer(to common.Address, amount *big.Int) []byte {
	enc, err := wrappedCollateral.abi.Pack("transfer", to, amount)
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
func (wrappedCollateral *WrappedCollateral) TryPackTransfer(to common.Address, amount *big.Int) ([]byte, error) {
	return wrappedCollateral.abi.Pack("transfer", to, amount)
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (wrappedCollateral *WrappedCollateral) UnpackTransfer(data []byte) (bool, error) {
	out, err := wrappedCollateral.abi.Unpack("transfer", data)
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
func (wrappedCollateral *WrappedCollateral) PackTransferFrom(from common.Address, to common.Address, amount *big.Int) []byte {
	enc, err := wrappedCollateral.abi.Pack("transferFrom", from, to, amount)
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
func (wrappedCollateral *WrappedCollateral) TryPackTransferFrom(from common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	return wrappedCollateral.abi.Pack("transferFrom", from, to, amount)
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (wrappedCollateral *WrappedCollateral) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := wrappedCollateral.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackUnderlying is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f307dc3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function underlying() view returns(address)
func (wrappedCollateral *WrappedCollateral) PackUnderlying() []byte {
	enc, err := wrappedCollateral.abi.Pack("underlying")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnderlying is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f307dc3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function underlying() view returns(address)
func (wrappedCollateral *WrappedCollateral) TryPackUnderlying() ([]byte, error) {
	return wrappedCollateral.abi.Pack("underlying")
}

// UnpackUnderlying is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (wrappedCollateral *WrappedCollateral) UnpackUnderlying(data []byte) (common.Address, error) {
	out, err := wrappedCollateral.abi.Unpack("underlying", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackUnwrap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x39f47693.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unwrap(address _to, uint256 _amount) returns()
func (wrappedCollateral *WrappedCollateral) PackUnwrap(to common.Address, amount *big.Int) []byte {
	enc, err := wrappedCollateral.abi.Pack("unwrap", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnwrap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x39f47693.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unwrap(address _to, uint256 _amount) returns()
func (wrappedCollateral *WrappedCollateral) TryPackUnwrap(to common.Address, amount *big.Int) ([]byte, error) {
	return wrappedCollateral.abi.Pack("unwrap", to, amount)
}

// PackWrap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf376c7a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function wrap(address _to, uint256 _amount) returns()
func (wrappedCollateral *WrappedCollateral) PackWrap(to common.Address, amount *big.Int) []byte {
	enc, err := wrappedCollateral.abi.Pack("wrap", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWrap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf376c7a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function wrap(address _to, uint256 _amount) returns()
func (wrappedCollateral *WrappedCollateral) TryPackWrap(to common.Address, amount *big.Int) ([]byte, error) {
	return wrappedCollateral.abi.Pack("wrap", to, amount)
}

// WrappedCollateralApproval represents a Approval event raised by the WrappedCollateral contract.
type WrappedCollateralApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const WrappedCollateralApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (WrappedCollateralApproval) ContractEventName() string {
	return WrappedCollateralApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (wrappedCollateral *WrappedCollateral) UnpackApprovalEvent(log *types.Log) (*WrappedCollateralApproval, error) {
	event := "Approval"
	if len(log.Topics) == 0 || log.Topics[0] != wrappedCollateral.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WrappedCollateralApproval)
	if len(log.Data) > 0 {
		if err := wrappedCollateral.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrappedCollateral.abi.Events[event].Inputs {
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

// WrappedCollateralTransfer represents a Transfer event raised by the WrappedCollateral contract.
type WrappedCollateralTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const WrappedCollateralTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (WrappedCollateralTransfer) ContractEventName() string {
	return WrappedCollateralTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (wrappedCollateral *WrappedCollateral) UnpackTransferEvent(log *types.Log) (*WrappedCollateralTransfer, error) {
	event := "Transfer"
	if len(log.Topics) == 0 || log.Topics[0] != wrappedCollateral.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WrappedCollateralTransfer)
	if len(log.Data) > 0 {
		if err := wrappedCollateral.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range wrappedCollateral.abi.Events[event].Inputs {
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
func (wrappedCollateral *WrappedCollateral) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], wrappedCollateral.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return wrappedCollateral.UnpackOnlyOwnerError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// WrappedCollateralOnlyOwner represents a OnlyOwner error raised by the WrappedCollateral contract.
type WrappedCollateralOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func WrappedCollateralOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (wrappedCollateral *WrappedCollateral) UnpackOnlyOwnerError(raw []byte) (*WrappedCollateralOnlyOwner, error) {
	out := new(WrappedCollateralOnlyOwner)
	if err := wrappedCollateral.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

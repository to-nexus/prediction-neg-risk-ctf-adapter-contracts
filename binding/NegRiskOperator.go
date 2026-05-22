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

// NegRiskOperatorMetaData contains all meta data concerning the NegRiskOperator contract.
var NegRiskOperatorMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nrAdapter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DELAY_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"name\":\"emergencyResolveQuestion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"}],\"name\":\"flagQuestion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"}],\"name\":\"flaggedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nrAdapter\",\"outputs\":[{\"internalType\":\"contractNegRiskAdapter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prepareCondition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeBips\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"prepareMarket\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"prepareQuestion\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"questionIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_payouts\",\"type\":\"uint256[]\"}],\"name\":\"reportPayouts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"}],\"name\":\"reportedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"}],\"name\":\"resolveQuestion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"}],\"name\":\"results\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"}],\"name\":\"unflagQuestion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MarketPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"QuestionEmergencyResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"}],\"name\":\"QuestionFlagged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"questionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"QuestionPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"QuestionReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"QuestionResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"}],\"name\":\"QuestionUnflagged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DelayPeriodNotOver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPayouts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEligibleForEmergencyResolution\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyFlagged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNegRiskAdapter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotFlagged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuestionAlreadyReported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuestionWithRequestIdAlreadyPrepared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResultNotAvailable\",\"type\":\"error\"}]",
	ID:  "NegRiskOperator",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161147e38038061147e83398101604081905261002e91610051565b335f908152602081905260409020600190556001600160a01b031660805261007e565b5f60208284031215610061575f5ffd5b81516001600160a01b0381168114610077575f5ffd5b9392505050565b6080516113cc6100b25f395f81816101ec01528181610638015281816109b001528181610e2d0152610fa601526113cc5ff3fe608060405234801561000f575f5ffd5b506004361061016e575f3560e01c80637b50d3b2116100d2578063c49298ac11610088578063dc89a19811610063578063dc89a19814610380578063e71a02e11461039f578063ead41243146103a6575f5ffd5b8063c49298ac14610347578063d46da3d51461035a578063d96ee7541461036d575f5ffd5b80638a0db615116100b85780638a0db6151461030d5780638bad0c0a14610320578063b8d89e0714610328575f5ffd5b80637b50d3b2146102da5780637dc0d1d0146102ed575f5ffd5b80634c6b25b1116101275780636e88c8fd1161010d5780636e88c8fd1461029557806370480275146102b45780637adbf973146102c7575f5ffd5b80634c6b25b1146102605780636b942f7c14610282575f5ffd5b806324d7806c1161015757806324d7806c1461019a57806325c0520a146101e7578063429b62e514610233575f5ffd5b80630aaf23fa146101725780631785f53c14610187575b5f5ffd5b610185610180366004611078565b6103b9565b005b6101856101953660046110b7565b610481565b6101d26101a83660046110b7565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205460011490565b60405190151581526020015b60405180910390f35b61020e7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101de565b6102526102413660046110b7565b5f6020819052908152604090205481565b6040519081526020016101de565b6101d261026e366004611078565b60036020525f908152604090205460ff1681565b610185610290366004611078565b61051a565b6102526102a3366004611078565b60056020525f908152604090205481565b6101856102c23660046110b7565b6106e6565b6101856102d53660046110b7565b610780565b6101856102e8366004611078565b61085f565b60015461020e9073ffffffffffffffffffffffffffffffffffffffff1681565b61025261031b36600461111c565b610929565b610185610a6f565b610252610336366004611078565b60046020525f908152604090205481565b610185610355366004611164565b610af1565b6101856103683660046111de565b610d23565b61018561037b366004611210565b505050565b61025261038e366004611078565b60026020525f908152604090205481565b6102525f81565b6102526103b4366004611240565b610eda565b335f90815260208190526040902054600114610401576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f818152600460205260408120549003610447576040517f015030c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f818152600460205260408082208290555182917f052435bc04fc49113a7bfd9198a92c0852ca622a621800f6da66d4b29b786c0591a250565b335f908152602081905260409020546001146104c9576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f818152602081905260408082208290555133917f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e91a350565b5f81815260046020526040902054819015610561576040517f18e6a4f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82815260056020526040812054908190036105a9576040517f563c5f3c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105b35f8261128f565b4210156105ec576040517fd0b72b4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260036020526040908190205490517fe200affd0000000000000000000000000000000000000000000000000000000081526004810185905260ff9091168015156024830152907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e200affd906044015f604051808303815f87803b15801561068e575f5ffd5b505af11580156106a0573d5f5f3e3d5ffd5b50505050837f5c3937ed929cd157b73b417381d743daf6e1ef65999e3ccb5dd64bc3247e28d6826040516106d8911515815260200190565b60405180910390a250505050565b335f9081526020819052604090205460011461072e576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f81815260208190526040808220600190555133917ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc91a350565b335f908152602081905260409020546001146107c8576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015473ffffffffffffffffffffffffffffffffffffffff1615610818576040517f8c7ee8d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b335f908152602081905260409020546001146108a7576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f818152600460205260409020548190156108ee576040517f18e6a4f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f828152600460205260408082204290555183917f2435a0347185933b12027c6f394a5fd9c03646dba233e956f50658719dfc0b3591a25050565b335f90815260208190526040812054600114610971576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f8a0db6150000000000000000000000000000000000000000000000000000000081525f9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690638a0db615906109e990889088908890600401611314565b6020604051808303815f875af1158015610a05573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a299190611336565b9050807f8138c0666fe0f752ff38486f542284f127aef02642c9c8db716ee1088839eeb0868686604051610a5f93929190611314565b60405180910390a2949350505050565b335f90815260208190526040902054600114610ab7576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f818152602081905260408082208290555182917f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e91a3565b60015473ffffffffffffffffffffffffffffffffffffffff163314610b42576040517f80fee10500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028114610b7c576040517f663493a000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82825f818110610b8f57610b8f61134d565b9050602002013590505f83836001818110610bac57610bac61134d565b9050602002013590508082610bc1919061128f565b600114610bfa576040517f663493a000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8581526002602052604090205480610c3f576040517fba0514c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8181526005602052604090205415610c84576040517facbb0dde00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83600114610c93575f610c96565b60015b5f83815260036020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016851515908117909155600583529281902042905580518b81529182019290925291925083917f504306b41b2531b3fd2bc5e1b32dc1fc87501906cfc63c1180e3873af20f0eae910160405180910390a250505050505050565b335f90815260208190526040902054600114610d6b576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281526004602052604081205490819003610db3576040517f015030c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610dbd5f8261128f565b421015610df6576040517fd0b72b4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fe200affd0000000000000000000000000000000000000000000000000000000081526004810184905282151560248201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e200affd906044015f604051808303815f87803b158015610e83575f5ffd5b505af1158015610e95573d5f5f3e3d5ffd5b50505050827fd1aea2ca9d3458614d11a93a203dd9fabbd3576aeb841422c46e235637333cb983604051610ecd911515815260200190565b60405180910390a2505050565b335f90815260208190526040812054600114610f22576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281526002602052604090205415610f67576040517f36640f2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f1d69b48d0000000000000000000000000000000000000000000000000000000081525f9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690631d69b48d90610fdf90899089908990600401611314565b6020604051808303815f875af1158015610ffb573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061101f9190611336565b5f84815260026020526040902081905590508281877fcdc45423ec79c60a3fe3de57272e598d71a4ec88822e822ac8e134184a8435aa8289896040516110679392919061137a565b60405180910390a495945050505050565b5f60208284031215611088575f5ffd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146110b2575f5ffd5b919050565b5f602082840312156110c7575f5ffd5b6110d08261108f565b9392505050565b5f5f83601f8401126110e7575f5ffd5b50813567ffffffffffffffff8111156110fe575f5ffd5b602083019150836020828501011115611115575f5ffd5b9250929050565b5f5f5f6040848603121561112e575f5ffd5b83359250602084013567ffffffffffffffff81111561114b575f5ffd5b611157868287016110d7565b9497909650939450505050565b5f5f5f60408486031215611176575f5ffd5b83359250602084013567ffffffffffffffff811115611193575f5ffd5b8401601f810186136111a3575f5ffd5b803567ffffffffffffffff8111156111b9575f5ffd5b8660208260051b84010111156111cd575f5ffd5b939660209190910195509293505050565b5f5f604083850312156111ef575f5ffd5b8235915060208301358015158114611205575f5ffd5b809150509250929050565b5f5f5f60608486031215611222575f5ffd5b61122b8461108f565b95602085013595506040909401359392505050565b5f5f5f5f60608587031215611253575f5ffd5b84359350602085013567ffffffffffffffff811115611270575f5ffd5b61127c878288016110d7565b9598909750949560400135949350505050565b808201808211156112c7577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b92915050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b838152604060208201525f61132d6040830184866112cd565b95945050505050565b5f60208284031215611346575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b60ff84168152604060208201525f61132d6040830184866112cd56fea26469706673582212207a02d2dfc51f82d9924f06d938bc3beaf6ff08ce93c9b2bf05f1e6396d845ecf64736f6c634300081e0033",
}

// NegRiskOperator is an auto generated Go binding around an Ethereum contract.
type NegRiskOperator struct {
	abi abi.ABI
}

// NewNegRiskOperator creates a new instance of NegRiskOperator.
func NewNegRiskOperator() *NegRiskOperator {
	parsed, err := NegRiskOperatorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &NegRiskOperator{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *NegRiskOperator) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _nrAdapter) returns()
func (negRiskOperator *NegRiskOperator) PackConstructor(_nrAdapter common.Address) []byte {
	enc, err := negRiskOperator.abi.Pack("", _nrAdapter)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDELAYPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe71a02e1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DELAY_PERIOD() view returns(uint256)
func (negRiskOperator *NegRiskOperator) PackDELAYPERIOD() []byte {
	enc, err := negRiskOperator.abi.Pack("DELAY_PERIOD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDELAYPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe71a02e1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function DELAY_PERIOD() view returns(uint256)
func (negRiskOperator *NegRiskOperator) TryPackDELAYPERIOD() ([]byte, error) {
	return negRiskOperator.abi.Pack("DELAY_PERIOD")
}

// UnpackDELAYPERIOD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe71a02e1.
//
// Solidity: function DELAY_PERIOD() view returns(uint256)
func (negRiskOperator *NegRiskOperator) UnpackDELAYPERIOD(data []byte) (*big.Int, error) {
	out, err := negRiskOperator.abi.Unpack("DELAY_PERIOD", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackAddAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70480275.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addAdmin(address admin) returns()
func (negRiskOperator *NegRiskOperator) PackAddAdmin(admin common.Address) []byte {
	enc, err := negRiskOperator.abi.Pack("addAdmin", admin)
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
func (negRiskOperator *NegRiskOperator) TryPackAddAdmin(admin common.Address) ([]byte, error) {
	return negRiskOperator.abi.Pack("addAdmin", admin)
}

// PackAdmins is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x429b62e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function admins(address ) view returns(uint256)
func (negRiskOperator *NegRiskOperator) PackAdmins(arg0 common.Address) []byte {
	enc, err := negRiskOperator.abi.Pack("admins", arg0)
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
func (negRiskOperator *NegRiskOperator) TryPackAdmins(arg0 common.Address) ([]byte, error) {
	return negRiskOperator.abi.Pack("admins", arg0)
}

// UnpackAdmins is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (negRiskOperator *NegRiskOperator) UnpackAdmins(data []byte) (*big.Int, error) {
	out, err := negRiskOperator.abi.Unpack("admins", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackEmergencyResolveQuestion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd46da3d5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function emergencyResolveQuestion(bytes32 _questionId, bool _result) returns()
func (negRiskOperator *NegRiskOperator) PackEmergencyResolveQuestion(questionId [32]byte, result bool) []byte {
	enc, err := negRiskOperator.abi.Pack("emergencyResolveQuestion", questionId, result)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEmergencyResolveQuestion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd46da3d5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function emergencyResolveQuestion(bytes32 _questionId, bool _result) returns()
func (negRiskOperator *NegRiskOperator) TryPackEmergencyResolveQuestion(questionId [32]byte, result bool) ([]byte, error) {
	return negRiskOperator.abi.Pack("emergencyResolveQuestion", questionId, result)
}

// PackFlagQuestion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b50d3b2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flagQuestion(bytes32 _questionId) returns()
func (negRiskOperator *NegRiskOperator) PackFlagQuestion(questionId [32]byte) []byte {
	enc, err := negRiskOperator.abi.Pack("flagQuestion", questionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFlagQuestion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b50d3b2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function flagQuestion(bytes32 _questionId) returns()
func (negRiskOperator *NegRiskOperator) TryPackFlagQuestion(questionId [32]byte) ([]byte, error) {
	return negRiskOperator.abi.Pack("flagQuestion", questionId)
}

// PackFlaggedAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb8d89e07.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flaggedAt(bytes32 _questionId) view returns(uint256)
func (negRiskOperator *NegRiskOperator) PackFlaggedAt(questionId [32]byte) []byte {
	enc, err := negRiskOperator.abi.Pack("flaggedAt", questionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFlaggedAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb8d89e07.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function flaggedAt(bytes32 _questionId) view returns(uint256)
func (negRiskOperator *NegRiskOperator) TryPackFlaggedAt(questionId [32]byte) ([]byte, error) {
	return negRiskOperator.abi.Pack("flaggedAt", questionId)
}

// UnpackFlaggedAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb8d89e07.
//
// Solidity: function flaggedAt(bytes32 _questionId) view returns(uint256)
func (negRiskOperator *NegRiskOperator) UnpackFlaggedAt(data []byte) (*big.Int, error) {
	out, err := negRiskOperator.abi.Unpack("flaggedAt", data)
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
func (negRiskOperator *NegRiskOperator) PackIsAdmin(addr common.Address) []byte {
	enc, err := negRiskOperator.abi.Pack("isAdmin", addr)
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
func (negRiskOperator *NegRiskOperator) TryPackIsAdmin(addr common.Address) ([]byte, error) {
	return negRiskOperator.abi.Pack("isAdmin", addr)
}

// UnpackIsAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (negRiskOperator *NegRiskOperator) UnpackIsAdmin(data []byte) (bool, error) {
	out, err := negRiskOperator.abi.Unpack("isAdmin", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackNrAdapter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25c0520a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nrAdapter() view returns(address)
func (negRiskOperator *NegRiskOperator) PackNrAdapter() []byte {
	enc, err := negRiskOperator.abi.Pack("nrAdapter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNrAdapter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25c0520a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function nrAdapter() view returns(address)
func (negRiskOperator *NegRiskOperator) TryPackNrAdapter() ([]byte, error) {
	return negRiskOperator.abi.Pack("nrAdapter")
}

// UnpackNrAdapter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x25c0520a.
//
// Solidity: function nrAdapter() view returns(address)
func (negRiskOperator *NegRiskOperator) UnpackNrAdapter(data []byte) (common.Address, error) {
	out, err := negRiskOperator.abi.Unpack("nrAdapter", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7dc0d1d0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function oracle() view returns(address)
func (negRiskOperator *NegRiskOperator) PackOracle() []byte {
	enc, err := negRiskOperator.abi.Pack("oracle")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7dc0d1d0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function oracle() view returns(address)
func (negRiskOperator *NegRiskOperator) TryPackOracle() ([]byte, error) {
	return negRiskOperator.abi.Pack("oracle")
}

// UnpackOracle is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (negRiskOperator *NegRiskOperator) UnpackOracle(data []byte) (common.Address, error) {
	out, err := negRiskOperator.abi.Unpack("oracle", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPrepareCondition is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd96ee754.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function prepareCondition(address , bytes32 , uint256 ) returns()
func (negRiskOperator *NegRiskOperator) PackPrepareCondition(arg0 common.Address, arg1 [32]byte, arg2 *big.Int) []byte {
	enc, err := negRiskOperator.abi.Pack("prepareCondition", arg0, arg1, arg2)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPrepareCondition is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd96ee754.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function prepareCondition(address , bytes32 , uint256 ) returns()
func (negRiskOperator *NegRiskOperator) TryPackPrepareCondition(arg0 common.Address, arg1 [32]byte, arg2 *big.Int) ([]byte, error) {
	return negRiskOperator.abi.Pack("prepareCondition", arg0, arg1, arg2)
}

// PackPrepareMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a0db615.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function prepareMarket(uint256 _feeBips, bytes _data) returns(bytes32)
func (negRiskOperator *NegRiskOperator) PackPrepareMarket(feeBips *big.Int, data []byte) []byte {
	enc, err := negRiskOperator.abi.Pack("prepareMarket", feeBips, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPrepareMarket is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a0db615.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function prepareMarket(uint256 _feeBips, bytes _data) returns(bytes32)
func (negRiskOperator *NegRiskOperator) TryPackPrepareMarket(feeBips *big.Int, data []byte) ([]byte, error) {
	return negRiskOperator.abi.Pack("prepareMarket", feeBips, data)
}

// UnpackPrepareMarket is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8a0db615.
//
// Solidity: function prepareMarket(uint256 _feeBips, bytes _data) returns(bytes32)
func (negRiskOperator *NegRiskOperator) UnpackPrepareMarket(data []byte) ([32]byte, error) {
	out, err := negRiskOperator.abi.Unpack("prepareMarket", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackPrepareQuestion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xead41243.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function prepareQuestion(bytes32 _marketId, bytes _data, bytes32 _requestId) returns(bytes32)
func (negRiskOperator *NegRiskOperator) PackPrepareQuestion(marketId [32]byte, data []byte, requestId [32]byte) []byte {
	enc, err := negRiskOperator.abi.Pack("prepareQuestion", marketId, data, requestId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPrepareQuestion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xead41243.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function prepareQuestion(bytes32 _marketId, bytes _data, bytes32 _requestId) returns(bytes32)
func (negRiskOperator *NegRiskOperator) TryPackPrepareQuestion(marketId [32]byte, data []byte, requestId [32]byte) ([]byte, error) {
	return negRiskOperator.abi.Pack("prepareQuestion", marketId, data, requestId)
}

// UnpackPrepareQuestion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xead41243.
//
// Solidity: function prepareQuestion(bytes32 _marketId, bytes _data, bytes32 _requestId) returns(bytes32)
func (negRiskOperator *NegRiskOperator) UnpackPrepareQuestion(data []byte) ([32]byte, error) {
	out, err := negRiskOperator.abi.Unpack("prepareQuestion", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackQuestionIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdc89a198.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function questionIds(bytes32 _requestId) view returns(bytes32)
func (negRiskOperator *NegRiskOperator) PackQuestionIds(requestId [32]byte) []byte {
	enc, err := negRiskOperator.abi.Pack("questionIds", requestId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuestionIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdc89a198.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function questionIds(bytes32 _requestId) view returns(bytes32)
func (negRiskOperator *NegRiskOperator) TryPackQuestionIds(requestId [32]byte) ([]byte, error) {
	return negRiskOperator.abi.Pack("questionIds", requestId)
}

// UnpackQuestionIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdc89a198.
//
// Solidity: function questionIds(bytes32 _requestId) view returns(bytes32)
func (negRiskOperator *NegRiskOperator) UnpackQuestionIds(data []byte) ([32]byte, error) {
	out, err := negRiskOperator.abi.Unpack("questionIds", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackRemoveAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1785f53c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeAdmin(address admin) returns()
func (negRiskOperator *NegRiskOperator) PackRemoveAdmin(admin common.Address) []byte {
	enc, err := negRiskOperator.abi.Pack("removeAdmin", admin)
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
func (negRiskOperator *NegRiskOperator) TryPackRemoveAdmin(admin common.Address) ([]byte, error) {
	return negRiskOperator.abi.Pack("removeAdmin", admin)
}

// PackRenounceAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8bad0c0a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceAdmin() returns()
func (negRiskOperator *NegRiskOperator) PackRenounceAdmin() []byte {
	enc, err := negRiskOperator.abi.Pack("renounceAdmin")
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
func (negRiskOperator *NegRiskOperator) TryPackRenounceAdmin() ([]byte, error) {
	return negRiskOperator.abi.Pack("renounceAdmin")
}

// PackReportPayouts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc49298ac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reportPayouts(bytes32 _requestId, uint256[] _payouts) returns()
func (negRiskOperator *NegRiskOperator) PackReportPayouts(requestId [32]byte, payouts []*big.Int) []byte {
	enc, err := negRiskOperator.abi.Pack("reportPayouts", requestId, payouts)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReportPayouts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc49298ac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reportPayouts(bytes32 _requestId, uint256[] _payouts) returns()
func (negRiskOperator *NegRiskOperator) TryPackReportPayouts(requestId [32]byte, payouts []*big.Int) ([]byte, error) {
	return negRiskOperator.abi.Pack("reportPayouts", requestId, payouts)
}

// PackReportedAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e88c8fd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reportedAt(bytes32 _questionId) view returns(uint256)
func (negRiskOperator *NegRiskOperator) PackReportedAt(questionId [32]byte) []byte {
	enc, err := negRiskOperator.abi.Pack("reportedAt", questionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReportedAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e88c8fd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reportedAt(bytes32 _questionId) view returns(uint256)
func (negRiskOperator *NegRiskOperator) TryPackReportedAt(questionId [32]byte) ([]byte, error) {
	return negRiskOperator.abi.Pack("reportedAt", questionId)
}

// UnpackReportedAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6e88c8fd.
//
// Solidity: function reportedAt(bytes32 _questionId) view returns(uint256)
func (negRiskOperator *NegRiskOperator) UnpackReportedAt(data []byte) (*big.Int, error) {
	out, err := negRiskOperator.abi.Unpack("reportedAt", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackResolveQuestion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6b942f7c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolveQuestion(bytes32 _questionId) returns()
func (negRiskOperator *NegRiskOperator) PackResolveQuestion(questionId [32]byte) []byte {
	enc, err := negRiskOperator.abi.Pack("resolveQuestion", questionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolveQuestion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6b942f7c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolveQuestion(bytes32 _questionId) returns()
func (negRiskOperator *NegRiskOperator) TryPackResolveQuestion(questionId [32]byte) ([]byte, error) {
	return negRiskOperator.abi.Pack("resolveQuestion", questionId)
}

// PackResults is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4c6b25b1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function results(bytes32 _questionId) view returns(bool)
func (negRiskOperator *NegRiskOperator) PackResults(questionId [32]byte) []byte {
	enc, err := negRiskOperator.abi.Pack("results", questionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResults is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4c6b25b1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function results(bytes32 _questionId) view returns(bool)
func (negRiskOperator *NegRiskOperator) TryPackResults(questionId [32]byte) ([]byte, error) {
	return negRiskOperator.abi.Pack("results", questionId)
}

// UnpackResults is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4c6b25b1.
//
// Solidity: function results(bytes32 _questionId) view returns(bool)
func (negRiskOperator *NegRiskOperator) UnpackResults(data []byte) (bool, error) {
	out, err := negRiskOperator.abi.Unpack("results", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSetOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7adbf973.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setOracle(address _oracle) returns()
func (negRiskOperator *NegRiskOperator) PackSetOracle(oracle common.Address) []byte {
	enc, err := negRiskOperator.abi.Pack("setOracle", oracle)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7adbf973.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setOracle(address _oracle) returns()
func (negRiskOperator *NegRiskOperator) TryPackSetOracle(oracle common.Address) ([]byte, error) {
	return negRiskOperator.abi.Pack("setOracle", oracle)
}

// PackUnflagQuestion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aaf23fa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unflagQuestion(bytes32 _questionId) returns()
func (negRiskOperator *NegRiskOperator) PackUnflagQuestion(questionId [32]byte) []byte {
	enc, err := negRiskOperator.abi.Pack("unflagQuestion", questionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnflagQuestion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aaf23fa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unflagQuestion(bytes32 _questionId) returns()
func (negRiskOperator *NegRiskOperator) TryPackUnflagQuestion(questionId [32]byte) ([]byte, error) {
	return negRiskOperator.abi.Pack("unflagQuestion", questionId)
}

// NegRiskOperatorMarketPrepared represents a MarketPrepared event raised by the NegRiskOperator contract.
type NegRiskOperatorMarketPrepared struct {
	MarketId [32]byte
	FeeBips  *big.Int
	Data     []byte
	Raw      *types.Log // Blockchain specific contextual infos
}

const NegRiskOperatorMarketPreparedEventName = "MarketPrepared"

// ContractEventName returns the user-defined event name.
func (NegRiskOperatorMarketPrepared) ContractEventName() string {
	return NegRiskOperatorMarketPreparedEventName
}

// UnpackMarketPreparedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MarketPrepared(bytes32 indexed marketId, uint256 feeBips, bytes data)
func (negRiskOperator *NegRiskOperator) UnpackMarketPreparedEvent(log *types.Log) (*NegRiskOperatorMarketPrepared, error) {
	event := "MarketPrepared"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskOperator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskOperatorMarketPrepared)
	if len(log.Data) > 0 {
		if err := negRiskOperator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskOperator.abi.Events[event].Inputs {
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

// NegRiskOperatorNewAdmin represents a NewAdmin event raised by the NegRiskOperator contract.
type NegRiskOperatorNewAdmin struct {
	Admin           common.Address
	NewAdminAddress common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const NegRiskOperatorNewAdminEventName = "NewAdmin"

// ContractEventName returns the user-defined event name.
func (NegRiskOperatorNewAdmin) ContractEventName() string {
	return NegRiskOperatorNewAdminEventName
}

// UnpackNewAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (negRiskOperator *NegRiskOperator) UnpackNewAdminEvent(log *types.Log) (*NegRiskOperatorNewAdmin, error) {
	event := "NewAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskOperator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskOperatorNewAdmin)
	if len(log.Data) > 0 {
		if err := negRiskOperator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskOperator.abi.Events[event].Inputs {
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

// NegRiskOperatorQuestionEmergencyResolved represents a QuestionEmergencyResolved event raised by the NegRiskOperator contract.
type NegRiskOperatorQuestionEmergencyResolved struct {
	QuestionId [32]byte
	Result     bool
	Raw        *types.Log // Blockchain specific contextual infos
}

const NegRiskOperatorQuestionEmergencyResolvedEventName = "QuestionEmergencyResolved"

// ContractEventName returns the user-defined event name.
func (NegRiskOperatorQuestionEmergencyResolved) ContractEventName() string {
	return NegRiskOperatorQuestionEmergencyResolvedEventName
}

// UnpackQuestionEmergencyResolvedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionEmergencyResolved(bytes32 indexed questionId, bool result)
func (negRiskOperator *NegRiskOperator) UnpackQuestionEmergencyResolvedEvent(log *types.Log) (*NegRiskOperatorQuestionEmergencyResolved, error) {
	event := "QuestionEmergencyResolved"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskOperator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskOperatorQuestionEmergencyResolved)
	if len(log.Data) > 0 {
		if err := negRiskOperator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskOperator.abi.Events[event].Inputs {
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

// NegRiskOperatorQuestionFlagged represents a QuestionFlagged event raised by the NegRiskOperator contract.
type NegRiskOperatorQuestionFlagged struct {
	QuestionId [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const NegRiskOperatorQuestionFlaggedEventName = "QuestionFlagged"

// ContractEventName returns the user-defined event name.
func (NegRiskOperatorQuestionFlagged) ContractEventName() string {
	return NegRiskOperatorQuestionFlaggedEventName
}

// UnpackQuestionFlaggedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionFlagged(bytes32 indexed questionId)
func (negRiskOperator *NegRiskOperator) UnpackQuestionFlaggedEvent(log *types.Log) (*NegRiskOperatorQuestionFlagged, error) {
	event := "QuestionFlagged"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskOperator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskOperatorQuestionFlagged)
	if len(log.Data) > 0 {
		if err := negRiskOperator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskOperator.abi.Events[event].Inputs {
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

// NegRiskOperatorQuestionPrepared represents a QuestionPrepared event raised by the NegRiskOperator contract.
type NegRiskOperatorQuestionPrepared struct {
	MarketId      [32]byte
	QuestionId    [32]byte
	RequestId     [32]byte
	QuestionIndex *big.Int
	Data          []byte
	Raw           *types.Log // Blockchain specific contextual infos
}

const NegRiskOperatorQuestionPreparedEventName = "QuestionPrepared"

// ContractEventName returns the user-defined event name.
func (NegRiskOperatorQuestionPrepared) ContractEventName() string {
	return NegRiskOperatorQuestionPreparedEventName
}

// UnpackQuestionPreparedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionPrepared(bytes32 indexed marketId, bytes32 indexed questionId, bytes32 indexed requestId, uint256 questionIndex, bytes data)
func (negRiskOperator *NegRiskOperator) UnpackQuestionPreparedEvent(log *types.Log) (*NegRiskOperatorQuestionPrepared, error) {
	event := "QuestionPrepared"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskOperator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskOperatorQuestionPrepared)
	if len(log.Data) > 0 {
		if err := negRiskOperator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskOperator.abi.Events[event].Inputs {
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

// NegRiskOperatorQuestionReported represents a QuestionReported event raised by the NegRiskOperator contract.
type NegRiskOperatorQuestionReported struct {
	QuestionId [32]byte
	RequestId  [32]byte
	Result     bool
	Raw        *types.Log // Blockchain specific contextual infos
}

const NegRiskOperatorQuestionReportedEventName = "QuestionReported"

// ContractEventName returns the user-defined event name.
func (NegRiskOperatorQuestionReported) ContractEventName() string {
	return NegRiskOperatorQuestionReportedEventName
}

// UnpackQuestionReportedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionReported(bytes32 indexed questionId, bytes32 requestId, bool result)
func (negRiskOperator *NegRiskOperator) UnpackQuestionReportedEvent(log *types.Log) (*NegRiskOperatorQuestionReported, error) {
	event := "QuestionReported"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskOperator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskOperatorQuestionReported)
	if len(log.Data) > 0 {
		if err := negRiskOperator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskOperator.abi.Events[event].Inputs {
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

// NegRiskOperatorQuestionResolved represents a QuestionResolved event raised by the NegRiskOperator contract.
type NegRiskOperatorQuestionResolved struct {
	QuestionId [32]byte
	Result     bool
	Raw        *types.Log // Blockchain specific contextual infos
}

const NegRiskOperatorQuestionResolvedEventName = "QuestionResolved"

// ContractEventName returns the user-defined event name.
func (NegRiskOperatorQuestionResolved) ContractEventName() string {
	return NegRiskOperatorQuestionResolvedEventName
}

// UnpackQuestionResolvedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionResolved(bytes32 indexed questionId, bool result)
func (negRiskOperator *NegRiskOperator) UnpackQuestionResolvedEvent(log *types.Log) (*NegRiskOperatorQuestionResolved, error) {
	event := "QuestionResolved"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskOperator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskOperatorQuestionResolved)
	if len(log.Data) > 0 {
		if err := negRiskOperator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskOperator.abi.Events[event].Inputs {
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

// NegRiskOperatorQuestionUnflagged represents a QuestionUnflagged event raised by the NegRiskOperator contract.
type NegRiskOperatorQuestionUnflagged struct {
	QuestionId [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const NegRiskOperatorQuestionUnflaggedEventName = "QuestionUnflagged"

// ContractEventName returns the user-defined event name.
func (NegRiskOperatorQuestionUnflagged) ContractEventName() string {
	return NegRiskOperatorQuestionUnflaggedEventName
}

// UnpackQuestionUnflaggedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionUnflagged(bytes32 indexed questionId)
func (negRiskOperator *NegRiskOperator) UnpackQuestionUnflaggedEvent(log *types.Log) (*NegRiskOperatorQuestionUnflagged, error) {
	event := "QuestionUnflagged"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskOperator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskOperatorQuestionUnflagged)
	if len(log.Data) > 0 {
		if err := negRiskOperator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskOperator.abi.Events[event].Inputs {
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

// NegRiskOperatorRemovedAdmin represents a RemovedAdmin event raised by the NegRiskOperator contract.
type NegRiskOperatorRemovedAdmin struct {
	Admin        common.Address
	RemovedAdmin common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const NegRiskOperatorRemovedAdminEventName = "RemovedAdmin"

// ContractEventName returns the user-defined event name.
func (NegRiskOperatorRemovedAdmin) ContractEventName() string {
	return NegRiskOperatorRemovedAdminEventName
}

// UnpackRemovedAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (negRiskOperator *NegRiskOperator) UnpackRemovedAdminEvent(log *types.Log) (*NegRiskOperatorRemovedAdmin, error) {
	event := "RemovedAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskOperator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskOperatorRemovedAdmin)
	if len(log.Data) > 0 {
		if err := negRiskOperator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskOperator.abi.Events[event].Inputs {
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
func (negRiskOperator *NegRiskOperator) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["DelayPeriodNotOver"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackDelayPeriodNotOverError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["InvalidPayouts"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackInvalidPayoutsError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["InvalidRequestId"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackInvalidRequestIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["NotAdmin"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackNotAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["NotEligibleForEmergencyResolution"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackNotEligibleForEmergencyResolutionError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["OnlyFlagged"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackOnlyFlaggedError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["OnlyNegRiskAdapter"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackOnlyNegRiskAdapterError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["OnlyNotFlagged"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackOnlyNotFlaggedError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["OnlyOracle"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackOnlyOracleError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["OracleAlreadyInitialized"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackOracleAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["QuestionAlreadyReported"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackQuestionAlreadyReportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["QuestionWithRequestIdAlreadyPrepared"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackQuestionWithRequestIdAlreadyPreparedError(raw[4:])
	}
	if bytes.Equal(raw[:4], negRiskOperator.abi.Errors["ResultNotAvailable"].ID.Bytes()[:4]) {
		return negRiskOperator.UnpackResultNotAvailableError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// NegRiskOperatorDelayPeriodNotOver represents a DelayPeriodNotOver error raised by the NegRiskOperator contract.
type NegRiskOperatorDelayPeriodNotOver struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DelayPeriodNotOver()
func NegRiskOperatorDelayPeriodNotOverErrorID() common.Hash {
	return common.HexToHash("0xd0b72b49fc6c8a3378d7e83cd872945ef37e8cd4c305acbd5c7f2679eca085cf")
}

// UnpackDelayPeriodNotOverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DelayPeriodNotOver()
func (negRiskOperator *NegRiskOperator) UnpackDelayPeriodNotOverError(raw []byte) (*NegRiskOperatorDelayPeriodNotOver, error) {
	out := new(NegRiskOperatorDelayPeriodNotOver)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "DelayPeriodNotOver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskOperatorInvalidPayouts represents a InvalidPayouts error raised by the NegRiskOperator contract.
type NegRiskOperatorInvalidPayouts struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPayouts()
func NegRiskOperatorInvalidPayoutsErrorID() common.Hash {
	return common.HexToHash("0x663493a0e55cd55fffa3aeddad0f26801e0b338aadf7eda63e49476e10b78d64")
}

// UnpackInvalidPayoutsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPayouts()
func (negRiskOperator *NegRiskOperator) UnpackInvalidPayoutsError(raw []byte) (*NegRiskOperatorInvalidPayouts, error) {
	out := new(NegRiskOperatorInvalidPayouts)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "InvalidPayouts", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskOperatorInvalidRequestId represents a InvalidRequestId error raised by the NegRiskOperator contract.
type NegRiskOperatorInvalidRequestId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRequestId()
func NegRiskOperatorInvalidRequestIdErrorID() common.Hash {
	return common.HexToHash("0xba0514c0d5ec80c22a6ebd3f2e6691e2f9ad3d1402978084361a7537d8546138")
}

// UnpackInvalidRequestIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRequestId()
func (negRiskOperator *NegRiskOperator) UnpackInvalidRequestIdError(raw []byte) (*NegRiskOperatorInvalidRequestId, error) {
	out := new(NegRiskOperatorInvalidRequestId)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "InvalidRequestId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskOperatorNotAdmin represents a NotAdmin error raised by the NegRiskOperator contract.
type NegRiskOperatorNotAdmin struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAdmin()
func NegRiskOperatorNotAdminErrorID() common.Hash {
	return common.HexToHash("0x7bfa4b9fb0cd3687c1f539f384b3f3f258f2c9aa9186353d0815413b508ed97d")
}

// UnpackNotAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAdmin()
func (negRiskOperator *NegRiskOperator) UnpackNotAdminError(raw []byte) (*NegRiskOperatorNotAdmin, error) {
	out := new(NegRiskOperatorNotAdmin)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "NotAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskOperatorNotEligibleForEmergencyResolution represents a NotEligibleForEmergencyResolution error raised by the NegRiskOperator contract.
type NegRiskOperatorNotEligibleForEmergencyResolution struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotEligibleForEmergencyResolution()
func NegRiskOperatorNotEligibleForEmergencyResolutionErrorID() common.Hash {
	return common.HexToHash("0x83a0b5b464362976f6274615de702a2bd4ab6e4f41c45f33a0cf9a3a0690ec94")
}

// UnpackNotEligibleForEmergencyResolutionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotEligibleForEmergencyResolution()
func (negRiskOperator *NegRiskOperator) UnpackNotEligibleForEmergencyResolutionError(raw []byte) (*NegRiskOperatorNotEligibleForEmergencyResolution, error) {
	out := new(NegRiskOperatorNotEligibleForEmergencyResolution)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "NotEligibleForEmergencyResolution", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskOperatorOnlyFlagged represents a OnlyFlagged error raised by the NegRiskOperator contract.
type NegRiskOperatorOnlyFlagged struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyFlagged()
func NegRiskOperatorOnlyFlaggedErrorID() common.Hash {
	return common.HexToHash("0x015030c371ffdd076a73ef8697f722d846ca225bed8fff55ed67bc6a51ef970a")
}

// UnpackOnlyFlaggedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyFlagged()
func (negRiskOperator *NegRiskOperator) UnpackOnlyFlaggedError(raw []byte) (*NegRiskOperatorOnlyFlagged, error) {
	out := new(NegRiskOperatorOnlyFlagged)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "OnlyFlagged", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskOperatorOnlyNegRiskAdapter represents a OnlyNegRiskAdapter error raised by the NegRiskOperator contract.
type NegRiskOperatorOnlyNegRiskAdapter struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyNegRiskAdapter()
func NegRiskOperatorOnlyNegRiskAdapterErrorID() common.Hash {
	return common.HexToHash("0x04322eca40e72e2f0396896e75d62b17f87841dedcdea662c214822fb726be61")
}

// UnpackOnlyNegRiskAdapterError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyNegRiskAdapter()
func (negRiskOperator *NegRiskOperator) UnpackOnlyNegRiskAdapterError(raw []byte) (*NegRiskOperatorOnlyNegRiskAdapter, error) {
	out := new(NegRiskOperatorOnlyNegRiskAdapter)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "OnlyNegRiskAdapter", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskOperatorOnlyNotFlagged represents a OnlyNotFlagged error raised by the NegRiskOperator contract.
type NegRiskOperatorOnlyNotFlagged struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyNotFlagged()
func NegRiskOperatorOnlyNotFlaggedErrorID() common.Hash {
	return common.HexToHash("0x18e6a4f7e35a581dfd19dfc24989f05c52119487a906b27998a96e3ec9e4f73a")
}

// UnpackOnlyNotFlaggedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyNotFlagged()
func (negRiskOperator *NegRiskOperator) UnpackOnlyNotFlaggedError(raw []byte) (*NegRiskOperatorOnlyNotFlagged, error) {
	out := new(NegRiskOperatorOnlyNotFlagged)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "OnlyNotFlagged", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskOperatorOnlyOracle represents a OnlyOracle error raised by the NegRiskOperator contract.
type NegRiskOperatorOnlyOracle struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOracle()
func NegRiskOperatorOnlyOracleErrorID() common.Hash {
	return common.HexToHash("0x80fee105c2072ed9e186f4c861f8af91e75494f49bfc0883cbebb2900e492022")
}

// UnpackOnlyOracleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOracle()
func (negRiskOperator *NegRiskOperator) UnpackOnlyOracleError(raw []byte) (*NegRiskOperatorOnlyOracle, error) {
	out := new(NegRiskOperatorOnlyOracle)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "OnlyOracle", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskOperatorOracleAlreadyInitialized represents a OracleAlreadyInitialized error raised by the NegRiskOperator contract.
type NegRiskOperatorOracleAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OracleAlreadyInitialized()
func NegRiskOperatorOracleAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0x8c7ee8d62051ba85340b61a993cd00a36caa1cdc178250423a8e72a3b4d2242f")
}

// UnpackOracleAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OracleAlreadyInitialized()
func (negRiskOperator *NegRiskOperator) UnpackOracleAlreadyInitializedError(raw []byte) (*NegRiskOperatorOracleAlreadyInitialized, error) {
	out := new(NegRiskOperatorOracleAlreadyInitialized)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "OracleAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskOperatorQuestionAlreadyReported represents a QuestionAlreadyReported error raised by the NegRiskOperator contract.
type NegRiskOperatorQuestionAlreadyReported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error QuestionAlreadyReported()
func NegRiskOperatorQuestionAlreadyReportedErrorID() common.Hash {
	return common.HexToHash("0xacbb0dde2c4c71afb28ec0bfff384432224a86ac65a44e6f1e628a0c7fea97d5")
}

// UnpackQuestionAlreadyReportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error QuestionAlreadyReported()
func (negRiskOperator *NegRiskOperator) UnpackQuestionAlreadyReportedError(raw []byte) (*NegRiskOperatorQuestionAlreadyReported, error) {
	out := new(NegRiskOperatorQuestionAlreadyReported)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "QuestionAlreadyReported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskOperatorQuestionWithRequestIdAlreadyPrepared represents a QuestionWithRequestIdAlreadyPrepared error raised by the NegRiskOperator contract.
type NegRiskOperatorQuestionWithRequestIdAlreadyPrepared struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error QuestionWithRequestIdAlreadyPrepared()
func NegRiskOperatorQuestionWithRequestIdAlreadyPreparedErrorID() common.Hash {
	return common.HexToHash("0x36640f2e126a42a104d1d5050329728b3a93e422515f2e82d8849eaad977bcec")
}

// UnpackQuestionWithRequestIdAlreadyPreparedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error QuestionWithRequestIdAlreadyPrepared()
func (negRiskOperator *NegRiskOperator) UnpackQuestionWithRequestIdAlreadyPreparedError(raw []byte) (*NegRiskOperatorQuestionWithRequestIdAlreadyPrepared, error) {
	out := new(NegRiskOperatorQuestionWithRequestIdAlreadyPrepared)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "QuestionWithRequestIdAlreadyPrepared", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NegRiskOperatorResultNotAvailable represents a ResultNotAvailable error raised by the NegRiskOperator contract.
type NegRiskOperatorResultNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ResultNotAvailable()
func NegRiskOperatorResultNotAvailableErrorID() common.Hash {
	return common.HexToHash("0x563c5f3c6fd5b7034e9c8fdbc9a4a6dbdb9318796feafbf4a3ddcb8535018881")
}

// UnpackResultNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ResultNotAvailable()
func (negRiskOperator *NegRiskOperator) UnpackResultNotAvailableError(raw []byte) (*NegRiskOperatorResultNotAvailable, error) {
	out := new(NegRiskOperatorResultNotAvailable)
	if err := negRiskOperator.abi.UnpackIntoInterface(out, "ResultNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

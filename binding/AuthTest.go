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

// StdInvariantFuzzArtifactSelector is an auto generated low-level Go binding around an user-defined struct.

// StdInvariantFuzzInterface is an auto generated low-level Go binding around an user-defined struct.

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.

// AuthTestMetaData contains all meta data concerning the AuthTest contract.
var AuthTestMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IS_TEST\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alice\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"brian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"carly\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeArtifacts\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"excludedArtifacts_\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"excludedContracts_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"excludedSenders_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetArtifactSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"artifact\",\"type\":\"string\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetArtifacts\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"targetedArtifacts_\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targetedContracts_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetInterfaces\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"artifacts\",\"type\":\"string[]\"}],\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targetedSenders_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_AddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_InitialAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_NotAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_OnlyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_RemoveAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_RenounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_revert_AddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_revert_OnlyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"log_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"log_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"log_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"log_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"log_named_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"log_named_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"log_named_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"}],\"name\":\"log_named_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"log_named_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"log_named_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"log_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"logs\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"}]",
	ID:  "AuthTest",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f8054909116909117905534801561002c575f5ffd5b50604051633a0278e960e11b8152602060048201526005602482015264616c69636560d81b6044820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90637404f1d2906064016080604051808303815f875af1158015610091573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100b591906102da565b51601f80546001600160a01b0390921661010002610100600160a81b0319909216919091179055604051633a0278e960e11b8152602060048201526005602482015264313934b0b760d91b6044820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90637404f1d2906064016080604051808303815f875af1158015610140573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061016491906102da565b51602080546001600160a01b0319166001600160a01b03909216919091178155604051633a0278e960e11b8152600481019190915260056024820152646361726c7960d81b6044820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90637404f1d2906064016080604051808303815f875af11580156101e9573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061020d91906102da565b51602180546001600160a01b0319166001600160a01b03909216919091179055604051633a0278e960e11b81526020600482015260056024820152643232bb34b760d91b6044820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90637404f1d2906064016080604051808303815f875af1158015610291573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102b591906102da565b51602280546001600160a01b0319166001600160a01b0390921691909117905561035c565b5f60808284031280156102eb575f5ffd5b50604051608081016001600160401b038111828210171561031a57634e487b7160e01b5f52604160045260245ffd5b60405282516001600160a01b0381168114610333575f5ffd5b815260208381015190820152604080840151908201526060928301519281019290925250919050565b612538806103695f395ff3fe608060405234801561000f575f5ffd5b506004361061019a575f3560e01c806385226c81116100e8578063da4a547311610093578063f5ac20011161006e578063f5ac2001146102f7578063f6e1ae5314610317578063fa7626d41461031f578063fb47e3a21461032c575f5ffd5b8063da4a5473146102df578063df9bbf7e146102e7578063e20c9f71146102ef575f5ffd5b8063b0464fdc116100c3578063b0464fdc146102b7578063b5508aa9146102bf578063ba414fa6146102c7575f5ffd5b806385226c8114610285578063916a17c61461029a578063960ead4b146102af575f5ffd5b80633e0a028a1161014857806366d9a9a01161012357806366d9a9a01461020b5780637168a7d514610220578063719b359d14610265575f5ffd5b80633e0a028a146101f35780633e5e3c23146101fb5780633f7286f414610203575f5ffd5b80632941d785116101785780632941d785146101ce5780632ade3880146101d657806333abd4be146101eb575f5ffd5b806304a536e51461019e5780630a9254e4146101a85780631ed7831c146101b0575b5f5ffd5b6101a6610351565b005b6101a66104ff565b6101b861065a565b6040516101c59190611ced565b60405180910390f35b6101a66106c7565b6101de610a41565b6040516101c59190611d91565b6101a6610b8a565b6101a6610fb0565b6101b8611070565b6101b86110db565b610213611146565b6040516101c59190611eff565b6022546102409073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c5565b6020546102409073ffffffffffffffffffffffffffffffffffffffff1681565b61028d6112bf565b6040516101c59190611f9b565b6102a261138a565b6040516101c59190612010565b6101a661148d565b6102a2611642565b61028d611745565b6102cf611810565b60405190151581526020016101c5565b6101a66118dc565b6101a66119d9565b6101b8611a9f565b6021546102409073ffffffffffffffffffffffffffffffffffffffff1681565b6101a6611b0a565b601f546102cf9060ff1681565b601f5461024090610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6020546040517fca669fa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156103cd575f5ffd5b505af11580156103df573d5f5f3e3d5ffd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527f7bfa4b9f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015610465575f5ffd5b505af1158015610477573d5f5f3e3d5ffd5b50506023546020546040517f7048027500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff918216600482015291169250637048027591506024015f604051808303815f87803b1580156104e7575f5ffd5b505af11580156104f9573d5f5f3e3d5ffd5b50505050565b601f546040517fca669fa700000000000000000000000000000000000000000000000000000000815261010090910473ffffffffffffffffffffffffffffffffffffffff166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801561057f575f5ffd5b505af1158015610591573d5f5f3e3d5ffd5b5050604080518082018252600481527f6175746800000000000000000000000000000000000000000000000000000000602090910152517fb9208574d39bc6b85a528191d39d983f3a0bc58ef7129343fde819f64f7268cd92506105f59150611ce0565b8190604051809103905ff5905080158015610612573d5f5f3e3d5ffd5b50602380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b606060168054806020026020016040519081016040528092919081815260200182805480156106bd57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610692575b5050505050905090565b601f546040517fca669fa700000000000000000000000000000000000000000000000000000000815261010090910473ffffffffffffffffffffffffffffffffffffffff166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015610747575f5ffd5b505af1158015610759573d5f5f3e3d5ffd5b50506023546020546040517f7048027500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff918216600482015291169250637048027591506024015f604051808303815f87803b1580156107c9575f5ffd5b505af11580156107db573d5f5f3e3d5ffd5b50506020546040517fca669fa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801561085b575f5ffd5b505af115801561086d573d5f5f3e3d5ffd5b5050602354601f546040517f1785f53c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff6101009092048216600482015291169250631785f53c91506024015f604051808303815f87803b1580156108e2575f5ffd5b505af11580156108f4573d5f5f3e3d5ffd5b5050602354601f546040517f24d7806c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff6101009092048216600482015261099b9450911691506324d7806c906024015b602060405180830381865afa158015610972573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061099691906120b2565b611b7a565b602354601f546040517f429b62e500000000000000000000000000000000000000000000000000000000815261010090910473ffffffffffffffffffffffffffffffffffffffff9081166004830152610a3f92169063429b62e5906024015b602060405180830381865afa158015610a15573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a3991906120d8565b5f611bfe565b565b6060601e805480602002602001604051908101604052809291908181526020015f905b82821015610b81575f848152602080822060408051808201825260028702909201805473ffffffffffffffffffffffffffffffffffffffff168352600181018054835181870281018701909452808452939591948681019491929084015b82821015610b6a578382905f5260205f20018054610adf906120ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610b0b906120ef565b8015610b565780601f10610b2d57610100808354040283529160200191610b56565b820191905f5260205f20905b815481529060010190602001808311610b3957829003601f168201915b505050505081526020019060010190610ac2565b505050508152505081526020019060010190610a64565b50505050905090565b737109709ecfa91a80626ff3989d68f67f5b1dd12d73ffffffffffffffffffffffffffffffffffffffff1663440ed10d6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610be3575f5ffd5b505af1158015610bf5573d5f5f3e3d5ffd5b5050602054601f5460405173ffffffffffffffffffffffffffffffffffffffff928316945061010090910490911691507ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc905f90a3601f546040517fca669fa700000000000000000000000000000000000000000000000000000000815261010090910473ffffffffffffffffffffffffffffffffffffffff166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015610cca575f5ffd5b505af1158015610cdc573d5f5f3e3d5ffd5b50506023546020546040517f7048027500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff918216600482015291169250637048027591506024015f604051808303815f87803b158015610d4c575f5ffd5b505af1158015610d5e573d5f5f3e3d5ffd5b50506020546040517fca669fa700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015610dde575f5ffd5b505af1158015610df0573d5f5f3e3d5ffd5b5050505060235f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637ff56f226040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610e5a575f5ffd5b505af1158015610e6c573d5f5f3e3d5ffd5b50506023546020546040517f24d7806c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152610f0e9450911691506324d7806c906024015b602060405180830381865afa158015610ee5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f0991906120b2565b611c89565b6023546020546040517f429b62e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152610a3f92919091169063429b62e5906024015b602060405180830381865afa158015610f85573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fa991906120d8565b6001611bfe565b6023546020546040517f24d7806c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526110109291909116906324d7806c90602401610957565b6023546020546040517f429b62e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152610a3f92919091169063429b62e5906024016109fa565b606060188054806020026020016040519081016040528092919081815260200182805480156106bd57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610692575050505050905090565b606060178054806020026020016040519081016040528092919081815260200182805480156106bd57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610692575050505050905090565b6060601b805480602002602001604051908101604052809291908181526020015f905b82821015610b81578382905f5260205f2090600202016040518060400160405290815f82018054611199906120ef565b80601f01602080910402602001604051908101604052809291908181526020018280546111c5906120ef565b80156112105780601f106111e757610100808354040283529160200191611210565b820191905f5260205f20905b8154815290600101906020018083116111f357829003601f168201915b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156112a757602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116112545790505b50505050508152505081526020019060010190611169565b6060601a805480602002602001604051908101604052809291908181526020015f905b82821015610b81578382905f5260205f200180546112ff906120ef565b80601f016020809104026020016040519081016040528092919081815260200182805461132b906120ef565b80156113765780601f1061134d57610100808354040283529160200191611376565b820191905f5260205f20905b81548152906001019060200180831161135957829003601f168201915b5050505050815260200190600101906112e2565b6060601d805480602002602001604051908101604052809291908181526020015f905b82821015610b81575f84815260209081902060408051808201825260028602909201805473ffffffffffffffffffffffffffffffffffffffff16835260018101805483518187028101870190945280845293949193858301939283018282801561147557602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116114225790505b505050505081525050815260200190600101906113ad565b737109709ecfa91a80626ff3989d68f67f5b1dd12d73ffffffffffffffffffffffffffffffffffffffff1663440ed10d6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156114e6575f5ffd5b505af11580156114f8573d5f5f3e3d5ffd5b5050601f5460405161010090910473ffffffffffffffffffffffffffffffffffffffff1692508291507f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e905f90a3601f546040517fca669fa700000000000000000000000000000000000000000000000000000000815261010090910473ffffffffffffffffffffffffffffffffffffffff166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156115c6575f5ffd5b505af11580156115d8573d5f5f3e3d5ffd5b5050505060235f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638bad0c0a6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156108e2575f5ffd5b6060601c805480602002602001604051908101604052809291908181526020015f905b82821015610b81575f84815260209081902060408051808201825260028602909201805473ffffffffffffffffffffffffffffffffffffffff16835260018101805483518187028101870190945280845293949193858301939283018282801561172d57602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116116da5790505b50505050508152505081526020019060010190611665565b60606019805480602002602001604051908101604052809291908181526020015f905b82821015610b81578382905f5260205f20018054611785906120ef565b80601f01602080910402602001604051908101604052809291908181526020018280546117b1906120ef565b80156117fc5780601f106117d3576101008083540402835291602001916117fc565b820191905f5260205f20905b8154815290600101906020018083116117df57829003601f168201915b505050505081526020019060010190611768565b6008545f9060ff16156118235750600190565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c6564000000000000000000000000000000000000000000000000000060248301525f9163667f9d7090604401602060405180830381865afa1580156118b1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118d591906120d8565b1415905090565b601f546040517fca669fa700000000000000000000000000000000000000000000000000000000815261010090910473ffffffffffffffffffffffffffffffffffffffff166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b5f604051808303815f87803b15801561195d575f5ffd5b505af115801561196f573d5f5f3e3d5ffd5b5050505060235f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637ff56f226040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156104e7575f5ffd5b602354601f546040517f24d7806c00000000000000000000000000000000000000000000000000000000815261010090910473ffffffffffffffffffffffffffffffffffffffff9081166004830152611a3c9216906324d7806c90602401610eca565b602354601f546040517f429b62e500000000000000000000000000000000000000000000000000000000815261010090910473ffffffffffffffffffffffffffffffffffffffff9081166004830152610a3f92169063429b62e590602401610f6a565b606060158054806020026020016040519081016040528092919081815260200182805480156106bd57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610692575050505050905090565b6040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527f7bfa4b9f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401611946565b8015611bfb576040517fa59828850000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a5982885906024015b5f6040518083038186803b158015611be4575f5ffd5b505afa158015611bf6573d5f5f3e3d5ffd5b505050505b50565b808214611c85576040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015f6040518083038186803b158015611c6e575f5ffd5b505afa158015611c80573d5f5f3e3d5ffd5b505050505b5050565b80611bfb576040517f0c9fd5810000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90630c9fd58190602401611bce565b6103c28061214183390190565b602080825282518282018190525f918401906040840190835b81811015611d3a57835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101611d06565b509095945050505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015611e97577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805173ffffffffffffffffffffffffffffffffffffffff168652602090810151604082880181905281519088018190529101906060600582901b8801810191908801905f5b81811015611e7d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8503018352611e67848651611d45565b6020958601959094509290920191600101611e2d565b509197505050602094850194929092019150600101611db7565b50929695505050505050565b5f8151808452602084019350602083015f5b82811015611ef55781517fffffffff0000000000000000000000000000000000000000000000000000000016865260209586019590910190600101611eb5565b5093949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015611e97577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160408752611f696040880182611d45565b9050602082015191508681036020880152611f848183611ea3565b965050506020938401939190910190600101611f25565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015611e97577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452611ffb858351611d45565b94506020938401939190910190600101611fc1565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015611e97577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815173ffffffffffffffffffffffffffffffffffffffff8151168652602081015190506040602087015261209c6040870182611ea3565b9550506020938401939190910190600101612036565b5f602082840312156120c2575f5ffd5b815180151581146120d1575f5ffd5b9392505050565b5f602082840312156120e8575f5ffd5b5051919050565b600181811c9082168061210357607f821691505b60208210810361213a577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5091905056fe6080604052348015600e575f5ffd5b50335f908152602081905260409020600190556103948061002e5f395ff3fe608060405234801561000f575f5ffd5b506004361061006f575f3560e01c8063704802751161004d57806370480275146101025780637ff56f22146101155780638bad0c0a1461011d575f5ffd5b80631785f53c1461007357806324d7806c14610088578063429b62e5146100d5575b5f5ffd5b610086610081366004610324565b610125565b005b6100c0610096366004610324565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205460011490565b60405190151581526020015b60405180910390f35b6100f46100e3366004610324565b5f6020819052908152604090205481565b6040519081526020016100cc565b610086610110366004610324565b6101be565b610086610258565b6100866102a2565b335f9081526020819052604090205460011461016d576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f818152602081905260408082208290555133917f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e91a350565b335f90815260208190526040902054600114610206576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81165f81815260208190526040808220600190555133917ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc91a350565b335f908152602081905260409020546001146102a0576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b335f908152602081905260409020546001146102ea576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f818152602081905260408082208290555182917f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e91a3565b5f60208284031215610334575f5ffd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610357575f5ffd5b939250505056fea2646970667358221220ec0bca1549a2127e06eb2d7978132d5a027bb3f87df474f759379389646dec3464736f6c634300081e0033a2646970667358221220ce36a11f9294c3dbe01fa34299fb4e3dac2226f2820bbd44a7eae32e712591b364736f6c634300081e0033",
}

// AuthTest is an auto generated Go binding around an Ethereum contract.
type AuthTest struct {
	abi abi.ABI
}

// NewAuthTest creates a new instance of AuthTest.
func NewAuthTest() *AuthTest {
	parsed, err := AuthTestMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AuthTest{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AuthTest) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackISTEST is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa7626d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function IS_TEST() view returns(bool)
func (authTest *AuthTest) PackISTEST() []byte {
	enc, err := authTest.abi.Pack("IS_TEST")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackISTEST is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa7626d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function IS_TEST() view returns(bool)
func (authTest *AuthTest) TryPackISTEST() ([]byte, error) {
	return authTest.abi.Pack("IS_TEST")
}

// UnpackISTEST is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (authTest *AuthTest) UnpackISTEST(data []byte) (bool, error) {
	out, err := authTest.abi.Unpack("IS_TEST", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackAlice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfb47e3a2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function alice() view returns(address)
func (authTest *AuthTest) PackAlice() []byte {
	enc, err := authTest.abi.Pack("alice")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAlice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfb47e3a2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function alice() view returns(address)
func (authTest *AuthTest) TryPackAlice() ([]byte, error) {
	return authTest.abi.Pack("alice")
}

// UnpackAlice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfb47e3a2.
//
// Solidity: function alice() view returns(address)
func (authTest *AuthTest) UnpackAlice(data []byte) (common.Address, error) {
	out, err := authTest.abi.Unpack("alice", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackBrian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x719b359d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function brian() view returns(address)
func (authTest *AuthTest) PackBrian() []byte {
	enc, err := authTest.abi.Pack("brian")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBrian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x719b359d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function brian() view returns(address)
func (authTest *AuthTest) TryPackBrian() ([]byte, error) {
	return authTest.abi.Pack("brian")
}

// UnpackBrian is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x719b359d.
//
// Solidity: function brian() view returns(address)
func (authTest *AuthTest) UnpackBrian(data []byte) (common.Address, error) {
	out, err := authTest.abi.Unpack("brian", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCarly is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5ac2001.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function carly() view returns(address)
func (authTest *AuthTest) PackCarly() []byte {
	enc, err := authTest.abi.Pack("carly")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCarly is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5ac2001.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function carly() view returns(address)
func (authTest *AuthTest) TryPackCarly() ([]byte, error) {
	return authTest.abi.Pack("carly")
}

// UnpackCarly is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf5ac2001.
//
// Solidity: function carly() view returns(address)
func (authTest *AuthTest) UnpackCarly(data []byte) (common.Address, error) {
	out, err := authTest.abi.Unpack("carly", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDevin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7168a7d5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function devin() view returns(address)
func (authTest *AuthTest) PackDevin() []byte {
	enc, err := authTest.abi.Pack("devin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDevin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7168a7d5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function devin() view returns(address)
func (authTest *AuthTest) TryPackDevin() ([]byte, error) {
	return authTest.abi.Pack("devin")
}

// UnpackDevin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7168a7d5.
//
// Solidity: function devin() view returns(address)
func (authTest *AuthTest) UnpackDevin(data []byte) (common.Address, error) {
	out, err := authTest.abi.Unpack("devin", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackExcludeArtifacts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb5508aa9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (authTest *AuthTest) PackExcludeArtifacts() []byte {
	enc, err := authTest.abi.Pack("excludeArtifacts")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExcludeArtifacts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb5508aa9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (authTest *AuthTest) TryPackExcludeArtifacts() ([]byte, error) {
	return authTest.abi.Pack("excludeArtifacts")
}

// UnpackExcludeArtifacts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (authTest *AuthTest) UnpackExcludeArtifacts(data []byte) ([]string, error) {
	out, err := authTest.abi.Unpack("excludeArtifacts", data)
	if err != nil {
		return *new([]string), err
	}
	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)
	return out0, nil
}

// PackExcludeContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe20c9f71.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (authTest *AuthTest) PackExcludeContracts() []byte {
	enc, err := authTest.abi.Pack("excludeContracts")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExcludeContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe20c9f71.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (authTest *AuthTest) TryPackExcludeContracts() ([]byte, error) {
	return authTest.abi.Pack("excludeContracts")
}

// UnpackExcludeContracts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (authTest *AuthTest) UnpackExcludeContracts(data []byte) ([]common.Address, error) {
	out, err := authTest.abi.Unpack("excludeContracts", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackExcludeSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb0464fdc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (authTest *AuthTest) PackExcludeSelectors() []byte {
	enc, err := authTest.abi.Pack("excludeSelectors")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExcludeSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb0464fdc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (authTest *AuthTest) TryPackExcludeSelectors() ([]byte, error) {
	return authTest.abi.Pack("excludeSelectors")
}

// UnpackExcludeSelectors is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (authTest *AuthTest) UnpackExcludeSelectors(data []byte) ([]StdInvariantFuzzSelector, error) {
	out, err := authTest.abi.Unpack("excludeSelectors", data)
	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}
	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)
	return out0, nil
}

// PackExcludeSenders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1ed7831c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (authTest *AuthTest) PackExcludeSenders() []byte {
	enc, err := authTest.abi.Pack("excludeSenders")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExcludeSenders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1ed7831c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (authTest *AuthTest) TryPackExcludeSenders() ([]byte, error) {
	return authTest.abi.Pack("excludeSenders")
}

// UnpackExcludeSenders is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (authTest *AuthTest) UnpackExcludeSenders(data []byte) ([]common.Address, error) {
	out, err := authTest.abi.Unpack("excludeSenders", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackFailed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba414fa6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function failed() view returns(bool)
func (authTest *AuthTest) PackFailed() []byte {
	enc, err := authTest.abi.Pack("failed")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFailed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba414fa6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function failed() view returns(bool)
func (authTest *AuthTest) TryPackFailed() ([]byte, error) {
	return authTest.abi.Pack("failed")
}

// UnpackFailed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (authTest *AuthTest) UnpackFailed(data []byte) (bool, error) {
	out, err := authTest.abi.Unpack("failed", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSetUp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a9254e4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setUp() returns()
func (authTest *AuthTest) PackSetUp() []byte {
	enc, err := authTest.abi.Pack("setUp")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetUp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a9254e4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setUp() returns()
func (authTest *AuthTest) TryPackSetUp() ([]byte, error) {
	return authTest.abi.Pack("setUp")
}

// PackTargetArtifactSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66d9a9a0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (authTest *AuthTest) PackTargetArtifactSelectors() []byte {
	enc, err := authTest.abi.Pack("targetArtifactSelectors")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTargetArtifactSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66d9a9a0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (authTest *AuthTest) TryPackTargetArtifactSelectors() ([]byte, error) {
	return authTest.abi.Pack("targetArtifactSelectors")
}

// UnpackTargetArtifactSelectors is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (authTest *AuthTest) UnpackTargetArtifactSelectors(data []byte) ([]StdInvariantFuzzArtifactSelector, error) {
	out, err := authTest.abi.Unpack("targetArtifactSelectors", data)
	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}
	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)
	return out0, nil
}

// PackTargetArtifacts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85226c81.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (authTest *AuthTest) PackTargetArtifacts() []byte {
	enc, err := authTest.abi.Pack("targetArtifacts")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTargetArtifacts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85226c81.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (authTest *AuthTest) TryPackTargetArtifacts() ([]byte, error) {
	return authTest.abi.Pack("targetArtifacts")
}

// UnpackTargetArtifacts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (authTest *AuthTest) UnpackTargetArtifacts(data []byte) ([]string, error) {
	out, err := authTest.abi.Unpack("targetArtifacts", data)
	if err != nil {
		return *new([]string), err
	}
	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)
	return out0, nil
}

// PackTargetContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f7286f4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (authTest *AuthTest) PackTargetContracts() []byte {
	enc, err := authTest.abi.Pack("targetContracts")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTargetContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f7286f4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (authTest *AuthTest) TryPackTargetContracts() ([]byte, error) {
	return authTest.abi.Pack("targetContracts")
}

// UnpackTargetContracts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (authTest *AuthTest) UnpackTargetContracts(data []byte) ([]common.Address, error) {
	out, err := authTest.abi.Unpack("targetContracts", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackTargetInterfaces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ade3880.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (authTest *AuthTest) PackTargetInterfaces() []byte {
	enc, err := authTest.abi.Pack("targetInterfaces")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTargetInterfaces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ade3880.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (authTest *AuthTest) TryPackTargetInterfaces() ([]byte, error) {
	return authTest.abi.Pack("targetInterfaces")
}

// UnpackTargetInterfaces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (authTest *AuthTest) UnpackTargetInterfaces(data []byte) ([]StdInvariantFuzzInterface, error) {
	out, err := authTest.abi.Unpack("targetInterfaces", data)
	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}
	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)
	return out0, nil
}

// PackTargetSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x916a17c6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (authTest *AuthTest) PackTargetSelectors() []byte {
	enc, err := authTest.abi.Pack("targetSelectors")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTargetSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x916a17c6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (authTest *AuthTest) TryPackTargetSelectors() ([]byte, error) {
	return authTest.abi.Pack("targetSelectors")
}

// UnpackTargetSelectors is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (authTest *AuthTest) UnpackTargetSelectors(data []byte) ([]StdInvariantFuzzSelector, error) {
	out, err := authTest.abi.Unpack("targetSelectors", data)
	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}
	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)
	return out0, nil
}

// PackTargetSenders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e5e3c23.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (authTest *AuthTest) PackTargetSenders() []byte {
	enc, err := authTest.abi.Pack("targetSenders")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTargetSenders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e5e3c23.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (authTest *AuthTest) TryPackTargetSenders() ([]byte, error) {
	return authTest.abi.Pack("targetSenders")
}

// UnpackTargetSenders is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (authTest *AuthTest) UnpackTargetSenders(data []byte) ([]common.Address, error) {
	out, err := authTest.abi.Unpack("targetSenders", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackTestAddAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x33abd4be.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_AddAdmin() returns()
func (authTest *AuthTest) PackTestAddAdmin() []byte {
	enc, err := authTest.abi.Pack("test_AddAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestAddAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x33abd4be.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_AddAdmin() returns()
func (authTest *AuthTest) TryPackTestAddAdmin() ([]byte, error) {
	return authTest.abi.Pack("test_AddAdmin")
}

// PackTestInitialAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdf9bbf7e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_InitialAdmin() returns()
func (authTest *AuthTest) PackTestInitialAdmin() []byte {
	enc, err := authTest.abi.Pack("test_InitialAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestInitialAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdf9bbf7e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_InitialAdmin() returns()
func (authTest *AuthTest) TryPackTestInitialAdmin() ([]byte, error) {
	return authTest.abi.Pack("test_InitialAdmin")
}

// PackTestNotAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e0a028a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_NotAdmin() returns()
func (authTest *AuthTest) PackTestNotAdmin() []byte {
	enc, err := authTest.abi.Pack("test_NotAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestNotAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e0a028a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_NotAdmin() returns()
func (authTest *AuthTest) TryPackTestNotAdmin() ([]byte, error) {
	return authTest.abi.Pack("test_NotAdmin")
}

// PackTestOnlyAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda4a5473.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_OnlyAdmin() returns()
func (authTest *AuthTest) PackTestOnlyAdmin() []byte {
	enc, err := authTest.abi.Pack("test_OnlyAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestOnlyAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda4a5473.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_OnlyAdmin() returns()
func (authTest *AuthTest) TryPackTestOnlyAdmin() ([]byte, error) {
	return authTest.abi.Pack("test_OnlyAdmin")
}

// PackTestRemoveAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2941d785.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_RemoveAdmin() returns()
func (authTest *AuthTest) PackTestRemoveAdmin() []byte {
	enc, err := authTest.abi.Pack("test_RemoveAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestRemoveAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2941d785.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_RemoveAdmin() returns()
func (authTest *AuthTest) TryPackTestRemoveAdmin() ([]byte, error) {
	return authTest.abi.Pack("test_RemoveAdmin")
}

// PackTestRenounceAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x960ead4b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_RenounceAdmin() returns()
func (authTest *AuthTest) PackTestRenounceAdmin() []byte {
	enc, err := authTest.abi.Pack("test_RenounceAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestRenounceAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x960ead4b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_RenounceAdmin() returns()
func (authTest *AuthTest) TryPackTestRenounceAdmin() ([]byte, error) {
	return authTest.abi.Pack("test_RenounceAdmin")
}

// PackTestRevertAddAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x04a536e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_revert_AddAdmin() returns()
func (authTest *AuthTest) PackTestRevertAddAdmin() []byte {
	enc, err := authTest.abi.Pack("test_revert_AddAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestRevertAddAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x04a536e5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_revert_AddAdmin() returns()
func (authTest *AuthTest) TryPackTestRevertAddAdmin() ([]byte, error) {
	return authTest.abi.Pack("test_revert_AddAdmin")
}

// PackTestRevertOnlyAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf6e1ae53.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_revert_OnlyAdmin() returns()
func (authTest *AuthTest) PackTestRevertOnlyAdmin() []byte {
	enc, err := authTest.abi.Pack("test_revert_OnlyAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestRevertOnlyAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf6e1ae53.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_revert_OnlyAdmin() returns()
func (authTest *AuthTest) TryPackTestRevertOnlyAdmin() ([]byte, error) {
	return authTest.abi.Pack("test_revert_OnlyAdmin")
}

// AuthTestNewAdmin represents a NewAdmin event raised by the AuthTest contract.
type AuthTestNewAdmin struct {
	Admin           common.Address
	NewAdminAddress common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const AuthTestNewAdminEventName = "NewAdmin"

// ContractEventName returns the user-defined event name.
func (AuthTestNewAdmin) ContractEventName() string {
	return AuthTestNewAdminEventName
}

// UnpackNewAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (authTest *AuthTest) UnpackNewAdminEvent(log *types.Log) (*AuthTestNewAdmin, error) {
	event := "NewAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestNewAdmin)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestRemovedAdmin represents a RemovedAdmin event raised by the AuthTest contract.
type AuthTestRemovedAdmin struct {
	Admin        common.Address
	RemovedAdmin common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const AuthTestRemovedAdminEventName = "RemovedAdmin"

// ContractEventName returns the user-defined event name.
func (AuthTestRemovedAdmin) ContractEventName() string {
	return AuthTestRemovedAdminEventName
}

// UnpackRemovedAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (authTest *AuthTest) UnpackRemovedAdminEvent(log *types.Log) (*AuthTestRemovedAdmin, error) {
	event := "RemovedAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestRemovedAdmin)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLog represents a log event raised by the AuthTest contract.
type AuthTestLog struct {
	Arg0 string
	Raw  *types.Log // Blockchain specific contextual infos
}

const AuthTestLogEventName = "log"

// ContractEventName returns the user-defined event name.
func (AuthTestLog) ContractEventName() string {
	return AuthTestLogEventName
}

// UnpackLogEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log(string arg0)
func (authTest *AuthTest) UnpackLogEvent(log *types.Log) (*AuthTestLog, error) {
	event := "log"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLog)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogAddress represents a log_address event raised by the AuthTest contract.
type AuthTestLogAddress struct {
	Arg0 common.Address
	Raw  *types.Log // Blockchain specific contextual infos
}

const AuthTestLogAddressEventName = "log_address"

// ContractEventName returns the user-defined event name.
func (AuthTestLogAddress) ContractEventName() string {
	return AuthTestLogAddressEventName
}

// UnpackLogAddressEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_address(address arg0)
func (authTest *AuthTest) UnpackLogAddressEvent(log *types.Log) (*AuthTestLogAddress, error) {
	event := "log_address"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogAddress)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogArray represents a log_array event raised by the AuthTest contract.
type AuthTestLogArray struct {
	Val []*big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const AuthTestLogArrayEventName = "log_array"

// ContractEventName returns the user-defined event name.
func (AuthTestLogArray) ContractEventName() string {
	return AuthTestLogArrayEventName
}

// UnpackLogArrayEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_array(uint256[] val)
func (authTest *AuthTest) UnpackLogArrayEvent(log *types.Log) (*AuthTestLogArray, error) {
	event := "log_array"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogArray)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogArray0 represents a log_array0 event raised by the AuthTest contract.
type AuthTestLogArray0 struct {
	Val []*big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const AuthTestLogArray0EventName = "log_array0"

// ContractEventName returns the user-defined event name.
func (AuthTestLogArray0) ContractEventName() string {
	return AuthTestLogArray0EventName
}

// UnpackLogArray0Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_array(int256[] val)
func (authTest *AuthTest) UnpackLogArray0Event(log *types.Log) (*AuthTestLogArray0, error) {
	event := "log_array0"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogArray0)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogArray1 represents a log_array1 event raised by the AuthTest contract.
type AuthTestLogArray1 struct {
	Val []common.Address
	Raw *types.Log // Blockchain specific contextual infos
}

const AuthTestLogArray1EventName = "log_array1"

// ContractEventName returns the user-defined event name.
func (AuthTestLogArray1) ContractEventName() string {
	return AuthTestLogArray1EventName
}

// UnpackLogArray1Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_array(address[] val)
func (authTest *AuthTest) UnpackLogArray1Event(log *types.Log) (*AuthTestLogArray1, error) {
	event := "log_array1"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogArray1)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogBytes represents a log_bytes event raised by the AuthTest contract.
type AuthTestLogBytes struct {
	Arg0 []byte
	Raw  *types.Log // Blockchain specific contextual infos
}

const AuthTestLogBytesEventName = "log_bytes"

// ContractEventName returns the user-defined event name.
func (AuthTestLogBytes) ContractEventName() string {
	return AuthTestLogBytesEventName
}

// UnpackLogBytesEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_bytes(bytes arg0)
func (authTest *AuthTest) UnpackLogBytesEvent(log *types.Log) (*AuthTestLogBytes, error) {
	event := "log_bytes"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogBytes)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogBytes32 represents a log_bytes32 event raised by the AuthTest contract.
type AuthTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  *types.Log // Blockchain specific contextual infos
}

const AuthTestLogBytes32EventName = "log_bytes32"

// ContractEventName returns the user-defined event name.
func (AuthTestLogBytes32) ContractEventName() string {
	return AuthTestLogBytes32EventName
}

// UnpackLogBytes32Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (authTest *AuthTest) UnpackLogBytes32Event(log *types.Log) (*AuthTestLogBytes32, error) {
	event := "log_bytes32"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogBytes32)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogInt represents a log_int event raised by the AuthTest contract.
type AuthTestLogInt struct {
	Arg0 *big.Int
	Raw  *types.Log // Blockchain specific contextual infos
}

const AuthTestLogIntEventName = "log_int"

// ContractEventName returns the user-defined event name.
func (AuthTestLogInt) ContractEventName() string {
	return AuthTestLogIntEventName
}

// UnpackLogIntEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_int(int256 arg0)
func (authTest *AuthTest) UnpackLogIntEvent(log *types.Log) (*AuthTestLogInt, error) {
	event := "log_int"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogInt)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogNamedAddress represents a log_named_address event raised by the AuthTest contract.
type AuthTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw *types.Log // Blockchain specific contextual infos
}

const AuthTestLogNamedAddressEventName = "log_named_address"

// ContractEventName returns the user-defined event name.
func (AuthTestLogNamedAddress) ContractEventName() string {
	return AuthTestLogNamedAddressEventName
}

// UnpackLogNamedAddressEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_address(string key, address val)
func (authTest *AuthTest) UnpackLogNamedAddressEvent(log *types.Log) (*AuthTestLogNamedAddress, error) {
	event := "log_named_address"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogNamedAddress)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogNamedArray represents a log_named_array event raised by the AuthTest contract.
type AuthTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const AuthTestLogNamedArrayEventName = "log_named_array"

// ContractEventName returns the user-defined event name.
func (AuthTestLogNamedArray) ContractEventName() string {
	return AuthTestLogNamedArrayEventName
}

// UnpackLogNamedArrayEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (authTest *AuthTest) UnpackLogNamedArrayEvent(log *types.Log) (*AuthTestLogNamedArray, error) {
	event := "log_named_array"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogNamedArray)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogNamedArray0 represents a log_named_array0 event raised by the AuthTest contract.
type AuthTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const AuthTestLogNamedArray0EventName = "log_named_array0"

// ContractEventName returns the user-defined event name.
func (AuthTestLogNamedArray0) ContractEventName() string {
	return AuthTestLogNamedArray0EventName
}

// UnpackLogNamedArray0Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_array(string key, int256[] val)
func (authTest *AuthTest) UnpackLogNamedArray0Event(log *types.Log) (*AuthTestLogNamedArray0, error) {
	event := "log_named_array0"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogNamedArray0)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogNamedArray1 represents a log_named_array1 event raised by the AuthTest contract.
type AuthTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw *types.Log // Blockchain specific contextual infos
}

const AuthTestLogNamedArray1EventName = "log_named_array1"

// ContractEventName returns the user-defined event name.
func (AuthTestLogNamedArray1) ContractEventName() string {
	return AuthTestLogNamedArray1EventName
}

// UnpackLogNamedArray1Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_array(string key, address[] val)
func (authTest *AuthTest) UnpackLogNamedArray1Event(log *types.Log) (*AuthTestLogNamedArray1, error) {
	event := "log_named_array1"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogNamedArray1)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogNamedBytes represents a log_named_bytes event raised by the AuthTest contract.
type AuthTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw *types.Log // Blockchain specific contextual infos
}

const AuthTestLogNamedBytesEventName = "log_named_bytes"

// ContractEventName returns the user-defined event name.
func (AuthTestLogNamedBytes) ContractEventName() string {
	return AuthTestLogNamedBytesEventName
}

// UnpackLogNamedBytesEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (authTest *AuthTest) UnpackLogNamedBytesEvent(log *types.Log) (*AuthTestLogNamedBytes, error) {
	event := "log_named_bytes"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogNamedBytes)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogNamedBytes32 represents a log_named_bytes32 event raised by the AuthTest contract.
type AuthTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw *types.Log // Blockchain specific contextual infos
}

const AuthTestLogNamedBytes32EventName = "log_named_bytes32"

// ContractEventName returns the user-defined event name.
func (AuthTestLogNamedBytes32) ContractEventName() string {
	return AuthTestLogNamedBytes32EventName
}

// UnpackLogNamedBytes32Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (authTest *AuthTest) UnpackLogNamedBytes32Event(log *types.Log) (*AuthTestLogNamedBytes32, error) {
	event := "log_named_bytes32"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogNamedBytes32)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogNamedDecimalInt represents a log_named_decimal_int event raised by the AuthTest contract.
type AuthTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const AuthTestLogNamedDecimalIntEventName = "log_named_decimal_int"

// ContractEventName returns the user-defined event name.
func (AuthTestLogNamedDecimalInt) ContractEventName() string {
	return AuthTestLogNamedDecimalIntEventName
}

// UnpackLogNamedDecimalIntEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (authTest *AuthTest) UnpackLogNamedDecimalIntEvent(log *types.Log) (*AuthTestLogNamedDecimalInt, error) {
	event := "log_named_decimal_int"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogNamedDecimalInt)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogNamedDecimalUint represents a log_named_decimal_uint event raised by the AuthTest contract.
type AuthTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const AuthTestLogNamedDecimalUintEventName = "log_named_decimal_uint"

// ContractEventName returns the user-defined event name.
func (AuthTestLogNamedDecimalUint) ContractEventName() string {
	return AuthTestLogNamedDecimalUintEventName
}

// UnpackLogNamedDecimalUintEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (authTest *AuthTest) UnpackLogNamedDecimalUintEvent(log *types.Log) (*AuthTestLogNamedDecimalUint, error) {
	event := "log_named_decimal_uint"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogNamedDecimalUint)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogNamedInt represents a log_named_int event raised by the AuthTest contract.
type AuthTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const AuthTestLogNamedIntEventName = "log_named_int"

// ContractEventName returns the user-defined event name.
func (AuthTestLogNamedInt) ContractEventName() string {
	return AuthTestLogNamedIntEventName
}

// UnpackLogNamedIntEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_int(string key, int256 val)
func (authTest *AuthTest) UnpackLogNamedIntEvent(log *types.Log) (*AuthTestLogNamedInt, error) {
	event := "log_named_int"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogNamedInt)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogNamedString represents a log_named_string event raised by the AuthTest contract.
type AuthTestLogNamedString struct {
	Key string
	Val string
	Raw *types.Log // Blockchain specific contextual infos
}

const AuthTestLogNamedStringEventName = "log_named_string"

// ContractEventName returns the user-defined event name.
func (AuthTestLogNamedString) ContractEventName() string {
	return AuthTestLogNamedStringEventName
}

// UnpackLogNamedStringEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_string(string key, string val)
func (authTest *AuthTest) UnpackLogNamedStringEvent(log *types.Log) (*AuthTestLogNamedString, error) {
	event := "log_named_string"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogNamedString)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogNamedUint represents a log_named_uint event raised by the AuthTest contract.
type AuthTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const AuthTestLogNamedUintEventName = "log_named_uint"

// ContractEventName returns the user-defined event name.
func (AuthTestLogNamedUint) ContractEventName() string {
	return AuthTestLogNamedUintEventName
}

// UnpackLogNamedUintEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (authTest *AuthTest) UnpackLogNamedUintEvent(log *types.Log) (*AuthTestLogNamedUint, error) {
	event := "log_named_uint"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogNamedUint)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogString represents a log_string event raised by the AuthTest contract.
type AuthTestLogString struct {
	Arg0 string
	Raw  *types.Log // Blockchain specific contextual infos
}

const AuthTestLogStringEventName = "log_string"

// ContractEventName returns the user-defined event name.
func (AuthTestLogString) ContractEventName() string {
	return AuthTestLogStringEventName
}

// UnpackLogStringEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_string(string arg0)
func (authTest *AuthTest) UnpackLogStringEvent(log *types.Log) (*AuthTestLogString, error) {
	event := "log_string"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogString)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogUint represents a log_uint event raised by the AuthTest contract.
type AuthTestLogUint struct {
	Arg0 *big.Int
	Raw  *types.Log // Blockchain specific contextual infos
}

const AuthTestLogUintEventName = "log_uint"

// ContractEventName returns the user-defined event name.
func (AuthTestLogUint) ContractEventName() string {
	return AuthTestLogUintEventName
}

// UnpackLogUintEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_uint(uint256 arg0)
func (authTest *AuthTest) UnpackLogUintEvent(log *types.Log) (*AuthTestLogUint, error) {
	event := "log_uint"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogUint)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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

// AuthTestLogs represents a logs event raised by the AuthTest contract.
type AuthTestLogs struct {
	Arg0 []byte
	Raw  *types.Log // Blockchain specific contextual infos
}

const AuthTestLogsEventName = "logs"

// ContractEventName returns the user-defined event name.
func (AuthTestLogs) ContractEventName() string {
	return AuthTestLogsEventName
}

// UnpackLogsEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event logs(bytes arg0)
func (authTest *AuthTest) UnpackLogsEvent(log *types.Log) (*AuthTestLogs, error) {
	event := "logs"
	if len(log.Topics) == 0 || log.Topics[0] != authTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AuthTestLogs)
	if len(log.Data) > 0 {
		if err := authTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range authTest.abi.Events[event].Inputs {
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
func (authTest *AuthTest) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], authTest.abi.Errors["NotAdmin"].ID.Bytes()[:4]) {
		return authTest.UnpackNotAdminError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// AuthTestNotAdmin represents a NotAdmin error raised by the AuthTest contract.
type AuthTestNotAdmin struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAdmin()
func AuthTestNotAdminErrorID() common.Hash {
	return common.HexToHash("0x7bfa4b9fb0cd3687c1f539f384b3f3f258f2c9aa9186353d0815413b508ed97d")
}

// UnpackNotAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAdmin()
func (authTest *AuthTest) UnpackNotAdminError(raw []byte) (*AuthTestNotAdmin, error) {
	out := new(AuthTestNotAdmin)
	if err := authTest.abi.UnpackIntoInterface(out, "NotAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

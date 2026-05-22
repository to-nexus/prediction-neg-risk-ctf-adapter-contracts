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

// CTHelpersTestMetaData contains all meta data concerning the CTHelpersTest contract.
var CTHelpersTestMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IS_TEST\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeArtifacts\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"excludedArtifacts_\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"excludedContracts_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"excludedSenders_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetArtifactSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"artifact\",\"type\":\"string\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetArtifacts\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"targetedArtifacts_\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targetedContracts_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetInterfaces\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"artifacts\",\"type\":\"string[]\"}],\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targetedSenders_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parent\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexSet\",\"type\":\"uint256\"}],\"name\":\"testFuzz_getCollectionId_matchesKeccakFormula\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"}],\"name\":\"testFuzz_getConditionId_matchesKeccakFormula\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"collectionId\",\"type\":\"bytes32\"}],\"name\":\"testFuzz_getPositionId_matchesKeccakFormula\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_getCollectionId_matchesKeccakFormula_nonZeroParent\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_getCollectionId_matchesKeccakFormula_parentZero_indexNo\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_getCollectionId_matchesKeccakFormula_parentZero_indexYes\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_getConditionId_matchesKeccakFormula_fixedVector\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test_getPositionId_matchesKeccakFormula_fixedVector\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"log_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"log_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"log_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"log_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"log_named_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"log_named_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"log_named_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"}],\"name\":\"log_named_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"log_named_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"log_named_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"log_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"logs\",\"type\":\"event\"}]",
	ID:  "CTHelpersTest",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602b575f5ffd5b50611861806100395f395ff3fe608060405234801561000f575f5ffd5b5060043610610163575f3560e01c806385226c81116100c7578063c37c15901161007d578063e20c9f7111610063578063e20c9f7114610261578063e950b36014610269578063fa7626d41461027c575f5ffd5b8063c37c159014610251578063cc778aa614610259575f5ffd5b8063b0464fdc116100ad578063b0464fdc14610229578063b5508aa914610231578063ba414fa614610239575f5ffd5b806385226c81146101ff578063916a17c614610214575f5ffd5b80633e5e3c231161011c5780635ba7781f116101025780635ba7781f146101cf57806365f0cb0a146101d757806366d9a9a0146101ea575f5ffd5b80633e5e3c23146101bf5780633f7286f4146101c7575f5ffd5b80631ed7831c1161014c5780631ed7831c146101795780632ade388014610197578063378ec141146101ac575f5ffd5b806302b11659146101675780631bef2b9e14610171575b5f5ffd5b61016f610289565b005b61016f61033d565b6101816103fb565b60405161018e9190611235565b60405180910390f35b61019f610468565b60405161018e91906112d9565b61016f6101ba3660046113eb565b6105b1565b6101816105fe565b610181610669565b61016f6106d4565b61016f6101e536600461143c565b610765565b6101f26107e9565b60405161018e91906114c0565b610207610962565b60405161018e919061155c565b61021c610a2d565b60405161018e91906115d1565b61021c610b30565b610207610c33565b610241610cfe565b604051901515815260200161018e565b61016f610dca565b61016f610e3c565b610181610f0d565b61016f610277366004611673565b610f78565b601f546102419060ff1681565b604080517fb5926f6a229bb8886d893d6c97c2f94d68a7ecdf2aa6d562e246a32d71f04d59602082018190527fb4e04d4580934c1bbc717556f6f05239e7acddb390f3c4d984d858d7ebb883349282018390526007606083018190529092915f906080016040516020818303038152906040528051906020012090505f610311858585610fde565b905061033681836040518060600160405280603d8152602001611771603d913961101b565b5050505050565b6040517fcafebabecafebabecafebabecafebabecafebabe00000000000000000000000060208201527f564c367f9a7f9ff4c7d911e0c844b505b3e5b5c35b099d979417fb3f59d94ce66034820181905260026054830181905273cafebabecafebabecafebabecafebabecafebabe925f906074016040516020818303038152906040528051906020012090505f6103d68585856110a8565b905061033681836040518060600160405280604081526020016117ae6040913961101b565b6060601680548060200260200160405190810160405280929190818152602001828054801561045e57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610433575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020015f905b828210156105a8575f848152602080822060408051808201825260028702909201805473ffffffffffffffffffffffffffffffffffffffff168352600181018054835181870281018701909452808452939591948681019491929084015b82821015610591578382905f5260205f20018054610506906116a3565b80601f0160208091040260200160405190810160405280929190818152602001828054610532906116a3565b801561057d5780601f106105545761010080835404028352916020019161057d565b820191905f5260205f20905b81548152906001019060200180831161056057829003601f168201915b5050505050815260200190600101906104e9565b50505050815250508152602001906001019061048b565b50505050905090565b6040805160208101859052908101839052606081018290525f906080016040516020818303038152906040528051906020012090505f6105f2858585610fde565b905061033681836110ef565b6060601880548060200260200160405190810160405280929190818152602001828054801561045e57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610433575050505050905090565b6060601780548060200260200160405190810160405280929190818152602001828054801561045e57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610433575050505050905090565b604080515f602082018190527fab9873d25e00587e46b953f5f5dd6863789d814f2ce2688c18ad8187eef81825928201839052600160608301819052916080016040516020818303038152906040528051906020012090505f61073a5f5f1b8585610fde565b905061075f81836040518060600160405280603e81526020016117ee603e913961101b565b50505050565b60408051606084901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166020808301829052603480840186905284518085039091018152605484018552805190820120607484019290925260888084018690528451808503909101815260a89093019093528151919092012061075f818361117b565b6060601b805480602002602001604051908101604052809291908181526020015f905b828210156105a8578382905f5260205f2090600202016040518060400160405290815f8201805461083c906116a3565b80601f0160208091040260200160405190810160405280929190818152602001828054610868906116a3565b80156108b35780601f1061088a576101008083540402835291602001916108b3565b820191905f5260205f20905b81548152906001019060200180831161089657829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561094a57602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116108f75790505b5050505050815250508152602001906001019061080c565b6060601a805480602002602001604051908101604052809291908181526020015f905b828210156105a8578382905f5260205f200180546109a2906116a3565b80601f01602080910402602001604051908101604052809291908181526020018280546109ce906116a3565b8015610a195780601f106109f057610100808354040283529160200191610a19565b820191905f5260205f20905b8154815290600101906020018083116109fc57829003601f168201915b505050505081526020019060010190610985565b6060601d805480602002602001604051908101604052809291908181526020015f905b828210156105a8575f84815260209081902060408051808201825260028602909201805473ffffffffffffffffffffffffffffffffffffffff168352600181018054835181870281018701909452808452939491938583019392830182828015610b1857602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411610ac55790505b50505050508152505081526020019060010190610a50565b6060601c805480602002602001604051908101604052809291908181526020015f905b828210156105a8575f84815260209081902060408051808201825260028602909201805473ffffffffffffffffffffffffffffffffffffffff168352600181018054835181870281018701909452808452939491938583019392830182828015610c1b57602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411610bc85790505b50505050508152505081526020019060010190610b53565b60606019805480602002602001604051908101604052809291908181526020015f905b828210156105a8578382905f5260205f20018054610c73906116a3565b80601f0160208091040260200160405190810160405280929190818152602001828054610c9f906116a3565b8015610cea5780601f10610cc157610100808354040283529160200191610cea565b820191905f5260205f20905b815481529060010190602001808311610ccd57829003601f168201915b505050505081526020019060010190610c56565b6008545f9060ff1615610d115750600190565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c6564000000000000000000000000000000000000000000000000000060248301525f9163667f9d7090604401602060405180830381865afa158015610d9f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dc391906116f4565b1415905090565b604080515f602082018190527fab9873d25e00587e46b953f5f5dd6863789d814f2ce2688c18ad8187eef81825928201839052600260608301819052916080016040516020818303038152906040528051906020012090505f610e305f5f1b8585610fde565b905061075f81836110ef565b604080517fdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef00000000000000000000000060208083018290527fd5a7feb54292f5faa3c35d065adf9076ee73b0b648f8ec30a49f410327f281e4603480850182905285518086039091018152605485018652805190830120607485019390935260888085018290528551808603909101815260a8909401909452825192019190912073deadbeefdeadbeefdeadbeefdeadbeefdeadbeef92919061075f81836040518060600160405280603e8152602001611733603e91396111da565b6060601580548060200260200160405190810160405280929190818152602001828054801561045e57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610433575050505050905090565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b16602082015260348101839052605481018290525f906074016040516020818303038152906040528051906020012090505f6105f28585856110a8565b6040805160208101859052908101839052606081018290525f906080015b6040516020818303038152906040528051906020012090509392505050565b8183146110a3576040517fc1fa1ed0000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c1fa1ed0906110769086908690869060040161170b565b5f6040518083038186803b15801561108c575f5ffd5b505afa15801561109e573d5f5f3e3d5ffd5b505050505b505050565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b16602082015260348101839052605481018290525f90607401610ffc565b808214611177576040517f7c84c69b0000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d90637c84c69b906044015b5f6040518083038186803b158015611160575f5ffd5b505afa158015611172573d5f5f3e3d5ffd5b505050505b5050565b808214611177576040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c549060440161114a565b8183146110a3576040517f88b44c85000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d906388b44c85906110769086908690869060040161170b565b602080825282518282018190525f918401906040840190835b8181101561128257835173ffffffffffffffffffffffffffffffffffffffff1683526020938401939092019160010161124e565b509095945050505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156113df577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805173ffffffffffffffffffffffffffffffffffffffff168652602090810151604082880181905281519088018190529101906060600582901b8801810191908801905f5b818110156113c5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a85030183526113af84865161128d565b6020958601959094509290920191600101611375565b5091975050506020948501949290920191506001016112ff565b50929695505050505050565b5f5f5f606084860312156113fd575f5ffd5b505081359360208301359350604090920135919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611437575f5ffd5b919050565b5f5f6040838503121561144d575f5ffd5b61145683611414565b946020939093013593505050565b5f8151808452602084019350602083015f5b828110156114b65781517fffffffff0000000000000000000000000000000000000000000000000000000016865260209586019590910190600101611476565b5093949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156113df577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261152a604088018261128d565b90506020820151915086810360208801526115458183611464565b9650505060209384019391909101906001016114e6565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156113df577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184526115bc85835161128d565b94506020938401939190910190600101611582565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156113df577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815173ffffffffffffffffffffffffffffffffffffffff8151168652602081015190506040602087015261165d6040870182611464565b95505060209384019391909101906001016115f7565b5f5f5f60608486031215611685575f5ffd5b61168e84611414565b95602085013595506040909401359392505050565b600181811c908216806116b757607f821691505b6020821081036116ee577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f60208284031215611704575f5ffd5b5051919050565b838152826020820152606060408201525f611729606083018461128d565b9594505050505056fe676574506f736974696f6e4964206d757374206d61746368207465616d204354205f676574506f736974696f6e4964206b656363616b20666f726d756c614e6f6e2d7a65726f20706172656e74206d75737420616c736f2075736520666c6174206b656363616b2c206e6f7420454320636f6d706f736974696f6e676574436f6e646974696f6e4964206d757374206d61746368207465616d204354205f676574436f6e646974696f6e4964206b656363616b20666f726d756c61676574436f6c6c656374696f6e496428302c206369642c203129206d757374206d61746368207465616d204354205f676574436f6c6c656374696f6e4964a2646970667358221220d9d8dd6cc7903fc61183f6b187874f3b9926852a34cd9db526674e68af2720de64736f6c634300081e0033",
}

// CTHelpersTest is an auto generated Go binding around an Ethereum contract.
type CTHelpersTest struct {
	abi abi.ABI
}

// NewCTHelpersTest creates a new instance of CTHelpersTest.
func NewCTHelpersTest() *CTHelpersTest {
	parsed, err := CTHelpersTestMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &CTHelpersTest{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *CTHelpersTest) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackISTEST is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa7626d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function IS_TEST() view returns(bool)
func (cTHelpersTest *CTHelpersTest) PackISTEST() []byte {
	enc, err := cTHelpersTest.abi.Pack("IS_TEST")
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
func (cTHelpersTest *CTHelpersTest) TryPackISTEST() ([]byte, error) {
	return cTHelpersTest.abi.Pack("IS_TEST")
}

// UnpackISTEST is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (cTHelpersTest *CTHelpersTest) UnpackISTEST(data []byte) (bool, error) {
	out, err := cTHelpersTest.abi.Unpack("IS_TEST", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackExcludeArtifacts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb5508aa9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (cTHelpersTest *CTHelpersTest) PackExcludeArtifacts() []byte {
	enc, err := cTHelpersTest.abi.Pack("excludeArtifacts")
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
func (cTHelpersTest *CTHelpersTest) TryPackExcludeArtifacts() ([]byte, error) {
	return cTHelpersTest.abi.Pack("excludeArtifacts")
}

// UnpackExcludeArtifacts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (cTHelpersTest *CTHelpersTest) UnpackExcludeArtifacts(data []byte) ([]string, error) {
	out, err := cTHelpersTest.abi.Unpack("excludeArtifacts", data)
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
func (cTHelpersTest *CTHelpersTest) PackExcludeContracts() []byte {
	enc, err := cTHelpersTest.abi.Pack("excludeContracts")
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
func (cTHelpersTest *CTHelpersTest) TryPackExcludeContracts() ([]byte, error) {
	return cTHelpersTest.abi.Pack("excludeContracts")
}

// UnpackExcludeContracts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (cTHelpersTest *CTHelpersTest) UnpackExcludeContracts(data []byte) ([]common.Address, error) {
	out, err := cTHelpersTest.abi.Unpack("excludeContracts", data)
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
func (cTHelpersTest *CTHelpersTest) PackExcludeSelectors() []byte {
	enc, err := cTHelpersTest.abi.Pack("excludeSelectors")
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
func (cTHelpersTest *CTHelpersTest) TryPackExcludeSelectors() ([]byte, error) {
	return cTHelpersTest.abi.Pack("excludeSelectors")
}

// UnpackExcludeSelectors is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (cTHelpersTest *CTHelpersTest) UnpackExcludeSelectors(data []byte) ([]StdInvariantFuzzSelector, error) {
	out, err := cTHelpersTest.abi.Unpack("excludeSelectors", data)
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
func (cTHelpersTest *CTHelpersTest) PackExcludeSenders() []byte {
	enc, err := cTHelpersTest.abi.Pack("excludeSenders")
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
func (cTHelpersTest *CTHelpersTest) TryPackExcludeSenders() ([]byte, error) {
	return cTHelpersTest.abi.Pack("excludeSenders")
}

// UnpackExcludeSenders is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (cTHelpersTest *CTHelpersTest) UnpackExcludeSenders(data []byte) ([]common.Address, error) {
	out, err := cTHelpersTest.abi.Unpack("excludeSenders", data)
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
func (cTHelpersTest *CTHelpersTest) PackFailed() []byte {
	enc, err := cTHelpersTest.abi.Pack("failed")
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
func (cTHelpersTest *CTHelpersTest) TryPackFailed() ([]byte, error) {
	return cTHelpersTest.abi.Pack("failed")
}

// UnpackFailed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (cTHelpersTest *CTHelpersTest) UnpackFailed(data []byte) (bool, error) {
	out, err := cTHelpersTest.abi.Unpack("failed", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTargetArtifactSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66d9a9a0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (cTHelpersTest *CTHelpersTest) PackTargetArtifactSelectors() []byte {
	enc, err := cTHelpersTest.abi.Pack("targetArtifactSelectors")
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
func (cTHelpersTest *CTHelpersTest) TryPackTargetArtifactSelectors() ([]byte, error) {
	return cTHelpersTest.abi.Pack("targetArtifactSelectors")
}

// UnpackTargetArtifactSelectors is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (cTHelpersTest *CTHelpersTest) UnpackTargetArtifactSelectors(data []byte) ([]StdInvariantFuzzArtifactSelector, error) {
	out, err := cTHelpersTest.abi.Unpack("targetArtifactSelectors", data)
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
func (cTHelpersTest *CTHelpersTest) PackTargetArtifacts() []byte {
	enc, err := cTHelpersTest.abi.Pack("targetArtifacts")
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
func (cTHelpersTest *CTHelpersTest) TryPackTargetArtifacts() ([]byte, error) {
	return cTHelpersTest.abi.Pack("targetArtifacts")
}

// UnpackTargetArtifacts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (cTHelpersTest *CTHelpersTest) UnpackTargetArtifacts(data []byte) ([]string, error) {
	out, err := cTHelpersTest.abi.Unpack("targetArtifacts", data)
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
func (cTHelpersTest *CTHelpersTest) PackTargetContracts() []byte {
	enc, err := cTHelpersTest.abi.Pack("targetContracts")
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
func (cTHelpersTest *CTHelpersTest) TryPackTargetContracts() ([]byte, error) {
	return cTHelpersTest.abi.Pack("targetContracts")
}

// UnpackTargetContracts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (cTHelpersTest *CTHelpersTest) UnpackTargetContracts(data []byte) ([]common.Address, error) {
	out, err := cTHelpersTest.abi.Unpack("targetContracts", data)
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
func (cTHelpersTest *CTHelpersTest) PackTargetInterfaces() []byte {
	enc, err := cTHelpersTest.abi.Pack("targetInterfaces")
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
func (cTHelpersTest *CTHelpersTest) TryPackTargetInterfaces() ([]byte, error) {
	return cTHelpersTest.abi.Pack("targetInterfaces")
}

// UnpackTargetInterfaces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (cTHelpersTest *CTHelpersTest) UnpackTargetInterfaces(data []byte) ([]StdInvariantFuzzInterface, error) {
	out, err := cTHelpersTest.abi.Unpack("targetInterfaces", data)
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
func (cTHelpersTest *CTHelpersTest) PackTargetSelectors() []byte {
	enc, err := cTHelpersTest.abi.Pack("targetSelectors")
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
func (cTHelpersTest *CTHelpersTest) TryPackTargetSelectors() ([]byte, error) {
	return cTHelpersTest.abi.Pack("targetSelectors")
}

// UnpackTargetSelectors is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (cTHelpersTest *CTHelpersTest) UnpackTargetSelectors(data []byte) ([]StdInvariantFuzzSelector, error) {
	out, err := cTHelpersTest.abi.Unpack("targetSelectors", data)
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
func (cTHelpersTest *CTHelpersTest) PackTargetSenders() []byte {
	enc, err := cTHelpersTest.abi.Pack("targetSenders")
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
func (cTHelpersTest *CTHelpersTest) TryPackTargetSenders() ([]byte, error) {
	return cTHelpersTest.abi.Pack("targetSenders")
}

// UnpackTargetSenders is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (cTHelpersTest *CTHelpersTest) UnpackTargetSenders(data []byte) ([]common.Address, error) {
	out, err := cTHelpersTest.abi.Unpack("targetSenders", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackTestFuzzGetCollectionIdMatchesKeccakFormula is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x378ec141.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function testFuzz_getCollectionId_matchesKeccakFormula(bytes32 parent, bytes32 conditionId, uint256 indexSet) pure returns()
func (cTHelpersTest *CTHelpersTest) PackTestFuzzGetCollectionIdMatchesKeccakFormula(parent [32]byte, conditionId [32]byte, indexSet *big.Int) []byte {
	enc, err := cTHelpersTest.abi.Pack("testFuzz_getCollectionId_matchesKeccakFormula", parent, conditionId, indexSet)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestFuzzGetCollectionIdMatchesKeccakFormula is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x378ec141.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function testFuzz_getCollectionId_matchesKeccakFormula(bytes32 parent, bytes32 conditionId, uint256 indexSet) pure returns()
func (cTHelpersTest *CTHelpersTest) TryPackTestFuzzGetCollectionIdMatchesKeccakFormula(parent [32]byte, conditionId [32]byte, indexSet *big.Int) ([]byte, error) {
	return cTHelpersTest.abi.Pack("testFuzz_getCollectionId_matchesKeccakFormula", parent, conditionId, indexSet)
}

// PackTestFuzzGetConditionIdMatchesKeccakFormula is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe950b360.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function testFuzz_getConditionId_matchesKeccakFormula(address oracle, bytes32 questionId, uint256 outcomeSlotCount) pure returns()
func (cTHelpersTest *CTHelpersTest) PackTestFuzzGetConditionIdMatchesKeccakFormula(oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) []byte {
	enc, err := cTHelpersTest.abi.Pack("testFuzz_getConditionId_matchesKeccakFormula", oracle, questionId, outcomeSlotCount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestFuzzGetConditionIdMatchesKeccakFormula is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe950b360.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function testFuzz_getConditionId_matchesKeccakFormula(address oracle, bytes32 questionId, uint256 outcomeSlotCount) pure returns()
func (cTHelpersTest *CTHelpersTest) TryPackTestFuzzGetConditionIdMatchesKeccakFormula(oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) ([]byte, error) {
	return cTHelpersTest.abi.Pack("testFuzz_getConditionId_matchesKeccakFormula", oracle, questionId, outcomeSlotCount)
}

// PackTestFuzzGetPositionIdMatchesKeccakFormula is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x65f0cb0a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function testFuzz_getPositionId_matchesKeccakFormula(address collateral, bytes32 collectionId) pure returns()
func (cTHelpersTest *CTHelpersTest) PackTestFuzzGetPositionIdMatchesKeccakFormula(collateral common.Address, collectionId [32]byte) []byte {
	enc, err := cTHelpersTest.abi.Pack("testFuzz_getPositionId_matchesKeccakFormula", collateral, collectionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestFuzzGetPositionIdMatchesKeccakFormula is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x65f0cb0a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function testFuzz_getPositionId_matchesKeccakFormula(address collateral, bytes32 collectionId) pure returns()
func (cTHelpersTest *CTHelpersTest) TryPackTestFuzzGetPositionIdMatchesKeccakFormula(collateral common.Address, collectionId [32]byte) ([]byte, error) {
	return cTHelpersTest.abi.Pack("testFuzz_getPositionId_matchesKeccakFormula", collateral, collectionId)
}

// PackTestGetCollectionIdMatchesKeccakFormulaNonZeroParent is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02b11659.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_getCollectionId_matchesKeccakFormula_nonZeroParent() pure returns()
func (cTHelpersTest *CTHelpersTest) PackTestGetCollectionIdMatchesKeccakFormulaNonZeroParent() []byte {
	enc, err := cTHelpersTest.abi.Pack("test_getCollectionId_matchesKeccakFormula_nonZeroParent")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestGetCollectionIdMatchesKeccakFormulaNonZeroParent is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02b11659.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_getCollectionId_matchesKeccakFormula_nonZeroParent() pure returns()
func (cTHelpersTest *CTHelpersTest) TryPackTestGetCollectionIdMatchesKeccakFormulaNonZeroParent() ([]byte, error) {
	return cTHelpersTest.abi.Pack("test_getCollectionId_matchesKeccakFormula_nonZeroParent")
}

// PackTestGetCollectionIdMatchesKeccakFormulaParentZeroIndexNo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc37c1590.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_getCollectionId_matchesKeccakFormula_parentZero_indexNo() pure returns()
func (cTHelpersTest *CTHelpersTest) PackTestGetCollectionIdMatchesKeccakFormulaParentZeroIndexNo() []byte {
	enc, err := cTHelpersTest.abi.Pack("test_getCollectionId_matchesKeccakFormula_parentZero_indexNo")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestGetCollectionIdMatchesKeccakFormulaParentZeroIndexNo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc37c1590.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_getCollectionId_matchesKeccakFormula_parentZero_indexNo() pure returns()
func (cTHelpersTest *CTHelpersTest) TryPackTestGetCollectionIdMatchesKeccakFormulaParentZeroIndexNo() ([]byte, error) {
	return cTHelpersTest.abi.Pack("test_getCollectionId_matchesKeccakFormula_parentZero_indexNo")
}

// PackTestGetCollectionIdMatchesKeccakFormulaParentZeroIndexYes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ba7781f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_getCollectionId_matchesKeccakFormula_parentZero_indexYes() pure returns()
func (cTHelpersTest *CTHelpersTest) PackTestGetCollectionIdMatchesKeccakFormulaParentZeroIndexYes() []byte {
	enc, err := cTHelpersTest.abi.Pack("test_getCollectionId_matchesKeccakFormula_parentZero_indexYes")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestGetCollectionIdMatchesKeccakFormulaParentZeroIndexYes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ba7781f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_getCollectionId_matchesKeccakFormula_parentZero_indexYes() pure returns()
func (cTHelpersTest *CTHelpersTest) TryPackTestGetCollectionIdMatchesKeccakFormulaParentZeroIndexYes() ([]byte, error) {
	return cTHelpersTest.abi.Pack("test_getCollectionId_matchesKeccakFormula_parentZero_indexYes")
}

// PackTestGetConditionIdMatchesKeccakFormulaFixedVector is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1bef2b9e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_getConditionId_matchesKeccakFormula_fixedVector() pure returns()
func (cTHelpersTest *CTHelpersTest) PackTestGetConditionIdMatchesKeccakFormulaFixedVector() []byte {
	enc, err := cTHelpersTest.abi.Pack("test_getConditionId_matchesKeccakFormula_fixedVector")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestGetConditionIdMatchesKeccakFormulaFixedVector is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1bef2b9e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_getConditionId_matchesKeccakFormula_fixedVector() pure returns()
func (cTHelpersTest *CTHelpersTest) TryPackTestGetConditionIdMatchesKeccakFormulaFixedVector() ([]byte, error) {
	return cTHelpersTest.abi.Pack("test_getConditionId_matchesKeccakFormula_fixedVector")
}

// PackTestGetPositionIdMatchesKeccakFormulaFixedVector is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcc778aa6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function test_getPositionId_matchesKeccakFormula_fixedVector() pure returns()
func (cTHelpersTest *CTHelpersTest) PackTestGetPositionIdMatchesKeccakFormulaFixedVector() []byte {
	enc, err := cTHelpersTest.abi.Pack("test_getPositionId_matchesKeccakFormula_fixedVector")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTestGetPositionIdMatchesKeccakFormulaFixedVector is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcc778aa6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function test_getPositionId_matchesKeccakFormula_fixedVector() pure returns()
func (cTHelpersTest *CTHelpersTest) TryPackTestGetPositionIdMatchesKeccakFormulaFixedVector() ([]byte, error) {
	return cTHelpersTest.abi.Pack("test_getPositionId_matchesKeccakFormula_fixedVector")
}

// CTHelpersTestLog represents a log event raised by the CTHelpersTest contract.
type CTHelpersTestLog struct {
	Arg0 string
	Raw  *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogEventName = "log"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLog) ContractEventName() string {
	return CTHelpersTestLogEventName
}

// UnpackLogEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log(string arg0)
func (cTHelpersTest *CTHelpersTest) UnpackLogEvent(log *types.Log) (*CTHelpersTestLog, error) {
	event := "log"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLog)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogAddress represents a log_address event raised by the CTHelpersTest contract.
type CTHelpersTestLogAddress struct {
	Arg0 common.Address
	Raw  *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogAddressEventName = "log_address"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogAddress) ContractEventName() string {
	return CTHelpersTestLogAddressEventName
}

// UnpackLogAddressEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_address(address arg0)
func (cTHelpersTest *CTHelpersTest) UnpackLogAddressEvent(log *types.Log) (*CTHelpersTestLogAddress, error) {
	event := "log_address"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogAddress)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogArray represents a log_array event raised by the CTHelpersTest contract.
type CTHelpersTestLogArray struct {
	Val []*big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogArrayEventName = "log_array"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogArray) ContractEventName() string {
	return CTHelpersTestLogArrayEventName
}

// UnpackLogArrayEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_array(uint256[] val)
func (cTHelpersTest *CTHelpersTest) UnpackLogArrayEvent(log *types.Log) (*CTHelpersTestLogArray, error) {
	event := "log_array"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogArray)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogArray0 represents a log_array0 event raised by the CTHelpersTest contract.
type CTHelpersTestLogArray0 struct {
	Val []*big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogArray0EventName = "log_array0"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogArray0) ContractEventName() string {
	return CTHelpersTestLogArray0EventName
}

// UnpackLogArray0Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_array(int256[] val)
func (cTHelpersTest *CTHelpersTest) UnpackLogArray0Event(log *types.Log) (*CTHelpersTestLogArray0, error) {
	event := "log_array0"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogArray0)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogArray1 represents a log_array1 event raised by the CTHelpersTest contract.
type CTHelpersTestLogArray1 struct {
	Val []common.Address
	Raw *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogArray1EventName = "log_array1"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogArray1) ContractEventName() string {
	return CTHelpersTestLogArray1EventName
}

// UnpackLogArray1Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_array(address[] val)
func (cTHelpersTest *CTHelpersTest) UnpackLogArray1Event(log *types.Log) (*CTHelpersTestLogArray1, error) {
	event := "log_array1"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogArray1)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogBytes represents a log_bytes event raised by the CTHelpersTest contract.
type CTHelpersTestLogBytes struct {
	Arg0 []byte
	Raw  *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogBytesEventName = "log_bytes"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogBytes) ContractEventName() string {
	return CTHelpersTestLogBytesEventName
}

// UnpackLogBytesEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_bytes(bytes arg0)
func (cTHelpersTest *CTHelpersTest) UnpackLogBytesEvent(log *types.Log) (*CTHelpersTestLogBytes, error) {
	event := "log_bytes"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogBytes)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogBytes32 represents a log_bytes32 event raised by the CTHelpersTest contract.
type CTHelpersTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogBytes32EventName = "log_bytes32"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogBytes32) ContractEventName() string {
	return CTHelpersTestLogBytes32EventName
}

// UnpackLogBytes32Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (cTHelpersTest *CTHelpersTest) UnpackLogBytes32Event(log *types.Log) (*CTHelpersTestLogBytes32, error) {
	event := "log_bytes32"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogBytes32)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogInt represents a log_int event raised by the CTHelpersTest contract.
type CTHelpersTestLogInt struct {
	Arg0 *big.Int
	Raw  *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogIntEventName = "log_int"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogInt) ContractEventName() string {
	return CTHelpersTestLogIntEventName
}

// UnpackLogIntEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_int(int256 arg0)
func (cTHelpersTest *CTHelpersTest) UnpackLogIntEvent(log *types.Log) (*CTHelpersTestLogInt, error) {
	event := "log_int"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogInt)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogNamedAddress represents a log_named_address event raised by the CTHelpersTest contract.
type CTHelpersTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogNamedAddressEventName = "log_named_address"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogNamedAddress) ContractEventName() string {
	return CTHelpersTestLogNamedAddressEventName
}

// UnpackLogNamedAddressEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_address(string key, address val)
func (cTHelpersTest *CTHelpersTest) UnpackLogNamedAddressEvent(log *types.Log) (*CTHelpersTestLogNamedAddress, error) {
	event := "log_named_address"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogNamedAddress)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogNamedArray represents a log_named_array event raised by the CTHelpersTest contract.
type CTHelpersTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogNamedArrayEventName = "log_named_array"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogNamedArray) ContractEventName() string {
	return CTHelpersTestLogNamedArrayEventName
}

// UnpackLogNamedArrayEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (cTHelpersTest *CTHelpersTest) UnpackLogNamedArrayEvent(log *types.Log) (*CTHelpersTestLogNamedArray, error) {
	event := "log_named_array"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogNamedArray)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogNamedArray0 represents a log_named_array0 event raised by the CTHelpersTest contract.
type CTHelpersTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogNamedArray0EventName = "log_named_array0"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogNamedArray0) ContractEventName() string {
	return CTHelpersTestLogNamedArray0EventName
}

// UnpackLogNamedArray0Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_array(string key, int256[] val)
func (cTHelpersTest *CTHelpersTest) UnpackLogNamedArray0Event(log *types.Log) (*CTHelpersTestLogNamedArray0, error) {
	event := "log_named_array0"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogNamedArray0)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogNamedArray1 represents a log_named_array1 event raised by the CTHelpersTest contract.
type CTHelpersTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogNamedArray1EventName = "log_named_array1"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogNamedArray1) ContractEventName() string {
	return CTHelpersTestLogNamedArray1EventName
}

// UnpackLogNamedArray1Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_array(string key, address[] val)
func (cTHelpersTest *CTHelpersTest) UnpackLogNamedArray1Event(log *types.Log) (*CTHelpersTestLogNamedArray1, error) {
	event := "log_named_array1"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogNamedArray1)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogNamedBytes represents a log_named_bytes event raised by the CTHelpersTest contract.
type CTHelpersTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogNamedBytesEventName = "log_named_bytes"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogNamedBytes) ContractEventName() string {
	return CTHelpersTestLogNamedBytesEventName
}

// UnpackLogNamedBytesEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (cTHelpersTest *CTHelpersTest) UnpackLogNamedBytesEvent(log *types.Log) (*CTHelpersTestLogNamedBytes, error) {
	event := "log_named_bytes"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogNamedBytes)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogNamedBytes32 represents a log_named_bytes32 event raised by the CTHelpersTest contract.
type CTHelpersTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogNamedBytes32EventName = "log_named_bytes32"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogNamedBytes32) ContractEventName() string {
	return CTHelpersTestLogNamedBytes32EventName
}

// UnpackLogNamedBytes32Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (cTHelpersTest *CTHelpersTest) UnpackLogNamedBytes32Event(log *types.Log) (*CTHelpersTestLogNamedBytes32, error) {
	event := "log_named_bytes32"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogNamedBytes32)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogNamedDecimalInt represents a log_named_decimal_int event raised by the CTHelpersTest contract.
type CTHelpersTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogNamedDecimalIntEventName = "log_named_decimal_int"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogNamedDecimalInt) ContractEventName() string {
	return CTHelpersTestLogNamedDecimalIntEventName
}

// UnpackLogNamedDecimalIntEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (cTHelpersTest *CTHelpersTest) UnpackLogNamedDecimalIntEvent(log *types.Log) (*CTHelpersTestLogNamedDecimalInt, error) {
	event := "log_named_decimal_int"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogNamedDecimalInt)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogNamedDecimalUint represents a log_named_decimal_uint event raised by the CTHelpersTest contract.
type CTHelpersTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogNamedDecimalUintEventName = "log_named_decimal_uint"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogNamedDecimalUint) ContractEventName() string {
	return CTHelpersTestLogNamedDecimalUintEventName
}

// UnpackLogNamedDecimalUintEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (cTHelpersTest *CTHelpersTest) UnpackLogNamedDecimalUintEvent(log *types.Log) (*CTHelpersTestLogNamedDecimalUint, error) {
	event := "log_named_decimal_uint"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogNamedDecimalUint)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogNamedInt represents a log_named_int event raised by the CTHelpersTest contract.
type CTHelpersTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogNamedIntEventName = "log_named_int"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogNamedInt) ContractEventName() string {
	return CTHelpersTestLogNamedIntEventName
}

// UnpackLogNamedIntEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_int(string key, int256 val)
func (cTHelpersTest *CTHelpersTest) UnpackLogNamedIntEvent(log *types.Log) (*CTHelpersTestLogNamedInt, error) {
	event := "log_named_int"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogNamedInt)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogNamedString represents a log_named_string event raised by the CTHelpersTest contract.
type CTHelpersTestLogNamedString struct {
	Key string
	Val string
	Raw *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogNamedStringEventName = "log_named_string"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogNamedString) ContractEventName() string {
	return CTHelpersTestLogNamedStringEventName
}

// UnpackLogNamedStringEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_string(string key, string val)
func (cTHelpersTest *CTHelpersTest) UnpackLogNamedStringEvent(log *types.Log) (*CTHelpersTestLogNamedString, error) {
	event := "log_named_string"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogNamedString)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogNamedUint represents a log_named_uint event raised by the CTHelpersTest contract.
type CTHelpersTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogNamedUintEventName = "log_named_uint"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogNamedUint) ContractEventName() string {
	return CTHelpersTestLogNamedUintEventName
}

// UnpackLogNamedUintEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (cTHelpersTest *CTHelpersTest) UnpackLogNamedUintEvent(log *types.Log) (*CTHelpersTestLogNamedUint, error) {
	event := "log_named_uint"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogNamedUint)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogString represents a log_string event raised by the CTHelpersTest contract.
type CTHelpersTestLogString struct {
	Arg0 string
	Raw  *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogStringEventName = "log_string"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogString) ContractEventName() string {
	return CTHelpersTestLogStringEventName
}

// UnpackLogStringEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_string(string arg0)
func (cTHelpersTest *CTHelpersTest) UnpackLogStringEvent(log *types.Log) (*CTHelpersTestLogString, error) {
	event := "log_string"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogString)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogUint represents a log_uint event raised by the CTHelpersTest contract.
type CTHelpersTestLogUint struct {
	Arg0 *big.Int
	Raw  *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogUintEventName = "log_uint"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogUint) ContractEventName() string {
	return CTHelpersTestLogUintEventName
}

// UnpackLogUintEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event log_uint(uint256 arg0)
func (cTHelpersTest *CTHelpersTest) UnpackLogUintEvent(log *types.Log) (*CTHelpersTestLogUint, error) {
	event := "log_uint"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogUint)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

// CTHelpersTestLogs represents a logs event raised by the CTHelpersTest contract.
type CTHelpersTestLogs struct {
	Arg0 []byte
	Raw  *types.Log // Blockchain specific contextual infos
}

const CTHelpersTestLogsEventName = "logs"

// ContractEventName returns the user-defined event name.
func (CTHelpersTestLogs) ContractEventName() string {
	return CTHelpersTestLogsEventName
}

// UnpackLogsEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event logs(bytes arg0)
func (cTHelpersTest *CTHelpersTest) UnpackLogsEvent(log *types.Log) (*CTHelpersTestLogs, error) {
	event := "logs"
	if len(log.Topics) == 0 || log.Topics[0] != cTHelpersTest.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CTHelpersTestLogs)
	if len(log.Data) > 0 {
		if err := cTHelpersTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cTHelpersTest.abi.Events[event].Inputs {
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

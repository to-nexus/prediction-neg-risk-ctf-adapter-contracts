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

// Order is an auto generated low-level Go binding around an user-defined struct.

// NegRiskFeeModuleMetaData contains all meta data concerning the NegRiskFeeModule contract.
var NegRiskFeeModuleMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_negRiskCtfExchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_negRiskAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ctf\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ctf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"contractIExchange\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"takerOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"makerOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"takerFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"makerFillAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeRate\",\"type\":\"uint256\"}],\"name\":\"matchOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"}]",
	ID:  "NegRiskFeeModule",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162001a3b38038062001a3b83398101604081905262000034916200022b565b3360009081526020818152604091829020600190556001600160a01b03851660808190528251635c1548fb60e01b8152925186939192635c1548fb92600480820193918290030181865afa15801562000091573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b7919062000275565b6001600160a01b031660a0816001600160a01b0316815250506080516001600160a01b0316633b521d786040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000111573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000137919062000275565b6001600160a01b0390811660c05260405163a22cb46560e01b8152848216600482015260016024820152908316915063a22cb46590604401600060405180830381600087803b1580156200018a57600080fd5b505af11580156200019f573d6000803e3d6000fd5b505060405163a22cb46560e01b8152306004820152600160248201526001600160a01b038416925063a22cb4659150604401600060405180830381600087803b158015620001ec57600080fd5b505af115801562000201573d6000803e3d6000fd5b505050505050506200029a565b80516001600160a01b03811681146200022657600080fd5b919050565b6000806000606084860312156200024157600080fd5b6200024c846200020e565b92506200025c602085016200020e565b91506200026c604085016200020e565b90509250925092565b6000602082840312156200028857600080fd5b62000293826200020e565b9392505050565b60805160a05160c05161174f620002ec6000396000818160f3015281816103b2015261092701526000818161028f015281816103d801526109010152600081816102680152610612015261174f6000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c80638bad0c0a11610081578063d2f7265a1161005b578063d2f7265a14610263578063d8dfeb451461028a578063f23a6e61146102b157600080fd5b80638bad0c0a146101dc578063bc197c81146101e4578063d2539b371461025057600080fd5b8063425c2096116100b2578063425c209614610188578063429b62e51461019b57806370480275146101c957600080fd5b80631785f53c146100d957806322a9339f146100ee57806324d7806c1461013f575b600080fd5b6100ec6100e7366004610db4565b6102c4565b005b6101157f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61017861014d366004610db4565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205460011490565b6040519015158152602001610136565b6100ec610196366004610dcf565b61035f565b6101bb6101a9366004610db4565b60006020819052908152604090205481565b604051908152602001610136565b6100ec6101d7366004610db4565b61046c565b6100ec610508565b61021f6101f2366004610e90565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091168152602001610136565b6100ec61025e36600461121a565b61058c565b6101157f000000000000000000000000000000000000000000000000000000000000000081565b6101157f000000000000000000000000000000000000000000000000000000000000000081565b61021f6102bf366004611323565b610691565b3360009081526020819052604090205460011461030d576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166000818152602081905260408082208290555133917f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e91a350565b336000908152602081905260409020546001146103a8576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082156103d6577f00000000000000000000000000000000000000000000000000000000000000006103f8565b7f00000000000000000000000000000000000000000000000000000000000000005b905061040781308686866106bd565b6040805173ffffffffffffffffffffffffffffffffffffffff808416825286166020820152908101849052606081018390527f6ce49f8691a80db5eb4f60cd55b14640529346a7ddf9bf8f77a423fa6a10bfdb9060800160405180910390a150505050565b336000908152602081905260409020546001146104b5576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116600081815260208190526040808220600190555133917ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc91a350565b33600090815260208190526040902054600114610551576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336000818152602081905260408082208290555182917f787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e91a3565b336000908152602081905260409020546001146105d5576040517f7bfa4b9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fe60f0c0500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e60f0c059061064d90889088908890889060040161155c565b600060405180830381600087803b15801561066757600080fd5b505af115801561067b573d6000803e3d6000fd5b5050505061068a848383610711565b5050505050565b7ff23a6e61000000000000000000000000000000000000000000000000000000005b9695505050505050565b801561068a57816000036107045773ffffffffffffffffffffffffffffffffffffffff841630146106f9576106f485858584610789565b61068a565b6106f485848361079b565b6106f485858585856107ab565b825160005b8181101561068a57828582815181106107315761073161162c565b602002602001015161012001511115610781576107818582815181106107595761075961162c565b60200260200101518583815181106107735761077361162c565b602002602001015185610857565b600101610716565b610795848484846109d1565b50505050565b6107a6838383610ac1565b505050565b6040517ff242432a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301528481166024830152604482018490526064820183905260a06084830152600060a483015286169063f242432a9060c401600060405180830381600087803b15801561083857600080fd5b505af115801561084c573d6000803e3d6000fd5b505050505050505050565b6101208301516000906108b0908383876101400151600181111561087d5761087d61139b565b14610888578561089b565b61089b868860a001518960c00151610b90565b8760a001518860c00151896101400151610bc1565b905060008085610140015160018111156108cc576108cc61139b565b146108d85760006108de565b84608001515b905060008086610140015160018111156108fa576108fa61139b565b14610925577f0000000000000000000000000000000000000000000000000000000000000000610947565b7f00000000000000000000000000000000000000000000000000000000000000005b905082156109c9576109608130886020015185876106bd565b6020808701516040805173ffffffffffffffffffffffffffffffffffffffff808616825290921692820192909252908101839052606081018490527f18fe0464eb77016dc4e227eb0d690e4002756d82b442143bbfb874548952b5f29060800160405180910390a15b505050505050565b60006040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff841660248201528260448201526020600060648360008a5af13d15601f3d116001600051141617169150508061068a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5452414e534645525f46524f4d5f4641494c454400000000000000000000000060448201526064015b60405180910390fd5b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152826024820152602060006044836000895af13d15601f3d1160016000511416171691505080610795576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f5452414e534645525f4641494c454400000000000000000000000000000000006044820152606401610ab8565b600082600003610ba257506000610bba565b82610bad838661168a565b610bb791906116c7565b90505b9392505050565b6000858711610bd2575060006106b3565b6000610be18887878787610c15565b905086600003610bf25790506106b3565b610bff8787878787610c15565b610c099082611702565b98975050505050505050565b60008515610cfd576000610c2a858585610d06565b9050600081118015610c445750670de0b6b3a76400008111155b15610cfb576000836001811115610c5d57610c5d61139b565b03610caf57610c6e6127108261168a565b86610c8a83610c8581670de0b6b3a7640000611702565b610d75565b610c94908a61168a565b610c9e919061168a565b610ca891906116c7565b9150610cfb565b610cc3670de0b6b3a764000061271061168a565b86610cda83610c8581670de0b6b3a7640000611702565b610ce4908a61168a565b610cee919061168a565b610cf891906116c7565b91505b505b95945050505050565b600080826001811115610d1b57610d1b61139b565b03610d535782600003610d2f576000610d4c565b82610d42670de0b6b3a76400008661168a565b610d4c91906116c7565b9050610bba565b83600003610d62576000610bb7565b83610bad670de0b6b3a76400008561168a565b6000818310610d845781610bba565b5090919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610daf57600080fd5b919050565b600060208284031215610dc657600080fd5b610bba82610d8b565b600080600060608486031215610de457600080fd5b610ded84610d8b565b95602085013595506040909401359392505050565b60008083601f840112610e1457600080fd5b50813567ffffffffffffffff811115610e2c57600080fd5b6020830191508360208260051b8501011115610e4757600080fd5b9250929050565b60008083601f840112610e6057600080fd5b50813567ffffffffffffffff811115610e7857600080fd5b602083019150836020828501011115610e4757600080fd5b60008060008060008060008060a0898b031215610eac57600080fd5b610eb589610d8b565b9750610ec360208a01610d8b565b9650604089013567ffffffffffffffff80821115610ee057600080fd5b610eec8c838d01610e02565b909850965060608b0135915080821115610f0557600080fd5b610f118c838d01610e02565b909650945060808b0135915080821115610f2a57600080fd5b50610f378b828c01610e4e565b999c989b5096995094979396929594505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516101a0810167ffffffffffffffff81118282101715610f9e57610f9e610f4b565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610feb57610feb610f4b565b604052919050565b803560028110610daf57600080fd5b803560038110610daf57600080fd5b600082601f83011261102257600080fd5b813567ffffffffffffffff81111561103c5761103c610f4b565b61106d60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610fa4565b81815284602083860101111561108257600080fd5b816020850160208301376000918101602001919091529392505050565b60006101a082840312156110b257600080fd5b6110ba610f7a565b9050813581526110cc60208301610d8b565b60208201526110dd60408301610d8b565b60408201526110ee60608301610d8b565b60608201526080820135608082015260a082013560a082015260c082013560c082015260e082013560e0820152610100808301358183015250610120808301358183015250610140611141818401610ff3565b90820152610160611153838201611002565b908201526101808281013567ffffffffffffffff81111561117357600080fd5b61117f85828601611011565b82840152505092915050565b600067ffffffffffffffff8211156111a5576111a5610f4b565b5060051b60200190565b600082601f8301126111c057600080fd5b813560206111d56111d08361118b565b610fa4565b82815260059290921b840181019181810190868411156111f457600080fd5b8286015b8481101561120f57803583529183019183016111f8565b509695505050505050565b600080600080600060a0868803121561123257600080fd5b853567ffffffffffffffff8082111561124a57600080fd5b61125689838a0161109f565b965060209150818801358181111561126d57600080fd5b8801601f81018a1361127e57600080fd5b803561128c6111d08261118b565b81815260059190911b8201840190848101908c8311156112ab57600080fd5b8584015b838110156112e3578035868111156112c75760008081fd5b6112d58f898389010161109f565b8452509186019186016112af565b50985050505060408801359450606088013591508082111561130457600080fd5b50611311888289016111af565b95989497509295608001359392505050565b60008060008060008060a0878903121561133c57600080fd5b61134587610d8b565b955061135360208801610d8b565b94506040870135935060608701359250608087013567ffffffffffffffff81111561137d57600080fd5b61138989828a01610e4e565b979a9699509497509295939492505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600281106113da576113da61139b565b9052565b600381106113da576113da61139b565b6000815180845260005b81811015611414576020818501810151868301820152016113f8565b81811115611426576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60006101a0825184526020830151611489602086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060408301516114b1604086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060608301516114d9606086018273ffffffffffffffffffffffffffffffffffffffff169052565b506080830151608085015260a083015160a085015260c083015160c085015260e083015160e08501526101008084015181860152506101208084015181860152506101408084015161152d828701826113ca565b505061016080840151611542828701826113de565b50506101808084015182828701526106b3838701826113ee565b60808152600061156f6080830187611459565b6020838203818501528187518084528284019150828160051b850101838a0160005b838110156115dd577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030185526115cb838351611459565b94860194925090850190600101611591565b50506040870189905286810360608801528751808252908401945091505081860160005b8281101561161d57815185529383019390830190600101611601565b50929998505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156116c2576116c261165b565b500290565b6000826116fd577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000828210156117145761171461165b565b50039056fea264697066735822122044d970fb603d9f3960e185bea73550a77e66316581897a5f7653dc3d92d181b364736f6c634300080f0033",
}

// NegRiskFeeModule is an auto generated Go binding around an Ethereum contract.
type NegRiskFeeModule struct {
	abi abi.ABI
}

// NewNegRiskFeeModule creates a new instance of NegRiskFeeModule.
func NewNegRiskFeeModule() *NegRiskFeeModule {
	parsed, err := NegRiskFeeModuleMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &NegRiskFeeModule{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *NegRiskFeeModule) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _negRiskCtfExchange, address _negRiskAdapter, address _ctf) returns()
func (negRiskFeeModule *NegRiskFeeModule) PackConstructor(_negRiskCtfExchange common.Address, _negRiskAdapter common.Address, _ctf common.Address) []byte {
	enc, err := negRiskFeeModule.abi.Pack("", _negRiskCtfExchange, _negRiskAdapter, _ctf)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAddAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70480275.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addAdmin(address admin) returns()
func (negRiskFeeModule *NegRiskFeeModule) PackAddAdmin(admin common.Address) []byte {
	enc, err := negRiskFeeModule.abi.Pack("addAdmin", admin)
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
func (negRiskFeeModule *NegRiskFeeModule) TryPackAddAdmin(admin common.Address) ([]byte, error) {
	return negRiskFeeModule.abi.Pack("addAdmin", admin)
}

// PackAdmins is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x429b62e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function admins(address ) view returns(uint256)
func (negRiskFeeModule *NegRiskFeeModule) PackAdmins(arg0 common.Address) []byte {
	enc, err := negRiskFeeModule.abi.Pack("admins", arg0)
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
func (negRiskFeeModule *NegRiskFeeModule) TryPackAdmins(arg0 common.Address) ([]byte, error) {
	return negRiskFeeModule.abi.Pack("admins", arg0)
}

// UnpackAdmins is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (negRiskFeeModule *NegRiskFeeModule) UnpackAdmins(data []byte) (*big.Int, error) {
	out, err := negRiskFeeModule.abi.Unpack("admins", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCollateral is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd8dfeb45.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function collateral() view returns(address)
func (negRiskFeeModule *NegRiskFeeModule) PackCollateral() []byte {
	enc, err := negRiskFeeModule.abi.Pack("collateral")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCollateral is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd8dfeb45.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function collateral() view returns(address)
func (negRiskFeeModule *NegRiskFeeModule) TryPackCollateral() ([]byte, error) {
	return negRiskFeeModule.abi.Pack("collateral")
}

// UnpackCollateral is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (negRiskFeeModule *NegRiskFeeModule) UnpackCollateral(data []byte) (common.Address, error) {
	out, err := negRiskFeeModule.abi.Unpack("collateral", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCtf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22a9339f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ctf() view returns(address)
func (negRiskFeeModule *NegRiskFeeModule) PackCtf() []byte {
	enc, err := negRiskFeeModule.abi.Pack("ctf")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCtf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22a9339f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ctf() view returns(address)
func (negRiskFeeModule *NegRiskFeeModule) TryPackCtf() ([]byte, error) {
	return negRiskFeeModule.abi.Pack("ctf")
}

// UnpackCtf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (negRiskFeeModule *NegRiskFeeModule) UnpackCtf(data []byte) (common.Address, error) {
	out, err := negRiskFeeModule.abi.Unpack("ctf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackExchange is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2f7265a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function exchange() view returns(address)
func (negRiskFeeModule *NegRiskFeeModule) PackExchange() []byte {
	enc, err := negRiskFeeModule.abi.Pack("exchange")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExchange is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2f7265a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function exchange() view returns(address)
func (negRiskFeeModule *NegRiskFeeModule) TryPackExchange() ([]byte, error) {
	return negRiskFeeModule.abi.Pack("exchange")
}

// UnpackExchange is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (negRiskFeeModule *NegRiskFeeModule) UnpackExchange(data []byte) (common.Address, error) {
	out, err := negRiskFeeModule.abi.Unpack("exchange", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackIsAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24d7806c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (negRiskFeeModule *NegRiskFeeModule) PackIsAdmin(addr common.Address) []byte {
	enc, err := negRiskFeeModule.abi.Pack("isAdmin", addr)
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
func (negRiskFeeModule *NegRiskFeeModule) TryPackIsAdmin(addr common.Address) ([]byte, error) {
	return negRiskFeeModule.abi.Pack("isAdmin", addr)
}

// UnpackIsAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (negRiskFeeModule *NegRiskFeeModule) UnpackIsAdmin(data []byte) (bool, error) {
	out, err := negRiskFeeModule.abi.Unpack("isAdmin", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackMatchOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2539b37.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts, uint256 makerFeeRate) returns()
func (negRiskFeeModule *NegRiskFeeModule) PackMatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int, makerFeeRate *big.Int) []byte {
	enc, err := negRiskFeeModule.abi.Pack("matchOrders", takerOrder, makerOrders, takerFillAmount, makerFillAmounts, makerFeeRate)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMatchOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2539b37.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts, uint256 makerFeeRate) returns()
func (negRiskFeeModule *NegRiskFeeModule) TryPackMatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int, makerFeeRate *big.Int) ([]byte, error) {
	return negRiskFeeModule.abi.Pack("matchOrders", takerOrder, makerOrders, takerFillAmount, makerFillAmounts, makerFeeRate)
}

// PackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (negRiskFeeModule *NegRiskFeeModule) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) []byte {
	enc, err := negRiskFeeModule.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
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
func (negRiskFeeModule *NegRiskFeeModule) TryPackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([]byte, error) {
	return negRiskFeeModule.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (negRiskFeeModule *NegRiskFeeModule) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := negRiskFeeModule.abi.Unpack("onERC1155BatchReceived", data)
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
func (negRiskFeeModule *NegRiskFeeModule) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) []byte {
	enc, err := negRiskFeeModule.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
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
func (negRiskFeeModule *NegRiskFeeModule) TryPackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return negRiskFeeModule.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (negRiskFeeModule *NegRiskFeeModule) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := negRiskFeeModule.abi.Unpack("onERC1155Received", data)
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
func (negRiskFeeModule *NegRiskFeeModule) PackRemoveAdmin(admin common.Address) []byte {
	enc, err := negRiskFeeModule.abi.Pack("removeAdmin", admin)
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
func (negRiskFeeModule *NegRiskFeeModule) TryPackRemoveAdmin(admin common.Address) ([]byte, error) {
	return negRiskFeeModule.abi.Pack("removeAdmin", admin)
}

// PackRenounceAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8bad0c0a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceAdmin() returns()
func (negRiskFeeModule *NegRiskFeeModule) PackRenounceAdmin() []byte {
	enc, err := negRiskFeeModule.abi.Pack("renounceAdmin")
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
func (negRiskFeeModule *NegRiskFeeModule) TryPackRenounceAdmin() ([]byte, error) {
	return negRiskFeeModule.abi.Pack("renounceAdmin")
}

// PackWithdrawFees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x425c2096.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function withdrawFees(address to, uint256 id, uint256 amount) returns()
func (negRiskFeeModule *NegRiskFeeModule) PackWithdrawFees(to common.Address, id *big.Int, amount *big.Int) []byte {
	enc, err := negRiskFeeModule.abi.Pack("withdrawFees", to, id, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWithdrawFees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x425c2096.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function withdrawFees(address to, uint256 id, uint256 amount) returns()
func (negRiskFeeModule *NegRiskFeeModule) TryPackWithdrawFees(to common.Address, id *big.Int, amount *big.Int) ([]byte, error) {
	return negRiskFeeModule.abi.Pack("withdrawFees", to, id, amount)
}

// NegRiskFeeModuleFeeRefunded represents a FeeRefunded event raised by the NegRiskFeeModule contract.
type NegRiskFeeModuleFeeRefunded struct {
	Token  common.Address
	To     common.Address
	Id     *big.Int
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const NegRiskFeeModuleFeeRefundedEventName = "FeeRefunded"

// ContractEventName returns the user-defined event name.
func (NegRiskFeeModuleFeeRefunded) ContractEventName() string {
	return NegRiskFeeModuleFeeRefundedEventName
}

// UnpackFeeRefundedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeRefunded(address token, address to, uint256 id, uint256 amount)
func (negRiskFeeModule *NegRiskFeeModule) UnpackFeeRefundedEvent(log *types.Log) (*NegRiskFeeModuleFeeRefunded, error) {
	event := "FeeRefunded"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskFeeModule.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskFeeModuleFeeRefunded)
	if len(log.Data) > 0 {
		if err := negRiskFeeModule.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskFeeModule.abi.Events[event].Inputs {
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

// NegRiskFeeModuleFeeWithdrawn represents a FeeWithdrawn event raised by the NegRiskFeeModule contract.
type NegRiskFeeModuleFeeWithdrawn struct {
	Token  common.Address
	To     common.Address
	Id     *big.Int
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const NegRiskFeeModuleFeeWithdrawnEventName = "FeeWithdrawn"

// ContractEventName returns the user-defined event name.
func (NegRiskFeeModuleFeeWithdrawn) ContractEventName() string {
	return NegRiskFeeModuleFeeWithdrawnEventName
}

// UnpackFeeWithdrawnEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeWithdrawn(address token, address to, uint256 id, uint256 amount)
func (negRiskFeeModule *NegRiskFeeModule) UnpackFeeWithdrawnEvent(log *types.Log) (*NegRiskFeeModuleFeeWithdrawn, error) {
	event := "FeeWithdrawn"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskFeeModule.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskFeeModuleFeeWithdrawn)
	if len(log.Data) > 0 {
		if err := negRiskFeeModule.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskFeeModule.abi.Events[event].Inputs {
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

// NegRiskFeeModuleNewAdmin represents a NewAdmin event raised by the NegRiskFeeModule contract.
type NegRiskFeeModuleNewAdmin struct {
	Admin           common.Address
	NewAdminAddress common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const NegRiskFeeModuleNewAdminEventName = "NewAdmin"

// ContractEventName returns the user-defined event name.
func (NegRiskFeeModuleNewAdmin) ContractEventName() string {
	return NegRiskFeeModuleNewAdminEventName
}

// UnpackNewAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (negRiskFeeModule *NegRiskFeeModule) UnpackNewAdminEvent(log *types.Log) (*NegRiskFeeModuleNewAdmin, error) {
	event := "NewAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskFeeModule.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskFeeModuleNewAdmin)
	if len(log.Data) > 0 {
		if err := negRiskFeeModule.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskFeeModule.abi.Events[event].Inputs {
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

// NegRiskFeeModuleRemovedAdmin represents a RemovedAdmin event raised by the NegRiskFeeModule contract.
type NegRiskFeeModuleRemovedAdmin struct {
	Admin        common.Address
	RemovedAdmin common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const NegRiskFeeModuleRemovedAdminEventName = "RemovedAdmin"

// ContractEventName returns the user-defined event name.
func (NegRiskFeeModuleRemovedAdmin) ContractEventName() string {
	return NegRiskFeeModuleRemovedAdminEventName
}

// UnpackRemovedAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (negRiskFeeModule *NegRiskFeeModule) UnpackRemovedAdminEvent(log *types.Log) (*NegRiskFeeModuleRemovedAdmin, error) {
	event := "RemovedAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != negRiskFeeModule.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NegRiskFeeModuleRemovedAdmin)
	if len(log.Data) > 0 {
		if err := negRiskFeeModule.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range negRiskFeeModule.abi.Events[event].Inputs {
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
func (negRiskFeeModule *NegRiskFeeModule) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], negRiskFeeModule.abi.Errors["NotAdmin"].ID.Bytes()[:4]) {
		return negRiskFeeModule.UnpackNotAdminError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// NegRiskFeeModuleNotAdmin represents a NotAdmin error raised by the NegRiskFeeModule contract.
type NegRiskFeeModuleNotAdmin struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAdmin()
func NegRiskFeeModuleNotAdminErrorID() common.Hash {
	return common.HexToHash("0x7bfa4b9fb0cd3687c1f539f384b3f3f258f2c9aa9186353d0815413b508ed97d")
}

// UnpackNotAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAdmin()
func (negRiskFeeModule *NegRiskFeeModule) UnpackNotAdminError(raw []byte) (*NegRiskFeeModuleNotAdmin, error) {
	out := new(NegRiskFeeModuleNotAdmin)
	if err := negRiskFeeModule.abi.UnpackIntoInterface(out, "NotAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

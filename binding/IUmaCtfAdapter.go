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

// QuestionData is an auto generated low-level Go binding around an user-defined struct.
type QuestionData struct {
	RequestTimestamp             *big.Int
	Reward                       *big.Int
	ProposalBond                 *big.Int
	Liveness                     *big.Int
	EmergencyResolutionTimestamp *big.Int
	Resolved                     bool
	Paused                       bool
	Reset                        bool
	RewardToken                  common.Address
	Creator                      common.Address
	AncillaryData                []byte
}

// IUmaCtfAdapterMetaData contains all meta data concerning the IUmaCtfAdapter contract.
var IUmaCtfAdapterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"flag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"getQuestion\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liveness\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"emergencyResolutionTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resolved\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reset\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ancillaryData\",\"type\":\"bytes\"}],\"internalType\":\"structQuestionData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"ancillaryData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liveness\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"ready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"resolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payouts\",\"type\":\"uint256[]\"}],\"name\":\"QuestionEmergencyResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"QuestionFlagged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ancillaryData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalBond\",\"type\":\"uint256\"}],\"name\":\"QuestionInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"QuestionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"QuestionReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"settledPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payouts\",\"type\":\"uint256[]\"}],\"name\":\"QuestionResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionID\",\"type\":\"bytes32\"}],\"name\":\"QuestionUnpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Flagged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAncillaryData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOOPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPayouts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFlagged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOptimisticOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotReadyToResolve\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Resolved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafetyPeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedToken\",\"type\":\"error\"}]",
	ID:  "IUmaCtfAdapter",
}

// IUmaCtfAdapter is an auto generated Go binding around an Ethereum contract.
type IUmaCtfAdapter struct {
	abi abi.ABI
}

// NewIUmaCtfAdapter creates a new instance of IUmaCtfAdapter.
func NewIUmaCtfAdapter() *IUmaCtfAdapter {
	parsed, err := IUmaCtfAdapterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IUmaCtfAdapter{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IUmaCtfAdapter) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackFlag is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78165a48.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flag(bytes32 questionID) returns()
func (iUmaCtfAdapter *IUmaCtfAdapter) PackFlag(questionID [32]byte) []byte {
	enc, err := iUmaCtfAdapter.abi.Pack("flag", questionID)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFlag is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78165a48.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function flag(bytes32 questionID) returns()
func (iUmaCtfAdapter *IUmaCtfAdapter) TryPackFlag(questionID [32]byte) ([]byte, error) {
	return iUmaCtfAdapter.abi.Pack("flag", questionID)
}

// PackGetQuestion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x58c039cd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getQuestion(bytes32 questionID) returns((uint256,uint256,uint256,uint256,uint256,bool,bool,bool,address,address,bytes))
func (iUmaCtfAdapter *IUmaCtfAdapter) PackGetQuestion(questionID [32]byte) []byte {
	enc, err := iUmaCtfAdapter.abi.Pack("getQuestion", questionID)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetQuestion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x58c039cd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getQuestion(bytes32 questionID) returns((uint256,uint256,uint256,uint256,uint256,bool,bool,bool,address,address,bytes))
func (iUmaCtfAdapter *IUmaCtfAdapter) TryPackGetQuestion(questionID [32]byte) ([]byte, error) {
	return iUmaCtfAdapter.abi.Pack("getQuestion", questionID)
}

// UnpackGetQuestion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x58c039cd.
//
// Solidity: function getQuestion(bytes32 questionID) returns((uint256,uint256,uint256,uint256,uint256,bool,bool,bool,address,address,bytes))
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackGetQuestion(data []byte) (QuestionData, error) {
	out, err := iUmaCtfAdapter.abi.Unpack("getQuestion", data)
	if err != nil {
		return *new(QuestionData), err
	}
	out0 := *abi.ConvertType(out[0], new(QuestionData)).(*QuestionData)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x185d1646.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond, uint256 liveness) returns(bytes32)
func (iUmaCtfAdapter *IUmaCtfAdapter) PackInitialize(ancillaryData []byte, rewardToken common.Address, reward *big.Int, proposalBond *big.Int, liveness *big.Int) []byte {
	enc, err := iUmaCtfAdapter.abi.Pack("initialize", ancillaryData, rewardToken, reward, proposalBond, liveness)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x185d1646.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond, uint256 liveness) returns(bytes32)
func (iUmaCtfAdapter *IUmaCtfAdapter) TryPackInitialize(ancillaryData []byte, rewardToken common.Address, reward *big.Int, proposalBond *big.Int, liveness *big.Int) ([]byte, error) {
	return iUmaCtfAdapter.abi.Pack("initialize", ancillaryData, rewardToken, reward, proposalBond, liveness)
}

// UnpackInitialize is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x185d1646.
//
// Solidity: function initialize(bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond, uint256 liveness) returns(bytes32)
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackInitialize(data []byte) ([32]byte, error) {
	out, err := iUmaCtfAdapter.abi.Unpack("initialize", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed56531a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pause(bytes32 questionID) returns()
func (iUmaCtfAdapter *IUmaCtfAdapter) PackPause(questionID [32]byte) []byte {
	enc, err := iUmaCtfAdapter.abi.Pack("pause", questionID)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed56531a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pause(bytes32 questionID) returns()
func (iUmaCtfAdapter *IUmaCtfAdapter) TryPackPause(questionID [32]byte) ([]byte, error) {
	return iUmaCtfAdapter.abi.Pack("pause", questionID)
}

// PackReady is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfcac49a2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ready(bytes32 questionID) view returns(bool)
func (iUmaCtfAdapter *IUmaCtfAdapter) PackReady(questionID [32]byte) []byte {
	enc, err := iUmaCtfAdapter.abi.Pack("ready", questionID)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReady is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfcac49a2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ready(bytes32 questionID) view returns(bool)
func (iUmaCtfAdapter *IUmaCtfAdapter) TryPackReady(questionID [32]byte) ([]byte, error) {
	return iUmaCtfAdapter.abi.Pack("ready", questionID)
}

// UnpackReady is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfcac49a2.
//
// Solidity: function ready(bytes32 questionID) view returns(bool)
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackReady(data []byte) (bool, error) {
	out, err := iUmaCtfAdapter.abi.Unpack("ready", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackReset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed3c7d40.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reset(bytes32 questionID) returns()
func (iUmaCtfAdapter *IUmaCtfAdapter) PackReset(questionID [32]byte) []byte {
	enc, err := iUmaCtfAdapter.abi.Pack("reset", questionID)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed3c7d40.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reset(bytes32 questionID) returns()
func (iUmaCtfAdapter *IUmaCtfAdapter) TryPackReset(questionID [32]byte) ([]byte, error) {
	return iUmaCtfAdapter.abi.Pack("reset", questionID)
}

// PackResolve is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c23bdf5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resolve(bytes32 questionID) returns()
func (iUmaCtfAdapter *IUmaCtfAdapter) PackResolve(questionID [32]byte) []byte {
	enc, err := iUmaCtfAdapter.abi.Pack("resolve", questionID)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResolve is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c23bdf5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resolve(bytes32 questionID) returns()
func (iUmaCtfAdapter *IUmaCtfAdapter) TryPackResolve(questionID [32]byte) ([]byte, error) {
	return iUmaCtfAdapter.abi.Pack("resolve", questionID)
}

// PackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f4dae9f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unpause(bytes32 questionID) returns()
func (iUmaCtfAdapter *IUmaCtfAdapter) PackUnpause(questionID [32]byte) []byte {
	enc, err := iUmaCtfAdapter.abi.Pack("unpause", questionID)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f4dae9f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unpause(bytes32 questionID) returns()
func (iUmaCtfAdapter *IUmaCtfAdapter) TryPackUnpause(questionID [32]byte) ([]byte, error) {
	return iUmaCtfAdapter.abi.Pack("unpause", questionID)
}

// IUmaCtfAdapterQuestionEmergencyResolved represents a QuestionEmergencyResolved event raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterQuestionEmergencyResolved struct {
	QuestionID [32]byte
	Payouts    []*big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const IUmaCtfAdapterQuestionEmergencyResolvedEventName = "QuestionEmergencyResolved"

// ContractEventName returns the user-defined event name.
func (IUmaCtfAdapterQuestionEmergencyResolved) ContractEventName() string {
	return IUmaCtfAdapterQuestionEmergencyResolvedEventName
}

// UnpackQuestionEmergencyResolvedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionEmergencyResolved(bytes32 indexed questionID, uint256[] payouts)
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackQuestionEmergencyResolvedEvent(log *types.Log) (*IUmaCtfAdapterQuestionEmergencyResolved, error) {
	event := "QuestionEmergencyResolved"
	if len(log.Topics) == 0 || log.Topics[0] != iUmaCtfAdapter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IUmaCtfAdapterQuestionEmergencyResolved)
	if len(log.Data) > 0 {
		if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iUmaCtfAdapter.abi.Events[event].Inputs {
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

// IUmaCtfAdapterQuestionFlagged represents a QuestionFlagged event raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterQuestionFlagged struct {
	QuestionID [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const IUmaCtfAdapterQuestionFlaggedEventName = "QuestionFlagged"

// ContractEventName returns the user-defined event name.
func (IUmaCtfAdapterQuestionFlagged) ContractEventName() string {
	return IUmaCtfAdapterQuestionFlaggedEventName
}

// UnpackQuestionFlaggedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionFlagged(bytes32 indexed questionID)
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackQuestionFlaggedEvent(log *types.Log) (*IUmaCtfAdapterQuestionFlagged, error) {
	event := "QuestionFlagged"
	if len(log.Topics) == 0 || log.Topics[0] != iUmaCtfAdapter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IUmaCtfAdapterQuestionFlagged)
	if len(log.Data) > 0 {
		if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iUmaCtfAdapter.abi.Events[event].Inputs {
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

// IUmaCtfAdapterQuestionInitialized represents a QuestionInitialized event raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterQuestionInitialized struct {
	QuestionID       [32]byte
	RequestTimestamp *big.Int
	Creator          common.Address
	AncillaryData    []byte
	RewardToken      common.Address
	Reward           *big.Int
	ProposalBond     *big.Int
	Raw              *types.Log // Blockchain specific contextual infos
}

const IUmaCtfAdapterQuestionInitializedEventName = "QuestionInitialized"

// ContractEventName returns the user-defined event name.
func (IUmaCtfAdapterQuestionInitialized) ContractEventName() string {
	return IUmaCtfAdapterQuestionInitializedEventName
}

// UnpackQuestionInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionInitialized(bytes32 indexed questionID, uint256 indexed requestTimestamp, address indexed creator, bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond)
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackQuestionInitializedEvent(log *types.Log) (*IUmaCtfAdapterQuestionInitialized, error) {
	event := "QuestionInitialized"
	if len(log.Topics) == 0 || log.Topics[0] != iUmaCtfAdapter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IUmaCtfAdapterQuestionInitialized)
	if len(log.Data) > 0 {
		if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iUmaCtfAdapter.abi.Events[event].Inputs {
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

// IUmaCtfAdapterQuestionPaused represents a QuestionPaused event raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterQuestionPaused struct {
	QuestionID [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const IUmaCtfAdapterQuestionPausedEventName = "QuestionPaused"

// ContractEventName returns the user-defined event name.
func (IUmaCtfAdapterQuestionPaused) ContractEventName() string {
	return IUmaCtfAdapterQuestionPausedEventName
}

// UnpackQuestionPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionPaused(bytes32 indexed questionID)
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackQuestionPausedEvent(log *types.Log) (*IUmaCtfAdapterQuestionPaused, error) {
	event := "QuestionPaused"
	if len(log.Topics) == 0 || log.Topics[0] != iUmaCtfAdapter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IUmaCtfAdapterQuestionPaused)
	if len(log.Data) > 0 {
		if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iUmaCtfAdapter.abi.Events[event].Inputs {
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

// IUmaCtfAdapterQuestionReset represents a QuestionReset event raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterQuestionReset struct {
	QuestionID [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const IUmaCtfAdapterQuestionResetEventName = "QuestionReset"

// ContractEventName returns the user-defined event name.
func (IUmaCtfAdapterQuestionReset) ContractEventName() string {
	return IUmaCtfAdapterQuestionResetEventName
}

// UnpackQuestionResetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionReset(bytes32 indexed questionID)
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackQuestionResetEvent(log *types.Log) (*IUmaCtfAdapterQuestionReset, error) {
	event := "QuestionReset"
	if len(log.Topics) == 0 || log.Topics[0] != iUmaCtfAdapter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IUmaCtfAdapterQuestionReset)
	if len(log.Data) > 0 {
		if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iUmaCtfAdapter.abi.Events[event].Inputs {
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

// IUmaCtfAdapterQuestionResolved represents a QuestionResolved event raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterQuestionResolved struct {
	QuestionID   [32]byte
	SettledPrice *big.Int
	Payouts      []*big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const IUmaCtfAdapterQuestionResolvedEventName = "QuestionResolved"

// ContractEventName returns the user-defined event name.
func (IUmaCtfAdapterQuestionResolved) ContractEventName() string {
	return IUmaCtfAdapterQuestionResolvedEventName
}

// UnpackQuestionResolvedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionResolved(bytes32 indexed questionID, int256 indexed settledPrice, uint256[] payouts)
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackQuestionResolvedEvent(log *types.Log) (*IUmaCtfAdapterQuestionResolved, error) {
	event := "QuestionResolved"
	if len(log.Topics) == 0 || log.Topics[0] != iUmaCtfAdapter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IUmaCtfAdapterQuestionResolved)
	if len(log.Data) > 0 {
		if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iUmaCtfAdapter.abi.Events[event].Inputs {
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

// IUmaCtfAdapterQuestionUnpaused represents a QuestionUnpaused event raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterQuestionUnpaused struct {
	QuestionID [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const IUmaCtfAdapterQuestionUnpausedEventName = "QuestionUnpaused"

// ContractEventName returns the user-defined event name.
func (IUmaCtfAdapterQuestionUnpaused) ContractEventName() string {
	return IUmaCtfAdapterQuestionUnpausedEventName
}

// UnpackQuestionUnpausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuestionUnpaused(bytes32 indexed questionID)
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackQuestionUnpausedEvent(log *types.Log) (*IUmaCtfAdapterQuestionUnpaused, error) {
	event := "QuestionUnpaused"
	if len(log.Topics) == 0 || log.Topics[0] != iUmaCtfAdapter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IUmaCtfAdapterQuestionUnpaused)
	if len(log.Data) > 0 {
		if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iUmaCtfAdapter.abi.Events[event].Inputs {
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
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["Flagged"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackFlaggedError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["Initialized"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["InvalidAncillaryData"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackInvalidAncillaryDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["InvalidOOPrice"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackInvalidOOPriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["InvalidPayouts"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackInvalidPayoutsError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["NotFlagged"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackNotFlaggedError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["NotInitialized"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackNotInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["NotOptimisticOracle"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackNotOptimisticOracleError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["NotReadyToResolve"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackNotReadyToResolveError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["Paused"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["PriceNotAvailable"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackPriceNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["Resolved"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackResolvedError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["SafetyPeriodNotPassed"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackSafetyPeriodNotPassedError(raw[4:])
	}
	if bytes.Equal(raw[:4], iUmaCtfAdapter.abi.Errors["UnsupportedToken"].ID.Bytes()[:4]) {
		return iUmaCtfAdapter.UnpackUnsupportedTokenError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// IUmaCtfAdapterFlagged represents a Flagged error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterFlagged struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Flagged()
func IUmaCtfAdapterFlaggedErrorID() common.Hash {
	return common.HexToHash("0xe8e3a259d3559d69e650ab9e9b3929a4a904d5807cdab761dd8081bc5745dda4")
}

// UnpackFlaggedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Flagged()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackFlaggedError(raw []byte) (*IUmaCtfAdapterFlagged, error) {
	out := new(IUmaCtfAdapterFlagged)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "Flagged", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterInitialized represents a Initialized error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Initialized()
func IUmaCtfAdapterInitializedErrorID() common.Hash {
	return common.HexToHash("0x5daa87a0e9463431830481fd4b6e3403442dfb9a12b9c07597e9f61d50b633c8")
}

// UnpackInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Initialized()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackInitializedError(raw []byte) (*IUmaCtfAdapterInitialized, error) {
	out := new(IUmaCtfAdapterInitialized)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "Initialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterInvalidAncillaryData represents a InvalidAncillaryData error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterInvalidAncillaryData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAncillaryData()
func IUmaCtfAdapterInvalidAncillaryDataErrorID() common.Hash {
	return common.HexToHash("0x9702d51282660b804d82783337777cfddba4253cabfa075a5b13c00f6b4fd920")
}

// UnpackInvalidAncillaryDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAncillaryData()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackInvalidAncillaryDataError(raw []byte) (*IUmaCtfAdapterInvalidAncillaryData, error) {
	out := new(IUmaCtfAdapterInvalidAncillaryData)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "InvalidAncillaryData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterInvalidOOPrice represents a InvalidOOPrice error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterInvalidOOPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidOOPrice()
func IUmaCtfAdapterInvalidOOPriceErrorID() common.Hash {
	return common.HexToHash("0x86c9649e674da7f8a8052dbe9c93f21fd8d5a257a615ae7f813b042d00545c32")
}

// UnpackInvalidOOPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidOOPrice()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackInvalidOOPriceError(raw []byte) (*IUmaCtfAdapterInvalidOOPrice, error) {
	out := new(IUmaCtfAdapterInvalidOOPrice)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "InvalidOOPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterInvalidPayouts represents a InvalidPayouts error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterInvalidPayouts struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPayouts()
func IUmaCtfAdapterInvalidPayoutsErrorID() common.Hash {
	return common.HexToHash("0x663493a0e55cd55fffa3aeddad0f26801e0b338aadf7eda63e49476e10b78d64")
}

// UnpackInvalidPayoutsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPayouts()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackInvalidPayoutsError(raw []byte) (*IUmaCtfAdapterInvalidPayouts, error) {
	out := new(IUmaCtfAdapterInvalidPayouts)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "InvalidPayouts", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterNotFlagged represents a NotFlagged error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterNotFlagged struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotFlagged()
func IUmaCtfAdapterNotFlaggedErrorID() common.Hash {
	return common.HexToHash("0xbb825d18bed9406d32cdc96a1172cbdda7f27b4225498d9d089479c165d12a98")
}

// UnpackNotFlaggedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotFlagged()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackNotFlaggedError(raw []byte) (*IUmaCtfAdapterNotFlagged, error) {
	out := new(IUmaCtfAdapterNotFlagged)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "NotFlagged", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterNotInitialized represents a NotInitialized error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterNotInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitialized()
func IUmaCtfAdapterNotInitializedErrorID() common.Hash {
	return common.HexToHash("0x87138d5c8c2e77cb9f25c07b03277aad63d22f6a05255580ec55d2c21666e734")
}

// UnpackNotInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitialized()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackNotInitializedError(raw []byte) (*IUmaCtfAdapterNotInitialized, error) {
	out := new(IUmaCtfAdapterNotInitialized)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "NotInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterNotOptimisticOracle represents a NotOptimisticOracle error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterNotOptimisticOracle struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOptimisticOracle()
func IUmaCtfAdapterNotOptimisticOracleErrorID() common.Hash {
	return common.HexToHash("0x05cef855ee9fde01aed6f4c15afcd5f818e21d8402641ca3bfc8709730739ced")
}

// UnpackNotOptimisticOracleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOptimisticOracle()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackNotOptimisticOracleError(raw []byte) (*IUmaCtfAdapterNotOptimisticOracle, error) {
	out := new(IUmaCtfAdapterNotOptimisticOracle)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "NotOptimisticOracle", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterNotReadyToResolve represents a NotReadyToResolve error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterNotReadyToResolve struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotReadyToResolve()
func IUmaCtfAdapterNotReadyToResolveErrorID() common.Hash {
	return common.HexToHash("0xb488fe4b570637504aa639ca85af31cc6cc84f1045b9f56c3637b28ff2e32601")
}

// UnpackNotReadyToResolveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotReadyToResolve()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackNotReadyToResolveError(raw []byte) (*IUmaCtfAdapterNotReadyToResolve, error) {
	out := new(IUmaCtfAdapterNotReadyToResolve)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "NotReadyToResolve", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterPaused represents a Paused error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Paused()
func IUmaCtfAdapterPausedErrorID() common.Hash {
	return common.HexToHash("0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752")
}

// UnpackPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Paused()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackPausedError(raw []byte) (*IUmaCtfAdapterPaused, error) {
	out := new(IUmaCtfAdapterPaused)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "Paused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterPriceNotAvailable represents a PriceNotAvailable error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterPriceNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PriceNotAvailable()
func IUmaCtfAdapterPriceNotAvailableErrorID() common.Hash {
	return common.HexToHash("0x579a4801aa6d7741d62846f455e170acc36453b8e699ca8c0134cf031a4e5fdb")
}

// UnpackPriceNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PriceNotAvailable()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackPriceNotAvailableError(raw []byte) (*IUmaCtfAdapterPriceNotAvailable, error) {
	out := new(IUmaCtfAdapterPriceNotAvailable)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "PriceNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterResolved represents a Resolved error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterResolved struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Resolved()
func IUmaCtfAdapterResolvedErrorID() common.Hash {
	return common.HexToHash("0xea00f1a0c354e8c71923a9a994e40cccad0bc6a1d8dfee4d081be694645d1610")
}

// UnpackResolvedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Resolved()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackResolvedError(raw []byte) (*IUmaCtfAdapterResolved, error) {
	out := new(IUmaCtfAdapterResolved)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "Resolved", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterSafetyPeriodNotPassed represents a SafetyPeriodNotPassed error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterSafetyPeriodNotPassed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafetyPeriodNotPassed()
func IUmaCtfAdapterSafetyPeriodNotPassedErrorID() common.Hash {
	return common.HexToHash("0x2a2c257c7f21a3b6b545cd806a2b50a5b9e38ad00993bc994fa1d58d92d38045")
}

// UnpackSafetyPeriodNotPassedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafetyPeriodNotPassed()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackSafetyPeriodNotPassedError(raw []byte) (*IUmaCtfAdapterSafetyPeriodNotPassed, error) {
	out := new(IUmaCtfAdapterSafetyPeriodNotPassed)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "SafetyPeriodNotPassed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IUmaCtfAdapterUnsupportedToken represents a UnsupportedToken error raised by the IUmaCtfAdapter contract.
type IUmaCtfAdapterUnsupportedToken struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedToken()
func IUmaCtfAdapterUnsupportedTokenErrorID() common.Hash {
	return common.HexToHash("0x6a1728823cfcc894fe1dcf37bfe71f201fb66b0b61862091f422023e22ea5ab9")
}

// UnpackUnsupportedTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedToken()
func (iUmaCtfAdapter *IUmaCtfAdapter) UnpackUnsupportedTokenError(raw []byte) (*IUmaCtfAdapterUnsupportedToken, error) {
	out := new(IUmaCtfAdapterUnsupportedToken)
	if err := iUmaCtfAdapter.abi.UnpackIntoInterface(out, "UnsupportedToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}

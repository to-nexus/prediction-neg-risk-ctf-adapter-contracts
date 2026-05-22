// Code generated - DO NOT EDIT.
// This file contains shared type definitions.

package binding

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type StdInvariantFuzzInterface struct {
	Addr      common.Address
	Artifacts []string
}


type StdInvariantFuzzArtifactSelector struct {
	Artifact  string
	Selectors [][4]byte
}


type Order struct {
	Salt          *big.Int
	Maker         common.Address
	Signer        common.Address
	Taker         common.Address
	TokenId       *big.Int
	MakerAmount   *big.Int
	TakerAmount   *big.Int
	Expiration    *big.Int
	Nonce         *big.Int
	FeeRateBps    *big.Int
	Side          uint8
	SignatureType uint8
	Signature     []byte
}


type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}



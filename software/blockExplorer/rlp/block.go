//// Copyright 2018 The plasma-research Authors
//// This file is part of the plasma-research library.
//// It's unstable. Currently is  not production ready. Structure will be changed potentially.
package rlp

import (
	"bytes"

	"github.com/ethereum/go-ethereum/rlp"
)

// Data structure that used for RLP encoding and decoding of Plasma blocks
type Block struct {
	BlockHeader  BlockHeader
	Transactions []Transaction
}

// PreviousBlockHash - 32 bytes
// MerkleRoot - 32 bytes
type BlockHeader struct {
	RSAAccumulator []byte // uint2048
	RSAChainProof  RSAInclusionProof
}

// Signatures may only contain one or two signatures
type Transaction struct {
	Signatures    []Signature
	MaxBlockIndex uint64
	Inputs        []Input
	Outputs       []Output
}

// Actual content of transaction in terms of UTXO model
type UnsignedTransactionContent struct {
	Inputs   []Input
	Outputs  []Output
	Metadata Metadata
}

// Represents transaction input in terms of UTXO model
// Input should refers to output of some previous transaction
// BlockIndex, TxIndex and OutputIndex helps to find out where that input are
//
// Fields:
// Owner - Ethereum address of the owner, 20 bytes
// BlockIndex  - index of the block that contains corresponding output
// TxIndex     - index of the transaction within the block
// OutputIndex - index of the output within transaction
// AssetId     - Id of asset in terms of multi asset Plasma implementation, 20 bytes.
// Amount      - up to 32 bytes
type Input struct {
	Owner      []byte
	BlockIndex uint64
	TxIndex    uint32
	Amount     Range
}

// Represents transaction output in terms of UTXO model
// Owner   - 20 bytes. Ethereum address of the owner
// AssetId - 20 bytes. Id of asset in terms of multi asset Plasma implementation.
// Amount  - up to 32 bytes
type Output struct {
	Owner  []byte
	Amount Range
}

// Signature 65 bytes long ECDSA signature encoded in RSV format
// R(32) bytes S(32) bytes  V(1) byte
type Signature struct {
	R []byte
	S []byte
	V []byte
}

type Metadata struct {
	MaxBlockId uint32
}

func EncodeToRLP(obj interface{}) ([]byte, error) {
	b := new(bytes.Buffer)
	err := rlp.Encode(b, obj)
	return b.Bytes(), err
}

func DecodeBlock(rlpEncodedBlock []byte) (Block, error) {
	var block = Block{}
	err := rlp.Decode(bytes.NewReader(rlpEncodedBlock), &block)
	return block, err
}

type RSAInclusionProof struct {
	A []byte // uint2048
	R []byte // uint256
}

type Range struct {
	Begin []byte // uint256
	End   []byte // uint256
}

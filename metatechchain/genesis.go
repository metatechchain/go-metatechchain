package metatechchain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/metatechchain/meta-base/hash"
	"github.com/metatechchain/meta-base/inter/idx"

	"github.com/metatechchain/go-metatechchain/inter"
	"github.com/metatechchain/go-metatechchain/metatechchain/genesis"
	"github.com/metatechchain/go-metatechchain/metatechchain/genesis/gpos"
)

type Genesis struct {
	Accounts    genesis.Accounts
	Storage     genesis.Storage
	Delegations genesis.Delegations
	Blocks      genesis.Blocks
	RawEvmItems genesis.RawEvmItems
	Validators  gpos.Validators

	FirstEpoch    idx.Epoch
	PrevEpochTime inter.Timestamp
	Time          inter.Timestamp
	ExtraData     []byte

	TotalSupply *big.Int

	DriverOwner common.Address

	Rules Rules

	Hash func() hash.Hash
}

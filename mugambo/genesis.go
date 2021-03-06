package mugambo

import (
	"math/big"

	"github.com/MugamboBC/mugambo-base/hash"
	"github.com/MugamboBC/mugambo-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"

	"github.com/topcoder1208/fantom-fork/inter"
	"github.com/topcoder1208/fantom-fork/mugambo/genesis"
	"github.com/topcoder1208/fantom-fork/mugambo/genesis/gpos"
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

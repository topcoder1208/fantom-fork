package blockproc

import (
	"github.com/MugamboBC/mugambo-base/inter/idx"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/topcoder1208/fantom-fork/evmcore"
	"github.com/topcoder1208/fantom-fork/inter"
	"github.com/topcoder1208/fantom-fork/mugambo"
)

type TxListener interface {
	OnNewLog(*types.Log)
	OnNewReceipt(tx *types.Transaction, r *types.Receipt, originator idx.ValidatorID)
	Finalize() BlockState
	Update(bs BlockState, es EpochState)
}

type TxListenerModule interface {
	Start(block BlockCtx, bs BlockState, es EpochState, statedb *state.StateDB) TxListener
}

type TxTransactor interface {
	PopInternalTxs(block BlockCtx, bs BlockState, es EpochState, sealing bool, statedb *state.StateDB) types.Transactions
}

type SealerProcessor interface {
	EpochSealing() bool
	SealEpoch() (BlockState, EpochState)
	Update(bs BlockState, es EpochState)
}

type SealerModule interface {
	Start(block BlockCtx, bs BlockState, es EpochState) SealerProcessor
}

type ConfirmedEventsProcessor interface {
	ProcessConfirmedEvent(inter.EventI)
	Finalize(block BlockCtx, blockSkipped bool) BlockState
}

type ConfirmedEventsModule interface {
	Start(bs BlockState, es EpochState) ConfirmedEventsProcessor
}

type EVMProcessor interface {
	Execute(txs types.Transactions, internal bool) types.Receipts
	Finalize() (evmBlock *evmcore.EvmBlock, skippedTxs []uint32, receipts types.Receipts)
}

type EVM interface {
	Start(block BlockCtx, statedb *state.StateDB, reader evmcore.DummyChain, onNewLog func(*types.Log), net mugambo.Rules) EVMProcessor
}

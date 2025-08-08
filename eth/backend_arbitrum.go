package eth

import (
	"context"

	"github.com/serenite11/go-ethereum/core"
	"github.com/serenite11/go-ethereum/core/state"
	"github.com/serenite11/go-ethereum/core/types"
	"github.com/serenite11/go-ethereum/core/vm"
	"github.com/serenite11/go-ethereum/eth/tracers"
	"github.com/serenite11/go-ethereum/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*types.Transaction, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return eth.stateAtTransaction(ctx, block, txIndex, reexec)
}

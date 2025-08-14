package eth

import (
	"context"

	"github.com/ethereum/go-arbitrum/core"
	"github.com/ethereum/go-arbitrum/core/state"
	"github.com/ethereum/go-arbitrum/core/types"
	"github.com/ethereum/go-arbitrum/core/vm"
	"github.com/ethereum/go-arbitrum/eth/tracers"
	"github.com/ethereum/go-arbitrum/ethdb"
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

package arbitrum

import (
	"context"

	"github.com/serenite11/go-ethereum/arbitrum_types"
	"github.com/serenite11/go-ethereum/core"
	"github.com/serenite11/go-ethereum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}

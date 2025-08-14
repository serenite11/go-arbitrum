package arbitrum

import (
	"context"

	"github.com/ethereum/go-arbitrum/common/hexutil"
	"github.com/ethereum/go-arbitrum/core"
	"github.com/ethereum/go-arbitrum/internal/ethapi"
	"github.com/ethereum/go-arbitrum/internal/ethapi/override"
	"github.com/ethereum/go-arbitrum/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, overrides *override.StateOverride, blockOverrides *override.BlockOverrides, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, overrides, blockOverrides, gasCap)
}

func NewRevertReason(result *core.ExecutionResult) error {
	return ethapi.NewRevertError(result)
}

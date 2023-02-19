package trie

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
)

type BlockchainData interface {
	GetBlock(ctx context.Context, blockHashStr string) (*types.Block, error)
	GetTxReceiptsOfABlock(ctx context.Context, blockHashStr string) (types.Receipts, error)
}

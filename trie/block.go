package trie

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type blockchainData struct {
	client *ethclient.Client
}

func NewBlockchainData(rawUrl string) BlockchainData {

	theClient, err := ethclient.Dial(rawUrl)

	if err != nil {
		panic(err)
	}

	return &blockchainData{
		client: theClient,
	}
}

func (b *blockchainData) GetBlock(ctx context.Context, blockHashStr string) (*types.Block, error) {
	return b.client.BlockByHash(ctx, common.HexToHash(blockHashStr))
}

func (b *blockchainData) GetTxReceiptsOfABlock(ctx context.Context, blockHashStr string) (types.Receipts, error) {
	theBlock, err := b.GetBlock(ctx, blockHashStr)
	if err != nil {
		return nil, err
	}

	var receipts = make(types.Receipts, len(theBlock.Transactions()), len(theBlock.Transactions()))

	for i, tx := range theBlock.Transactions() {
		fmt.Println("index: ", i)
		//fmt.Println("tx hash: ", tx.Hash().String())

		if i == 119 {
			fmt.Printf("\nthe tx hash of transaction with index %d is %s \n", i, tx.Hash().String())
		}

		if tx.Hash().String() == "0x4c8da45a7a7e7d4a358b5c7650031b55abc0a6038b87de7ee04a1195f21b2f58" {
			fmt.Printf("\ntransaction with id: %s has index of %d\n", tx.Hash().String(), i)
		}

		theReceipt, err := b.client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, err
		}

		receipts[i] = theReceipt

		//b, err := json.Marshal(theReceipt)

		//fmt.Println("theReceipt:", string(b))
	}

	return receipts, nil
}

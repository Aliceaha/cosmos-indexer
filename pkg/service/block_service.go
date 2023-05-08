package service

import (
	"fmt"

	"github.com/Biskwit/cosmos-indexer/pkg/db"
	"github.com/Biskwit/cosmos-indexer/pkg/models"
)

func CreateBlock(block models.BlockResponse, blockchain string) {
	rawQuery := fmt.Sprintf(`
		CREATE block:%s CONTENT {
			hash: "%s",
			height: "%s",
			time: "%s",
			num_txs: %d,
			chain_id: "%s"
		};
		`, block.Block.Header.Height, block.BlockID.Hash, block.Block.Header.Height, block.Block.Header.Time, len(block.Block.Data.Txs), block.Block.Header.ChainID)
	db.Request(blockchain, rawQuery)
}

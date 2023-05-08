package modules

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/Biskwit/cosmos-indexer/pkg/models"
	"github.com/Biskwit/cosmos-indexer/pkg/service"
	"github.com/Biskwit/cosmos-indexer/pkg/utils"
)

func Realtime(blockchain string) {
	actualBlock := ""
	for {
		resp, err := http.Get(os.Getenv("COSMOS_API") + "/blocks/latest")
		if err != nil {
			utils.Logger.Error("Error:", err)
			continue
		}
		var result models.BlockResponse
		_ = json.NewDecoder(resp.Body).Decode(&result)
		block := result.BlockID.Hash
		if block != actualBlock {
			utils.Logger.Info("new block: ", block)
			actualBlock = result.BlockID.Hash
			go service.CreateBlock(result, blockchain)

			for _, tx := range result.Block.Data.Txs {
				decodedTx := utils.DecodeTx(tx)
				utils.Logger.Info("	└ tx: ", decodedTx)
				go service.CreateTx(decodedTx, blockchain)
			}
		}
		time.Sleep(4 * time.Second)
	}

}

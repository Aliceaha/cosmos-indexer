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
		block := result.Block.Header.Height
		if block != actualBlock {
			utils.Logger.Info("new block: ", result.Block.Header.Height)
			actualBlock = block
			go service.CreateBlock(result, blockchain)
			go service.CreateTx(result.Block.Data.Txs, blockchain)
		}
		time.Sleep(2 * time.Second)
	}

}

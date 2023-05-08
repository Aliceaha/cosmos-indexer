package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Biskwit/cosmos-indexer/pkg/db"
	"github.com/Biskwit/cosmos-indexer/pkg/models"
	"github.com/Biskwit/cosmos-indexer/pkg/utils"
)

func CreateTx(decodedTx string, blockchain string) {
	respTx, errTx := http.Get(os.Getenv("COSMOS_API") + "/cosmos/tx/v1beta1/txs/" + decodedTx)
	if errTx != nil {
		utils.Logger.Error("Error:", errTx)
		return
	}

	var tx models.Tx
	_ = json.NewDecoder(respTx.Body).Decode(&tx)
	rawQuery := fmt.Sprintf(`
		CREATE tx:%s CONTENT {
			hash: "%s",
			height: "%s",
			logs: %+v,
			msgs: "%s",
			signatures: "%s",
			data: "%s",
			gas_used: "%s",
			gas_wanted: "%s",
			code: %d,
			timestamp: "%s",
			memo: "%s"
		};
	`, tx.TxResponse.TxHash, tx.TxResponse.TxHash, tx.TxResponse.Height, tx.TxResponse.Logs, tx.Tx.Body.Messages, tx.Tx.Signatures, tx.TxResponse.Data, tx.TxResponse.GasUsed, tx.TxResponse.GasWanted, tx.TxResponse.Code, tx.TxResponse.Timestamp, tx.Tx.Body.Memo)
	utils.Logger.Info(rawQuery)
	db.Request(blockchain, rawQuery)
}

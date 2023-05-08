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

func CreateTx(txs []string, blockchain string) {
	for _, txString := range txs {
		go func(txString string) {
			decodedTx := utils.DecodeTx(txString)
			utils.Logger.Info("	└ tx: ", decodedTx)
			respTx, errTx := http.Get(os.Getenv("COSMOS_API") + "/cosmos/tx/v1beta1/txs/" + decodedTx)
			if errTx != nil {
				utils.Logger.Error("Error:", errTx)
				return
			}

			var tx models.Tx
			_ = json.NewDecoder(respTx.Body).Decode(&tx)

			logsByte, _ := json.Marshal(tx.TxResponse.Logs)
			signaturesByte, _ := json.Marshal(tx.Tx.Signatures)
			msgByte, _ := json.Marshal(tx.Tx.Body.Messages)
			rawQuery := fmt.Sprintf(`
			CREATE tx:%s CONTENT {
				height: "%s",
				logs: %+v,
				msgs: %+v,
				signatures: %+v,
				data: "%s",
				gas_used: "%s",
				gas_wanted: "%s",
				code: %d,
				timestamp: "%s",
				memo: "%s"
			};
		`, tx.TxResponse.TxHash, tx.TxResponse.Height, string(logsByte), string(msgByte), string(signaturesByte), tx.TxResponse.Data, tx.TxResponse.GasUsed, tx.TxResponse.GasWanted, tx.TxResponse.Code, tx.TxResponse.Timestamp, tx.Tx.Body.Memo)
			db.Request(blockchain, rawQuery)
		}(txString)
	}

}

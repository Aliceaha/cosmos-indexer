package main

import (
	"os"
	"time"

	"github.com/Biskwit/cosmos-indexer/pkg/modules"
	"github.com/Biskwit/cosmos-indexer/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	// go build indexer.go
	// ./indexer realtime cosmos_mainnet
	err := godotenv.Load(".env")
	if err != nil {
		utils.Logger.Fatal("Error loading environment variables file")
	}
	utils.InitLogger()
	args := os.Args[1:]
	if len(args) <= 1 {
		panic("no args provided")
	}
	mode := args[0]
	blockchain := args[1]

	utils.Logger.Info("starting cosmos indexer ", blockchain)

	switch mode {
	case "realtime":
		utils.Logger.Info("mode: realtime")
		modules.Realtime(blockchain)
	case "ingest":
		utils.Logger.Info("mode: ingest") //TODO: ingest from genesis
		modules.Ingest(blockchain)
	case "full":
		utils.Logger.Info("mode: full") //TODO: ingest from genesis
		modules.Realtime(blockchain)
		time.Sleep(4 * time.Second) // wait for realtime to start properly
		modules.Ingest(blockchain)
	default:
		panic("no args provided")
	}
}

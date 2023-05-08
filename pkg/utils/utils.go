package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	log "github.com/sirupsen/logrus"
)

var (
	// Info is the logger for info messages
	Logger = log.New()
)

func InitLogger() {
	// Set the default log level to info
	Logger.SetLevel(log.InfoLevel)
}

func DecodeTx(tx string) string {
	txBz, err := base64.StdEncoding.DecodeString(tx)
	if err != nil {
		Logger.Error("Error:", err)
	}
	converted := []byte(txBz)

	// hash the byte slice and return the resulting string
	hasher := sha256.New()
	hasher.Write(converted)

	return hex.EncodeToString(hasher.Sum(nil))
}

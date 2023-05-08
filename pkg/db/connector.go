package db

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"os"

	"github.com/Biskwit/cosmos-indexer/pkg/utils"
)

func Request(blockchain string, statement string) string {
	utils.Logger.Debug(statement)
	// Set the endpoint URL
	url := os.Getenv("SURREALDB") + "/sql"

	// Set the request body
	body := []byte(statement)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		utils.Logger.Error("Error creating request:", err)
		return ""
	}

	// Set the headers
	req.Header.Set("NS", "cosmos")
	req.Header.Set("DB", blockchain)
	req.Header.Set("Accept", "application/json")

	// Set the basic auth credentials
	username := "root"
	password := os.Getenv("SURREALDB_PWD")
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", basicAuth)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.Logger.Error("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	// Print the response status code and body
	utils.Logger.Debug("Response Status:", resp.Status)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	utils.Logger.Debug("Response Body:", buf.String())
	return buf.String()
}

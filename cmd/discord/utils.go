package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"web3.warehouse/mefpmonitor/types"
)

func SendMessage(url string, message types.Message) {
	payload := new(bytes.Buffer)

	err := json.NewEncoder(payload).Encode(message)
	if err != nil {
		fmt.Printf("Encoder error: %s\n", err)
		LogError("Encoder error: " + err.Error())
	}

	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		fmt.Printf("Webhook req error: %s\n", err)
		LogError("Webhook req error: " + err.Error())
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("decoding res error: %s\n", err)
			LogError("decoding res error: " + err.Error())
		}

		fmt.Print(string(responseBody))
	}
}

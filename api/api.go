package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	discord "web3.warehouse/mefpmonitor/cmd/discord"
	"web3.warehouse/mefpmonitor/types"
	"web3.warehouse/mefpmonitor/utils"
)

func GetCollectionStats(collectionName string) types.CollectionData {
	var collection types.CollectionData
	endpoint := utils.GetConfigFromJson().MagicEdenAPIEndpoint

	resp, apiErr := http.Get(endpoint + "/collections/" + collectionName + "/stats")
	if apiErr != nil {
		fmt.Printf("me api error: %s\n", apiErr)
		discord.LogError("me api error: " + apiErr.Error())
	}

	respBody, decErr := io.ReadAll(resp.Body)

	if decErr != nil {
		fmt.Printf("Decoding response error: %s\n", decErr)
		discord.LogError("Decoding response error: " + decErr.Error())
	}

	unmarshalErr := json.Unmarshal(respBody, &collection)

	if unmarshalErr != nil {
		fmt.Printf("Converting json to obj err: %s\n", unmarshalErr)
		discord.LogError("Converting json to obj err: " + unmarshalErr.Error())
	}

	if collection.Symbol != "" {
		fmt.Printf("Collection Name: %s\n", collection.Symbol)
		fmt.Printf("Collection FP: %d\n", collection.FloorPrice)
		fmt.Printf("Collection Avg24h: %f\n", 0.0)
		fmt.Printf("Collection ListedCount: %d\n", collection.ListedCount)
		fmt.Printf("Collection Volume: %d\n", 0)
	} else {
		fmt.Printf("GetCollectionStats error")
		discord.LogError("GetCollectionStats error")
	}

	fmt.Printf("\n")

	return collection
}

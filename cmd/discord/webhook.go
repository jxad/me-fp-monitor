package cmd

import (
	"strconv"

	"web3.warehouse/mefpmonitor/types"
	"web3.warehouse/mefpmonitor/utils"
)

func SendPriceAlert(collectionData types.CollectionData, priceMonitored string, alertTitle string) {
	username := "FP Monitor"
	url := utils.GetConfigFromJson().PriceAlertWebhook

	fp := strconv.Itoa(collectionData.FloorPrice)
	listedCount := strconv.Itoa(collectionData.ListedCount)

	embedFields := []types.Field{
		{
			Name:   "Price Monitored",
			Value:  priceMonitored,
			Inline: false,
		},

		{
			Name:   "Floor Price",
			Value:  utils.ConvertLamportsToSol(fp),
			Inline: false,
		},

		{
			Name:   "Listed Count",
			Value:  listedCount,
			Inline: false,
		},
	}

	title := collectionData.Symbol + " " + alertTitle

	embeds := []types.Embed{
		{
			Title:  &title,
			Fields: &embedFields,
		},
	}

	message := types.Message{
		Username: &username,
		Embeds:   &embeds,
	}

	SendMessage(url, message)
}

func LogError(err string) {
	url := utils.GetConfigFromJson().PriceAlertWebhook
	username := "Error Logger"

	embedFields := []types.Field{
		{
			Name:   "Error Log",
			Value:  err,
			Inline: false,
		},
	}

	title := "ME FP Monitor Error"

	embeds := []types.Embed{
		{
			Title:  &title,
			Fields: &embedFields,
		},
	}

	message := types.Message{
		Username: &username,
		Embeds:   &embeds,
	}

	SendMessage(url, message)
}

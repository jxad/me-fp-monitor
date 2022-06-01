package cmd

import (
	"strconv"

	"web3.warehouse/mefpmonitor/types"
	"web3.warehouse/mefpmonitor/utils"
)

func SendPriceAlert(collectionData types.CollectionData, priceMonitored string, alertTitle string) {
	username := "FP Monitor"
	url := utils.GetConfigFromJson().PriceAlertWebhook

	embedFields := []types.Field{
		{
			Name:   "Price Monitored",
			Value:  priceMonitored + " ◎",
			Inline: false,
		},

		{
			Name:   "Floor Price",
			Value:  utils.ConvertLamportsToSol(collectionData.FloorPrice) + " ◎",
			Inline: false,
		},

		{
			Name:   "Listed Count",
			Value:  strconv.Itoa(collectionData.ListedCount),
			Inline: false,
		},

		{
			Name:   "24H AVG Price",
			Value:  utils.ConvertLamportsToSol(collectionData.AvgPrice24Hr) + " ◎",
			Inline: false,
		},

		{
			Name:   "Total Volume",
			Value:  utils.ConvertLamportsToSol(collectionData.VolumeAll) + " ◎",
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
	url := utils.GetConfigFromJson().ErrorWebhook
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

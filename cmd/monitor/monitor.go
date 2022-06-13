package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gosuri/uilive"
	"web3.warehouse/mefpmonitor/api"
	discord "web3.warehouse/mefpmonitor/cmd/discord"
	"web3.warehouse/mefpmonitor/types"
	"web3.warehouse/mefpmonitor/utils"
)

func StartSingleCollectionMonitor(name string, upAlert types.AlertCondition, downAlert types.AlertCondition, delay int) {
	for {
		CheckCollection(name, upAlert, downAlert)
		StartTimer(delay)
	}
}

func StartMultipleCollectionMonitor(delay int) {
	for {
		collections := utils.GetCollectionsFromJson()
		fixedDelay := delay / len(collections)

		for i := 0; i < len(collections); i++ {
			CheckCollection(collections[i].Symbol, collections[i].UpAlert, collections[i].DownAlert)
			StartTimer(fixedDelay)
		}
	}
}

func CheckCollection(name string, upAlert types.AlertCondition, downAlert types.AlertCondition) {
	collectionData := api.GetCollectionStats(name)
	if collectionData.Symbol != "" {
		fp := utils.ConvertLamportsToSol(collectionData.FloorPrice)
		var upAlertPrice float64
		var downAlertPrice float64
		var collectionFP float64

		upAlertPrice, err := strconv.ParseFloat(upAlert.Price, 64)
		if err != nil {
			fmt.Println("CheckCollection error: " + err.Error())
			return
		}

		downAlertPrice, err = strconv.ParseFloat(downAlert.Price, 64)
		if err != nil {
			fmt.Println("CheckCollection error: " + err.Error())
			return
		}

		collectionFP, err = strconv.ParseFloat(fp, 64)
		if err != nil {
			fmt.Println("CheckCollection error: " + err.Error())
			return
		}

		if upAlert.Enabled && (collectionFP > upAlertPrice) {
			discord.SendPriceAlert(collectionData, upAlert.Price, "FP Price is Up!!")
		}

		if downAlert.Enabled && (collectionFP < downAlertPrice) {
			discord.SendPriceAlert(collectionData, downAlert.Price, "FP Price is Down!!")
		}
	}
}

func StartTimer(delay int) {
	writer := uilive.New()
	// start listening for updates and render
	writer.Start()
	for i := delay; i >= 0; i-- {
		fmt.Fprintf(writer, "Next request in: %d seconds\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Fprintln(writer, "Sending new request....")
	fmt.Printf("\n")
	writer.Stop() // flush and stop rendering
}

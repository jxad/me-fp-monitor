package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gosuri/uilive"
	"web3.warehouse/mefpmonitor/api"
	discord "web3.warehouse/mefpmonitor/cmd/discord"
	"web3.warehouse/mefpmonitor/utils"
)

func StartSingleCollectionMonitor(name string, upOrDown string, price string, delay int) {
	for {
		counter := CheckCollection(name, upOrDown, price)

		if counter >= 3 {
			fmt.Printf("3 alerts sent. Stopping monitor...")
			os.Exit(0)
		}

		StartTimer(delay)
	}
}

func StartMultipleCollectionMonitor(delay int) {
	for {
		collections := utils.GetCollectionsFromJson()
		fixedDelay := delay / len(collections)

		for i := 0; i < len(collections); i++ {
			CheckCollection(collections[i].Symbol, collections[i].UpOrDown, collections[i].Price)
			StartTimer(fixedDelay)
		}
	}
}

func CheckCollection(name string, upOrDown string, price string) int {
	var counter int = 0
	collectionData := api.GetCollectionStats(name)
	if collectionData.Symbol != "" {
		fp := utils.ConvertLamportsToSol(collectionData.FloorPrice)
		var userPrice float64
		var collectionPrice float64

		if userPriceFloat, err := strconv.ParseFloat(price, 64); err == nil {
			userPrice = userPriceFloat
		} else {
			log.Fatalf("Error - Invalid FP inserted")
			discord.LogError("Error - Invalid FP inserted")
		}

		if fpPriceFloat, err := strconv.ParseFloat(fp, 64); err == nil {
			collectionPrice = fpPriceFloat
		} else {
			log.Fatalf("Error - Invalid Collection FP (memonitor.go => row 31)")
			discord.LogError("Error - Invalid Collection FP (memonitor.go => row 31)")
		}

		switch upOrDown {
		case "UP":
			if collectionPrice > userPrice {
				counter += 1
				discord.SendPriceAlert(collectionData, price, "FP Price is Up!!")
			}
		case "DOWN":
			if userPrice > collectionPrice {
				counter += 1
				discord.SendPriceAlert(collectionData, price, "FP Price is Down!!")
			}
		}
	}

	return counter
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

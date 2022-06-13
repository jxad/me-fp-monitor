package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"web3.warehouse/mefpmonitor/types"
	"web3.warehouse/mefpmonitor/utils"
)

func AddCollection() {
	var symbol string
	var up bool
	var down bool
	var upAlertPrice string
	var downAlertPrice string

	symbol = utils.InputString("Name")

	up = utils.PriceAlertConditionMenu("Up")
	if up {
		upAlertPrice = utils.InputString("Up Alert Price")
	}

	down = utils.PriceAlertConditionMenu("Down")
	if down {
		downAlertPrice = utils.InputString("Down Alert Price")
	}

	collection := types.Collection{
		Symbol: symbol,
		UpAlert: types.AlertCondition{
			Enabled: up,
			Price:   upAlertPrice,
		},
		DownAlert: types.AlertCondition{
			Enabled: up,
			Price:   downAlertPrice,
		},
	}

	collections := utils.GetCollectionsFromJson()

	collections = append(collections, collection)

	file, _ := json.MarshalIndent(collections, "", " ")

	_ = ioutil.WriteFile(utils.COLLECTIONS_JSON_FILE, file, 0644)
}

func RemoveCollection() {
	ViewCollections()

	collections := utils.GetCollectionsFromJson()

	index := utils.InputInteger("Collection index")
	index -= 1

	if index >= 0 && len(collections) >= index+1 {
		name := collections[index].Symbol
		copy(collections[index:], collections[index+1:])
		collections[len(collections)-1] = types.Collection{}
		collections = collections[:len(collections)-1]

		file, _ := json.MarshalIndent(collections, "", " ")

		_ = ioutil.WriteFile(utils.COLLECTIONS_JSON_FILE, file, 0644)

		fmt.Printf("%s removed\n", name)
	} else {
		fmt.Printf("Collection not found\n")
	}
}

func ViewCollections() {
	collections := utils.GetCollectionsFromJson()

	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Symbol", "Up Alert", "Down Alert"})

	counter := 0
	for _, collection := range collections {
		counter += 1
		var up string = "Disabled"
		var down string = "Disabled"

		if collection.UpAlert.Enabled {
			up = collection.UpAlert.Price
		}

		if collection.UpAlert.Enabled {
			down = collection.DownAlert.Price
		}

		t.AppendRow(table.Row{
			counter, collection.Symbol, up, down},
		)
	}

	t.Render()
}

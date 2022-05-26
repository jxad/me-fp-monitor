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
	var price string
	var upordown string

	symbol = utils.InputString("Name")
	price = utils.InputString("Price")
	upordown = utils.PriceAlertConditionMenu()

	collection := types.Collection{
		Symbol:   symbol,
		Price:    price,
		UpOrDown: upordown,
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
	t.AppendHeader(table.Row{"Id", "Symbol", "Price", "Up Or Down"})

	counter := 0
	for _, collection := range collections {
		counter += 1
		t.AppendRow(table.Row{
			counter, collection.Symbol, collection.Price, collection.UpOrDown},
		)
	}

	t.Render()
}

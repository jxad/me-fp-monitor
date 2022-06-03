package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/common-nighthawk/go-figure"
	menu "web3.warehouse/mefpmonitor/cmd/menu"
	"web3.warehouse/mefpmonitor/types"
	"web3.warehouse/mefpmonitor/utils"
)

func main() {
	CheckJsonFiles()

	myFigure := figure.NewColorFigure("IO.wrhs", "", "purple", true)
	myFigure.Print()
	fmt.Print("                                             @jxadd\n")
	menu.StartMenu()
}

func CheckJsonFiles() {
	if _, err := os.Stat(utils.COLLECTIONS_JSON_FILE); err != nil {
		emptyCollection := types.Collection{}
		file, _ := json.MarshalIndent(emptyCollection, "", " ")
		_ = ioutil.WriteFile(utils.COLLECTIONS_JSON_FILE, file, 0644)
	}

	if _, err := os.Stat(utils.CONFIG_JSON_FILE); err != nil {
		emptyConfig := types.Configuration{}
		file, _ := json.MarshalIndent(emptyConfig, "", " ")
		_ = ioutil.WriteFile(utils.CONFIG_JSON_FILE, file, 0644)
	}
}

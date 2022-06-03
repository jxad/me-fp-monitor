package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	monitor "web3.warehouse/mefpmonitor/cmd/monitor"
	"web3.warehouse/mefpmonitor/utils"
)

const DEFAULT_DELAY_SINGLE_COLLECTION_MONITOR = 30
const DEFAULT_DELAY_MULTIPLE_COLLECTION_MONITOR = 600

func StartMenu() {
	var choices = []string{"Start One Collection Monitor", "Multiple Collection Monitor"}

	prompt := promptui.Select{
		Label: "Main Menu",
		Items: choices,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case choices[0]:
		SingleCollectionMonitorMenu()
	case choices[1]:
		MultipleCollectionMonitorMenu()
	}
}

func SingleCollectionMonitorMenu() {
	var collectionName string
	var upordown string
	var priceDiff string
	var delay int

	collectionName = utils.InputString("Collection name")
	upordown = utils.PriceAlertConditionMenu()
	priceDiff = utils.InputString("Price")

	//GetDelay
	delay = utils.InputInteger("Enter delay (seconds)")
	if delay < 0 {
		fmt.Printf("Delay too low, using default delay (%d seconds)\n", DEFAULT_DELAY_SINGLE_COLLECTION_MONITOR)
		delay = DEFAULT_DELAY_SINGLE_COLLECTION_MONITOR
	}

	monitor.StartSingleCollectionMonitor(collectionName, upordown, priceDiff, delay)
}

func MultipleCollectionMonitorMenu() {
	var choices = []string{"Start Monitor", "Settings", "Go Back"}

	prompt := promptui.Select{
		Label: "Multiple Collection Monitor Menu",
		Items: choices,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case choices[0]:
		delay := utils.InputInteger("Enter single request delay (seconds)")
		if delay < 0 {
			fmt.Printf("Delay too low, using default delay (%d seconds)\n", DEFAULT_DELAY_MULTIPLE_COLLECTION_MONITOR)
			monitor.StartMultipleCollectionMonitor(DEFAULT_DELAY_MULTIPLE_COLLECTION_MONITOR)
		} else {
			monitor.StartMultipleCollectionMonitor(delay)
		}
	case choices[1]:
		MultipleCollectionSettingsMenu()
	case choices[2]:
		StartMenu()
	}
}

func MultipleCollectionSettingsMenu() {
	var choices = []string{"Add Collection", "Remove Collection", "View Collections", "Go Back"}

	prompt := promptui.Select{
		Label: "Collections Settings Menu",
		Items: choices,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case choices[0]:
		monitor.AddCollection()
		MultipleCollectionSettingsMenu()
	case choices[1]:
		monitor.RemoveCollection()
		MultipleCollectionSettingsMenu()
	case choices[2]:
		monitor.ViewCollections()
		MultipleCollectionSettingsMenu()
	case choices[3]:
		MultipleCollectionMonitorMenu()
	}
}

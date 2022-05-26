package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"web3.warehouse/mefpmonitor/types"
)

const COLLECTIONS_JSON_FILE = "collections.json"
const CONFIG_JSON_FILE = "config.json"
const SOL_LAMPORTS = 1000000000

func GetCollectionsFromJson() []types.Collection {
	jsonFile, err := os.Open(COLLECTIONS_JSON_FILE)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var collections []types.Collection

	json.Unmarshal([]byte(byteValue), &collections)

	return collections
}

func GetConfigFromJson() types.Configuration {
	jsonFile, err := os.Open(CONFIG_JSON_FILE)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config types.Configuration

	json.Unmarshal([]byte(byteValue), &config)

	return config
}

func ConvertLamportsToSol(fp string) string {
	fpValue, err := strconv.ParseFloat(fp, 64)
	if err != nil {
		return "error"
	}

	solValue := fpValue / SOL_LAMPORTS

	return fmt.Sprintf(fmt.Sprintf("%f", solValue))
}

func InputInteger(label string) int {
	validate := func(input string) error {
		_, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	res, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1
	}

	value, err := strconv.Atoi(res)
	if err != nil {
		return -1
	}

	return value
}

func InputString(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func PriceAlertConditionMenu() string {
	var choices = []string{"Up", "Down"}

	prompt := promptui.Select{
		Label: "Price Alert Condition",
		Items: choices,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "error"
	}

	switch result {
	case choices[0]:
		return "UP"
	case choices[1]:
		return "DOWN"
	default:
		return "error"
	}
}

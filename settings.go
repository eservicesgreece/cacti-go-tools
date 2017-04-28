package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func dumpConfig() {
	for setting, value := range viper.AllSettings() {
		if checkIfIsMap(value) {
			fmt.Println(setting)
			for option, oValue := range value.(map[string]interface{}) {
				fmt.Println(option, "=", oValue)
			}
		} else {
			fmt.Println("Setting:", setting, "=", value)
		}
	}
}

func setupConfig() {
	viper.SetConfigName("cacti-go-tools") // name of config file (without extension)
	viper.SetConfigType("json")           // Set type to json
	// Set Configuration File paths
	viper.AddConfigPath("/etc/cacti-go-tools/")
	viper.AddConfigPath("$HOME/.cacti-go-tools")
	viper.AddConfigPath(".")

	//Fetch Configuration
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		if *installtype != "" {
			downloadConfig("/etc/cacti-go-tools/", *installconfurl)
			os.Exit(0)
		} else {
			fmt.Println("Config file not found, please create it under /etc/cacti-go-tools/cacti-go-tools.json or run cacti-go-tools install config")
			os.Exit(1)
		}
	}
}

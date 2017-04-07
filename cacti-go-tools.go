package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {

	//Setup flag parsing
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Version(version)
	appFlags := kingpin.Parse()

	// check if we are executed by a user
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		interactive = true
	} else {
		interactive = false
	}

	viper.SetConfigName("cacti-go-tools") // name of config file (without extension)
	viper.SetConfigType("json")           // Set type to json
	//Set Configuration File paths
	viper.AddConfigPath("/etc/cacti-go-tools/")
	viper.AddConfigPath("$HOME/.cacti-go-tools")
	viper.AddConfigPath(".")

	//Fetch Configuration
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		if *installconf != "" {
			downloadConfig("/etc/cacti-go-tools/", *installconfurl)
			os.Exit(0)
		} else {
			fmt.Println("Config file not found, please create it under /etc/cacti-go-tools/cacti-go-tools.json or run cacti-go-tools install config")
			os.Exit(1)
		}
	}

	//Flag Actions
	switch appFlags {
	case "config":
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
		break
	case "engine":
		switch *enginetype {
		case "nginx":
			fmt.Printf(nginxStatus(fetchURL(makeURL(viper.GetString("phpfpm.uri"), viper.GetString("phpfpm.path")))))
			break
		case "phpfpm":
			fmt.Printf(phpfpmStatus(fetchURL(makeURL(viper.GetString("phpfpm.uri"), viper.GetString("phpfpm.path")))))
			break
		case "pagespeed":
			break
		case "bind":
			switch *engineoptions {
			case "requests":
				fmt.Printf(bindStatus(combinePath(viper.GetString("bind.uri"), viper.GetString("bind.filename")), "requests"))
				break
			case "queries":
				fmt.Printf(bindStatus(combinePath(viper.GetString("bind.uri"), viper.GetString("bind.filename")), "queries"))
				break
			case "nsstats":
				fmt.Printf(bindStatus(combinePath(viper.GetString("bind.uri"), viper.GetString("bind.filename")), "nsstats"))
				break
			default:
				fmt.Printf(bindStatus(combinePath(viper.GetString("bind.uri"), viper.GetString("bind.filename")), "queries"))
				break
			}
		default:
			fmt.Println("Engine " + *enginetype + " does not exist.")
			break
		}
	case "test test":
		fmt.Println(combinePath(`\test\`, `\filename`))
		fmt.Printf(makeURL(viper.GetString("phpfpm.uri"), viper.GetString("phpfpm.path")))
		break
	case "test nginxtest":
		fmt.Printf("nginx snmp tests " + *nginxtesthost)
		break
	case "install":
		if *installconf != "" {
			downloadConfig("/etc/cacti-go-tools/", *installconfurl)
		}
		if *installbin != "" {
			copyFile("cacti-go-tools", "/usr/bin/", "cacti-go-tools")
		}
	default:
		fmt.Printf("Mistakes were made.")
		break
	}
}

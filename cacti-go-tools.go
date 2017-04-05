package main

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

var version string
var buildstamp string
var hash string

//Set Commands, Flags and Args
var (
	app   = kingpin.New("cacti-go-tools", "Data Collection Engine")
	curl  = kingpin.Command("url", "Acquisition URL")
	curla = curl.Arg("Status URL", "Status URL").Required().URL()

	config = kingpin.Command("config", "Show Configuration")

	engine        = kingpin.Command("engine", "Acquisition Engine")
	enginetype    = engine.Arg("enginetype", "Supported Engines: nginx, php-fpm").Required().HintOptions("nginx php-fpm pagespeed bind").String()
	engineoptions = engine.Arg("engine options", "engine options").String()
)

func main() {

	viper.SetConfigName("cacti-go-tools") // name of config file (without extension)
	viper.SetConfigType("json")           // Set type to json
	//Find Configuration File
	viper.AddConfigPath("/etc/cacti-go-tools/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.cacti-go-tools") // call multiple times to add many search paths
	viper.AddConfigPath(".")                     // optionally look for config in the working directory

	//Fetch Configuration
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	//Setup flag parsing
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Version(version)
	kingpin.Parse()

	//Flag Actions
	switch kingpin.Parse() {
	case "config":
		fmt.Println("nginx")
		fmt.Println("URI :", viper.GetString("nginx.uri"))
		fmt.Println("Path :", viper.GetString("nginx.path"))
		fmt.Println("phpfpm")
		fmt.Println("URI :", viper.GetString("phpfpm.uri"))
		fmt.Println("Path :", viper.GetString("phpfpm.path"))
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
				fmt.Printf(bindStatus(makeURL(viper.GetString("bind.uri"), viper.GetString("bind.filename")), "requests"))
				break
			case "queries":
				fmt.Printf(bindStatus(makeURL(viper.GetString("bind.uri"), viper.GetString("bind.filename")), "queries"))
				break
			case "nsstats":
				fmt.Printf(bindStatus(makeURL(viper.GetString("bind.uri"), viper.GetString("bind.filename")), "nsstats"))
				break
			default:
				fmt.Printf(bindStatus(makeURL(viper.GetString("bind.uri"), viper.GetString("bind.filename")), "queries"))
				break
			}
		default:
			fmt.Println("Engine " + *enginetype + " does not exist.")
			break
		}
	default:
		fmt.Printf("Mistakes were made.")
		break
	}
}

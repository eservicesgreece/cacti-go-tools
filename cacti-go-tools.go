package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

var version string
var buildstamp string
var hash string

var interactive bool

//Set Commands, Flags and Args
var (
	app   = kingpin.New("cacti-go-tools", "Data Collection Engine")
	curl  = kingpin.Command("url", "Acquisition URL")
	curla = curl.Arg("Status URL", "Status URL").Required().URL()

	config = kingpin.Command("config", "Show Configuration")

	engine        = kingpin.Command("engine", "Acquisition Engine")
	enginetype    = engine.Arg("enginetype", "Supported Engines: nginx, php-fpm").Required().HintOptions("nginx php-fpm pagespeed bind").String()
	engineoptions = engine.Arg("engine options", "engine options").String()

	test          = kingpin.Command("test", "Testing tools")
	nginxtest     = test.Command("nginx", "Test SNMP Acquisition")
	nginxtesthost = nginxtest.Arg("host", "Host to test").Required().String()
	testuser      = test.Command("test", "test")

	install        = kingpin.Command("install", "Install cacti-go-tools")
	installconf    = install.Arg("config", "Installs default configuration").String()
	installconfurl = install.Flag("configurl", "Configuration file URL").String()
	installbin     = install.Arg("binary", "Copies the cacti-go-tools binary in /usr/local/bin").String()
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
	//Find Configuration File
	viper.AddConfigPath("/etc/cacti-go-tools/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.cacti-go-tools") // call multiple times to add many search paths
	viper.AddConfigPath(".")                     // optionally look for config in the working directory

	//Fetch Configuration
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
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

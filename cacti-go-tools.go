//go:generate goversioninfo
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	fullversion := "cacti-go-tools (c) 2017 eServices Greece | https://eservices-greece.com\nVersion: " + version + " | Build: " + buildstamp + "\nBuild on: " + hash + "\nGITHUB: https://github.com/eservicesgreece/cacti-go-tools"
	kingpin.Version(fullversion)               //set our version
	kingpin.CommandLine.HelpFlag.Short('h')    //enable help short
	kingpin.CommandLine.VersionFlag.Short('v') //enable version short

	var appFlags = kingpin.Parse() //Setup flag parsing
	setupConfig()                  //Setup config file details, parse and fetch if needed

	logLvl, _ := logrus.ParseLevel(viper.GetString("logging.level"))
	logfile, _ := os.OpenFile(combinePath(viper.GetString("logging.uri"), viper.GetString("logging.path")), os.O_WRONLY|os.O_CREATE, 0755)

	if viper.GetBool("logging.enabled") == false {
		log.Out = ioutil.Discard
	}

	logrus.SetLevel(logLvl)
	logrus.SetFormatter(new(logrus.TextFormatter))
	logrus.SetOutput(logfile)

	//Flag Actions
	switch appFlags {
	case "config":
		dumpConfig()
		break
	case "ver":
		fmt.Println(version)
		break
	case "engine":
		switch *enginetype {
		case "nginx":
			fmt.Printf(nginxStatus(fetchURL(makeURL(viper.GetString("nginx.uri"), viper.GetString("nginx.path")))))
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
		case "ntp":
			fmt.Printf(loopStats(combinePath(viper.GetString("ntp.uri"), viper.GetString("ntp.filename"))))
			break
		default:
			fmt.Println("Engine " + *enginetype + " does not exist.")
			break
		}
	case "test test":

		break
	case "test nginxtest":
		fmt.Printf("nginx snmp tests " + *nginxtesthost)
		break
	case "install":
		switch *installtype {
		case "config":
			downloadConfig("/etc/cacti-go-tools/", *installconfurl)
			break
		case "binary":
			copyFile("cacti-go-tools", "/usr/bin/", "cacti-go-tools")
		}
	default:
		fmt.Printf("Mistakes were made.")
		break
	}
}

package main

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var version string
var buildstamp string
var hash string

var interactive = terminal.IsTerminal(int(os.Stdout.Fd())) // check if we are executed by a user

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

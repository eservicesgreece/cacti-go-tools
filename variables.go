package main

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var version string
var buildstamp string
var hash string

var log = logrus.New()

var cactiConf = "cacti-go-tools.json"

var ( //Set Commands, Flags and Args
	app = kingpin.New("cacti-go-tools", "Data Collection Engine")
	ver = app.Flag("ver", "Ver").Bool()

	config = kingpin.Command("config", "Show Configuration")

	engine        = kingpin.Command("engine", "Acquisition Engine")
	enginetype    = engine.Arg("enginetype", "Supported Engines: nginx, php-fpm, bind, ntp").Required().HintOptions("nginx php-fpm pagespeed bind ntp").String()
	engineoptions = engine.Arg("engine options", "engine options").String()

	test          = kingpin.Command("test", "Testing tools")
	nginxtest     = test.Command("nginx", "Test SNMP Acquisition")
	nginxtesthost = nginxtest.Arg("host", "Host to test").Required().String()
	testtest      = test.Command("test", "test")

	install        = kingpin.Command("install", "Install cacti-go-tools")
	installtype    = install.Arg("config or binary", "Installs default configuration or binary").String()
	installconfurl = install.Flag("configurl", "Configuration file URL").String()
)

package main

import (
	"net/url"
	"path"
	"path/filepath"
	"reflect"

	"fmt"

	"github.com/spf13/viper"
)

func combinePath(userPath, file string) string {
	return filepath.Join(userPath, file)
}

func makeURL(uri, filePath string) string {
	url, err := url.Parse(uri)
	if err != nil {
		fmt.Printf("Incorrect URI: %v", uri)
	}

	url.Path = path.Join(url.Path, filePath)
	return url.String()
}

func onemptyreturnzero(tag string, engine string) string {
	if (tag == "") && (viper.GetBool(engine + ".onemptyreturnzero")) {
		return "0"
	}
	return tag
}

func checkIfIsMap(typeFC interface{}) bool {
	if reflect.ValueOf(typeFC).Kind() == reflect.Map {
		return true
	}
	return false
}

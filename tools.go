//go:generate goversioninfo 
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"reflect"

	"strings"

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

func logLastLine(file string) []string {
	var oldOffset int64

	logFile, err := os.Open(file)
	if err != nil {
		fmt.Printf("File: " + file + " not found.")
		os.Exit(1)
	}

	defer logFile.Close()

	reader := bufio.NewReader(logFile)

	lastLineSize := 0

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		lastLineSize = len(line)
	}

	logStat, err := os.Stat(file)
	buffer := make([]byte, lastLineSize)

	offset := logStat.Size() - int64(lastLineSize+1)
	numRead, err := logFile.ReadAt(buffer, offset)

	if oldOffset != offset {
		buffer = buffer[:numRead]
		oldOffset = offset
	}
	return strings.Fields(string(buffer))
}

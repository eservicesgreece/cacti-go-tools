package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchConfig(confURL string) {
	var cactiConf = "cacti-go-tools.json"
	//create our file
	fileOut, err := os.Create(cactiConf)
	if err != nil {
		fmt.Printf("File Error: %v", err)
	}
	defer fileOut.Close()

	_, err = os.Stat(cactiConf)
	if err != nil {
	} else {
		fileBody, err := http.Get(confURL)
		if err != nil {
			fmt.Printf("Download Error: %v", err)
		} else {
			io.Copy(fileOut, fileBody.Body)
			fmt.Printf("File Downloaded")
		}
		defer fileBody.Body.Close()
	}
}

func copyFile(srcFile, dstPath, dstFile string) bool {

	err := os.MkdirAll(dstPath, os.ModeDir)
	if err != nil {
		fmt.Printf("Cant create path: %v", dstPath)
	}

	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Printf("Source File Error: %v", err)
		defer src.Close()
		return false
	}
	dst, err := os.Create(dstFile)
	if err != nil {
		fmt.Printf("Destination File Error: %v", err)
		defer dst.Close()
		return false
	}
	_, err = io.Copy(src, dst)
	if err != nil {
		fmt.Printf("File Copy Error: %v", err)
		return false
	}

	return true
}

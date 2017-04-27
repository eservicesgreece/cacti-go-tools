package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mitchellh/go-homedir"
)

func getUsersHome() string {
	homePath, err := homedir.Dir()
	if err != nil {

	}
	return (homePath)
}

func downloadConfig(location, url string) {
	fmt.Println("Downloading Default Config")
	if url != "" {
		fetchConfig(url)
		copyFile(combinePath(getUsersHome(), cactiConf), location, cactiConf)
	} else {
		fetchConfig("https://raw.githubusercontent.com/eservicesgreece/cacti-go-tools/master/cacti-go-tools.json")
		copyFile(combinePath(getUsersHome(), cactiConf), location, cactiConf)
	}
}

func fetchConfig(confURL string) {
	//create our file
	fmt.Println(combinePath(getUsersHome(), cactiConf))
	fileOut, err := os.Create(combinePath(getUsersHome(), cactiConf))
	if err != nil {
		fmt.Println("File Creation Error: %v", err)
	}
	defer fileOut.Close()

	_, err = os.Stat(combinePath(getUsersHome(), cactiConf))
	if err != nil {
		fmt.Println("File Error: %v", err)
	} else {
		fileBody, err := http.Get(confURL)
		if err != nil {
			fmt.Println("Download Error: %v", err)
		} else {
			io.Copy(fileOut, fileBody.Body)
			fmt.Println("File Downloaded")
		}
		defer fileBody.Body.Close()
	}
	//	defer fileOut.Close()
}

func copyFile(srcFile, dstPath, dstFile string) bool {

	err := os.MkdirAll(dstPath, os.ModeDir)
	if err != nil {
		fmt.Println("Cant create path: %v", dstPath)
	}

	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("Source File Error: %v", err)
		defer src.Close()
		return false
	}

	dst, err := os.Create(combinePath(dstPath, dstFile))
	if err != nil {
		fmt.Printf("Destination File Error: %v", err)
		defer dst.Close()
		return false
	}

	fmt.Println(srcFile)
	fmt.Println(combinePath(dstPath, dstFile))
	_, err = io.Copy(src, dst)
	if err != nil {
		fmt.Println("File Copy Error: %v", err)
		return false
	}

	err = dst.Close()
	if err != nil {
		fmt.Println("Error: %v", err)
		return false
	}
	return true
}

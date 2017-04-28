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

	err := os.Remove(combinePath(getUsersHome(), cactiConf))
	if err != nil {
		fmt.Println(err)
	}
}

func fetchConfig(confURL string) {
	//create our file
	fmt.Println(combinePath(getUsersHome(), cactiConf))
	fileOut, err := os.Create(combinePath(getUsersHome(), cactiConf))
	if err != nil {
		fmt.Printf("File Creation Error: %v\n", err)
	}
	defer fileOut.Close()

	_, err = os.Stat(combinePath(getUsersHome(), cactiConf))
	if err != nil {
		fmt.Printf("File Error: %v\n", err)
	} else {
		fileBody, err := http.Get(confURL)
		if err != nil {
			fmt.Printf("Download Error: %v\n", err)
		} else {
			io.Copy(fileOut, fileBody.Body)
			fmt.Println("File Downloaded")
		}
		defer fileBody.Body.Close()
	}
}

func copyFile(srcFile, dstPath, dstFile string) bool {

	err := os.MkdirAll(dstPath, os.ModeDir)
	if err != nil {
		fmt.Printf("Cant create path: %v\n", dstPath)
	}

	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Printf("Source File Error: %v\n", err)
		defer src.Close()
		return false
	}

	dst, err := os.Create(combinePath(dstPath, dstFile))
	if err != nil {
		fmt.Printf("Destination File Error: %v\n", err)
		defer dst.Close()
		return false
	}

	_, err = io.Copy(dst, src)
	if err != nil {
		fmt.Printf("File Copy Error: %v\n", err)
		return false
	}

	err = dst.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return false
	}
	return true
}

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func fetchURL(statURL string) string {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	var reply []byte
	resp, err := netClient.Get(statURL)

	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		reply, _ = ioutil.ReadAll(resp.Body)
	}

	if resp.StatusCode == 200 {
		return string(reply)
	}
	return "http/" + resp.Status

}

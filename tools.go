package main

import "github.com/spf13/viper"

func makeURL(uri, path string) string {
	rURL := ""

	if path[len(path)-1:] == `/` {
		path = path[:len(path)-1]
	}

	if (uri[len(uri)-1:] == `/`) && (path[:1] == `/`) {
		rURL = uri[:len(uri)-1] + path
	} else {
		if (uri[len(uri)-1:] == `/`) || (path[:1] == `/`) {
			rURL = uri + path
		} else {
			rURL = uri + `/` + path
		}
	}
	return rURL
}

func onemptyreturnzero(tag []byte, engine string) []byte {
	if (tag == nil) && (viper.GetBool(engine + ".onemptyreturnzero")) {
		return []byte("0")
	}
	return tag
}

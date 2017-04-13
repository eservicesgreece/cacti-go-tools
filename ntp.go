package main

import "strings"

func loopStats(file string) string {
	return strings.Join(logLastLine(file), "\n")
}

package main

import (
	"regexp"
	"strings"
)

func phpfpmStatus(httpReply string) string {
	//find all numbers
	var connections = regexp.MustCompile(`\d+`)
	//match all new lines
	var removeNewLines = regexp.MustCompile(`\n`)
	//match everything up to and the accepted con string
	var removeHeader = regexp.MustCompile(`^.*accepted conn.`)

	//Return http status code on error
	if httpReply[:5] == "http/" {
		return httpReply
	}

	//trim all whitespace and remove new lines
	cleanReply := removeNewLines.ReplaceAllString(strings.TrimSpace(httpReply), "")
	//remove everything that matches regexp
	cleanReply = removeHeader.ReplaceAllString(cleanReply, "")
	//Remove everything but the numbers
	phpfpmValues := connections.FindAllStringSubmatch(cleanReply, -1)
	//iterate over all values and create []string
	statusValues := make([]string, len(phpfpmValues))
	for i := range statusValues {
		statusValues[i] = phpfpmValues[i][0]
	}
	//Join all values with a new line intersected
	return strings.Join(statusValues, "\n")
}

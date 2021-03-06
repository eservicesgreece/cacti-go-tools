package main

import (
	"regexp"
	"strings"
)

func nginxStatus(httpReply string) string {
	//Find all Numbers in string
	var connections = regexp.MustCompile(`\d+`)
	//Match all new lines
	var removeNewLines = regexp.MustCompile(`\n`)

	//Return http status code on error
	if httpReply[:5] == "http/" {
		return httpReply
	}

	//trim all whitespace and remove new lines
	cleanReply := removeNewLines.ReplaceAllString(strings.TrimSpace(httpReply), "")
	//Remove everything but the numbers
	nginxValues := connections.FindAllStringSubmatch(cleanReply, -1)
	//iterate over all values and create []string
	statusValues := make([]string, len(nginxValues))
	for i := range statusValues {
		statusValues[i] = nginxValues[i][0]
	}
	//Join all values with a new line intersected
	return strings.Join(statusValues, "\n")
}

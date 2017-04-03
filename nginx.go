package main

import "regexp"
import "strings"

func nginxStatus(httpReply string) string {
	var connections = regexp.MustCompile(`\d+`)
	var removeNewLines = regexp.MustCompile(`\n`)

	if httpReply[:5] == "http/" {
		return httpReply
	}

	cleanReply := removeNewLines.ReplaceAllString(strings.TrimSpace(httpReply), "")

	nginxValues := connections.FindAllStringSubmatch(cleanReply, -1)

	statusValues := make([]string, len(nginxValues))
	for i := range statusValues {
		statusValues[i] = nginxValues[i][0]
	}

	return strings.Join(statusValues, "\n")

}

func phpfpmStatus(httpReply string) string {

	return "bar"
}

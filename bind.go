package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func parseStatistics(stats []string, statsFile string) []string {
	var statusValues []string
	bindstats, err := ioutil.ReadFile(statsFile)
	if err != nil {
		fmt.Println("File " + statsFile + " Not Found")
	} else {
		for _, element := range stats {
			_ = element
			idTag := regexp.MustCompile(`(?m)(\d+\s` + element + `$)`)
			tag := idTag.FindAllString(string(bindstats), -1)

			removeTag := regexp.MustCompile(`\d+ `)
			cleanTag := removeTag.FindString(tag[len(tag)-1])
			cleanTag = onemptyreturnzero(cleanTag, "bind")

			statusValues = append(statusValues, string(cleanTag))
		}
	}
	return statusValues
}

func bindStatus(statisticsfile, qType string) string {
	incommingRequests := []string{"QUERY", "STATUS", "NOTIFY", "UPDATE"}
	incommingQueries := []string{"A", "NS", "CNAME", "SOA", "PTR", "HINFO", "MX", "TXT", "AAAA", "EID", "SRV", "NAPTR", "A6", "DNAME", "DS", "RRSIG", "DNSKEY", "SPF", "AXFR", "ANY", "Others"}
	nsStatistics := []string{`IPv4 requests received`, `IPv6 requests received`, `requests with EDNS\(0\) received`, `TCP requests received`, `auth queries rejected`, `recursive queries rejected`, `transfer requests rejected`, `update requests rejected`, `responses sent`, `truncated responses sent`, `responses with EDNS\(0\) sent`, `queries resulted in successful answer`, `queries resulted in authoritative answer`, `queries resulted in nxrrset`, `queries resulted in NXDOMAIN`, `other query failures`}

	var statusValues []string

	switch qType {
	case "requests":
		statusValues = parseStatistics(incommingRequests, statisticsfile)
		break
	case "queries":
		statusValues = parseStatistics(incommingQueries, statisticsfile)
		break
	case "nsstats":
		statusValues = parseStatistics(nsStatistics, statisticsfile)
		break
	}
	return strings.Join(statusValues, "\n")
}

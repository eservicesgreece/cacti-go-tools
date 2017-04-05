package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func bindStatus(statisticsfile, qType string) string {
	incommingRequests := []string{"QUERY", "STATUS", "NOTIFY", "UPDATE"}
	incommingQueries := []string{"A", "AAAA", "NS", "CNAME", "SOA", "PTR", "HINFO", "MX", "TXT", "EID", "SRV", "NAPTR", "A6", "DNAME", "DS", "RRSIG", "DNSKEY", "SPF", "AXFR", "ANY", "Others"}
	nsStatistics := []string{`IPv4 requests received`, `IPv6 requests received`, `requests with EDNS\(0\) received`, `TCP requests received`, `auth queries rejected`, `recursive queries rejected`, `transfer requests rejected`, `update requests rejected`, `responses sent`, `truncated responses sent`, `responses with EDNS\(0\) sent`, `queries resulted in successful answer`, `queries resulted in authoritative answer`, `queries resulted in nxrrset`, `queries resulted in NXDOMAIN`, `other query failures`}

	var statusValues []string

	bindstats, err := ioutil.ReadFile(statisticsfile)
	if err != nil {
		fmt.Println("File " + statisticsfile + " Not Found")
	} else {

		switch qType {
		case "requests":
			for _, element := range incommingRequests {
				idTag := regexp.MustCompile(`(?m:\d+ ` + element + "$)")
				tag := idTag.Find(bindstats)

				removeTag := regexp.MustCompile(`\d+ `)
				cleanTag := removeTag.Find(tag)

				statusValues = append(statusValues, string(cleanTag))
			}
			break
		case "queries":
			for _, element := range incommingQueries {
				idTag := regexp.MustCompile(`(?m:\d+ ` + element + "$)")
				tag := idTag.Find(bindstats)

				removeTag := regexp.MustCompile(`\d+ `)
				cleanTag := removeTag.Find(tag)

				statusValues = append(statusValues, string(cleanTag))
			}
			break
		case "nsstats":
			for _, element := range nsStatistics {
				idTag := regexp.MustCompile(`(?m:\d+ ` + element + "$)")
				tag := idTag.Find(bindstats)

				removeTag := regexp.MustCompile(`\d+ `)
				cleanTag := removeTag.Find(tag)

				statusValues = append(statusValues, string(cleanTag))
			}
			break
		}

	}
	return strings.Join(statusValues, "\n")
}

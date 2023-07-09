package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("domain,hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord \n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func checkDomain(domain string) {
	var hasMx, hasSPF, hasDMARC bool

	var spfRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("err:%v", err)
	}

	if len(mxRecord) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("err:%v", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc" + domain)
	if err != nil {
		log.Printf("error: %v", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v,%v, %v, %v,%v, %v", domain, hasMx, hasSPF, spfRecord,hasDMARC,dmarcRecord)
}

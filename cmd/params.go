package cmd

import (
	"flag"
	"fmt"
	"log"
	"time"
)

type parameters struct {
	CredentialFilePath string
	StartTimeOfRange   time.Time
	EndTimeOfRange     time.Time
}

var params parameters

func ParseParameters() {
	startDateStr, endDateStr := "", ""

	flag.StringVar(&params.CredentialFilePath, "credential-file", DefaultCredentialFilePath, "(Option) Download client_secret_*.json from Google Developer Console, and specifiled path")
	flag.StringVar(&startDateStr, "start-date", "", fmt.Sprintf("(Option) Starting date of the range for the sync (Default: %s)", time.Now().Format("2006/01/02")))
	flag.StringVar(&endDateStr, "end-date", "", fmt.Sprintf("(Option) Ending date of the range for the sync (Default: %s)", time.Now().Add(DefaultDateRangeLength).Format("2006/01/02")))
	flag.Parse()

	var date time.Time
	var err error

	if startDateStr == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006/01/02", startDateStr)
		if err != nil {
			log.Fatalf("Invalid date format (Expect: 2006/01/02)")
		}
	}
	params.StartTimeOfRange = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)

	if endDateStr == "" {
		date = time.Now().Add(DefaultDateRangeLength)
	} else {
		date, err = time.Parse("2006/01/02", startDateStr)
		if err != nil {
			log.Fatalf("Invalid date format (Expect: 2006/01/02)")
		}
	}
	date = date.AddDate(0, 0, 1)
	params.EndTimeOfRange = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)

	fmt.Println(params)
}

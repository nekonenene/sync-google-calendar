package cmd

import (
	"fmt"
	"log"
	"time"
)

func Exec() {
	ParseParameters()

	service, err := GetService()
	if err != nil {
		log.Fatalf("Failed to get a service: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := service.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}

	fmt.Println("今後のイベント：")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}
}

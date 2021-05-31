package cmd

import (
	"fmt"
	"log"
	"time"
)

func Exec() {
	ParseParameters()

	fmt.Println("まずは同期元の Google カレンダーとの連携をおこないます")
	service, err := GetService()
	if err != nil {
		log.Fatalf("Failed to get a service: %v", err)
	}

	startTime := params.StartTimeOfRange.Format(time.RFC3339)
	endTime := params.EndTimeOfRange.Format(time.RFC3339)

	events, err :=
		service.Events.
			List("primary").
			ShowDeleted(false).
			SingleEvents(true).
			TimeMin(startTime).
			TimeMax(endTime).
			MaxResults(2500).
			OrderBy("startTime").
			Do()

	if err != nil {
		log.Fatalf("Unable to retrieve user's events: %v", err)
	}

	if len(events.Items) == 0 {
		fmt.Println("カレンダーにイベントがありませんでした。終了します")
	}

	fmt.Println("以下のイベントを同期します：")
	for _, item := range events.Items {
		date := item.Start.DateTime
		if date == "" {
			date = item.Start.Date
		}
		fmt.Printf("%v (%v)\n", item.Summary, date)
	}
}

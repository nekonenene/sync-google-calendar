package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"google.golang.org/api/calendar/v3"
)

func Exec() {
	ParseParameters()

	fmt.Println("まずは同期元の Google カレンダーとの連携をおこないます")
	tokenPath := ""
	if params.UseTokenCache {
		tokenPath = FromAccountTokenPath
	}
	fromAccountService, err := GetService(tokenPath)
	if err != nil {
		log.Fatalf("Failed to get a service: %v", err)
	}

	startTime := params.StartTimeOfRange.Format(time.RFC3339)
	endTime := params.EndTimeOfRange.Format(time.RFC3339)

	fromAccountEvents, err :=
		fromAccountService.Events.
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

	if len(fromAccountEvents.Items) == 0 {
		fmt.Println("カレンダーにイベントがありませんでした。終了します")
	}

	fmt.Println("イベントの取得が完了しました")

	fmt.Println("次に予定作成をおこなう Google カレンダーとの連携をおこないます")
	tokenPath = ""
	if params.UseTokenCache {
		tokenPath = ToAccountTokenPath
	}
	toAccountService, err := GetService(tokenPath)
	if err != nil {
		log.Fatalf("Failed to get a service: %v", err)
	}

	toAccountEvents, err :=
		toAccountService.Events.
			List("primary").
			ShowDeleted(false).
			SingleEvents(true).
			TimeMin(startTime).
			TimeMax(endTime).
			MaxResults(2500).
			OrderBy("startTime").
			Do()

	// Choose events to copy
	var createEvents []*calendar.Event
	for _, fromAccountEvent := range fromAccountEvents.Items {
		startTime := fromAccountEvent.Start.DateTime
		endTime := fromAccountEvent.End.DateTime
		if startTime == "" || endTime == "" {
			continue // Ignore all day event
		}

		existDuplicateEvent := false
		for _, toAccountEvent := range toAccountEvents.Items {
			if startTime == toAccountEvent.Start.DateTime && endTime == toAccountEvent.End.DateTime {
				existDuplicateEvent = true
				break
			}
		}

		if !existDuplicateEvent {
			createEvents = append(createEvents, fromAccountEvent)
		}
	}

	if len(createEvents) == 0 {
		fmt.Println("作成すべきイベントがありませんでした（同じ開始・終了時刻のイベントや終日イベントは作成されません）")
		fmt.Println("終了します")
		return
	}

	fmt.Println("\n以下のイベントを作成します：")
	for _, event := range createEvents {
		date := event.Start.DateTime
		if date == "" {
			date = event.Start.Date
		}
		fmt.Printf("%s (Starting at: %s)\n", event.Summary, date)
	}

	fmt.Printf("\nよろしいですか？[y|n]: ")
	var response string
	fmt.Scanln(&response)
	if !strings.EqualFold(response, "y") && !strings.EqualFold(response, "yes") {
		fmt.Println("終了します")
		return
	}

	for _, event := range createEvents {
		var summary, description string
		if params.TitleOverwrite != "" {
			summary = params.TitleOverwrite
		} else {
			summary = event.Summary
		}
		summary = params.TitlePrefix + summary

		if params.DescriptionOverwrite != "" {
			description = params.DescriptionOverwrite
		} else {
			description = event.Description
		}

		insertEvent := &calendar.Event{
			Start:          event.Start,
			End:            event.End,
			Summary:        summary,
			Description:    description,
			Location:       event.Location,
			ConferenceData: event.ConferenceData,
		}

		_, err := toAccountService.Events.Insert("primary", insertEvent).Do()
		if err != nil {
			log.Fatalf("Failed to insert an event: %v", err)
		}
	}
	fmt.Println("完了しました")
}

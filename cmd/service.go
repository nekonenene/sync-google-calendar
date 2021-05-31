package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Get Google Calendar service
func GetService() (*calendar.Service, error) {
	b, err := ioutil.ReadFile(params.CredentialFilePath)
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
	if err != nil {
		return nil, err
	}

	tok := getTokenFromWeb(config)
	ctx := context.Background()
	service, err := calendar.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, tok)))
	if err != nil {
		return nil, err
	}

	return service, nil
}

// Get token through OAuth2 authorization using web browser
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Open the following URL using your web browser, and allow accessing to your Google Calendar: \n%v\n\n", authURL)
	openBrowser(authURL)

	fmt.Printf("After authentication, please input the code: ")

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Open URL using the default web browser
func openBrowser(url string) {
	os := runtime.GOOS
	switch os {
	case "windows":
		exec.Command("start", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	case "linux":
		exec.Command("xgd-open", url).Start()
	}
}

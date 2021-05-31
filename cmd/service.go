package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Get Google Calendar service
func GetService(tokenPath string) (*calendar.Service, error) {
	b, err := ioutil.ReadFile(params.CredentialFilePath)
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
	if err != nil {
		return nil, err
	}

	var tok *oauth2.Token
	if tokenPath == "" {
		tok = getTokenFromWeb(config)
	} else {
		tok = getTokenFromFile(config, tokenPath)
	}

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

// Get token from local file or OAuth2 authorization
func getTokenFromFile(config *oauth2.Config, tokenPath string) *oauth2.Token {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first time.
	tok, err := tokenFromFile(tokenPath)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokenPath, tok)
	}
	return tok
}

// Get token from local
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Store token
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
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

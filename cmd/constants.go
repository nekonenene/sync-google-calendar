package cmd

import "time"

const (
	DefaultCredentialFilePath = "./client_secret.json"
	DefaultDateRangeLength    = 24 * time.Hour * 6
	FromAccountTokenPath      = "./token_from.json"
	ToAccountTokenPath        = "./token_to.json"
)

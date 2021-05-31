package cmd

import (
	"flag"
)

type parameters struct {
	CredentialFilePath string
	TokenFilePath      string
}

var params parameters

func ParseParameters() {
	flag.StringVar(&params.CredentialFilePath, "credential-file", DefaultCredentialFilePath, "(Option) Download client_secret_*.json from Google Developer Console, and specifiled path")
	flag.StringVar(&params.TokenFilePath, "token-file", DefaultTokenFilePath, "(Option) If you want to use your token file, specifiled path")
	flag.Parse()
}

package cmd

import (
	"flag"
)

type parameters struct {
	CredentialFilePath string
}

var params parameters

func ParseParameters() {
	flag.StringVar(&params.CredentialFilePath, "credential-file", DefaultCredentialFilePath, "(Option) Download client_secret_*.json from Google Developer Console, and specifiled path")
	flag.Parse()
}

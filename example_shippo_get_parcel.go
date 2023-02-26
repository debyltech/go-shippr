package main

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/debyltech/go-shippr/shippo"
)

func main() {
	apiKey := flag.String("apikey", "", "api key for XPS")
	flag.Parse()

	if *apiKey == "" {
		log.Fatal("missing -apikey")
	}

	client := shippo.NewClient(*apiKey)

	response, err := client.GetParcel("5b37410dde6d4d9ea707ab1ba303e9c4")
	if err != nil {
		log.Fatal(err)
	}
	jsonPretty, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(jsonPretty))
}

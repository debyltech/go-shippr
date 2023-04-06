package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/debyltech/go-shippr/shippo"
)

func main() {
	apiKey := flag.String("apikey", "", "api key for Shippo")
	flag.Parse()

	if *apiKey == "" {
		log.Fatal("missing -apikey")
	}

	client := shippo.NewClient(*apiKey)

	response, err := client.GetRatesForShipmentId("5d3dbe982334420495893a530ca3b302")
	if err != nil {
		log.Fatal(err)
	}
	jsonPretty, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(jsonPretty))
}

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

	request := shippo.LiveRateRequest{
		AddressFrom: shippo.Address{
			Name:       "de Byl Technologies LLC",
			Address1:   "176 Lull Rd",
			City:       "Weare",
			State:      "NH",
			Country:    "US",
			PostalCode: "03281",
		},
		AddressTo: shippo.Address{
			Name:       "Mrs Hippo",
			Address1:   "354 Patch Rd",
			City:       "Henniker",
			State:      "NH",
			Country:    "US",
			PostalCode: "03242",
			Phone:      "6035055790",
		},
		LineItems: []shippo.LineItem{{
			Quantity:           1,
			TotalPrice:         "49.99",
			Currency:           "USD",
			Weight:             "150",
			WeightUnit:         "g",
			Title:              "Deflector",
			ManufactureCountry: "US",
			Sku:                "DBD123",
		}},
		Parcel: shippo.Parcel{
			Length:       "12",
			Width:        "12",
			Height:       "10",
			DistanceUnit: "cm",
			Weight:       "150",
			WeightUnit:   "g",
		},
	}

	response, err := client.GetLiveRates(request)
	if err != nil {
		log.Fatal(err)
	}
	jsonPretty, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(jsonPretty))
}

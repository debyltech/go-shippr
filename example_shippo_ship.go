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

	request := shippo.ShipmentRequest{
		CarrierAccount:    "fa27e44374774bdea1819f4325f1347a",
		ServiceLevelToken: "usps_priority",
		Shipment: shippo.Shipment{
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
				Address1:   "965 Mission St.",
				City:       "San Francisco",
				State:      "CA",
				Country:    "US",
				PostalCode: "94105",
			},
			Parcels: []shippo.Parcel{{
				Length:       "12",
				Width:        "12",
				Height:       "10",
				DistanceUnit: "cm",
				Weight:       "150",
				WeightUnit:   "g",
			}},
		}}

	response, err := client.CreateShipment(request)
	if err != nil {
		log.Fatal(err)
	}
	jsonPretty, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(jsonPretty))
}

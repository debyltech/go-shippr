package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

	request := shippo.Shipment{
		AddressFrom: shippo.Address{
			Name:       "de Byl Technologies LLC",
			Address1:   "176 Lull Rd",
			City:       "Weare",
			State:      "NH",
			Country:    "US",
			PostalCode: "03281",
			Phone:      "6037484015",
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
		Parcels: []shippo.Parcel{{
			Length:       "6",
			Width:        "6",
			Height:       "4",
			DistanceUnit: "in",
			Weight:       "0.5",
			WeightUnit:   "lb",
		}},
	}

	response, err := client.CreateShipment(request)
	if err != nil {
		log.Fatal(err)
	}
	jsonPretty, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = client.AwaitQueuedFinished(response.Id)
	if err != nil {
		log.Fatal(err)
	}

	rates, err := client.GetRatesForShipmentId(response.Id)
	if err != nil {
		log.Fatal(err)
	}
	jsonPretty, err = json.MarshalIndent(rates, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonPretty))
	fmt.Println(response.Id)
}

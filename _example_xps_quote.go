package main

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/debyltech/go-shippr/xps"
)

func main() {
	apiKey := flag.String("apikey", "", "api key for XPS")
	customerId := flag.String("customerid", "", "customerId for XPS")
	flag.Parse()

	if *apiKey == "" {
		log.Fatal("missing -apikey")
	}
	if *customerId == "" {
		log.Fatal("missing -customerid")
	}

	client := xps.NewClient(*apiKey, *customerId)

	pieces := []xps.Piece{{
		Weight:          "120",
		Length:          "10",
		Width:           "10",
		Height:          "10",
		InsuranceAmount: "0",
	}}

	quoteRequest := xps.ShipmentRequest{
		CarrierCode:     "usps",
		ServiceCode:     "usps_priority",
		PackageTypeCode: "usps_custom_package",
		Sender: xps.QuoteAddressSender{
			Country: "US",
			ZIP:     "03281",
		},
		Receiver: xps.QuoteAddressReceiver{
			City:    "San Francisco",
			Country: "US",
			ZIP:     "94105",
			Email:   "test@debyltech.com",
		},
		Residential:     true,
		WeightUnit:      "g",
		DimensionUnit:   "cm",
		Currency:        xps.CurrencyUSD,
		CustomsCurrency: xps.CurrencyUSD,
		Pieces:          pieces,
		Billing: xps.ShipmentBillingQuote{
			Party: xps.ShipmentBillingPartySender,
		},
	}

	response, err := client.GetQuote(quoteRequest)
	if err != nil {
		log.Fatal(err)
	}
	jsonPretty, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(jsonPretty))
}

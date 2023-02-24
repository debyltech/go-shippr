package main

import (
	"flag"
	"io"
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

	quoteRequest := xps.QuoteRequest{
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
		Billing: xps.QuoteBilling{
			Party: xps.QuoteBillingPartySender,
		},
	}

	response, err := client.GetQuote(quoteRequest)
	if err != nil {
		log.Fatal(err)
	}
	if response.Status != "200 OK" {
		log.Printf("bad response status %s", response.Status)
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", string(bodyBytes))
}

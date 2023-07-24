package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/debyltech/go-shippr/shippo"
)

func main() {
	apiKey := flag.String("apikey", "", "api key for XPS")
	taxIdFlag := flag.String("taxid", "", "either EIN or IOSS")
	flag.Parse()

	taxId := strings.ToLower(*taxIdFlag)

	if *apiKey == "" {
		log.Fatal("missing -apikey")
	}
	if taxId != "ioss" || taxId != "ein" {
		log.Fatal("missing or invalid -taxid")
	}

	toAddress := shippo.Address{
		Name:       "Madcat Services",
		Address1:   "211g Castle Road",
		Address2:   "ATTN: Conor Mulholland",
		City:       "Randalstown",
		State:      "Antrim and Newtownabbey",
		Country:    "IE",
		PostalCode: "TA6 4XY",
		Phone:      "07941782556",
	}

	client := shippo.NewClient(*apiKey)

	customsItem, err := client.CreateCustomsItem(shippo.CustomsItem{
		Description:   "Wind Deflector",
		Quantity:      1,
		NetWeight:     "210",
		WeightUnit:    "g",
		Currency:      "USD",
		ValueAmount:   "39.99",
		OriginCountry: "US",
		Metadata:      "order:DBT-1020",
	})
	if err != nil {
		log.Fatal(err)
	}
	customsDeclarationRequest := shippo.CustomsDeclaration{
		Certify:           true,
		CertifySigner:     "Bastian de Byl",
		Items:             []string{customsItem.Id},
		NonDeliveryOption: shippo.NONDELIV_RETURN,
		ContentsType:      shippo.CONTYP_MERCH,
		Incoterm:          shippo.INCO_DDU,
	}

	/* Handle special Canadial EEL/PFC */
	if strings.ToLower(toAddress.Country) == "ca" {
		customsDeclarationRequest.EELPFC = shippo.EEL_NOEEI3036
	} else {
		customsDeclarationRequest.EELPFC = shippo.EEL_NOEEI3037a
	}

	switch taxId {
	case "ioss":
		customsDeclarationRequest.ExporterIdentification = shippo.ExporterIdentification{
			TaxId: shippo.CustomsTaxId{
				Number: "IM7240004756",
				Type:   shippo.TAX_IOSS,
			},
		}
		break

	case "ein":
		customsDeclarationRequest.ExporterIdentification = shippo.ExporterIdentification{
			TaxId: shippo.CustomsTaxId{
				Number: "922297356",
				Type:   shippo.TAX_EIN,
			},
		}
		break
	}

	if customsDeclarationRequest.ExporterIdentification.TaxId.Number == "" {
		log.Fatal("customs declarationm exporter null")
	}

	declaration, err := client.CreateCustomsDeclaration()
	if err != nil {
		log.Fatal(err)
	}

	request := shippo.Shipment{
		AddressFrom: shippo.Address{
			Name:       "de Byl Technologies LLC",
			Address1:   "1 Hardy Rd",
			Address2:   "PMB 199",
			City:       "Bedford",
			State:      "NH",
			Country:    "US",
			PostalCode: "03110",
			Phone:      "6034160859",
		},
		AddressTo: toAddress,
		Parcels: []shippo.Parcel{{
			Length:       "19",
			Width:        "19",
			Height:       "16.5",
			DistanceUnit: "cm",
			Weight:       "210",
			WeightUnit:   "g",
		}},
		CustomsDeclaration: declaration.Id,
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

	shipmentResponse, err := client.GetShipmentById(response.Id)
	if err != nil {
		log.Fatal(err)
	}
	jsonShipmentResponse, err := json.MarshalIndent(shipmentResponse, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonShipmentResponse))

	rates, err := client.GetRatesForShipmentId(response.Id)
	if err != nil {
		log.Fatal(err)
	}
	jsonPretty, err = json.MarshalIndent(rates, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonPretty))
	fmt.Println("Shipment ID:", response.Id)
}

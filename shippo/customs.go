package shippo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	helper "github.com/debyltech/go-helpers/json"
)

type EELPFC string
type NonDeliveryOption string
type ContentsType string
type Incoterm string
type TaxType string

const (
	CustomsItemUri        string = BaseUri + "/customs/items"
	CustomsDeclarationUri        = BaseUri + "/customs/declarations"

	// EEL/PFC Codes
	EEL_NOEEI3037a EELPFC = "NOEEI_30_37_a"
	EEL_NOEEI3037h        = "NOEEI_30_37_h"
	EEL_NOEEI3036         = "NOEEI_30_36"
	PFC_AESITN            = "AES_ITN"

	// Non-Delivery Options
	NONDELIV_ABANDON NonDeliveryOption = "ABANDON"
	NONDELIV_RETURN  NonDeliveryOption = "RETURN"

	// Customs Declaration Content Types
	CONTYP_DOC      ContentsType = "DOCUMENTS"
	CONTYP_GIFT                  = "GIFT"
	CONTYP_SAMPLE                = "SAMPLE"
	CONTYP_MERCH                 = "MERCHANDISE"
	CONTYP_HUMDON                = "HUMANITARIAN_DONATION"
	CONTYP_RETMERCH              = "RETURN_MERCHANDISE"
	CONTYP_OTHER                 = "OTHER"

	// Customs Incoterm
	INCO_DDU Incoterm = "DDU"
	INCO_DDP          = "DDP"

	// Tax Types
	TAX_EIN  TaxType = "EIN"
	TAX_VAT          = "VAT"
	TAX_IOSS         = "IOSS"
	TAX_ARN          = "ARN"
)

type CustomsItem struct {
	Id            string    `json:"object_id,omitempty"`
	Created       time.Time `json:"object_created,omitempty"`
	Description   string    `json:"description"`
	Quantity      int       `json:"quantity"`
	NetWeight     string    `json:"net_weight"`
	WeightUnit    string    `json:"mass_unit"`
	ValueAmount   float64   `json:"value_amount"`
	Currency      string    `json:"value_currency"`
	OriginCountry string    `json:"origin_country"`
	SKUCode       string    `json:"sku_code,oitempty"`
	ECCNEAR99     string    `json:"eccn_ear99,omitempty"`
	TariffNumber  string    `json:"tariff_number,omitempty"`
	Metadata      string    `json:"metadata"`
}

type CustomsTaxId struct {
	Number string  `json:"number"`
	Type   TaxType `json:"type"`
}

type ExporterIdentification struct {
	TaxId CustomsTaxId `json:"tax_id"`
}

type CustomsDeclaration struct {
	Id                     string                 `json:"object_id,omitempty"`
	Created                time.Time              `json:"object_created,omitempty"`
	CertifySigner          string                 `json:"certify_signer"`
	Certify                bool                   `json:"certify"`
	Items                  []string               `json:"items"`
	NonDeliveryOption      NonDeliveryOption      `json:"non_delivery_option"`
	ContentsType           ContentsType           `json:"contents_type"`
	ContentsExplanation    string                 `json:"contents_explanation,omitempty"`
	ExporterReference      string                 `json:"exporter_reference,omitempty"`
	ImporterReference      string                 `json:"importer_reference,omitempty"`
	Invoice                string                 `json:"invoice,omitempty"`
	License                string                 `json:"license,omitempty"`
	Certificate            string                 `json:"certificate,omitempty"`
	Notes                  string                 `json:"notes,omitempty"`
	EELPFC                 EELPFC                 `json:"eel_pfc,omitempty"`
	AESITN                 string                 `json:"aes_itn,omitempty"`
	Incoterm               Incoterm               `json:"incoterm"`
	VatCollected           bool                   `json:"is_vat_collected,omitempty"`
	B13aFilingOption       string                 `json:"b13a_filing_option,omitempty"`
	B13aNumber             string                 `json:"b13a_number,omitempty"`
	ExporterIdentification ExporterIdentification `json:"export_identification"`
	Metadata               string                 `json:"metadata"`
}

func (c *Client) CreateCustomsItem(item CustomsItem) (*CustomsItem, error) {
	data := url.Values{}
	data.Set("description", item.Description)
	data.Set("quantity", fmt.Sprintf("%d", item.Quantity))
	data.Set("net_weight", item.NetWeight)
	data.Set("mass_unit", item.WeightUnit)
	data.Set("value_amount", fmt.Sprintf("%.2f", item.ValueAmount))
	data.Set("value_currency", item.Currency)
	data.Set("tariff_number", item.TariffNumber)
	data.Set("origin_country", item.OriginCountry)
	data.Set("metadata", item.Metadata)
	data.Set("sku_code", item.SKUCode)

	request, err := http.NewRequest("POST", CustomsItemUri, bytes.NewBuffer([]byte(data.Encode())))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	request.Header.Set("Authorization", fmt.Sprintf("%s %s", BasicAuth, c.ApiKey))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var responseItem CustomsItem
	err = json.NewDecoder(response.Body).Decode(&responseItem)
	if err != nil {
		return nil, err
	}

	return &responseItem, nil
}

func (c *Client) CreatecustomsDeclaration(request CustomsDeclaration) (*CustomsDeclaration, error) {
	response, err := helper.Post(CustomsDeclarationUri, BasicAuth, c.ApiKey, request)
	if err != nil {
		return nil, err
	}
	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var responseDeclaration CustomsDeclaration
	err = json.NewDecoder(response.Body).Decode(&responseDeclaration)
	if err != nil {
		return nil, err
	}

	return &responseDeclaration, nil
}

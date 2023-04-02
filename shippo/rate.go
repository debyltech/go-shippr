package shippo

import (
	"encoding/json"
	"time"

	helper "github.com/debyltech/go-helpers/json"
)

const (
	GenerateRateUri string = BaseUri + "/live-rates"
)

type Rate struct {
	Id               string       `json:"object_id"`
	CarrierAccountId string       `json:"carrier_account"`
	ShipmentId       string       `json:"shipment"`
	Created          time.Time    `json:"object_created"`
	Attributes       []string     `json:"attributes"`
	Amount           string       `json:"amount"`
	Currency         string       `json:"currency"`
	AmountLocal      string       `json:"amount_local"`
	CurrencyLocal    string       `json:"currency_local"`
	EstimatedDays    int          `json:"estimated_days"`
	Provider         string       `json:"provider"`
	ServiceLevel     ServiceLevel `json:"servicelevel"`
	ProviderImage75  string       `json:"provider_image_75"`
	ProviderImage200 string       `json:"provider_image_200"`
}

type LiveRateRequest struct {
	AddressFrom Address    `json:"address_from"`
	AddressTo   Address    `json:"address_to"`
	LineItems   []LineItem `json:"line_items"`
	Parcel      Parcel     `json:"parcel,omitempty"`
}

type LiveRateResult struct {
	Amount        string `json:"amount"`
	AmountLocal   string `json:"amount_local"`
	Currency      string `json:"currency"`
	CurrencyLocal string `json:"currency_local"`
	EstimatedDays int    `json:"estimated_days"`
	Title         string `json:"title"`
}

type LiveRateResponse struct {
	Rates []LiveRateResult `json:"results"`
}

func (c *Client) GetLiveRates(request LiveRateRequest) (*LiveRateResponse, error) {
	response, err := helper.Post(GenerateRateUri, BasicAuth, c.ApiKey, request)
	if err != nil {
		return nil, err
	}
	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var rate LiveRateResponse
	err = json.NewDecoder(response.Body).Decode(&rate)
	if err != nil {
		return nil, err
	}

	return &rate, nil
}

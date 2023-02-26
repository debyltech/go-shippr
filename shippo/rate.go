package shippo

import (
	"encoding/json"

	helper "github.com/debyltech/go-helpers/json"
)

type LineItem struct {
	Quantity           int    `json:"quantity"`
	TotalPrice         string `json:"total_price"`
	Currency           string `json:"currency"`
	Weight             string `json:"weight"`
	WeightUnit         string `json:"weight_unit"`
	Title              string `json:"title"`
	ManufactureCountry string `json:"manufacture_country"`
	Sku                string `json:"sku"`
}

type RateRequest struct {
	AddressFrom Address    `json:"address_from"`
	AddressTo   Address    `json:"address_to"`
	LineItems   []LineItem `json:"line_items"`
	Parcel      Parcel     `json:"parcel"`
}

type RateResult struct {
	Amount        string `json:"amount"`
	AmountLocal   string `json:"amount_local"`
	Currency      string `json:"currency"`
	CurrencyLocal string `json:"currency_local"`
	EstimatedDays int    `json:"estimated_days"`
	Title         string `json:"title"`
}

type RateResponse struct {
	Rates []RateResult `json:"results"`
}

var (
	GenerateRateUri string = BaseUri + "/live-rates"
)

func (c *Client) GenerateRates(request RateRequest) (*RateResponse, error) {
	response, err := helper.Post(GenerateRateUri, BasicAuth, c.ApiKey, request)
	if err != nil {
		return nil, err
	}
	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var rate RateResponse
	err = json.NewDecoder(response.Body).Decode(&rate)
	if err != nil {
		return nil, err
	}

	return &rate, nil
}

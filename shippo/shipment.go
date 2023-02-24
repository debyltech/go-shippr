package shippo

import (
	"encoding/json"

	helper "github.com/debyltech/go-helpers"
)

type Shipment struct {
	AddressFrom Address  `json:"address_from"`
	AddressTo   Address  `json:"address_to"`
	Parcels     []Parcel `json:"parcels"`
}

type ShipmentRequest struct {
	CarrierAccount    string   `json:"carrier_account"`
	ServiceLevelToken string   `json:"servicelevel_token"`
	Shipment          Shipment `json:"shipment"`
}

type ShipmentRate struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	AmountLocal   string `json:"amount_local"`
	CurrencyLocal string `json:"currency_local"`
	Provider      string `json:"provider"`
}

type ShipmentResponse struct {
	Status         string       `json:"status"`
	Rate           ShipmentRate `json:"rate"`
	TrackingNumber string       `json:"tracking_number"`
	TrackingUrl    string       `json:"tracking_url_provider"`
	LabelUrl       string       `json:"label_url"`
}

var (
	ShipmentUri string = BaseUri + "/transactions"
)

func (c *Client) CreateShipment(req ShipmentRequest) (*ShipmentResponse, error) {
	response, err := helper.Post(ShipmentUri, BasicAuth, c.ApiKey, req)
	if err != nil {
		return nil, err
	}
	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var res ShipmentResponse
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
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

type CreateShipmentRequest struct {
	CarrierAccount    string   `json:"carrier_account"`
	ServiceLevelToken string   `json:"servicelevel_token"`
	Shipment          Shipment `json:"shipment"`
	Metadata          string   `json:"metadata,omitempty"`
}

type ShipmentRate struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	AmountLocal   string `json:"amount_local"`
	CurrencyLocal string `json:"currency_local"`
	Provider      string `json:"provider"`
}

type CreateShipmentResponse struct {
	Status         string       `json:"status"`
	Rate           ShipmentRate `json:"rate"`
	TrackingNumber string       `json:"tracking_number"`
	TrackingUrl    string       `json:"tracking_url_provider"`
	LabelUrl       string       `json:"label_url"`
}

type TransactionResponse struct {
	Id             string `json:"object_id"`
	State          string `json:"object_state"`
	Status         string `json:"status"`
	TrackingNumber string `json:"tracking_number"`
	TrackingUrl    string `json:"tracking_url_provider"`
	LabelUrl       string `json:"label_url"`
	Metadata       string `json:"metadata,omitempty"`
}

type TransactionsResponse struct {
	Transactions []TransactionResponse `json:"results"`
}

var (
	TransactionUri string = BaseUri + "/transactions"
	ShipmentsUri   string = BaseUri + "/shipments"
)

func (c *Client) CreateShipment(req CreateShipmentRequest) (*CreateShipmentResponse, error) {
	response, err := helper.Post(TransactionUri, BasicAuth, c.ApiKey, req)
	if err != nil {
		return nil, err
	}
	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var res CreateShipmentResponse
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ListShipments() (*TransactionsResponse, error) {
	response, err := helper.Get(TransactionUri, BasicAuth, c.ApiKey, map[string]string{"status": "SUCCESS"})
	if err != nil {
		return nil, err
	}
	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var res TransactionsResponse
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

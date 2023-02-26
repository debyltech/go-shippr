package shippo

import (
	"encoding/json"
	"time"

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
	LabelFileType     string   `json:"label_file_type,omitempty"`
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
	Id             string    `json:"object_id"`
	Created        time.Time `json:"object_created,omitempty"`
	Eta            time.Time `json:"eta,omitempty"`
	State          string    `json:"object_state"`
	Status         string    `json:"status"`
	TrackingNumber string    `json:"tracking_number"`
	TrackingUrl    string    `json:"tracking_url_provider"`
	LabelUrl       string    `json:"label_url"`
	Parcel         string    `json:"parcel"`
	Metadata       string    `json:"metadata,omitempty"`
}

type TransactionsResponse struct {
	Transactions []TransactionResponse `json:"results"`
}

var (
	TransactionUri string = BaseUri + "/transactions"
	ShipmentsUri   string = BaseUri + "/shipments"
	ParcelsUri     string = BaseUri + "/parcels"
)

func (c *Client) GetParcel(id string) (*Parcel, error) {
	response, err := helper.Get(ParcelsUri+"/"+id, BasicAuth, c.ApiKey, nil)
	if err != nil {
		return nil, err
	}
	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var parcel Parcel
	err = json.NewDecoder(response.Body).Decode(&parcel)
	if err != nil {
		return nil, err
	}

	return &parcel, nil
}

func (c *Client) CreateShipment(request CreateShipmentRequest) (*CreateShipmentResponse, error) {
	if request.LabelFileType == "" {
		request.LabelFileType = "PNG"
	}

	response, err := helper.Post(TransactionUri, BasicAuth, c.ApiKey, request)
	if err != nil {
		return nil, err
	}
	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var shipment CreateShipmentResponse
	err = json.NewDecoder(response.Body).Decode(&shipment)
	if err != nil {
		return nil, err
	}

	return &shipment, nil
}

func (c *Client) ListShipments() (*TransactionsResponse, error) {
	response, err := helper.Get(TransactionUri, BasicAuth, c.ApiKey, nil)
	if err != nil {
		return nil, err
	}
	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var transactions TransactionsResponse
	err = json.NewDecoder(response.Body).Decode(&transactions)
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (c *Client) GetShipment(id string) (*TransactionResponse, error) {
	response, err := helper.Get(TransactionUri+"/"+id, BasicAuth, c.ApiKey, nil)
	if err != nil {
		return nil, err
	}
	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var transaction TransactionResponse
	err = json.NewDecoder(response.Body).Decode(&transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

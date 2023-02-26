package shippo

import (
	"encoding/base64"
	"encoding/json"
	"time"

	helper "github.com/debyltech/go-helpers/json"
	"github.com/skip2/go-qrcode"
)

type CreateTransactionRequest struct {
	CarrierAccount    string   `json:"carrier_account"`
	ServiceLevelToken string   `json:"servicelevel_token"`
	Shipment          Shipment `json:"shipment"`
	LabelFileType     string   `json:"label_file_type,omitempty"`
	Metadata          string   `json:"metadata,omitempty"`
}

type TransactionRate struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	AmountLocal   string `json:"amount_local"`
	CurrencyLocal string `json:"currency_local"`
	Provider      string `json:"provider"`
}

type CreateTransactionResponse struct {
	Status         string          `json:"status"`
	Rate           TransactionRate `json:"rate"`
	TrackingNumber string          `json:"tracking_number"`
	TrackingUrl    string          `json:"tracking_url_provider"`
	LabelUrl       string          `json:"label_url"`
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
	ParcelsUri     string = BaseUri + "/parcels"
)

func (t *TransactionResponse) TransactionPNGBase64() (string, error) {
	img, err := qrcode.Encode("transaction:"+t.Id, qrcode.Medium, 128)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(img), nil
}

func (c *Client) CreateTransaction(request CreateTransactionRequest) (*CreateTransactionResponse, error) {
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

	var shipment CreateTransactionResponse
	err = json.NewDecoder(response.Body).Decode(&shipment)
	if err != nil {
		return nil, err
	}

	return &shipment, nil
}

func (c *Client) ListTransactions() (*TransactionsResponse, error) {
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

func (c *Client) GetTransaction(id string) (*TransactionResponse, error) {
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

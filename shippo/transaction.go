package shippo

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	helper "github.com/debyltech/go-helpers/json"
	"github.com/skip2/go-qrcode"
)

const (
	TransactionUri string = BaseUri + "/transactions"
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
	Id             string          `json:"object_id"`
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
	Messages       []any     `json:"messages,omitempty"`
}

type TransactionsResponse struct {
	Transactions []TransactionResponse `json:"results"`
}

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

func (c *Client) CreateLabelWithRateId(rateId string, labelFileType string) (*TransactionResponse, error) {
	data := url.Values{}
	data.Set("rate", rateId)
	data.Set("label_file_type", labelFileType)
	data.Set("async", "false")

	request, err := http.NewRequest("POST", TransactionUri, bytes.NewBuffer([]byte(data.Encode())))
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

	var transaction TransactionResponse
	err = json.NewDecoder(response.Body).Decode(&transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

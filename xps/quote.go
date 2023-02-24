package xps

import (
	"encoding/json"
	"fmt"

	helper "github.com/debyltech/go-helpers"
)

type QuoteSurcharge struct {
	Description string `json:"description"`
	Amount      string `json:"amount"`
}

type QuoteResponse struct {
	Currency        string           `json:"currency"`
	CustomsCurrency string           `json:"customsCurrency"`
	TotalAmount     string           `json:"totalAmount"`
	BaseAmount      string           `json:"baseAmount"`
	Surcharges      []QuoteSurcharge `json:"surcharges"`
	Zone            string           `json:"zone"`
}

type QuoteAddressSender struct {
	Country string `json:"country"`
	ZIP     string `json:"zip"`
}

type QuoteAddressReceiver struct {
	City    string `json:"city"`
	Country string `json:"country"`
	ZIP     string `json:"zip"`
	Email   string `json:"email"`
}

type QuotingOption struct {
	Name            string `json:"name"`
	CarrierCode     string `json:"carrierCode"`
	ServiceCode     string `json:"serviceCode"`
	PackageTypeCode string `json:"packageTypeCode"`
}

type QuotingOptions struct {
	Options []QuotingOption `json:"integratedQuotingOptions"`
}

const (
	QuoteURIFmt              string = "/restapi/v1/customers/%s/quote"
	QuotingOptionsURIFmt     string = "/restapi/v1/customers/%s/integratedQuotingOptions"
	QuotingOptionNotFoundFmt string = "quoting option not found for carrierCode:'%s', serviceCode:'%s', packageTypeCode:'%s'"
)

func (c *Client) getBillingEndpoint() string {
	return BaseUri + fmt.Sprintf(QuoteURIFmt, c.CustomerId)
}

func (c *Client) getQuotingOptionsEndpoint() string {
	return BaseUri + fmt.Sprintf(QuotingOptionsURIFmt, c.CustomerId)
}

func (c *Client) GetQuotingOptions() (*QuotingOptions, error) {
	response, err := helper.Get(c.getQuotingOptionsEndpoint(), BasicAuth, c.ApiKey, nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var options QuotingOptions
	err = json.NewDecoder(response.Body).Decode(&options)
	if err != nil {
		return nil, err
	}

	return &options, nil
}

func (c *Client) HasQuotingOption(request ShipmentRequest) error {
	options, err := c.GetQuotingOptions()
	if err != nil {
		return err
	}

	for _, v := range options.Options {
		if v.CarrierCode == request.CarrierCode && v.ServiceCode == request.ServiceCode && v.PackageTypeCode == request.PackageTypeCode {
			return nil
		}
	}

	return fmt.Errorf(QuotingOptionNotFoundFmt, request.CarrierCode, request.ServiceCode, request.PackageTypeCode)
}

func (c *Client) GetQuote(request ShipmentRequest) (*QuoteResponse, error) {
	err := c.HasQuotingOption(request)
	if err != nil {
		return nil, err
	}

	response, err := helper.Post(c.getBillingEndpoint(), BasicAuth, c.ApiKey, request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var quote QuoteResponse
	err = json.NewDecoder(response.Body).Decode(&quote)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}

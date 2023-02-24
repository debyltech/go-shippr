package xps

import (
	"encoding/json"
	"fmt"
	"net/http"

	helper "github.com/debyltech/go-helpers"
)

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

type QuoteBillingParty string

type QuoteBilling struct {
	Party QuoteBillingParty `json:"party"`
}

type QuoteRequest struct {
	CarrierCode         string               `json:"carrierCode"`
	ServiceCode         string               `json:"serviceCode"`
	PackageTypeCode     string               `json:"packageTypeCode"`
	Sender              QuoteAddressSender   `json:"sender"`
	Receiver            QuoteAddressReceiver `json:"receiver"`
	Residential         bool                 `json:"residential"`
	SignatureOptionCode *string              `json:"signatureOptionCode"`
	ContentDescription  string               `json:"contentDescription"`
	WeightUnit          string               `json:"weightUnit"`
	DimensionUnit       string               `json:"dimUnit"`
	Currency            Currency             `json:"currency"`
	CustomsCurrency     Currency             `json:"customsCurrency"`
	Pieces              []Piece              `json:"pieces"`
	Billing             QuoteBilling         `json:"billing"`
}

type QuoteOption struct {
	Name            string `json:"name"`
	CarrierCode     string `json:"carrierCode"`
	ServiceCode     string `json:"serviceCode"`
	PackageTypeCode string `json:"packageTypeCode"`
}

type QuoteOptions struct {
	Options []QuoteOption `json:"integratedQuotingOptions"`
}

const (
	QuoteBillingPartySender   QuoteBillingParty = "sender"
	QuoteBillingPartyReceiver QuoteBillingParty = "receiver"
	QuoteURIFmt               string            = "/restapi/v1/customers/%s/quote"
	QuotingOptionsURIFmt      string            = "/restapi/v1/customers/%s/integratedQuotingOptions"

	QuotingOptionNotFoundFmt string = "quoting option not found for carrierCode:'%s', serviceCode:'%s', packageTypeCode:'%s'"
)

func (c *Client) GetBillingEndpoint() string {
	return BaseURI + fmt.Sprintf(QuoteURIFmt, c.CustomerId)
}

func (c *Client) GetQuotingOptionsEndpoint() string {
	return BaseURI + fmt.Sprintf(QuotingOptionsURIFmt, c.CustomerId)
}

func (c *Client) GetQuotingOptions() (*QuoteOptions, error) {
	response, err := helper.Get(c.GetQuotingOptionsEndpoint(), BasicAuth, c.ApiKey, nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var options QuoteOptions
	err = json.NewDecoder(response.Body).Decode(&options)
	if err != nil {
		return nil, err
	}

	return &options, nil
}

func (c *Client) HasQuotingOption(request QuoteRequest) error {
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

func (c *Client) GetQuote(request QuoteRequest) (*http.Response, error) {
	err := c.HasQuotingOption(request)
	if err != nil {
		return nil, err
	}

	response, err := helper.Post(c.GetBillingEndpoint(), BasicAuth, c.ApiKey, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

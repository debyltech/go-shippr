package xps

import (
	"encoding/json"
	"fmt"
	"io"

	helper "github.com/debyltech/go-helpers"
)

type Address struct {
	Name     string `json:"name"`
	Company  string `json:"company"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	ZIP      string `json:"zip"`
	Country  string `json:"country"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type ShipmentResponse struct {
	BookNumber     string `json:"bookNumber"`
	TrackingNumber string `json:"trackingNumber"`
	PrepayBalance  string `json:"prepayBalance"`
	Zone           string `json:"zone"`
}

const (
	CreateShipmentUriFmt   string = "/restapi/v1/customers/%s/shipments"
	GetShipmentLabelUriFmt string = "/restapi/v1/customers/%s/shipments/%s/label"
)

func (c *Client) getShipmentEndpoint() string {
	return BaseUri + fmt.Sprintf(CreateShipmentUriFmt, c.CustomerId)
}

func (c *Client) getShipmentLabelEndpoint(bookNumber string) string {
	return BaseUri + fmt.Sprintf(GetShipmentLabelUriFmt, c.CustomerId, bookNumber)
}

func (c *Client) CreateShipmentDomestic(request ShipmentRequest) (*ShipmentResponse, error) {
	err := c.HasQuotingOption(request)
	if err != nil {
		return nil, err
	}

	response, err := helper.Post(c.getShipmentEndpoint(), BasicAuth, c.ApiKey, request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	var shipment ShipmentResponse
	err = json.NewDecoder(response.Body).Decode(&shipment)
	if err != nil {
		return nil, err
	}

	return &shipment, nil
}

func (c *Client) GetShipmentLabel(bookNumber string) ([]byte, error) {
	response, err := helper.Get(c.getShipmentLabelEndpoint(bookNumber), BasicAuth, c.ApiKey, nil)
	if err != nil {
		return nil, err
	}

	if response.Header.Get("Content-Type") != "application/pdf" {
		return nil, fmt.Errorf("expected 'application/pdf' response, got '%s'", response.Header.Get("Content-Type"))
	}

	pdfBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return pdfBytes, nil
}

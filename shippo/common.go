package shippo

import (
	"fmt"
	"io"
	"net/http"
)

type Address struct {
	Id         string `json:"object_id,omitempty"`
	Name       string `json:"name"`
	Company    string `json:"company"`
	Street     string `json:"street_no"`
	Address1   string `json:"street1"`
	Address2   string `json:"street2"`
	Address3   string `json:"street3"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"zip"`
	Country    string `json:"country"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
}

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

type ServiceLevel struct {
	Name               string `json:"name"`
	Token              string `json:"token"`
	Terms              string `json:"terms"`
	ExtendedToken      string `json:"extended_token"`
	ParentServiceLevel string `json:"parent_servicelevel,omitempty"`
}

const (
	BaseUri   string = "https://api.goshippo.com"
	BasicAuth string = "ShippoToken"
)

func HandleResponseStatus(response *http.Response) error {
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return nil
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return fmt.Errorf("status: %s error: %s", response.Status, string(bodyBytes))
}

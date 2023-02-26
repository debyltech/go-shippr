package shippo

import (
	"fmt"
	"io"
	"net/http"
)

type Address struct {
	Name       string `json:"name"`
	Company    string `json:"company"`
	Address1   string `json:"street1"`
	Address2   string `json:"street_no"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"zip"`
	Country    string `json:"country"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
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

package shippo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
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

type AddressValidationResultMessage struct {
	Source string `json:"source"`
	Code   string `json:"code"`
	Type   string `json:"type"`
	Text   string `json:"text"`
}

type AddressValidationResult struct {
	IsValid  bool                             `json:"is_valid"`
	Messages []AddressValidationResultMessage `json:"messages"`
}

type AddressValidationResponse struct {
	Created          time.Time               `json:"object_created"`
	IsComplete       bool                    `json:"is_complete"`
	ValidationResult AddressValidationResult `json:"validation_results"`
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
	BaseUri              string = "https://api.goshippo.com"
	BasicAuth            string = "ShippoToken"
	AddressValidationUri string = BaseUri + "/addresses"
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

func (c *Client) ValidateAddress(address Address) (bool, error) {
	data := url.Values{}
	data.Set("name", address.Name)
	data.Set("street1", address.Address1)
	data.Set("city", address.City)
	data.Set("state", address.State)
	data.Set("zip", address.PostalCode)
	data.Set("country", address.Country)
	data.Set("email", address.Email)
	data.Set("validate", "true")

	request, err := http.NewRequest("POST", AddressValidationUri, bytes.NewBuffer([]byte(data.Encode())))
	if err != nil {
		return false, err
	}

	client := &http.Client{}

	request.Header.Set("Authorization", fmt.Sprintf("%s %s", BasicAuth, c.ApiKey))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	var responseValidation AddressValidationResponse
	if err := json.NewDecoder(response.Body).Decode(&responseValidation); err != nil {
		return false, err
	}

	return responseValidation.ValidationResult.IsValid, nil
}

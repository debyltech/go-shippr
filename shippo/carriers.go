package shippo

import (
	"encoding/json"

	helper "github.com/debyltech/go-helpers"
)

type ServiceLevel struct {
	Token               string `json:"token"`
	Name                string `json:"name"`
	SupportReturnLabels bool   `json:"supports_return_labels"`
}

type Account struct {
	CarrierId     string            `json:"object_id"`
	Carrier       string            `json:"carrier"`
	CarrierName   string            `json:"carrier_name"`
	ServiceLevels []ServiceLevel    `json:"service_levels"`
	AccountId     string            `json:"account_id"`
	Parameters    map[string]string `json:"parameters"`
	Active        bool              `json:"active"`
}

type CarrierResponse struct {
	Results []Account `json:"results"`
}

var (
	CarrierAccountsUri string = BaseUri + "/carrier_accounts"
)

func (c *Client) ListCarrierAccounts(carrier string) (*CarrierResponse, error) {
	urlQueries := make(map[string]string)
	urlQueries["service_levels"] = "true"
	if carrier != "" {
		urlQueries["carrier"] = carrier
	}
	response, err := helper.Get(CarrierAccountsUri, BasicAuth, c.ApiKey, urlQueries)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var carriers CarrierResponse
	err = json.NewDecoder(response.Body).Decode(&carriers)
	if err != nil {
		return nil, err
	}

	return &carriers, nil
}

func (c *Client) ListAllCarrierAccounts() (*CarrierResponse, error) {
	return c.ListCarrierAccounts("")
}

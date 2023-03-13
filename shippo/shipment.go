package shippo

import (
	"encoding/base64"
	"encoding/json"

	helper "github.com/debyltech/go-helpers/json"
	"github.com/skip2/go-qrcode"
)

const (
	ShipmentsUri string = BaseUri + "/shipments"
)

type Shipment struct {
	Id          string   `json:"object_id,omitempty"`
	AddressFrom Address  `json:"address_from"`
	AddressTo   Address  `json:"address_to"`
	Parcels     []Parcel `json:"parcels"`
	Metadata    string   `json:"metadata,omitempty"`
}

func (s *Shipment) ShipmentPNGBase64() (string, error) {
	img, err := qrcode.Encode("shipment:"+s.Id, qrcode.Medium, 128)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(img), nil
}

func (c *Client) CreateShipment(request Shipment) (*Shipment, error) {
	response, err := helper.Post(ShipmentsUri, BasicAuth, c.ApiKey, request)
	if err != nil {
		return nil, err
	}
	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var shipment Shipment
	err = json.NewDecoder(response.Body).Decode(&shipment)
	if err != nil {
		return nil, err
	}

	return &shipment, nil
}

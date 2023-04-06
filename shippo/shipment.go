package shippo

import (
	"encoding/base64"
	"encoding/json"
	"time"

	helper "github.com/debyltech/go-helpers/json"
	"github.com/skip2/go-qrcode"
)

const (
	ShipmentsUri string = BaseUri + "/shipments"
)

type Shipment struct {
	Id                 string    `json:"object_id,omitempty"`
	Created            time.Time `json:"object_created,omitempty"`
	AddressFrom        Address   `json:"address_from"`
	AddressTo          Address   `json:"address_to"`
	Parcels            []Parcel  `json:"parcels"`
	CustomsDeclaration any       `json:"customs_declaration,omitempty"`
	Metadata           string    `json:"metadata,omitempty"`
	Rates              []Rate    `json:"rates,omitempty"`
	Messages           []any     `json:"messages"`
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

func (c *Client) GetShipmentById(id string) (*Shipment, error) {
	var shipment Shipment
	response, err := helper.Get(ShipmentsUri+"/"+id, BasicAuth, c.ApiKey, nil)
	if err != nil {
		return nil, err
	}

	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&shipment)
	if err != nil {
		return nil, err
	}

	return &shipment, nil
}

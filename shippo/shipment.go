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
	Id                 string              `json:"object_id,omitempty"`
	Status             string              `json:"status"`
	Created            time.Time           `json:"object_created,omitempty"`
	AddressFrom        Address             `json:"address_from"`
	AddressTo          Address             `json:"address_to"`
	Parcels            []Parcel            `json:"parcels"`
	CustomsDeclaration *CustomsDeclaration `json:"customs_declaration,omitempty"`
	Metadata           string              `json:"metadata,omitempty"`
	Rates              []Rate              `json:"rates,omitempty"`
	Messages           []any               `json:"messages"`
}

type ListShipmentsResponse struct {
	Next      *string    `json:"next"`
	Previous  *string    `json:"previous"`
	Shipments []Shipment `json:"results"`
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

func (c *Client) ListShipments(query map[string]string) (*ListShipmentsResponse, error) {
	var shipments ListShipmentsResponse
	response, err := helper.Get(ShipmentsUri, BasicAuth, c.ApiKey, query)
	if err != nil {
		return nil, err
	}

	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&shipments)
	if err != nil {
		return nil, err
	}

	return &shipments, nil
}

func (c *Client) AwaitQueuedFinished(id string) error {
	var err error
	var shipment *Shipment = &Shipment{Status: "QUEUED"}

	for shipment.Status == "QUEUED" {
		shipment, err = c.GetShipmentById(id)
		if err != nil {
			return err
		}
	}

	return nil
}

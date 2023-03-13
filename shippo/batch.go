package shippo

import "time"

type Batch struct {
	Id          string  `json:"object_id"`
	Shipment    *string `json:"shipment,omitempty"`
	Transaction *string `json:"transaction,omitempty"`
	Metadata    string  `json:"metadata"`
}

type BatchShipments struct {
	Next      *string `json:"next"`
	Shipments []Batch `json:"results"`
}

type CreateBatchRequest struct {
	Id                  string    `json:"object_id"`
	Created             time.Time `json:"object_created,omitempty"`
	Updated             time.Time `json:"object_updated,omitempty"`
	DefaultCarrier      string    `json:"default_carrier_account"`
	DefaultServiceLevel string    `json:"default_servicelevel_token"`
	LabelFileType       string    `json:"label_filetype"`
}

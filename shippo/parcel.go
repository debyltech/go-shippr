package shippo

import (
	"encoding/json"

	helper "github.com/debyltech/go-helpers/json"
)

type Parcel struct {
	Id           string `json:"object_id,omitempty"`
	Length       string `json:"length"`
	Width        string `json:"width"`
	Height       string `json:"height"`
	DistanceUnit string `json:"distance_unit"`
	Weight       string `json:"weight"`
	WeightUnit   string `json:"mass_unit"`
	LineItems    []any  `json:"line_items,omitempty"`
}

func (c *Client) GetParcel(id string) (*Parcel, error) {
	response, err := helper.Get(ParcelsUri+"/"+id, BasicAuth, c.ApiKey, nil)
	if err != nil {
		return nil, err
	}
	err = HandleResponseStatus(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var parcel Parcel
	err = json.NewDecoder(response.Body).Decode(&parcel)
	if err != nil {
		return nil, err
	}

	return &parcel, nil
}

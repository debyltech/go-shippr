package xps

// REDO THIS

/*
type Address struct {
	Name       string `json:"name"`
	Company    string `json:"company"`
	Address1   string `json:"street1"`
	Address2   string `json:"street_no"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"zip"`
	Country    string `json:"country"`
}

type Parcel struct {
	Length       string `json:"length"`
	Width        string `json:"width"`
	Height       string `json:"height"`
	DistanceUnit string `json:"distance_unit"`
	Weight       string `json:"weight"`
	MassUnit     string `json:"mass_unit"`
}

const (
	BaseUri   string = "https://api.goshippo.com"
	BasicAuth string = "ShippoToken"
)

func HandleResponseStatus(response *http.Response) error {
	if response.Status == "200 OK" {
		return nil
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return fmt.Errorf("status: %s error: %s", response.Status, string(bodyBytes))
}
*/

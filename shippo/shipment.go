package shippo

var (
	ShipmentsUri string = BaseUri + "/shipments"
)

type Shipment struct {
	AddressFrom Address  `json:"address_from"`
	AddressTo   Address  `json:"address_to"`
	Parcels     []Parcel `json:"parcels"`
}

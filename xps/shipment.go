package xps

type ShipmentBillingParty string

type ShipmentBillingQuote struct {
	Party ShipmentBillingParty `json:"party"`
}

type ShipmentBillingFull struct {
	Party   ShipmentBillingParty `json:"party"`
	Account *string              `json:"account"`
	Country *string              `json:"country"`
	ZIP     *string              `json:"zip"`
}

type ShipmentRequest struct {
	CarrierCode         string      `json:"carrierCode"`
	ServiceCode         string      `json:"serviceCode"`
	PackageTypeCode     string      `json:"packageTypeCode"`
	Sender              interface{} `json:"sender"`   // Expected types are QuoteAddressSender, or Address
	Receiver            interface{} `json:"receiver"` // Expected types are QuoteAddressReceiver, or Address
	Residential         bool        `json:"residential"`
	SignatureOptionCode *string     `json:"signatureOptionCode"` // Can be nil if signature not requested
	ContentDescription  string      `json:"contentDescription"`
	WeightUnit          string      `json:"weightUnit"`
	DimensionUnit       string      `json:"dimUnit"`
	Currency            Currency    `json:"currency"`
	CustomsCurrency     Currency    `json:"customsCurrency"`
	Pieces              []Piece     `json:"pieces"`
	LabelImageFormat    *string     `json:"labelImageFormat,omitempty"`
	Billing             interface{} `json:"billing"` // Expected types are QuoteBilling, or ShipmentBilling
}

const (
	ShipmentBillingPartySender   ShipmentBillingParty = "sender"
	ShipmentBillingPartyReceiver ShipmentBillingParty = "receiver"
)

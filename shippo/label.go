package shippo

type Label struct {
	SnipcartOrderToken  string `json:"order_token"`
	ShippoTransactionId string `json:"transaction_id"`
	TrackingNumber      string `json:"tracking_number"`
	TrackingUrl         string `json:"tracking_url_provider"`
}

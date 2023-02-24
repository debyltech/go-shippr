package xps

type Client struct {
	ApiKey     string
	CustomerId string
}

func NewClient(apiKey string, customerId string) Client {
	return Client{
		ApiKey:     apiKey,
		CustomerId: customerId,
	}
}

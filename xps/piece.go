package xps

type Piece struct {
	Weight          string  `json:"weight"`
	Length          string  `json:"length"`
	Width           string  `json:"width"`
	Height          string  `json:"height"`
	InsuranceAmount string  `json:"insuranceAmount"`
	DeclaredValue   *string `json:"declaredValue"`
}

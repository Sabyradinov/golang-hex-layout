package model

// ShipmentRequest is a struct that represents the request for delivery
type ShipmentRequest struct {
	Id                 int64   `json:"id"`
	ProductId          int64   `json:"productId"`
	Address            string  `json:"address"`
	From               string  `json:"from"`
	To                 string  `json:"to"`
	Amount             float64 `json:"amount"`
	Currency           string  `json:"currency"`
	Commission         float64 `json:"commission"`
	CommissionCurrency string  `json:"commissionCurrency"`
}

package entity

// Shipment structure to describe shipment table
type Shipment struct {
	Id                 int64   `json:"id" gorm:"column:id"`
	ProductId          int64   `json:"productId" gorm:"column:product_id"`
	Address            string  `json:"address" gorm:"column:address"`
	From               string  `json:"from" gorm:"column:from"`
	To                 string  `json:"to" gorm:"column:to"`
	Amount             float64 `json:"amount" gorm:"column:amount"`
	Currency           string  `json:"currency" gorm:"column:currency"`
	Commission         float64 `json:"commission" gorm:"column:commission"`
	CommissionCurrency string  `json:"commissionCurrency" gorm:"column:commission_currency"`
}

func (Shipment) TableName() string {
	return "public.Shipment"
}

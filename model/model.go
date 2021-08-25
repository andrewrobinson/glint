package model

import "time"

type Customer struct {
	FirstName    string
	LastName     string
	Email        string
	Description  string
	Amount       float64
	AmountGBP    float64
	FromCurrency string
	ToCurrency   string
	Rate         float64
	Date         time.Time
}

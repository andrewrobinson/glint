package model

import (
	"log"
	"strconv"
	"time"

	"github.com/luno/luno-go/decimal"
)

type Customer struct {
	FirstName    string
	LastName     string
	Email        string
	Description  string
	Amount       decimal.Decimal
	AmountGBP    decimal.Decimal
	FromCurrency string
	ToCurrency   string
	Rate         float64
	Date         time.Time
}

func BuildCustomerFromCsvRow(line []string) Customer {

	amountString := line[5]
	rateString := line[8]
	dateString := line[9]

	rate, err := strconv.ParseFloat(rateString, 64)

	if err != nil {
		log.Fatal(err)
	}

	amount, err := decimal.NewFromString(amountString)

	if err != nil {
		log.Fatal(err)
	}

	date, err := time.Parse("02/01/2006 15:04", dateString)

	if err != nil {
		log.Fatal(err)
	}

	return Customer{
		FirstName:    line[0],
		LastName:     line[1],
		Email:        line[2],
		Description:  line[3],
		Amount:       amount,
		AmountGBP:    amount.MulInt64(rate),
		FromCurrency: line[6],
		ToCurrency:   line[7],
		Rate:         rate,
		Date:         date,
	}

}

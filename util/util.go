package util

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/andrewrobinson/glint/model"
)

func BuildCustomer(line []string) model.Customer {

	amountString := line[5]
	rateString := line[8]
	dateString := line[9]

	rate, err := strconv.ParseFloat(rateString, 64)

	if err != nil {
		log.Fatal(err)
	}

	amount, err := strconv.ParseFloat(amountString, 64)

	if err != nil {
		log.Fatal(err)
	}

	date, err := time.Parse("02/01/2006 15:04", dateString)

	if err != nil {
		log.Fatal(err)
	}

	return model.Customer{
		FirstName:    line[0],
		LastName:     line[1],
		Email:        line[2],
		Description:  line[3],
		Amount:       amount,
		AmountGBP:    amount * rate,
		FromCurrency: line[6],
		ToCurrency:   line[7],
		Rate:         rate,
		Date:         date,
	}

}

func DateInAugust2020(date time.Time) bool {

	start, _ := time.Parse(time.RFC822, "01 Aug 2020 10:00 UTC")
	end, _ := time.Parse(time.RFC822, "31 Aug 2020 23:59 UTC")

	if inTimeSpan(start, end, date) {
		fmt.Printf("date:%+v is in August 2020\n", date)
	}
	// } else {
	// 	fmt.Printf("date:%+v is NOT in August 2020\n", date)
	// }

	return inTimeSpan(start, end, date)
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func PrintTopSpends(rows []model.Customer) {
	for _, row := range rows {
		fmt.Printf("email:%s, amountGBP:%.2f\n", row.Email, row.AmountGBP)
	}

}

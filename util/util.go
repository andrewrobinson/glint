package util

import (
	"fmt"
	"time"

	"github.com/andrewrobinson/glint/model"
)

func GetCardSpendsInAugust2020(data [][]string) []model.Customer {

	rows := make([]model.Customer, 0)

	for i, line := range data {

		description := line[3]

		// skip the 1st row of the csv
		// filter by description CARD SPEND
		if i == 0 || description != "CARD SPEND" {
			continue
		}

		// and normalise amount to GBP
		customer := model.BuildCustomerFromCsvRow(line)

		// filter by date in Aug 2020
		if dateInAugust2020(customer.Date) {
			rows = append(rows, customer)
		}

	}
	return rows
}

func GetTopSpends(rows []model.Customer, limit int) []model.Customer {

	if len(rows) >= limit {
		return rows[0:limit]
	} else {
		return rows
	}

}

func dateInAugust2020(date time.Time) bool {

	start, _ := time.Parse(time.RFC3339, "2020-08-01T00:00:00Z")
	end, _ := time.Parse(time.RFC3339, "2020-08-31T23:59:59Z")

	return inTimeSpan(start, end, date)
}

func inTimeSpan(start, end, check time.Time) bool {
	//this I got off the internet
	// return check.After(start) && check.Before(end)

	//but my unit tests made me change it to this
	return (check.After(start) && check.Before(end)) || check.Equal(start) || check.Equal(end)
}

func PrintTopSpends(rows []model.Customer) {
	for _, row := range rows {
		fmt.Printf("email:%s, amountGBP:%.2f\n", row.Email, row.AmountGBP)
	}

}

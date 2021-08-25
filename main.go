package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/andrewrobinson/glint/model"
	"github.com/andrewrobinson/glint/util"
)

func main() {

	f, err := os.Open("sample-transactions.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// CARD SPENDs in August 2020
	rows := getMatchingRows(data)

	// sort by amountGBP
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].AmountGBP > rows[j].AmountGBP
	})

	// return top 5 spends
	topSpends := getTopSpends(rows, 5)

	fmt.Println("top 5 Spends")
	util.PrintTopSpends(topSpends)

}

func getMatchingRows(data [][]string) []model.Customer {

	rows := make([]model.Customer, 0)

	for i, line := range data {

		description := line[3]

		// skip the 1st row of the csv
		// filter by description CARD SPEND
		if i == 0 || description != "CARD SPEND" {
			continue
		}

		// and normalise amount to GBP
		customer := util.BuildCustomer(line)

		// filter by date in Aug 2020
		if util.DateInAugust2020(customer.Date) {
			rows = append(rows, customer)
		}

	}
	return rows
}

func getTopSpends(rows []model.Customer, limit int) []model.Customer {

	if len(rows) >= limit {
		return rows[0:limit]
	} else {
		return rows
	}

}

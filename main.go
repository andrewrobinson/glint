package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type Customer struct {
	firstName    string
	lastName     string
	email        string
	description  string
	amount       float64
	amountGBP    float64
	fromCurrency string
	toCurrency   string
	rate         float64
	date         time.Time
}

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
		return rows[i].amountGBP > rows[j].amountGBP
	})

	// return top 5 spends
	topSpends := getTopSpends(rows, 5)

	fmt.Println("top 5 Spends")
	printTopSpends(topSpends)

	// fmt.Printf("len rows:%d\n", len(rows))
	// fmt.Printf("len topSpends:%d\n", len(topSpends))

}

func printTopSpends(rows []Customer) {
	for _, row := range rows {
		fmt.Printf("email:%s, amountGBP:%.2f\n", row.email, row.amountGBP)
	}

}

func getMatchingRows(data [][]string) []Customer {

	rows := make([]Customer, 0)

	for i, line := range data {

		//TODO - just for debugging
		if i == 6 {
			return rows
		}

		description := line[3]

		// skip the 1st row of the csv
		// filter by description CARD SPEND
		if i == 0 || description != "CARD SPEND" {
			continue
		}

		// and normalise amount to GBP
		customer := buildCustomer(line)

		// filter by date in Aug 2020
		if dateInAugust2020(customer) {
			rows = append(rows, customer)
		}

	}
	return rows
}

func buildCustomer(line []string) Customer {

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

	return Customer{
		firstName:    line[0],
		lastName:     line[1],
		email:        line[2],
		description:  line[3],
		amount:       amount,
		amountGBP:    amount * rate,
		fromCurrency: line[6],
		toCurrency:   line[7],
		rate:         rate,
		date:         date,
	}

}

func dateInAugust2020(customer Customer) bool {
	//TODO - impl
	return true
}

func getTopSpends(rows []Customer, limit int) []Customer {

	if len(rows) >= limit {
		return rows[0:limit]
	} else {
		return rows
	}

}

// func tmp() {

// 	type Student struct {
// 		name  string
// 		score int
// 	}

// 	students := []Student{

// 		Student{name: "John", score: 45},
// 		Student{name: "Bill", score: 68},
// 		Student{name: "Sam", score: 98},
// 		Student{name: "Julia", score: 87},
// 		Student{name: "Tom", score: 91},
// 		Student{name: "Martin", score: 71},
// 	}

// 	sort.Slice(students, func(i, j int) bool {
// 		return students[i].score < students[j].score
// 	})

// 	fmt.Println(students)

// 	sort.Slice(students, func(i, j int) bool {
// 		return students[i].name > students[j].name
// 	})

// 	fmt.Println(students)

// }

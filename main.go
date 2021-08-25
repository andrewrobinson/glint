package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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

func buildCustomer(line []string) Customer {

	amountString := line[5]
	rateString := line[8]
	dateString := line[9]

	// fmt.Printf("rateString: %s\n", rateString)

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

	// fmt.Printf("date parsed:%s\n", date)
	// fmt.Printf("rate parsed:%f\n", rate)

	// First name,Last name,Email,Description,Merchant code,Amount,From Currency,To Currency,Rate,Date

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

	rows := make([]Customer, 0)

	for i, line := range data {

		description := line[3]

		if i == 0 || description != "CARD SPEND" {
			continue
		}

		if i == 6 {
			break
		}

		// 	//filter by description CARD SPEND
		// 	//filter by date = Aug 2020
		// 	//normalise amount to GBP
		// 	//sort by amount
		// 	//return top 5

		customer := buildCustomer(line)

		if dateInAugust2020(customer) {
			rows = append(rows, customer)
		}

	}

	fmt.Printf("%d\n", len(rows))

}

func dateInAugust2020(customer Customer) bool {
	return true
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

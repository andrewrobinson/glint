package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"

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

	rows := util.GetCardSpendsInAugust2020(data)

	// sort by amountGBP
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].AmountGBP > rows[j].AmountGBP
	})

	// return max 5 from the sorted rows
	topSpends := util.GetTopSpends(rows, 5)

	fmt.Println("top 5 Card Spends in August 2020")
	util.PrintTopSpends(topSpends)

}

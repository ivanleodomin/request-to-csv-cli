package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ivanleodomin/request-to-csv-cli/beers"
	"github.com/ivanleodomin/request-to-csv-cli/csv"
)

func main() {
	beers, err := beers.GetBeers()

	if err != nil {
		log.Fatal(err)
	}

	headers := []string{"Abv", "BrewersTips", "Description", "Ebc", "FirstBrewed", "FoodPairing"}

	var lines [][]string

	for _, b := range beers {

		lines = append(lines, []string{
			fmt.Sprint(b.Abv),
			b.BrewersTips,
			b.Description,
			fmt.Sprint(b.Ebc),
			b.FirstBrewed,
			fmt.Sprintf("[%s]", strings.Join(b.FoodPairing, "-")),
		})
	}

	csv.ParseToCsv("beers.csv", headers, lines)
}

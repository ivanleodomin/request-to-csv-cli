package main

import (
	"fmt"

	"github.com/ivanleodomin/request-to-csv-cli/beers"
)

func main() {
	beers, _ := beers.GetBeers()
	fmt.Println((*beers)[0].Description)
}

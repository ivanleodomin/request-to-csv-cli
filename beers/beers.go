package beers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const url = "https://api.punkapi.com/v2/beers?per_page=80"

type Beer struct {
	Abv         int      `json:"abv"`
	BrewersTips string   `json:"brewers_tips"`
	Description string   `json:"description"`
	Ebc         int      `json:"ebc"`
	FirstBrewed string   `json:"first_brewed"`
	FoodPairing []string `json:"food_pairing"`
}

func GetBeers() ([]Beer, error) {

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var beers []Beer
	json.Unmarshal(body, &beers)

	return beers, nil
}

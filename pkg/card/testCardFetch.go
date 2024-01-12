package card

import (
	"io"
	"net/http"
)

var API_URL = "https://db.ygoprodeck.com/api/v7/cardinfo.php"

// A test function for fetching a card from the API
func TestCardFetch(name string) {
	resp, err := http.Get(API_URL + "?name=" + name)
	if err != nil {
		panic("Error fetching card")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Error reading response body")
	}

	
	newCards := ParseCards(body)
	println(newCards[0].Name) // At this point, we have a card struct with the data from the API
}
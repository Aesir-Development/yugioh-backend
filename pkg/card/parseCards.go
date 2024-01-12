package card

import (
	"encoding/json"
)

type CardWrapper struct {
	Data []Card `json:"data"`
}

// ParseCard takes card JSON and parses it into a Card struct
func ParseCards(cardJSON []byte) []Card {

	println("Parsing card")

	var cards CardWrapper
	err := json.Unmarshal(cardJSON, &cards)
	if err != nil {
		panic(err)
	}

	return cards.Data
}
	
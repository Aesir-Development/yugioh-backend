package deck

import (
	"encoding/json"
)

// NOTE: This uses the YGOPRODeck deck builder, and parses the deck to a JSON string
// that can be used to save the deck to the DB


// Parse deck to JSON string
func ParseDeck(deck Deck) string {
	deckJSON, err := json.Marshal(deck)
	if err != nil {
		panic(err)
	}

	return string(deckJSON)
}

// TODO - Parse string to deck from YGOPRODeck deck builder
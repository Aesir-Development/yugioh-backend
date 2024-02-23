package dbConnection

// FIXME

/*
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	decks "github.com/Aesir-Development/yugioh-backend/pkg/deck"
)

type deck struct {
	ID int `json:"id"`
}

type deckCard struct {
	DeckID int `json:"deck_id"`
	CardID int `json:"card_id"`
	IsExtraDeck bool `json:"is_extra_deck"`
}

// Saving the deck to the DB
func SaveDeck(deck decks.Deck) (int, error) {
	res, err := DB.Exec("INSER INTO decks () VALUES ()")
	if err != nil {
		return fmt.Errorf("error inserting deck: %s", err)
	}	

	deckID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert ID: %s", err)
	}

	for _, card := range deck.Deck {
		_, err := DB.Exec("INSERT INTO deck_cards (deck_id, card_id, is_extra_deck) VALUES (?, ?, ?)", deckID, card.ID, false)
		if err != nil {
			return int{}, fmt.Errorf("error inserting deck card: %s", err)
		}

	}

	return nil
}
*/

// Parsing the deck to DB structs
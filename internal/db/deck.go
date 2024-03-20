package dbConnection

import (
	"github.com/Aesir-Development/yugioh-backend/internal/user"
	"github.com/Aesir-Development/yugioh-backend/pkg/card"
	"github.com/Aesir-Development/yugioh-backend/pkg/deck"

	"fmt"
)

// NOTE - Everything untested

// Update an existing deck
func UpdateDeck(deck deck.Deck, id int64) error {
	
	for _, card := range deck.Deck {
		// Main deck cards
		err := InsertCardToDB(card, id, "main_deck")
		if err != nil {
			return fmt.Errorf("error inserting card into main deck: %s", err)
		}
	}

	for _, card := range deck.ExtraDeck {
		// Extra deck cards
		err := InsertCardToDB(card, id, "extra_deck")
		if err != nil {
			return fmt.Errorf("error inserting card into extra deck: %s", err)
		}
	}

	return nil
}

// Saves deck to user's saved decks list
func SaveDeck(user_id int64, deck_id int64) error {
	_, err := DB.Exec("INSERT INTO saved_decks (user_id, deck_id) VALUES (?, ?)", user_id, deck_id)
	if err != nil {
		return fmt.Errorf("error saving deck: %s", err)
	}

	return nil
}

// Inserts an empty deck into DB, returns the ID of the new deck
func InsertDeckToDB(user user.User) (int64, error) {
	res, err := DB.Exec("INSERT INTO decks")
	if err != nil {
		return 0, fmt.Errorf("error creating decks table: %s", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert id: %s", err)
	}

	return id, nil
}

// Inserts a card into a deck in the DB
func InsertCardToDB(card card.Card, id int64, deckType string) error {
	_, err := DB.Exec("INSERT INTO deck_cards (deck_id, card_id, card_type) VALUES (?, ?, ?)", id, card.ID, deckType)
	if err != nil {
		return fmt.Errorf("error inserting card into deck: %s", err)
	}

	return nil
}
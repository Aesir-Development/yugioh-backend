package dbConnection

import (
	"net/http"
	"io"
	"github.com/Aesir-Development/yugioh-backend/pkg/card"
	"encoding/json"
)

// CARDS TABLE
/*
	"id INT NOT NULL AUTO_INCREMENT," +
	"name VARCHAR(255) NOT NULL," +
	"type VARCHAR(255) NOT NULL," +
	"frame_type VARCHAR(255) NOT NULL," +
	"description VARCHAR(255) NOT NULL," +
	"attack INT NOT NULL," +
	"defense INT NOT NULL," +
	"level INT NOT NULL," +
	"race VARCHAR(255) NOT NULL," +
	"attribute VARCHAR(255) NOT NULL," +
	"card_sets JSON," +
	"card_images JSON," +
	"card_prices JSON," +
	"PRIMARY KEY (id)" +
*/

// GetCards - Get all cards from API
func GetCards() []card.Card {
	resp, err := http.Get("https://db.ygoprodeck.com/api/v7/cardinfo.php")
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Error reading response body")
	}

	cards := card.ParseCards(body)

	return cards
}

// SaveCards - Save cards to DB
func SaveCards(cards []card.Card) {
	for _, card := range cards {
		cardImagesJSON := CardImagesToJSON(card.CardImages)
		cardSetsJSON := CardSetsToJSON(card.CardSets)
		cardPricesJSON := CardPricesToJSON(card.CardPrices)
	
		_, err := DB.Exec("INSERT INTO cards (name, type, frame_type, description, attack, defense, level, race, attribute, card_sets, card_images, card_prices) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		card.Name, card.Type, card.FrameType, card.Description, card.Attack, card.Defense, card.Level, card.Race, card.Attribute, cardSetsJSON, cardImagesJSON, cardPricesJSON)
	
		if err != nil {
			panic(err)
		}
	}
}

// SaveCard - Save a single card to DB
func SaveCard(card card.Card) {

	cardImagesJSON := CardImagesToJSON(card.CardImages)
	cardSetsJSON := CardSetsToJSON(card.CardSets)
	cardPricesJSON := CardPricesToJSON(card.CardPrices)

	_, err := DB.Exec("INSERT INTO cards (name, type, frame_type, description, attack, defense, level, race, attribute, card_sets, card_images, card_prices) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	card.Name, card.Type, card.FrameType, card.Description, card.Attack, card.Defense, card.Level, card.Race, card.Attribute, cardSetsJSON, cardImagesJSON, cardPricesJSON)

	if err != nil {
		panic(err)
	}
}

func CardImagesToJSON(images []card.CardImage) string {
	json, err := json.Marshal(images)
	if err != nil {
		panic(err)
	}
	return string(json)
}

func CardSetsToJSON(sets []card.CardSet) string {
	json, err := json.Marshal(sets)
	if err != nil {
		panic(err)
	}
	return string(json)
}

func CardPricesToJSON(prices []card.CardPrice) string {
	json, err := json.Marshal(prices)
	if err != nil {
		panic(err)
	}
	return string(json)
}
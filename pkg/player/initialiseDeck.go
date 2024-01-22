package player

import (
	"github.com/Aesir-Development/yugioh-backend/pkg/card"
	// conn "github.com/Aesir-Development/yugioh-backend/internal/db" 
)

// Here we get the deck from API and assign it to the player
func (p *Player) InitialiseDeck() {
	println("Initialising deck")
	p.Deck = FetchDeck()
}

// FetchDeck fetches a deck from the database.
func FetchDeck() []card.CardState {
	// A test deck for now
	// TODO - Fetch the deck from the database

	deck := []card.CardState{
		{
			Card: card.Card{
				Name: "Blue-Eyes White Dragon",
				Description: "This legendary dragon is a powerful engine of destruction.",
				Attack: 3000,
				Defense: 2500,
				Level: 8,
				Attribute: "LIGHT",
				Type: "Dragon",
			},
			Position: card.Deck,
		},
		{
			Card: card.Card{
				Name: "Dark Magician",
				Description: "The ultimate wizard in terms of attack and defense.",
				Attack: 2500,
				Defense: 2100,
				Level: 7,
				Attribute: "DARK",
				Type: "Spellcaster",
			},
			Position: card.Deck,
		},
		{
			Card: card.Card{
				Name: "Red-Eyes Black Dragon",
				Description: "A ferocious dragon with a deadly attack.",
				Attack: 2400,
				Defense: 2000,
				Level: 7,
				Attribute: "DARK",
				Type: "Dragon",
			},
			Position: card.Deck,
		},
	}
	return deck
}

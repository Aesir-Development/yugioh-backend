package deck

import (
	"github.com/Aesir-Development/yugioh-backend/pkg/card"
)

// NOTE: This uses the YGOPRODeck deck builder
// https://ygoprodeck.com/deckbuilder/

type Deck struct {
	Deck []card.CardState `json:"deck"`
	ExtraDeck []card.CardState `json:"extraDeck"`
	SideDeck []card.CardState `json:"sideDeck"`
}
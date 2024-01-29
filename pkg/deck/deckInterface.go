package deck

import (
	"github.com/Aesir-Development/yugioh-backend/pkg/card"
)

// NOTE - This will be using our own deck builder in the future

type Deck struct {
	Deck []card.CardState `json:"deck"`
	ExtraDeck []card.CardState `json:"extraDeck"`
	SideDeck []card.CardState `json:"sideDeck"`
}
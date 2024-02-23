package deck

import (
	"github.com/Aesir-Development/yugioh-backend/pkg/card"
)

// NOTE - This will be using our own deck builder in the future

type Deck struct {
	Deck []card.Card `json:"deck"`
	ExtraDeck []card.Card `json:"extraDeck"`
}
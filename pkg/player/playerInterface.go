package player

import (
	"github.com/Aesir-Development/yugioh-backend/pkg/card"
)


// Player struct
type Player struct {
	Deck []card.CardState
	Hand []card.CardState
	Graveyard []card.CardState
	ExtraDeck []card.CardState
	Field []card.CardState
}
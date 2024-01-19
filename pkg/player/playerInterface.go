package player

import (
	"github.com/Aesir-Development/yugioh-backend/internal/duel"
)


// Player struct
type Player struct {
	Deck []duel.CardState
	Hand []duel.CardState
	Graveyard []duel.CardState
	ExtraDeck []duel.CardState
	Field []duel.CardState
}
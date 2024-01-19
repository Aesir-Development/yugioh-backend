package duel

import (
)


// DuelManager is the main struct for the duel package
type DuelManager struct {
	Player1 *PlayerManager
	Player2 *PlayerManager

	CurrentState GameState
}

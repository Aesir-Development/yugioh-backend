package duel

import (
	"github.com/Aesir-Development/yugioh-backend/pkg/player"
)

type GameState int64

const (
	Waiting GameState = iota // When you're waiting for the duel to start
	PrepPhase // When you're in the preparation phase
	MainPhase // When you're in the main phase
	BattlePhase // When you're in the battle phase
	EndPhase // When you're in the end phase
)

// DuelManager is the main struct for the duel package
type DuelManager struct {
	Player1 *player.Player
	Player2 *player.Player

	CurrentState GameState
	ActivePlayer bool // True if player 1 is active, false if player 2 is active
}

// NewDuelManager creates a new DuelManager struct
func NewDuelManager(player1 *player.Player, player2 *player.Player) *DuelManager {
	return &DuelManager{
		Player1: player1,
		Player2: player2,
		CurrentState: Waiting,
		ActivePlayer: true,
	}
}

// StartDuel starts a duel between two players
func (d *DuelManager) StartDuel() {
	println("Starting duel")
	d.CurrentState = PrepPhase
}
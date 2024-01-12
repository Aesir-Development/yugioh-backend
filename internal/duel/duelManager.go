package duel

import (
	"github.com/Aesir-Development/yugioh-backend/pkg/state"
)


// DuelManager is the main struct for the duel package
type DuelManager struct {
	// TODO - add the player structs here
	cardManager *state.StateManager

	// Playfield test
	Card1 *state.StateManager
	Card2 *state.StateManager
}


// NewDuelManager creates a new DuelManager struct
func NewDuelManager() *DuelManager {

	// TODO - initialize the player structs here
	
	cardManager := state.NewStateManager()

	return &DuelManager{
		cardManager: cardManager,
	}

}


// StartDuel starts a new duel
func (d *DuelManager) StartDuel() {
	// TODO - implement this function
}

func (d *DuelManager) AddPlayer() {
	// TODO - implement this function
}

// A test function to test basic battle between 2 cards.
func (d *DuelManager) SelectCardsTest(card1 int, card2 int) {
	// Find cards by ID
	// TODO - implement this function

	// Create a new state manager for each card
	// TODO - implement this function

	// Set the cards to the duel manager
	// TODO - implement this function

	// Set the cards to face up
	// TODO - implement this function
}

// A test function to test basic battle between 2 cards.
func (d *DuelManager) BattleTest() {

}
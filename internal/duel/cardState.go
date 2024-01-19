package duel

import "github.com/Aesir-Development/yugioh-backend/pkg/card"

// Implementation of the StateManager interface
type CardState struct {
	card.Card
	FaceUp bool
	CanAttack bool
	CanChangePosition bool
	CanBeAttacked bool	
}

// Implementation of the StateManager interface
func (s *CardState) ToggleFaceUp() bool {
	s.FaceUp = !s.FaceUp
	return s.FaceUp
}

func (s *CardState) ToggleCanAttack() bool {
	s.CanAttack = !s.CanAttack
	return s.CanAttack
}

func (s *CardState) ToggleCanChangePosition() bool {
	s.CanChangePosition = !s.CanChangePosition
	return s.CanChangePosition
}

func (s *CardState) ToggleCanBeAttacked() bool {
	s.CanBeAttacked = !s.CanBeAttacked
	return s.CanBeAttacked
}

// NewStateManager creates a new StateManager struct
func NewCardState() *CardState {
	return &CardState{
		FaceUp: false,
		CanAttack: false,
		CanChangePosition: false,
		CanBeAttacked: false,
	}
}
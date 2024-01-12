package state


// Implementation of the StateManager interface
type StateManager struct {
	FaceUp bool
	CanAttack bool
	CanChangePosition bool
	CanBeAttacked bool	
}

// Implementation of the StateManager interface
func (s *StateManager) ToggleFaceUp() bool {
	s.FaceUp = !s.FaceUp
	return s.FaceUp
}

func (s *StateManager) ToggleCanAttack() bool {
	s.CanAttack = !s.CanAttack
	return s.CanAttack
}

func (s *StateManager) ToggleCanChangePosition() bool {
	s.CanChangePosition = !s.CanChangePosition
	return s.CanChangePosition
}

func (s *StateManager) ToggleCanBeAttacked() bool {
	s.CanBeAttacked = !s.CanBeAttacked
	return s.CanBeAttacked
}

// NewStateManager creates a new StateManager struct
func NewStateManager() *StateManager {
	return &StateManager{
		FaceUp: false,
		CanAttack: false,
		CanChangePosition: false,
		CanBeAttacked: false,
	}
}
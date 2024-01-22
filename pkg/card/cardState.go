package card

type CardPosition int64

const (
	Attack CardPosition = iota
	Defense
	Hand
	Graveyard
	ExtraDeck
	Field
	Deck
)

// Implementation of the StateManager interface
type CardState struct {
	Card
	FaceUp bool
	CanAttack bool
	CanChangePosition bool
	CanBeAttacked bool
	Position CardPosition
}
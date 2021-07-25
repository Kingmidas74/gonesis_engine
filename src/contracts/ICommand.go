package contracts

type ICommand interface {
	GetIdentifier() int
	IsInterrupts() bool
	Handle(terrain ITerrain, agent IAgent) int
}

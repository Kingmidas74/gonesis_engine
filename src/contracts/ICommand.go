package contracts

type ICommand interface {
	IsInterrupts() bool
	Handle(terrain ITerrain, agent IAgent) int
}

package contracts

type IReproduction interface {
	TryToReproduce(agent IAgent, terrain ITerrain) bool
}

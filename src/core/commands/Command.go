package commands

import "github.com/Kingmidas74/gonesis/contracts"

type Command struct {
	IsInterrupt bool
	Handler     func(contracts.ITerrain, contracts.IAgent) int
}

func (this *Command) IsInterrupts() bool {
	return this.IsInterrupt
}

func (this *Command) Handle(terrain contracts.ITerrain, agent contracts.IAgent) int {
	return this.Handler(terrain, agent)
}

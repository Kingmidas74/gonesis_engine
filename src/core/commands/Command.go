package commands

import "github.com/Kingmidas74/gonesis/contracts"

type Command struct {
	IsInterrupt bool
	Identifier  int
	Handler     func(contracts.ITerrain, contracts.IAgent)
}

func (this *Command) IsInterrupts() bool {
	return this.IsInterrupt
}

func (this *Command) GetIdentifier() int {
	return this.Identifier
}

func (this *Command) Handle(terrain contracts.ITerrain, agent contracts.IAgent) {
	this.Handler(terrain, agent)
}

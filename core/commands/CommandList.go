package commands

import "github.com/Kingmidas74/gonesis_engine/contracts"

type CommandList struct {
	Commands map[int]contracts.ICommand
}

func (this *CommandList) GetCommandByIdentifier(identifier int) contracts.ICommand {
	return this.Commands[identifier]
}

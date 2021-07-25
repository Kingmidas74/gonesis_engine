package commands

import "github.com/Kingmidas74/gonesis/contracts"

type CommandList struct {
	Commands []contracts.ICommand
}

func (this *CommandList) GetCommandByIdentifier(identifier int) contracts.ICommand {
	for i := range this.Commands {
		if this.Commands[i].GetIdentifier() == identifier {
			return this.Commands[i]
		}
	}
	return nil
}

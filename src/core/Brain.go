package core

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/core/commands"
	"github.com/Kingmidas74/gonesis/utils"
)

type Brain struct {
	CommandList    commands.CommandList
	Commands       []int
	CurrentAddress int
}

func (this *Brain) SetAddress(address int) {
	this.CurrentAddress = utils.ModLikePython(address, len(this.Commands))
}

func (this *Brain) GetCurrentCommand() contracts.ICommand {
	currentCommandIdentifier := this.GetCommandIdentifier(this.CurrentAddress)
	currentCommand := this.CommandList.GetCommandByIdentifier(currentCommandIdentifier)

	if currentCommand == nil {
		this.MoveAddressOn(currentCommandIdentifier)
	}

	return currentCommand
}

func (this *Brain) MoveAddressOn(delta int) {
	this.SetAddress(this.CurrentAddress + delta)
}

func (this *Brain) GetCommandIdentifier(address int) int {
	return this.Commands[utils.ModLikePython(address, len(this.Commands))]
}

func (this *Brain) GetCurrentAddress() int {
	return this.CurrentAddress
}

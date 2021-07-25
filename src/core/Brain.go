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

func (this *Brain) GetCurrentCommand() contracts.ICommand {
	currentCommandIdentifier := this.GetCommandIdentifier(this.CurrentAddress)
	currentCommand := this.CommandList.GetCommandByIdentifier(currentCommandIdentifier)

	if currentCommand == nil {
		this.MoveAddressOn(currentCommandIdentifier)
	}

	return currentCommand
}

func (this *Brain) MoveAddressOn(delta int) {
	this.SetCurrentAddress(this.CurrentAddress + delta)
}

func (this *Brain) GetCommandIdentifier(address int) int {
	return this.Commands[utils.ModLikePython(address, len(this.Commands))]
}

func (this *Brain) GetCurrentAddress() int {
	return this.CurrentAddress
}

func (this *Brain) SetCurrentAddress(address int) {
	this.CurrentAddress = utils.ModLikePython(address, len(this.Commands))
}

func (this *Brain) Clone() contracts.IBrain {
	newBrain := *this
	newBrain.SetCurrentAddress(0)
	return &newBrain
}

func (this *Brain) Mutate() {
	randomIndex := utils.RandomIntBetween(0, len(this.Commands)-1)
	this.Commands[randomIndex] = utils.RandomIntBetween(0, len(this.Commands)-1)
}

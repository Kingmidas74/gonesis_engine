package gonesis

type Brain struct {
	CommandList    CommandList
	Commands       []int
	CurrentAddress int
}

func (this *Brain) SetAddress(address int) {
	this.CurrentAddress = modLikePython(address, len(this.Commands))
}

func (this *Brain) GetCurrentCommand() *Command {
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
	return this.Commands[modLikePython(address, len(this.Commands))]
}

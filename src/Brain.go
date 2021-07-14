package gonesis

type Brain struct {
	CommandList           []Command
	Commands              []int
	currentCommandAddress int
}

func (this *Brain) MoveAddressOn(delta int) {
	this.SetAddress(this.currentCommandAddress + delta)
}

func (this *Brain) SetAddress(address int) {
	this.currentCommandAddress = modLikePython(address, len(this.Commands))
}

func (this *Brain) SearchCurrentCommand() *Command {
	for i := 0; i < len(this.CommandList); i++ {
		if this.CommandList[i].Identifier == this.GetCommandIdentifier(this.currentCommandAddress) {
			return &this.CommandList[i]
		}
	}
	return nil
}

func (this *Brain) GetCurrentCommand() *Command {
	currentCommand := this.SearchCurrentCommand()

	for currentCommand == nil {
		this.SetAddress(this.GetCommandIdentifier(this.currentCommandAddress))
		currentCommand = this.SearchCurrentCommand()
	}
	return currentCommand

}

func (this *Brain) GetCommandIdentifier(address int) int {
	return this.Commands[modLikePython(address, len(this.Commands))]
}

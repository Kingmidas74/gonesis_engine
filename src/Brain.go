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

func (this *Brain) GetCurrentCommand() *Command {

	commandIndex := this.Commands[this.currentCommandAddress]
	for commandIndex > len(this.CommandList)-1 {
		this.MoveAddressOn(commandIndex)
		commandIndex = this.Commands[this.currentCommandAddress]
	}
	return &this.CommandList[this.Commands[this.currentCommandAddress]]

}

func (this *Brain) GetCommandIdentifier(address int) int {
	return this.Commands[modLikePython(address, len(this.Commands))]
}

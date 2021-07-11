package gonesis

type Brain struct {
	Commands              []Command
	currentCommandAddress int
}

func (this *Brain) MoveAddressOn(delta int) {
	this.SetAddress(this.currentCommandAddress + delta)
}

func (this *Brain) SetAddress(address int) {
	this.currentCommandAddress = modLikePython(address, len(this.Commands))
}

func (this *Brain) GetCommand() Command {
	return this.Commands[this.currentCommandAddress]
}

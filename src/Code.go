package gonesis


type Code struct {
	Commands       *CommandList
	currentAddress int

}

func (this *Code) getCurrentCommand() Command  {
	command := this.Commands.getCurrentCommand()
	this.Commands.moveNext()
	this.currentAddress = this.Commands.currentIndex
	return command
}

func (this *Code) jumpTo(delta int)  {
	this.currentAddress+=delta
	//this.commands.currentIndex =
}
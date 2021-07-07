package gonesis

type CommandList struct {
	commands []Command
	length int
	currentIndex int
}

func CommandListConstructor() *CommandList {
	list := CommandList{}
	list.length = 0
	list.currentIndex = 0
	list.commands = make([]Command, 0)
	return &list
}

func (this *CommandList) getCurrentCommand() Command  {
	if this.currentIndex<0 || len(this.commands)==0 {
		panic("index out of range")
	}
	return this.commands[this.currentIndex]
}

func (this *CommandList) moveNext() bool {
	this.currentIndex++
	if this.currentIndex>=this.length {
		this.currentIndex=0
	}
	return true
}

func (this *CommandList) reset() bool {
	this.currentIndex = 0
	if this.currentIndex>=this.length {
		this.currentIndex=0
	}
	return true
}

func (this *CommandList) AppendCommand(command Command) bool {
	this.commands = append(this.commands, command)
	this.length++
	return true
}
package gonesis

type Brain struct {
	Commands              []Command
	CurrentCommandAddress int
}

func (this *Brain) MoveAddressOn(delta int) {
	if delta < 0 {
		panic("delta should be non negative")
	}
	this.CurrentCommandAddress = (this.CurrentCommandAddress + delta) % len(this.Commands)
}

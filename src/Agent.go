package gonesis

type Agent struct {
	Brain
	Direction Direction

	Energy int
}

func (this *Agent) nextDay(world *World, maxSteps int) {
	for step := 0; step < maxSteps && this.Energy > 0; step++ {
		command := this.GetCommand()
		command.Handler(world, this)
		this.Energy--
		if command.IsInterrupts {
			break
		}
	}
}

func (this *Agent) isAlive() bool {
	return this.Energy > 0
}

package gonesis

type Agent struct {
	Brain
	Coords    Coords
	Direction Direction

	Energy int
}

func (this *Agent) nextDay(world *World, maxSteps int) {
	for step, command := 0, this.Commands[this.CurrentCommandAddress]; step < maxSteps && !command.IsInterrupts; step++ {
		command.Handler(world, this)
	}
}

func (this *Agent) isAlive() bool {
	return this.Energy > 0
}

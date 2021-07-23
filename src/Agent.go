package gonesis

type Agent struct {
	Coords
	Brain

	Energy     int
	Generation int
}

func (this *Agent) NextDay(world *World, maxSteps int) {
	for step := 0; step < maxSteps && this.Energy > 0; step++ {
		this.Energy--
		command := this.GetCurrentCommand()
		if command == nil {
			continue
		}
		command.Handler(world, this)
		if command.IsInterrupts {
			break
		}
	}
}

func (this *Agent) IsAlive() bool {
	return this.Energy > 0
}

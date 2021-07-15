package gonesis

type Agent struct {
	Coords
	Brain

	Energy     int
	Generation int
}

func (this *Agent) NextDay(world *World, maxSteps int) {
	for step := 0; step < maxSteps && this.Energy > 0; step++ {
		command := this.GetCurrentCommand()
		command.Handler(world, this)

		this.Energy--

		if command.IsInterrupts {
			break
		}
	}
}

func (this *Agent) IsAlive() bool {
	return this.Energy > 0
}

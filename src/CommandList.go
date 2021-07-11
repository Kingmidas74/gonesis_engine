package gonesis

var (
	moveCommand = Command{
		IsInterrupts: true,
		Handler: func(world *World, agent *Agent) {
			agent.Brain.MoveAddressOn(1)
		},
	}
	eatCommand = Command{
		IsInterrupts: true,
		Handler: func(world *World, agent *Agent) {
			agent.Brain.MoveAddressOn(1)
		},
	}
)

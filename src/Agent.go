package gonesis


type Direction byte
const (
	up = 0
	right = 1
	down = 2
	left = 3
)

type Agent struct {

	code Code

	X int
	Y int
	direction Direction

	visibilityRadius int

	energy int
}

func AgentConstructor(code Code, energy int, X,Y int, direction Direction, visibilityRadius int) Agent {

	if visibilityRadius<1 || X+Y<0 || energy<1 {
		panic("wrong agent properties")
	}

	agent := Agent{}
	agent.energy = energy
	agent.code = code
	agent.X, agent.Y, agent.direction = X, Y, direction
	agent.visibilityRadius = visibilityRadius
	return agent
}

func (this* Agent) action(maxSteps int) {
	for currentCommand, step := this.code.getCurrentCommand(), 0; step < maxSteps; step++ {
		currentCommand.handler(this)
		if currentCommand.isAbortable {
			break
		}
	}
}

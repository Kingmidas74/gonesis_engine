package gonesis

type World struct {
	Terrain
	populationController Population

	tomorrowWorld *World
}

func (this *World) Action(agents []Agent, maxDays, maxSteps int) {
	agents = this.populationController.Reproduction(agents)
	for maxDays > 0 {
		agents = this.runDay(agents, maxSteps)
	}
}

func (this *World) runDay(agents []Agent, maxSteps int) []Agent {
	if this.tomorrowWorld == nil {
		this.tomorrowWorld = this
	}
	for i := 0; i < len(agents); i++ {
		agents[i].nextDay(this, maxSteps)
	}
	return agents
}

func (this *World) evaluateAgents(agents []Agent) []Agent {
	livingAgents := this.filterLivingAgents(agents)
	livingAgents = this.populationController.Reproduction(livingAgents)
	return livingAgents
}

func (this *World) filterLivingAgents(agents []Agent) []Agent {
	result := make([]Agent, 0)
	for _, agent := range agents {
		if agent.isAlive() {
			result = append(result, agent)
		}
	}
	return result
}

func (this *World) GetCellInfo(latitude int, longitude int) *Cell {

	return &this.Cells[latitude*this.LongitudesCount+longitude]
}

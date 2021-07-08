package gonesis

type World struct {
	Cells                []Cell
	LatitudesCount       int
	LongitudesCount      int
	populationController Population
}

func (this *World) Action(agents []Agent, maxDays, maxSteps int) {
	agents = this.populationController.Reproduction(agents)
	for maxDays > 0 {
		agents = this.runDay(agents, maxSteps)
	}
}

func (this *World) runDay(agents []Agent, maxSteps int) []Agent {
	for _, agent := range agents {
		agent.nextDay(this, maxSteps)
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

package gonesis

type World struct {
	Terrain
	populationController Population
}

func (this *World) Action(agents []Agent, maxDays, maxSteps int, callback func(*Terrain, int)) {
	if len(agents) == 0 {
		return
	}
	passDays := 0
	for true {
		isNewGeneration, newAgents, err := this.evaluateAgents(agents)
		if err != nil {
			return
		} else if isNewGeneration {
			agents = newAgents
			this.placeAgents(agents)
		}
		callback(&this.Terrain, passDays)
		this.runDay(agents, maxSteps)

		maxDays--
		passDays++
	}
}

func (this *World) runDay(agents []Agent, maxSteps int) {
	for i := 0; i < len(agents); i++ {
		agents[i].NextDay(this, maxSteps)
	}
}

func (this *World) evaluateAgents(agents []Agent) (bool, []Agent, error) {
	return this.populationController.Reproduction(this.filterLivingAgents(agents))
}

func (this *World) filterLivingAgents(agents []Agent) []Agent {
	result := make([]Agent, 0)
	for i := 0; i < len(agents); i++ {
		if agents[i].IsAlive() {
			result = append(result, agents[i])
		}
	}
	return result
}

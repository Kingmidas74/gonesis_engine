package gonesis

type World struct {
	Terrain
	populationController Population
}

func (this *World) Action(agents []Agent, maxDays, maxSteps int, callback func(*Terrain, int)) {
	agents = this.populationController.Reproduction(agents)
	for maxDays > 0 {
		this.Terrain.placeAgents(agents)
		this.runDay(agents, maxSteps)
		callback(&this.Terrain, maxDays)
		//agents = this.evaluateAgents(agents)
		maxDays--
	}
}

func (this *World) runDay(agents []Agent, maxSteps int) {
	for i := 0; i < len(agents); i++ {
		agents[i].nextDay(this, maxSteps)
	}
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

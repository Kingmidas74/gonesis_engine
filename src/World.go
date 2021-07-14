package gonesis

type World struct {
	Terrain
	populationController Population
}

func (this *World) Action(agents []Agent, maxDays, maxSteps int, callback func(*Terrain, int)) {
	passDays := 0
	for true {
		if len(agents) == 0 {
			break
		}
		callback(&this.Terrain, passDays)
		isNewGeneration, newAgents, err := this.evaluateAgents(agents)
		if isNewGeneration {
			agents = newAgents
			this.Terrain.placeAgents(agents)
		} else if err != nil {
			break
		}
		this.runDay(agents, maxSteps)
		//agents = this.evaluateAgents(agents)
		maxDays--
		passDays++
	}
}

func (this *World) runDay(agents []Agent, maxSteps int) {
	for i := 0; i < len(agents); i++ {
		agents[i].nextDay(this, maxSteps)
	}
}

func (this *World) evaluateAgents(agents []Agent) (bool, []Agent, error) {
	livingAgents := this.filterLivingAgents(agents)
	if len(livingAgents) == 0 {
		panic("end")
	}
	return this.populationController.Reproduction(livingAgents)
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

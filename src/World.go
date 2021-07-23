package gonesis

type World struct {
	Terrain
	populationController Population
}

func (this *World) Action(maxSteps int, callback func(*Terrain, int)) {
	//this.placeAgents(agents)
	passDays := 0
	for true {
		livingAgents := this.filterLivingAgents()
		if len(livingAgents) == 0 {
			return
		}
		this.cleanDeath()
		callback(&this.Terrain, passDays)
		this.runDay(livingAgents, maxSteps)
		passDays++
	}
}

func (this *World) runDay(agents []*Agent, maxSteps int) {
	for i := 0; i < len(agents); i++ {
		agents[i].NextDay(this, maxSteps)
	}
}

func (this *World) filterLivingAgents() []*Agent {
	result := make([]*Agent, 0)
	for _, cell := range this.Cells {
		if cell.Agent != nil && cell.Agent.IsAlive() {
			result = append(result, cell.Agent)
		}
	}
	return result
}

package gonesis

import "github.com/Kingmidas74/gonesis/contracts"

type World struct {
	contracts.ITerrain
}

func (this *World) Action(maxSteps int, callback func(contracts.ITerrain, int)) {
	passDays := 0
	for true {
		livingAgents := this.filterLivingAgents()
		if len(livingAgents) == 0 {
			return
		}
		this.cleanDeath()
		callback(this.ITerrain, passDays)
		this.runDay(livingAgents, maxSteps)
		passDays++
	}
}

func (this *World) runDay(agents []contracts.IAgent, maxSteps int) {
	for i := 0; i < len(agents); i++ {
		agents[i].NextDay(this.ITerrain, maxSteps)
	}
}

func (this *World) filterLivingAgents() []contracts.IAgent {
	result := make([]contracts.IAgent, 0)
	for _, cell := range this.GetCells() {
		if cell.GetAgent() != nil && cell.GetAgent().IsAlive() {
			result = append(result, cell.GetAgent())
		}
	}
	return result
}

func (this *World) cleanDeath() {
	for y := 0; y < this.GetHeight(); y++ {
		for x := 0; x < this.GetWidth(); x++ {
			currentCell := this.GetCell(x, y)
			if currentCell.GetAgent() != nil && !currentCell.GetAgent().IsAlive() {
				currentCell.SetAgent(nil)
				currentCell.SetCellType(contracts.EmptyCell)
			}
		}
	}
}

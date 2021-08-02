package gonesis_engine

import "github.com/Kingmidas74/gonesis_engine/contracts"

type World struct {
	contracts.ITerrain
}

func (this *World) Action(maxSteps int, callback func(contracts.ITerrain, int)) {
	callback(this.ITerrain, 0)
	for currentDay, livingAgents := 1, this.filterLivingAgents(); len(livingAgents) > 0; currentDay, livingAgents = currentDay+1, this.filterLivingAgents() {
		this.runDay(livingAgents, maxSteps)
		this.cleanDeath()
		callback(this.ITerrain, currentDay)
	}
}

func (this *World) runDay(agents []contracts.IAgent, maxSteps int) {
	for _, agent := range agents {
		agent.NextDay(this.ITerrain, maxSteps)
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

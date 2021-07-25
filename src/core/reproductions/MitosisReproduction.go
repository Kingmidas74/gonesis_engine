package reproductions

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/utils"
)

type MitosisReproduction struct {
	ReproductionPower   int
	GenerationPower     int
	MutationProbability int
}

func (this *MitosisReproduction) TryToReproduce(agent contracts.IAgent, terrain contracts.ITerrain) bool {
	if children, freeCells := this.makeChildren(agent), this.findEmptyNeighborsCell(agent, terrain); len(children) > 0 && len(freeCells) > 0 {

		placedChildrenCount := 0

		for placedChildrenCount < len(children) {
			freeCells[placedChildrenCount].SetCellType(contracts.LockedCell)
			children[placedChildrenCount].SetX(freeCells[placedChildrenCount].GetX())
			children[placedChildrenCount].SetY(freeCells[placedChildrenCount].GetY())
			freeCells[placedChildrenCount].SetAgent(children[placedChildrenCount])
			placedChildrenCount++
		}
		return true
	}
	return false
}

func (this *MitosisReproduction) findEmptyNeighborsCell(agent contracts.IAgent, terrain contracts.ITerrain) []contracts.ICell {
	cells := terrain.GetNeighbors(agent.GetX(), agent.GetY())
	result := make([]contracts.ICell, 0)

	for i := 0; i < len(cells); i++ {
		if cell := cells[i]; cell.GetCellType() == contracts.EmptyCell {
			result = append(result, cell)
		}
	}
	return result
}

func (this *MitosisReproduction) makeChildren(agent contracts.IAgent) []contracts.IAgent {
	children := make([]contracts.IAgent, 0)
	if agent.GetEnergy() >= this.ReproductionPower {
		for len(children) < this.GenerationPower {
			child := agent.MakeChild()
			child.SetEnergy(this.ReproductionPower / this.GenerationPower)
			child.SetGeneration(agent.GetGeneration() + 1)
			childBrain := agent.GetBrain().Clone()

			if utils.RandomIntBetween(0, 100) < this.MutationProbability {
				childBrain.Mutate()
			}

			childBrain.SetCurrentAddress(0)
			child.SetBrain(childBrain)
			children = append(children, child)
		}
		agent.SetEnergy(0)
	}
	return children
}

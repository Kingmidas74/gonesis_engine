package reproductions

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/utils"
)

type BuddingReproduction struct {
	ReproductionPower   int
	MutationProbability int
}

func (this *BuddingReproduction) TryToReproduce(agent contracts.IAgent, terrain contracts.ITerrain) bool {
	if child, freeCell := this.makeChild(agent), this.findEmptyNeighborCell(agent, terrain); child != nil && freeCell != nil {
		freeCell.SetCellType(contracts.LockedCell)
		child.SetX(freeCell.GetX())
		child.SetY(freeCell.GetY())
		freeCell.SetAgent(child)
		return true
	}
	return false
}

func (this *BuddingReproduction) findEmptyNeighborCell(agent contracts.IAgent, terrain contracts.ITerrain) contracts.ICell {
	cells := terrain.GetNeighbors(agent.GetX(), agent.GetY())

	for i := 0; i < len(cells); i++ {
		if cell := cells[i]; cell.GetCellType() == contracts.EmptyCell {
			return cell
		}
	}
	return nil
}

func (this *BuddingReproduction) makeChild(agent contracts.IAgent) contracts.IAgent {
	if agent.GetEnergy() >= this.ReproductionPower {
		child := agent.MakeChild()
		child.SetEnergy(this.ReproductionPower)
		child.SetGeneration(agent.GetGeneration() + 1)
		childBrain := agent.GetBrain().Clone()

		if utils.RandomIntBetween(0, 100) < this.MutationProbability {
			childBrain.Mutate()
		}

		childBrain.SetCurrentAddress(0)
		child.SetBrain(childBrain)
		agent.SetEnergy(agent.GetEnergy() - this.ReproductionPower)
		return child
	}
	return nil
}

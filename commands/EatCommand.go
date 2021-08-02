package commands

import (
	"github.com/Kingmidas74/gonesis_engine/contracts"
	"github.com/Kingmidas74/gonesis_engine/core/commands"
)

type EatCommand struct {
	commands.Command
}

func (this *EatCommand) Handle(terrain contracts.ITerrain, agent contracts.IAgent) int {
	currentCell := terrain.GetCell(agent.GetX(), agent.GetY())

	argument := agent.GetBrain().GetCommandIdentifier(agent.GetBrain().GetCurrentAddress() + 1)
	targetCell := terrain.GetNeighbor(currentCell.GetX(), currentCell.GetY(), argument)
	stepLength := this.calculateStepLength(targetCell)

	if targetCell.GetCellType() == contracts.OrganicCell {
		agent.SetX(targetCell.GetX())
		agent.SetY(targetCell.GetY())
		agent.SetEnergy(agent.GetEnergy() + targetCell.GetCost())
		targetCell.SetAgent(agent)
		targetCell.SetCellType(contracts.LockedCell)
		currentCell.SetAgent(nil)
		currentCell.SetCellType(contracts.EmptyCell)
		currentCell.SetCost(0)
	}

	return stepLength
}

func (this *EatCommand) calculateStepLength(cell contracts.ICell) int {

	if cell.GetCellType() == contracts.ObstacleCell {
		return 2
	}
	if cell.GetCellType() == contracts.LockedCell {
		return 3
	}
	if cell.GetCellType() == contracts.OrganicCell {
		if cell.GetCost() < 0 {
			return 1
		}
		return 4
	}
	return 5
}

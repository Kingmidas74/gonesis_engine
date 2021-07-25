package commands

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/core/commands"
)

type EatCommand struct {
	commands.Command
}

func (this *EatCommand) Handle(terrain contracts.ITerrain, agent contracts.IAgent) {
	currentCell := terrain.GetCell(agent.GetX(), agent.GetY())

	argument := agent.GetBrain().GetCommandIdentifier(agent.GetBrain().GetCurrentAddress() + 1)
	targetCell := terrain.GetNeighbor(currentCell.GetX(), currentCell.GetY(), argument)

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

	agent.GetBrain().MoveAddressOn(2)
}

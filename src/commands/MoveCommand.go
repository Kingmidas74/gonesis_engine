package commands

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/core/commands"
)

type MoveCommand struct {
	commands.Command
}

func (this *MoveCommand) Handle(terrain contracts.ITerrain, agent contracts.IAgent) {
	currentCell := terrain.GetCell(agent.GetX(), agent.GetY())

	argument := agent.GetBrain().GetCommandIdentifier(agent.GetBrain().GetCurrentAddress() + 1)
	targetCell := terrain.GetNeighbor(currentCell.GetX(), currentCell.GetY(), argument)

	if targetCell.GetCellType() == contracts.EmptyCell {
		agent.SetX(targetCell.GetX())
		agent.SetY(targetCell.GetY())
		targetCell.SetAgent(agent)
		targetCell.SetCellType(contracts.LockedCell)
		currentCell.SetAgent(nil)
		currentCell.SetCellType(contracts.EmptyCell)

	}

	agent.GetBrain().MoveAddressOn(2)
}

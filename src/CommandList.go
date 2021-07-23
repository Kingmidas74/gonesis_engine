package gonesis

type CommandList struct {
	Commands []Command
}

func (this *CommandList) GetCommandByIdentifier(identifier int) *Command {
	for i := range this.Commands {
		if this.Commands[i].Identifier == identifier {
			return &this.Commands[i]
		}
	}
	return nil
}

func getTargetCell(world *World, agent *Agent, direction int) *Cell {

	targetX := agent.Coords.X
	targetY := agent.Coords.Y

	switch direction {
	case NorthDirection:
		targetY -= 1
	case NorthEastDirection:
		targetY -= 1
		targetX += 1
	case EastDirection:
		targetX += 1
	case SouthEastDirection:
		targetX += 1
		targetY += 1
	case SouthDirection:
		targetY += 1
	case SouthWestDirection:
		targetX -= 1
		targetY += 1
	case WestDirection:
		targetX -= 1
	case NorthWestDirection:
		targetX -= 1
		targetY -= 1
	}
	return world.GetCell(targetX, targetY)
}

var (
	moveCommand = Command{
		IsInterrupts: true,
		Identifier:   0,
		Handler: func(world *World, agent *Agent) {
			currentCell := world.GetCell(agent.X, agent.Y)

			argument := agent.Brain.GetCommandIdentifier(agent.Brain.CurrentAddress + 1)
			direction := modLikePython(argument, 8)
			targetCell := getTargetCell(world, agent, direction)

			if targetCell.CellType == EmptyCell {
				agent.X = targetCell.X
				agent.Y = targetCell.Y
				targetCell.Agent = agent
				targetCell.CellType = LockedCell
				currentCell.Agent = nil
				currentCell.CellType = EmptyCell

			}

			agent.Brain.MoveAddressOn(2)
		},
	}
	eatCommand = Command{
		IsInterrupts: false,
		Identifier:   1,
		Handler: func(world *World, agent *Agent) {
			currentCell := world.GetCell(agent.X, agent.Y)

			argument := agent.Brain.GetCommandIdentifier(agent.Brain.CurrentAddress + 1)
			direction := modLikePython(argument, 8)
			targetCell := getTargetCell(world, agent, direction)

			if targetCell.CellType == OrganicCell {
				agent.X = targetCell.X
				agent.Y = targetCell.Y
				agent.Energy += targetCell.Cost
				targetCell.Agent = agent
				targetCell.CellType = LockedCell
				currentCell.Agent = nil
				currentCell.CellType = EmptyCell
				currentCell.Cost = 0
			}

			agent.Brain.MoveAddressOn(2)
		},
	}
	waitCommand = Command{
		IsInterrupts: false,
		Identifier:   2,
		Handler: func(world *World, agent *Agent) {
			agent.Brain.MoveAddressOn(1)
		},
	}
)

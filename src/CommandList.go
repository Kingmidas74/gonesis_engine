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
	case North:
		targetY -= 1
	case NorthEast:
		targetY -= 1
		targetX += 1
	case East:
		targetX += 1
	case SouthEast:
		targetX += 1
		targetY += 1
	case South:
		targetY += 1
	case SouthWest:
		targetX -= 1
		targetY += 1
	case West:
		targetX -= 1
	case NorthWest:
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

			argument := 0 //agent.Brain.GetCommandIdentifier(agent.Brain.currentCommandAddress + 1)
			direction := modLikePython(argument, 8)
			targetCell := getTargetCell(world, agent, direction)

			if targetCell.CellType == Empty {
				agent.X = targetCell.X
				agent.Y = targetCell.Y
				targetCell.Agent = agent
				targetCell.CellType = Locked
				currentCell.Agent = nil
				currentCell.CellType = Empty

			}

			agent.Brain.MoveAddressOn(2)
		},
	}
	eatCommand = Command{
		IsInterrupts: false,
		Identifier:   1,
		Handler: func(world *World, agent *Agent) {
			currentCell := world.GetCell(agent.X, agent.Y)

			argument := 0 // agent.Brain.GetCommandIdentifier(agent.Brain.currentCommandAddress + 1)
			direction := modLikePython(argument, 8)
			targetCell := getTargetCell(world, agent, direction)

			if targetCell.CellType == Organic {
				agent.X = targetCell.X
				agent.Y = targetCell.Y
				agent.Energy += targetCell.Cost
				targetCell.Agent = agent
				targetCell.CellType = Locked
				currentCell.Agent = nil
				currentCell.CellType = Empty
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

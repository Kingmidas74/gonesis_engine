package gonesis

func getTargetCell(world *World, agent *Agent, direction int) *Cell {

	targetLatitude := agent.Coords.Latitude
	targetLongitude := agent.Coords.Longitude

	switch direction {
	case North:
		targetLatitude -= 1
	case NorthEast:
		targetLatitude -= 1
		targetLongitude += 1
	case East:
		targetLongitude += 1
	case SouthEast:
		targetLatitude += 1
		targetLongitude += 1
	case South:
		targetLatitude += 1
	case SouthWest:
		targetLatitude += 1
		targetLongitude -= 1
	case West:
		targetLongitude -= 1
	case NorthWest:
		targetLatitude -= 1
		targetLongitude -= 1
	}
	return world.GetCell(targetLatitude, targetLongitude)
}

var (
	moveCommand = Command{
		IsInterrupts: true,
		Handler: func(world *World, agent *Agent) {
			currentCell := world.GetCell(agent.Latitude, agent.Longitude)

			argument := agent.Brain.GetCommandIdentifier(agent.Brain.currentCommandAddress + 1)
			direction := modLikePython(argument, 8)
			targetCell := getTargetCell(world, agent, direction)

			if targetCell.CellType == Empty {
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
		Handler: func(world *World, agent *Agent) {

			argument := agent.Brain.GetCommandIdentifier(agent.Brain.currentCommandAddress + 1)
			direction := modLikePython(argument, 8)
			targetCell := getTargetCell(world, agent, direction)

			if targetCell.CellType == Organic {
				targetCell.CellType = Empty
				agent.Energy += targetCell.Cost
			}

			agent.Brain.MoveAddressOn(2)
		},
	}
	waitCommand = Command{
		IsInterrupts: false,
		Handler: func(world *World, agent *Agent) {
			agent.Brain.MoveAddressOn(1)
		},
	}
)

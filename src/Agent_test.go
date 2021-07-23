package gonesis

import "testing"

func TestAgent_IsAlive_Success(t *testing.T) {
	agent := Agent{Energy: -1}

	if agent.IsAlive() {
		t.Errorf("Agent shouldn't be alive with negative energy")
	}
}

func TestNextDay(t *testing.T) {
	agent := Agent{
		Energy: 3,
		Brain: Brain{
			Commands: []int{0, 0, 0, 0},
		},
		Coords: Coords{
			X: 0,
			Y: 0,
		},
	}

	for i := 0; i < 4; i++ {
		command := Command{}
		command.IsInterrupts = i == 2
		command.Identifier = i
		command.Handler = func(world *World, currentAgent *Agent) {
			currentAgent.MoveAddressOn(1)
		}

		agent.CommandList.Commands = append(agent.CommandList.Commands, command)
	}

	agent.NextDay(&World{
		Terrain: Terrain{
			Cells: []Cell{Cell{
				Coords: Coords{
					X: 0,
					Y: 0,
				},
				CellType: LockedCell,
				Cost:     0,
				Agent:    &agent,
			}},
			Width:  1,
			Height: 1,
		},
		populationController: Population{},
	}, 10)

	if agent.IsAlive() {
		t.Errorf("Agent shouldn't be alive with negative energy")
	}
}

func TestAgent_FindEmptyNeighborCell(t *testing.T) {
	agent := Agent{
		Energy: 3,
		Brain: Brain{
			Commands: []int{0, 0, 0, 0},
		},
		Coords: Coords{
			X: 2,
			Y: 2,
		},
	}

	terrain := Terrain{
		Cells:  make([]Cell, 0),
		Width:  3,
		Height: 3,
	}
	for y := 0; y < terrain.Height; y++ {
		for x := 0; x < terrain.Width; x++ {
			terrain.Cells = append(terrain.Cells, Cell{
				Coords: Coords{
					X: x,
					Y: y,
				},
				CellType: LockedCell,
			})
		}
	}
	terrain.Cells[8].CellType = LockedCell
	terrain.Cells[8].Agent = &agent

	terrain.Cells[0].CellType = EmptyCell

	world := World{
		Terrain:              terrain,
		populationController: Population{},
	}

	freeCell := agent.FindEmptyNeighborCell(&world)

	if freeCell.X+freeCell.Y != 0 {
		t.Errorf("Wrong free cell")
	}

}

func TestAgent_MakeChild_SuccessWhenCanMakeAChild(t *testing.T) {
	agent := Agent{
		Coords: Coords{
			X: 2,
			Y: 2,
		},
		Brain: Brain{
			Commands: []int{0, 0, 0, 0},
		},
		Energy:            12,
		ReproductionPower: 10,
	}

	child := agent.MakeChild()

	if child == nil || agent.Energy != 2 {
		t.Errorf("Wrong child")
	}

}

func TestAgent_MakeChild_FailedWhenNotEnoughEnergy(t *testing.T) {
	agent := Agent{
		Coords: Coords{
			X: 2,
			Y: 2,
		},
		Brain: Brain{
			Commands: []int{0, 0, 0, 0},
		},
		Energy:            2,
		ReproductionPower: 10,
	}

	child := agent.MakeChild()

	if child != nil {
		t.Errorf("Wrong child")
	}

}

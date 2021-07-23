package gonesis

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRunDay(t *testing.T) {
	randCommandIndices := []int{3, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 2, 2, 1, 1, 1, 2, 2, 3}
	agents := make([]*Agent, 0)
	for i := 0; i < 2; i++ {
		rand.Shuffle(len(randCommandIndices), func(i, j int) {
			randCommandIndices[i], randCommandIndices[j] = randCommandIndices[j], randCommandIndices[i]
		})
		agents = append(agents, &Agent{
			Energy: 20,
			Brain: Brain{
				CommandList: CommandList{Commands: []Command{moveCommand, eatCommand}},
				Commands:    randCommandIndices,
			},
		})
	}

	terrain := Terrain{}
	terrain.Generate(5, 5, 30, 10)
	terrain.placeAgents(agents)

	world := World{
		Terrain: terrain,
		populationController: Population{
			NextGenerationLine:  1,
			MutationProbability: 0,
			Size:                len(agents),
		},
	}
	for y := 0; y < terrain.Height; y++ {
		for x := 0; x < terrain.Width; x++ {
			currentCell := terrain.GetCell(x, y)
			fmt.Printf("%2d %2d | ", currentCell.CellType, currentCell.Cost)
		}
		println()
	}
	drawFrame(&terrain, 1)
	world.runDay(agents, 3)

}

func TestWorld_Action(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	//	randCommandIndices := rand.Perm(10)

	agents := make([]*Agent, 0)
	for i := 0; i < 1; i++ {
		agents = append(agents, &Agent{
			Energy: 10,
			Brain: Brain{
				CommandList: CommandList{Commands: []Command{moveCommand, eatCommand}},
				Commands: []int{
					0, 4, //down
					1, 4, //eat down
					0, 2, //right
					1, 2, //eat right
					8,
				},
			},
		})
	}

	terrain := Terrain{}
	terrain.Generate(25, 25, 10, 16)
	terrain.placeAgents(agents)

	world := World{
		Terrain: terrain,
		populationController: Population{
			NextGenerationLine:  2,
			MutationProbability: 0,
			Size:                10,
		},
	}

	world.Action(50, drawFrame)

}

func TestWorld_Action_SpecificAgent(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	//randCommandIndices := rand.Perm(10)

	agents := make([]Agent, 0)
	agents = append(agents, Agent{
		Energy:            2,
		ReproductionPower: 50,
		Brain: Brain{
			CommandList: CommandList{Commands: []Command{moveCommand, eatCommand}},
			Commands: []int{
				0, 4, //down
				1, 4, //eat down
				0, 2, //right
				1, 2, //eat right
				8,
			},
		},
		Coords: Coords{
			X: 1,
			Y: 0,
		},
	})

	terrain := Terrain{
		Cells:  make([]Cell, 0),
		Width:  10,
		Height: 5,
	}
	for y := 0; y < terrain.Height; y++ {
		for x := 0; x < terrain.Width; x++ {
			terrain.Cells = append(terrain.Cells, Cell{
				Coords: Coords{
					X: x,
					Y: y,
				},
				CellType: EmptyCell,
				Cost:     0,
			})
		}
	}
	terrain.Cells[1].CellType = LockedCell
	terrain.Cells[1].Agent = &agents[0]

	terrain.Cells[21].CellType = OrganicCell
	terrain.Cells[21].Cost = 3

	terrain.Cells[23].CellType = OrganicCell
	terrain.Cells[23].Cost = 3

	terrain.Cells[43].CellType = OrganicCell
	terrain.Cells[43].Cost = 3

	terrain.Cells[45].CellType = OrganicCell
	terrain.Cells[45].Cost = 3

	terrain.Cells[15].CellType = OrganicCell
	terrain.Cells[15].Cost = 3

	terrain.Cells[17].CellType = OrganicCell
	terrain.Cells[17].Cost = 3

	terrain.Cells[37].CellType = OrganicCell
	terrain.Cells[37].Cost = 3

	terrain.Cells[39].CellType = OrganicCell
	terrain.Cells[39].Cost = 3

	terrain.Cells[9].CellType = OrganicCell
	terrain.Cells[9].Cost = 3

	world := World{
		Terrain: terrain,
		populationController: Population{
			NextGenerationLine:  0,
			MutationProbability: 0,
			Size:                1,
		},
	}

	world.Action(1, drawFrame)

}

func TestWorld_Action_SpecificAgentWithChild(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	//randCommandIndices := rand.Perm(10)

	agents := make([]Agent, 0)
	agents = append(agents, Agent{
		Energy:            2,
		ReproductionPower: 10,
		Brain: Brain{
			CommandList: CommandList{Commands: []Command{moveCommand, eatCommand}},
			Commands: []int{
				0, 4, //down
				1, 4, //eat down
				0, 2, //right
				1, 2, //eat right
				8,
			},
		},
		Coords: Coords{
			X: 1,
			Y: 0,
		},
	})

	terrain := Terrain{
		Cells:  make([]Cell, 0),
		Width:  10,
		Height: 5,
	}
	for y := 0; y < terrain.Height; y++ {
		for x := 0; x < terrain.Width; x++ {
			terrain.Cells = append(terrain.Cells, Cell{
				Coords: Coords{
					X: x,
					Y: y,
				},
				CellType: EmptyCell,
				Cost:     0,
			})
		}
	}
	terrain.Cells[1].CellType = LockedCell
	terrain.Cells[1].Agent = &agents[0]

	terrain.Cells[21].CellType = OrganicCell
	terrain.Cells[21].Cost = 4

	terrain.Cells[23].CellType = OrganicCell
	terrain.Cells[23].Cost = 4

	terrain.Cells[43].CellType = OrganicCell
	terrain.Cells[43].Cost = 4

	terrain.Cells[45].CellType = OrganicCell
	terrain.Cells[45].Cost = 4

	terrain.Cells[15].CellType = OrganicCell
	terrain.Cells[15].Cost = 4

	terrain.Cells[17].CellType = OrganicCell
	terrain.Cells[17].Cost = 4

	terrain.Cells[37].CellType = OrganicCell
	terrain.Cells[37].Cost = 4

	terrain.Cells[39].CellType = OrganicCell
	terrain.Cells[39].Cost = 4

	terrain.Cells[9].CellType = OrganicCell
	terrain.Cells[9].Cost = 4

	world := World{
		Terrain: terrain,
		populationController: Population{
			NextGenerationLine:  0,
			MutationProbability: 0,
			Size:                1,
		},
	}

	world.Action(1, drawFrame)

}

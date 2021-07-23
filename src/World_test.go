package gonesis

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFilterLivingAgents(t *testing.T) {
	agents := make([]Agent, 0)
	for i := 0; i < 6; i++ {
		agent := Agent{}
		agent.Energy = i % 2
		agents = append(agents, agent)
	}
	world := World{}
	livingAgents := world.filterLivingAgents(agents)
	if len(livingAgents)*2 != len(agents) {
		t.Errorf("filtration of living is wrong")
	}
}

func TestEvaluateAgents(t *testing.T) {
	agents := make([]Agent, 0)
	for i := 0; i < 6; i++ {
		agents = append(agents, Agent{
			Energy: i % 2,
		})
	}

	world := World{
		Terrain: Terrain{
			Cells:  nil,
			Width:  1024,
			Height: 1024,
		},
		populationController: Population{
			NextGenerationLine:  8,
			MutationProbability: 0,
			Size:                64,
		},
	}
	_, livingAgents, err := world.evaluateAgents(agents)
	if len(livingAgents) != world.populationController.Size {
		t.Errorf("evaluation failed")
	} else if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRunDay(t *testing.T) {
	randCommandIndices := []int{3, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 2, 2, 1, 1, 1, 2, 2, 3}
	agents := make([]Agent, 0)
	for i := 0; i < 2; i++ {
		rand.Shuffle(len(randCommandIndices), func(i, j int) {
			randCommandIndices[i], randCommandIndices[j] = randCommandIndices[j], randCommandIndices[i]
		})
		agents = append(agents, Agent{
			Energy: 20,
			Brain: Brain{
				CommandList: CommandList{Commands: []Command{moveCommand, eatCommand}},
				Commands:    randCommandIndices,
			},
		})
	}

	terrain := Terrain{}
	terrain.Generate(5, 5, agents, 30, 10)

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

	randCommandIndices := rand.Perm(10)

	agents := make([]Agent, 0)
	for i := 0; i < 1; i++ {
		agents = append(agents, Agent{
			Energy: 10,
			Brain: Brain{
				CommandList: CommandList{Commands: []Command{moveCommand, eatCommand}},
				Commands:    randCommandIndices,
			},
		})
	}

	terrain := Terrain{}
	terrain.Generate(25, 25, agents, 10, 16)

	world := World{
		Terrain: terrain,
		populationController: Population{
			NextGenerationLine:  2,
			MutationProbability: 0,
			Size:                10,
		},
	}

	world.Action(agents, 50, drawFrame)

}

func TestWorld_Action_SpecificAgent(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	//randCommandIndices := rand.Perm(10)

	agents := make([]Agent, 0)
	agents = append(agents, Agent{
		Energy: 2,
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
				CellType: Empty,
				Cost:     0,
			})
		}
	}
	terrain.Cells[1].CellType = Locked
	terrain.Cells[1].Agent = &agents[0]

	terrain.Cells[21].CellType = Organic
	terrain.Cells[21].Cost = 3

	terrain.Cells[23].CellType = Organic
	terrain.Cells[23].Cost = 3

	terrain.Cells[43].CellType = Organic
	terrain.Cells[43].Cost = 3

	terrain.Cells[45].CellType = Organic
	terrain.Cells[45].Cost = 3

	terrain.Cells[15].CellType = Organic
	terrain.Cells[15].Cost = 3

	terrain.Cells[17].CellType = Organic
	terrain.Cells[17].Cost = 3

	terrain.Cells[37].CellType = Organic
	terrain.Cells[37].Cost = 3

	terrain.Cells[39].CellType = Organic
	terrain.Cells[39].Cost = 3

	terrain.Cells[9].CellType = Organic
	terrain.Cells[9].Cost = 3

	world := World{
		Terrain: terrain,
		populationController: Population{
			NextGenerationLine:  0,
			MutationProbability: 0,
			Size:                1,
		},
	}

	world.Action(agents, 1, drawFrame)

}

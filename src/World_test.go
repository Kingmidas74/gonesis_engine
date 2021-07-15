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

	randCommandIndices := []int{0, 0, 0, 0, 1, 0, 1, 0, 0, 0}

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
			NextGenerationLine:  1,
			MutationProbability: 0,
			Size:                2,
		},
	}

	world.Action(agents, 50, 1, drawFrame)

}

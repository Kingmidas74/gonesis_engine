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
			Cells:           nil,
			LatitudesCount:  1024,
			LongitudesCount: 1024,
		},
		populationController: Population{
			NextGenerationLine:  8,
			MutationProbability: 0,
			Size:                64,
		},
	}
	livingAgents := world.evaluateAgents(agents)
	if len(livingAgents) != world.populationController.Size {
		t.Errorf("evaluation failed")
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
				CommandList: []Command{moveCommand, eatCommand},
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
	for currentLatitude := 0; currentLatitude < terrain.LatitudesCount; currentLatitude++ {
		for currentLongitude := 0; currentLongitude < terrain.LongitudesCount; currentLongitude++ {
			currentCell := terrain.GetCell(currentLatitude, currentLongitude)
			fmt.Printf("%2d %2d | ", currentCell.CellType, currentCell.Cost)
		}
		println()
	}

	world.runDay(agents, 3)

}

func TestWorld_Action(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	randCommandIndices := []int{3, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 2, 2, 1, 1, 1, 2, 2, 3}

	agents := make([]Agent, 0)
	for i := 0; i < 1; i++ {
		rand.Shuffle(len(randCommandIndices), func(i, j int) {
			randCommandIndices[i], randCommandIndices[j] = randCommandIndices[j], randCommandIndices[i]
		})
		agents = append(agents, Agent{
			Energy: 20,
			Brain: Brain{
				CommandList: []Command{waitCommand, moveCommand, eatCommand},
				Commands:    randCommandIndices,
			},
		})
		fmt.Printf("%v", randCommandIndices)
	}

	terrain := Terrain{}
	terrain.Generate(15, 15, agents, 30, 15)

	world := World{
		Terrain: terrain,
		populationController: Population{
			NextGenerationLine:  1,
			MutationProbability: 0,
			Size:                8,
		},
	}

	world.Action(agents, 3, 3, drawFrame)

}

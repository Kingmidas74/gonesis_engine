package gonesis

import (
	"fmt"
	"testing"
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
	agents := make([]Agent, 0)
	for i := 0; i < 2; i++ {
		agents = append(agents, Agent{
			Energy: 20,
			Brain: Brain{
				Commands: []Command{moveCommand, eatCommand},
			},
		})
	}

	terrain := Terrain{}
	terrain.Generate(5, 5, 30, 10)

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

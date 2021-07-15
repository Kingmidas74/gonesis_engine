package gonesis

import (
	"errors"
)

type Population struct {
	NextGenerationLine  int
	MutationProbability float64
	Size                int
}

func (this *Population) Reproduction(parents []Agent) (bool, []Agent, error) {
	if len(parents) == 0 {
		return false, nil, errors.New("not enough agents")
	} else if len(parents) > this.NextGenerationLine {
		return false, parents, nil
	}

	newGeneration := make([]Agent, 0)

	parents = this.ReCreateMinimumParents(parents)

	for i := 0; i < len(parents); i++ {
		newGeneration = append(newGeneration, parents[i])
	}

	for len(newGeneration) < this.Size {
		parentIndex := randomIntBetween(0, len(parents)-1)
		child := this.makeNewAgent(parents[parentIndex])
		newGeneration = append(newGeneration, child)
	}

	return true, newGeneration, nil
}

func (this *Population) ReCreateMinimumParents(agents []Agent) []Agent {
	for len(agents) < this.NextGenerationLine {
		parentIndex := randomIntBetween(0, len(agents)-1)

		parent := this.makeNewAgent(agents[parentIndex])

		agents = append(agents, parent)
	}
	return agents
}

func (this *Population) makeNewAgent(parent Agent) Agent {
	child := parent
	child.Energy = 20
	child.CurrentAddress = 0
	child.Generation = parent.Generation + 1
	return child
}

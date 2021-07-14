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
		return false, parents, errors.New("there are all dead")
	}

	if len(parents) >= this.NextGenerationLine {
		return false, parents, nil
	}

	newGeneration := make([]Agent, 0)

	parents = this.reCreateMinimumParents(parents)

	for i := 0; i < len(parents); i++ {
		newGeneration = append(newGeneration, parents[i])
	}

	for len(newGeneration) < this.Size {
		parentIndex := randomIntBetween(0, len(parents)) - 1
		if parentIndex < 0 || parentIndex > len(parents)-1 {
			parentIndex = 0
		}
		child := this.makeNewAgent(parents[parentIndex], true)
		newGeneration = append(newGeneration, child)
	}

	return true, newGeneration, nil
}

func (this *Population) reCreateMinimumParents(agents []Agent) []Agent {
	for len(agents) < this.NextGenerationLine {
		parentIndex := randomIntBetween(0, len(agents)) - 1
		if parentIndex < 0 || parentIndex > len(agents)-1 {
			parentIndex = 0
		}
		parent := this.makeNewAgent(agents[parentIndex], false)
		agents = append(agents, parent)
	}
	return agents
}

func (this *Population) makeNewAgent(parent Agent, isChild bool) Agent {
	child := parent
	if isChild {
		child.Energy = 20
		child.currentCommandAddress = 0
		child.Generation = parent.Generation + 1
	}
	return child
}

package gonesis

import "math/rand"

type Population struct {
	NextGenerationLine  int
	MutationProbability float64
	Size                int
}

func (this *Population) Reproduction(parents []Agent) []Agent {
	if len(parents) >= this.NextGenerationLine {
		return parents
	}

	newGeneration := make([]Agent, 0)

	parents = this.reCreateMinimumParents(parents)

	for len(newGeneration) < this.Size {
		parentIndex := rand.Intn(len(parents))
		child := this.makeNewAgent(parents[parentIndex])
		newGeneration = append(newGeneration, child)
	}

	return newGeneration
}

func (this *Population) reCreateMinimumParents(agents []Agent) []Agent {
	for len(agents) < this.NextGenerationLine {
		parentIndex := rand.Intn(len(agents))
		parent := this.makeNewAgent(agents[parentIndex])
		agents = append(agents, parent)
	}
	return agents
}

func (this *Population) makeNewAgent(parent Agent) Agent {
	child := parent
	return child
}

package gonesis

import (
	"math/rand"
	"testing"
	"time"
)

func TestMakeNewAgent(t *testing.T) {
	population := Population{}
	agent := Agent{}
	agent.Energy = 3
	newAgent := population.makeNewAgent(agent, false)
	newAgent.Energy = 2
	if agent.Energy == newAgent.Energy {
		t.Errorf("Agent was reproduce wrong! Link is same")
	}
}

func TestReCreateMinimumParents(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	population := Population{}
	population.NextGenerationLine = 8
	agent := Agent{}
	agent.Energy = 3
	agents := make([]Agent, 0)
	agents = append(agents, agent)
	firstColonists := population.reCreateMinimumParents(agents)
	if len(firstColonists) != population.NextGenerationLine {
		t.Errorf("Not enough colonists")
	}
}

func TestReproduction(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	population := Population{}
	population.NextGenerationLine = 8
	population.Size = 64
	agent := Agent{}
	agent.Energy = 3
	agents := make([]Agent, 0)
	agents = append(agents, agent)
	_, firstColonists, err := population.Reproduction(agents)
	if len(firstColonists) != population.Size {
		t.Errorf("Not enough colonists")
	} else if err != nil {
		t.Errorf(err.Error())
	}
}

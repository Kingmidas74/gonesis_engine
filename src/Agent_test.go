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

	agent.NextDay(&World{}, 10)

	if agent.IsAlive() {
		t.Errorf("Agent shouldn't be alive with negative energy")
	}
}

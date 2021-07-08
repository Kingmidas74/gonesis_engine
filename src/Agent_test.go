package gonesis

import "testing"

func TestIsAlive(t *testing.T) {
	agent := Agent{}
	agent.Energy = -1

	isAlive := agent.isAlive()

	if isAlive {
		t.Errorf("Agent shouldn't be alive with negative energy")
	}
}

func TestNextDay(t *testing.T) {
	agent := Agent{}
	agent.Energy = 3

	for i := 0; i < 4; i++ {
		command := Command{}
		command.Identifier = i
		command.IsInterrupts = i == 2
		command.Handler = func(world *World, currentAgent *Agent) {
			currentAgent.MoveAddressOn(1)
		}

		agent.Commands = append(agent.Commands, command)
	}

	agent.nextDay(&World{}, 10)

	isAlive := agent.isAlive()

	if isAlive {
		t.Errorf("Agent shouldn't be alive with negative energy")
	}
}

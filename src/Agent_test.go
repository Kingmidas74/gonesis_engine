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

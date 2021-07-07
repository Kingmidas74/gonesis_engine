package gonesis

import (
	"fmt"
	"testing"
)

func TestCodeChain(t *testing.T) {

	agentsCount := 2


	agents := make([]Agent,0)


	for i := 0; i < agentsCount; i++ {
		currentCode := Code{}
		currentCode.Commands = CommandListConstructor()
		currentCode.Commands.AppendCommand(CommandConstructor(1,true, func(agent *Agent) bool {
			println("u",agent.X, agent.Y, agent.energy)
			agent.code.jumpTo(1)
			return true
		}))
		currentCode.Commands.AppendCommand(CommandConstructor(2,true, func(agent *Agent) bool {
			println("ur",agent.X, agent.Y, agent.energy)
			agent.code.jumpTo(1)
			return true
		}))
		currentCode.Commands.AppendCommand(CommandConstructor(6,true, func(agent *Agent) bool {
			println("d",agent.X, agent.Y, agent.energy)
			agent.code.jumpTo(1)
			return true
		}))
		currentCode.Commands.AppendCommand(CommandConstructor(8,false, func(agent *Agent) bool {
			println("eat", agent.X, agent.Y, agent.energy)
			agent.code.jumpTo(2)
			return true
		}))
		currentCode.currentAddress = 0
		agents = append(agents, AgentConstructor(currentCode,1,i,i, up,1))
	}

	for i, agent := range agents {
		fmt.Printf("agent %d", i)
		agent.action(10)
	}
}

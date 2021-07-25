package agents

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/core/primitives"
)

type Agent struct {
	contracts.IBrain
	primitives.Coords
	contracts.IReproduction

	Energy int

	ReproductionPower int
	Generation        int
	Age               int
}

func (this *Agent) GetGeneration() int {
	return this.Generation
}

func (this *Agent) SetGeneration(generation int) {
	this.Generation = generation
}

func (this *Agent) GetEnergy() int {
	return this.Energy
}

func (this *Agent) SetEnergy(energy int) {
	this.Energy = energy
}

func (this *Agent) GetBrain() contracts.IBrain {
	return this.IBrain
}

func (this *Agent) SetBrain(brain contracts.IBrain) {
	this.IBrain = brain
}

func (this *Agent) GetAge() int {
	return this.Age
}

func (this *Agent) IsAlive() bool {
	return this.Energy > 0
}

func (this *Agent) NextDay(terrain contracts.ITerrain, maxSteps int) {
	this.Age++
	for step := 0; step < maxSteps && this.Energy > 0; step++ {
		this.Energy--
		command := this.GetCurrentCommand()
		if command == nil {
			continue
		}
		command.Handle(terrain, this)
		if command.IsInterrupts() {
			break
		}
	}
	this.TryToReproduce(this, terrain)
}
func (this *Agent) MakeChild() contracts.IAgent {
	return &Agent{
		IReproduction: this.IReproduction,
	}
}

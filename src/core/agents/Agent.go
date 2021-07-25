package agents

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/core"
	"github.com/Kingmidas74/gonesis/core/primitives"
)

type Agent struct {
	core.Brain
	primitives.Coords

	Energy int

	ReproductionPower int
	Generation        int
}

func (this *Agent) GetGeneration() int {
	return this.Generation
}

func (this *Agent) SetGeneration(generation int) {
	this.Generation = generation
}

func (this *Agent) GetX() int {
	return this.X
}

func (this *Agent) SetX(x int) {
	this.X = x
}

func (this *Agent) GetY() int {
	return this.Y
}

func (this *Agent) SetY(y int) {
	this.Y = y
}

func (this *Agent) GetEnergy() int {
	return this.Energy
}

func (this *Agent) SetEnergy(energy int) {
	this.Energy = energy
}

func (this *Agent) GetBrain() contracts.IBrain {
	return &this.Brain
}

func (this *Agent) IsAlive() bool {
	return this.Energy > 0
}

func (this *Agent) NextDay(terrain contracts.ITerrain, maxSteps int) {
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
	this.tryToReproduce(terrain)
}

func (this *Agent) tryToReproduce(terrain contracts.ITerrain) {
	if child, freeCell := this.makeChild(), this.findEmptyNeighborCell(terrain); child != nil && freeCell != nil {
		freeCell.SetCellType(contracts.LockedCell)
		child.SetX(freeCell.GetX())
		child.SetY(freeCell.GetY())
		freeCell.SetAgent(child)

	}
}

func (this *Agent) makeChild() *Agent {
	if this.Energy >= this.ReproductionPower {
		child := &Agent{
			Brain:             this.Brain,
			Energy:            this.ReproductionPower,
			ReproductionPower: this.ReproductionPower,
			Generation:        this.Generation + 1,
		}
		child.CurrentAddress = 0

		this.Energy -= this.ReproductionPower
		return child
	}
	return nil
}

func (this *Agent) findEmptyNeighborCell(terrain contracts.ITerrain) contracts.ICell {
	cells := terrain.GetNeighbors(this.X, this.Y)

	for i := 0; i < len(cells); i++ {
		if cell := cells[i]; cell.GetCellType() == contracts.EmptyCell {
			return cell
		}
	}
	return nil
}

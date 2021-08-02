package terrains

import (
	"github.com/Kingmidas74/gonesis_engine/contracts"
	"github.com/Kingmidas74/gonesis_engine/core/primitives"
)

type Cell struct {
	primitives.Coords

	CellType contracts.CellType
	Cost     int
	Agent    contracts.IAgent
}

func (this *Cell) GetCellType() contracts.CellType {
	return this.CellType
}

func (this *Cell) SetCellType(cellType contracts.CellType) {
	this.CellType = cellType
}

func (this *Cell) GetAgent() contracts.IAgent {
	return this.Agent
}

func (this *Cell) SetAgent(agent contracts.IAgent) {
	this.Agent = agent
}

func (this *Cell) GetCost() int {
	return this.Cost
}

func (this *Cell) SetCost(cost int) {
	this.Cost = cost
}

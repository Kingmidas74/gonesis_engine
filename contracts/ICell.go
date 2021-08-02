package contracts

type ICell interface {
	GetCellType() CellType
	SetCellType(cellType CellType)

	GetAgent() IAgent
	SetAgent(agent IAgent)

	SetX(x int)
	GetX() int

	SetY(y int)
	GetY() int

	GetCost() int
	SetCost(cost int)
}

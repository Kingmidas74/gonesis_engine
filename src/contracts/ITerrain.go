package contracts

type CellType int

const (
	EmptyCell    CellType = 0
	LockedCell   CellType = 1
	OrganicCell  CellType = 2
	ObstacleCell CellType = 3
)

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

type ITerrain interface {
	GetWidth() int
	GetHeight() int
	GetCells() []ICell

	GetCell(x, y int) ICell

	GetNeighbor(x, y int, direction int) ICell
	GetNeighbors(x, y int) []ICell
}

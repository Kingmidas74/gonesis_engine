package contracts

type CellType int

const (
	EmptyCell    CellType = 0
	LockedCell   CellType = 1
	OrganicCell  CellType = 2
	ObstacleCell CellType = 3
)

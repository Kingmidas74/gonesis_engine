package gonesis

type CellType byte

const (
	EmptyCell    = 0
	LockedCell   = 1
	OrganicCell  = 2
	ObstacleCell = 3
)

type Cell struct {
	Coords
	CellType CellType
	Cost     int
	Agent    *Agent
}

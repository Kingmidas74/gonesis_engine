package gonesis

type CellType byte

const (
	Empty    = 0
	Locked   = 1
	Organic  = 2
	Obstacle = 3
)

type Cell struct {
	Coords
	CellType CellType
	Cost     int
	Agent    *Agent
}

package gonesis

type CellType byte

const (
	Empty    = 0
	Organic  = 1
	Obstacle = 2
)

type Cell struct {
	CellType CellType
	Cost     int
	Agent    *Agent
}

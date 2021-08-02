package contracts

type ITerrain interface {
	GetWidth() int
	GetHeight() int
	GetCells() []ICell

	GetCell(x, y int) ICell

	GetNeighbor(x, y int, direction int) ICell
	GetNeighbors(x, y int) []ICell
}

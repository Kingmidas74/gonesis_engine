package contracts

type MazeBuilder[G MazeGenerator] interface {
	SetWidth(width int) MazeBuilder[G]
	SetHeight(height int) MazeBuilder[G]
	FirstFilled(flag bool) MazeBuilder[G]
	SetRequiredEmptyCells(requiredEmptyCells int) MazeBuilder[G]
	Build() (Maze, error)
}

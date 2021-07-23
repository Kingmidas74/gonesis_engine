package gonesis

type Terrain struct {
	Cells  []Cell
	Width  int
	Height int
}

func (this *Terrain) transformX(x int) int {
	return modLikePython(x, this.Width)
}

func (this *Terrain) transformY(y int) int {
	return modLikePython(y, this.Height)
}

func (this *Terrain) getCellIndex(x, y int) int {
	x = this.transformX(x)
	y = this.transformY(y)
	return y*this.Width + x
}

func (this *Terrain) GetCell(x, y int) *Cell {
	return &this.Cells[this.getCellIndex(x, y)]
}

func (this *Terrain) SetCell(x, y int, cell Cell) {
	this.Cells[this.getCellIndex(x, y)] = cell
}

func (this *Terrain) Generate(width, height int, agents []Agent, organicProbability int, emptyCellsCount int) {
	this.Width, this.Height = width, height

	this.Cells = make([]Cell, width*height)
	organicCellsCount := 0

	if agents != nil {
		emptyCellsCount += len(agents)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			currentCell := Cell{
				Coords: Coords{
					X: x,
					Y: y,
				},
			}

			currentOrganicProbability := randomIntBetween(1, 100)
			if organicCellsCount < width*height-(emptyCellsCount) && organicProbability > currentOrganicProbability {
				currentCell.CellType = Organic
				currentCell.Cost = randomIntBetween(1, 10)
				organicCellsCount++
			}
			this.SetCell(x, y, currentCell)
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == 0 || y == 0 || x == width-1 || y == height-1 {
				currentCell := this.GetCell(x, y)
				currentCell.CellType = 0
				currentCell.Cost = 0
			}
		}
	}

	this.placeAgents(agents)
}

func (this *Terrain) placeAgents(agents []Agent) {
	if agents == nil {
		return
	}

	for i := 0; i < len(this.Cells); i++ {
		if this.Cells[i].CellType == Locked || this.Cells[i].Agent != nil {
			this.Cells[i].Agent = nil
			this.Cells[i].CellType = Empty
		}
	}

	placedAgentsCount := 0

	for placedAgentsCount < len(agents) {
		for y := 0; y < this.Height; y++ {
			for x := 0; x < this.Width; x++ {
				if (x == 0 || y == 0 || x == this.Width-1 || y == this.Height-1) && randomIntBetween(0, 101) > 50 && placedAgentsCount < len(agents) {
					currentCell := this.GetCell(x, y)

					agent := &agents[placedAgentsCount]
					currentCell.Agent = agent
					agent.X = x
					agent.Y = y
					currentCell.CellType = Locked
					placedAgentsCount++

				}
			}
		}
	}

}

func (this *Terrain) cleanDeath() {
	for y := 0; y < this.Height; y++ {
		for x := 0; x < this.Width; x++ {
			currentCell := this.GetCell(x, y)
			if currentCell.Agent != nil && !currentCell.Agent.IsAlive() {
				currentCell.Agent = nil
				currentCell.CellType = Empty
			}
		}
	}
}

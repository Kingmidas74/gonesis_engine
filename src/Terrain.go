package gonesis

type Terrain struct {
	Cells           []Cell
	LatitudesCount  int
	LongitudesCount int
}

func (this *Terrain) transformLatitude(latitude int) int {
	return modLikePython(latitude, this.LatitudesCount)
}

func (this *Terrain) transformLongitude(longitude int) int {
	return modLikePython(longitude, this.LongitudesCount)
}

func (this *Terrain) getCellIndex(latitude, longitude int) int {
	latitude = this.transformLatitude(latitude)
	longitude = this.transformLongitude(longitude)
	return latitude*this.LongitudesCount + longitude
}

func (this *Terrain) GetCell(latitude, longitude int) *Cell {
	return &this.Cells[this.getCellIndex(latitude, longitude)]
}

func (this *Terrain) SetCell(latitude, longitude int, cell Cell) {
	this.Cells[this.getCellIndex(latitude, longitude)] = cell
}

func (this *Terrain) Generate(latitudesCount, longitudesCount int, agents []Agent, organicProbability int, emptyCellsCount int) {
	this.LatitudesCount, this.LongitudesCount = latitudesCount, longitudesCount

	this.Cells = make([]Cell, latitudesCount*longitudesCount)
	organicCellsCount := 0

	if agents != nil {
		emptyCellsCount += len(agents)
	}

	for currentLatitude := 0; currentLatitude < latitudesCount; currentLatitude++ {
		for currentLongitude := 0; currentLongitude < longitudesCount; currentLongitude++ {
			currentCell := Cell{}

			currentOrganicProbability := randomIntBetween(1, 100)
			if organicCellsCount < latitudesCount*longitudesCount-(emptyCellsCount) && organicProbability > currentOrganicProbability {
				currentCell.CellType = Organic
				currentCell.Cost = randomIntBetween(1, 10)
				organicCellsCount++
			}
			this.SetCell(currentLatitude, currentLongitude, currentCell)
		}
	}

	this.placeAgents(agents)
}

func (this *Terrain) placeAgents(agents []Agent) {
	if agents == nil {
		return
	}

	for i := 0; i < this.LatitudesCount*this.LongitudesCount; i++ {
		if this.Cells[i].CellType == Locked {
			this.Cells[i].Agent = nil
			this.Cells[i].CellType = Empty
		}
	}

	placedAgentsCount := 0
	for currentLatitude := 0; currentLatitude < this.LatitudesCount; currentLatitude++ {
		for currentLongitude := 0; currentLongitude < this.LongitudesCount; currentLongitude++ {
			if currentCell := this.GetCell(currentLatitude, currentLongitude); currentCell.CellType == Empty && placedAgentsCount < len(agents) {
				agent := &agents[placedAgentsCount]
				currentCell.Agent = agent
				agent.Latitude = currentLatitude
				agent.Longitude = currentLongitude
				currentCell.CellType = Locked
				placedAgentsCount++
			}
		}
	}
}

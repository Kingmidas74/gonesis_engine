package terrainGenerators

import (
	"github.com/Kingmidas74/gonesis_engine/contracts"
	"github.com/Kingmidas74/gonesis_engine/core/primitives"
	"github.com/Kingmidas74/gonesis_engine/core/terrains"
)

type TerrainGenerator struct {
}

func (this *TerrainGenerator) GenerateMatrix(width, height int) []bool {
	result := make([]bool, 0)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			result = append(result, false)
		}
	}
	result[0] = true
	return result
}

func (this *TerrainGenerator) GenerateTerrain(terrainType contracts.TerrainType, matrix []bool, width, height int) contracts.ITerrain {

	cells := make([]contracts.ICell, 0)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			currentCell := &terrains.Cell{
				Coords: primitives.Coords{
					X: x,
					Y: y,
				},
				CellType: contracts.EmptyCell,
			}
			if matrix[y*width+x] == false {
				currentCell.SetCellType(contracts.ObstacleCell)
			}
			cells = append(cells, currentCell)
		}
	}

	terrain := terrains.Terrain{
		Cells:  cells,
		Width:  width,
		Height: height,
	}

	var result contracts.ITerrain

	switch terrainType {
	case contracts.MooreTerrain:
		result = &terrains.MooreTerrain{Terrain: terrain}
	case contracts.NeumannTerrain:
		result = &terrains.NeumannTerrain{Terrain: terrain}
	case contracts.HexTerrain:
		result = &terrains.HexTerrain{Terrain: terrain}
	}
	return result
}

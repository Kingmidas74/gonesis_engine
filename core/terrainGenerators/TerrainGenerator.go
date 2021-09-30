package terrainGenerators

import (
	"github.com/Kingmidas74/gonesis_engine/contracts"
	"github.com/Kingmidas74/gonesis_engine/core/primitives"
	"github.com/Kingmidas74/gonesis_engine/core/terrains"
)

type TerrainGenerator struct {
}

func (this *TerrainGenerator) GenerateMatrix(width, height int) []bool {
	result := make([]bool, width*height)

	for y := 0; y < height; y = y + 2 {
		for x := 0; x < width; x = x + 2 {
			result[y*width+x] = true
		}
	}
	result[0] = true
	return result
}

func (this *TerrainGenerator) GenerateTerrain(terrainType contracts.TerrainType, matrix []bool, width, height int) contracts.ITerrain {

	terrain := terrains.Terrain{
		Cells:  this.TransformMatrixToCells(matrix, width, height),
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

func (this *TerrainGenerator) TransformMatrixToCells(matrix []bool, width, height int) []contracts.ICell {

	cells := make([]contracts.ICell, width*height)

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
			if x+1 == width || y+1 == height || x == 0 || y == 0 {
				currentCell.SetCellType(contracts.EmptyCell)
			}
			cells[y*width+x] = currentCell
		}
	}

	return cells
}

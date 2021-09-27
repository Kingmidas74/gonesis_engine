package terrainGenerators

import (
	"github.com/Kingmidas74/gonesis_engine/contracts"
	"github.com/Kingmidas74/gonesis_engine/core/primitives"
	"github.com/Kingmidas74/gonesis_engine/core/terrains"
	"github.com/Kingmidas74/gonesis_engine/utils"
)

type TerrainGenerator struct {
}

func (this *TerrainGenerator) GenerateMatrix(width, height int) []bool {
	result := make([]bool, 0)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			result = append(result, utils.ModLikePython(x, 2) == 0 && utils.ModLikePython(y, 2) == 0)
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
			if x+1 == width || y+1 == height || x == 0 || y == 0 {
				currentCell.SetCellType(contracts.EmptyCell)
			}
			cells = append(cells, currentCell)
		}
	}

	return cells
}

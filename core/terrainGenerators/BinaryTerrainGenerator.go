package terrainGenerators

import (
	"github.com/Kingmidas74/gonesis_engine/contracts"
	"github.com/Kingmidas74/gonesis_engine/utils"
)

type BinaryTerrainGenerator struct {
	TerrainGenerator
}

func (this *BinaryTerrainGenerator) Generate(terrainType contracts.TerrainType, width, height int) contracts.ITerrain {

	matrix := this.GenerateMatrix(width, height)

	for y := 0; y < height; y = y + 2 {
		for x := 0; x < width; x = x + 2 {

			if y == 0 {
				if x+1 < width {
					matrix[y*width+x+1] = true
				}
				continue
			}

			direction := utils.RandomIntBetween(0, 1)
			if direction == 0 {
				matrix[(y-1)*width+x] = true
				continue
			}

			if x+1 >= width {
				continue
			}

			matrix[y*width+x+1] = true
		}
	}

	return this.GenerateTerrain(terrainType, matrix, width, height)
}

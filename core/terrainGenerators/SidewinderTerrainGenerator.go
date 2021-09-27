package terrainGenerators

import (
	"github.com/Kingmidas74/gonesis_engine/contracts"
	"github.com/Kingmidas74/gonesis_engine/utils"
)

type SidewinderTerrainGenerator struct {
	TerrainGenerator
}

func (this *SidewinderTerrainGenerator) Generate(terrainType contracts.TerrainType, width, height int) contracts.ITerrain {

	matrix := this.GenerateMatrix(width, height)

	for y := 0; y < height; y = y + 2 {
		runset := make([]int, 0)
		for x := 0; x < width; x = x + 2 {

			if y == 0 {
				if x+1 < width {
					matrix[y*width+x+1] = true
				}
				continue
			}

			runset = append(runset, x)

			direction := utils.RandomIntBetween(0, 1)

			if direction == 1 && x+1 < width {
				matrix[y*width+x+1] = true
				continue
			}

			randX := runset[utils.RandomIntBetween(0, len(runset)-1)]
			matrix[(y-1)*width+randX] = true
			runset = make([]int, 0)
		}
	}

	return this.GenerateTerrain(terrainType, matrix, width, height)
}

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

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch utils.RandomIntBetween(0, 1) {
			case 0:
				if x+1 < width && y*width+x+1 < width*height {
					matrix[y*width+x+1] = true
				}
				break
			case 1:
				if y+1 < height && (y+1)*width+x < width*height {
					matrix[(y+1)*width+x] = true
				}
				break
			}
		}
	}

	return this.GenerateTerrain(terrainType, matrix, width, height)
}

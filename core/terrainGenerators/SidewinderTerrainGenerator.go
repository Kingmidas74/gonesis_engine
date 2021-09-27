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
	sx := 0
	sy := 0
	cx := sx

	for y := sy; y < height; y++ {
		for x := sx; x < width; x++ {
			if y != sy {
				if utils.RandomIntBetween(0, 1) == 0 && x+1 < width {
					matrix[y*width+x+1] = true
				} else {
					randX := utils.RandomIntBetween(cx, x)
					if y-1 >= 0 {
						matrix[(y-1)*width+randX] = true

						if x+1 < width {
							cx = x
						} else {
							cx = sx
						}
					}
				}
			} else {
				if x+1 < width {
					matrix[y*width+x+1] = true
				}
			}
		}
	}

	return this.GenerateTerrain(terrainType, matrix, width, height)
}

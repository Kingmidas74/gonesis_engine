package terrainGenerators

import (
	"github.com/Kingmidas74/gonesis_engine/contracts"
	"github.com/Kingmidas74/gonesis_engine/utils"
)

type AldousBroderTerrainGenerator struct {
	TerrainGenerator
}

func (this *AldousBroderTerrainGenerator) Generate(terrainType contracts.TerrainType, width, height int) contracts.ITerrain {

	matrix := this.GenerateMatrix(width, height)

	for i := range matrix {
		matrix[i] = false
	}
	matrix[0] = true

	actors := make([]int, 0)

	actorsCount := 7

	for i := 0; i < actorsCount; i++ {
		actors = append(actors, 0)
		actors = append(actors, 0)
	}

	directions := make(map[int][4]int)
	directions[0] = [4]int{0, -2, 0, 1}
	directions[1] = [4]int{2, 0, -1, 0}
	directions[2] = [4]int{0, 2, 0, -1}
	directions[3] = [4]int{-2, 0, 1, 0}

	isFinished := false

	for !isFinished {
		for i := 0; i < actorsCount; i++ {

			direction := directions[utils.RandomIntBetween(0, 3)]

			if actors[i*2]+direction[0] >= 0 && actors[i*2]+direction[0] < width && actors[i*2+1]+direction[1] >= 0 && actors[i*2+1]+direction[1] < height {

				actors[i*2] += direction[0]
				actors[i*2+1] += direction[1]

				if matrix[actors[i*2+1]*width+actors[i*2]] == false {
					matrix[actors[i*2+1]*width+actors[i*2]] = true
					matrix[(actors[i*2+1]+direction[3])*width+actors[i*2]+direction[2]] = true
				}
			}
		}

		isFinished = true
		for y := 0; y < height; y = y + 2 {
			for x := 0; x < width; x = x + 2 {
				isFinished = isFinished && matrix[y*width+x]
			}
		}
	}

	return this.GenerateTerrain(terrainType, matrix, width, height)
}

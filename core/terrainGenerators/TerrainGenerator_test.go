package terrainGenerators

import (
	"github.com/Kingmidas74/gonesis_engine/contracts"
	"math/rand"
	"testing"
	"time"
)

func TestBinaryTerrainGenerator_Generate(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	terrainGenerator := BinaryTerrainGenerator{}

	terrain := terrainGenerator.Generate(contracts.MooreTerrain, 5, 5)

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			print(terrain.GetCell(x, y).GetCellType())
			print(" ")
		}
		println()
	}

}

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

	size := 500

	terrain := terrainGenerator.Generate(contracts.MooreTerrain, size, size)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			print(terrain.GetCell(x, y).GetCellType())
			print(" ")
		}
		println()
	}

}

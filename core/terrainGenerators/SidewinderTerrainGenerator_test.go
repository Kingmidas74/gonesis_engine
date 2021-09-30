package terrainGenerators

import (
	"github.com/Kingmidas74/gonesis_engine/contracts"
	"math/rand"
	"testing"
	"time"
)

func TestSidewinderTerrainGenerator_Generate(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	terrainGenerator := SidewinderTerrainGenerator{}

	size := 101

	terrainGenerator.Generate(contracts.MooreTerrain, size, size)

}

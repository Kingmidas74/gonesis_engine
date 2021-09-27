package terrainGenerators

import (
	"github.com/Kingmidas74/gonesis_engine/contracts"
	"math/rand"
	"testing"
	"time"
)

func TestAldousBroderTerrainGenerator_Generate(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	terrainGenerator := AldousBroderTerrainGenerator{}

	size := 128

	terrainGenerator.Generate(contracts.MooreTerrain, size, size)

}

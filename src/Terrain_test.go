package gonesis

import (
	"math/rand"
	"testing"
	"time"
)

func TestTerrain_SetCell(t *testing.T) {
	latitudeCount := 5
	longitudeCount := 5
	organicProbability := 50
	emptyCellsCount := 1

	terrain := Terrain{}
	terrain.Generate(latitudeCount, longitudeCount, organicProbability, emptyCellsCount)

	terrain.SetCell(-2, 5, Cell{CellType: Obstacle})

	cellType := terrain.GetCell(-2, 5).CellType

	if cellType != Obstacle {
		t.Errorf("set cell works incorrect")
	}
}

func TestGenerate(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	agents := make([]Agent, 0)
	for i := 0; i < 2; i++ {
		agents = append(agents, Agent{
			Energy: 20,
		})
	}

	latitudeCount := 2560 * 2
	longitudeCount := 1440 * 2
	organicProbability := 0
	emptyCellsCount := 100

	terrain := Terrain{}
	terrain.Generate(latitudeCount, longitudeCount, organicProbability, emptyCellsCount)

	if len(terrain.Cells) != latitudeCount*longitudeCount {
		t.Errorf("terrain generate error")
	}

	freeSpace := 0
	for currentLatitude := 0; currentLatitude < latitudeCount; currentLatitude++ {
		for currentLongitude := 0; currentLongitude < longitudeCount; currentLongitude++ {
			if terrain.Cells[currentLatitude*longitudeCount+currentLongitude].CellType == Empty {
				if terrain.Cells[currentLatitude*longitudeCount+currentLongitude].Cost != 0 {
					t.Errorf("Free space shouldn't cost anything")
				}
				freeSpace++
			}
		}
	}

	if freeSpace < emptyCellsCount {
		t.Errorf("not enought free space")
	}

	drawFrame(&terrain, 1)
}

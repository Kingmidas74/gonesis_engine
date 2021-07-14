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
	terrain.Generate(latitudeCount, longitudeCount, nil, organicProbability, emptyCellsCount)

	terrain.SetCell(-2, 5, Cell{CellType: Obstacle})

	cellType := terrain.GetCell(-2, 5).CellType

	if cellType != Obstacle {
		t.Errorf("set cell works incorrect")
	}
}

func TestGenerateWithAgent_Success(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	agents := make([]Agent, 0)
	for i := 0; i < 2; i++ {
		agents = append(agents, Agent{
			Energy: 20,
		})
	}

	latitudeCount := 1024
	longitudeCount := 768
	organicProbability := 5
	emptyCellsCount := 20

	terrain := Terrain{}
	terrain.Generate(latitudeCount, longitudeCount, agents, organicProbability, emptyCellsCount)

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

	if &agents[0] != terrain.GetCell(agents[0].Latitude, agents[0].Longitude).Agent {
		t.Errorf("agents weren't linked")
	}

	drawFrame(&terrain, 1)
}

func TestGenerateWithoutAgents_Success(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	latitudeCount := 1024
	longitudeCount := 768
	organicProbability := 5
	emptyCellsCount := 20

	terrain := Terrain{}
	terrain.Generate(latitudeCount, longitudeCount, nil, organicProbability, emptyCellsCount)

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

	drawFrame(&terrain, 2)
}

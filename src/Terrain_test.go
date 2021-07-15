package gonesis

import (
	"math/rand"
	"testing"
	"time"
)

func TestTerrain_SetCell(t *testing.T) {
	width := 5
	height := 5
	organicProbability := 50
	emptyCellsCount := 1

	terrain := Terrain{}
	terrain.Generate(width, height, nil, organicProbability, emptyCellsCount)

	terrain.SetCell(-2, 5, Cell{CellType: Obstacle})

	cellType := terrain.GetCell(-2, 5).CellType

	if cellType != Obstacle {
		t.Errorf("set cell works incorrect")
	}
}

func TestGenerateWithAgent_Success(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	agents := make([]Agent, 0)
	for i := 0; i < 10; i++ {
		agents = append(agents, Agent{
			Energy: 20,
		})
	}

	width := 10
	height := 10
	organicProbability := 50
	emptyCellsCount := 20

	terrain := Terrain{}
	terrain.Generate(width, height, agents, organicProbability, emptyCellsCount)

	if len(terrain.Cells) != width*height {
		t.Errorf("terrain generate error")
	}

	freeSpace := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if currentCell := terrain.GetCell(x, y); currentCell.CellType == Empty {
				if currentCell.Cost != 0 {
					t.Errorf("Free space shouldn't cost anything")
				}
				freeSpace++
			}
		}
	}

	if freeSpace < emptyCellsCount {
		t.Errorf("not enought free space")
	}

	if &agents[0] != terrain.GetCell(agents[0].X, agents[0].Y).Agent {
		t.Errorf("agents weren't linked")
	}

	drawFrame(&terrain, 1)
}

func TestGenerateWithoutAgents_Success(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	width := 10
	height := 10
	organicProbability := 50
	emptyCellsCount := 20

	terrain := Terrain{}
	terrain.Generate(width, height, nil, organicProbability, emptyCellsCount)

	if len(terrain.Cells) != width*height {
		t.Errorf("terrain generate error")
	}

	freeSpace := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if currentCell := terrain.GetCell(x, y); currentCell.CellType == Empty {
				if currentCell.Cost != 0 {
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

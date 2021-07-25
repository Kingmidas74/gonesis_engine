package terrains

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/utils"
)

type NeumannTerrain struct {
	Terrain
}

func (this *NeumannTerrain) GetNeighbor(x, y int, direction int) contracts.ICell {
	direction = utils.ModLikePython(direction, 4)
	targetX := x
	targetY := y

	switch direction {
	case 0:
		targetY -= 1
	case 1:
		targetX += 1
	case 2:
		targetY += 1
	case 3:
		targetX -= 1
	}
	return this.GetCell(targetX, targetY)
}

func (this *NeumannTerrain) GetNeighbors(x, y int) []contracts.ICell {
	result := make([]contracts.ICell, 0)

	coords := []int{
		x, y - 1,

		x - 1, y,
		x + 1, y,

		x, y + 1,
	}
	for i := 0; i < len(coords); i = i + 2 {
		result = append(result, this.GetCell(coords[i], coords[i+1]))
	}
	return result
}

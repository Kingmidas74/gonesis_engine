package terrains

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/utils"
)

type MooreTerrain struct {
	Terrain
}

func (this *MooreTerrain) GetNeighbor(x, y int, direction int) contracts.ICell {
	direction = utils.ModLikePython(direction, 8)
	targetX := x
	targetY := y

	switch direction {
	case 0:
		targetY -= 1
	case 1:
		targetY -= 1
		targetX += 1
	case 2:
		targetX += 1
	case 3:
		targetX += 1
		targetY += 1
	case 4:
		targetY += 1
	case 5:
		targetX -= 1
		targetY += 1
	case 6:
		targetX -= 1
	case 7:
		targetX -= 1
		targetY -= 1
	}
	return this.GetCell(targetX, targetY)
}

func (this *MooreTerrain) GetNeighbors(x, y int) []contracts.ICell {
	result := make([]contracts.ICell, 0)

	coords := []int{
		x - 1, y - 1,
		x, y - 1,
		x + 1, y - 1,

		x - 1, y,
		x + 1, y,

		x - 1, y + 1,
		x, y + 1,
		x + 1, y + 1,
	}
	for i := 0; i < len(coords); i = i + 2 {
		result = append(result, this.GetCell(coords[i], coords[i+1]))
	}
	return result
}

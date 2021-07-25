package terrains

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/utils"
)

type NeumannDirection int

const (
	NeumannUp    NeumannDirection = 0
	NeumannRight NeumannDirection = 1
	NeumannDown  NeumannDirection = 2
	NeumannLeft  NeumannDirection = 3
)

type NeumannTerrain struct {
	Terrain
}

func (this *NeumannTerrain) GetNeighbor(x, y int, direction int) contracts.ICell {
	multiples := this.getCoordsMultiples()
	neumannDirection := NeumannDirection(utils.ModLikePython(direction, len(multiples)))
	return this.GetCell(x+multiples[neumannDirection][0], y+multiples[neumannDirection][1])
}

func (this *NeumannTerrain) GetNeighbors(x, y int) []contracts.ICell {
	result := make([]contracts.ICell, 0)

	for _, coords := range this.getCoordsMultiples() {
		result = append(result, this.GetCell(x+coords[0], y+coords[1]))
	}
	return result
}

func (this *NeumannTerrain) getCoordsMultiples() map[NeumannDirection][2]int {
	multiples := make(map[NeumannDirection][2]int)
	multiples[NeumannUp] = [2]int{0, -1}
	multiples[NeumannRight] = [2]int{1, 0}
	multiples[NeumannDown] = [2]int{0, 1}
	multiples[NeumannLeft] = [2]int{-1, 0}
	return multiples
}

package terrains

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/utils"
)

type NeumannDirection int

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
	const (
		Up    NeumannDirection = 0
		Right NeumannDirection = 1
		Down  NeumannDirection = 2
		Left  NeumannDirection = 3
	)

	multiples := make(map[NeumannDirection][2]int)
	multiples[Up] = [2]int{0, -1}
	multiples[Right] = [2]int{1, 0}
	multiples[Down] = [2]int{0, 1}
	multiples[Left] = [2]int{-1, 0}
	return multiples
}

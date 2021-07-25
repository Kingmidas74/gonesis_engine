package terrains

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/utils"
)

type HexDirection int

const (
	HexUpLeft    HexDirection = 0
	HexUpRight   HexDirection = 1
	HexRight     HexDirection = 2
	HexDownRight HexDirection = 3
	HexDownLeft  HexDirection = 4
	HexLeft      HexDirection = 5
)

type HexTerrain struct {
	Terrain
}

func (this *HexTerrain) GetNeighbor(x, y int, direction int) contracts.ICell {
	multiples := this.getCoordsMultiples()
	hexDirection := HexDirection(utils.ModLikePython(direction, len(multiples)))
	return this.GetCell(x+multiples[hexDirection][0], y+multiples[hexDirection][1])
}

func (this *HexTerrain) GetNeighbors(x, y int) []contracts.ICell {
	result := make([]contracts.ICell, 0)

	for _, coords := range this.getCoordsMultiples() {
		result = append(result, this.GetCell(x+coords[0], y+coords[1]))
	}
	return result
}

func (this *HexTerrain) getCoordsMultiples() map[HexDirection][2]int {
	multiples := make(map[HexDirection][2]int)
	multiples[HexUpLeft] = [2]int{-1, -1}
	multiples[HexUpRight] = [2]int{0, -1}
	multiples[HexRight] = [2]int{1, 0}
	multiples[HexDownRight] = [2]int{0, 1}
	multiples[HexDownLeft] = [2]int{-1, 1}
	multiples[HexLeft] = [2]int{-1, 0}
	return multiples
}

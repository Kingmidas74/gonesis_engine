package terrains

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/utils"
)

type HexDirection int

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
	const (
		UpLeft    HexDirection = 0
		UpRight   HexDirection = 1
		Right     HexDirection = 2
		DownRight HexDirection = 3
		DownLeft  HexDirection = 4
		Left      HexDirection = 5
	)

	multiples := make(map[HexDirection][2]int)
	multiples[UpLeft] = [2]int{-1, -1}
	multiples[UpRight] = [2]int{0, -1}
	multiples[Right] = [2]int{1, 0}
	multiples[DownRight] = [2]int{0, 1}
	multiples[DownLeft] = [2]int{-1, 1}
	multiples[Left] = [2]int{-1, 0}
	return multiples
}

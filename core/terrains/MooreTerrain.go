package terrains

import (
	"github.com/Kingmidas74/gonesis_engine/contracts"
	"github.com/Kingmidas74/gonesis_engine/utils"
)

type MooreDirection int

type MooreTerrain struct {
	Terrain
}

func (this *MooreTerrain) GetNeighbor(x, y int, direction int) contracts.ICell {
	multiples := this.getCoordsMultiples()
	mooreDirection := MooreDirection(utils.ModLikePython(direction, len(multiples)))
	return this.GetCell(x+multiples[mooreDirection][0], y+multiples[mooreDirection][1])
}

func (this *MooreTerrain) GetNeighbors(x, y int) []contracts.ICell {
	result := make([]contracts.ICell, 0)

	for _, coords := range this.getCoordsMultiples() {
		result = append(result, this.GetCell(x+coords[0], y+coords[1]))
	}
	return result
}

func (this *MooreTerrain) getCoordsMultiples() map[MooreDirection][2]int {
	const (
		Up        MooreDirection = 0
		UpRight   MooreDirection = 1
		Right     MooreDirection = 2
		DownRight MooreDirection = 3
		Down      MooreDirection = 4
		DownLeft  MooreDirection = 5
		Left      MooreDirection = 6
		UpLeft    MooreDirection = 7
	)

	multiples := make(map[MooreDirection][2]int)
	multiples[Up] = [2]int{0, -1}
	multiples[UpRight] = [2]int{1, -1}
	multiples[Right] = [2]int{1, 0}
	multiples[DownRight] = [2]int{1, 1}
	multiples[Down] = [2]int{0, 1}
	multiples[DownLeft] = [2]int{-1, 1}
	multiples[Left] = [2]int{-1, 0}
	multiples[UpLeft] = [2]int{-1, -1}
	return multiples
}

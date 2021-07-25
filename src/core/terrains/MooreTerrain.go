package terrains

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/utils"
)

type MooreDirection int

const (
	MooreUp        MooreDirection = 0
	MooreUpRight   MooreDirection = 1
	MooreRight     MooreDirection = 2
	MooreDownRight MooreDirection = 3
	MooreDown      MooreDirection = 4
	MooreDownLeft  MooreDirection = 5
	MooreLeft      MooreDirection = 6
	MooreUpLeft    MooreDirection = 7
)

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
	multiples := make(map[MooreDirection][2]int)
	multiples[MooreUp] = [2]int{0, -1}
	multiples[MooreUpRight] = [2]int{1, -1}
	multiples[MooreRight] = [2]int{1, 0}
	multiples[MooreDownRight] = [2]int{1, 1}
	multiples[MooreDown] = [2]int{0, 1}
	multiples[MooreDownLeft] = [2]int{-1, 1}
	multiples[MooreLeft] = [2]int{-1, 0}
	multiples[MooreUpLeft] = [2]int{-1, -1}
	return multiples
}

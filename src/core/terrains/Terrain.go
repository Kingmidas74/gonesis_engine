package terrains

import (
	"github.com/Kingmidas74/gonesis/contracts"
	"github.com/Kingmidas74/gonesis/utils"
)

type Terrain struct {
	Cells  []contracts.ICell
	Width  int
	Height int
}

func (this *Terrain) transformX(x int) int {
	return utils.ModLikePython(x, this.Width)
}

func (this *Terrain) transformY(y int) int {
	return utils.ModLikePython(y, this.Height)
}

func (this *Terrain) getCellIndex(x, y int) int {
	x = this.transformX(x)
	y = this.transformY(y)
	return y*this.Width + x
}

func (this *Terrain) GetCell(x, y int) contracts.ICell {
	return this.Cells[this.getCellIndex(x, y)]
}

func (this *Terrain) GetCells() []contracts.ICell {
	return this.Cells
}

func (this *Terrain) GetWidth() int {
	return this.Width
}

func (this *Terrain) GetHeight() int {
	return this.Height
}

package gonesis

type Agent struct {
	Coords
	Brain

	Energy int

	ReproductionPower int
	Generation        int
}

func (this *Agent) NextDay(world *World, maxSteps int) {
	for step := 0; step < maxSteps && this.Energy > 0; step++ {
		this.Energy--
		command := this.GetCurrentCommand()
		if command == nil {
			continue
		}
		command.Handler(world, this)
		if command.IsInterrupts {
			break
		}
	}
	this.TryToReproduce(world)
}

func (this *Agent) TryToReproduce(world *World) {
	if child, freeCell := this.MakeChild(), this.FindEmptyNeighborCell(world); child != nil && freeCell != nil {
		freeCell.CellType = LockedCell
		freeCell.Agent = child
		child.X = freeCell.X
		child.Y = freeCell.Y
	}
}

func (this *Agent) IsAlive() bool {
	return this.Energy > 0
}

func (this *Agent) MakeChild() *Agent {
	if this.Energy >= this.ReproductionPower {
		child := *this
		child.CurrentAddress = 0
		child.Energy = this.ReproductionPower
		child.Generation = this.Generation + 1
		this.Energy -= this.ReproductionPower
		return &child
	}
	return nil
}

func (this *Agent) FindEmptyNeighborCell(world *World) *Cell {
	coords := []int{
		this.Coords.X - 1, this.Coords.Y - 1,
		this.Coords.X, this.Coords.Y - 1,
		this.Coords.X + 1, this.Coords.Y - 1,

		this.Coords.X - 1, this.Coords.Y,
		this.Coords.X + 1, this.Coords.Y,

		this.Coords.X - 1, this.Coords.Y + 1,
		this.Coords.X, this.Coords.Y + 1,
		this.Coords.X + 1, this.Coords.Y + 1,
	}

	for i := 0; i < len(coords); i = i + 2 {
		if cell := world.GetCell(coords[i], coords[i+1]); cell.CellType == EmptyCell {
			return cell
		}
	}
	return nil
}

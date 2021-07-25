package contracts

type IBrain interface {
	GetCommandIdentifier(address int) int
	GetCurrentAddress() int
	MoveAddressOn(delta int)
}

type IAgent interface {
	GetGeneration() int
	SetGeneration(generation int)

	SetX(x int)
	GetX() int

	SetY(y int)
	GetY() int

	SetEnergy(energy int)
	GetEnergy() int

	NextDay(terrain ITerrain, maxSteps int)
	IsAlive() bool

	GetBrain() IBrain
}

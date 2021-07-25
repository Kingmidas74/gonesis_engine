package contracts

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
	GetAge() int
}

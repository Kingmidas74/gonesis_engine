package contracts

type IWorld interface {
	Action(maxSteps int, callback func(ITerrain, int))
}

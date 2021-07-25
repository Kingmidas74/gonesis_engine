package contracts

type IBrain interface {
	GetCommandIdentifier(address int) int
	GetCurrentAddress() int
	SetCurrentAddress(address int)
	MoveAddressOn(delta int)
	GetCurrentCommand() ICommand

	Clone() IBrain
	Mutate()
}

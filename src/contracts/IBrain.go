package contracts

type IBrain interface {
	GetCommandIdentifier(address int) int
	GetCurrentAddress() int
	MoveAddressOn(delta int)
}

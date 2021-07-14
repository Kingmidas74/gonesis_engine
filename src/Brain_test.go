package gonesis

import "testing"

func TestMoveAddressOn_SuccessWheDeltaPositive(t *testing.T) {
	brain := Brain{}
	for i := 0; i < 6; i++ {
		command := Command{}
		brain.CommandList = append(brain.CommandList, command)
	}
	brain.MoveAddressOn(1)
	brain.MoveAddressOn(2)
	if brain.currentCommandAddress != 3 {
		t.Errorf("Wrong move address %d", brain.currentCommandAddress)
	}
}

func TestMoveAddressOn_SuccessWheDeltaNegative(t *testing.T) {
	brain := Brain{}
	for i := 0; i < 6; i++ {
		command := Command{}
		brain.CommandList = append(brain.CommandList, command)
	}
	brain.MoveAddressOn(2)
	brain.MoveAddressOn(-4)
	if brain.currentCommandAddress != 4 {
		t.Errorf("Wrong move address %d", brain.currentCommandAddress)
	}
}

func TestSetAddress_SuccessWhenAddressBiggerThanLength(t *testing.T) {
	brain := Brain{}
	for i := 0; i < 6; i++ {
		command := Command{}
		brain.CommandList = append(brain.CommandList, command)
	}
	brain.SetAddress(8)
	if brain.currentCommandAddress != 2 {
		t.Errorf("Wrong set address %d", brain.currentCommandAddress)
	}
}

func TestSetAddress_SuccessWhenAddressLessZero(t *testing.T) {
	brain := Brain{}
	for i := 0; i < 6; i++ {
		command := Command{}
		brain.CommandList = append(brain.CommandList, command)
	}
	brain.SetAddress(-1)
	if brain.currentCommandAddress != 5 {
		t.Errorf("Wrong set address %d", brain.currentCommandAddress)
	}
}

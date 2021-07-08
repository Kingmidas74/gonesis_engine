package gonesis

import "testing"

func TestMoveAddressOn(t *testing.T) {
	brain := Brain{}
	for i := 0; i < 6; i++ {
		command := Command{}
		command.Identifier = i
		brain.Commands = append(brain.Commands, command)
	}
	brain.CurrentCommandAddress = 4
	brain.MoveAddressOn(2)
	if brain.CurrentCommandAddress != 0 {
		t.Errorf("%d", brain.CurrentCommandAddress)
	}
}

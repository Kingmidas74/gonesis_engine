package gonesis

import "testing"

func TestBrain_SetAddress_SuccessWhenAddressLessZero(t *testing.T) {
	commandsCount := 6
	delta := -1
	brain := Brain{}
	for i := 0; i < commandsCount; i++ {
		brain.Commands = append(brain.Commands, i)
	}
	brain.SetAddress(delta)
	if brain.CurrentAddress != modLikePython(delta, commandsCount) {
		t.Errorf("Wrong set address %d", brain.CurrentAddress)
	}
}

func TestBrain_SetAddress_SuccessWhenAddressGreatLength(t *testing.T) {
	commandsCount := 6
	delta := 7
	brain := Brain{}
	for i := 0; i < commandsCount; i++ {
		brain.Commands = append(brain.Commands, i)
	}
	brain.SetAddress(delta)
	if brain.CurrentAddress != modLikePython(delta, commandsCount) {
		t.Errorf("Wrong set address %d", brain.CurrentAddress)
	}
}

func TestBrain_MoveAddressOn_SuccessWhenDeltaGreatZero(t *testing.T) {
	commandsCount := 6
	delta := 7
	initialAddress := 2
	brain := Brain{
		CurrentAddress: initialAddress,
	}
	for i := 0; i < commandsCount; i++ {
		brain.Commands = append(brain.Commands, i)
	}
	brain.MoveAddressOn(delta)
	if brain.CurrentAddress != modLikePython(delta+initialAddress, commandsCount) {
		t.Errorf("Wrong set address %d", brain.CurrentAddress)
	}
}

func TestBrain_MoveAddressOn_SuccessWhenDeltaLessZero(t *testing.T) {
	commandsCount := 6
	delta := -4
	initialAddress := 2
	brain := Brain{
		CurrentAddress: initialAddress,
	}
	for i := 0; i < commandsCount; i++ {
		brain.Commands = append(brain.Commands, i)
	}
	brain.MoveAddressOn(delta)
	if brain.CurrentAddress != 4 {
		t.Errorf("Wrong set address %d", brain.CurrentAddress)
	}
}

func TestBrain_GetCurrentCommand_Success(t *testing.T) {
	commandsCount := 6
	initialAddress := 2
	brain := Brain{
		CommandList: CommandList{Commands: []Command{moveCommand}},
	}
	for i := 0; i < commandsCount; i++ {
		brain.Commands = append(brain.Commands, commandsCount-i-1)
	}
	brain.SetAddress(initialAddress)
	currentCommand := brain.GetCurrentCommand().Identifier
	if currentCommand != 0 {
		t.Errorf("Wrong set address %d", brain.CurrentAddress)
	}
}

func TestBrain_GetCommand_Success(t *testing.T) {
	commandsCount := 6
	initialAddress := 2
	brain := Brain{}
	for i := 0; i < commandsCount; i++ {
		brain.Commands = append(brain.Commands, i)
	}
	currentCommand := brain.GetCommandIdentifier(initialAddress - 4)
	if currentCommand != commandsCount-initialAddress {
		t.Errorf("Wrong set address %d", brain.CurrentAddress)
	}
}

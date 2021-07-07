package gonesis


type Command struct {
	code int
	isAbortable bool
	handler func(*Agent) bool
}

func CommandConstructor(code int, isAbortable bool, handler func(*Agent)bool) Command {
	command := Command{}
	command.code = code
	command.isAbortable = isAbortable
	command.handler = handler
	return command
}

package gonesis

type Command struct {
	IsInterrupts bool
	Handler      func(*World, *Agent)
}

package gonesis

type Command struct {
	Identifier   int
	IsInterrupts bool
	Handler      func(*World, *Agent)
}

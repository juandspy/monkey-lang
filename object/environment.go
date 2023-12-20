package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment is used to store the variables and bindings
type Environment struct {
	store map[string]Object
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

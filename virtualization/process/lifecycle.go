package process

// New creates a new process
func New(name string, program []byte) *Process {
	return &Process{
		ID:    generateNewProcessID(),
		State: StateNew,
		Name:  name,
	}
}

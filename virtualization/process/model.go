package process

type State uint8

const (

	// StateUnknown indicates the state of this process is unknown.
	//
	// It should be considered a fatal error for a Process to be found in this state
	StateUnknown State = iota

	// StateNew indicates that Process is in the process of being created
	StateNew

	// StateReady indicates that Process has been allocated all resources needed to run,
	// but the CPU is not currently working on its instructions
	StateReady

	// StateRunning indicates that Process instruction is currently being executed by the CPU
	StateRunning

	// StateWaiting indicates that Process execution is currently paused due to
	// - waiting for some resource to become available (e.g., CPU time, disk access, etc)
	// - context switch
	// - waiting for some events to occur (e.g., input device, inter-process messages, timer to go off)
	StateWaiting

	// StateTerminated indicates the program has stopped running either by completion
	// or by outright termination as an offending program
	StateTerminated
)

type Process struct {
	ID    uint32
	State State
	Name  string
}

type AddressSpace struct {
	InitialSize uint16
}

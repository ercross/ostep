package process

// Since processID need only to be unique for each system restart,
// then a simple ID generator such as this is sufficient
var lastProcessID uint32 = 0

func generateNewProcessID() uint32 {
	return lastProcessID + 1
}

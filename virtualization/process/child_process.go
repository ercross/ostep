package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

/*
	Fork current process using ForkExec

	System call fork() is not available in Go because fork()
	was not designed to be thread-safe.
	See https://stackoverflow.com/questions/28370646/how-do-i-fork-a-go-process/28371586#28371586
*/

func syscallForkProcess() {
	parentPid := os.Getpid()
	fmt.Printf("Parent PID: %d\n", parentPid)

	command := "/bin/ls"
	commandArguments := []string{"ls", "-l", "/"}
	childPid, err := syscall.ForkExec(command, commandArguments, nil)
	if err != nil {
		log.Fatalf("failed to fork process. Err: %v", err)
	}

	fmt.Printf("Child process PID: %d\n", childPid)
}

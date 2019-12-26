package softfiles

import (
	"fmt"
	"io"
	"os/exec"
)

func ForceStop(cmd *exec.Cmd) {
	fmt.Println("Killing Process...")
	err := cmd.Process.Kill()
	PanicIf(err)
	fmt.Println("Done!")
}
func ForceRestart(cmd *exec.Cmd) (io.Reader, io.WriteCloser, *exec.Cmd) {
	fmt.Println("Killing Process...")
	err := cmd.Process.Kill()
	PanicIf(err)
	fmt.Println("Starting Process!")
	//command
	outPipe, err := cmd.StdoutPipe()
	PanicIf(err)
	inPipe, err := cmd.StdinPipe()
	PanicIf(err)
	fmt.Println("Starting Server...")
	err = cmd.Start()
	PanicIf(err)
	fmt.Println("Server started!")
	return outPipe, inPipe, cmd
}

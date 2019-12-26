package softfiles

import (
	"fmt"
	"io"
	"os/exec"
	"time"
)

/*func cycle(normReader io.Reader, errReader io.Reader){

	for x := 0; x > 5; x++{
	}

}*/
func getIfServerTurnsOff(cmd *exec.Cmd) bool {
	done := make(chan bool)
	go func(chan bool) {
		fmt.Println("Waiting for server to turn off!")
		err := cmd.Wait()
		PanicIf(err)
		done <- true
	}(done)
	fmt.Println("Waiting 5 seconds for process to stop...")
	time.Sleep(5 * time.Second)
	if <-done == true {
		fmt.Println("Server has stopped")
		return true
	} else {
		fmt.Println("Server has failed to quit. Please run gsm force-restart")
		return false
	}
}
func Stop(closer io.WriteCloser, cmd *exec.Cmd) {
	_, err := closer.Write([]byte("stop"))
	PanicIf(err)
	getIfServerTurnsOff(cmd)
}
func Restart(closer io.WriteCloser, cmd *exec.Cmd) (io.Reader, io.WriteCloser, *exec.Cmd) {
	_, err := closer.Write([]byte("stop"))
	PanicIf(err)
	if getIfServerTurnsOff(cmd) {
		command := cmd
		outPipe, err := command.StdoutPipe()
		PanicIf(err)
		inPipe, err := command.StdinPipe()
		PanicIf(err)
		fmt.Println("Starting Server...")
		err = command.Start()
		PanicIf(err)
		fmt.Println("Server started!")
		return outPipe, inPipe, command
	}
	/*go func() {
		err = cmd.Wait()
		PanicIf(err)
	}()*/

}

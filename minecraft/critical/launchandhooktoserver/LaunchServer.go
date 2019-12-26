package launchandhooktoserver

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

/*func ticker(tick time.Duration) {
	ticker := time.NewTicker(tick)
	defer ticker.Stop()
	done := make(chan bool)
	sleep := 1 * time.Second
	go func() {
		time.Sleep(sleep)
		done <- true
	}()
	ticks := 0
	for {
		select {
		case <-done:
			fmt.Printf("%v × %v ticks in %v\n", ticks, tick, sleep)
			return
		case <-ticker.C:
			ticks++
		}
	}
}*/

//Log checks and logs a error
/*func Parallelize(functions ...func()) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))

	defer waitGroup.Wait()

	for _, function := range functions {
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(function)
	}
}*/
func PanicIf(err error) {
	if err != nil {
		log.Fatal(err)
		//fmt.Println("error")
		//panic(err)
	}
}

func CopyClose(toPipe io.WriteCloser, fromPipe io.Reader) {
	defer toPipe.Close()
	_, err := io.Copy(toPipe, fromPipe)
	PanicIf(err)
}

func ReadAndPrint(pipe io.Reader) {
	scanner := bufio.NewScanner(pipe)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
}
func LaunchAndGrabInput(serverStartScript string) {
	//var age int
	//serverStartScript = "/home/henry/Desktop/Server Test/MinecraftServerManagerStart.sh"
	//fmt.Printf("Server Jar is: " + serverStartScript)
	startCommand := exec.Command("/bin/sh", serverStartScript)
	cmdStdoutPipe, err := startCommand.StdoutPipe()
	PanicIf(err)
	cmdStdinPipe, err := startCommand.StdinPipe()
	PanicIf(err)
	go CopyClose(cmdStdinPipe, os.Stdin)
	go ReadAndPrint(cmdStdoutPipe)
	err = startCommand.Start()
	PanicIf(err)
	err = startCommand.Wait()
	PanicIf(err)
}

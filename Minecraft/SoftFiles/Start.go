package SoftFiles

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

//type Cmd exec.Cmd
func Start(serverJar string, serverDir string) (io.Reader, io.WriteCloser, *exec.Cmd) {
	command := getCmd(serverJar)
	command.Dir = serverDir
	outpipe, err := command.StdoutPipe()
	PanicIf(err)
	inpipe, err := command.StdinPipe()
	PanicIf(err)
	fmt.Println("Starting Server...")
	err = command.Start()
	PanicIf(err)
	fmt.Println("Server started!")
	return outpipe, inpipe, command
}
func getCmd(serverJar string) *exec.Cmd {
	return exec.Command("java", "-jar", serverJar)
}
func PanicIf(err error) {
	if err != nil {
		log.Fatal(err)
		//fmt.Println("error")
		//panic(err)
	}
}

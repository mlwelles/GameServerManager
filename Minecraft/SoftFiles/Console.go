package SoftFiles

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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
func Console(inPipe io.WriteCloser, outPipe io.Reader) {
	go CopyClose(inPipe, os.Stdin)
	go ReadAndPrint(outPipe)
}

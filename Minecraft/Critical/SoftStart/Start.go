package SoftStart

import (
	"log"
	"os/exec/"
)

//type Cmd exec.Cmd

func PanicIf(err error) {
	if err != nil {
		log.Fatal(err)
		//fmt.Println("error")
		//panic(err)
	}
}

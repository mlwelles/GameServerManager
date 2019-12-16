package GameServerManager

import "os"

func main() {
	//TODO: Develop Methods for Soft Starting and hooking into command
	if os.Args != nil {
		a := os.Args[1]
		if a == "start" {
			//soft start
		}
		if a == "console" {
			//hook into command
		}
		if a == "stop" {
			//send stop command to inpipe
		}
		if a == "restart" {
			//send stop command, then re-run
		}
		if a == "force-restart" {

		}
	}
}

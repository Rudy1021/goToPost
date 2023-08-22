package main

import (
	"fmt"
	"goToPost/runner"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print(`Usage: main <arguments> <ip> <CollectionName>
-t                  convert to thunder-client
-p                  convert to postman
-s                  convert to swagger`)
		return
	}

	switch os.Args[1] {
	case "-t":
		runner.UseThunder()
	case "-p":
		runner.UsePostman()
	case "-s":
	default:
		fmt.Print(`Usage: main <arguments> <ip> <CollectionName>
		-t                  convert to thunder-client
		-p                  convert to postman
		-s                  convert to swagger`)
		return
	}

}

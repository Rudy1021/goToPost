package main

import (
	"fmt"
	"goToPost/runner"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print(`
GoToPost gtp is a tool that generate router's url to json for Thunder Client or Postman.

Usage: 

        gtp <arguments> <ip> <CollectionName>

Arguments:

        -t        Convert to Thunder Client
        -p        Convert to Postman
Please visit https://github.com/Rudy1021/goToPost for more information.`)
		return
	}

	switch os.Args[1] {
	case "-t":
		runner.UseThunder()
	case "-p":
		runner.UsePostman()
	default:
		fmt.Print(`Usage: gtp <arguments> <ip> <CollectionName>
		-t                  convert to thunder-client
		-p                  convert to postman`)
		return
	}

}

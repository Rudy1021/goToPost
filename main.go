package main

import (
	"flag"
	"fmt"
	"goToPost/runner"
	"os"
)

func main() {
	thunderPtr := flag.Bool("t", false, "Convert to Thunder Client")
	postmanPtr := flag.Bool("p", false, "Convert to Postman")
	configPtr := flag.Bool("config", false, "Path to the configuration file")

	flag.Parse()

	// 如果提供了 -config 參數，則使用它
	baseUrl := os.Args[2]

	fileName := os.Args[3]
	if *configPtr {
		fmt.Println("Configuration file:", *configPtr)
		baseUrl = os.Args[3]

		fileName = os.Args[4]
		// 在這裡加入讀取和處理配置文件的代碼
	}

	if !*thunderPtr && !*postmanPtr {
		fmt.Print(`
GoToPost gtp is a tool that generate router's url to json for Thunder Client or Postman.

Usage: 

        gtp <arguments> <ip:port> <CollectionName>

Arguments:

        -t        Convert to Thunder Client
        -p        Convert to Postman
				-config   load configuration file
Please visit https://github.com/Rudy1021/goToPost for more information.`)
		return
	}

	if *thunderPtr {
		runner.UseThunder()
		return
	}

	if *postmanPtr {
		runner.UsePostman(baseUrl, fileName, *configPtr)
		return
	}
}

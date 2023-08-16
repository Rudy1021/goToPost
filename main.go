package main

import (
	"encoding/json"
	"fmt"
	"goToPost/models"
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print(`Usage: main <arguments> <ip> <CollectionName>
-t                  convert to thunder-client
-p                  convert to postman
-s                  convert to swagger`)
		return
	}

	exportJson := models.Thunder{}
	now := time.Now().UTC()
	baseUrl := os.Args[2]

	exportJson.Client = "Thunder Client"
	exportJson.CollectionName = os.Args[3]
	exportJson.DateExported = now
	exportJson.Version = "1.1"
	exportJson.Folders = []string{}

	switch os.Args[1] {
	case "-t":
	case "-p":
	case "-s":
	default:
		fmt.Print(`Usage: main <arguments> <ip> <CollectionName>
		-t                  convert to thunder-client
		-p                  convert to postman
		-s                  convert to swagger`)
		return
	}

	// 打開檔案，取得檔案指標
	filePath := "router.go"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("無法開啟檔案:", err)
		return
	}
	defer file.Close() // 確保在函式結束時關閉檔案

	// 讀取檔案內容
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("無法讀取檔案內容:", err)
		return
	}

	// 定義正則表達式
	//"/Attendee/Select"
	// reForNameAndActions := regexp.MustCompile(`\.([A-Z]+)\s*\("(/?[^"]+)",\s*\w+\)`)

	///"getFunctions"
	// reForActionsWithName := regexp.MustCompile(`\.([A-Z]+)\s*\("([^"]+)",\s*([^)]+)\)`)
	reForGroup := regexp.MustCompile(`(.+)\s:=\s\w+\.Group\(\"(.+)\"\)`)

	matchesForGroup := reForGroup.FindAllStringSubmatch(string(content), -1)
	// for _, match := range matches {
	// 	if len(match) == 3 {
	// 		httpMethod := match[1]
	// 		url := match[2]
	// 		fmt.Printf("HTTP Method: %s\n", httpMethod)
	// 		fmt.Printf("URL: %s\n", url)
	// 		// request := models.Request{}
	// 		// request.Method = httpMethod
	// 		// request.Url = baseUrl + url
	// 		// request.Name = url
	// 		// request.SortNum = 10000
	// 		// request.Created = now
	// 		// request.Modified = now
	// 		// request.Headers = []string{}
	// 		// request.Params = []string{}
	// 		// request.Tests = []string{}
	// 		// exportJson.Requests = append(exportJson.Requests, request)
	// 	}
	// }

	// for _, match := range matches {
	// 	if len(match) == 4 {
	// 		httpMethod := match[1]
	// 		url := match[2]
	// 		handler := match[3]
	// 		fmt.Printf("HTTP Method: %s\n", httpMethod)
	// 		fmt.Printf("URL: %s\n", url)
	// 		fmt.Printf("Handler: %s\n", handler)
	// 		request := models.Request{}
	// 		request.Method = httpMethod
	// 		request.Url = baseUrl + url
	// 		request.Name = handler
	// 		request.SortNum = 10000
	// 		request.Created = now
	// 		request.Modified = now
	// 		request.Headers = []string{}
	// 		request.Params = []string{}
	// 		request.Tests = []string{}
	// 		exportJson.Requests = append(exportJson.Requests, request)
	// 	}
	// }

	for _, group := range matchesForGroup {
		if len(group) == 3 {
			// groupName := group[1]
			groupRouter := group[2]
			reForGroupWithActionsAndName := regexp.MustCompile(group[1] + `\.([A-Z]+)\s*\("([^"]+)",\s*([^)]+)\)`)
			matchesWithGroup := reForGroupWithActionsAndName.FindAllStringSubmatch(string(content), -1)
			// fmt.Printf("Group Name: %s\n", groupName)
			// fmt.Printf("Group Router: %s\n", groupRouter)

			for _, route := range matchesWithGroup {
				if len(route) == 4 {
					httpMethods := route[1]
					apiRoutes := route[2]
					handler := route[3]
					if apiRoutes[0:1] != "/" {
						apiRoutes = "/" + apiRoutes
					}

					// fmt.Printf("Http Methods: %s\n", httpMethods)
					// fmt.Printf("Api Routes: %s\n", groupRouter+apiRoutes)
					request := models.Request{}
					request.Method = httpMethods
					request.Url = baseUrl + "/" + groupRouter + apiRoutes
					request.Name = handler
					request.SortNum = 10000
					request.Created = now
					request.Modified = now
					request.Headers = []string{}
					request.Params = []string{}
					request.Tests = []string{}
					exportJson.Requests = append(exportJson.Requests, request)
				}
			}
		}
	}

	jsonData, _ := json.Marshal(exportJson)

	// fmt.Println(string(jsonData))

	err = os.WriteFile("thunder-collection_"+exportJson.CollectionName+".json", jsonData, 0777)
	if err != nil {
		log.Fatal("Error writing JSON file:", err)
		return
	}

	fmt.Println("JSON data saved to output.json")
}

package runner

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

func UseThunder() {

	exportJson := models.Thunder{}
	now := time.Now().UTC()
	baseUrl := os.Args[2]

	exportJson.Client = "Thunder Client"
	exportJson.CollectionName = os.Args[3]
	exportJson.DateExported = now
	exportJson.Version = "1.1"
	exportJson.Folders = []string{}

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

	reForGroup := regexp.MustCompile(`(.+)\s:=\s\w+\.Group\(\"(.+)\"\)`)

	matchesForGroup := reForGroup.FindAllStringSubmatch(string(content), -1)

	if len(matchesForGroup) != 0 {
		for _, group := range matchesForGroup {
			if len(group) == 3 {
				groupRouter := group[2]
				reForGroupWithActionsAndName := regexp.MustCompile(group[1] + `\.([A-Z]+)\s*\("([^"]+)",\s*([^)]+)\)`)
				matchesWithGroup := reForGroupWithActionsAndName.FindAllStringSubmatch(string(content), -1)

				for _, route := range matchesWithGroup {
					if len(route) == 4 {
						httpMethods := route[1]
						apiRoutes := route[2]
						handler := route[3]
						if apiRoutes[0:1] != "/" {
							apiRoutes = "/" + apiRoutes
						}

						request := models.RequestOfThunder{}
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
	} else {
		///"getFunctions"
		reForActionsWithName := regexp.MustCompile(`\.([A-Z]+)\s*\("([^"]+)",\s*([^)]+)\)`)

		matchForActionsWithName := reForActionsWithName.FindAllStringSubmatch(string(content), -1)

		for _, routes := range matchForActionsWithName {
			if len(routes) == 4 {
				httpMethods := routes[1]
				url := routes[2]
				handler := routes[3]
				request := models.RequestOfThunder{}
				request.Method = httpMethods
				request.Url = baseUrl + url
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

	jsonData, _ := json.Marshal(exportJson)

	err = os.WriteFile("thunder-collection_"+exportJson.CollectionName+".json", jsonData, 0777)
	if err != nil {
		log.Fatal("Error writing JSON file:", err)
		return
	}

	fmt.Println("JSON data saved to thunder-collection_" + exportJson.CollectionName + ".json")
}

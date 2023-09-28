package runner

import (
	"encoding/json"
	"fmt"
	"goToPost/models"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func UsePostman() {

	exportJson := models.Postman{}
	baseUrl := os.Args[2]

	exportJson.Info.Name = os.Args[3]
	exportJson.Info.Schema = "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"

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

	reForActionsWithName := regexp.MustCompile(`\.([A-Z]+)\s*\("([^"]+)",\s*([^)]+)\)`)

	reForGroup := regexp.MustCompile(`(.+)\s:=\s\w+\.Group\(\"(.+)\"\)`)

	matchesForGroup := reForGroup.FindAllStringSubmatch(string(content), -1)

	urlRegex := regexp.MustCompile(`^(https?://)?([^:/]+)(:\d+)?`)

	urlMatch := urlRegex.FindStringSubmatch(baseUrl)

	host := ""
	protocol := ""
	port := ""

	if len(urlMatch) == 4 {
		protocol = urlMatch[1]
		host = urlMatch[2]
		port = urlMatch[3]
		if protocol == "" {
			protocol = "http://"
		}
		if port != "" {
			port = port[1:] // Remove the leading ":"
		}
	} else {
		fmt.Println("Invalid URL format:", baseUrl)
	}

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

						fullURL := baseUrl + groupRouter + apiRoutes
						path := strings.Split((groupRouter + apiRoutes), "/")

						postmanItem := models.Item{}
						postmanItem.Name = handler
						postmanItem.Request.Method = httpMethods
						postmanItem.Request.Headers = []string{}
						postmanItem.Response = []string{}
						postmanItem.Request.Url.Raw = fullURL
						postmanItem.Request.Url.Protocol = protocol
						postmanItem.Request.Url.Host = strings.Split(host, ".")
						postmanItem.Request.Url.Path = path
						postmanItem.Request.Url.Port = port

						exportJson.Item = append(exportJson.Item, postmanItem)

					}
				}
			}
		}
	} else {

		matchForActionsWithName := reForActionsWithName.FindAllStringSubmatch(string(content), -1)

		for _, routes := range matchForActionsWithName {
			if len(routes) == 4 {
				httpMethods := routes[1]
				apiRoutes := routes[2]
				handler := routes[3]

				if apiRoutes[0:1] != "/" {
					apiRoutes = "/" + apiRoutes
				}

				fullURL := baseUrl + apiRoutes
				path := strings.Split((apiRoutes), "/")

				postmanItem := models.Item{}
				postmanItem.Name = handler
				postmanItem.Request.Method = httpMethods
				postmanItem.Request.Headers = []string{}
				postmanItem.Response = []string{}
				postmanItem.Request.Url.Raw = fullURL
				postmanItem.Request.Url.Protocol = protocol
				postmanItem.Request.Url.Host = strings.Split(host, ".")
				postmanItem.Request.Url.Path = path
				postmanItem.Request.Url.Port = port

				exportJson.Item = append(exportJson.Item, postmanItem)

			}
		}
	}

	jsonData, _ := json.Marshal(exportJson)

	err = os.WriteFile("postman-collection_"+exportJson.Info.Name+".json", jsonData, 0777)
	if err != nil {
		log.Fatal("Error writing JSON file:", err)
		return
	}

	fmt.Println("JSON data saved to postman-collection_" + exportJson.Info.Name + ".json")
}

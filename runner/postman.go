package runner

import (
	"encoding/json"
	"fmt"
	"goToPost/loader"
	"goToPost/models"
	"log"
	"os"
	"regexp"
	"strings"
)

func UsePostman(baseUrl, fileName string, useConfigFile bool) {

	exportJson := models.Postman{}

	exportJson.Info.Name = fileName
	exportJson.Info.Schema = "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"

	// 讀取路由文件
	content, error := loader.LoadRouter()

	header := []models.Headers{}

	if error != nil {
		fmt.Println("Error loading router file:", error)
		return
	}
	if useConfigFile {
		config, err := loader.LoadConfig()

		if err != nil {
			fmt.Println("Error loading config file:", err)
			return
		}

		for _, cookie := range config {
			tempHeader := models.Headers{}
			tempHeader.Key = "Cookie"
			tempHeader.Value = fmt.Sprintf("%s=%s", cookie.Name, cookie.Value)
			tempHeader.Type = "text"
			header = append(header, tempHeader)
		}
	}
	// 定義正則表達式
	//"/Attendee/Select"
	// reForNameAndActions := regexp.MustCompile(`\.([A-Z]+)\s*\("(/?[^"]+)",\s*\w+\)`)

	reForGroup := regexp.MustCompile(`(.+)\s:=\s\w+\.Group\(\"(.+)\"\)`)

	groupMatches := reForGroup.FindAllStringSubmatch(string(content), -1)

	urlRegex := regexp.MustCompile(`^(https?://)?([^:/]+)(:\d+)?`)

	urlMatches := urlRegex.FindStringSubmatch(baseUrl)

	host := ""
	protocol := ""
	port := ""

	if len(urlMatches) == 4 {
		protocol = urlMatches[1]
		host = urlMatches[2]
		port = urlMatches[3]
		if protocol == "" {
			protocol = "http"
		}
		if port != "" {
			port = port[1:] // Remove the leading ":"
		}
	} else {
		fmt.Println("Invalid URL format:", baseUrl)
	}

	if len(groupMatches) != 0 {
		for _, group := range groupMatches {
			if len(group) == 3 {
				groupRouter := group[2]
				reForGroupWithActionsAndName := regexp.MustCompile(group[1] + `\.([A-Z]+)\s*\("([^"]+)",\s*([^)]+)\)`)
				groupMatchesWithActionsAndName := reForGroupWithActionsAndName.FindAllStringSubmatch(string(content), -1)

				for _, route := range groupMatchesWithActionsAndName {
					if len(route) == 4 {
						httpMethods := route[1]
						apiRoutes := route[2]
						handler := route[3]
						path := strings.Split((groupRouter + apiRoutes), "/")
						if apiRoutes[0:1] != "/" {
							apiRoutes = "/" + apiRoutes
						}

						fullURL := protocol + "://" + baseUrl + "/" + groupRouter + apiRoutes

						postmanItem := models.Item{}
						postmanItem.Name = handler
						postmanItem.Request.Method = httpMethods
						postmanItem.Request.Headers = header
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

		reForActionsWithName := regexp.MustCompile(`\.([A-Z]+)\s*\("([^"]+)",\s*([^)]+)\)`)

		matchForActionsWithName := reForActionsWithName.FindAllStringSubmatch(string(content), -1)

		for _, routes := range matchForActionsWithName {
			if len(routes) == 4 {
				httpMethods := routes[1]
				apiRoutes := routes[2]
				handler := routes[3]
				path := strings.Split((apiRoutes), "/")
				if apiRoutes[0:1] != "/" {
					apiRoutes = "/" + apiRoutes
				}
				fullURL := protocol + "://" + baseUrl + apiRoutes

				postmanItem := models.Item{}
				postmanItem.Name = handler
				postmanItem.Request.Method = httpMethods
				postmanItem.Request.Headers = header
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

	err := os.WriteFile(exportJson.Info.Name+".postman-collection"+".json", jsonData, 0777)
	if err != nil {
		log.Fatal("Error writing JSON file:", err)
		return
	}

	fmt.Println("JSON data saved to " + exportJson.Info.Name + ".postman-collection" + ".json")
}

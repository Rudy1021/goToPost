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

	exportJson := models.Thunder{}
	now := time.Now().UTC()
	baseUrl := "127.0.0.1:5487"

	exportJson.Client = "Thunder Client"
	exportJson.CollectionName = "test"
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
	re := regexp.MustCompile(`router\.([A-Z]+)\s*\("(/?[^"]+)",\s*\w+\)`)

	matches := re.FindAllStringSubmatch(string(content), -1)
	for _, match := range matches {
		if len(match) == 3 {
			httpMethod := match[1]
			url := match[2]
			// fmt.Printf("HTTP Method: %s\n", httpMethod)
			// fmt.Printf("URL: %s\n", url)
			request := models.Request{}
			request.Method = httpMethod
			request.Url = baseUrl + url
			request.Name = url
			request.SortNum = 10000
			request.Created = now
			request.Modified = now
			request.Headers = []string{}
			request.Params = []string{}
			request.Tests = []string{}
			exportJson.Requests = append(exportJson.Requests, request)
		}
	}

	jsonData, _ := json.Marshal(exportJson)

	// fmt.Println(string(jsonData))

	err = os.WriteFile("output.json", jsonData, 0777)
	if err != nil {
		log.Fatal("Error writing JSON file:", err)
		return
	}

	fmt.Println("JSON data saved to output.json")
}

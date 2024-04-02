package loader

import (
	"fmt"
	"io"
	"os"
)

func LoadRouter() ([]byte, error) {
	// 在這裡加入讀取路由的代碼
	filePath := "router.go"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("無法開啟檔案:", err)
		return []byte{}, err
	}
	defer file.Close() // 確保在函式結束時關閉檔案

	// 讀取檔案內容
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("無法讀取檔案內容:", err)
		return []byte{}, err
	}
	return content, nil
}

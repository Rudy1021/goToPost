package loader

import (
	"fmt"
	"goToPost/models"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfig() ([]models.Cookie, error) {
	// Read the YAML file
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println("無法讀取檔案內容:", err)
		return nil, err
	}

	// Unmarshal the YAML data into a struct
	config := models.Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("無法解析 YAML 檔案:", err)
		return nil, err
	}

	return config.Cookies, nil
}

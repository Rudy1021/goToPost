package models

import "time"

type Thunder struct {
	Client         string             `json:"client"`
	CollectionName string             `json:"collectionName"`
	DateExported   time.Time          `json:"dateExported"`
	Version        string             `json:"version"`
	Folders        []string           `json:"folders"`
	Requests       []RequestOfThunder `json:"requests"`
}

type RequestOfThunder struct {
	ContainerId string    `json:"containerId"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	Method      string    `json:"method"`
	SortNum     int       `json:"sortNum"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
	Headers     []string  `json:"headers"`
	Params      []string  `json:"params"`
	Tests       []string  `json:"tests"`
}

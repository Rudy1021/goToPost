package models

import "time"

type Thunder struct {
	Client         string    `json:"client"`
	CollectionName string    `json:"collectionName"`
	DateExported   time.Time `json:"dateExported"`
	Version        string    `json:"version"`
	Folders        []string  `json:"folders"`
	Requests       []Request `json:"requests"`
}

type Request struct {
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

/*
  "requests": [
    {
      "containerId": "",
      "name": "getApsOutWoSelect",
      "url": "127.0.0.1:5487/api/getApsOutWoSelect",
      "method": "GET",
      "sortNum": 10000,
      "created": "2023-04-14T01:45:32.211Z",
      "modified": "2023-04-14T01:45:32.211Z",
      "headers": [
        {
          "name": "Cookie",
          "value": "plantID=JU"
        }
      ],
      "params": [],
      "tests": []
    },
*/

package models

type Postman struct {
	Info info   `json:"info"`
	Item []Item `json:"item"`
}

type info struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

type Item struct {
	Name     string           `json:"name"`
	Request  RequestOfPostman `json:"request"`
	Response []string         `json:"response"`
}

type RequestOfPostman struct {
	Method  string    `json:"method"`
	Headers []Headers `json:"headers"`
	Url     url       `json:"url"`
}

type url struct {
	Raw      string   `json:"raw"`
	Protocol string   `json:"protocol"`
	Host     []string `json:"host"`
	Path     []string `json:"path"`
	Port     string   `json:"port"`
}

type Headers struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

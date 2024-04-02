package models

type Cookie struct {
	Name   string `yaml:"name"`
	Value  string `yaml:"value"`
	Global bool   `yaml:"global"`
	Path   string `yaml:"path"`
}
type Config struct {
	Cookies []Cookie `yaml:"cookies"`
}

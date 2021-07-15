package configs

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

type AppConfig struct {
	Servers []Server
}
type Server struct {
	Hostname string
	Name string
}

func(config *AppConfig) parse(data []byte) error {
	return yaml.Unmarshal(data, config)
}
func LoadConfig() AppConfig {
	filename, _ := filepath.Abs("../system-design/yaml/app.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var appconfig AppConfig
	if err := appconfig.parse(yamlFile); err != nil {
		log.Fatal(err)
	}
	return appconfig
}
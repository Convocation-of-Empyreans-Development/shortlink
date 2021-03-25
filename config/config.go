package config

import (
	"encoding/json"
	_ "io"
	"io/ioutil"
	"os"
)

type RedirectionConfig struct {
	Mapping map[string]string `json:"mapping"`
}

func LoadConfig(filename string) (config RedirectionConfig) {
	configFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	configJson, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configJson, &config)
	if err != nil {
		panic(err)
	}
	return config
}

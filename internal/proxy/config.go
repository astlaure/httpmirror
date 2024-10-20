package proxy

import (
	"encoding/json"
	"io"
	"os"
)

type (
	ServiceTarget struct {
		Active  string `json:"active"`
		Preview string `json:"preview"`
	}

	Service struct {
		Name   string        `json:"name"`
		Path   string        `json:"path"`
		Target ServiceTarget `json:"target"`
	}

	Config struct {
		Services []Service `json:"services"`
	}
)

var Configuration = LoadConfig()

func LoadConfig() Config {
	configFile, err := os.Open("resources/proxy.config.json")

	if err != nil {
		panic(err)
	}

	defer configFile.Close()

	byteData, err := io.ReadAll(configFile)
	var proxyConfig Config

	json.Unmarshal(byteData, &proxyConfig)

	return proxyConfig
}

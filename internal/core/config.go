package core

import (
	"encoding/json"
	"io"
	"os"
)

type (
	AppConfig struct {
		Frontend string `json:"frontend"`
	}
)

var Config AppConfig

func ReadConfigFile() {
	configFile, err := os.Open("resources/app.config.json")

	if err != nil {
		panic(err)
	}

	defer configFile.Close()

	byteData, err := io.ReadAll(configFile)

	json.Unmarshal(byteData, &Config)
}

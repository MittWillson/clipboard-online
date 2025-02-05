package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

const ConfigFile = "config.json"
const LogFile = "log.txt"

type Config struct {
	Port     string       `json:"port"`
	LogLevel logrus.Level `json:"logLevel"`
}

var DefaultConfig = Config{
	Port:     "8086",
	LogLevel: log.WarnLevel,
}

func loadConfig(path string) (*Config, error) {
	if isExistFile(path) {
		return loadConfigFromFile(path)
	}
	if err := createConfigFile(path); err != nil {
		return nil, err
	}
	return &DefaultConfig, nil
}

func loadConfigFromFile(path string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func createConfigFile(path string) error {
	defaultConfigJSON, err := json.MarshalIndent(DefaultConfig, "", "  ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(path, []byte(defaultConfigJSON), 0644); err != nil {
		return err
	}
	return nil
}

func isExistFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

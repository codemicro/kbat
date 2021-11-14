package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

const configFileName = "settings.json"

func getConfigDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	if currentUser.HomeDir == "" {
		return "", errors.New("user has no home directory")
	}

	configDir := filepath.Join(currentUser.HomeDir, ".config", "kbat")

	if err := os.MkdirAll(configDir, 0777); err != nil && !errors.Is(err, os.ErrExist) {
		return "", err
	}

	return configDir, nil
}

type Config struct {
	Editor             string
	RepositoryLocation string
}

func (c *Config) Save() error {

	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	configFile := filepath.Join(configDir, configFileName)
	return ioutil.WriteFile(configFile, jsonData, 0644)
}

func LoadConfig() (*Config, error) {

	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	configFile := filepath.Join(configDir, configFileName)

	fcont, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	conf := new(Config)
	if err := json.Unmarshal(fcont, conf); err != nil {
		return nil, err
	}

	return conf, nil
}

package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	token    string `yaml:"token"`
	chatID   int64  `yaml:"chatid"`
	host     string `yaml:"host"`
	username string `yaml:"username"`
	password string `yaml:"password"`
}

func ReadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	config := Config{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

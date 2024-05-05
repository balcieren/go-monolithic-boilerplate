package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Common struct {
	API struct {
		PORT string `json:"port" yaml:"port"`
	} `json:"api" yaml:"api"`
}

func NewCommon() (*Common, error) {
	common := Common{}

	filename, err := filepath.Abs("./config.yaml")
	if err != nil {
		return nil, err
	}

	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(yamlFile, &common); err != nil {
		return nil, err
	}

	return &common, nil
}

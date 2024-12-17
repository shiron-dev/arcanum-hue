package converter

import (
	"os"

	"github.com/shiron-dev/arcanum-hue/internal/converter/model"
	"gopkg.in/yaml.v3"
)

func LoadConfigPath(path string) (*model.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return loadConfig(data)
}

func loadConfig(data []byte) (*model.Config, error) {
	var config model.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

package converter

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ThemeType string

const (
	ThemeTypeLight ThemeType = "light"
	ThemeTypeDark  ThemeType = "dark"
)

type Config struct {
	Name   string       `yaml:"name"`
	Colors []ColorTheme `yaml:"colorThemes"`
}

type ColorTheme struct {
	Type  ThemeType   `yaml:"type"`
	Model ColorsModel `yaml:"colors"`
}

type ColorsModel struct {
	Foreground              string `yaml:"foreground"`              // #D7D7D7
	SecondaryForeground     string `yaml:"secondaryForeground"`     // #9D9D9D
	Background              string `yaml:"background"`              // #181818
	SecondaryBackground     string `yaml:"secondaryBackground"`     // #313131
	BackgroundHighlight     string `yaml:"backgroundHighlight"`     // #0078D4
	PeekHighlightBackground string `yaml:"peekHighlightBackground"` // #BB800966
	BorderForeground        string `yaml:"borderForeground"`        // #2B2B2B
	BorderBackground        string `yaml:"borderBackground"`        // #3C3C3C
	EditorForeground        string `yaml:"editorForeground"`        // #CCCCCC
	EditorBackground        string `yaml:"editorBackground"`        // #1F1F1F
	DeleteForeground        string `yaml:"deleteForeground"`        // #F85149
	Border2                 string `yaml:"border2"`                 // #616161
	TextLinkForeground      string `yaml:"textLinkForeground"`      // #4DA3FF
}

func LoadConfigPath(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return loadConfig(data)
}

func loadConfig(data []byte) (*Config, error) {
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

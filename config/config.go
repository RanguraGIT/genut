package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type directories_config struct {
	Directories struct {
		Skip    []string `yaml:"skip"`
		Process []string `yaml:"process"`
		Others  bool     `yaml:"others,omitempty"`
	} `yaml:"directories"`
}

var (
	con     directories_config
	skip    []string
	process []string
	others  bool
)

func NewDirectoriesConfig(path string) *directories_config {
	configFile, err := os.Open(filepath.Join(path, ".genut.yml"))
	if err != nil {
		return &directories_config{
			Directories: struct {
				Skip    []string "yaml:\"skip\""
				Process []string "yaml:\"process\""
				Others  bool     "yaml:\"others,omitempty\""
			}{
				Skip:    []string{"vendor", "mocks"},
				Process: []string{"service", "repository", "usecase", "repository", "pkg", "infrastructure"},
				Others:  false,
			},
		}
	}

	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&con)
	if err != nil {
		fmt.Println("Error decode config:", err)
		return nil
	}

	return &directories_config{
		Directories: con.Directories,
	}
}

// function to check and load the config file
func (c *directories_config) LoadConfig() (skip []string, process []string, others bool) {
	others = false
	if c.Directories.Others {
		others = c.Directories.Others
	}

	return c.Directories.Skip, c.Directories.Process, others
}

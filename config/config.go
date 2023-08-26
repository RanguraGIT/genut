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

// function to check and load the config file
func (c *directories_config) LoadMockConfig() (skip []string, process []string, others bool) {
	skipped := make([]string, 0, len(c.Directories.Skip))
	for _, s := range c.Directories.Skip {
		if s != "mocks" {
			skipped = append(skipped, s)
		}
	}
	c.Directories.Skip = skipped

	c.Directories.Process = []string{"mocks"}

	others = false
	if c.Directories.Others {
		others = c.Directories.Others
	}

	return c.Directories.Skip, c.Directories.Process, others
}

// function to get the skip directories
func (c *directories_config) GetSkip() []string {
	return c.Directories.Skip
}

// function to get the process directories
func (c *directories_config) GetProcess() []string {
	return c.Directories.Process
}

// function to get the others directories
func (c *directories_config) GetOthers() bool {
	return c.Directories.Others
}

package helper

import (
	"path/filepath"
	"strings"
)

type contains_config struct {
	slice []string
	str   string
}

func NewContainsConfig(slice []string, str string) *contains_config {
	return &contains_config{
		slice: slice,
		str:   str,
	}
}

// function to check if the file contains in skip or process directories
func (c *contains_config) Contains() bool {
	for _, s := range c.slice {
		if s == c.str {
			return true
		}
	}
	return false
}

// function to check if the file contains in skip or process directories
func (c *contains_config) ContainsAny() bool {
	dir := strings.Split(c.str, string(filepath.Separator))
	for _, s := range c.slice {
		for _, d := range dir {
			if d == s {
				return true
			}
		}
	}
	return false
}

// function to get directory name
func (c *contains_config) GetDir() string {
	components := strings.Split(c.str, string(filepath.Separator))
	for _, dir := range c.slice {
		for _, component := range components {
			if component == dir {
				return dir
			}
		}
	}
	return "others"
}

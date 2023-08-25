package helper

import (
	"os"
	"strings"
)

type filename_config struct {
	origin string
	pkg    string
}

var mockName string

func NewFilenameConfig(origin string, pkg string) *filename_config {
	return &filename_config{
		origin: origin,
		pkg:    pkg,
	}
}

// function to generate mock filename
func (c *filename_config) Generate() string {
	fileDir := strings.SplitAfter(c.origin, c.pkg)

	if len(fileDir) < 2 {
		mockName = fileDir[0]
	} else {
		mockName = fileDir[1]
	}

	file := strings.ReplaceAll(mockName, string(os.PathSeparator), "_")
	return file[1:]
}

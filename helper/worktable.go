package helper

import "os"

// function to get current working directory
func Worktable() (string, error) {
	cwd, err := os.Getwd()
	return cwd, err
}

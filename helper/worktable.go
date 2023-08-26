package helper

import "os"

func Worktable() (string, error) {
	cwd, err := os.Getwd()
	return cwd, err
}

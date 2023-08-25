package checker

import (
	"fmt"
	"os/exec"
)

// function to check if mockgen is installed
func Mockgen() bool {
	cmd := exec.Command("mockgen", "-version")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

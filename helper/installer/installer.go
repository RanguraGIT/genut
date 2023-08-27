package installer

import (
	"fmt"
	"os/exec"
	"strings"
)

// function to install mockgen
func Mockgen() bool {
	cmd := exec.Command("go", "install", "github.com/golang/mock/mockgen")
	out, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(out), "malformed module path") {
			fmt.Println("go install malformed module path.")
		}
		fmt.Println("installing mockgen, You may need to install manually!")
		return false
	}
	return true
}

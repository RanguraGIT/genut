package main

import (
	"flag"
	"fmt"

	"github.com/RanguraGIT/genut/genut"
)

// Sub command list
func main() {
	genFlag := flag.Bool("mocking", false, "The command can be used to generate a new project")
	flag.Parse()

	if *genFlag {
		genut.GenMockgen()
	} else {
		fmt.Println("genut")
	}
}

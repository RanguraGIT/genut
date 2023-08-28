package main

import (
	"fmt"
	"os"

	mocks "github.com/RanguraGIT/genut/genut/mocks"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "genut",
		Short: "Genut is a tool to ease DG project dev sec ops",
	}

	var generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate code",
		Run: func(cmd *cobra.Command, args []string) {
			mock, _ := cmd.Flags().GetBool("mocks")

			if mock {
				mocks.GenMockgen()
			}

			if !mock {
				fmt.Println("No actions selected. Use --mocks.")
			}
		},
	}

	// Add flags to the generate command
	generateCmd.Flags().Bool("mocks", false, "Generate mockgen")

	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/RanguraGIT/genut/genut"
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
			mock, _ := cmd.Flags().GetBool("mocking")
			wrap, _ := cmd.Flags().GetBool("wrapper")

			if mock {
				genut.GenMockgen()
			}

			if wrap {
				genut.GenWrapper()
			}

			if !mock && !wrap {
				fmt.Println("No actions selected. Use --mocking and/or --wrapper.")
			}
		},
	}

	// Add flags to the generate command
	generateCmd.Flags().Bool("mocking", false, "Generate mockgen")
	generateCmd.Flags().Bool("wrapper", false, "Generate wrapper")

	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

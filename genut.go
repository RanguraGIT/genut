package main

import (
	"fmt"
	"os"

	mocks "github.com/RanguraGIT/genut/genut/mocks"
	"github.com/RanguraGIT/genut/helper"
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
			conf, _ := cmd.Flags().GetBool("config")

			if mock {
				mocks.GenMockgen()
			}

			if conf {
				helper.GenConfig()
			}

			if !mock && !conf {
				fmt.Println("No actions selected. Use --mocks and/or --config.")
			}
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of MyApp",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Version 1.0")
		},
	}

	// Add flags to the generate command
	generateCmd.Flags().Bool("mocks", false, "Generate mockgen")
	generateCmd.Flags().Bool("config", false, "Generate config file")

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

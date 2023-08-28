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
		RunE: func(cmd *cobra.Command, args []string) error {
			configFlag, _ := cmd.Flags().GetBool("config")
			versionFlag, _ := cmd.Flags().GetBool("version")

			if versionFlag {
				fmt.Println("Genut v1.0.0")
			} else if configFlag {
				helper.GenConfig()
			} else {
				help()
			}

			return nil
		},
	}

	rootCmd.PersistentFlags().BoolP("config", "c", false, "Generate config file to root project")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version of Genut")

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "tCreate new project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cooming soon")
		},
	}

	rootCmd.AddCommand(createCmd)

	var installCmd = &cobra.Command{
		Use:   "install",
		Short: "Install new service",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cooming soon")
		},
	}

	rootCmd.AddCommand(installCmd)

	var mockCmd = &cobra.Command{
		Use:   "mocks",
		Short: "Generate code",
		Run: func(cmd *cobra.Command, args []string) {
			mocks.GenMockgen()
		},
	}

	rootCmd.AddCommand(mockCmd)

	var preCommitCmd = &cobra.Command{
		Use:   "pre-commit",
		Short: "Installing pre-commit configuration",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cooming soon")
		},
	}

	rootCmd.AddCommand(preCommitCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func help() {
	fmt.Println()
	fmt.Println("Description:")
	fmt.Println("  Genut is a tool to ease DG project devsecops")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  genut [command / flags]")
	fmt.Println()
	fmt.Println("Available Flags:")
	fmt.Println("  --config   -c\t\tGenerate config file to root project")
	fmt.Println("  --version  -v\t\tPrint the version number of Genut")
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println("  create  [project]\tCreate new project")
	fmt.Println("  install [service]\tCreate new service")
	fmt.Println("  mocks\t\t\tGenerate mocks")
	fmt.Println("  pre-commit\t\tInstalling pre-commit configuration")
	fmt.Println()
}

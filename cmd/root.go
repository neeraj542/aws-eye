package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aws-eye",
	Short: "A lightweight, interactive AWS EC2 utility",
	Long: `aws-eye is a simple, interactive AWS utility that fetches EC2 instance 
details in a clean, readable format. It provides both interactive and 
flag-based modes for querying EC2 instances.`,
	Version: "1.0.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate("{{.Version}}\n")
}


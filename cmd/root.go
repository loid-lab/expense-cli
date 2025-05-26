package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tracker",
	Short: "A lightweight CLI to add, update and delete expenses",
	Long:  "Little Expense Tracker is a simple CLI tool for managing your personal expenses.",
}

// Execute adds all child commands and runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

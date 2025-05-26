package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an Expense",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Expense updated")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

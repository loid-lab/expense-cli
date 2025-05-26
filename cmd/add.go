package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new Expense",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Expense added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
